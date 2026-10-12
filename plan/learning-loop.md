# Learning loop — 2026-10-12

The cycle a learn-and-improve cluster runs. Not a plan decision; a record of the shape the run just took.

## Cycle

1. **Learn.** Read the memory summaries and the docs that hold current decisions: `plan/autonomous-loop.md`, `plan/git-workflow.md`, `PROTOCOL.md`, `plan.md`'s decided sections, `plan/analysis/*.md`. Note the gaps those docs already admit (a `Cut` line, a `Not in 0.1.0` list, a stub function).
2. **Relearn.** Re-scan the live repo state against the docs read in step 1: run the loader's own tests (`go test ./internal/catalog/...`), grep for the schema a parser actually enforces, list directories that a task names but that do not exist on disk. Trust code and `go test` output over a doc's wording when they disagree.
3. **Analyse.** Send the same spec — the relevant doc excerpts plus the confirmed gaps — to several OmniRoute models in parallel. Each model proposes independently; no model sees another's answer.
4. **Suggest.** Collect each model's raw output unedited. A worker that returns an empty body, hits `finish_reason=length` before a complete answer, or returns a 4xx is marked discarded, not silently dropped.
5. **Consensus.** Compare the surviving outputs. A unanimous or majority answer wins as-is. A model that proposes something a decided refusal rule forbids (for example, building toward Akash or A2A without an explicit deployment task) is overruled even if more than one worker proposed it — a refusal rule outranks a worker vote.
6. **Implement.** The driver, not a worker, writes the one consensus file the task named, inside the paths the task allowed. No worker has write access; a worker's output is text to compare, never a diff to apply directly.
7. **Commit.** The orchestrator commits, after `go test ./...` (or the task's named check) passes, following `plan/git-workflow.md`'s one-branch-one-commit rule. A learn-and-improve cluster does not commit, branch, or push — that is this file's one hard boundary.

## Memory rule

Every cycle writes its proposals and outcomes to memory under a key shaped `learning-loop-YYYYMMDD` (OpenViking, `viking://~/memories`, or the equivalent flat ledger when OpenViking is unavailable per `plan/needs.md`). The entry holds: which docs were read, which gaps were confirmed against code, which models answered, the consensus result, and the file path it was applied to. A cycle that skips this write has not finished; "suggest" without a durable record is not done, per `plan/GOALS.md`'s "a plan sentence is not done" rule.

## Types

- `GapClaim{ doc string; line string; confirmedBy string }` — a gap, and the command or file that confirmed it still holds.
- `WorkerResult{ model string; status "ok"|"discarded"; content string }`
- `ConsensusResult{ winner string; overruledBy []string; appliedPath string }`

## Functions

- `Learn(docs []string) []GapClaim`
- `Relearn(claims []GapClaim) []GapClaim` — re-verify, drop a claim the code no longer supports.
- `Analyse(claims []GapClaim, models []string) []WorkerResult`
- `Consensus(results []WorkerResult, refusals []string) ConsensusResult`
- `Implement(c ConsensusResult, allowedPaths []string) error` — refuses a path outside the task's list.

## Invariant

A consensus result is applied only to a path the task named. A refusal rule beats a worker majority. No worker writes a file; only the driver does.

## File path

`plan/learning-loop.md` (this file).
