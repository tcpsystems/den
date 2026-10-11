# Append-only store

This section is the store. One writer. Seven fields. Put appends. Query reads what was stored. A lead cut is AppendRefusal, not a new section and not stop. Fast and recent memory are local directories on this machine. The OmniRoute copy is backup. Backup is not the read path.

Implements the decided memory rules: one store, later section supersedes, nothing deleted; the append-only progress line; `Put` / `Query` as store and fetch; a lead cut keeps the current section and a refusal row. The archived mesh upsert (`ON CONFLICT DO UPDATE`) is not this store.

## Module

`Store` lives in this directory. Embeddings stay on local port 8091. They do not choose the next section. `room.say` is not the store.

Store gate: Section, Function, Read, Do, Write filled; Next not durable until a lead lets it pass; Rail may be empty. Superpowers' Rail wait is not this put.

Query returns the stored seven fields. Steer body and ledger body stay out.

## Types

```go
type SectionID string
type RailSwitch string
type TaskID string

type NextRef struct {
    Kind string // "id", "stop", or ""
    ID   SectionID
}

type SectionRecord struct {
    Section  SectionID
    Function string
    Read     []string
    Do       []string
    Write    string
    Next     NextRef
    Rail     RailSwitch
}

type ProgressRow struct {
    At       time.Time
    Task     TaskID
    State    string // stored | cut | commit-refuse | failed | stop
    Artifact string
    Next     SectionID
}

type RefusalRow struct {
    At      time.Time
    Task    TaskID
    State   string // cut | commit-refuse
    Refused NextRef
    Reason  string
}

type MemoryLayer struct {
    Tier, Scope, Key, Value, Source string
    Timestamp                       time.Time
}

type Store struct {
    root string // local data dir
    home string // this machine; never BackupRoot
}

const BackupRoot = "tcp:/opt/harness-memory"

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
    ErrNotFound      = errors.New("store: section not found")
    ErrBackupRead    = errors.New("store: backup is not the read path")
)
```

## Functions

```go
func Open(root string) (*Store, error)
// Refuses root under BackupRoot (ErrBackupRead).

func (s *Store) Put(rec SectionRecord) error
// Requires Section, Function, Read, Do, Write non-empty.
// Appends one JSONL row of the seven fields. Updates current.
// Refuses a second put of the same Section (ErrSecondStore).
// Refuses in-place rewrite of an existing Write key (ErrRewriteCorpus).
// Refuses a shorter body for that key (ErrShrink).
// Refuses Write == "" (ErrEmptyWrite). Does not read backup.

func (s *Store) Query(id SectionID) (SectionRecord, error)
// Latest stored record for id. Empty id is current.
// Returns the seven fields. No steer line. No ledger body.
// Does not open BackupRoot.

func (s *Store) AppendRefusal(row RefusalRow) error
// Appends kind cut | commit-refuse. Current stays.
// Does not Put a successor. State "stop" is ErrCutNotStop.

func (s *Store) AppendProgress(row ProgressRow) error
// One JSONL append: At, Task, State, Artifact, Next. No overwrite.

func (s *Store) ReadFast() ([]MemoryLayer, error)
// Reads ~/.local/share/harness/memory/fast on this machine.

func (s *Store) ReadRecent() ([]MemoryLayer, error)
// Reads ~/.local/share/harness/memory/recent on this machine.

func ReadBackup() ([]MemoryLayer, error)
// Always ErrBackupRead. Backup is not the read path.

func ChildWrite(parent, proposed SectionRecord) error
func EncodeBody(rec SectionRecord) ([]byte, error) // exactly seven fields
func DecodeBody(body []byte) (SectionRecord, error)
```

## Invariants

1. One store. One writer (mutex, append, fsync). No second kb. No checkpointer.
2. Section record is Section, Function, Read, Do, Write, Next, Rail. No eighth field.
3. Put requires Section, Function, Read, Do, Write. Rail may be empty. Next is not durable until a lead lets it pass.
4. Same Section id a second time is ErrSecondStore. In-place edit is ErrRewriteCorpus. Empty Write is ErrEmptyWrite. Shorter replacement body is ErrShrink. Delete is ErrDelete.
5. A later Section id supersedes current. The previous row stays. Nothing is deleted.
6. A lead cut of Next calls AppendRefusal, not Put. Current stays. Not a clean stop. Not a new section.
7. A child packet cannot write the parent's Write, Next, or Rail (ErrChildWrite).
8. Query returns the stored section, not steer, not the ledger body.
9. Fetch is after store. Query with no prior Put is ErrNotFound.
10. Fast and recent memory are local. BackupRoot is never opened to remember.

## Files

- `sections/store/store.go` — Put, Query, AppendRefusal, ReadFast, ReadRecent, ReadBackup
- `sections/store/store_test.go`
- `sections/store/LEAD.md` — this file

Data under the store root: `sections.jsonl`, `refusals.jsonl`, `progress.jsonl`, `current`. Fast memory: `~/.local/share/harness/memory/fast`. Recent: `~/.local/share/harness/memory/recent`. Backup: `tcp:/opt/harness-memory/` via `scripts/backup-fast-memory.sh`. A seat does not read that host.

## Adopt

Adopt the seven-field record. Write is the one key. Put appends. Query reads the append log. Progress is one JSONL append of timestamp, task id, state, artifact path, next id.

Adopt the store gate: Section, Function, Read, Do, Write filled; Next not durable until a lead; Rail empty allowed.

Adopt the lead-cut rule: mesh keeps the current section and a refusal row. Lead does not write the corpus.

Adopt local fast and recent as the recall path. Adopt backup as copy-only.

## Refuse

Refuse an eighth field. Refuse `room.say` as the store. Refuse a second store of the same Section, an empty Write, a shrink, a delete, and an in-place corpus rewrite.

Refuse treating a lead cut as NEXT stop or as a new section.

Refuse a child writing the parent's Write, Next, or Rail.

Refuse putting the steer line or the ledger body into `SectionRecord`.

Refuse the mesh choosing a model.

Refuse reading `tcp:/opt/harness-memory` to remember.

## Types

`SectionRecord`, `SectionID`, `NextRef`, `RailSwitch`, `ProgressRow`, `RefusalRow`, `MemoryLayer`, `Store`

## Functions

`Open`, `Put`, `Query`, `AppendRefusal`, `AppendProgress`, `ReadFast`, `ReadRecent`, `ReadBackup`, `ChildWrite`, `EncodeBody`, `DecodeBody`

## Invariants

One store, seven fields, append-only progress, later supersedes, no delete, cut is refusal not stop, fast memory local, backup is not the read path.

## File path

`sections/store/store.go`

## Adopt or refuse, against the seven fields already in plan.md

| Field | Store |
| --- | --- |
| Section | required on Put; key identity; second put refused |
| Function | required on Put |
| Read | required on Put; stored, not executed here |
| Do | required on Put; stored, not executed here |
| Write | required; the one key |
| Next | stored as id or stop; a cut is AppendRefusal, not this field becoming stop |
| Rail | stored with the section when set; may be empty; seat cannot write it |
