#!/usr/bin/env bash
# Local CI: format, vet, test, race test (best-effort), build the binaries.
set -euo pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "$REPO_ROOT"

echo "==> gofmt"
FMT_OUT="$(gofmt -l .)"
if [[ -n "$FMT_OUT" ]]; then
	echo "gofmt FAILED: the following files need formatting:"
	echo "$FMT_OUT"
	exit 1
fi

echo "==> go vet"
if ! go vet ./...; then
	echo "go vet FAILED"
	exit 1
fi

echo "==> go test"
if ! go test ./...; then
	echo "go test FAILED"
	exit 1
fi

echo "==> go test -race"
if ! go test -race ./internal/... ./cmd/...; then
	if [[ "$(go env CGO_ENABLED)" == "0" ]]; then
		echo "WARNING: race detector unsupported (CGO_ENABLED=0); skipping race failure"
	else
		echo "go test -race FAILED"
		exit 1
	fi
fi

echo "==> go build"
if ! go build ./cmd/denctl ./cmd/dentui; then
	echo "go build FAILED"
	exit 1
fi

echo "CI OK"
