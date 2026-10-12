# Harness plan

The mesh is the memory. It fetches a section, stores the current one, and holds the rules between them. A model only answers the packet it is given.

## 1. Architecture

```
you
  rails: steer a live step, approve a commit, no delete
    five leads
      section path, stored on the mesh
        one task, one seat, that seat's own department
          one direct model
            SECTION, FINDING, NEXT, PROPOSAL
              a lead may cut NEXT
                the next section is stored
```

A seat does not search the repo to find out what to do. The section names the files, the function, and the next id.

## 2. Five leads

They judge what enters this file. They do not delete. They do not write the corpus.

| Lead | Model | Owns |
| --- | --- | --- |
| critic | `claude/claude-opus-5` | Architecture. Veto on the shape of the loop |
| ada | `grok-cli/grok-4.7` | Memory rules. What a section must contain |
| quinn | `gemini/gemini-3.1-pro-preview` | Model policy. Which direct ids stay |
| steer-rails | `codex/gpt-6-astra` | Rails and steer. What only you can flip |
| remember-facts | `claude/claude-sonnet-5` | The section path and what "current" means |

`remember-facts` dissented: it would drop the mesh and call seats from the lead. The other finished leads kept the mesh as the bus. The bus stays. It is not a model router.

## 3. Memory rules

A section record has these fields and no others.

| Field | Rule |
| --- | --- |
| Section | Stable id. One job |
| Function | One sentence |
| Read | The only files. Read-only |
| Do | Steps, in order |
| Write | One mesh kb key |
| Next | The next section id, or stop |
| Rail | Empty, or a switch only you flip |

Store and fetch go through the mesh kb. Embeddings for OpenViking stay on local port 8091. The combo `deerflow-memory` is not a memory model. It recorded 11 calls, 0 successes, about 547 seconds each.

Current knowledge can be replaced by a later section. Foxcore is the current state, not a text that cannot change. A seat still may not delete a file.

## 4. Model rules

One seat, one direct model id, through OmniRoute. No `auto/*`. No `deerflow-*` combo.

The fifty use the fast, cheap, high-success set from the provider board: Mistral, DeepSeek Flash, Gemini Flash, Gemma, Haiku, ModelScope. A dead id fails that seat and names the id.

The five leads use the five models in the table above. That is the only place the stronger models run.

What the board showed, and the seats confirmed: a combo walks hundreds of targets and the gateway pegs. Qwen's token plan is out of monthly quota until 14 Oct and averages 33 seconds. Ollama Cloud succeeds about 60 percent of the time. Mistral is about 97 percent at a fifth of a second. DeepSeek Flash is about 93 percent at a third of a second.

## 5. The other seats

They research, debate, and analyse. Each has its own department, so it cannot take another seat's task. Each returns four lines: `SECTION`, `FINDING`, `NEXT`, `PROPOSAL`.

| Seats | Work |
| --- | --- |
| connect-channels, connect-identity | How one lead sits on many surfaces |
| remember-skills, remember-sessions | How a finished function becomes the next section |
| schedule-cron, schedule-park | How a parked thought fires later |
| delegate-children, delegate-rpc, delegate-council | When a child runs, and what it cannot touch |
| search-web, search-browser, search-media, search-models | Which direct model and which search a section may use |
| experiment-local, experiment-docker, experiment-ssh, experiment-remote | Isolation for that section. No socket on a guest |
| noor, sable, ivo, rhee | Turn three findings into the next section |
| bram, cass, dell, echo, jules, kit, lane, moss, ren, sol, tess, uma, ash, bay, cole, drew, indra, joss, keir, lux, maren, nico, opal, pax | One assigned section each. Read only |

Read-only roots are the checkouts named in `plan/nest.md`. Xenode is not on this machine.

## 6. Rails

You steer a live step by adding a line. The run is not cancelled. You approve a commit that widens a read root, pushes, or deletes. No seat has a delete grant.

## 7. Cut

Routers and combos for this council. Any delete or overwrite of the corpus. `deerflow-memory` as the memory model. Shared departments that let one seat run another's task.

