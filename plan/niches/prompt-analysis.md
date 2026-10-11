# Seat prompt analysis

The injected prompt is `prompts/seat.md`. It is the only text every seat sees before the packet. This note is what that file holds, what it drops, and the smallest fix.

## What holds

- The model id is given. A combo is used only when the packet names it.
- Fast and recent memory are local. The OmniRoute host is not the read path.
- One skill file, then one reference. The catalog stays out of the packet.
- The seven fields are named. A child cannot write the parent's Write, Next, or Rail.
- Stop is the packet's Write, or Next equal to stop.
- A failed tool is a refusal row, not a new plan.

## What it drops

- The safety limits that live in `plan/niches/INDEX.md` are not in the prompt. A seat that only reads `prompts/seat.md` never sees them.
- The match runner is not named. A seat does not know that a premise submits one disallowed action and the spectator writes both actions.
- "Shell is argv" uses the word shell for a rule that refuses a shell.
- The skill path is relative. A worktree whose current directory is not the repo root opens the wrong file.
- Check, test, and reeval from `plan/architecture.md` are not in the prompt, so a seat can stop at a note.
- `skills/select-skill` is a route row, but the prompt does not say to use it when the packet does not name a skill.

## Fix

Keep the prompt short. Add four lines, not the niche catalog.

1. Paths are repo-root paths.
2. Safety limits are copied and not dropped. Name argv, the section root, Rail, and the owned-host test rule.
3. If the packet names a skill, open that. Otherwise open `skills/use-harness/SKILL.md` and one row. Use `skills/select-skill/SKILL.md` when the packet does not name one.
4. A match is `sections/security/match.py`. The writeup is the premise action and the defense action.

Do not paste the niche index into the prompt. That is the catalog, and the route skill already points at it.
