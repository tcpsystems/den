# Improvements consensus — 2026-10-12

Draft proposal from a worker-consensus pass (4 OmniRoute models, see `/tmp/harness-wave/taskH1/`). Ordered by impact: autonomy, then concurrency, then cost. Not a decided plan line; a reviewer still needs to accept each one against `plan.md`'s seven fields.

1. Goal: Implement `Loop.Step()` to run one tool call, record its check, and decrement the budget.
   Why: `internal/loop/loop.go` is an empty stub (`return nil`); `autonomous-loop.md` steps 3-4-7 have no code behind them.
2. Goal: Add a chat-completions client that calls OmniRoute from the runtime itself.
   Why: no code in `internal/` or `cmd/` calls `https://ai.tcp.systems/v1/chat/completions`; `plan/GOALS.md` "Not in 0.1.0" confirms "no model call is made by the harness binaries."
3. Goal: Wire `inject.One`'s returned path into an actual loaded skill body in the context packet.
   Why: `internal/inject/inject.go` returns a path but nothing in `cmd/` reads the file; `plan/analysis/codemap/gaps.md` section 2 names this gap directly.
4. Goal: Add a command that generates a keypair, calls `grant.Sign`, and writes a grant file for a seat.
   Why: `internal/grant/grant.go` has `Sign`/`Verify` but `cmd/denmcp` only reads a grant an operator already wrote; no seat can request scope (`plan/analysis/codemap/gaps.md` section 3).
5. Goal: Enforce the concurrency cap as "the smallest limit among CPU, memory, provider quota, and the task budget."
   Why: `autonomous-loop.md` line 27 states the rule; no code anywhere computes or enforces it, so high concurrency is unbounded in practice.
6. Goal: Permanently exclude `deerflow-memory` as a callable model id in the catalog.
   Why: `plan.md` line 49 and section 7 "Cut" record 11 calls, 0 successes, ~547s each — a decided refusal that the catalog does not yet encode as a hard block.
7. Goal: Build a one-local-benchmark-row mechanism (even a simple pass/fail + latency log) that gates skill promotion.
   Why: `plan/analysis/skill-adoption.md` step 4 requires "one local benchmark row proves the skill helps" before promotion, but no script exists; `plan/GOALS.md` confirms "the injector does not rank."
8. Goal: Route `denfetch`/`denmcp` secret reads through `internal/vault`'s grant-gated reveal instead of the process environment.
   Why: `autonomous-loop.md` line 32 says credentials go only to Vaultwarden; `internal/vault/vault.go` exists but is bypassed (`plan/analysis/codemap/gaps.md` section 5).
9. Goal: Make `denctl` call `catalog.Lead` and `grant.Verify` at seat start instead of staying an observer.
   Why: `plan/GOALS.md` "Not in 0.1.0" states `denctl` calls none of `inject.One`, `catalog.Lead`, or `grant.Verify` today.
10. Goal: Add a cost/latency ledger per tool call so "Observe... cost, and latency" (loop step 1) has real numbers to read.
    Why: `autonomous-loop.md` step 1 names cost and latency as loop inputs, but no counter exists outside the one `deerflow-memory` anecdote in `plan.md`.