## 8. Running

Hub `127.0.0.1:8787`. Fifty seats connected. This file is the packet they are judging and filling. A lead may cut a `NEXT` before it is stored.

## Who wins

The spectator may write, steer, and add seats. Consensus wins. A sentence enters this file as decided only when more than one seat states the same scaffolding claim and none of the leads cuts it. One reply, including mine, stays a draft.

## Consensus so far

These claims showed up in more than one finished reply. They are scaffolding, not model names.

- A packet is handed to the seat. The seat does not go looking for the job.
- An inline plan names the files and the order before any write.
- The loop is: receive the packet and the tool list, make one call, read the result, repeat until a final message.
- The progress record is append-only: timestamp, task id, state, artifact path, next id.
- Memory is layered: tier, scope, key, value, source, timestamp. A later section supersedes. Nothing is deleted.
- The human surface has slots for input, history, tool activity, and review.

## Not consensus

`sable` read the Goose recipe as a cooking recipe. That reply does not count. Claims that are only quotas, auth failures, or missing model ids do not count. `remember-facts` still dissents on the mesh as the bus. The other leads kept the bus. That dissent stays open until another lead joins it.

## Seat expansion

These lines are from finished seats. They are not a build.

**mesh-docs.** The mesh is already the bus, the store, and the fetch. It does not choose a model. `room.say` is not the store. Bulk text goes to `kb.put`. Fetch is `kb.query`, `room.tail`, `task.get`.

**memory-docs.** One store appends a section, fetches the latest, and refuses a second store. Embeddings stay on the local embedder. A language pass that leaves this store for a combo is refused.

**skill-docs.** Skills stay discoverable by catalog frontmatter. One body loads only when the description matches. Harness skills live in `skills/`.

**tool-docs.** Three tool steps, each with name, inputs, outputs, params, and next. The result attaches to the packet. It does not pick the next seat. CLI is argv only, no shell, no delete. MCP discovers, then calls.

**order-docs.** Build order they named: packet, then store, then fetch, then tool step, then skill load, then steer, then rail, then commit gate. Fetch before store leaves the first index with nowhere to live.

**fox-docs.** A Foxcore seat stays inside one repo prompt and the contract that repo already speaks. One of fox, foxd, or floki. Research stops at the brief. Development is a feature branch and a pull request. It does not merge, and it does not open a second product in the same seat.

**critic, on the ledger.** The progress record is one append-only file: timestamp, task id, state, artifact path, next task id.

**ada, on the edit.** The inline plan names the files and the order before any write.

**lux, on the loop.** The seat receives context and a tool list, emits one call, reads the result, and repeats until a final message.

**known-docs.** The pieces that repeat across the named harnesses are plan, workspace, change, close, record, and split. A recipe or a skill that is never replayed is only a note.

**said-docs.** Two LINE sentences become one section by pairing them on the rail, in the order they were said. Neither sentence is deleted. On a disagreement the shared field is kept and both sentences stay.

**tcp-docs.** The tree on this machine has four parts: `llm-intel`, `tcp-passkey`, `nginx`, and `systemd`. The brief is what fences a seat. A grant path does not, on these host-side workers.

**steer-docs.** A steer line enters the live step. The in-flight action finishes. Wider read root, push, and delete stay shut until a human opens one.

**advise-live.** Suggestion: the live page at :8790 is a spectator. Graph harvests seat state and completed LINE candidates; chat without a seat says to union; chat with a seat submits one new task to that department. Neither path is the section store or a live-step steer.

**review-rules, chase-rules.** Draw the packet gate between the model and the mesh, and keep the human rail as a side record that a seat packet cannot write.

**source-guard, source-hay.** A packet does not reach the model until a gate has run, and a tool result takes the same gate before store.

**chase-ledger, review-ledger.** task.submit appends the progress line as the self-contained body of one new task, and task.claim hands that body to the next seat as the packet so the seat never opens a file.

**chase-halt, review-gate.** A model call that produced nothing has one successor, a terminal failed task with a reason and no evidence; a halt and a lost lease are not that same successor.

