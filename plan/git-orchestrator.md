# Local Git Orchestrator & Swarm Workflow

This document defines the "Local Commit Orchestrator" pattern for the Den harness. It enables high-concurrency multi-agent swarms to work on a single feature branch while maintaining strict semantic versioning, token efficiency, and code quality through automated verification and merging.

## 1. The Swarm Cluster Model
Instead of isolated branches for every minor task, workers are grouped into **Clusters** targeting a specific `wave/YYYYMMDD-feature` branch.
- **Role Assignment**: The orchestrator assigns specific file paths or logical units (e.g., `internal/fetch`, `cmd/dentui`) to individual workers within the cluster.
- **Shared Context**: All workers in a cluster share the same `BASE.md` and `plan/` context but operate in isolated git worktrees (`~/Projects/harness-plan-wt/`) to prevent filesystem collisions during execution.

## 2. Token-Efficient Intelligence
To retain intelligence while minimizing token burn:
- **Caveman Compression**: Strips conversational filler from all agent outputs before they enter the commit message or ledger.
- **RTK Tool-Slack**: Uses `denread` with precise line offsets to avoid dumping entire files into the LLM context.
- **Deep Knowledge Cache**: Hits against `sections/store/memory-map.md` and OpenViking embeddings are prioritized over re-deriving known patterns.

## 3. The Orchestrator Loop
The Orchestrator acts as the local "Merge Master" for the swarm:
1. **Ingest**: Workers submit their changes as local commits in their worktrees.
2. **Verify**: The Orchestrator runs `go test ./...` and `bash -n` on scripts for every incoming commit.
3. **Feedback & Fix**: 
   - If tests fail, the Orchestrator provides targeted feedback to the specific worker.
   - For minor syntax or linting errors, the Orchestrator applies the fix directly using `search_replace`.
4. **Semantic Squash**: The Orchestrator merges verified work into the main feature branch, ensuring all commits follow the `feat(scope):`, `fix(scope):`, `chore(scope):` standard.
5. **Finalize**: Once the cluster's goals are met, the branch is cleaned up and marked as `ready-to-push`.

## 4. Persistence & Resilience
All orchestration decisions, feedback logs, and consensus records are written to flat Markdown ledgers in `research/drafts/` and `plan/analysis/`. This ensures that if the runtime restarts or OpenViking rebuilds, the swarm's progress and reasoning are fully recoverable.
