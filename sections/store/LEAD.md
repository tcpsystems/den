# Append-only store

Draft. Researcher `r3` (grok-4.6) owns this file.

Implements the decided memory rules: one mesh store, seven fields, later section supersedes, nothing deleted; the append-only progress line (timestamp, task id, state, artifact path, next id); `kb.put` / `kb.query` as store and fetch; `room.say` is not the store; a lead cut of Next keeps the current section and a refusal row.

Grunts `r3g1`–`r3g5` were empty at write. This draft still names the mesh symbols they were assigned: PutKb, QueryKb, refusal row, checkpointer-is-not-store, single writer, empty-write and shrink errors.

## Module

`SectionStore` sits on `store.Store.PutKb` and `store.Store.QueryKb`. It does not add a second database. Embeddings stay on local port 8091. The LangGraph checkpointer, zvec, OmniMemory, Qwen Memory, Mem0, and `deerflow-memory` are not this store.

Store gate (notes: graphify, alibaba-platform, alibaba-agents, six notes that leave Rail empty). Superpowers' Rail wait is the first implementation tool, not this put.

Next fetch of a stored section returns the stored seven fields, including Rail when it is set (graphify, AgentScope, Spring). Steer body and ledger body stay out. The ledger is a pointer on `ProgressRow` (superpowers pointer-only, applied to the side record only).

## Types

```go
// SectionRecord is the section. Seven fields, no eighth.
type SectionRecord struct {
    Section  SectionID   // stable id, one job
    Function string      // one sentence
    Read     []string    // the only files, read-only
    Do       []string    // steps, in order
    Write    string      // one mesh kb key; PutSection uses this as KbEntry.Key
    Next     NextRef     // next section id, or stop; not durable until a lead lets it pass
    Rail     RailSwitch  // empty, or a switch only the human flips
}

type SectionID string

// NextRef.Kind is "id", "stop", or "". Empty is unset. "stop" is a clean stop.
type NextRef struct {
    Kind string
    ID   SectionID
}

type RailSwitch string // "" or a named human switch. A seat packet cannot write it.

// ProgressRow is the decided ledger line. One append, never an edit.
type ProgressRow struct {
    At       time.Time
    Task     protocol.TaskID
    State    string     // stored | cut | commit-refuse | failed | stop
    Artifact string     // KbEntry.Key or path
    Next     SectionID  // empty on cut
}

// RefusalRow is a lead cut. Not a new section. Not NEXT stop.
type RefusalRow struct {
    At      time.Time
    Task    protocol.TaskID
    State   string // cut | commit-refuse
    Refused NextRef
    Reason  string
}

// MemoryLayer is the decided overlay. A later section supersedes. Nothing is deleted.
type MemoryLayer struct {
    Tier, Scope, Key, Value, Source string
    Timestamp                       time.Time
}

type SectionStore interface {
    PutSection(ctx context.Context, rec SectionRecord) error
    FetchLatest(ctx context.Context, id SectionID) (SectionRecord, error)
    AppendProgress(ctx context.Context, row ProgressRow) error
    AppendRefusal(ctx context.Context, row RefusalRow) error
}

var (
    ErrEmptyWrite    = errors.New("store: Write is empty")
    ErrEmptyBody     = errors.New("store: section body is empty")
    ErrMissingFields = errors.New("store: Section, Function, Read, or Do is empty")
    ErrShrink        = errors.New("store: refused to shrink stored section")
    ErrSecondStore   = errors.New("store: second store of current section")
    ErrRewriteCorpus = errors.New("store: in-place rewrite of stored section")
    ErrDelete        = errors.New("store: delete is refused")
    ErrCutNotStop    = errors.New("store: lead cut is not a clean stop")
    ErrChildWrite    = errors.New("store: child cannot write parent Write, Next, or Rail")
)
```

Wire to mesh types already on disk: `protocol.KbEntry`, `protocol.KbPutParams`, `protocol.KbQueryParams`, `protocol.KbQueryResult`, `protocol.Task`.

## Functions

```go
func (s *sectionStore) PutSection(ctx context.Context, rec SectionRecord) error
// Requires rec.Section, Function, Read, Do, Write non-empty.
// Next may be named; it is not durable until a lead lets it pass.
// Rail may be empty.
// Key = rec.Write. Body = JSON of the seven fields.
// Calls store.Store.PutKb. Refuses a second put of the same Section.
// Refuses a shorter body for an existing key (ErrShrink).
// Refuses Write == "" (ErrEmptyWrite) and empty JSON (ErrEmptyBody).
// Does not call store.Store.AppendMessage. Does not call OmniMemory.Remember.

func (s *sectionStore) FetchLatest(ctx context.Context, id SectionID) (SectionRecord, error)
// store.Store.QueryKb by key, latest KbEntry.At.
// Returns the stored seven fields. Rail rides if it was stored.
// Does not return a steer line. Does not return a ledger body.
// A pointer to a ProgressRow is allowed when that row's Task matches.

func (s *sectionStore) AppendProgress(ctx context.Context, row ProgressRow) error
// service.EventSink.Append kind "section.progress". JSONL, mutex, fsync.
// Fields: At, Task, State, Artifact, Next. No delete. No overwrite.

func (s *sectionStore) AppendRefusal(ctx context.Context, row RefusalRow) error
// service.EventSink.Append kind "section.cut".
// Stores the refused Next and Reason. Does not PutSection the successor.
// Current section stays. Returns ErrCutNotStop if State is "stop".

func EncodeBody(rec SectionRecord) ([]byte, error) // exactly the seven fields
func DecodeBody(e protocol.KbEntry) (SectionRecord, error)
```

