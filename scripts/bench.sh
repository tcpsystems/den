#!/usr/bin/env bash
# Benchmarks the denctl /api/state endpoint, catalog.LoadSkills, and
# console.Clusters scan time. Starts its own denctl on an ephemeral port
# if nothing is already listening on 127.0.0.1:8793.
set -euo pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "$REPO_ROOT"

DENCTL_PID=""
LOADSKILLS_DIR="$REPO_ROOT/scripts/.bench-loadskills"
CLUSTERS_DIR="$REPO_ROOT/scripts/.bench-clusters"

cleanup() {
	if [[ -n "$DENCTL_PID" ]] && kill -0 "$DENCTL_PID" 2>/dev/null; then
		kill "$DENCTL_PID" 2>/dev/null || true
		wait "$DENCTL_PID" 2>/dev/null || true
	fi
	rm -rf "$LOADSKILLS_DIR" "$CLUSTERS_DIR"
}
trap cleanup EXIT

# ---- build denctl ----
go build -o /tmp/denctl-k ./cmd/denctl

# ---- pick a target: reuse a live denctl on 8793, else start our own ----
DEFAULT_ADDR="127.0.0.1:8793"
code="$(curl -s -o /dev/null -w '%{http_code}' --max-time 1 "http://$DEFAULT_ADDR/" || true)"
if [[ "$code" == "200" ]]; then
	TARGET_ADDR="$DEFAULT_ADDR"
else
	PORT=18793
	while ss -ltn "( sport = :$PORT )" 2>/dev/null | grep -q LISTEN; do
		PORT=$((PORT + 1))
	done
	TARGET_ADDR="127.0.0.1:$PORT"
	/tmp/denctl-k -repo "$REPO_ROOT" -addr "$TARGET_ADDR" >/tmp/denctl-k.log 2>&1 &
	DENCTL_PID=$!
	up=0
	for _ in $(seq 1 50); do
		code="$(curl -s -o /dev/null -w '%{http_code}' --max-time 1 "http://$TARGET_ADDR/" || true)"
		if [[ "$code" == "200" ]]; then
			up=1
			break
		fi
		sleep 0.1
	done
	if [[ "$up" -ne 1 ]]; then
		echo "denctl did not come up on $TARGET_ADDR" >&2
		exit 1
	fi
fi

# ---- 50 sequential /api/state requests, collect time_total in seconds ----
mkdir -p /tmp/harness-wave
SAMPLES_FILE="$(mktemp)"
for _ in $(seq 1 50); do
	curl -s -o /dev/null -w '%{time_total}\n' "http://$TARGET_ADDR/api/state" >>"$SAMPLES_FILE"
done

# p50/p95 on the nearest-rank method: rank = ceil(pct * N), 1-indexed after
# a numeric ascending sort.
read -r STATE_P50_MS STATE_P95_MS <<<"$(python3 - "$SAMPLES_FILE" <<'PY'
import sys
vals = sorted(float(l) * 1000 for l in open(sys.argv[1]) if l.strip())
n = len(vals)
def pct(p):
    rank = max(1, -(-int(p * n * 100) // 100))  # ceil(p*n)
    return vals[rank - 1]
print(f"{pct(0.50):.3f} {pct(0.95):.3f}")
PY
)"
rm -f "$SAMPLES_FILE"

# ---- catalog.LoadSkills timing, via a throwaway helper inside the module ----
LOAD_SKILLS_MS="null"
mkdir -p "$LOADSKILLS_DIR"
cat >"$LOADSKILLS_DIR/main.go" <<'GOEOF'
package main

import (
	"fmt"
	"time"

	"github.com/tcpsystems/den/internal/catalog"
)

func main() {
	idx := catalog.Builtin()
	start := time.Now()
	if err := catalog.LoadSkills(&idx, "."); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%.3f\n", float64(time.Since(start).Microseconds())/1000.0)
}
GOEOF
if out="$(go run "$LOADSKILLS_DIR" 2>/tmp/bench-loadskills.err)"; then
	if [[ "$out" =~ ^[0-9.]+$ ]]; then
		LOAD_SKILLS_MS="$out"
	fi
fi

# ---- console.Clusters scan timing over /tmp/harness-wave ----
CLUSTERS_SCAN_MS="null"
mkdir -p "$CLUSTERS_DIR"
cat >"$CLUSTERS_DIR/main.go" <<'GOEOF'
package main

import (
	"fmt"
	"time"

	"github.com/tcpsystems/den/internal/console"
)

func main() {
	start := time.Now()
	console.Clusters("/tmp/harness-wave")
	fmt.Printf("%.3f\n", float64(time.Since(start).Microseconds())/1000.0)
}
GOEOF
if out="$(go run "$CLUSTERS_DIR" 2>/tmp/bench-clusters.err)"; then
	if [[ "$out" =~ ^[0-9.]+$ ]]; then
		CLUSTERS_SCAN_MS="$out"
	fi
fi

TS="$(date +%s)"
JSON="$(python3 -c "
import json
print(json.dumps({
    'timestamp': $TS,
    'state_p50_ms': $STATE_P50_MS,
    'state_p95_ms': $STATE_P95_MS,
    'state_samples': 50,
    'load_skills_ms': $LOAD_SKILLS_MS,
    'clusters_scan_ms': $CLUSTERS_SCAN_MS,
    'target_addr': '$TARGET_ADDR',
}))
")"

echo "$JSON"
mkdir -p /tmp/harness-wave
echo "$JSON" >"/tmp/harness-wave/bench-$TS.json"
