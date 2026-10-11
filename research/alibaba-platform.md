# Alibaba platform

This note expands the harness plan. It is not a second plan, and it does not replace Seat expansion. One finished note stays a Suggestion. A line moves into the plan only when a second finished note states the same scaffolding claim and no lead cuts it.

The mesh stays the bus, the store, and the fetch. It does not choose a model. The section record still cannot grow past Section, Function, Read, Do, Write, Next, and Rail. AirLLM stays an optional later disk-stream inference slot, not the loop. Nothing below installs a runtime.

Pages opened for this note: the nine repository homes, Spring AI Alibaba's human-in-the-loop page, Open Code Review's architecture page, Lynxe's developer quick start, and the skill-up DeepSeek Harness plugin readme. Lynxe's `design.md` was opened and is a frontend store sketch, not the agent loop.

## What it is

**Spring AI Alibaba** (`alibaba/spring-ai-alibaba`) is a Java agent framework. The Agent Framework sits on a graph runtime that checkpoints long-running state. Built-in shapes are `SequentialAgent`, `ParallelAgent`, `RoutingAgent`, and `LoopAgent`. Context engineering is hooks and interceptors: human-in-the-loop, compaction, editing, model and tool call limits, tool retry, planning, and dynamic tool selection. Admin and Studio are a visual builder, tracer, and evaluator. A2A is offered through Nacos. The model is a caller the application picks. The graph does not have to be the model router, but `RoutingAgent` is a router if you mount it as one.

**AssistantAgent** (`spring-ai-alibaba/AssistantAgent`) is an enterprise assistant on that framework. It generates code and runs it in a GraalVM sandbox, instead of only calling a fixed tool list. An evaluation graph classifies intent before the prompt is built. Experience is one model with COMMON, REACT, and TOOL rows, convertible to and from a Skill, disclosed progressively, and fetched with `search_exp` then `read_exp`. FastIntent skips the model call when a stored experience matches. Learning extracts new experience from history. None of that is the section store.

**Lynxe** (`spring-ai-alibaba/Lynxe`, originally JManus) is a Java plan runner for exploratory jobs that still need a fixed shape. Func-Agent mode is a saved plan template: steps, agent configuration, and tool selection, published before any execute call. Execution is a separate request (`executeByToolNameAsync` or the sync twin) and returns a `planId`. `PlanExecutionRecord` holds `currentPlanId`, `rootPlanId`, `parentPlanId`, status, `agentExecutionSequence`, `userInputWaitState`, `summary`, `structureResult`, and nested `subPlanExecutionRecords`. A child step is an `AgentExecutionRecord` with its own `stepId`, result, and `thinkActRecords`. Human input is `POST /api/executor/submit-input/{planId}` while the same plan waits. A delete of an execution record is a mark, not a database delete. `design.md` only lays out Pinia owners for the UI.

**skill-up** (`alibaba/skill-up`) evaluates a Skill, an agent, or a workspace against declarative cases, then grades the reply and the file changes. skill-upper drives an evaluate, diagnose, fix, rerun conversation. The DeepSeek Harness plugin adds isolated report directories, an opt-in observer, and an approval gate before a regression case may be written. That loop reads evidence and rewrites a Skill. It is not a step on the live packet.

**open-code-review** (`alibaba/open-code-review`) reviews a diff or a file scan. Engineering picks the files, the rules, and the groups. The model comments inside a group. A five-gate filter runs before any sub-agent. Each group is its own sub-agent, its own message buffer, and, when the diff is large, a plan call that is not allowed to use tools. Comments are positioned and filtered outside the model, then written to JSON or to an append-only JSONL session log. The session log is not the git tree. A failing group warns and stops; it does not retry itself, and it does not rewrite the corpus.

**ROLL** (`alibaba/ROLL`) is a reinforcement-learning library for training language models: RLVR, agentic rollouts, distill, DPO, PPO, GRPO, and the rest. Ray schedules training roles. vLLM and SGLang are the rollout engines. It writes checkpoints of weights. It is not a seat, and it does not fetch a section.