Existing mesh calls this module uses, and does not replace:

- `(*store.Store).PutKb(ctx, protocol.KbEntry) error`
- `(*store.Store).QueryKb(ctx, protocol.KbQueryParams) ([]protocol.KbEntry, error)`
- `(*jsonlEventLog).Append(ctx, kind string, payload any) error`

Existing mesh calls this module does not use as the section store:

- `(*store.Store).AppendMessage` — `room.say`
- `(*store.Store).CreateTask` — `task.submit` hands a packet, it does not put a section
- `agent.OmniMemory.Remember` / `Recall`

## Invariants

1. One store. One writer (`BEGIN IMMEDIATE` on PutKb). No zvec, no second kb, no checkpointer.
2. Section record is Section, Function, Read, Do, Write, Next, Rail. No eighth field.
3. PutSection requires Section, Function, Read, Do, Write. Rail may be empty. Next is not durable until a lead lets it pass.
4. Same Section id a second time is ErrSecondStore. In-place edit of a stored body is ErrRewriteCorpus. Empty Write is ErrEmptyWrite. Shorter replacement body is ErrShrink. Delete is ErrDelete.
5. A later Section id supersedes current. The previous kb row stays. Nothing is deleted.
6. A lead cut of Next calls AppendRefusal, not PutSection. Current section stays. Not a clean stop. Not a new section.
7. A child packet cannot write the parent's Write, Next, or Rail (ErrChildWrite).
8. FetchLatest returns the stored section, not steer, not the ledger body. Progress is pointer-only.
9. Fetch is after store. A query with no prior put has nowhere to live.
10. Embeddings stay on 8091. They do not choose the next section.

## Files

`internal/store/section.go` — this module, on the mesh already at `~/Projects/agent-mesh`.

Wraps `internal/store/store.go` (`PutKb`, `QueryKb`) and `internal/service/eventlog.go` (`Append` into `events.jsonl`). Types `protocol.KbEntry` and `protocol.Task` live in `internal/protocol/protocol.go`. Schema: table `kb` in `internal/store/schema.sql`. Data: mesh `-data` dir, `mesh.db` plus `events.jsonl`.

## Adopt

Adopt the seven-field record. Write is the one `KbEntry.Key`. Put is `PutKb`. Fetch is `QueryKb`. Progress is one JSONL append of timestamp, task id, state, artifact path, next id.

Adopt the store gate from graphify, alibaba-platform, and alibaba-agents: Section, Function, Read, Do, Write filled; Next not durable until a lead; Rail empty allowed. Superpowers' Rail-before-first-tool is not this gate.

Adopt the lead-cut rule already repeated in the plan: mesh keeps the current section and a refusal row. Lead does not write the corpus.

Adopt the stored-section fetch from graphify, AgentScope, and Spring: Rail rides with the section when stored. Steer and ledger bodies stay out. Superpowers pointer-only applies to `ProgressRow`, not to dropping Rail from the record.

## Refuse

Refuse an eighth field. Refuse `room.say` / `AppendMessage` as the store. Refuse OmniMemory, Qwen Memory, Mem0, zvec, Graphify-as-memory, `deerflow-memory`, LangGraph checkpointer, rtp-llm prefix cache, and AirLLM as this store.

Refuse a second store of the same Section, an empty Write, a shrink, a delete, and an in-place corpus rewrite.

Refuse treating a lead cut as NEXT stop or as a new section.

Refuse a child or a lead writing the parent's Write, Next, or Rail.

Refuse putting the steer line or the ledger body into `SectionRecord`.

Refuse the mesh choosing a model.

## Types

`SectionRecord`, `SectionID`, `NextRef`, `RailSwitch`, `ProgressRow`, `RefusalRow`, `MemoryLayer`, `SectionStore`, `protocol.KbEntry`, `protocol.Task`

## Functions

`PutSection`, `FetchLatest`, `AppendProgress`, `AppendRefusal`, `EncodeBody`, `DecodeBody`, `store.Store.PutKb`, `store.Store.QueryKb`, `jsonlEventLog.Append`

## Invariants

One store, seven fields, append-only progress, later supersedes, no delete, cut is refusal not stop.

## File path

`internal/store/section.go`

## Adopt or refuse, against the seven fields already in plan.md

| Field | Store |
| --- | --- |
| Section | required on PutSection; key identity; second put refused |
| Function | required on PutSection |
| Read | required on PutSection; stored, not executed here |
| Do | required on PutSection; stored, not executed here |
| Write | required; the one `KbEntry.Key` |
| Next | stored as id or stop; not durable until a lead lets it pass; a cut is AppendRefusal, not this field becoming stop |
| Rail | stored with the section when set; may be empty; seat cannot write it |
