#!/bin/bash
# Copy fast and recent memory to the OmniRoute host. Seats still read the local copies.
set -euo pipefail
local_root="${HOME}/.local/share/harness/memory"
remote="${HARNESS_MEMORY_REMOTE:-tcp:/opt/harness-memory}"
mkdir -p "${local_root}/fast" "${local_root}/recent"
ssh -o ControlMaster=no "${remote%%:*}" "mkdir -p ${remote#*:}"
rsync -a --exclude '.git' "${local_root}/fast" "${local_root}/recent" "${remote}/"
echo "backed up fast and recent to ${remote}"
