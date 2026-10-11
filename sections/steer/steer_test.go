package steer

import (
	"errors"
	"testing"
	"time"
)

func sampleStep(inFlight bool) *LiveStep {
	step := &LiveStep{
		TaskID: "task-1",
		Section: Section{
			Section:  "steer-1",
			Function: "attach a steer line",
			Read:     []string{"sections/steer/LEAD.md"},
			Do:       []string{"apply", "finish"},
			Write:    "kb/steer-1",
			Next:     "",
			Rail:     "",
		},
	}
	if inFlight {
		step.InFlight = &ToolCall{
			Name:   "shell",
			Args:   "echo hi",
			Output: "hi",
		}
	}
	return step
}

func TestApplySteerAttachesWithoutCancel(t *testing.T) {
	step := sampleStep(true)
	line := SteerLine{TaskID: "task-1", Text: "prefer the short path", At: time.Unix(1, 0).UTC()}
	if err := ApplySteer(step, line); err != nil {
		t.Fatal(err)
	}
	if step.Steer == nil || step.Steer.Text != "prefer the short path" {
		t.Fatalf("steer = %+v", step.Steer)
	}
	if step.InFlight == nil || step.InFlight.Name != "shell" {
		t.Fatalf("in-flight was cancelled: %+v", step.InFlight)
	}
	if step.Section.Section != "steer-1" || step.Section.Next != "" {
		t.Fatalf("section changed: %+v", step.Section)
	}
}

func TestApplySteerRefusesEmptyAndMismatch(t *testing.T) {
	step := sampleStep(false)
	if err := ApplySteer(nil, SteerLine{Text: "x"}); !errors.Is(err, ErrNilStep) {
		t.Fatalf("nil step: %v", err)
	}
	if err := ApplySteer(step, SteerLine{TaskID: "task-1", Text: "  "}); !errors.Is(err, ErrEmptyLine) {
		t.Fatalf("empty: %v", err)
	}
	if err := ApplySteer(step, SteerLine{TaskID: "other", Text: "go"}); !errors.Is(err, ErrTaskMismatch) {
		t.Fatalf("mismatch: %v", err)
	}
}

func TestInjectSpectatorLineSamePath(t *testing.T) {
	step := sampleStep(true)
	if err := InjectSpectatorLine(step, SteerLine{Text: "mid-run note"}); err != nil {
		t.Fatal(err)
	}
	if step.Steer == nil || step.InFlight == nil {
		t.Fatalf("inject halted or cleared: steer=%v inflight=%v", step.Steer, step.InFlight)
	}
}

func TestFinishInFlightCompletesWithoutClearingSteer(t *testing.T) {
	step := sampleStep(true)
	if err := ApplySteer(step, SteerLine{Text: "keep going"}); err != nil {
		t.Fatal(err)
	}
	res, err := FinishInFlight(step)
	if err != nil {
		t.Fatal(err)
	}
	if res == nil || res.Name != "shell" || res.Output != "hi" {
		t.Fatalf("result = %+v", res)
	}
	if step.InFlight != nil {
		t.Fatalf("still in flight: %+v", step.InFlight)
	}
	if step.Steer == nil || step.Steer.Text != "keep going" {
		t.Fatalf("steer cleared: %+v", step.Steer)
	}
	if step.Section.Section != "steer-1" {
		t.Fatalf("section changed: %+v", step.Section)
	}
}

func TestFinishInFlightIdleAndToolError(t *testing.T) {
	step := sampleStep(false)
	if _, err := FinishInFlight(step); !errors.Is(err, ErrIdle) {
		t.Fatalf("idle: %v", err)
	}
	boom := errors.New("tool failed")
	step.InFlight = &ToolCall{Name: "shell", Err: boom}
	res, err := FinishInFlight(step)
	if !errors.Is(err, boom) {
		t.Fatalf("tool err: %v", err)
	}
	if res == nil || res.Err != boom.Error() {
		t.Fatalf("result = %+v", res)
	}
	if step.InFlight != nil {
		t.Fatal("in-flight not cleared after error")
	}
}

func TestFinishThenSteerAppliesToLaterCalls(t *testing.T) {
	step := sampleStep(true)
	if err := ApplySteer(step, SteerLine{Text: "use /tmp"}); err != nil {
		t.Fatal(err)
	}
	if _, err := FinishInFlight(step); err != nil {
		t.Fatal(err)
	}
	// Later call sees the steer; ApplySteer did not cancel the earlier tool.
	if step.Steer == nil || step.Steer.Text != "use /tmp" {
		t.Fatalf("steer missing for later call: %+v", step.Steer)
	}
	if step.InFlight != nil {
		t.Fatal("prior tool still marked in flight")
	}
}

func TestHumanRailAndFetchOmitSteerBody(t *testing.T) {
	human := Human{ID: "h1"}
	sec := Section{Section: "steer-1", Rail: ""}
	if err := FlipRail(human, &sec, "approve-design"); err != nil {
		t.Fatal(err)
	}
	if sec.Rail != "approve-design" {
		t.Fatalf("rail = %q", sec.Rail)
	}
	if err := FlipRail(Human{}, &sec, ""); !errors.Is(err, ErrNilHuman) {
		t.Fatalf("seat flip: %v", err)
	}
	if err := WriteSideRecord(human, SideRecord{Kind: "ruling", Body: "ok"}); err != nil {
		t.Fatal(err)
	}
	if err := WriteSideRecord(Human{}, SideRecord{Kind: "ruling", Body: "ok"}); !errors.Is(err, ErrNilHuman) {
		t.Fatalf("seat side: %v", err)
	}
	got, err := FetchNext(sec, &LedgerPointer{EntryID: "e1"})
	if err != nil {
		t.Fatal(err)
	}
	if got.Section.Rail != "approve-design" {
		t.Fatalf("rail missing on fetch: %+v", got)
	}
	if got.Ledger == nil || got.Ledger.EntryID != "e1" {
		t.Fatalf("ledger: %+v", got.Ledger)
	}
	// NextFetch type has no SteerLine field; body stays out by construction.
}

func TestApproveCommitHumanOnly(t *testing.T) {
	if err := ApproveCommit(Human{ID: "h1"}, CommitGrant{Push: true}); err != nil {
		t.Fatal(err)
	}
	if err := SeatApproveCommit(CommitGrant{WidenRead: true}); !errors.Is(err, ErrSeatAuthority) {
		t.Fatalf("seat grant: %v", err)
	}
	if err := SeatApproveCommit(CommitGrant{Delete: true}); !errors.Is(err, ErrDeleteGrant) {
		t.Fatalf("delete: %v", err)
	}
}
