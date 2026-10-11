---
name: develop-skill
description: Grow the skill set by one. Discover, assess, and append a single skill reference. Do not paste a catalog into the packet.
---

# Develop one skill

Run this loop once. Then stop. A second pass is a new packet.

1. Discover. List skill names already on the machine. Do not read every body.
2. Analyse. One sentence: which name matches the task.
3. Evaluate. If a den skill or an installed skill already matches, attach that path. Do not write a new one.
4. Append. If none match, add one row to `skills/select-skill/SKILL.md` and one short `SKILL.md`. One job.
5. Reiterate only when the check fails. Fix that skill. Do not start a second topic.
6. Solidify. The new skill has a description, one path, and a check. Then it is done.

## Already installed

Use these. Do not copy their text into den.

| Task | Path |
| --- | --- |
| Cloudflare Workers, Wrangler, Durable Objects, D1-adjacent docs | `~/.grok/skills/cloudflare`, `workers-best-practices`, `wrangler`, `durable-objects` |
| Web performance | `~/.agents/skills/web-perf` |
| Vercel deploy, React composition, cost | `~/.agents/skills/deploy-to-vercel`, `vercel-composition-patterns`, `vercel-optimize`, `vercel-react-best-practices` |
| Page and prose review | `~/.agents/skills/web-design-guidelines`, `writing-guidelines` |
| Fetch a public page | `~/.agents/skills/use-tinyfish` |
| This harness | `skills/use-harness`, `skills/select-skill`, `skills/onboard` |

GitHub has no skill directory here. Use the `gh` command. Do not invent a token file.

MCP discovery is `meshctl suite check` and the `mcp_servers` entries in `~/.grok/config.toml`. Report the server name and whether it is enabled. Do not print a bearer token.

## Check

The packet grew by one skill path, or by zero because a match already existed. The catalog was not pasted into the prompt.
