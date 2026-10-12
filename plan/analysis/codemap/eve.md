# Codemap — Eve notes versus den code

Wave `codemap-20261012`, seat `c7`. Research map only. Sources are the four Eve notes named below, plus `internal/inject`, `internal/grant`, and `internal/orch` / `cmd/denorch` in this worktree. None of the four notes were missing.

## directory.md

Source: `/home/kivox/Projects/harness-plan-wt/eve-20261012-r1/research/drafts/eve/directory.md`.

Copy the path-as-name rule: one directory, filename is the id, no second registry that can drift, and keep standing text small so procedure lives in a skill loaded by description.

Already in code: `inject.One` returns one catalog skill path when a grant scope is `skill:<id>` and the seat matches (`internal/inject/inject.go`); no match or a bad grant returns `skills/use-harness/SKILL.md` and false. It does not rank or embed.

Still a note: `agent/instructions.md` as a complete agent, `skills.txt` progressive discovery, `tools/*.ts` with the filename as the tool name, and dropping `agent.ts`. The catalog is a JSON index, not a compiled agent directory.

## capabilities.md

Source: `/home/kivox/Projects/harness-plan-wt/eve-20261012-r2/research/drafts/eve-product/capabilities.md`.

Copy, in that note's order: checkpoint a finished step and replay it, a typed approval before a side effect, and one cheap eval before accepting a seat. Do not copy channels, schedules, Vercel Sandbox, or Connect.

Already in code: `grant.Verify` rejects a bad parent signature, the wrong seat, a missing scope, and a clock outside `nbf`/`exp` (`internal/grant/grant.go`). `denorch take` runs checks (default `go test ./...`) and merges only a clean seat (`internal/orch/wave.go`).

Still a note: step journals, replay, an approval tool the model must call, and the continuation token. There is no crash resume of an in-flight seat.

## delegation.md

Source: `/home/kivox/Projects/harness-plan-wt/eve-20261012-r3/research/drafts/eve-agents/delegation.md`.

Copy the parent/child split den already aims at: the child gets a fresh context and its own instructions, the parent is the combiner, writes do not overlap, and a child cannot delegate. Do not copy `taskId` continuation, the 32-task cap, or Workflow `ctx.agent`.

Already in code: `denorch open` makes one wave branch plus one branch and worktree per seat (`wave--seat`), and `Commit` refuses paths outside `Owns` (`internal/orch/wave.go`, `cmd/denorch/main.go`). The package comment says it does not push.

Still a note: dynamic and remote subagent kinds, the task receipt, `task.result`, and the turn that waits until children finish. Seats are git workers, not model-called children.

## borrow.md

Source: `/home/kivox/Projects/harness-plan-wt/eve-20261012-r4/plan/analysis/eve/borrow.md`.

Copy the five designs in its build order: agent-as-directory, durable steps, a permission lease, subagent delegation, then one eval before take. That note already says not to port Workflow SDK, Vercel Sandbox, channels, Connect, `agent.ts`, or Zod.

Already in code: the lease slice is the closest. `grant.Contract` is one seat, one parent, a reason, a scope, and a deadline, signed with ed25519, with no secret in the bytes. `inject.One` is the skill-path half of "description in the prompt, body on load." `denorch` `open`, `commit`, and `take` are the seat orchestrator that note tells den to keep.

Still a note: `agent/instructions.md`, `skills.txt`, `steps/<id>/journal.jsonl`, `Seat.Check`, and a child grant signed by the parent. `take` has no per-seat check list yet, and nothing auto-revokes a lease except the expiry clock.
