# Alibaba AgentScope, Qwen-Agent, Tongyi DeepResearch

One source, 2026-10-11. This note feeds the harness plan. It is not a second plan, and it does not change decided lines. Nothing here is decided until another finished note states the same scaffolding claim and no lead cuts it.

The mesh stays the bus, the store, and the fetch. It does not choose a model. The section record stays Section, Function, Read, Do, Write, Next, and Rail. These pages do not add a field.

Pages opened for this note:

- https://github.com/agentscope-ai/agentscope
- https://docs.agentscope.io/en/versions/2.0.10dev/building-blocks/agent/overview
- https://docs.agentscope.io/en/versions/2.0.10dev/building-blocks/agent/interrupt-agent
- https://docs.agentscope.io/en/versions/2.0.10dev/building-blocks/agent/human-in-the-loop
- https://docs.agentscope.io/en/versions/2.0.10dev/building-blocks/sop
- https://docs.agentscope.io/en/versions/2.0.10dev/building-blocks/pipeline/team
- https://docs.agentscope.io/en/versions/2.0.10dev/building-blocks/message-and-event
- https://docs.agentscope.io/en/versions/2.0.10dev/building-blocks/permission-system/overview
- https://docs.agentscope.io/en/versions/2.0.10dev/building-blocks/middleware
- https://github.com/QwenLM/Qwen-Agent
- https://qwenlm.github.io/Qwen-Agent/en/guide/
- https://raw.githubusercontent.com/QwenLM/Qwen-Agent/main/qwen_agent/agent.py
- https://raw.githubusercontent.com/QwenLM/Qwen-Agent/main/qwen_agent/agents/fncall_agent.py
- https://raw.githubusercontent.com/QwenLM/Qwen-Agent/main/qwen_agent/llm/fncall_prompts/nous_fncall_prompt.py
- https://raw.githubusercontent.com/QwenLM/Qwen-Agent/main/qwen_agent/tools/base.py
- https://github.com/Alibaba-NLP/DeepResearch
- https://github.com/Alibaba-NLP/DeepResearch/tree/main/inference
- https://github.com/Alibaba-NLP/DeepResearch/tree/main/Agent
- https://raw.githubusercontent.com/Alibaba-NLP/DeepResearch/main/inference/react_agent.py
- https://raw.githubusercontent.com/Alibaba-NLP/DeepResearch/main/inference/prompt.py
- https://raw.githubusercontent.com/Alibaba-NLP/DeepResearch/main/FAQ.md

## What it is

AgentScope 2.0 is a Python agent framework. The overview calls `Agent` a stateless reasoning-acting loop: it accepts messages or events, calls tools, can force a structured result, and can pause and resume. Around that loop it documents a toolkit, a permission check on every tool call, middleware, a message distinct from an event, a team pipeline, and an experimental standard operating procedure. The service features named on the repository home (channels, a skill hub, schedules, realtime voice) are a product around the loop. They are not the loop.

Qwen-Agent is a Python framework aimed at Qwen. An `Agent` receives a message list and either calls a model or a tool. `FnCallAgent` repeats that until a model turn uses no tool, or the call budget runs out. The default tool text asks for one JSON object with a function name and arguments inside `<tool_call>` tags. A tool result is wrapped as `<tool_response>` and fed back as a user message. The same repository also ships MCP, a Docker code interpreter, RAG, and a Gradio UI. The guide page repeats that purpose and shows an `Assistant` plus `WebUI`. It does not define a section record.

Tongyi DeepResearch is a 30.5B research model and an inference script, not a second harness. The repository home says inference can follow ReAct, or an IterResearch "Heavy" mode that scales at test time. The `inference/` tree that was opened contains `react_agent.py`, `prompt.py`, and the tool files. It does not contain a Heavy-mode runner. `MultiTurnReactAgent` starts from a system prompt plus the question, calls the model, runs at most one tool call when the text contains `<tool_call>`, appends `<tool_response>`, and stops when the text contains `<answer>`. The `Agent/` tree only shows the directory names `AgentFounder` and `AgentScaler`. Those names were not opened. The FAQ only points at local deployment or the hosted Bailian app.

None of the three chooses the seat's model for this harness. AgentScope's `ModelRouterMiddleware` does pick a chat model from the latest user text and swaps `agent.model` for that reply. That picker stays off this mesh.

## Best practices

### What must be non-empty before a tool or a store

