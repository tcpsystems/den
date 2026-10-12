#!/usr/bin/env bash
# Checks the last 10 commits on the current branch against the convention
# `type(scope): subject` and scans for API-key-looking strings. Does not
# stop at the first failure; prints a full report and exits 1 if any
# commit failed. Does not use -e, so every commit is checked.
set -uo pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "$REPO_ROOT"

TYPES='feat|fix|docs|test|refactor|perf|chore'
# \x1f (unit separator) splits hash from subject even if the subject has
# colons or spaces.
COMMITS="$(git log -n 10 --pretty=format:'%H%x1f%s')"

overall_fail=0
declare -a report=()

while IFS=$'\x1f' read -r hash subject; do
	[[ -z "$hash" ]] && continue
	short="${hash:0:7}"
	fail=0
	reasons=()

	# Whole subject line must match type(scope): subject and be <=72 chars.
	# "Whole subject line" (not just the text after the colon) is the
	# conventional-commits community convention and is what is checked here.
	if [[ ! "$subject" =~ ^(${TYPES})(\([a-zA-Z0-9_./-]+\))?:\ .+$ ]]; then
		fail=1
		reasons+=("does not match 'type(scope): subject'")
	fi
	if [[ "$subject" =~ [A-Z] ]]; then
		fail=1
		reasons+=("subject has uppercase letters")
	fi
	if (( ${#subject} > 72 )); then
		fail=1
		reasons+=("subject is ${#subject} chars, want <=72")
	fi

	# Scan subject + full body for secret-shaped strings. Report only the
	# rule name, never the matched text.
	body="$(git show -s --format=%B "$hash")"
	if echo "$body" | grep -qE 'sk-[A-Za-z0-9]{16,}'; then
		fail=1
		reasons+=("message matches sk-... secret pattern")
	fi
	if echo "$body" | grep -qE 'doo_v1_[A-Za-z0-9]{16,}'; then
		fail=1
		reasons+=("message matches doo_v1_... secret pattern")
	fi
	if echo "$body" | grep -qE 'ac\.sk\.[A-Za-z0-9._-]{10,}'; then
		fail=1
		reasons+=("message matches ac.sk..... secret pattern")
	fi
	if echo "$body" | grep -qE '[0-9a-fA-F]{20,}'; then
		fail=1
		reasons+=("message contains a 20+ hex-char run")
	fi

	if [[ "$fail" -eq 0 ]]; then
		echo "PASS $short ${subject:0:60}"
	else
		echo "FAIL $short ${subject:0:60}"
		overall_fail=1
		for r in "${reasons[@]}"; do
			report+=("$short: $r")
		done
	fi
done <<<"$COMMITS"

echo
echo "=== summary ==="
if [[ "${#report[@]}" -eq 0 ]]; then
	echo "all 10 commits pass"
else
	for line in "${report[@]}"; do
		echo "$line"
	done
fi

exit "$overall_fail"
