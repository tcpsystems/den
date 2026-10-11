# Packet schema and validator

Draft. Researcher `r1`. Not a decided line. Grunts `r1g1` through `r1g5` had empty claims. No grunt sentence is followed.

## Decided rule

Implements plan.md §3 and the consensus line: a section record is seven fields and no others, and the packet is handed to the seat. `chase-ledger` / `review-ledger`: `task.claim` hands `protocol.Task.Body` so the seat never opens a file. `review-rules` / `chase-rules`: the human side record is not a field the seat can write. `tool-docs`: a tool result attaches and does not set Next.

Does not implement either open. The fill-gate (which fields must be non-empty before a tool or a store) stays open. The next-fetch open (pointer only, or Rail riding with the section) stays open. Superpowers does not win the gate in this module. plan.md §3 does: Rail is empty, or a switch only the human flips.

## Types

```go
type SectionID string

type RailSwitch string // "" or one switch token. Not an object.

type SectionPacket struct {
    Section  SectionID  `json:"Section"`
    Function string     `json:"Function"`
    Read     []string   `json:"Read"`
    Do       []string   `json:"Do"`
    Write    string     `json:"Write"` // one kb key, or ""
    Next     string     `json:"Next"`  // "", "stop", or one SectionID
    Rail     RailSwitch `json:"Rail"`
}

type Actor string // "seat" | "human"

type ToolAttach struct {
    Name   string
    Result string
}

type Turn struct {
    Packet SectionPacket
    Tool   *ToolAttach // live only; not a section field
}

type Refusal struct {
    Code  string // "extra-field" | "missing-field" | "shape" | "seat-rail"
    Field string
}
```

## Functions

`Parse(body string) (SectionPacket, *Refusal)` decodes `protocol.Task.Body`. The JSON object must hold exactly `Section`, `Function`, `Read`, `Do`, `Write`, `Next`, `Rail`. Unknown keys return `extra-field`. Absent keys return `missing-field`. It does not call `kb.query`, `task.get`, or open `Read`.

`Validate(p SectionPacket) *Refusal` checks shape only. `Section` is one non-empty token with no whitespace. `Function`, if set, is one string with no newline; empty is legal. `Read` and `Do` are arrays of strings; empty is legal; null is `shape`. `Write` is one string or empty, never an array or a comma-joined list. `Next` is `""`, `"stop"`, or one `SectionID`, never an array. `Rail` is `""` or one token with no whitespace, never an object or array. Empty `Function`, `Read`, `Do`, `Write`, `Next`, and `Rail` pass. That pass does not tell a later gate to accept them.

`ValidateSeat(actor Actor, handed, proposed SectionPacket) *Refusal` runs `Validate(proposed)`. If `actor` is `seat` and `proposed.Rail != handed.Rail`, code `seat-rail`. A seat cannot set Rail and cannot clear it. Only `human` may change Rail, and the new value still has to pass `Validate`.

`Attach(p SectionPacket, tool *ToolAttach) Turn` puts the result on `Turn.Tool`. It does not write `Next`, `Write`, or `Rail`.

`HandOffBody(task protocol.Task) (SectionPacket, *Refusal)` is `Parse(task.Body)`. `Task.ID`, `Title`, `Dept`, `Role`, `State`, `Owns`, `DependsOn`, `Lease`, and `Evidence` stay on the mesh task. They are not copied into `SectionPacket`.

No `CutNext`. A lead cut is not a transform of this type and not a clean stop. This module does not call `Store.PutKb` or `Store.CreateTask`.

## Invariants

- Seven keys, no eighth. A skill body, catalog, model id, steer line, ledger body, approval object, checkpoint, or token counter as a key is `extra-field`.
- Bytes a later `kb.put` stores are `SectionPacket` only. `ToolAttach` is stripped before that write. When `Write` is set, `protocol.KbEntry.Key` is that key. `KbEntry.Tags` are not section fields.
- `Parse` and `Validate` do not read paths in `Read`. They do not fill a missing path from the repo.
- `protocol.TaskState` is not `Next`. It stays the mesh state: `pending`, `claimed`, `working`, `suspended`, `completed`, `failed`.
- Deer-flow is not imported. `ThreadStateResponse.next` (`[]string` of node names), `values`, `checkpoint`, `checkpoint_id`, `parent_checkpoint_id`, `tasks`, and `metadata` are a second store. `RunCreateRequest.input`, `checkpoint`, `interrupt_before`, `interrupt_after`, and `conversation_references` are not the seven fields. `RunResponse` counters and `thread_id` are not fields. `ThreadGoalRequest.objective` is not `Function`.
- `room.say` is not the store and not the packet.

## File path

`sections/packet/LEAD.md`. The symbols are drafts in this file. They are not added to `agent-mesh` and not added to `deer-flow`.

## Adopt or refuse

- Section. Adopt `SectionPacket.Section` as the stable id. Refuse `protocol.Task.ID` and deer-flow `thread_id` as that id.
- Function. Adopt one sentence string on the packet. Refuse deriving it from a paragraph, from `ThreadGoalRequest.objective`, or from `RunCreateRequest.input`.
- Read. Adopt `[]string`, read-only, possibly empty. Refuse opening those paths to discover the job. Refuse a repo walk and `conversation_references`.
- Do. Adopt ordered `[]string`, possibly empty. Refuse graph edges and generated code as Do. Refuse a tool-argument check here. That check is the open gate, not this schema.
- Write. Adopt one kb key or empty. Refuse a key list, a file-path set, a checkpoint id, and `room.say`.
- Next. Adopt `""`, `stop`, or one section id. Refuse `DependsOn`, `ThreadStateResponse.next`, and `TaskState`. Refuse deleting Next inside this type.
- Rail. Adopt empty or one human switch. Refuse a seat write that changes it. Refuse a side-record object, a steer line, or a ledger body stored in Rail. Refuse making the first tool wait on Rail. Superpowers loses that point here; plan.md §3 wins. Neither side of the next-fetch open is adopted.

MCP JSON-RPC and the A2A task object are refused as the packet. Those grunt notes were empty and those pages were not opened. A tool call stays one loop call. Its result is `ToolAttach`, not a field. Elicitation is not Rail.
