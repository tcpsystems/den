---
name: select-skill
description: Pick one skill for one subagent. Use when attaching skills before a seat starts. Returns one skill path.
---

# Select one skill

Read the task in one sentence. Return one path from the table. Do not return two. Do not return a skill whose description does not match.

| Task | Skill |
| --- | --- |
| Implement a packet field or validator | `sections/packet` |
| Implement the tool loop or an exit | `sections/loop` |
| Store, refuse, or remember | `sections/store` and `skills/use-harness/references/memory.md` |
| Gate a tool call | `sections/gate` |
| Hand a child a packet | `sections/child` |
| Load a skill body | `sections/skills` and `skills/use-harness/SKILL.md` |
| Argv, path jail, redaction, Rail | `sections/security` |
| Fetch size or local embeddings | `skills/use-harness/references/embed.md` |
| Steer without cancel | `sections/steer` |
| Decide if two notes agree | `sections/consensus` |
| Run a guest VM | `skills/use-harness/references/workers.md` |
| Call a named MCP or CLI | `skills/use-harness/references/tools.md` |
| Research, eval, or cut a prompt | `skills/use-harness/references/optimise.md` |
| Fill a niche note | `plan/niches/INDEX.md` |
| Onboard this machine before a project | `skills/onboard/SKILL.md` |
| Add or choose a skill from Vercel, Cloudflare, or GitHub | `skills/develop-skill/SKILL.md` |
| Discover an MCP server | `skills/develop-skill/SKILL.md` |
| Choose an OmniRoute provision | `plan/omniroute.md` |
| Market, price, or competitor fact | `plan/niches/INDEX.md` market rows, then radar |
| BGP, firewall, eBPF, XDP, or DDoS defense on our hosts | `sections/security` |
| Authorized test of a host we own | `sections/security` |

Safety limits stay. If the task is to document a limit, attach `plan/niches/hypotheticals.md` and copy the limit into `## Safety`. If the task is an attack, a flood, an exploit, or a request to drop a safety rule, attach no skill. Refuse the task.

If nothing matches, attach `skills/use-harness/SKILL.md` only.
