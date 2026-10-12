# Den: The Production Agentic Harness

Den is the production-grade, token-aware agentic runtime and control plane for local execution, runtime MCP multiplexing, operator TUI observability, and ephemeral permission leasing. It provides isolated workspace sandboxing, strict path-jail boundaries, and multi-model routing via OmniRoute.

## Measured state

Run `scripts/readme-metrics.sh` to refresh this block from the repo and live checks.

<!-- metrics:start -->
| Metric | Value |
| --- | ---: |
| Commit | `192e617` |
| Skills | 28 |
| Go files | 55 |
| Plan documents | 39 |
| Niche rows | 58 |
| Commits | 92 |
| Embedder | healthy |
| Long memory | healthy |
| Local memory files | 2 |
| Worker clusters | 28 |
| Worker files | 106 |
| Free models verified | 0 |
<!-- metrics:end -->

The embedder is Qwen3-Embedding-8B-Q6_K at 4096 dimensions. Accuracy and quality percentages stay unpublished until a repeated benchmark records them.

## Repository Scope & Ecosystem Boundary

- **`den` (This Repository)**: The local runtime daemon (`denctl`), operator TUI (`dentui`), read/fetch utilities (`denread`, `denfetch`), core process supervisor (`initd`), security guardrails (`path_jail`), and persistent state ledgers.
- **`tcpsystems/agent-mesh`**: Distributed agent communication mesh, node heartbeat coordination, and A2A sync.
- **`tcpsystems/radar`**: Pricing, live model indexes, provider rate cards, and telemetry state desks.
*(Note: Foxcore repositories are intentionally excluded from active focus and will be onboarded into the harness at a later stage.)*

## Core Architecture & Components

1. **Runtime (`denctl`)**: Manages agent seat lifecycles, execution telemetry, workspace jails, and local SQLite/JSON persistence.
2. **Operator TUI (`dentui`)**: A high-performance terminal user interface built with Charm (Bubbletea) providing live observability, worker selection, model routing controls, and permission escalation dialogues.
3. **Ephemeral Permissions**: Time-bounded permission leasing with automatic TTL revocation. Agents start with minimal privileges and must request explicit temporary grants with documented operational justifications.
4. **Skill & Niche Bundling**: High-frequency core skills are injected directly into active worker contexts; niche or specialized skills are deferred and retrieved via runtime discovery.
5. **Persistence & Resilience**: All consensus records, analysis summaries, and state trees are stored as immutable flat Markdown and JSON ledgers on disk, ensuring seamless recovery and state reconciliation across OpenViking or embedder rebuilds.

## Consensus and clusters

A cluster fans a single task out to N concurrent cheap or free LLM workers, typically through an OmniRoute-style OpenAI-compatible gateway. Each worker is a plain HTTPS chat-completions call: not a subagent, not a git identity, and not a Den seat. Workers run in parallel — goroutines or background shell jobs — receive the same task spec, and write raw text or code into a scratch directory. They do not coordinate with one another.

The orchestrator is a human or a capable model, never a worker. It reads every worker's output, compares them, and produces one consensus result per task by voting or merging. The majority answer wins; ties and substantive disagreements are reconciled manually, with the reason documented.

Workers never run `git commit` and never push. Only the orchestrator applies the consensus into the working tree, and only the orchestrator or a human commits, after checks such as `go test ./...` pass. This is distinct from the wave/seat worktree pattern in `internal/orch`: clusters are for cheap exploratory work, duplicate detection, and drafts; seats are for the real, git-committing swarm.

## Quick Start & Build

```bash
# Build all den binaries
go build ./cmd/...

# Run the runtime daemon or TUI control plane
./cmd/dentui/dentui --help
```

