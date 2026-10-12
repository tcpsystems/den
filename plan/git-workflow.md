# Git Workflow, Version Control & Semantic Standards

Every change within the Den harness ecosystem—whether runtime code, TUI controls, skill packs, or persisted analysis ledgers—must be tracked, versioned, and committed strictly through standardized Git workflows. 

## 1. Commit Message Conventions (Conventional Commits)
All commits must follow the conventional commits specification using semantic prefixes:
- `feat(<scope>):` New features, agent capabilities, or runtime components.
- `fix(<scope>):` Bug fixes, permission leash corrections, or jail escapes.
- `chore(<scope>):` Maintenance tasks, dependency updates, build scripts, or cleanup.
- `refactor(<scope>):` Code restructuring without changing external behavior.
- `docs(<scope>):` Documentation updates, plan expansions, or research notes.
- `test(<scope>):` Adding or correcting test suites and verification scripts.
- `perf(<scope>):` Performance optimizations (e.g. embedding speed, streaming).
- `style(<scope>):` Formatting, linting, or TUI visual styling.

## 2. Semantic Release Versioning
Releases and tags follow Semantic Versioning (`MAJOR.MINOR.PATCH`):
- **MAJOR**: Breaking changes to runtime APIs, protocol formats, or jail security rules.
- **MINOR**: Backward-compatible new features, added worker capabilities, or new skill packs.
- **PATCH**: Backward-compatible bug fixes, performance tweaks, or documentation updates.

## 3. Branching & Worktree Discipline
- **`main` / `develop`**: Protected branches. Direct commits to `main` are strictly prohibited.
- **Feature/Wave Branches**: Work must be performed in isolated branches (e.g. `wave/YYYYMMDD-<feature>`) or git worktrees (`~/Projects/harness-plan-wt/`).
- **Atomic Commits**: Each commit must represent a single, verifiable logical unit of change accompanied by passing test output.

## 4. Backup

A finished, checked commit is pushed to the existing GitHub remote at the end of its wave. Secrets, credentials, caches, and unaccepted research are not committed. A conflict returns to the owning worker. The orchestrator does not rewrite the branch.

## 5. Implementation order

1. Consolidate orchestrator and commit checking.
2. Index every skill and verify its grant.
3. Fix memory identity and retention.
4. Trace and reduce `/api/state` latency.
5. Add the three-width screenshot check.
6. Add the read-only OmniRoute evidence client.

## 6. Completion Pipeline
Close out a wave in this order:
- Cluster lands: worker/consensus results land in a scratch dir.
- Orchestrator reviews consensus diffs.
- One conventional commit per task (`feat`/`fix`/`docs`, matching §1 prefixes) on the wave branch.
- `go test ./...` must pass before anything merges.
- Fast-forward-only merge of finished branch drafts onto the wave branch.
- Consolidation pass: dedupe and canonicalize what landed.
- Write an assessment note to OpenViking memory, key pattern `assessment-wave-YYYYMMDD`.
- Next project handoff order: radar first, then ideology, skills, fox, foxd, floki.