The seven fields do not all have to be filled before the first tool. They also do not all have to be filled before every store.

Qwen's `BaseTool._verify_json_format_args` runs before the tool body. A string that is not JSON raises `Parameters must be formatted as a valid JSON!`. A parameter marked `required` and missing from the object raises `Parameters %s is required!`. `is_tool_schema` separately demands that the tool definition itself have a non-empty `name`, a `description`, `parameters`, and a `required` list. The Nous template says the call is `{"name": <function-name>, "arguments": <args-json-object>}`. DeepResearch uses that shape in `prompt.py` and, in `react_agent.py`, refuses to call a tool when the JSON does not contain `name` and `arguments`. The failed text is still appended as a `<tool_response>` user message: `Error: Tool call is not a valid JSON. Tool call must contain a valid "name" and "arguments" field.` The tool function does not run. An empty model body is retried inside `call_server` and is not treated as a successful turn.

Mapped onto the section, that gate sits in Do. Do is non-empty at the first tool only when it holds valid JSON arguments and every argument that tool marks required. Write, Next, and Rail are not read by this check. Read is not universally required. In the DeepResearch prompt, `search` and `google_scholar` require `query`, `visit` requires `url` and `goal`, `parse_file` requires `files`, and `PythonInterpreter` has an empty `required` list with the code placed after the JSON. An empty Read is valid when the tool's own required list is not a file list. Qwen's file helper copies files only `if files`; an empty file list skips the copy rather than inventing a path.

The step's identity is a different gate, and it comes from the AgentScope SOP, not from the tool schema. An `SOPStep` requires `subject` (a short name) and `description` ("What the step must achieve: the destination, not the route"). The description tells the executor the destination. The route stays with the executor. Section and Function are those two strings, and they exist before the step runs. They are not the tool's `name` and `description`, and they must not be replaced by them.

A store that advances Next is stricter than the first tool. The SOP engine gives the executor the step description and whatever the previous step handed over. The executor then returns a `handover` through structured output, "the account the following steps will read." The refusal table says a reply that ends without a handover has "No structured output, so nothing was handed on," and that case spends an attempt. "Only handovers cross between steps; files and conversations do not." So Write must be non-empty before a store that sets Next. Until that handover exists, Write and Next may stay empty while tools run. Rail may stay empty the whole time: `verifier` defaults to `None`, and then "the step passes as soon as the executor hands over."

A bad tool call is not that store. DeepResearch keeps the error inside the message list and continues the loop. It does not open a new section for the error.

### When a lead cuts NEXT

A cut of NEXT is a refused handover, not an interrupt of the live tool, and not a write by the lead.

The SOP page says judging belongs to the verifier, and "do not write it as a step of its own." The verifier sees the description, the step input, and the handover, and answers with `passed` and `message`. A refusal "clears the handover and sends `message` back to the executor." The step goes back to `PENDING`. If refusals reach `max_attempts`, the step and the run are marked `FAILED`. `SOPRunState` holds "the run's inputs and each step's phase, handover, and verdicts." The page says developers can store that state and rebuild the engine from it. The verifier is not the writer. In this harness the mesh is that store. The lead corresponds to the verifier: it can cut the proposed next, and it does not write the corpus.

The record the cut leaves is the pending run state: the inputs, the phase `PENDING` (or `FAILED` only after the attempt budget), an empty handover, and the verdict message. Next is not stored. The empty handover is not a new section. It is why the next section does not start.

`UserInterruptEvent` is a different act. On the SOP page it "abandon[s] the parked call. The attempt is dropped without counting as a refusal, and the step goes back to `PENDING`." The interrupt page says a paused agent then discards the pending confirmation, synthesises a `ToolResultBlock` marked interrupted for each pending tool call, emits `ReplyEndEvent`, and is ready for new input. That writes an interrupted tool result onto the current reply. It does not leave a verdict on an unstored NEXT. This harness already says a steer line lets the in-flight action finish. A NEXT cut happens before the successor is stored. Use the refusal record for that cut. Do not use the interrupt record for it.

DeepResearch agrees about who writes, and disagrees about the shape of the leftover if copied naively. `_run` always returns a dict with `question`, `answer`, `messages`, `prediction`, and `termination`. The `answer` value was read from the input item; the README says that field is the ground-truth reference for evaluation, not the model's write. On a miss, `prediction` is `No answer found.` and `termination` is `answer not found` or `exceed available llm calls`. The function returns that dict. It does not write a corpus. The useful part is the named leftover. The five keys must not become section fields. The mesh already has an append-only progress line. A cut fills that line's state with the verdict and leaves the next id unset.

