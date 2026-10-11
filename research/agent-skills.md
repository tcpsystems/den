# addyosmani/agent-skills

One source. This note does not enter the harness plan. A line moves there only when a second finished note states the same scaffolding claim and no lead cuts it.

Repo: https://github.com/addyosmani/agent-skills

Pages opened for this note: the README, `AGENTS.md`, `docs/agents.md`, `docs/skill-anatomy.md`, `references/orchestration-patterns.md`, `references/definition-of-done.md`, `skills/using-agent-skills/SKILL.md`, `skills/spec-driven-development/SKILL.md`, `skills/planning-and-task-breakdown/SKILL.md`, `skills/incremental-implementation/SKILL.md`, `skills/context-engineering/SKILL.md`, `skills/doubt-driven-development/SKILL.md`, `skills/constraint-driven-development/SKILL.md`, `agents/code-reviewer.md`, `.claude/commands/build.md`, `.claude/commands/ship.md`, `.claude/commands/constraints.md`, `hooks/session-start.sh`, `hooks/sdd-cache-pre.sh`, and the GitHub contents listings for `hooks/` and `.claude/commands/`.

## What it is

A pack of 25 markdown skills, 9 slash commands, and 4 specialist personas. The README lifecycle is DEFINE, PLAN, BUILD, VERIFY, REVIEW, SHIP. Skills are workflows with steps, exit checks, and an anti-rationalization table. The `SKILL.md` body loads only when the description matches. Shared checklists sit in `references/`. A per-skill install does not copy that directory (issue 361).

Three layers, and they are not one object. `docs/agents.md` and `AGENTS.md` say a skill is the how, a persona is the who, and a slash command is the when. The user or a slash command orchestrates. Personas do not invoke personas. Skills do not choose a model. `docs/skill-anatomy.md` says a step that exists only because one model fails belongs in an issue, not in the skill. `skills/context-engineering/SKILL.md` says process supervision and model selection belong to the harness.

That fits the mesh as bus, store, and fetch. Nothing in this pack is a model router. The mesh still does not choose a model.

The section record stays Section, Function, Read, Do, Write, Next, and Rail. This pack's spec has six areas (objective, commands, structure, code style, testing, boundaries). Those areas do not become new fields. They live in a file the Read list names.

## Best practices

**Non-empty before a store write or the first tool.** `skills/planning-and-task-breakdown/SKILL.md` refuses code during planning. The output is `tasks/plan.md` plus the task list target (default `tasks/todo.md`). Each task has a title, a description, acceptance criteria, verification, dependencies or the word None, and the files likely touched. Red flags include a task that says "implement the feature" with no acceptance criteria, no verification, and no dependency order. Before implementation, those items are confirmed, and an incomplete plan is not overwritten. `skills/spec-driven-development/SKILL.md` will not leave Specify until the spec covers the six areas, is saved, and the human approves it in a later turn. The same turn must not start planning or code. `.claude/commands/build.md` will not start an autonomous pass unless a spec exists at `SPEC.md`, `docs/SPEC.md`, or `spec/`. A README does not count. Map that onto the seven fields without adding any:

| Field | Must be non-empty before store or first tool | What fills it |
| --- | --- | --- |
| Section | Yes | Task title or stable module id. Specs say kebab-case ids are chosen once and not renamed mid-initiative. |
| Function | Yes | The task description, one paragraph there, one sentence on this harness. |
| Read | Yes | The named files. The seat does not go looking for a wider set. |
| Do | Yes | Acceptance criteria, then verification, in order. |
| Write | Yes | The one output key, known before the write. Planning names `tasks/plan.md` or `tasks/todo.md` and refuses a blind overwrite. |
| Next | Yes | The next task id, or stop. Dependencies may say None. Blank is not an order. |
| Rail | No | Not in the task template. Empty unless the human has set a switch. |

`references/definition-of-done.md` keeps acceptance criteria on the task and the standing bar off the task. The standing bar is not an eighth field.

