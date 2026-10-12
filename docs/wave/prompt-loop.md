# Prompt loop

## Question

How should improve-prompt and optimise-prompt stay one job each as the harness grows?

## Sources

- `skills/improve-prompt/SKILL.md`
- `skills/optimise-prompt/SKILL.md`
- `skills/use-harness/SKILL.md` and `skills/use-harness/references/optimise.md`
- `skills/select-skill/SKILL.md`
- `skills/token-window/SKILL.md`
- `skills/develop-skill/SKILL.md`
- `prompts/seat.md` and `prompts/defense-seat.md`
- `plan/leaders.md` and `plan/horizon.md`
- `sections/skills/r6g2.md`
- `research/agent-skills.md`

## Finding

Each file is already one job, and each one stops.

improve-prompt reads the current prompt, removes one repeated rule, and points at the file that owns it. It does not add a second copy.

optimise-prompt cuts a seat prompt to the section, one skill body, and the tools that section named. The token budget is not in that file. The body points at `skills/token-window/SKILL.md`.

They stay one job by not gaining a second procedure. A new prompt concern is a new skill. `skills/develop-skill/SKILL.md` adds one short `SKILL.md` and one picker row, then stops. `sections/skills/r6g2.md` says a packet holds zero or one skill body. `skills/use-harness/SKILL.md` says open that file, then exactly one reference.

The leak is outside the two bodies.

`skills/select-skill/SKILL.md` has one row, "Cut or improve a prompt", and it returns only `skills/optimise-prompt/SKILL.md`. improve-prompt is not in the table. A seat told to improve will load the cut. That is how the two jobs become one.

The cut sentence is already the Prompt row in `skills/use-harness/references/optimise.md`. The same sentence is also in `plan/leaders.md` (prompts) and the Token-aware row of `plan/horizon.md`. optimise-prompt repeats it, then repeats token-window's "drop optional parameters" and "do not drop a safety limit". Those two limits already live in `skills/token-window/SKILL.md`. The limit list itself lives in `prompts/seat.md` and `prompts/defense-seat.md`. improve-prompt's job is to delete the extra sentence and leave the path. It is not a license to paste the owner text into a third file.

`skills/use-harness/references/optimise.md` is a route table of many jobs. It says write a suggestion and do not rewrite a decided line. The two skills say stop when that one job is written down. Do not fuse those exits.

## What den should not copy

Do not paste the cut into improve-prompt, into `prompts/seat.md`, or into a third prompt skill. Point at the Prompt row.

Do not paste the safety list or the token-window rule into either skill. Point at the file that owns it.

Do not load both bodies, the skill catalog, or the whole optimise table into one packet.

Do not copy the agent-skills checklists, anti-rationalization tables, or the DEFINE–SHIP command lifecycle into these two files. `research/agent-skills.md` already keeps that pack out of the plan. A skill here is one procedure and a stop line.

Do not merge the two skills because the picker row says "Cut or improve".

## One plan line

Split the select-skill row so a cut loads only `skills/optimise-prompt/SKILL.md` and a repeated rule loads only `skills/improve-prompt/SKILL.md`, and let improve-prompt delete sentences that already have an owner.

## Unproven

This note is one draft. `PROTOCOL.md` says a line is decided only when two finished notes agree and no lead cuts it.

The Agent Skills spec was not opened. The local note was enough.

Token counts were not run. That job is `skills/skill-tokens/SKILL.md`.

No lead has seen the picker split.