### What a parent hands a child

`TeamPipeline` gives the child one prompt, and nothing else.

The leader calls `TeamAssign` with `member` and `prompt`. The page says `prompt` is "the complete task for the member" and "Members cannot see the leader's context, so everything the task needs must be written here." The pipeline "hands `prompt` to the member as a user message." The member works in its own context. "The leader receives only each member's final reply, never its intermediate reasoning or tool calls." That reply is "text and data blocks." Several assignments in one round return in call order. Members never talk to each other. `reset_members` defaults to true: after the leader reply ends, member context is cleared.

A user message, on the message page, can contain only text and data. It cannot carry the parent's tool calls, tool results, or hints. Write, Next, and Rail are therefore not in the child packet unless a parent copies them into the prompt. This harness does not copy them. The child may not read or write the parent's Write, the parent's Next, or the parent's Rail. The child returns a final reply. The parent may later put that reply into the parent's own Write. The child does not set the parent's Next and does not flip the parent's Rail.

`TeamAssign` "skips permission checks." The same sentence says the member's own tools "are still checked under their own permissions." A human confirm for a member is routed back to that member by `reply_id`, and the leader waits on that assignment. The skip on the assignment itself does not fit this harness. Every tool call here is gated, including the handoff. The child's own checks stay.

Qwen's base `Agent` stores a `description` "which will be used for multi_agent," and `FnCallAgent` keeps a `Memory` object for files. Neither of those was a child packet on the pages opened. The child packet to adopt is the TeamAssign prompt, not Qwen's file memory and not the parent's message list.

DeepResearch's `visit` tool is not a child seat. Its schema requires `url` and `goal`, and `EXTRACTOR_PROMPT` then asks for JSON with `rational`, `evidence`, and `summary`. That is a tool result. It is not a second section, and those three words are not new section fields.

### Steer, the human side record, and Rail

They are three objects. The next fetch of the section includes only the section's Rail from among them.

The message page separates them. A `Msg` is "the unit of inter-agent communication and persistence." A user message holds text and data only. A `HintBlock` is "hint information used to guide the LLM." When it is sent to the model it is converted into a user message, and the page says to wrap it in tags "to avoid confusion with user inputs." A `HintBlockEvent` is persisted and replayable. A `RequireUserConfirmEvent` is not stored as message content; reconstructing the reply "updates tool call states to `ASKING`." `finished_reason` of `None` means the reply is paused, not finished.

The human-in-the-loop page makes the confirm event an input of its own. The agent pauses, emits `RequireUserConfirmEvent`, and `reply` returns a waiting message with `finished_reason` `None`. Resume is a `UserConfirmResultEvent` carrying `confirm_results` for that `reply_id`. A denied call "produce[s] an error result visible to the LLM." Accepted suggested rules "are persisted to the permission engine," which is not the reply body. `AskUser` writes `output` for the model and the same answers into `metadata` "so a program can branch." Those are two shapes of one human answer, and neither is the section.

The permission overview is the switch. A mode is "a global policy" "set at configuration time, switchable at runtime." Rules are a second component. Tool checks are a third. `EXPLORE` denies modifications. `DONT_ASK` turns an ask into a deny when nobody is there to answer. That mode is the Rail analogue: empty, or a switch only the human flips. It is not the hint, and it is not the confirm event.

The team and SOP pages list `reply_stream` inputs as different types: a message that starts work, a confirm result, an external execution result, and `UserInterruptEvent`. A steer line is the hint on the live reply. It is not a new section, it is not a confirm, and it is not an interrupt. The plan already says a steer does not cancel the run. An interrupt ends the parked reply. Do not fetch those as the same object.

What a later fetch may include:

- A replay of the current reply may include the persisted hint and the tool state, including `ASKING`, denied, or interrupted. That replay is the live reply, not the next section.
- `SOPRunState` includes inputs, phase, handover, and verdicts. The page says it "covers only the procedure's own state." Agent context is saved separately. The confirm event is not a field of that state. The permission mode is not a field of that state.
- This harness already stores Rail on the section. The next section fetch therefore includes that Rail field, because it is part of the section record. It does not include the steer hint and it does not include the confirm event or the permission-engine rules. Those stay side records. A seat packet still cannot write the human side record.

