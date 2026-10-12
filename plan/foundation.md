# Foundation

This is the building block. Later work expands one section of it. It does not reopen `plan/BASE.md`. The second analysis has not been started.

Four readers looked through this session's stopped subagents and the research notes. 140 completed outputs were digested. 79 cancelled or failed outputs were empty. An empty output is not a finding.

## Block

den is the harness (`tcpsystems/den`, checkout `~/Projects/harness-plan`, branch `develop`). tcp.systems is the system it folds into. Foxcore (`fox`, `foxd`, `floki`) is the company that system operates. radar (`tcpsystems/radar`, private) is the price and tracker catalog. The three brains stay separate. Names are `plan/names.md`. Processes are `plan/stack.md`.

The user talks to one operator. The operator writes a packet and starts seats. A seat does the named function and returns. The human starts the next analysis.

## Runtime

One packet is one turn. The turn ends. It does not resume inside itself.

1. The operator writes the seven fields. Rail is empty unless the human side record already flipped it.
2. `select-skill` returns one path. No match returns `use-harness` only. The catalog is not copied in.
3. The seat reads the Read paths. A path outside the section root is a refusal from `path_jail`.
4. The seat makes one tool call. The call is an argv vector. `run_argv` refuses a shell metacharacter. The call runs only after the jail accepts it.
5. The result is appended. A repeat of the same Function, Read paths, and skill path reads the stored file and does not call the tool again.
6. The seat writes Write, or a refusal row, and returns. Next is the operator's next packet, not a jump the seat takes.

A child is a new packet. The parent turn is over when the child starts. The child does not see the parent transcript.

The model server owns attention, the KV cache, and sampling. The harness owns the packet, the skill file, the ledger compare, and the tool result. A thought that does not change the packet is not a step.

## Security, in this order

Prefer the earlier control. A later one does not replace it.

1. Path jail. `sections/security/match.py` `path_jail`. The root is the packet root, not a process-global directory.
2. Argv. `run_argv`. One vector. No shell string.
3. Rail. `set_rail` accepts a flip only from the human side record. A prompt, a tool result, and a child cannot flip it.
4. Refusal row. A refused call is recorded. It is not repaired, not retried, and not turned into a success.
5. Scope. No secret in the packet, the vault, or `/api/state`. A private host is not fetched. OmniRoute runs only when the packet names it.

A yes from a person does not add a path. The jail still has to accept the path.

## No progress

Backtrack when the turn is not moving the Function. Stop. Do not take another lap to look busy.

| Sign | What the seat does |
| --- | --- |
| The same Read set and the same tool argv appear again | Return the stored result. Do not call. |
| The tool was refused | Write the refusal. Stop. Do not rephrase the argv. |
| The provider stopped for length, or the arguments were repaired | Do not run the call. Write an error. Stop. |
| Two notes in a row add no Write and no refusal | Stop and return the packet to the operator. |
| The seat is choosing a model, a skill rank, or a next seat | Stop. Those are not the Function. |
| A thought names a new tool that the packet did not name | Stop. One call was the budget. |

A backtrack writes one line: the sign, and the field that did not change. It does not delete the earlier row. The operator writes the next packet. The seat does not invent a new goal to keep the turn alive.

## Watch

A seat that is running can be watched. A seat that has returned is a row, not a process. On 2026-10-12 the research seats had returned. Their notes are the `wave/20261012-*` branches. The hub at `127.0.0.1:8787` was up with `connected: 0`. `denctl` was not running. The live view on `127.0.0.1:8790` was not running.

| Point | What it is | Where it is read |
| --- | --- | --- |
| Goal | The Function field. One sentence. It does not change during the turn. | The packet |
| Test | The check named in the skill. Pass or fail. A fail is a refusal row. | The seat result |
| Metric | One number the test already prints: pass, fail, or the control that refused. No second score. | The row |
| Checkpoint | A commit on the seat's own branch. The working tree of `develop` is not a checkpoint. | `git log` on that branch |
| Timeline | Append-only lines: accepted, running, refused, no-progress, returned. | The mesh event log, or `GET /api/events` when `denctl` is up |

