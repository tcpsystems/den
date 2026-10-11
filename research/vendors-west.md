# Vendors west

Anthropic, OpenAI, Google, Vercel, and Cloudflare. Skills, guides, guards, and harnesses only. This note answers four holes in the existing harness plan. It does not open a second plan. One note stays a suggestion until a second finished note states the same scaffolding claim and no lead cuts it.

Pages opened for this note:

- https://github.com/anthropics/skills/blob/main/README.md
- https://github.com/anthropics/skills/blob/main/spec/agent-skills-spec.md which only points at the spec
- https://agentskills.io/specification
- https://github.com/openai/openai-agents-python/blob/main/docs/guardrails.md
- https://github.com/openai/openai-agents-python/blob/main/docs/handoffs.md
- https://github.com/openai/openai-agents-python/blob/main/docs/human_in_the_loop.md
- https://github.com/openai/openai-agents-python/blob/main/docs/tools.md
- https://github.com/google/adk-docs/blob/main/docs/callbacks/types-of-callbacks.md
- https://github.com/google/adk-docs/blob/main/docs/agents/custom-agents.md
- https://github.com/vercel/ai/blob/main/content/docs/03-ai-sdk-core/15-tools-and-tool-calling.mdx
- https://developers.cloudflare.com/agents/runtime/lifecycle/state/
- https://developers.cloudflare.com/agents/runtime/communication/readonly-connections/

The section record stays Section, Function, Read, Do, Write, Next, and Rail. Nothing below adds an eighth field. The mesh stays the bus, the store, and the fetch. It does not choose a model.

## What it is

Anthropic's skills repo is a folder of `SKILL.md` packs plus a one-line pointer. The README says a skill is a folder with YAML frontmatter and instructions, and "the frontmatter requires only two fields": `name` and `description`. The spec file in that repo now says the format lives at agentskills.io. That spec says `name` and `description` are required, `description` must be non-empty and at most 1024 characters, and `license`, `compatibility`, `metadata`, and `allowed-tools` are optional. The body has no required shape. Loading is three steps: metadata at startup, the `SKILL.md` body when the skill is activated, and `scripts/`, `references/`, or `assets/` only when needed.

OpenAI's Agents SDK is a harness of agents, tools, handoffs, guardrails, sessions, and human approval. Input guardrails run on the first agent. Output guardrails run on the agent that produces the final output. Tool guardrails run on every guarded function-tool call, before execution and again after it. A blocking input guardrail finishes before the agent starts, so a tripwire means the agent never executes tools. An output tripwire rejects the candidate final output. The runner asks the session to keep already-completed tool items and a data-free placeholder, and it does not retain the rejected payload. Handoffs transfer the conversation. Agents-as-tools do not. Human approval is a separate `RunState` of pending tool calls, not a field of the tool result.

Google ADK puts a callback on each stage. `before_agent_callback` can skip the agent before core logic. `before_tool_callback` runs after the model has named a tool and before `run_async`. Only a `None` return lets the tool run. A returned dict, including `{}`, skips the tool and becomes the tool result. `AgentTool` runs another agent inside the parent's flow, captures the final response, returns that response as the tool result, and also forwards state and artifact changes into the parent context. A custom agent's sub-agents share `ctx.session.state`, and a sub-agent's `output_key` is how a later step usually stores a result.

Vercel AI SDK tools have an `inputSchema` that both the model and the SDK use to validate a call. `description` on a tool is optional. `execute` is optional because a call can be forwarded instead of run. `toolApproval` is a separate decision from the schema and from `execute`. `'denied'` records approval request and response parts, then surfaces a denied tool output. `'user-approval'` returns a `tool-approval-request` and does not execute on that call. The manual path is two model calls: the first asks, the second either executes or tells the model the denial.

Cloudflare Agents persist JSON state in the agent's own SQLite through `setState`. `initialState` applies only when nothing is stored yet. With no `initialState`, `this.state` is `undefined`. `validateStateChange(nextState, source)` runs before persistence and broadcast. Throwing aborts the update. `onStateChanged(state, source)` is only a notification. `source` is `"server"` or a client `Connection`, and it is not a field of `state`. A readonly connection receives state broadcasts and may call RPC that does not write, but `setState` throws `Error("Connection is readonly")`. The readonly flag lives on the connection attachment. `connection.state` never exposes it and cannot overwrite it.

## Best practices

These are the four scaffolding lines a seat could adopt. Each one is still a suggestion from this note alone.

**Packet gate.** Before a store write or the first tool, Section and Function must be non-empty. Read, Do, Write, Next, and Rail may be empty. An empty tool gate is a refusal, not permission to write or to run the tool.

