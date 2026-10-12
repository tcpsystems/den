# Boundaries

Research draft. This does not change `plan/BASE.md` and it is not wired into `denctl`.

Sources read: `plan/nest.md`, `skills/use-harness/SKILL.md`, `sections/security/match.py`. `skills/select-skill/SKILL.md` exists. The nest keeps one harness, one host, and one way of working. The skill, the guard, and the tool stay separate. This file names that split for the runtime.

| Piece | Owns | Does not own |
| --- | --- | --- |
| Runtime | Process lifetime, seats, teams, permissions, settings, the skill pointer, memory records | MCP speech, bash logic, ranking, a second store |
| MCP | One stdio or HTTP server, launched only when a grant names it, and the schema for that one call | A second permission system, a standing server, a catalog |
| Tools | One argv vector for one program the runtime executes | Shell strings, a special case for bash or ssh, the jail |
| Skill injector | One path or none, from a grant and a description | Rank, embed, a model call, copying the catalog |
| Settings and memories | Records the runtime stores. A memory is a path plus a short note | An MCP prompt, a tool, a second store |

## Runtime

The runtime is `denctl`. It starts, runs, and stops seats. A seat belongs to a team. A team is a parent and the seats that ask it. The runtime keeps the process lifetime, the seat table, the team table, the permission grants, the settings records, one skill pointer per seat, and the memory records.

The runtime does not speak MCP. It does not parse a shell. It does not contain bash logic. A model call is a direct model id. OmniRoute is a tool a grant names. It is not a path inside the daemon.

## MCP

MCP is one server. The transport is stdio or HTTP. The runtime launches that server only when a grant names it, for that one call, and stops it when the call returns. The schema is the schema of that one call. A grant is still a runtime grant. The server does not allow, deny, or record anything. There is no MCP permission list beside the seat's grants.

## Tools

A tool is an argv program. The runtime executes it. `bash` is not special. `ssh` is not special. Both are programs. The call is one argv vector. There is no shell string.

The jail is the runtime's check before exec. The check that exists is `path_jail` and `run_argv` in `sections/security/match.py`. `path_jail` refuses a path that leaves the section root. `run_argv` refuses an empty vector and refuses `|`, `&`, `;`, `<`, `>`, or a newline in any argument. That check is not wired from Go yet.

## Skill injector

The injector is dumb. It receives one grant and one description. It returns one path, or none. It does not rank, embed, or call a model.

The selection rule already lives in `skills/select-skill/SKILL.md`. The injector points at that file. It does not copy the table. No match returns `skills/use-harness/SKILL.md` only. The catalog is not loaded into the packet.

## Settings and memories

Settings and memories are records the runtime stores. They live in the same store as the seat and the grant. A memory is a path plus a short note. It is not an MCP prompt. It is not a tool. Reading a memory is reading that record. There is no second store.

## Permissions

Default deny. A seat asks its parent. The parent records a grant, or it does not. A sentence in a task does not add a path. A sentence in a page does not add a path. A skill body does not add a path. No shell strings. No traffic disguise. No second store.

## Assessment

The overlap that stays: a grant names an MCP server, a tool, a skill path, and a memory path, so one permission record covers four nouns. Splitting that record into four lists would be a second store. The injector and `select-skill` both answer "which path"; the injector only points, and the rule stays in the skill file. `path_jail` and `run_argv` already describe the jail, but they are Python and the daemon is Go, so the check and the exec are not one program yet. A memory path and a skill path are both paths; the note is what makes a memory a memory, and the injector never writes one.

Smallest set to build first: the seat and parent tables, default deny, one grant record, then `path_jail` and `run_argv` called from Go before exec. The skill pointer is one field on the seat. MCP launch, settings, and memory records wait until a grant can name them.