**ROCK** (`alibaba/ROCK`) is the environment kit under that training stack. Admin schedules, a worker runs, a rocklet proxies. A sandbox is a Docker-backed session with a memory cap and a CPU cap. The GEM shape is `make`, `reset`, `step`, `close`: observation in, action in, observation and a terminal flag out. The sandbox does not hold the training corpus. Bash and chat are other action protocols on the same isolation boundary.

**zvec** (`alibaba/zvec`) is an in-process vector library: dense and sparse vectors, full-text, hybrid filters, a write-ahead log, many readers, one writer. v0.7 is also a file store inside ReMe and a local `zvec-grep` CLI. It ranks by similarity. It does not name the next section, and it does not choose a model.

**rtp-llm** (`alibaba/rtp-llm`) is Alibaba's inference engine: paged and flash attention, weight-only quantisation, LoRA, multimodal inputs, prefix cache, system-prompt cache, speculative decoding. A prefix cache is a serving optimisation. It is not a section, not a kb key, and not the bus.

## Best practices

These four answers stay inside the seven fields. They are how this pack would fill the open holes. They are not decided lines.

### What must be non-empty before a store write or the first tool

Open Code Review never hands the model a blank file set. The diff provider materialises the change, then `whyExcluded` drops binaries, user excludes, unsupported extensions, and default paths, and `selectFiles` also drops deletes and diffs that are too large. Only what survives is grouped. The plan phase, when it runs, is one `PLAN_TASK` call with no `Tools` field. The checklist comes back as `plan_guidance` before the main loop's first tool. Lynxe refuses the other order: the developer quick start says the Func-Agent template, including plan steps and tool selection, must exist and be published before `executeByToolNameAsync`. The execute call does not invent the steps.

Map that onto the section record. Before the mesh stores the section, and before that section's first tool, **Section, Function, Read, and Do are non-empty**. Section is the id (`planId`, or the group key). Function is the one-sentence job (the published tool description, or the rule the filter already bound). Read is the only files, already chosen by the filter, not by the model. Do is the ordered steps (the saved plan, or the checklist). **Write** is the one kb key and must be non-empty before the store that uses it; it is not filled by the tool. **Next** may already name a successor, because the section names the next id, but it is not durable until a lead lets it pass. **Rail** may stay empty. Empty Rail is the decided shape of that field, not a missing packet.

Do not copy Open Code Review's skip of the plan on a small diff. This harness already requires the inline plan before any write. A short Do is still Do. A packet with an empty Read or an empty Do does not reach the model and does not get a store write.

### When a lead cuts NEXT

Spring AI Alibaba's human-in-the-loop hook runs after the model responds and before any tool call. `approve` runs the call unchanged. `edit` changes the call, then runs it. `reject` does not run the call; the hook synthesises an explanation back into the conversation. The graph state, including `InterruptionMetadata` (tool name, arguments, description, decision), is saved by the checkpointer, not by the tool and not by the model. Resume uses the same thread id and an empty input, because the checkpoint already holds the state. skill-up draws the same line one step later: a candidate observation is not a case file; listing it does not authorise a write; an approved observation id is required; the writer never overwrites an existing file, and a failed validate rolls the new file back. Open Code Review's review filter drops a comment after the model emitted it, and the git tree is untouched. A sub-agent that produces nothing leaves a warning on that group. The warning is not a retry inside the agent, and it is not a clean empty success.

The lead still does not write the corpus. The mesh stores the cut, because the mesh is the store. What it stores is a side row, the same family as the checkpoint and the JSONL event: timestamp, task id, state `cut` (or the existing commit-refuse), the refused next id, and the reason. It does not store a new section whose Next is the cut target. The current section stays the current section. The cut is not a clean stop. A clean stop remains only NEXT stop, or a section that already holds both rail sentences. The next fetch of the section returns that current section, not the refused successor.

### What a parent hands a child

