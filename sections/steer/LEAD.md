# Steer without killing the run

Draft. Researcher `r9` (grok-4.5) owns this file.

Implements decided plan §6 Rails and steer-docs: a steer line enters the live step; the run is not cancelled; the in-flight action finishes. Wider read root, push, and delete stay shut until a human opens one.

## Module

Live steer and human rail objects on the mesh. Not a second bus. Not a model chooser.

## Types

```go
type SteerLine struct {
    TaskID    string
    Text      string
    At        time.Time
}

type SideRecord struct {
    TaskID string
    Kind   string // approval | reject | ruling | refuse
    Body   string
    At     time.Time
}

type LedgerPointer struct {
    EntryID string // touches this task only
}

type RailSwitch string // "" | human-set token; never seat-written

type Section struct {
    Section  string
    Function string
    Read     []string
    Do       []string
    Write    string
    Next     string
    Rail     RailSwitch
}

type LiveStep struct {
    TaskID   string
    Section  Section
    InFlight *ToolCall // nil when idle
    Steer    *SteerLine
}

type NextFetch struct {
    Section Section        // Rail rides when already stored (may be "")
    Ledger  *LedgerPointer // optional; body never included
    // no SteerLine, no SideRecord body
}

type CommitGrant struct {
    WidenRead bool
    Push      bool
    Delete    bool
}
```

## Functions

```go
func ApplySteer(step *LiveStep, line SteerLine) error
// Attach line to the live step. Does not cancel. Does not open a new section.

func FinishInFlight(step *LiveStep) (*ToolResult, error)
// In-flight tool completes before the steer text is applied to later calls.

func InjectSpectatorLine(step *LiveStep, line SteerLine) error
// Mid-run human line. Same path as ApplySteer. No halt. No new Section id.

func WriteSideRecord(human Human, rec SideRecord) error
// Human only. Seat packet path returns refuse.

func FlipRail(human Human, sec *Section, sw RailSwitch) error
// Human only. Empty is legal. Seat or prompt cannot flip.

func FetchNext(store Store, id string) (NextFetch, error)
// Returns stored Section (Rail included as stored). Optional LedgerPointer
// when a side entry touches this task. Omits SteerLine and SideRecord bodies.

func ApproveCommit(human Human, g CommitGrant) error
// Opens widen / push / delete. Default all false until human sets one.
```

## Invariants

- SteerLine, SideRecord, and Section.Rail are three objects. No collapse.
- ApplySteer / InjectSpectatorLine never cancel the run and never store a successor Section.
- FinishInFlight runs to completion; steer does not abort the current tool.
- A seat packet cannot write SideRecord, cannot FlipRail, cannot ApproveCommit.
- FetchNext never returns SteerLine or SideRecord.Body; ledger is pointer-only.
- Rail rides with the stored section (empty or switch). Notes followed: graphify, alibaba-agents, alibaba-platform. Refuse topics putting the live steer line into NextFetch.
- No eighth section field. Marked draft until a second finished note states the same and no lead cuts it.

## Files

`mesh/steer.go`

## Adopt

- Decided §6 / steer-docs: steer enters live step; in-flight finishes; run continues.
- Three-object split (verified superpowers line on steer vs ledger vs Rail).
- NextFetch: Rail switch rides with the section; steer and side-record bodies stay out; ledger pointer only when it touches this task.
- Spectator inject is ApplySteer, not a new section and not a halt.
- CommitGrant stays human-only for widen, push, delete.

## Refuse

- Cancel or interrupt as the steer path (UserInterruptEvent is not ApplySteer).
- Seat or client-supplied SideRecord / approval record treated as Rail.
- Pasting SteerLine, SideRecord body, or finish menu into a child or NextFetch.
- Topics rule that NextFetch carries the live steer line.
- Growing Section past the seven fields.
- Any seat delete grant.
