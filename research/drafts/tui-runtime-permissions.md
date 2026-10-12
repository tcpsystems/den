# Runtime, TUI, Permissions, and Ephemeral Grants (Beta/Alpha Draft)

This note records the consensus on expanding `den` into a production-grade beta harness with a fully capable TUI, direct worker dispatch, permission leasing, and persistent disk state.

## 1. Core Architecture Split
- **Runtime (`denctl`)**: Owns the state, the agent teams, seat dispatch, permission hierarchies, and local ledger storage.
- **TUI (`dentui`)**: Not just a passive viewer. It is the operator control plane. It connects to the Unix socket (or loopback API) to inspect health, select models, configure OmniRoute/A2A routes, and dispatch worker swarms.
- **WebAPI / WebApp**: Strictly observability. It provides a real-time read-only stream of events, traces, and metrics.
- **Agent Execution**: The runtime and agent are unified. There is no separate agent binary; the daemon spawns worker routines/processes governed by the runtime's sandbox/jail rules.

## 2. Capability Access & Routing
- **OmniRoute & A2A**: The runtime holds the OmniRoute base URL and API keys securely. Workers inherit access through configured route profiles (`auto/coding`, `rookyard`, etc.) rather than raw tokens. A2A tasks use the daemon's internal stream handlers.
- **Tools vs. MCP**: MCP servers are spawned on-demand as subprocesses when a seat's grant explicitly calls for them. Local argv tools (`denread`, `denfetch`, git) run inside the path jail (`sections/security/match.py`).

## 3. Dynamic Permission Leasing (The Ephemeral Grant Model)
Agents start with a minimal default permission set defined by the session orchestrator. 
- **Requesting Access**: If a worker encounters a task requiring a restricted path, network destination, or tool, it submits a structured permission request containing a clear operational reason.
- **Delegation & Escalation**: The parent orchestrator (or human operator) evaluates the request. If authorized, the parent grants the permission *with a hard timeout* (e.g., valid for 15 minutes or 2 tool calls).
- **Automatic Revocation**: Once the timeout expires or the subtask completes, the runtime automatically revokes the temporary permission. This minimizes the window of risk for accidental damage or hallucinations.

## 4. Skill Injection & Knowledge Access
- **Dumb Skill Injector**: The orchestrator injects candidate skills based on initial task classification. 
- **Pull-Based Exploration**: Workers can query the global skill catalog, niche indexes, and OpenViking/local memory nodes (`viking://~/memories`). If a worker needs specific deep knowledge, it requests it through the runtime API, which filters and returns the relevant context snippet.

## 5. Persistence & Rebuild Resilience
To survive OpenViking resets, embedding model version swaps, or database re-initializations, all core harness state (plans, consensus records, permission graphs, agent ledgers) is stored as **flat Markdown and structured JSON files under `.den/persistence/` or `research/`**. 
- Git acts as the ultimate backup layer.
- The system can cold-boot from disk state without losing the operational consensus history.