Open Code Review's parent (`Agent.dispatchSubtasks`) does not hand a child the whole review. The child gets one group: at most ten files, that group's diffs, and the rules merged for those files. It has its own goroutine and its own message buffer. Files outside the group may be named as `change_files` and may be opened with `file_read_diff` or `code_search` for understanding. Findings there are off-limits as comment targets. The prompt says to ignore issues outside the diffs it was given. A failure stays in that group. Lynxe keeps `parentPlanId` and `rootPlanId` on the record, and the child is a nested `subPlanExecutionRecords` entry with its own `stepId` and result. The child does not own the parent's plan id. ROCK's sandbox is the same split at the process boundary: the admin schedules, the sandbox receives a session and one action, and it does not receive the admin's store.

The packet a parent hands a child is a new section body: that child's **Section, Function, Read, and Do** only. The child may not read or write the parent's **Write**, **Next**, or **Rail**. It may not append the parent's kb key, may not retarget the parent's successor, and may not flip or clear the parent's switch. Context outside the child's Read, if it is passed at all, is understanding-only and cannot become a write target. Spring AI's nested-agent example can set `includeContents` so the child sees the parent graph's message history, and it shares one checkpoint saver. Refuse that sharing for this packet. The child's tool result comes back to the parent as a result on the parent's packet. It does not pick the next seat, and it does not become the parent's Next.

### Steer line, human rail side record, and section Rail

These are three objects. The pack never stores them as one blob.

The **steer line** is the live edit. In the human-in-the-loop hook, `edit` changes the pending tool call and the same run continues; `approve` lets the in-flight call finish. Lynxe's `submit-input` fills `userInputWaitState` on the same `planId` and does not start a second plan. That matches the rail you already have: a line enters the live step, the in-flight action finishes, the run is not cancelled. The steer line is not a field of the section, and it is gone when the step finishes. The next section fetch does not include it.

The **human rail side record** is the checkpoint row: interruption metadata, the skill-up candidate observation, the Open Code Review JSONL event and the warning array. The checkpointer writes it. The observer writes it. The lead does not, and the seat's packet cannot. skill-up says review is a separate boundary from the case write: collection creates a candidate; approval of an exact observation id is a later call; only then may `write_observation_case` run. The side record may be fetched as ledger (`task.get`, the append-only progress file), read-only. It is not a section field. The next fetch of the section packet does not include it, or a seat could write it.

The **section Rail** is the switch on the section record. It stays empty, or it holds a switch only the human flips. It is not the steer line and not the side record. It is the only one of the three that may ride in the next section fetch. A seat still cannot write it. Wider read root, push, and delete stay shut until that human switch opens one.

## Adopt

Suggestion, this note only. A seat may adopt these four sentences. They do not enter the plan until a second finished note repeats them and no lead cuts them.

1. Before a store write or the first tool, Section, Function, Read, and Do are already non-empty; Write names the one kb key before that store; Next may be named but is not durable until a lead lets it pass; Rail may stay empty.
2. When a lead cuts NEXT, the mesh stores the refused next id and the reason as a side row, the lead does not write the corpus, and the cut leaves the current section plus a commit-refuse row rather than a new section or a clean stop.
3. A parent hands a child only that child's Section, Function, Read, and Do, and the child may not touch the parent's Write, Next, or Rail.
4. The steer line, the human rail side record, and the section Rail field are three objects; the next section fetch may include only the Rail switch, the side record stays on the ledger, and the steer line is not part of that fetch.

Two narrower agreements, already decided, are not new claims. Open Code Review's append-only JSONL session and Lynxe's mark-instead-of-delete match the append-only progress record and the ban on delete. AssistantAgent's `search_exp` then `read_exp`, and skill-up's refusal to treat a whole Skill body as the eval config, match the catalog rule: one body loads only when the description matches. Do not reopen those lines.

## Refuse

Refuse `RoutingAgent`, graph conditional edges, and any A2A directory as the way the next seat is chosen. The successor is the section's Next, after a lead has let it pass. The mesh does not choose a model, and a tool result does not pick the next seat.

Refuse handing a child the parent's message history, the parent's checkpoint, or the parent's Write, Next, and Rail. Spring AI's nested example shares one saver and can pass parent contents. That sharing does not fit this packet.

