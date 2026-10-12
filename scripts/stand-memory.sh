#!/usr/bin/env bash
# Stand up the memory path den already uses.
#
# 1. llama-embed on 127.0.0.1:8091 (Qwen3-Embedding-8B Q6_K, RX 6600).
# 2. OpenViking on 127.0.0.1:1933, from the existing compose file.
# 3. mcp-hub on 127.0.0.1:37373, if that user unit exists.
#
# This script does not install a second memory product. It does not print keys.
# It refuses to start the embedder while the display pin is still the RX 6600.
set -euo pipefail

status_file=/run/gpu-display/status
embed_unit=llama-embed.service
compose=/home/kivox/.config/deerflow/openviking-compose.yaml
project=openviking

die() { printf 'stand-memory: %s\n' "$1" >&2; exit 1; }

gpu_pin() {
  if [[ ! -r "$status_file" ]]; then
    printf 'unknown'
    return
  fi
  # The status file is a short key=value list written by gpu-display-select.
  awk -F= '$1=="device" || $1=="mode" { print $1"="$2 }' "$status_file" | tr '\n' ' '
  printf '\n'
}

if gpu_pin | grep -q 'amd-card'; then
  die "display pin is still the RX 6600. Move the cable to the 1050 Ti before starting the embedder."
fi

if ! systemctl --user is-active --quiet "$embed_unit"; then
  systemctl --user start "$embed_unit"
fi

ready=0
for _ in $(seq 1 40); do
  if curl -fsS --max-time 2 http://127.0.0.1:8091/health >/dev/null 2>&1; then
    ready=1
    break
  fi
  sleep 1
done
[[ "$ready" == 1 ]] || die "embedder did not answer on 127.0.0.1:8091"

if [[ ! -f "$compose" ]]; then
  die "missing $compose"
fi
docker compose -p "$project" -f "$compose" up -d

if systemctl --user cat mcp-hub.service >/dev/null 2>&1; then
  systemctl --user start mcp-hub.service
fi

printf 'embed  http://127.0.0.1:8091/health\n'
curl -fsS --max-time 3 http://127.0.0.1:8091/health || true
printf '\nopenviking\n'
curl -fsS --max-time 3 http://127.0.0.1:1933/health || true
printf '\nmcp-hub %s\n' "$(systemctl --user is-active mcp-hub.service 2>/dev/null || printf 'absent')"
