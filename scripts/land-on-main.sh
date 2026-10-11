#!/bin/bash
# Land one worker branch onto local main.
# Refuses the merge when the diff is not exactly that worker's file.
# Does not touch plan.md. Does not push.
set -euo pipefail
ROOT=/home/kivox/Projects/harness-plan
LOG=/tmp/harness-plan-land.log
branch=${1:?branch}
cd "$ROOT"

file_for() {
  local b=$1 slug
  case $b in
    r1|r1g*) slug=packet ;;
    r2|r2g*) slug=loop ;;
    r3|r3g*) slug=store ;;
    r4|r4g*) slug=gate ;;
    r5|r5g*) slug=child ;;
    r6|r6g*) slug=skills ;;
    r7|r7g*) slug=security ;;
    r8|r8g*) slug=efficiency ;;
    r9|r9g*) slug=steer ;;
    r10|r10g*) slug=consensus ;;
    *) echo "unknown branch $b" >&2; return 1 ;;
  esac
  case $b in
    *g*) printf 'sections/%s/%s.md\n' "$slug" "$b" ;;
    *) printf 'sections/%s/LEAD.md\n' "$slug" ;;
  esac
}

allowed=$(file_for "$branch")
ahead=$(git rev-list --count main.."$branch")
if [ "$ahead" = 0 ]; then
  echo "skip $branch already on main" | tee -a "$LOG"
  exit 0
fi

names=$(git diff --name-only main..."$branch")
if [ "$names" != "$allowed" ]; then
  echo "refuse $branch touches:" | tee -a "$LOG"
  echo "$names" | tee -a "$LOG"
  echo "allowed only: $allowed" | tee -a "$LOG"
  exit 1
fi

if ! git diff --quiet || ! git diff --cached --quiet; then
  echo "refuse $branch: main worktree is dirty" | tee -a "$LOG"
  exit 1
fi

flock -w 180 "$ROOT/.git/land.lock" git merge --no-ff --no-edit "$branch" -m "land $branch

Only $allowed. plan.md unchanged."
echo "landed $branch $allowed" | tee -a "$LOG"
