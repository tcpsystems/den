# Steer without killing the run

This section is live steer. A steer line enters the live step. The run is not cancelled. The in-flight tool finishes. Wider read root, push, and delete stay shut until a human opens one.

Implements decided plan §6 Rails and steer-docs. Grunt files `r9g1` through `r9g5` had empty claims. No grunt sentence is followed.

## Module

`ApplySteer` and `FinishInFlight` live in this directory. Not a second bus. Not a model chooser. SteerLine, SideRecord, and Section.Rail stay three objects.

## Types

Defined in `sections/steer/steer.go`.

`SteerLine` is `TaskID`, `Text`, and `At`. It is a live human utterance. It is not a section field.

`SideRecord` is `TaskID`, `Kind`, `Body`, and `At`. Kind is `approval`, `reject`, `ruling`, or `refuse`. It is not Rail.

`LedgerPointer` is `EntryID` only. Body is never included.

`RailSwitch` is `""` or a named switch only the human flips.

`Section` holds the seven fields: `Section`, `Function`, `Read`, `Do`, `Write`, `Next`, `Rail`. No eighth field.

`ToolCall` is the in-flight tool: `Name`, `Args`, `Output`, `Err`.

`ToolResult` is `Name`, `Output`, and `Err` after the tool completes.

`LiveStep` is `TaskID`, `Section`, `InFlight`, and `Steer`. `InFlight` is nil when idle.

`NextFetch` is the stored `Section` (Rail as stored) and an optional `LedgerPointer`. It has no SteerLine and no SideRecord body.

`CommitGrant` is `WidenRead`, `Push`, and `Delete`. Default all false until a human sets one.

`Human` marks a human-only call.

## Functions

`ApplySteer(step *LiveStep, line SteerLine) error` attaches the line to the live step. It does not cancel. It does not clear `InFlight`. It does not change `Section` and does not open a new section. Empty text is `ErrEmptyLine`. A TaskID mismatch is `ErrTaskMismatch`.

`FinishInFlight(step *LiveStep) (*ToolResult, error)` completes the in-flight tool and clears `InFlight`. It does not clear `Steer`. It does not change `Section`. It is not a cancel. Idle is `ErrIdle`. A tool error is returned after `InFlight` is cleared so a later call can see the steer.

`InjectSpectatorLine(step *LiveStep, line SteerLine) error` is the mid-run human line. Same path as `ApplySteer`. No halt. No new Section id.

`WriteSideRecord(human Human, rec SideRecord) error` is human only. A seat path returns `ErrSeatAuthority` / `ErrNilHuman`.

`FlipRail(human Human, sec *Section, sw RailSwitch) error` is human only. Empty is legal. Seat or prompt cannot flip.

`FetchNext(sec Section, ledger *LedgerPointer) (NextFetch, error)` returns the stored section (Rail included as stored) and an optional ledger pointer. Omits SteerLine and SideRecord bodies.

`ApproveCommit(human Human, g CommitGrant) error` opens widen / push / delete. Human only.

`SeatApproveCommit(g CommitGrant) error` always refuses. Delete is `ErrDeleteGrant`; any other seat grant is `ErrSeatAuthority`.

## Invariants

- SteerLine, SideRecord, and Section.Rail are three objects. No collapse.
- ApplySteer / InjectSpectatorLine never cancel the run and never store a successor Section.
- FinishInFlight runs to completion; steer does not abort the current tool.
- A seat packet cannot write SideRecord, cannot FlipRail, cannot ApproveCommit.
- FetchNext never returns SteerLine or SideRecord.Body; ledger is pointer-only.
- Rail rides with the stored section (empty or switch). Notes followed: graphify, alibaba-agents, alibaba-platform. Refuse topics putting the live steer line into NextFetch.
- No eighth section field.

## Files

- `sections/steer/steer.go` — ApplySteer, FinishInFlight, InjectSpectatorLine
- `sections/steer/steer_test.go`
- `sections/steer/LEAD.md` — this file

The symbols are not added to `agent-mesh`.

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

## Types

`SteerLine`, `SideRecord`, `LedgerPointer`, `RailSwitch`, `Section`, `ToolCall`, `ToolResult`, `LiveStep`, `NextFetch`, `CommitGrant`, `Human`

## Functions

`ApplySteer`, `FinishInFlight`, `InjectSpectatorLine`, `WriteSideRecord`, `FlipRail`, `FetchNext`, `ApproveCommit`, `SeatApproveCommit`

## Invariants

Steer enters the live step; in-flight finishes; not a cancel; three objects; seat cannot flip Rail or write SideRecord; NextFetch omits steer and side-record bodies.

## File path

`sections/steer/steer.go`

## Adopt or refuse, against the seven fields already in plan.md

| Field | Steer |
| --- | --- |
| Section | unchanged by ApplySteer and FinishInFlight; no successor opened |
| Function | unchanged |
| Read | unchanged |
| Do | unchanged |
| Write | unchanged; not a side-record body |
| Next | unchanged; steer is not Next |
| Rail | human FlipRail only; not a SteerLine; rides on NextFetch when stored |
