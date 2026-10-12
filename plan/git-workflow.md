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
