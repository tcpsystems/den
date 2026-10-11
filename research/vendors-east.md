# Vendors east

This note deepens `plan.md`. It is not a second plan. Nothing here is decided. A line moves into the plan only when two finished notes state it and no lead cuts it. One page, including two pages inside this file, stays a suggestion.

The mesh stays the bus, the store, and the fetch. It does not choose a model. AirLLM is not in this pack and is not the loop. Quota essays, auth failures, and missing model ids are ignored.

Pages below were opened. A docs host that returned only a title is not cited.

## What it is

### DeepSeek Harness

`deepseek-ai/deepseek-harness` (`dsh`) is the org's pinned harness: everything is a plugin, on Cordis, MIT, developer preview, local web UI on `127.0.0.1:3080`. Docs: [architecture](https://github.com/deepseek-ai/deepseek-harness/blob/master/docs/architecture.md), [agent lifecycle](https://github.com/deepseek-ai/deepseek-harness/blob/master/docs/agent-lifecycle.md), [subagents](https://github.com/deepseek-ai/deepseek-harness/blob/master/docs/subsystems/subagent.md).

A turn is an inbox claim, then `agent/pre-step`, then at most one admitted step. The pre-step listener may rewrite or reject the claim. A reject, an empty first claim, or a failed pre-step removes the batch and spends no step. A durable turn still closes. System prompt and user text are not committed when the async admit is cancelled, and an empty claim never becomes a step. Replay lives on `session/event`. Live `agent/*` is coordination, not the corpus.

