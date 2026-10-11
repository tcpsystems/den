# Superpowers

One source, one note. This file does not edit the harness plan and it is not a second plan. A line here becomes decided only when another finished note states the same scaffolding claim and no lead cuts it. Until then every line below is a Suggestion.

The mesh stays the bus, the store, and the fetch. It does not choose a model. The section record stays Section, Function, Read, Do, Write, Next, and Rail. Nothing in this note adds a field.

Pages opened for this note: the repo README, the skills tree, the hooks tree, `hooks/hooks.json`, `hooks/session-start`, `AGENTS.md`, and the skills `using-superpowers`, `brainstorming`, `writing-plans`, `subagent-driven-development` (including `implementer-prompt.md` and `task-reviewer-prompt.md`), `executing-plans`, `requesting-code-review`, `verification-before-completion`, `finishing-a-development-branch`, and `dispatching-parallel-agents`. Issue threads and the separate eval repo were not opened.

## What it is

Superpowers is a skill library plus a session-start guard for coding agents. The README describes it as a methodology: the agent does not start by writing code, it draws a spec out of the conversation, shows that spec in short chunks, and only after sign-off writes an implementation plan a junior could follow. On "go" it runs `subagent-driven-development` or `executing-plans`. Skills are supposed to trigger on their own. The README's basic order is brainstorming, then a git worktree, then writing-plans, then one of the two executors, then test-driven development, then code review, then finishing the branch. It says the agent checks for a relevant skill before any task, and that critical review issues block progress.

The guard in this repo is small. `hooks/hooks.json` registers one `SessionStart` command, matcher `startup|clear|compact`, and `hooks/session-start` injects the full text of `using-superpowers` into the session context. There is no hook that inspects a packet, blocks a tool, or writes a store. The Pi and Muse install notes say that same bootstrap is injected at startup and, for Pi, again after compaction. The Hermes note says a session that compacts with no post-compaction hook loses the bootstrap. The skill text itself tells a dispatched subagent to ignore `using-superpowers`. So the injection is the parent session's guard, not the child's packet.

`using-superpowers` says to invoke a relevant skill before any response or action, including clarifying questions, exploring, and checking files. Process skills come before implementation skills. User instructions (`AGENTS.md` and a direct request) take precedence over skills, and a skill workflow is skipped only when the human has explicitly said so.

## Best practices

### What must be filled before a store or the first tool

Map the plan's words onto the seven fields. Do not grow the record.

`writing-plans` is for a multi-step task and it runs before touching code. The stored plan has to start with a one-sentence Goal, which is Function. Each task names exact Create, Modify, and Test paths, which is Read. Each step is one action with a checkable result, in order, which is Do. The save path is one file under `docs/superpowers/plans/`, which is Write, the one key. The plan's own identity (the feature name and `Task N`) is Section. A line that decides nothing is not a filled field: the skill names `TBD`, "handle edge cases", "add appropriate validation", and "write tests for the above" as the opposite of a plan. A step is done when the implementer can write exactly one reasonable thing from it. The implementer writes the body; a plan that is longer than the code it describes has written the code instead. The README's shorter claim that every task contains complete code is the loose summary. The skill is the tighter rule, and it is the one that fits a packet that names files and order without becoming a transcript.

That store can be written before the human answers. The skill then says to wait for the review before implementation, even if an execution method was already chosen. `brainstorming` is the earlier gate on the same switch: nothing gets built until the human approves the full description, in chat for a small change and as a written design for a project. Building there includes product code, scaffolding, installing dependencies, creating a project, and invoking an implementation skill. Reading and exploring are allowed before that approval. The design, when there is one, is committed under `docs/superpowers/specs/` before it is handed over, and a builder check runs before the hand-over. Two exceptions drop the gate: the human opted out of the questions, or they asked for a throwaway spike. A spike skips the plan and the review, and keeping its output is a new request that comes back through the gate.

Next is not one of those required fields. Neither skill puts a next-section id on the packet before the save or before the first implementation tool. The handoff asks which executor to use. It does not store the following task.

Rail is the human switch, and it is not required to be non-empty for the plan store. It is required before the first implementation tool when this section is a build: the waited-for plan review, or the design approval. An empty Rail still means there is no switch. A quick, clear request in `brainstorming` is done without a design document. Between tasks, both executors refuse a fresh "should I continue?" So an empty Rail on a later task is normal. Consent to start on `main` or `master` is the other switch the executors name, and it stays empty unless that consent was given.

A completion store is a second write, not the plan store. `verification-before-completion` forbids a completion claim, a commit, or a move to the next task without a fresh command and its output in that same turn. `executing-plans` runs the task's test command and appends `Task N: complete` only if it passes. A failing run records nothing, and the task is not complete. The five fields above are still the packet. The missing piece is evidence that Do ran, not an eighth field.

