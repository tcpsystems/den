#!/bin/bash
# Apply worker commits onto local main as they appear.
# Stops after 6 hours, or sooner when every worker branch has a commit that is on main.
# A branch that has not committed yet is still pending. It is not "already merged".
set -u
ROOT=/home/kivox/Projects/harness-plan
LOG=/tmp/harness-plan-land.log
cd "$ROOT"
end=$((SECONDS + 21600))
base=18ab6fb
branches=$(git for-each-ref --format='%(refname:short)' refs/heads | grep -E '^r(10|[1-9])(g[1-5])?$' || true)
while [ "$SECONDS" -lt "$end" ]; do
  pending=0
  landed=0
  for b in $branches; do
    tip=$(git rev-parse "$b")
    ahead=$(git rev-list --count main.."$b" || echo 1)
    if [ "$ahead" != 0 ]; then
      pending=$((pending + 1))
      "$ROOT/scripts/land-on-main.sh" "$b" || true
    elif [ "$tip" != "$base" ] && git merge-base --is-ancestor "$b" main; then
      landed=$((landed + 1))
    else
      pending=$((pending + 1))
    fi
  done
  echo "$(date -Is) landed=$landed pending=$pending" >> "$LOG"
  if [ "$pending" = 0 ]; then
    echo "all worker branches have a commit on main" >> "$LOG"
    exit 0
  fi
  sleep 20
done
echo "poll-land stopped after 6 hours" >> "$LOG"
exit 0
