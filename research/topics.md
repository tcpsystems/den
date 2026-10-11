# GitHub topics: harness, skills, agent-harness

Niche note for the harness plan. Not a second plan. Nothing here is decided. A line can move into the plan only when a second finished note states the same scaffolding claim and no lead cuts it. This file is one source.

Opened on 2026-10-11:

- https://github.com/topics/harness
- https://github.com/topics/skills
- https://github.com/topics/agent-harness
- https://github.com/ai-boost/awesome-harness-engineering
- https://github.com/ai-boost/awesome-harness-engineering/blob/main/templates/PLAN.md
- https://github.com/ai-boost/awesome-harness-engineering/blob/main/templates/HARNESS_CHECKLIST.md
- https://raw.githubusercontent.com/ai-boost/awesome-harness-engineering/main/templates/PLAN.md
- https://github.com/bradAGI/awesome-cli-coding-agents
- https://github.com/the911fund/skill-of-skills

The mesh stays the bus, the store, and the fetch. It does not choose a model. The section record still cannot grow past Section, Function, Read, Do, Write, Next, and Rail. Counts below are what those pages showed when opened, not a quota essay.

## What it is

The `agent-harness` topic page defines the object this plan is building. Agent harnesses are the runtime scaffolding around agents. They usually combine context delivery, tool interfaces, planning state, memory, sandboxes, permissions, evaluation, and observability so a longer task stays repeatable, safe, and traceable. The page listed 1,535 public repositories.

The `harness` topic is the wider label, 2,393 repositories, with no definition paragraph of its own. Its first page is a catalog: a long-horizon superagent, swarm runtimes, code-review pipelines, typed agent frameworks, and memory products. That catalog is not a loop. The highest-starred row describes sandboxes, memories, tools, skills, subagents, and a message gateway. This plan already has a local harness and a mesh. This note does not replace the mesh with that product.

The `skills` topic is larger still, 16,290 repositories, and mostly a catalog of skill packs, plugin directories, and graphs. A skill on that page is a discoverable body with a platform marker (`SKILL.md`, a rules file, or `AGENTS.md`). It is not a section record and it is not a model id. The deep note for the top skills framework belongs in `superpowers.md`, not here.

`awesome-harness-engineering` says harness engineering is the scaffolding around the model: context delivery, tool interfaces, planning artifacts, verification loops, memory systems, and sandboxes. The list is about the harness, not the model. Its contents keep planning, permissions, memory, and human-in-the-loop as separate headings. Its own template set is four files: `AGENTS.md`, `PLAN.md`, `IMPLEMENT.md`, and `HARNESS_CHECKLIST.md`.

`awesome-cli-coding-agents` (page date 2026-10-05) lists 130-plus terminal agents plus the harnesses that orchestrate, sandbox, or extend them. The useful split on that page is agents versus session managers, orchestrators, and infrastructure. Plan mode, a permission engine, an approval inbox, and a review gate before landing show up as separate features, often on different tools.

`skill-of-skills` is a quality-ranked directory, not a harness. The README said 1,400-plus skills and 420-plus at a curated tier, counts as of 2026-10-11 18:25 UTC. Ranking is 45 percent structural quality, then star velocity, star gravity, recency, social activity, and user ratings, times a reputation factor, with a small penalty when dependency vulnerabilities are at least five. The page says the scores are evidence, not endorsements. Curated means gotchas, examples, and verification. Medium risk means extended permissions, and the page names shell access and subagents as that extension. Tools under quality 40 are retired after 14 days below the threshold. That directory must not become the fetch.

## Best practices

These are the scaffolding readings that touch the four open holes. Each one is a suggestion from this note.

### Which fields must be non-empty

The checklist is the gate, and the plan template is the shape. Both were opened. The checklist says a failing item is a blocker. Before the task starts it wants `PLAN.md` for a non-trivial task, scope boundaries written down (in scope and out of scope), milestones with explicit verification commands, and eval criteria written before the task starts, not after. It also says tool schemas stay minimal: no optional field the agent will not use. Long-lived state (plans, decisions, progress) lives in files, not in the prompt. Context is scoped to this task, not the whole tree.

The plan template is created at the start of a task. Its task line is one sentence: what is being built or fixed. Milestones are an ordered list, and a milestone is not marked complete until its verification gate passes. Notes are a running log: append entries, do not overwrite. Context, approach, open questions, and risks are extra headings on that template. They are not extra fields on this section record.

Folded into the seven fields, without growing the record:

| Template demand | Section field | Empty before the first tool or the store write? |
| --- | --- | --- |
| The artifact exists and has an identity | Section | No. Must be non-empty |
| One sentence task | Function | No. Must be non-empty |
| In-scope list. Out of scope is whatever Read does not name | Read | No. Must be non-empty |
| Ordered milestones and the checks that close them | Do | No. Must be non-empty |
| The one output key, filled when there is a product | Write | Yes. May be empty |
| Successor id, or stop | Next | No. Must be an id or stop, not blank |
| Human switch | Rail | Yes. May be empty |

Next is the tense mapping. The template never uses the word Next. This seat maps "eval criteria and a verification end exist before the task starts" onto Next so the decided line stays true: the section already names the next id. A blank Next would make the first tool a search for what to do, which this plan already refuses. Write stays empty until there is a key to store. Rail stays empty until a human flips it. Neither is a start condition.

A store write of the section is the same gate as the first tool. The checklist puts the plan in a file before execution. Storing a section that is missing Section, Function, Read, Do, or a non-blank Next would store a prompt, not a plan. Write and Rail are the two fields the checklist does not demand before the start.

`awesome-cli-coding-agents` is the second opened page on the same point, weaker but the same shape. Several rows separate plan mode from the tool loop (a plan mode before unattended runs, a plan mode beside a permission engine). That is plan-before-tool, not a second schema. One note still cannot decide it.

### When a lead cuts Next

The human-in-the-loop section of the awesome list describes a pause that persists, not a reviewer who rewrites the work. LangGraph pauses the agent mid-loop, persists state, and resumes only after human review. Dify suspends at a critical decision and routes the continuation from the human action: approve, reject, or escalate. A separate row on the same list says a mechanical gate, not a prompt request, checks the turn before that turn can complete.

The CLI list says the same thing from the landing side. One control plane hands work across a planner, an executor, and a human reviewer, and changes pass a diff review gate before they land. Another keeps supervisor review against acceptance criteria and stores the run evidence on disk. The reviewer does not land the change.

Read against this plan, a lead cut of Next is that reject. The lead does not write the corpus, so the lead does not store the next section and does not rewrite the current one. The mesh, which is already the store, appends the cut onto the progress record that already exists: timestamp, task id, state, artifact path, next id. State names the cut. The next id on that row is the current section's id, not a new section. No new section body is stored. The current section stays as it was, with its Next unadvanced. The run is not a clean stop, and it is not the terminal failed task already reserved for a model call that produced nothing. Deny-and-continue on the same permissions section is the matching recovery: the cut does not halt the loop.

What is stored: one progress row for the cut. Who stores it: the mesh. What the cut leaves: the current section, unchanged, plus that row. The lead writes neither.

### What a parent hands a child

The context section of the awesome list is the clearest parent-child rule on these pages. An isolated subagent starts fresh. The page calls that isolation a firewall. A forked subagent inherits the supervisor's full conversation, with the trailing tool call excised and the task description rewritten as a user message. The same paragraph's rule is to fork a worker that continues an in-progress investigation, and to isolate a verifier that must judge independently. The planning section adds that every handoff needs a typed schema, a constrained action schema, and an explicit boundary check. A survey blurb on the same page says more than four agent-to-agent handoffs almost always fail.

The CLI list names the same split without the theory. One terminal agent has three subagent modes, Fork, Cold, and Explore, beside a permission engine and a plan mode. Cold and Explore are the fresh starts. Fork is the inheritance mode.

This plan's child is another department with its own task. It is the isolated case, not the continuing worker. The parent therefore hands a new packet whose only filled fields are that child's Section, Function, Read, and Do, plus that child's own Next as an id or stop. The parent's Write, Next, and Rail are not in the packet. The child cannot touch them, because they are not present and the boundary check refuses a write to them. The child's Write and Rail start empty, under the same start gate as any other section. A forked copy of the parent conversation would hand the child the parent's output key, the parent's successor, and the parent's human switch. That is the handoff this seat refuses.

### Steer line, side record, and Rail

They are three objects. The awesome list already separates them into three primitives: planning artifacts, permissions and authorization, and human-in-the-loop. The `agent-harness` topic sentence separates planning state from permissions the same way. The CLI list separates them in the product rows: an approval inbox is not the session being steered, and a mobile approval that carries the command text is not the plan file.

Mapped onto this plan:

| Object | What it is | Who writes it | May the next fetch include it? |
| --- | --- | --- | --- |
| Steer line | A hard signal for the live step | The human, during the step | Yes, the one live line. If it is not in the loaded context, it did not happen |
| Human rail side record | The approval and reject log | The mesh, as a side record a seat packet cannot write | No. It stays in the store |
| Section Rail | Empty, or a switch only the human flips | The human, on the section | Yes. It is a field of the section being fetched |