## Adopt

Adopt the tool gate on Do. Before the first tool, Do must be valid JSON and must include every argument that tool marks required. Otherwise the tool does not run. The error can sit on the current reply as the tool observation. It is not a new section. Section and Function stay the step's short name and one-sentence destination, already set before the step runs. Read may be empty unless that tool's required list is a file list. Write, Next, and Rail may be empty at the first tool.

Adopt the handover gate on Write. A store that would set Next requires a non-empty Write. No handover means nothing is handed on, and Next does not start. Files and the conversation do not cross the step boundary. Rail may remain empty when nobody has to judge the step.

Adopt the verifier cut. When a lead cuts NEXT, the lead writes no corpus. The mesh stores the current step as pending, with the handover cleared and the verdict kept. That record is not a successor section and not a clean stop. The in-flight tool is not cancelled by this cut.

Adopt the child prompt. The parent hands the child one self-contained user message. The child cannot see the parent's context, so it cannot touch the parent's Write, Next, or Rail. Only the child's final text and data come back. The child's tools stay under the child's own permission checks. Member context does not have to survive the parent's reply.

Adopt the three-way split. The steer line is a hint on the live reply. The human rail side record is the confirm event plus any rule the human accepts into the permission side store. The section Rail is the switch field on the section. The next section fetch may include that Rail field and must not include the hint or the confirm event.

## Refuse

Refuse `ModelRouterMiddleware`. The middleware page says it reads the latest user text, asks a router model to pick a candidate, and replaces `agent.model` for that reply. The mesh does not choose a model. One seat keeps the direct id it was given.

Refuse Qwen `Memory`, AgentScope `Mem0Middleware`, and `AgenticMemoryMiddleware` as the section store. `FnCallAgent` builds a `Memory` for files beside the message list. The middleware page lists long-term memory as another backend. This harness has one store. A later section supersedes. Nothing is deleted. Embeddings stay on the local embedder.

Refuse skipping permission on `TeamAssign`. The handoff is a tool call. It takes the same gate as any other tool. Keep the rule that the child's own tools are checked as the child's.

Refuse using `UserInterruptEvent` as a NEXT cut or as a steer. A steer enters the live step and the in-flight action finishes. An interrupt discards the parked call and writes an interrupted tool result. A NEXT cut leaves a pending step with an empty handover and a verdict, stored by the mesh.

Refuse growing the section. Do not add DeepResearch's `rational`, `evidence`, and `summary`. Do not add `prediction`, `termination`, `messages`, or the evaluation `answer`. Do not store the transcript on the section. A named miss stays on the progress line this plan already has.

Refuse letting the child, the verifier, or a lead write the corpus. The verifier "only judges, never writes." The child returns a final reply and does not set the parent's Write, Next, or Rail. A denied tool becomes an error result on the current reply, not a deleted section and not a new one.

Refuse Qwen group-host selection and AgentScope team membership as a model or department router. A host that picks the next speaker from context would choose work the mesh did not hand over. Departments stay on the seats. The parent names one child in the prompt because the section already named that child.

Refuse the code interpreter as the seat. The Qwen README says the Docker interpreter mounts one working directory and is still to be used with caution, and it says the math executor is not sandboxed and is for local testing only. Isolation for a section stays the rule this plan already has. No socket on a guest.

## Park

Park end-to-end reinforcement learning, the synthetic-data pipeline, continual pre-training, and the IterResearch Heavy mode. The DeepResearch home describes all four. The opened `inference/` tree implements the ReAct loop only. Heavy mode is test-time scaling, which is another loop after a packet exists, not a step on this packet. `AgentFounder` and `AgentScaler` were not opened; the directory names sit on the training side of that repository. Leave them there.

Park the SOP engine as a later runner for a procedure whose step count and owners are known before it starts. The SOP page says to use it only when both are knowable without running, and when someone actually checks the step. This council's next id is not that fixed list. The live packet stays one handed section, one tool call, then the result. The handover rule and the verifier cut can be adopted now without mounting the engine.

Park AgentScope's channels, schedules, realtime voice, RAG service, and skill hub. They are serving features on the repository home. They do not decide the section.

Park a second save of agent context beside `SOPRunState`. The SOP page already splits them. This harness does not need that split as a second corpus. The section fetch remains the mesh fetch.

AirLLM is not in this pack. It stays an optional later disk-stream inference slot, not the loop.