Anthropic and the spec it points at say the same thing: the record is not a skill unless `name` and `description` are present, and `description` must be non-empty. Optional frontmatter and the body are not required for that identity. OpenAI says a callable approval rule is not invoked when tool arguments are missing, empty, whitespace, malformed, or not an object, and the call then requires manual approval. A hosted sticky approval is persisted only when both identity fields are non-empty. Google says an empty dict from `before_tool_callback` skips the tool and `{}` becomes the result. Return `None`, not `{}`, or the tool does not run. Cloudflare will persist empty arrays and nulls inside a state object that passes `validateStateChange`, and it will persist nothing if that hook throws or if `setState` is never called. Emptiness of Read, Do, Write, Next, or Rail is therefore legal inside the seven fields. It does not authorize the first tool, and it does not authorize a write that the validator rejects.

Map the names and stop. Section is `name`. Function is the one-sentence `description` of what this section does and when it runs. Read can hold the optional `allowed-tools` list or the file list, and it may be empty. Do is the body, loaded when the section is activated, and it may be empty at the gate. Write is the one mesh kb key and may be empty until a real store. Next may be empty until a later step fills it or a lead cuts it. Rail may be empty. Do not add `license`, `compatibility`, or `metadata` as section fields.

**Cut NEXT.** When a lead cuts NEXT, the mesh keeps the previous section, the lead writes no corpus, and the cut leaves a data-free refusal in place of the rejected next payload. The store writes that refusal. The lead does not.

Cloudflare: `validateStateChange` runs before persistence, and throwing aborts the update, so the rejected next state is not saved and the previous state remains. The hook is not a corpus writer. OpenAI: an output tripwire "rejects the candidate final output." The runner persists completed tool items plus a placeholder such as "Output withheld by an output guardrail." and does not retain the original rejected payload in the session or in `RunState`. A blocking input tripwire raises and the agent never starts. Vercel: `'denied'` emits the approval request and the approval response, includes the reason on that response, and does not execute the tool. Google: the tool's `run_async` is skipped. The dict the callback returns is the tool result, and the tool itself does not write.

If the lead does not write the corpus, the record left behind is the previous section plus that refusal. The refusal names the cut and withholds the rejected NEXT payload. It is not a new section, not a clean stop, and not a Write the lead authored. The mesh is the writer because the mesh is the store. A lead that only judges does not get a kb key.

**Parent and child.** A parent hands a child a structured input for the child's own Section, Function, Read, and Do. The child may not write the parent's Write, may not set the parent's Next, and may not flip the parent's Rail. The child returns a result. The parent keeps those three.

OpenAI's tools guide says agents-as-tools exist so a central agent can orchestrate specialists "instead of handing off control." The nested run "is not inherited automatically" from the parent conversation. The default child input is one string field, `input`, or a schema the parent declares with `parameters`. A `custom_output_extractor` returns a value to the central agent. The same guide's handoff chapter is the opposite: a handoff makes it "as though the new agent takes over the conversation, and gets to see the entire previous conversation history," and `input_type` "does not replace the next agent's main input." Input guardrails do not run again on the receiving agent. Cloudflare's readonly connection receives the current state and may call a method that does not write, but any `setState`, including one hidden inside a callable, throws. The child can read. The child cannot persist. Google's `AgentTool` does return the child's final response as a tool result inside the parent flow, which matches the return half. The same paragraph says the framework "forwards any state/artifact changes back to the parent's context," and the custom-agent page shares `ctx.session.state` across sub-agents, with `output_key` writing into that shared map. That forward is how a child would touch the parent's Write, Next, or Rail. Do not take that half.

**Three rail objects.** The steer line, the human rail side record, and the section Rail field are three objects. The next fetch may include the section, and therefore the Rail field. It may not include the steer line as a section field, and it may not include a writable copy of the human side record.

Cloudflare keeps the persisted object, the writer, and the permission flag apart. `this.state` is what a later read returns. `source` arrives beside it. The readonly flag is on the connection attachment, survives hibernation, and is hidden from `connection.state`. OpenAI keeps the pending call on `RunResult.interruptions` (`tool_name` and `arguments`), and keeps the human decision on server-owned `RunState` via `approve` or `reject`. The guide says the server must not accept replacement tool calls, arguments, approval records, or serialized state from the client, and must not take the reviewer's identity from the request body. A rejection message can be shown back into the run. That message is not the approval record and not the section. Vercel keeps `toolApproval` apart from `inputSchema` and from `execute`. A denial is an approval-response part plus a denied tool output. A manual steer emits `tool-approval-request` and waits. `'not-applicable'` means there is no approval object at all.

So the next fetch of a stored section includes Rail, which may be empty. The steer line stays on the live step until a human resolves it. It is not Next and it is not fetched as the next section. The human side record stays outside the seven fields. The seat packet cannot write it. A read-only notice that a cut happened may travel with the refusal row from the cut rule above. That notice is not a second Rail and not a license for the seat to flip the switch.