A websocket, when `denctl` is running, streams those timeline lines for one section. It does not start a seat, does not merge, and does not write the packet. The read-only mesh page at `127.0.0.1:8787` is the same kind of window for hub tasks. Port 8790 is the graph view, and it is a separate process. If that process is down, the timeline is the git log and the hub health JSON.

Watching is not progress. A seat that only emits thoughts has hit the no-progress exit. The stream shows that exit. It does not keep the turn open.

## Confirmed

Two or more notes, or `plan/BASE.md`, already state these.

- The packet has seven fields: Section, Function, Read, Do, Write, Next, Rail. It does not grow an eighth field.
- The packet is handed to the seat. An inline plan comes before a write. One tool call, then the result, until a final message.
- The mesh is bus, store, and fetch. It does not choose a model or the next seat.
- One memory store. A later section supersedes. Progress is append-only. A seat does not delete.
- One skill body, loaded because the description matches. The skill does not name a model.
- CLI is argv only. Paths stay inside the section root.
- A child gets its own packet. It does not write the parent's Write, Next, or Rail. It does not inherit the parent's conversation. It does not dispatch another child.
- A lead who cuts Next does not write the corpus. The mesh keeps the current section and a refusal row. The cut is not a clean stop and not a new section.
- Rail stays empty until a human flips it. A seat cannot write that flip. A steer line is a message into the live step. The block reason stays on the human side.
- Foxcore stays one product repo and one pull request. fox owns desired state on Cloudflare Workers, D1, and Durable Objects. foxd and floki dial out. den does not write fox desired state.
- Hetzner is lab and staging. The tcp price desk is research. It does not set the Foxcore price.
- Training loops, harness-edit loops, and AirLLM stay off the live packet. AirLLM is a later disk-stream slot. The RX 6600 runs one embedding model.
- Embeddings are `127.0.0.1:8091` on core. From tcp the hostname is `http://core.ts.tcp.systems:8091`. If that embedder is down, the seat waits.
- OmniRoute at `https://ai.tcp.systems` is a tool. It is not the default model hop. The fixed combo is `harness-winner`. A catalog walk is not a route.
- Fast and recent memory are local files. OpenViking at `127.0.0.1:1933` is the long store. A backup may copy fast and recent to `tcp:/opt/harness-memory`. That host is not the read path.

## Measured once

Re-assess before expanding. These are seat reports, not a fresh probe.

| Fact | Source |
| --- | --- |
| LangChain is a library on LangGraph. Fit 1/5 as the fox brain on Workers, 3/5 as a desk experiment. | LangChain brief |
| LangGraph is an on-demand workflow library. The hosted product name is LangSmith Deployment. | LangGraph brief |
| Langfuse beside the desk, 4/5. LangSmith on production DeerFlow, off. Self-hosted LangSmith on the desk, 2/5. fox tracing, 1/5. | tracing briefs |
| DeerFlow `38c55c6` has no A2A server. Subagents are in-chat workers and do not delegate. | DeerFlow audit |
| OmniRoute 3.8.51 had `a2aEnabled` false and memory injection off when that seat looked. | OmniRoute brief |
| `origin/section/packet` reported tests passed. `section/store` and `section/steer` are pushed. None of the three are merged. | seat notes and `develop` |
| `r*` branches are single unreviewed drafts. Placeholder drafts were called out on `r4g4`, `r6g4`, and `r8g4`. | half-finished review |
| Live Caddy doors are the table in `plan/names.md`. Apex redirects to `chat.tcp.systems`. | host read, 2026-10-11 |
| tcp.systems is a `cax41` with the 160 GB disk kept. | host resize |

## Borrow