A child start carries a `prompt` (content blocks delivered as the child's user message), a parent pointer, an optional tool filter, an optional depth cap, and an optional persona. The filter removes named tools from the child prompt and refuses them at execution. Parent delivery returns a settlement notice. Running local children can receive a later steering message. A provider may also seed the child with a prefix of the parent's log. That seed is optional and is the wrong default for this plan.

### Reasonix

Reasonix here is the community DeepSeek CLI at `esengine/DeepSeek-Reasonix`, not a `deepseek-ai` repo. The opened page is [SUBAGENT_PROFILES.md](https://github.com/esengine/DeepSeek-Reasonix/blob/main-v2/docs/SUBAGENT_PROFILES.md).

A profile is a manual skill with `runAs: subagent`. The child gets the profile prompt and the task. The parent keeps the task and the child's final answer, not the child's working context. A review or security-review child also gets a compact facts pack: confirmed decisions, an evidence summary, and file anchors. `write_paths` names the files that child may write; omitted paths become a whole-workspace claim. `read-only: true` strips writer tools. `allowed-tools` is an allowlist and cannot bypass permission or sandbox deny rules. `try` is the read-only runner. The same call can also name a model and an effort. That model pick is not this mesh's job.

### Qwen

The `QwenLM` org page lists `qwen-code` (terminal coding agent), `Qwen-Live-Harness` (realtime audio and video shell over a background harness), and `Qwen-MM-Plugins`. This note uses the code CLI and the live harness. Qwen-Agent stays with the Alibaba agents note.

[Approval mode](https://github.com/QwenLM/qwen-code/blob/main/docs/users/features/approval-mode.md) is a session switch: Plan (read-only, shell not run), Ask Permissions, Auto-Edit, Auto, YOLO. Plan is entered by `/plan` or the mode cycle, not by the wording of a planning request. `permissions.deny` still blocks before the Auto classifier. A deny on disk is not rewritten when Auto strips an over-broad allow for the session.

[Multi-agent coordination](https://github.com/QwenLM/qwen-code/blob/main/docs/users/features/multi-agent-coordination.md) is experimental. `/coordinate` gives investigation workers a read-only tool set. One writer may sit in a leader-owned worktree. The leader is the only merge authority. Teammates share a task list and can message each other. If the team flag is off, the same command falls back to foreground agents that report only to the leader. Arena is several models on one task with a winner. That is a model contest, not a fetch.

[Qwen-Live-Harness](https://github.com/QwenLM/Qwen-Live-Harness) is a desktop daemon plus a host. The daemon talks to Qwen Omni and may hand work to Qwen Code, Qoder, Codex, Claude Code, or Gemini CLI. File changes require that background harness. The app does not take over an arbitrary terminal. Memory libraries are local; consolidation and embeddings are cloud calls. Proactive monitors watch audio or video and are documented as not a safety alarm.

### Kimi and Moonshot

`MoonshotAI/kimi-code` is the current coding CLI. The archived `kimi-cli` points at it. The opened agent guide is [Agents and Sub-Agents](https://github.com/MoonshotAI/kimi-code/blob/main/docs/en/customization/agents.md) (raw body; the docs host shell was empty).

The main agent passes a task description. The child cannot see the parent history. Only the final result returns. Built-in `explore` does not modify files. Built-in `plan` has no shell. The three built-ins cannot dispatch further children, so a chain stops unless a custom agent opts in. Dispatch is three stages: dispatch, approval, collection. An always-allow rule the parent has accepted is copied onto every child. Agent frontmatter requires `description`; `name` may default to the file name. `tools` and `disallowedTools` are checked again before execution and are not the same control as permission rules. Each child has its own `wire.jsonl` under the session directory. That file is a debug trace, not the section store.

The GitHub README also names skills, lifecycle hooks, and ACP. Those guides were not opened, so they are not used as evidence.

### Cursor

`cursor/plugins` is a marketplace of standalone plugins. Each plugin may carry `skills/`, `rules/`, and `mcp.json` as separate trees ([README](https://github.com/cursor/plugins/blob/main/README.md)). Rules are not skills.

[Orchestrate](https://github.com/cursor/plugins/blob/main/orchestrate/skills/orchestrate/SKILL.md) fans a goal across cloud agents. Planners publish tasks and do not code. Workers do not talk to each other. Each level sees only its children's handoffs. A script and `plan.json` own the spawn and wait loop so a long agent loop is not the store. The kickoff command takes a `--model` flag. That flag is a model choice and stays out of this mesh.

### xAI Grok Build

`xai-org/grok-build` is the coding-agent harness and TUI. Opened guides: [skills](https://github.com/xai-org/grok-build/blob/main/crates/codegen/xai-grok-pager/docs/user-guide/08-skills.md), [hooks](https://github.com/xai-org/grok-build/blob/main/crates/codegen/xai-grok-pager/docs/user-guide/10-hooks.md), [project rules](https://github.com/xai-org/grok-build/blob/main/crates/codegen/xai-grok-pager/docs/user-guide/12-project-rules.md), [subagents](https://github.com/xai-org/grok-build/blob/main/crates/codegen/xai-grok-pager/docs/user-guide/16-subagents.md), [plan mode](https://github.com/xai-org/grok-build/blob/main/crates/codegen/xai-grok-pager/docs/user-guide/19-plan-mode.md).

A skill is a `SKILL.md`. `description` decides automatic load. The body is inlined only when it applies, and only the first 25,000 tokens. `name` and `description` may be omitted and then defaulted. That default is weaker than an explicit Function.

Hooks can deny a tool (`PreToolUse`), block a prompt (`UserPromptSubmit`), or refuse a stop. A block reason is shown to the person and is not added to model context, the on-disk session, or the session summary. After a restart the blocked bubble is gone because nothing was stored. Fail-open is the default: only an explicit deny blocks. A schema-invalid rewrite also blocks.

Plan mode writes `plan.md` and rejects every other edit until the person approves. The mode flag is stored apart from the plan file and survives restart. Compaction keeps a reminder that the mode is on. An empty or missing plan still opens the approval surface. Quit abandons the plan and turns the mode off. Revision notes that the person sends are delivered into the live plan. Subagents are not covered by the parent's plan-mode edit gate: a write-capable child starts with plan mode inactive and can edit while the parent cannot.

A child is spawned with `prompt` and a short `description`. It has its own context. `explore` and `plan` cannot edit. Only the top session may spawn; depth is one. `steer` joins the current turn at the next safe point. `queue` waits for a later turn. The child view cannot change the parent's permission mode.

Project rules (`AGENTS.md` and `rules/*.md`) are standing instructions loaded into context. They are not a steer line and not a section field.

## Best practices

The section record cannot grow past Section, Function, Read, Do, Write, Next, and Rail. These vendors never name that record. The mapping is this note's suggestion.

**Before a store write or the first tool.** DeepSeek will not open a step on an empty or rejected claim, and it will not commit the prompt or the user text for that failed admit. Map the claim onto Section, Function, Read, and Do: all four non-empty. Rail may stay empty; an empty switch is not an empty claim. Next may stay unset until the seat answers; it is not a precondition of the first tool. Write must be the one mesh key when the action is the store itself. A read-only first tool does not invent a second write key. Grok's empty `plan.md` still opens approval and still blocks every other edit, which matches a packet that can be judged before a corpus write. Do not copy Grok's fallback that invents `description` from the first paragraph. Function has to be the explicit sentence.

**When a lead cuts NEXT.** The lead returns the reject and does not write the corpus. DeepSeek's driver, not the rejecting listener, removes the claimed batch and closes a durable turn that spent no step. That closed turn is the cut's record. It is not a clean stop and it is not the next section. Grok's blocked prompt is harsher: the reason is shown and then stored nowhere, so a restart cannot fetch it. This plan already wants a refuse row rather than silence. Prefer the DeepSeek closed turn, stored by the mesh, over a cut that leaves no record. One page does not decide that. A second note has to agree, and no lead has cut it.

**What a parent hands a child.** Kimi hands only the task description and takes back only the final answer. The child cannot see the parent history, and a built-in child cannot dispatch another child. Reasonix adds a compact facts pack for a review child and a `write_paths` list that is the child's own write set. Map that onto the fixed fields: the child packet may carry Section, Function, Read, and Do for its own job. It may not write the parent's Write key, may not set the parent's Next, and may not flip the parent's Rail. `allowed-tools` does not bypass the parent's permission or sandbox. Qwen's read-only teammates match the same write ban, and the leader stays the only merge. DeepSeek's optional parent-log seed, Kimi's copy of always-allow rules, and Grok's plan-mode hole (a child editing while the parent is gated) are the three ways a child touches something it must not.

**Steer, side record, and Rail.** They are three objects.

- The steer line is a message sent into the live step. Grok `steer` joins the current turn. Plan-mode revision notes are delivered the same way. The next fetch of that live step may include it. The in-flight action is not cancelled.
- The human rail side record is the approval surface and a block reason. Grok shows the block reason to the person and does not put it in model context or the session record. A seat packet cannot write that record. The next fetch must not include it.
- The section Rail field is the switch, empty or flipped only by the person. Grok persists plan-mode state apart from `plan.md` and may put a short reminder in the next compacted context. The next fetch may include that switch. It may not include the side-record prose, and it may not treat `AGENTS.md` or a skill body as the Rail field.

## Adopt

- The DeepSeek pre-step gate: empty or rejected claim, no step, no corpus commit. Mesh stores the closed turn. The lead does not.
- The child packet from Kimi, tightened by Reasonix: task text in, final answer out, review child gets a compact facts pack, writers stripped when read-only, allowlist is not a bypass, built-in children do not spawn.
- The three-way split from Grok plan mode and hooks: steer enters the live step, the block reason stays on the human side, the mode flag is a separate switch the next fetch may remind.
- Qwen Plan as a read-only mode that a sentence alone does not enter, and `permissions.deny` ahead of any classifier.
- Cursor's split of skills, rules, and MCP into different trees, and Orchestrate's rule that a planner publishes tasks and does not code, while a worker returns one handoff and does not see its siblings.
- Skill discovery that loads a body only when the description matches. That is already the plan. This pack only repeats it.

## Refuse

- Any vendor path that chooses the model: DeepSeek `agentOptions`, Reasonix `model` on `task`, Grok's spawn `model` argument, Qwen Arena's winner, Cursor `--model`, Qwen Live handing the job to whichever background CLI is installed. The mesh does not choose a model. One seat keeps its direct id.
- Seeding a child with the parent's log, copying the parent's always-allow rules onto the child, or letting a child edit while the parent's plan gate is on.
- A shared team task list the child can write. That is the parent's Next and the parent's task store.
- Grok fail-open hooks as the rail. A crash or a timeout must not let the tool run. This plan's gate fails closed.
- A blocked prompt that is stored nowhere. The cut needs the closed-turn record above, or the existing commit-refuse row. Silence is not a cut.
- Defaulting a missing Function or description from the first paragraph, or from a file name.
- Continual-learning style rewrite of `AGENTS.md`, and any plugin that overwrites the corpus. Leads do not write the corpus. Seats do not delete.
- Qwen Live, Herdr, or a second coding CLI as the loop. They are other harnesses. File writes stay behind the human rail.
- Treating project rules, skill bodies, and the section Rail as one object.

## Park

- DeepSeek agent teams, Qwen Agent Team, and Cursor Orchestrate's cloud script, Slack mirror, and Andon reaction. Useful later as a spectator or a scheduler. Not a step on this packet. Not a second store.
- Grok worktree isolation and Qwen's single leader-owned worktree. Later isolation, after the packet gate exists. Not a socket on a guest, and not a model router.
- Reasonix richer skills (`references/`, `scripts/`) and Grok personas with input and output contracts. Later, when a child packet is already fixed. Do not grow the section record to hold them.
- Qwen Live proactive monitors and local memory libraries. Later desk features. Not the fetch.
- Plugin marketplaces (Cursor, Grok, Kimi). Catalog discovery is already decided. Installing a marketplace is a later loop.
- Hook shells that format or notify after a tool. Later. The only hook shape this packet needs is the fail-closed gate before the tool and before the store.
- AirLLM. Optional later disk-stream inference. Not this loop.