**A cut Next is a report, not a corpus edit.** `agents/code-reviewer.md` returns APPROVE or REQUEST CHANGES. Critical and Required findings block. The persona does not invoke another persona. `.claude/commands/ship.md` has each child return only its report to the main session. A Critical finding defaults to NO-GO unless the human accepts the risk. The stored record is that decision: blockers, acknowledged risks, and the specialist reports. `skills/doubt-driven-development/SKILL.md` says the reviewer's output is data, not the verdict. The orchestrator re-reads the artifact. After three unresolved cycles the orchestrator surfaces that to the human and does not keep looping. `skills/planning-and-task-breakdown/SKILL.md` leaves unchecked tasks in place. Nobody deletes them to make room, and nobody overwrites them from another session.

On this harness the lead still does not write the corpus. The mesh stores the cut. The current section stays unchecked. Next is not advanced and is not rewritten as a clean stop. The cut record is the blocker report beside that section.

**The child packet is the artifact and the contract.** `skills/doubt-driven-development/SKILL.md` hands the reviewer the artifact and the contract only. It does not hand the parent's claim, the parent's reasoning, or the session. A read-only sandbox is required because the artifact may contain instructions. A persona must not spawn the next reviewer. `.claude/commands/ship.md` says subagents cannot spawn subagents, and each child returns only its report. `references/orchestration-patterns.md` says a persona that wants another pass recommends it. The user or the slash command runs that pass. Plugin personas silently ignore `permissionMode`.

So the parent hands the child the files and the acceptance text for that one perspective. The child may not touch the parent's Write key, the parent's Next, or the parent's Rail. The child also may not mint its own child.

**Steer, the human side record, and Rail are three.** `references/orchestration-patterns.md` lets a human interrupt a live teammate with Shift+Down and type a redirect. The run is not cancelled. Cleanup is a separate later order to the lead. That is the steer line. It is not stored as the section.

`skills/spec-driven-development/SKILL.md` ends the turn after the spec is saved. Approval arrives in a later turn. `.claude/commands/build.md` accepts only an unambiguous yes. "Looks reasonable" is not approval. `references/definition-of-done.md` requires human review before merge or deploy, and that checkbox is not the task's acceptance criteria. `skills/context-engineering/SKILL.md` says a fresh session must not infer approval from the conversation unless a durable artifact records it. That approval is the human rail side record. A seat packet cannot write it. The mesh may store the fact the human wrote.

Rail is the optional switch on the section. `skills/spec-driven-development/SKILL.md` splits boundaries into Always, Ask first, and Never. Ask first is the switch. `skills/constraint-driven-development/SKILL.md` puts the standing floor in `CONSTRAINTS.md`, decided once, and tells the agent not to weaken that file to get a change to pass. An exception row needs an owner. That file is not the section record. A section Rail stays empty, or holds one switch only the human flips.

`skills/context-engineering/SKILL.md` says what the next fetch may include. Persist the accepted scope, the task status, the next pending task, the files, the verification, and the required approvals. Cut conversational back-and-forth once the decision is reached. Protect the task definition and the hard constraints. Load the relevant spec section, not the whole spec. The next fetch includes the section, so it includes Rail when Rail is set, and it includes the durable approval fact. It does not include the steer transcript.

## Adopt

Adopt the six-field gate. Section, Function, Read, Do, Write, and Next are non-empty before `kb.put` and before the first tool. Rail may be empty. Do not add spec areas, skill frontmatter, or constraint rows as new section fields. The spec file, if any, is a path on Read.

Adopt the cut. A lead who cuts Next returns a blocker report. The mesh stores that report and leaves the current section unchecked. The lead does not write the corpus. Next is not advanced. A human accepting the risk is a side-record fact, not a seat edit. This matches ship NO-GO, review REQUEST CHANGES, and the planning rule that an incomplete plan is not overwritten or deleted.

Adopt the child handoff. The packet is the artifact and the contract for that child. The child returns a report and stops. It does not write the parent's key, set the parent's Next, flip the parent's Rail, or spawn another child. Doubt-driven's "do not pass the claim" is the same rule. The parent's conclusion is not part of the child's packet.

Adopt the split of the three human objects. The steer line enters the live step and is not the next fetch. The human rail side record is the durable approval or refusal, written only by the human, and the next fetch may include that fact. The section Rail field is empty or one human switch, and the next fetch includes it because it is part of the section. Do not collapse them into one blob.