**chase-refuse, said-docs.** A commit passes only for NEXT stop or a section that already holds both rail sentences, and any other NEXT skips the lasting write and leaves a commit-refuse row instead of a clean stop.

**source-opensearchvl, source-erl.** A miss is a named successor, not silence and not a clean stop.

**source-evolver, source-meta.** Scaffold a later loop after the first harness exists: traces diagnose, propose, evaluate; it is after commit, not a step on the packet.

**explore-live, advise-live.** The live page is a spectator: graph reads seat state and LINE candidates; chat says to union or submits one new task to a clicked seat, and neither path is the section store.

**source-nano, source-pi.** Every tool call is gated before it runs.

**explore-fox, fox-docs.** Start on FOXCORE-CLOUD.md and one repo prompt; land features on a new branch from develop or main, then a pull request into that integration branch.

## Suggestions

**Suggestion:** source-rllm and source-meta already state the same later-loop rule. A training or harness-edit loop stays off this harness until a packet or its traces exist, it only reads that record, and it does not train or install a trainer. Meta-Harness also says that loop is after commit, not a step on the live packet. Treat ERL's second attempt and OpenSearch-VL's training recipe the same way. Do not mount their retry, their decorator, or their visual tool catalog on this packet.

## Research workflow

Owner instruction. The next research run expands this file. It does not open a second plan. The slash-command deep-research whose query was skill names is stopped. Do not resume it.

- Read this file whole before any write.
- Append under **Seat expansion** or **Suggestions**.
- Do not delete. Do not rewrite a decided line above this heading.
- A sentence is decided only when two or more finished seats state the same scaffolding claim and no lead cuts it. One research reply stays a **Suggestion**.
- Do not build the harness.

Source pack the run must read before it writes:

- GitHub topics: `harness`, `skills`, `agent-harness`.
- Named repos: `obra/superpowers`, `farion1231/cc-switch`, `Graphify-Labs/graphify`, `addyosmani/agent-skills`.
- Alibaba: org `alibaba` (`open-code-review`, `skill-up`, `ROLL`, `ROCK`, `zvec`, `rtp-llm`); `QwenLM/Qwen-Agent`; `agentscope-ai/agentscope`; `Alibaba-NLP/DeepResearch`; `alibaba/spring-ai-alibaba`; `spring-ai-alibaba/AssistantAgent`; `spring-ai-alibaba/Lynxe`.
- Vendor skills, guides, guards, and harnesses: Vercel, Anthropic, Google, OpenAI, DeepSeek, Qwen, Kimi, Moonshot, Reasonix, Cursor, Grok, xAI, Cloudflare.
- Already on disk: `/tmp/harness-ten/*.md`. AirLLM is an optional later disk-stream inference slot, not the loop.

The notes from that run are in `research/`. The runner that produced them stays outside this repo.

Phases: inventory this file, research the source pack, verify each proposed line against this file, append only **Suggestion** lines. Output is a scratch expansion draft plus a union summary. No software.

**Suggestion:** Deeper notes for this expansion live in the session `research/` directory, one file per niche. They feed this file. They are not a second plan. A line moves here only after two finished notes state it and no lead cuts it. The owner asked for the research run: data, best practices, and every niche, still under these rules.

## Niche notes

The run `harness-plan-expand` finished. One line survived its verifier. Twelve candidates did not. The depth is in the session `research/` notes. These lines are suggestions until a second note states the same claim and no lead cuts it.

**superpowers, verified.** The steer line, the ruling ledger, and the section Rail are three objects, and the next fetch includes none of their bodies. It may carry only a pointer to a ledger entry that touches this task. Source: `obra/superpowers` `skills/subagent-driven-development/SKILL.md`.

**topics, superpowers, cc-switch, graphify, alibaba-agents, alibaba-platform, vendors-west.** A parent hands a child its own packet. The child does not write the parent's Write, Next, or Rail.

**topics, graphify, alibaba-agents, alibaba-platform, vendors-west.** A lead who cuts Next does not write the corpus. The mesh keeps the current section and a refusal row. The cut is not a clean stop and not a new section.

