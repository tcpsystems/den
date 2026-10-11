#!/bin/bash
# Bootstrap a fresh Debian host so the harness can run worker microVMs.
# The mesh and the RX 6600 embedder are not installed here. They are separate.
set -euo pipefail

if [ "$(id -u)" -eq 0 ]; then
  echo "Run this as the user who will own the worker VMs, not as root." >&2
  exit 1
fi

if ! command -v sudo >/dev/null; then
  echo "sudo is required." >&2
  exit 1
fi

sudo apt-get update
sudo apt-get install -y --no-install-recommends \
  ca-certificates curl qemu-kvm
sudo usermod -aG kvm "$USER" || true

if ! command -v smolvm >/dev/null; then
  curl -fsSL https://smolmachines.com/install.sh -o /tmp/smolvm-install.sh
  bash /tmp/smolvm-install.sh --no-modify-path
fi

install_dir="${HOME}/.local/bin"
if [ ! -x "${install_dir}/smolvm" ] && [ -x "${HOME}/.smolvm/bin/smolvm" ]; then
  mkdir -p "${install_dir}"
  ln -sfn "${HOME}/.smolvm/bin/smolvm" "${install_dir}/smolvm"
fi

echo "smolvm is installed. A worker VM is: smolvm machine create --net --name <seat>"
echo "Embeddings, when a GPU is present, listen on 127.0.0.1:8091. Do not put a second model on that GPU."
echo "Log out and back in if the kvm group was just added."