### When a lead cuts NEXT

The controller cuts the move to the next task. The implementer does not. `subagent-driven-development` says never to move on while a review still has open Critical or Important issues that are neither fixed nor parked with a ruling at the five-round cap. Until then the controller appends a ledger line, `Task N: fix round R/5` or `Task N: parked — finding — Ruling: ...`, and it does not append `Task N: complete`. The complete line is what lets the next task start, and it is withheld. `requesting-code-review` says the same block in shorter form: fix Critical immediately, fix Important before proceeding, and never proceed with either still open. Minor findings are ledgered and do not enter the loop.

The lead does not write the corpus. The plan file stays the argument. If the plan itself is wrong, the controller rules, writes `Ruling: what — why — cost if wrong` into that plan's ledger, and carries the ruling in the next dispatch. It does not edit the plan text to hide the cut. "Deviating from the plan without a ledgered ruling is a decision made in secret" is the paired line in `executing-plans`. A silent discard is forbidden. The human sees the rulings because the controller copies every `Ruling:` line into the final message. A ruling that dies with the scratch workspace was a secret.

Who stores the cut: the controller, in that plan's ledger (`<repo>/.superpowers/sdd/<plan>/progress.md`), whose first line is the plan path. The implementer writes a report file the controller named and returns a short status. It does not append the ledger. Another plan's directory is never read or written.

What the cut leaves when the lead does not write the corpus: the plan is unchanged, Next is not stored, and the record is the ledger row that is not `Task N: complete`. On the stricter executor, a failed check leaves even less: `executing-plans` keeps the test output in the scratch workspace and appends no line at all. Either way there is no clean stop and no stored next id. The task is resumed, not followed by the next section.

### What a parent hands a child

`subagent-driven-development` says a subagent must not inherit the session's context or history. The controller constructs the packet. The dispatch is five parts: one line on where the task fits, the brief path as the requirements with the exact values, interfaces and decisions from earlier tasks that the brief cannot know, the controller's resolution of an ambiguity, and the report path plus the report contract. Exact values live only in the brief. A fresh subagent gets its task, the interfaces it touches, and the global constraints, and nothing else. It must not be told to read the whole plan. Prior-task summaries are not pasted in. A real session's dispatch that was almost all pasted history is the example the skill warns against.

Those pieces still fit the seven fields. The brief is this task's Function, Read, and Do. Global constraints are already implied by every task's requirements, so they ride inside Do, not as a new field. The interfaces from earlier tasks are the Produces and Consumes lines inside that same brief. The report path is scratch the parent named. It is not the parent's Write key, and it is not a second section key.

The child may not touch the parent's Write, Next, or Rail.

- Write. The parent owns the plan file and the ledger. The child writes the report to the given path and, for an implementer, commits the task's files. It does not choose a store key, does not append `Task N: complete`, and does not open another plan's workspace. The reviewer child is read-only: it does not mutate the worktree, the index, HEAD, or the branch. Its packet is three paths (brief, report, diff package) plus the constraints copied verbatim. The diff file is the view of the change. It does not crawl the repo to find the job.
- Next. The child does not read the rest of the plan, so it cannot see or set the next task. It never dispatches a helper or a reviewer. Review arrives from the controller after the report. A reviewer the worker spawns counts for nothing. "More tasks remain?" is the controller's question.
- Rail. The four stops stay on the controller: an irreversible or destructive operation, a security-sensitive action, a side effect outside the worktree that norms say to ask about first (merge, push to a shared branch, publish), and a plan so broken that every path forward is a guess. The child escalates to the controller with `BLOCKED` or `NEEDS_CONTEXT`. It does not present the finish menu and it does not flip the switch. `using-superpowers` tells that child to ignore the session skill, so the parent's bootstrap is not part of the fetch.

`dispatching-parallel-agents` repeats the same handoff shape for a different case: one agent per independent problem, constructed context, no inherited history, and no parallel work when there is shared state. The executor skill is stricter on one plan: never dispatch two implementers in parallel, because they conflict. One task, one child.

### Steer, side record, and Rail

They are three objects.

The steer line is a live human utterance during a step. Both executors refuse turning that into a between-task check-in. "Should I continue?" is named as waste. The current task keeps going. Rulings are decided and ledgered, not parked on a question. `brainstorming` uses "steer" in a different sense: offering a menu of options steers the human, so the first move is a question that gets them describing. That is not a field on the section, and it is not the live-step line. The next fetch does not include either one. Session history is what the skill says never to hand over, and the session-start injection is the parent's guard, which the child is told to ignore.

The human rail side record is the ledger of rulings, fix rounds, parked findings, and deferred minors. The controller writes it. A child packet cannot. The human reads it at the end, in the "Rulings I made" list, not by opening the section. The next fetch may include a pointer to one ledger entry, and only when an earlier parked finding touches this task. The final reviewer is pointed at the ruling lines so it can weigh them. The body of the ledger is not pasted into the child. `requesting-code-review` says the same thing for the reviewer: precisely crafted context, never the session's history.

