# Plan versus code

Read 2026-10-12 from this worktree. A plan sentence is not an implementation. A row is shipped only when a Go file in this tree defines the behavior. `plan/foundation.md` is not in this worktree (the changelog names it; the file is absent). Section leads, names only: `sections/child/LEAD.md`, `sections/consensus/LEAD.md`, `sections/efficiency/LEAD.md`, `sections/gate/LEAD.md`, `sections/loop/LEAD.md`, `sections/packet/LEAD.md`, `sections/security/LEAD.md`, `sections/skills/LEAD.md`, `sections/steer/LEAD.md`, `sections/store/LEAD.md`.

## Plan docs that describe shipped code

These four docs name behavior that a Go package in this tree actually has. The doc is still prose. The package is the code.

| Plan doc | What the prose claims that the package does | Go package |
| --- | --- | --- |
| `plan/protocols.md` | Fixed names: identity, role, niche, profile, skill, tool, mcp, plugin, memory, context, grant, model, seat, wave. `catalog.Lead` binds with no required MCP. A grant is ed25519 over canonical JSON (seat, parent, reason, scope, deadline). `denorch open` is one branch and worktree per seat; `denorch commit` then `denorch take`. | `internal/catalog` (`terms.go` `Terms`, `catalog.go` `Builtin`, `Lead`). `internal/grant` (`grant.go` `Contract`, `Sign`, `Verify`). `internal/orch` (`types.go` `Seat`, `Wave`; `wave.go` `Open`, `Commit`, `Take`). `cmd/denorch`. |
| `plan/READY.md` | Checked rows: protocols file, `catalog.Lead` / `TestLeadBindsWithoutRequiredMCP`, `internal/grant` `TestSignVerifyScopeAndExpiry`, MCP catalog with no required provider, `internal/orch` `Open`, `cmd/den` init, `go test ./...`. | Same three packages, plus `internal/mcp` (`spawn.go` `Spawn` starts only `mcp:<id>` the grant names). `mcp/catalog.json` is data, not a Go package. `cmd/den` is named and is not a directory in this tree (`cmd/` has `dencat`, `denctl`, `denfetch`, `denmcp`, `denorch`, `dentui`). |
| `plan/analysis/consensus.md` | 2026-10-12 record: `Builtin` and `Lead` staff the harness; evidence cited is `internal/catalog/catalog.go`. | `internal/catalog` only. The "Accept" line is a note, not a second package. |
| `plan/git-workflow.md` | Conventional prefixes `feat|fix|chore|refactor|docs|test|perf|style`, and one worktree per seat under `~/Projects/harness-plan-wt/`. | `internal/orch` (`commit.go` `conventional`, `CheckMessage`; `wave.go` `Open` writes `harness-plan-wt/<wave>-<seat>`). Role type is `docs` / `feat` / `test` (`roleType`). Semantic-release `MAJOR.MINOR.PATCH` and a protected `main` are in this doc and have no Go. |

`internal/inject` is shipped and no plan doc describes it. `inject.go` `One` returns the catalog path for the first `skill:<id>` on the grant, or `skills/use-harness/SKILL.md`. Tests: `inject_test.go`. Callers are the package tests only.

## Plan docs with no Go package

No `package` under `internal/` or `cmd/` implements these. Some name a binary that exists for a different job (`cmd/denctl`, `cmd/dentui`). That binary is not the doc.

| Plan doc | Work the prose describes | No Go package for |
| --- | --- | --- |
| `plan/BASE.md` | Architectural base. Seven packet fields. Memory tiers. Operator SSE. Loop. Safety copy list. "Done" / "not merged" claims. | Packet type, memory store, turn loop, Rail. `cmd/denctl` and `internal/console` exist; they are not this base. |
| `plan/architecture.md` | Seat packet, one skill, one tool, smolvm guest, fox/foxd/floki split. | Seat loop, guest, edge agents. |
| `plan/controller.md` | Minimal controller: Go API, static Svelte page, loopback `127.0.0.1:8793`, start/stop after a confirm. | Svelte app. Start/stop button. `cmd/denctl` listens; it does not implement this page. |
| `plan/git-orchestrator.md` | Cluster merge master: Caveman compression, RTK tool-slack, `search_replace` auto-fix, semantic squash, `ready-to-push`, ledgers in `research/drafts/`. | Compression, auto-fix, squash, push marker. `internal/orch` merges; it does not do those steps. |
| `plan/horizon.md` | Long-horizon holds: reproducible, secure, maintainable, DRY, persistent, contextual, memory-aware. | A horizon package. The table points at other plan files and `sections/security/match.py`. |
| `plan/leaders.md` | Big leader, orchestrator, section seat, council, crutch seats. | Leader or council process. |
| `plan/needs.md` | Controller product is absent. `r*` drafts. Requirements list. | The later controller. This file says `console/server.py` is a sketch. |
| `plan/nest.md` | Paths and remotes for harness, radar, fox, foxd, floki, ideology, agent-mesh. Host service list. | Nest registry. Paths are notes. |
| `plan/omniroute.md` | When a combo, batch, A2A, MCP, compression, embedding, or memory read is allowed. | Provision gate. `internal/mcp` only spawns a granted server. |
| `plan/operator.md` | Operator surface, Unix socket `den.sock`, TUI, web limits, soul/persona record. | Operator package. `cmd/denctl` and `cmd/dentui` are binaries; this doc is the wish for them. |
| `plan/standup.md` | How to get the harness up. Mesh, embedder, OpenViking, OmniRoute checks. | Stand-up runner. |
| `plan/start.md` | Onboarding order: wants, onboard note, memory, analysis, then seats. | Onboard command. |
| `plan/structures.md` | Optional LangChain, LangGraph, Deep Agents, OpenViking structures. | Those runtimes. |
| `plan/analysis/master-architecture.md` | Caveman, RTK, prompt cache, background degradation, swarm cluster blueprint. | Token pipeline. No package reads this file. |
| `plan/analysis/skill-adoption.md` | Capability-to-source matrix (ECC, Picrew, and others). | Adopted skills. The file says a source stays a suggestion. |
| `plan/niches/INDEX.md` | Niche map. A line becomes code only after two notes and a check. | Niche fillers. |
| `plan/niches/exercises.md` | Match exercises on hosts we own. | Exercise runner. |
| `plan/niches/hypotheticals.md` | Hypothetical control checks. | Hypothetical runner. |
| `plan/niches/prompt-analysis.md` | Notes on `prompts/seat.md`. | Prompt injector. `internal/inject` does not read this file. |

Section leads are drafts, not packages. `sections/security/match.py` is Python, not a Go package. `sections/workers/dev-env.md` is a worker note, not a lead, and has no package.

## Shipped Go this map does not call a plan

`internal/catalog`, `internal/grant`, `internal/orch`, `internal/inject` are the four asked. Also in the tree, and not treated as those plan docs: `internal/console`, `internal/fetch`, `internal/mcp`, `internal/vault`, `cmd/dencat`, `cmd/denctl`, `cmd/denfetch`, `cmd/denmcp`, `cmd/denorch`, `cmd/dentui`.