The skills section of the awesome list records the steering study: passive documentation is ignored, and a hard steering signal placed in the loaded context (a skill file, an error, a prompt) is what changes behavior. The line on that page is that if the guidance was not in the loaded context, it did not happen. That is why the next fetch includes the live steer line. It does not include the whole side record. The checklist says context is scoped to this task, and long-lived decisions live in files, not in the prompt. The side record is that file. Fetch returns the section, which includes Rail when a human has set it, and the one steer line that applies now. A seat still cannot write Rail, the steer line, or the side record.

Collapsing the three into one object would either drop the steer (so it did not happen) or copy the side log into the packet (so the fetch is no longer scoped). This seat refuses that collapse.

## Adopt

Adopt, as suggestions of this note only:

- Before a store write or the first tool, Section, Function, Read, and Do are non-empty, and Next is an id or stop. Write and Rail may be empty. No eighth field is added to hold context, approach, risks, or open questions.
- Out of scope is the complement of Read. The checklist wants it written down. This record writes it by naming the only files, not by adding a field.
- A lead cut of Next stores one progress row on the mesh and stores no next section. The lead writes nothing. The current section stays unadvanced. The cut is not a clean stop and not a corpus edit.
- A parent hands a child a fresh packet. The child does not receive, and may not write, the parent's Write, Next, or Rail.
- The steer line, the human rail side record, and the section Rail field stay three objects. The next fetch includes the live steer line and the section Rail. It does not include the side record.
- Notes and the progress record stay append-only. The plan template says append and do not overwrite. That matches the ledger already in the plan.
- Destructive confirmation stays a human gate. The checklist says destructive operations require explicit confirmation. This plan already keeps delete shut until a human opens it.
- A skill remains a catalog body loaded when its description matches. This pack adds only that a curated skill shows gotchas, examples, and verification. That is a quality bar on the body, not a new section field.

## Refuse

- Do not let the mesh, a topic catalog, or a skill directory choose a model. Rows that pin a provider per role, or that route by a learned tier, are model selection. The mesh is bus, store, and fetch.
- Do not treat a directory rank as an endorsement or as a packet field. The skill directory says its scores are evidence, not endorsements.
- Do not let a seat or a lead rewrite the stored section or the corpus. The template's "update the plan throughout execution" is refused where it would overwrite. Append a progress row, or store a later section. Do not edit the current one in place.
- Do not let a reviewer agent write the next section. The same awesome page records that reviewer agents can degrade performance, and this plan already says leads do not write the corpus.
- Do not fork the parent's conversation into the child. Isolation is the firewall. Fork is inheritance of Write, Next, and Rail.
- Do not chain more than one parent-to-child handoff as a habit. The page's survey line is that more than four agent-to-agent handoffs almost always fail. One child packet is the pattern. A pipeline of handoffs is not.
- Do not treat a natural-language prohibition as the rail. The permissions section says almost none of the sampled prose security rules are backed by a real control, and that "do not" is not a deny. Rail is a switch the human flips, or it is empty. The side record is the store's log. The steer line is loaded context. None of the three is a sentence the model is trusted to obey.
- Do not add optional fields the seat will not use. The checklist forbids them, and the section record is already closed at seven.
- Do not make the skill directory, a knowledge graph, or a second memory product the store or the fetch.
- Do not mount a self-editing harness, a skill optimizer, or a trace-trained loop on this packet. Those rows exist on the awesome list. They are the later loop already refused on the live packet.

## Park

Park as a later loop, after a packet and its traces exist:

- The skill directory's nightly rescore, 14-day retirement under quality 40, and curated bar (gotchas, examples, verification). Useful once skills are replayed. Not a field on the section, and not a reason to load every body.
- Forked subagents that continue one investigation. The page recommends fork for that case and isolate for a verifier. This packet's child is the isolated case. Fork stays parked until a later loop has a same-department continuation, which this plan does not.
- The four human-approval shapes on the awesome list (a blanket hook, a per-tool context, an async third-party approval, a protocol elicitation). The live shape is already a steer line plus a side record plus Rail. Picking a transport is later.
- Eval harnesses, trace debuggers, and compaction studies listed on the same pages. They diagnose after commit. They are not a step on the packet.
- Extra plan headings (context, approach, open questions, risks). If a later section needs them, they are prose inside Do or a later section. They are not new fields.
- Template files as repo artifacts (`PLAN.md`, `IMPLEMENT.md`, `AGENTS.md`). This plan's packet is the section on the mesh. Copying those filenames into a repo is a later packaging choice, not the loop.

## Holes this note does not close

No page in this pack names Section, Function, Read, Do, Write, Next, and Rail. The mapping above is this seat's reading of the checklist, the plan template, the isolation paragraph, and the human-in-the-loop section. A second note has to state the same mapping before any of it is decided. A lead can still cut it.
