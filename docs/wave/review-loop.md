# Review loop

## Question

How should den review parallel workers? The loop is plan, implement, assess, test, evaluate, improve, reassess, feedback, approval, push. A later reviewer with the harness rules decides merge or return. Workers are already on separate `wave/20261012-*` branches from `94a9722`. `develop` is dirty and seats must not write it.

## Sources

- https://code.claude.com/docs/en/code-review — Anthropic Code Review. Multiple agents read one diff plus surrounding code in parallel. Each looks for a different class of issue. A verification step checks candidates against actual code behavior and drops false positives. Findings are severity-tagged (important, nit, pre-existing) and do not approve or block. The check run stays neutral. `CLAUDE.md` and `REVIEW.md` tune what is flagged. A local `/code-review` command reviews a diff without the hosted service.
- https://github.com/anthropics/claude-code/tree/main/plugins/code-review — the plugin behind that command. Four agents audit one PR in parallel. Only findings at confidence 80 or above are posted. A score of 0 to 100 is assigned, and below 80 is a false positive.

## Finding

Review the worker branches. Do not review `develop`.

Each seat already planned, implemented, assessed, tested, evaluated, improved, and reassessed on its own branch. The reviewer does not rerun that inner loop and does not edit the branch. It reads `git diff 94a9722..<branch>` and the harness rules (`DEVELOP.md`: review the diff against the base, do not land on `main`; one section per branch; Rail stays a human side record).

One reviewer packet, four read-only passes on that same diff, then one verdict:

1. Scope. Files stay inside the assigned section. No secrets. No write to `develop`, `main`, or another seat's path.
2. Behavior. Logic, edges, and regressions against the packet's Write. A claim is kept only if the diff or a test shows it.
3. Safety. Path jail, argv-only tools, Rail not flipped by a prompt or a child packet.
4. Proof. The seat's test matches the Write. A missing or mismatched test is a return, not a silent pass.

Tag each finding important, nit, or pre-existing. Pre-existing does not return the branch. Important does. Nits go in the note and do not block. Drop anything the diff does not support. The four passes do not vote. The later reviewer writes one line: `merge` or `return`. Merge is a decision for the land step. Push stays on the worker's own branch. Approval is not a commit onto `develop`.

## What den should not copy

Do not copy Anthropic's hosted fleet, GitHub check-run, `@claude review` trigger, reaction tuning, or Team/Enterprise gate. Do not copy the plugin's four aspect agents as a skill catalog. Do not let a score, a bot comment, or a neutral check approve or block. Do not treat a review pass as a license to write `develop`.

## One plan line

A later reviewer reads each `wave/20261012-*` diff from `94a9722` in four read-only passes, drops unverified findings, and returns `merge` or `return` without writing `develop`.

## Unproven

Confidence 80 is the plugin's cutoff, not a measured den threshold. Parallel aspect agents versus one reviewer with four passes is untested here. No wave branch was reviewed while writing this note.
