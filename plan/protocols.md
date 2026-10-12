# Runtime protocols

These names are fixed. A new idea gets a new name. It does not reuse one of these.

| Name | Means | Does not mean |
| --- | --- | --- |
| identity | Stable who. Id, name, avatar. | Role, niche, or model. |
| role | What the seat may commit and own on this wave. | A job title in prose. |
| niche | The subject the seat is staffed for. | A skill body. |
| profile | The bound bundle: identity, role, niche, model, skills, mcp, memory. | A running process. |
| skill | One markdown playbook, loaded when the step matches. | A binary or an MCP server. |
| tool | One argv inside the path jail. | An MCP server. |
| mcp | An indexed server started only when a grant names it. | A skill. |
| plugin | A directory that adds catalog rows. | A change to the vocabulary. |
| memory | Facts the seat may read. Disk ledger first. | The embedder process. |
| context | The packet for one step. One skill body. | The full catalog. |
| grant | A signed contract. Seat, parent, reason, scope, deadline. ed25519. | A role, or a secret value. |
| model | A route id. | The gateway. |
| seat | One worker on one branch in one wave. | A profile. |
| wave | The shared branch the orchestrator merges seats onto. | A chat. |

## Multi-agent

1. The session orchestrator writes a profile for each seat.
2. `denorch open` gives each seat its own branch and worktree.
3. The seat reads its context: identity, one skill, the task, and the paths it owns.
4. The seat commits with `denorch commit`. The role picks the commit type.
5. The orchestrator runs checks and `denorch take` merges onto the wave.
6. A grant is a signed contract. The parent signs the canonical JSON with ed25519. The runtime pins the parent public key. The contract names the seat, the reason, the scope, and the deadline. Scope items look like `mcp:exa` or `secret:NAME`. The value stays in the runtime store. A prompt sees the name only.

## MCP providers

`mcp/catalog.json` lists many providers. None is required. SearXNG stays the local search engine. `denfetch` stays as it is until a later expansion. `denmcp spawn` starts the granted containers in parallel. Adding a provider is a new catalog row, not a change of the default.

## Release

`VERSION` and `plan/RELEASE.md` name what the binaries do. `plan/GOALS.md` names what is still unwired. A goal moves only when a test or a command shows it.

## Extending

Add `plugins/<id>/plugin.json` with new skill, tool, mcp, or niche rows. Do not add a fifth role until the four roles cannot say the job. Do not put Foxcore repos in this catalog.