Named structures. They are not installed.

| Structure | Use when a packet names it |
| --- | --- |
| Reasonix facts pack | A review child gets confirmed decisions, an evidence summary, and file anchors. Task in, final answer out. `write_paths` is that child's write set. |
| Kimi child | No parent history. Built-in children do not spawn. |
| DeepSeek pre-step | An empty or rejected claim spends no step and does not commit the corpus. |
| Qwen Plan | Read-only mode is a switch. A sentence does not enter it. Deny rules sit ahead of any classifier. |
| Cursor trees | Skills, rules, and MCP are different trees. A planner publishes tasks. A worker returns one handoff and does not see its siblings. |
| agentskills.io | Catalog is name plus description. One body loads after selection. |
| Vercel `ToolLoopAgent` | One generate or stream turn. Stop on a final message or the step cap. |
| Ponytail ladder | Need, stdlib, native, installed dependency, then write. |
| ECC loop | Acceptance first, then eval, then freeze, audit, shrink, or replay. |
| A2A card | Only for a specialist that is its own service. Seat-to-seat stays the packet. |
| MCP | `tools/list`, then `tools/call`. Discover, then one call. |

## Refused

- A model picker inside the mesh, a combo that walks the whole catalog, or a vendor flag that chooses the model for the seat.
- A second store: a graph database as memory, a vector side store, an OmniRoute memory bag, or a tool side file.
- Seeding a child with the parent's log, always-allow rules, or write rights.
- A hook that fails open. A blocked prompt stored nowhere. Silence as a cut.
- A seat that deletes, interprets generated code as the tool, or runs a training pass on the live packet.
- LangChain or LangGraph as the fox control plane. DeerFlow as the public product brain. agent-mesh as a second brain.
- Production tracing shipped to a cloud that copies prompts off the tailnet.
- Loading every skill body. Putting a second GPU model beside the embedder.
- A Git forge or a notes server on tcp.
- A retired label that was removed from `nginx/tcp.systems.conf` in radar. It is not a place in `plan/names.md`.

## Empty

79 subagents were cancelled or failed and returned no text. Their assignments are not results. The gaps they leave, still unwritten:

- Persistence research before any new OmniRoute store.
- OmniRoute metric names and local-embedding research, as their own notes.
- Niche documents for cloud, market, virt, and the security-limit writeup.
- The second-pass `g2` seats for packet, store, steer, skills, gate, security, efficiency, consensus, loop, and child.
- A single-writer store proof, a Rail inject proof, and a memory-storage map.

Completed research files already exist beside the session plan. They are suggestions until two notes agree. `verified-lines.md` kept one suggestion and dropped twelve.

## Expand from here

Use one verb from `plan/names.md`. One subject. One output file `plan/analysis/<subject>.md`. Stop. Ascend into `plan/BASE.md` only after a second note agrees.

| Slot | Starts from | Status |
| --- | --- | --- |
| Names and doors | `plan/names.md` | written |
| Process split | `plan/stack.md` | written. `denctl` socket-only and `denapi` are edited and do not compile yet. |
| Operator | `plan/operator.md` | written |
| Research suggestions | session `research/*.md` | one source each |
| Section Go | `section/packet`, `section/store`, `section/steer` | pushed, not merged |
| Ecology | host capacity on core and tcp | re-assess |
| Economy | radar price desk, one unit per note | not started |
| OmniRoute | `plan/omniroute.md` and `harness-winner` | settings exist, optimisation pass not started |
| Multi-agent | one orchestrator, four readers, one skeptic | this foundation used four readers |

## File anchors

| Path | Role |
| --- | --- |
| `plan/foundation.md` | This block. |
| `plan/names.md` | Grammar. |
| `plan/BASE.md` | Closed base. |
| `plan/stack.md` | Daemon, API, TUI. |
| `PROMPT.md` | Reasonix profile prompt for the pass the human starts. |