The section Rail is the switch only the human flips. On a build section it is the design or plan approval before the first implementation tool. It is also explicit consent before implementation on `main` or `master`. At the end, `finishing-a-development-branch` will not merge, push, or discard on its own. After a green suite it presents a menu and waits: merge locally, push and open a pull request, or keep the branch. Discard is not on that menu. It runs only after the human asks, and only the typed word `discard` authorizes the deletion. A failing suite does not reach the menu. An empty Rail means none of these switches is set, which is the normal state of a mid-plan task. The child's fetch does not include the parent's menu. Global constraints copied into the brief are spec text inside Do, not this switch.

So the next fetch includes this task's Section, Function, Read, and Do, the parent-named report path, and at most a pointer into the side record. It does not include the steer line, the ledger body, or the Rail switch. Next stays off that packet until the controller has a `Task N: complete` line.

## Adopt

Suggestion. Before the plan store, Section, Function, Read, Do, and Write are non-empty and are not placeholders. The first implementation tool on a build section also waits until Rail is the human's approval. Next is not required for either gate. A completion line is stored only after a fresh check of Do passes.

Suggestion. When a lead cuts NEXT, the lead stores a ledger row (a fix round or a parked ruling) or, if the check failed first, stores no completion line at all. The lead does not write the plan corpus. The record left is the missing `Task N: complete` line, so no next id is stored. The child does not store that cut.

Suggestion. A parent hands a child only that task's brief, the interfaces the brief cannot know, the binding constraints inside that brief, and the report path. The child may not touch the parent's Write, Next, or Rail: it does not write the plan or the ledger, it does not read or choose the next task, and it does not flip the human switch or dispatch the reviewer.

Suggestion. The steer line, the ruling ledger, and the section Rail stay three objects. The next fetch includes none of their bodies. A pointer to a ledger entry is allowed only when that entry touches this task.

These four are the same scaffolding claims as the JSON at the end of the research turn. They are still one source.

## Refuse

Refuse letting the controller pick or escalate the child's model. The executor skill says to name a model on every dispatch, to use a cheaper tier for transcription, and to bump a more capable model on fix rounds 4 and 5. That chooses a model. The mesh does not. A seat already has its direct id. The fix cap and the ledger ruling can stay without the bump.

Refuse a second implementer on the same plan at the same time. The skill forbids parallel implementation dispatches because the writes conflict. Batching several tasks into one brief is the same breach of one task, one seat. Independent investigations with no shared Write are a different skill and are not this packet.

Refuse any seat delete. The executor deletes the plan workspace when the final review is clean (`rm -rf` of the scratch directory) and says git history is then the record. The finishing skill deletes a worktree on merge, and on the typed word `discard` it force-deletes the branch. The README's test-driven step deletes code that was written before a test. A seat has no delete grant. The progress ledger is the record that has to survive, not a scratch directory the controller throws away. Showing the human a refused worktree removal, and waiting, is the rail. Doing the removal is not a seat action.

Refuse a spike, or any other write, that skips the plan. `brainstorming` allows a throwaway spike with no tests and no review, then says keeping it is a new request. An inline plan before any write is already the rule. The spike does not become a path on this packet.

Refuse pasting the session, the steer line, the ledger body, or the finish menu into the child. The skill's own failure mode is a dispatch that is almost all history. Also refuse treating `using-superpowers` as the child's job. A subagent is told to ignore it.

Refuse importing the repo's contribution policy as packet fields. `AGENTS.md` requires a filled pull-request template, a search of open and closed requests, a disclosed model and harness, and a human looking at the full diff before submit. The human-looking-at-the-diff part is already the finish menu's push switch. The rest is that repo's review desk, not a section.

Refuse the visual-companion fetch and its version telemetry. It is optional, off when `SUPERPOWERS_DISABLE_TELEMETRY` is set, and it is not a store, a fetch of a section, or a rail.

## Park

Park `diagnosing-superpowers` and the eval harness `AGENTS.md` points at. The README says that skill reads a transcript after a session has gone wrong and can export a scrubbed bundle. `AGENTS.md` says skill changes are graded by a separate drill harness. That is a later loop: it reads a record that already exists, it runs after the work, and it is not a step on the live packet. This note did not open the eval repo. Do not mount it, and do not add a trace field to the section.

Park marketplace installers, per-harness plugin commands, and the commercial-support offer. They put the same skills on someone else's agent. They do not change what a packet must hold.

Park `dispatching-parallel-agents` as a later isolated read, only where there is no shared Write and no shared Next. It is not how a plan's tasks advance. The executor already forbids that shape on one plan.