Refuse FastIntent and any experience row that answers without the packet. AssistantAgent will skip the model call for a familiar scenario. This loop still receives the packet and the tool list, makes one call, and reads the result. A stored experience is not a section.

Refuse code-as-action as the tool step. AssistantAgent generates code and executes it in a Graal sandbox. This harness's CLI is argv only, with no shell and no delete. Generated code is not an argv list.

Refuse skill-up's automatic Skill install. If `skills` is omitted and a `SKILL.md` sits above the config, skill-up installs it. Here a body loads only when the description matches. Also refuse letting skill-upper edit a Skill, or write a regression case, on the live packet. The plugin itself says it does not edit Skills and that an unapproved observation cannot become a case. Keep that refusal. Do not weaken it into "the seat may evolve the skill if the judge is cheap."

Refuse Open Code Review as the seat's search. `ocr scan` and `code_search` let the model walk a repository. A seat does not search the repo to find out what to do. Read is the only files. Delegation mode, where the host model reviews with OCR only selecting files, still lets that host model run. It does not replace the handed packet.

Refuse zvec, ReMe, or `zvec-grep` as the store or the fetch. The mesh kb is the store. Fetch is `kb.query`, `room.tail`, and `task.get`. A similarity rank would choose what comes back, which is a second fetch, and it must not choose the next section or the model. Embeddings stay on the local embedder. zvec's single writer and write-ahead log do not make it the section store.

Refuse rtp-llm's prefix cache and system-prompt cache as memory. They are serving caches. They are not a section, not a kb key, and not a substitute for the mesh.

Refuse ROLL's weight checkpoint, ROCK's admin database, and Lynxe's plan database as a second corpus. Leads do not write the corpus. A training checkpoint is not a section.

Refuse a retry inside the seat when the model returns nothing. Open Code Review already isolates a failed group and leaves the retry to the outer pipeline. This harness already has a different successor for an empty model call: one terminal failed task, with a reason and no evidence. Do not import OCR's three empty rounds, ROLL's rollout retry, or Spring AI's tool-retry hook as that successor.

## Park

Park as a later loop, after a packet or its traces exist, and not as a step on the live packet.

skill-up and skill-upper. The eval, the isolated report directory, the before/after `result.json`, and the approval-gated regression case are a useful later loop. They read evidence. They do not run on the seat that is still producing the section. The observer stays off until that later loop, and even then a candidate is not a section write.

ROLL, and ROCK as its training ground. Agentic rollouts, reward, distill, and GEM `step` are how a later loop would train. They do not train and they do not install a trainer on this harness. ROCK's admin, worker, and rocklet are a fleet scheduler. This mesh is already the bus. Do not put a second scheduler under it.

ROCK's sandbox image, memory cap, and CPU cap may be re-read when an experiment seat needs a guest. They stay parked now. The experiment seats already own isolation, and a guest still gets no socket. Parking ROCK does not open a Bash session on the packet.

rtp-llm. Park it beside AirLLM as an optional later inference engine. It is not the loop, not the bus, and not a model chooser. Prefix cache stays inside that engine if it is ever used. It never becomes the fetch.

zvec. Park it as an optional later similarity index beside the local embedder, still not the section fetch, still not allowed to name Next or to pick a model. Hybrid search and `zvec-grep` stay in that later slot. They are not how a seat finds its job.

AssistantAgent's learning extractor, experience CRUD, and Skill conversion. Park them with the other harness-edit loops. Progressive disclosure is already how skills load. The converter that rewrites a Skill from a trace is the later loop, not this packet.

Spring AI Admin, Studio, and the graph's PlantUML export. Park them as a spectator view, the same kind of view as the live page: they may show state, and they are not the section store and not a live-step steer. A visual router is still a router. Leave it parked.

Lynxe's right-hand step panel and `userInputWaitState` are a useful picture of the three objects above. Do not port the plan-template service. The section record is already the plan, and it cannot grow an eighth field to hold wait state. Wait state, if it is kept, is the side record.
