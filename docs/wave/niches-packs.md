# Niches as memory and skill packs

## Question

How should den grow niches (design, webperf, Go, Svelte, shadcn, Mantine, DNS, crypto, email, copy) as memory and skill packs without loading the catalog into the packet?

## Sources

- `plan/niches/INDEX.md`. One agent, one niche file, one skill. A line becomes implementation only after two finished notes agree and a check passes.
- `skills/use-harness/SKILL.md`. Read this file, then one reference. Do not load the catalog into the packet. Skills load one body when the description matches.
- `skills/select-skill/SKILL.md`. Return one path. Do not return a skill whose description does not match. If nothing matches, attach `skills/use-harness/SKILL.md` only.
- `skills/develop-skill/SKILL.md`. List names, do not read every body. If one matches, attach it. If none match, add one row and one short `SKILL.md`. The catalog is not pasted into the prompt. Web performance is already `~/.agents/skills/web-perf`. Do not copy that text into den.
- `plan/niches/prompt-analysis.md`. Keep the seat prompt short. Do not paste the niche index into it.
- `plan/BASE.md` and `skills/use-harness/references/memory.md`. Fast and recent are local files. OpenViking is the long store. The OmniRoute host is the backup, not the read path.
- `research/vendors-west.md`. The skills spec's on-demand load is already decided. Do not reopen it. `license`, `compatibility`, `metadata`, and `allowed-tools` stay out of the seven fields.
- External, one source: [Agent Skills specification](https://agentskills.io/specification). A skill is a directory with `SKILL.md`. Agents load `name` and `description` at startup, then the body only after the skill is activated. `scripts/`, `references/`, and `assets/` load on demand. Confirmed by TinyFish search and fetch on 2026-10-12.

## Finding

A niche names one skill. The catalog is the index and the route table. A packet carries one path, not the table.

Grow each named niche as its own pack, one packet at a time, through `skills/develop-skill`:

1. Match an installed skill by description first (`web-perf` already covers webperf). Attach that path. Do not vendor its body.
2. If none matches, add one `SKILL.md` whose description says what it does and when to use it, plus one row in `skills/select-skill`. One job. Design, Go, Svelte, shadcn, Mantine, DNS, crypto, email, and copy are nine later packets, not one catalog write.
3. Keep the long procedure in `references/` and scripts on disk. The seat opens the body, then one reference.
4. Global memory is a promoted principle: one short line in local fast memory, after two notes agree. It is not the skill body, not the niche index, and not a paste into the next packet. The body stays the skill. OpenViking keeps the long note. A later note supersedes. Nothing is deleted.

The seat prompt stays as it is: open the named skill, or `select-skill` then one path. Safety limits stay copied. The niche index stays out.

## What den should not copy

- The Agent Skills startup scan of every skill's `name` and `description` into the seat. den already routes with one table and one path. Do not add a second catalog to the packet.
- Optional spec fields (`license`, `compatibility`, `metadata`, `allowed-tools`) as packet fields.
- Installed skill text (`web-perf`, Vercel, Cloudflare) copied into the repo.
- A curator that rewrites skill bodies from traces.
- A niche that names two skills, or a pack that loads design plus webperf plus the UI kits in one seat.

## One plan line

One niche, one skill, loaded when its description matches; promote a principle into local fast memory, and leave the body and the catalog on disk.

## Unproven

- No check yet that `select-skill` returns a single path for these ten names. webperf is the only one with an installed path in `develop-skill`. The other nine have no `SKILL.md` and no niche file.
- Not shown that a promoted principle in fast memory is read on the next seat without pulling the skill body with it.
- The spec's activation rule is external. den's own rule is the local files above. This note does not move the spec into the plan.