**Open on the next fetch.** Superpowers, the verified line, keeps all three bodies out of the next fetch. AgentScope, Spring AI Alibaba, and Graphify let the section Rail ride with the section, and keep the steer line and the confirm event out. The topics note would include the live steer line. That split stays open.

**Suggestion:** `research/topics.md`. GitHub topics `harness`, `skills`, and `agent-harness`. A skill body can carry gotchas, examples, and verification. That is a quality bar, not an eighth field. A prose "do not" is not the rail. Skill optimizers and trace-trained loops stay parked.

**Suggestion:** `research/superpowers.md`. Before a plan store, Section, Function, Read, Do, and Write are non-empty and are not placeholders. The first implementation tool waits until Rail is the human's approval. No seat delete, including a clean-review `rm` of the scratch plan. A spike that skips the plan is refused. Diagnosis of a finished transcript stays a later loop.

**Suggestion:** `research/cc-switch.md`. cc-switch is an operator switcher for provider keys, prompts, and skills. Park it outside the packet. Failover to the next provider is the bus choosing a model, so it stays off this loop. A same-name skill is kept as its own copy, not overwritten. Session-file delete stays shut.

**Suggestion:** `research/graphify.md`. Graphify is not the mesh, the memory, or a model chooser. A later local code parse may be a read inside a path the section already named. Semantic extraction, confidence fields, hyperedges, and a lesson file stay off the section record. An empty or smaller graph does not overwrite a stored one.

**Suggestion:** `research/agent-skills.md`. `addyosmani/agent-skills` matches the catalog rule: name and description in frontmatter, one body when the description matches, and the procedure does not name a model. Refuse injecting the whole catalog into the packet, persona-calls-persona, and `/build auto` as this loop.

**Suggestion:** `research/alibaba-agents.md`. AgentScope and Qwen-Agent: a tool runs only when Do holds the required arguments. A store that would set Next needs a non-empty Write. The child gets one self-contained message and returns its final text. Refuse `ModelRouterMiddleware`, Qwen `Memory`, and Mem0 as a second store. Tongyi DeepResearch's RL, synthetic data, and IterResearch Heavy stay a later loop.

**Suggestion:** `research/alibaba-platform.md`. Refuse Spring `RoutingAgent`, graph edges, and A2A as the way the next seat is chosen. Refuse generated code as the tool step. Refuse `zvec` as the store. Refuse skill-up auto-install. Park skill-up's eval, ROLL, and ROCK as later loops. Park `rtp-llm` beside AirLLM as an inference engine.

**Suggestion:** `research/vendors-west.md`. OpenAI's check finishes before the tool runs. Cloudflare throws before persist. Agents-as-tools stay the parent-child shape. Handoffs that transfer the conversation are refused. Vercel's schema, approval, and execute stay three objects. A client-supplied approval record is not Rail. A second model call does not repair a lead cut.

**Suggestion:** `research/vendors-east.md`. DeepSeek, Reasonix, Qwen, Kimi, Cursor, and Grok each have a path that picks a model. Those paths stay off the mesh. A child gets task text in and a final answer out, and an allowlist is not a bypass. Fail-open hooks are refused. Qwen Live and a second coding CLI are other harnesses, not this loop.

## Cross-note

`research/cross-note.md` compares the nine notes. No new sources. The nine researchers were stopped once the notes were on disk and they had stayed on their first turn.

**Open, the fetch.** Superpowers keeps the steer line, the side record, and the Rail body out of the next fetch, with a pointer only. Graphify, AgentScope, and Spring let the Rail switch ride with the stored section and keep the other two out. Topics alone would also fetch the live steer line.

**Open, the gate.** The notes do not agree which of the seven fields must be non-empty. They do agree not to add an eighth field. Six notes leave Rail empty at the gate. Superpowers is the dissent: the first implementation tool waits until Rail is the human's approval.

**Suggestion:** Leave both opens as opens. The child-packet rule and the lead-cut rule are the two lines that already repeat. Do not promote the field gate or the fetch contents until a second pass states one of them and a lead does not cut it.
