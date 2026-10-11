// Package steer attaches a human line to a live step.
// ApplySteer does not cancel. FinishInFlight lets the in-flight tool complete.
// SteerLine, SideRecord, and Section.Rail stay three objects.
package steer

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrNilStep       = errors.New("steer: live step is nil")
	ErrEmptyLine     = errors.New("steer: steer line text is empty")
	ErrTaskMismatch  = errors.New("steer: TaskID does not match the live step")
	ErrIdle          = errors.New("steer: no in-flight tool")
	ErrSeatAuthority = errors.New("steer: seat cannot write SideRecord, Rail, or CommitGrant")
	ErrNilSection    = errors.New("steer: section is nil")
	ErrNilHuman      = errors.New("steer: human is required")
	ErrDeleteGrant   = errors.New("steer: delete grant is refused for a seat")
)

// RailSwitch is empty, or a named switch only the human flips.
type RailSwitch string

// Section is the seven fields. No eighth. Rail is not a SteerLine.
type Section struct {
	Section  string
	Function string
	Read     []string
	Do       []string
	Write    string
	Next     string
	Rail     RailSwitch
}

// SteerLine is a live human utterance during a step. It is not a section field.
type SteerLine struct {
	TaskID string
	Text   string
	At     time.Time
}

// SideRecord is a human ruling. It is not Rail and not a SteerLine.
type SideRecord struct {
	TaskID string
	Kind   string // approval | reject | ruling | refuse
	Body   string
	At     time.Time
}

// LedgerPointer names a side entry that touches this task. Body is never included.
type LedgerPointer struct {
	EntryID string
}

// ToolCall is the in-flight tool on a live step. Nil InFlight means idle.
type ToolCall struct {
	Name   string
	Args   string
	Output string
	Err    error
}

// ToolResult is what FinishInFlight returns after the tool completes.
type ToolResult struct {
	Name   string
	Output string
	Err    string
}

// LiveStep is one running section step. Steer attaches here. It does not cancel.
type LiveStep struct {
	TaskID   string
	Section  Section
	InFlight *ToolCall
	Steer    *SteerLine
}

// NextFetch is what a later fetch may return for this task.
// Rail rides with the stored section. SteerLine and SideRecord bodies stay out.
type NextFetch struct {
	Section Section
	Ledger  *LedgerPointer
}

// CommitGrant opens widen / push / delete. Default all false until a human sets one.
type CommitGrant struct {
	WidenRead bool
	Push      bool
	Delete    bool
}

// Human marks a human-only call. A seat packet path must not construct one.
type Human struct {
	ID string
}

// ApplySteer attaches line to the live step.
// It does not cancel the run, does not clear InFlight, and does not open a new section.
func ApplySteer(step *LiveStep, line SteerLine) error {
	if step == nil {
		return ErrNilStep
	}
	text := strings.TrimSpace(line.Text)
	if text == "" {
		return ErrEmptyLine
	}
	if line.TaskID != "" && step.TaskID != "" && line.TaskID != step.TaskID {
		return ErrTaskMismatch
	}
	at := line.At
	if at.IsZero() {
		at = time.Now().UTC()
	}
	task := line.TaskID
	if task == "" {
		task = step.TaskID
	}
	// Copy. Leave InFlight and Section alone.
	step.Steer = &SteerLine{
		TaskID: task,
		Text:   text,
		At:     at,
	}
	return nil
}

// InjectSpectatorLine is the mid-run human line. Same path as ApplySteer.
// No halt. No new Section id.
func InjectSpectatorLine(step *LiveStep, line SteerLine) error {
	return ApplySteer(step, line)
}

// FinishInFlight completes the in-flight tool and clears InFlight.
// It does not clear Steer. It does not change Section. It is not a cancel.
func FinishInFlight(step *LiveStep) (*ToolResult, error) {
	if step == nil {
		return nil, ErrNilStep
	}
	call := step.InFlight
	if call == nil {
		return nil, ErrIdle
	}
	res := &ToolResult{
		Name:   call.Name,
		Output: call.Output,
	}
	var err error
	if call.Err != nil {
		res.Err = call.Err.Error()
		err = call.Err
	}
	step.InFlight = nil
	// Steer stays for later calls. Section is unchanged.
	return res, err
}

// WriteSideRecord stores a human ruling. A seat path returns ErrSeatAuthority.
func WriteSideRecord(human Human, rec SideRecord) error {
	if human.ID == "" {
		return ErrNilHuman
	}
	if strings.TrimSpace(rec.Kind) == "" || strings.TrimSpace(rec.Body) == "" {
		return ErrEmptyLine
	}
	return nil
}

// FlipRail sets the human switch on the section. Empty is legal. Seat cannot flip.
func FlipRail(human Human, sec *Section, sw RailSwitch) error {
	if human.ID == "" {
		return ErrNilHuman
	}
	if sec == nil {
		return ErrNilSection
	}
	sec.Rail = sw
	return nil
}

// FetchNext returns the stored section (Rail as stored) and an optional ledger pointer.
// It omits SteerLine and SideRecord bodies.
func FetchNext(sec Section, ledger *LedgerPointer) (NextFetch, error) {
	out := NextFetch{Section: sec}
	if ledger != nil && ledger.EntryID != "" {
		out.Ledger = &LedgerPointer{EntryID: ledger.EntryID}
	}
	return out, nil
}

// ApproveCommit opens widen / push / delete. Human only. Delete stays human-only.
func ApproveCommit(human Human, g CommitGrant) error {
	if human.ID == "" {
		return ErrNilHuman
	}
	_ = g
	return nil
}

// SeatApproveCommit is the seat path. Always refused.
func SeatApproveCommit(g CommitGrant) error {
	if g.Delete {
		return ErrDeleteGrant
	}
	return ErrSeatAuthority
}
