# Packet schema and validator

Implements plan.md §3 and the consensus line: a section record is seven fields and no others, and the packet is handed to the seat. `chase-ledger` / `review-ledger`: `task.claim` hands the task body so the seat never opens a file. `review-rules` / `chase-rules`: the human side record is not a field the seat can write. `tool-docs`: a tool result attaches and does not set Next.

Does not implement either open. The fill-gate (which fields must be non-empty before a tool or a store) stays open. The next-fetch open (pointer only, or Rail riding with the section) stays open. Superpowers does not win the gate in this module. plan.md §3 does: Rail is empty, or a switch only the human flips.

Grunt files `r1g1` through `r1g5` had empty claims. No grunt sentence is followed.

## Types

Defined in `sections/packet/packet.go`.

`SectionID` is one stable id: a non-empty token with no whitespace.

`RailSwitch` is `""` or one switch token. It is not an object.

`SectionPacket` holds `Section`, `Function`, `Read`, `Do`, `Write`, `Next`, and `Rail`. JSON keys are those names. `Write` is one kb key or empty. `Next` is `""`, `stop`, or one section id.

`Actor` is `seat` or `human`.

`ToolAttach` is `Name` and `Result`. It is live on a turn and is not a section field.

`Turn` is `Packet` plus `Tool`. Marshaling a turn writes the packet only.

`Refusal` is `Code` and `Field`. `Code` is `extra-field`, `missing-field`, `shape`, or `seat-rail`. A nil `*Refusal` means the check passed.

`Task` is the mesh handoff record. `HandOffBody` reads `Body` only. `ID`, `Title`, `Dept`, `Role`, `State`, `Owns`, `DependsOn`, `Lease`, and `Evidence` stay on the task.

## Functions

`Parse(body string) (SectionPacket, *Refusal)` decodes the task body. The JSON object must hold exactly `Section`, `Function`, `Read`, `Do`, `Write`, `Next`, and `Rail`. An unknown key returns `extra-field`, including when a required key is also absent. An absent key, with no unknown key, returns `missing-field` for the first missing name in that order. A duplicate key, a non-object, a null, an array, or any other wrong JSON type returns `shape`. `Parse` does not call `kb.query` or `task.get`, and it does not open `Read`.

`Validate(p SectionPacket) *Refusal` checks shape only. `Section` is one non-empty token with no whitespace. `Function`, if set, is one string with no newline; empty is legal. `Read` and `Do` are arrays of strings; empty is legal; a nil slice is `shape`. `Write` is one string or empty; a comma-joined list is `shape`. `Next` is `""`, `"stop"`, or one section id. `Rail` is `""` or one token with no whitespace. Empty `Function`, `Read`, `Do`, `Write`, `Next`, and `Rail` pass. That pass does not tell a later gate to accept them.

`ValidateSeat(actor Actor, handed, proposed SectionPacket) *Refusal` runs `Validate(proposed)`. If that passes and `actor` is `seat` and `proposed.Rail != handed.Rail`, the code is `seat-rail`. A seat cannot set Rail and cannot clear it. Only `human` may change Rail, and the new value still has to pass `Validate`. Any other actor is `shape` on `Actor`.

`Attach(p SectionPacket, tool *ToolAttach) Turn` puts the result on `Turn.Tool`. It does not write `Next`, `Write`, or `Rail`.

`HandOffBody(task Task) (SectionPacket, *Refusal)` is `Parse(task.Body)`.

`SectionPacket.MarshalJSON` writes the seven fields after `Validate`. `Turn.MarshalJSON` writes that packet and drops `ToolAttach`. `SectionPacket.UnmarshalJSON` is `Parse`.

No `CutNext`. A lead cut is not a transform of this type and not a clean stop. This package does not call `Store.PutKb` or `Store.CreateTask`. It does not import `agent-mesh` or deer-flow.

## Invariants

- Seven keys, no eighth. A skill body, catalog, model id, steer line, ledger body, approval object, checkpoint, or token counter as a key is `extra-field`.
- Bytes a later `kb.put` stores are `SectionPacket` only. `ToolAttach` is stripped before that write. When `Write` is set, that string is the kb key. Tags are not section fields.
- `Parse` and `Validate` do not read paths in `Read`. They do not fill a missing path from the repo.
- Mesh task state is not `Next`. `pending`, `claimed`, `working`, `suspended`, `completed`, and `failed` stay on `Task.State`. A key named `TaskState` or `State` is `extra-field`.
- Deer-flow is not imported. `values`, `checkpoint`, `checkpoint_id`, `parent_checkpoint_id`, `tasks`, `metadata`, `input`, `interrupt_before`, `interrupt_after`, `conversation_references`, `thread_id`, and `objective` are `extra-field`, not the seven fields.
- `room.say` is not the store and not the packet. As a key it is `extra-field`.

## File path

`sections/packet/packet.go`

`sections/packet/packet_test.go`

The symbols are not added to `agent-mesh` and not added to `deer-flow`.

## Adopt or refuse

- Section. Adopt `SectionPacket.Section` as the stable id. Refuse `Task.ID` and deer-flow `thread_id` as that id.
- Function. Adopt one string on the packet, with no newline. Refuse deriving it from a paragraph, from an objective field, or from a run input field.
- Read. Adopt `[]string`, read-only, possibly empty. Refuse opening those paths to discover the job. Refuse a repo walk and `conversation_references`.
- Do. Adopt ordered `[]string`, possibly empty. Refuse graph edges and generated code as Do. Refuse a tool-argument check here. That check is the open gate, not this schema.
- Write. Adopt one kb key or empty. Refuse a key list, a comma-joined list, a file-path set, a checkpoint id, and `room.say`.
- Next. Adopt `""`, `stop`, or one section id. Refuse `DependsOn`, a list of node names, and mesh task state. Refuse deleting Next inside this type.
- Rail. Adopt empty or one human switch. Refuse a seat write that changes it. Refuse a side-record object, a steer line, or a ledger body stored in Rail. Refuse making the first tool wait on Rail. Superpowers loses that point here; plan.md §3 wins. Neither side of the next-fetch open is adopted.

MCP JSON-RPC and the A2A task object are refused as the packet. Those grunt notes were empty and those pages were not opened. A tool call stays one loop call. Its result is `ToolAttach`, not a field. Elicitation is not Rail.
