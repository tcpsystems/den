# Registry layout

Written 2026-10-12. This file is research. It does not change `plan/BASE.md`, and it does not add a skill, a niche, a runtime, or a tool.

## What is copied

Two public layouts already split a shared catalog from a project pin.

shadcn/ui keeps a registry as data. The documented schema is `registry.json`. A large registry can be split across several of those files with `include`. A public GitHub repository can be a source. The project file `components.json` is where one app says which registries it uses. Docs: https://ui.shadcn.com/docs/registry/registry-json

The skills CLI installs one pack into the agent that is running. `npx skills add vercel-labs/agent-skills` is the documented form. The CLI source is https://github.com/vercel-labs/skills and the docs are https://www.skills.sh/docs. Packs bundle public and private skills into one install command. The site says each skill should be read before install, and that a listing is not a security guarantee.

den already has the same split in words. `plan/niches/skills-index.md` is a path list. `skills/` inside this checkout is the project pin. `~/.grok/skills` and `~/.agents/skills` are the global pin on this machine. A packet loads one path. It does not load a catalog.

## Namespaces

A name is `source/item`. The source is a registry the human added. The item is one file.

| Scope | Where it lives | Who writes it |
| --- | --- | --- |
| Global | A registry file outside the checkout, one row per item, with a URL and a licence line | The human, or a note that two later notes have agreed |
| Project | `skills/<name>/SKILL.md` and one row in the project index | A seat, one file, after the licence line is read |
| Packet | One path | `select-skill`, or the human name, which wins |

A project row shadows a global row with the same item name. A packet never sees both.

## Item

One row. No second copy of a rule that already has a file.

| Field | Meaning |
| --- | --- |
| name | `source/item` |
| kind | skill, profile, niche, or note |
| path | File inside the project, or empty when the row is only global |
| url | Page or repository the body came from |
| licence | The line the reader quoted |
| check | The local pass or fail, or empty |

An empty check means the row is a candidate. A candidate is not loaded.

## Install

The shape of `npx skills add` is the shape to keep: one source, one command, files land in the project pin. The command is one argv vector. It refuses a shell string. It does not run a post-install script from the registry.

Order:

1. Read the registry row.
2. Read the licence on the url.
3. Write one file under the project pin.
4. Append one index row.
5. Stop.

A second item is a second command.

## What this note does not do

It does not vendor a registry. It does not rank items. It does not add a telemetry call. skills.sh ranks by anonymous install counts; that ranking is theirs, not a score on this machine. It does not change `develop-skill` or `select-skill` until two notes agree.