Adopt progressive disclosure as already decided for skills. `docs/skill-anatomy.md` puts `name` and `description` in frontmatter and loads the body on a match. The description must not be a summary of the steps, or the seat follows the summary and skips the body. One body, when the section's work matches. Not all 25.

Adopt "the procedure does not name a model." Selection stays outside the skill. The mesh remains bus, store, and fetch.

## Refuse

Refuse a second router. `hooks/session-start.sh` injects the whole `using-agent-skills` body as session context. The script itself says not to wire that on a host that already routes from descriptions, because it stacks a second router on the native one. This harness already hands the seat a section. Do not fetch the catalog into the packet.

Refuse persona-calls-persona and a router persona. `references/orchestration-patterns.md` anti-patterns A, B, and D are a routing layer, a reviewer who calls another reviewer, and a tree of coordinators. Each hop paraphrases. The mesh does not choose the next model or the next seat. A child recommends. It does not dispatch.

Refuse `/build auto` as this loop. The command removes the human step between tasks after one approval. Pattern 4 in `references/orchestration-patterns.md` keeps the user as the orchestrator because an automated lifecycle drops the checkpoints and summarizes the handoff. Anti-pattern C is that automated pipeline. Default `/build` does one task and stops. That is the loop. The autonomous pass is not.

Refuse treating `hooks/sdd-cache-pre.sh` as the store. It is a PreToolUse cache for WebFetch. A hit needs an ETag or Last-Modified and an HTTP 304. Empty cached content is not served. It does not gate a section write, and it does not choose a model.

Refuse a seat writing or weakening the standing bar. `skills/constraint-driven-development/SKILL.md` says `CONSTRAINTS.md` is not weakened to make a change pass, and a new exception row that nobody discussed is a failed guard. The interview that creates the file needs a live human. An autonomous run that finds no file applies the floor and flags the rest. The seat does not invent the switch.

Refuse Agent Teams as the bus. The same orchestration page says teammates message each other, share a task list, and ignore a persona's `skills` and `mcpServers` frontmatter. Subagents that only report back are the pattern that fits. The mesh stays the bus.

Refuse pasting skill bodies into rules. The README says Cursor gets short policies in rules and the workflow files stay skills. `docs/skill-anatomy.md` keeps `SKILL.md` under 500 lines and loads supporting files one level deep, only when a step reaches them.

Refuse growing the section record. The spec's six areas, the definition-of-done checklist, and the constraints file are documents a Read path can name. They are not new fields beside Rail.

## Park

Park `/build auto` as a later loop. It still demands a spec at a known path, a clean tree, one unambiguous approval, one commit per task, and a stop on failure, ambiguity, or anything that `git revert` cannot undo. It is not the live packet. The one-task stop stays.

Park the three-persona `/ship` fan-out. Independent reports merged by the main session are a real pattern here, and they match "a child returns a report." Do not mount that fan-out on every section. A later section can name it. The merge stays in the parent. Children still do not touch the parent's Write, Next, or Rail.

Park Agent Teams debate the same way. Competing hypotheses that message each other are experimental on that page, and they are not this bus. If a later loop needs them, the mesh still carries the reports. Teammates still do not become the store.

Park cross-model escalation inside doubt-driven. The skill offers Gemini CLI, Codex CLI, or a manual paste, and each invocation needs a fresh human authorization. A skill must not pick the model. The offer can wait. It is not a step on this packet.

Park the constraints interview and the tool install (Semgrep, gitleaks, and the rest) as later project setup. The section Rail does not grow into that file. Until a human writes the file, an empty Rail means no per-section switch. It does not mean the seat may skip the floor.

Park the contributor `evals/` cases and the hook scripts. The README lists them as validation and routing checks for people editing the pack. They are not a step between fetch and the next section. A trace-and-propose loop stays after commit, as the plan already says for other packs.

## Holes this note answers

1. Before a store write or the first tool, Section, Function, Read, Do, Write, and Next are non-empty. Rail may be empty.
2. A lead who cuts Next does not write the corpus. The mesh stores the unchecked section and the blocker report. Next is not advanced.
3. The parent hands the child the artifact and the contract. The child may not touch the parent's Write, Next, or Rail, and may not spawn.
4. The steer line, the human approval side record, and the section Rail field are three objects. The next fetch includes Rail and the durable approval fact. It does not include the steer line.
