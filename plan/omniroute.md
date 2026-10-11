# OmniRoute provisions

OmniRoute is a tool the packet names. It is not the seat's model hop. A provision is used only when the rule below says yes. Otherwise it stays off.

| Provision | Use when | Do not use when | Skill |
| --- | --- | --- | --- |
| Direct model id | The packet names one model | A combo would add a hop for a single call | `prompts/seat.md` |
| Combo `harness-winner` | The packet says the call may move among the scored models | The seat has a direct id, or latency matters more than a retry | `plan/omniroute.md` |
| Batch | Many completions are already written, and none need a tool in the middle | The seat is in a tool loop | `skills/use-harness/references/tools.md` |
| A2A | The tool is a remote agent card at `https://ai.tcp.systems/api/a2a` | The other party is a den seat. Seats talk on the mesh | `skills/use-harness/references/tools.md` |
| MCP | The step needs a gateway tool, and `meshctl suite check` shows it | A local CLI already does the step | `skills/use-harness/references/tools.md` |
| Agent-skill catalog | The task is about OmniRoute itself and `select-skill` names one OmniRoute skill id | The task is a den section. Use a den skill | `skills/select-skill` |
| Compression or prompt rewrite | The leader already cut the prompt and the packet still asks for a smaller body | As the default on every call | `skills/use-harness/references/optimise.md` |
| Embeddings | Never. Vectors are `127.0.0.1:8091` | Always | `skills/use-harness/references/embed.md` |
| Memory read | Never. Fast and recent are local. OpenViking is local | Always | `skills/use-harness/references/memory.md` |
| Persistence copy | After a local note is written, the backup script copies fast and recent to the OmniRoute host | As the seat's read path, or before the note exists | `scripts/backup-fast-memory.sh` |
| Model score | After local probes in `scripts/score-harness-winner.py` | To pick the seat's model inside the loop | `plan/omniroute.md` |

## Calculated use

1. If the packet names a model id, call that id. Stop.
2. If the packet names `harness-winner`, call the combo. A model that fails three times gets two more tries. Five failures remove it from the combo. Survivors keep their local score as weight.
3. If the next step is a tool, do not batch and do not hand the turn to A2A.
4. If the other party is a den seat, send a mesh packet. Do not open an A2A task.
5. If the step is memory or an embedding, stay on this machine.

Agent skills that ship with OmniRoute, including routing and A2A notes, are references for the rows above. Loading one of them does not turn OmniRoute on.
