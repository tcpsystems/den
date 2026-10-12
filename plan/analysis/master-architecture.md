# Harness Master Architecture, Token Optimisation, & Swarm Orchestration Plan

This document establishes the comprehensive technical blueprint for the Den harness, detailing multi-agent swarm collaboration, token conservation (Caveman, RTK tool-slack compression, caching), role-based branch commits, and orchestrator-led verification and merging.

---

## 1. Token Optimisation & Intelligence Retention

To maximize reasoning efficiency while minimizing token waste, Den enforces strict compression and caching strategies across all agent turns:
- **Caveman & Session-Dedup**: Eliminates redundant conversational filler, conversational boilerplate, and repeated context headers.
- **RTK (Runtime Token-Slack Reduction)**: Tool outputs and file reads are truncated or filtered to exact line ranges via `denread` and `denfetch`, preventing entire file dumps into prompt contexts.
- **Prompt Optimization & Caching**: System prompts and static skill packs are pinned in cache (`x-omniroute-session` and prompt-caching affinity) to achieve high cache hit rates across iterative turns.
- **Background Degradation**: Heavy tasks (such as title generation, summarization, and tag labeling) are automatically routed to cheap/fast models (`auto/coding:cheap`), preserving premium model capacity for deep code synthesis and architectural reasoning.

---

## 2. Swarm Cluster Roles & Shared Branch Workflows

The multi-agent swarm operates across dedicated clusters, all committing to a shared feature branch (e.g., `wave/20261012-swarm-hardening`) utilizing strict semantic conventional commits:

1. **Research Cluster (`researcher`)**:
   - Scouts documentation, public repos, and MCP tools using `denfetch` and TinyFish.
   - Outputs raw, verified, source-backed data into `research/` or `plan/niches/`.
   - Commit prefix: `docs(research): ...` or `feat(research): ...`

2. **Planning & Architecture Cluster (`architect`)**:
   - Designs system specifications, state ledgers, security jails, and permission leasing logic.
   - Commit prefix: `docs(plan): ...` or `refactor(arch): ...`

3. **Implementation & Tooling Cluster (`builder`)**:
   - Writes Go binaries (`den`, `denctl`, `dentui`, `denread`, `denfetch`), security matchers (`path_jail`), and runtime daemon loops.
   - Ensures all Go code passes `go test ./...`.
   - Commit prefix: `feat(den): ...` or `fix(runtime): ...`

4. **Verification & Review Cluster (`reviewer`)**:
   - Audits code against security policies, token limits, and test coverage.
   - Commit prefix: `test(verify): ...` or `style(review): ...`

---

## 3. Orchestrator Feedback & Merge Pipeline

- **Autonomous Role Execution**: Workers commit directly to the shared branch with conventional commit messages reflecting their assigned role.
- **Orchestrator Analysis Loop**: The orchestrator reviews incoming commits, verifies test outputs via `run_terminal_command`, applies minor fixes directly, or feeds constructive critique back to the swarm cluster.
- **Final Integration**: Once all cluster tasks pass verification, the orchestrator packages the branch into a clean, ready-to-push state adhering strictly to semantic versioning and git standards.