## Adopt

Adopt the four lines in Best practices. They fit the seven fields and the rules already in the plan: the packet is handed over, a lead may cut NEXT before it is stored, leads do not write the corpus, a child has things it cannot touch, and the human rail is a side record a seat packet cannot write.

Adopt OpenAI's blocking order for the first tool. The check finishes before the tool runs. Parallel guardrails that let the tool start and then cancel it are the wrong shape, because the plan's in-flight steer does not cancel a step that has already begun, and a tool that already wrote is no longer a gate.

Adopt Cloudflare's "throw before persist" as the store half of a cut, and OpenAI's data-free placeholder as the record the cut leaves. Together they say who stores it: the mesh stores the previous section and the refusal, and the lead stores neither the rejected NEXT nor the corpus.

Adopt agents-as-tools, not handoffs, as the parent-child shape. The parent remains the seat that owns Write, Next, and Rail.

Adopt Vercel's split of schema, approval, and execute as the reason the three rail objects stay three. Schema validation is the packet gate. Approval is the human side record. Execute is the tool. None of them is the section Rail field.

## Refuse

Refuse a handoff that transfers the conversation. The child would see the parent's history, including tool writes, and would become the agent that produces the final output. That is the child setting Next and touching Write. OpenAI also says tool guardrails do not apply to the handoff call itself, so the gate the plan wants on every tool would not wrap the transfer.

Refuse Google's shared `ctx.session.state` and the `AgentTool` forward of state and artifact changes into the parent. The child may return a result. It may not write the parent's Write key, choose the parent's Next, or flip the parent's Rail. A sub-agent `output_key` on the shared map is a hidden Write.

Refuse treating an empty dict, empty tool arguments, or a missing `initialState` as permission. Google turns `{}` into a skip. OpenAI will not run the approval callable on empty arguments. Cloudflare leaves `this.state` undefined until a real `setState`. The first tool waits for an explicit allow, which is `None` in ADK and a passed guardrail or approval in the other two. The store waits for a write that passes validation.

Refuse growing the section past the seven fields. The skills spec allows `license`, `compatibility`, `metadata`, and `allowed-tools`. Those stay out of the record. `allowed-tools` can be expressed inside Read. The others are not Section, Function, Read, Do, Write, Next, or Rail. Cloudflare's `source` and readonly flag stay outside the fetched state for the same reason. OpenAI's `RunState` approval record stays outside it too.

Refuse letting the mesh pick a model. Cloudflare's state guide shows an agent method that both reads SQL history and calls a named model. That call sits in the agent, which is the seat, not in `setState`. Copying it into the store would make the bus a router. The mesh fetches and stores. The seat already has one direct model id.

Refuse a client-supplied approval record. OpenAI says a schema check is not authentication, and the server must apply `approve` or `reject` only to items it already stored. Vercel emits the approval response from the SDK. A seat cannot write the human side record into the packet and then fetch it back as if it were Rail.

Refuse Cloudflare's caveat that a readonly callable still runs side effects that happen before `setState`. A child that must not touch the parent's Write must be stopped before the side effect, not after the email or the charge. The check belongs before the first tool and before the store write.

Refuse Vercel's second model call as the way a cut is repaired. After a denial, that SDK "informs the model that approval was denied" and lets the same loop continue. A lead cut is not a hint the seat uses to author a new NEXT. The refusal row is the end of that proposal.

## Park

Park as a later loop any vendor path that edits the harness, the skill body, or the routing after the commit. The skills spec's on-demand `scripts/` and `references/` load is already the decided skill rule: one body loads when the description matches. Do not reopen it. Do not build a curator that rewrites skills from traces on this packet.

Park OpenAI's `always_approve` and `always_reject` sticky decisions. They survive for the rest of one run and are serialized inside `RunState`. They are not the section Rail field, and they are not a human switch that outlives the run. A later loop can decide whether a human Rail should default the next section. This packet does not.

Park Vercel's `stopWhen` and step-count loop as a harness implementation. The plan already says the seat makes one call, reads the result, and repeats until a final message. The SDK's multi-step driver is that loop inside a vendor library, not a new seat and not a model choice for the mesh.

Park Google's pattern of handing the skip dict back to the same model as a tool result the model may answer. It is useful telemetry. On this harness the lead cut is stored by the mesh as a refusal and is not a prompt for the seat to write the corpus or the next id.

Park Cloudflare workflow helpers that `mergeAgentState` or `resetAgentState` from a background step. A later scheduled loop might append a section. It is not the live packet, and a reset would fight the rule that nothing is deleted and a later section supersedes.
