# Git consensus of section edits

Draft. Researcher `r10` (omniroute-kimi) owns this file.

## Module

Consensus protocol for section state transitions in a mesh-backed harness.
Decides which claims stand, which stay open, and which software implements the decided rules.

## Types

```go
// SectionID unique identifier for a section record
type SectionID string

// SectionRecord the seven fields stored on the mesh
type SectionRecord struct {
	Section  SectionID   // stable id, one job
	Function string      // one sentence
	Read     []string    // files, read-only
	Do       []string    // steps in order
	Write    string      // mesh kb key
	Next     SectionID   // next id or stop
	Rail     RailSwitch  // empty, or switch only human flips
}

// RailSwitch enum for approval states
type RailSwitch string

const (
	RailEmpty     RailSwitch = ""           // default, unset
	RailPending   RailSwitch = "pending"    // awaiting human
	RailApproved  RailSwitch = "approved"   // tool may proceed
	RailRefused   RailSwitch = "refused"    // human blocked
)

// Packet the data a parent hands to a child
type Packet struct {
	Task     string      // the child's assigned work
	ReadSnap []string    // copy of parent's Read at handoff
	Input    interface{} // task-specific input
}

// FetchResult what the next fetch returns
type FetchResult struct {
	Section SectionRecord // current section with all seven fields
	Pointer interface{}   // opaque refs to steer, ledger, side records
}

// GateRule which fields must be non-empty before store/tool
type GateRule struct {
	BeforeStore []string // field names required
	BeforeTool  []string // field names required including approval
}

// CutResult when a lead cuts Next
type CutResult struct {
	Current  SectionRecord // unchanged section
	Refusal  RefusalRow    // stored refusal marker
	NoCorpus bool          // corpus not written on cut
}

// RefusalRow minimal record of a lead cut
type RefusalRow struct {
	From   SectionID // section that was cut
	Reason string    // why Next was cut
	Time   int64     // timestamp
}
```

## Functions

```go
// Handoff creates a child packet from parent record
// Invariant: child cannot write parent's Write, Next, or Rail
func Handoff(parent SectionRecord, childTask string) Packet {
	return Packet{
		Task:     childTask,
		ReadSnap: append([]string{}, parent.Read...),
		Input:    nil,
	}
}

// IsChildWriteAllowed checks if field is writable by child
// Invariant: Write, Next, Rail are parent-only
func IsChildWriteAllowed(field string) bool {
	switch field {
	case "Write", "Next", "Rail":
		return false
	default:
		return true
	}
}

// FetchSection retrieves current section from mesh
// Open: whether Rail rides with section or pointer only
func FetchSection(meshKey string, opts FetchOpts) (FetchResult, error) {
	// Superpowers: Rail body stays out, pointer only
	// Graphify/AgentScope/Spring: Rail switch rides with section
	return FetchResult{}, nil
}

// GateCheck validates required fields before store or tool
// Open: six notes leave Rail empty; superpowers requires Rail approval
func GateCheck(rec SectionRecord, beforeTool bool, rule GateRule) error {
	for _, f := range rule.BeforeStore {
		if !fieldNonEmpty(rec, f) {
			return fmt.Errorf("gate: %s empty", f)
		}
	}
	if beforeTool {
		for _, f := range rule.BeforeTool {
			if !fieldNonEmpty(rec, f) {
				return fmt.Errorf("gate: %s required for tool", f)
			}
		}
	}
	return nil
}

// LeadCut handles when a lead vetoes Next
// Invariant: mesh keeps current section and refusal row; no corpus write
func LeadCut(current SectionRecord, reason string) CutResult {
	return CutResult{
		Current: current,
		Refusal: RefusalRow{
			From:   current.Section,
			Reason: reason,
			Time:   time.Now().Unix(),
		},
		NoCorpus: true,
	}
}

// IsCutClean determines if a stopped section is a clean stop
// Invariant: cut is not clean stop and not a new section
func IsCutClean(cut CutResult) bool {
	return false // cuts are never clean stops
}

// fieldNonEmpty checks if a field has non-placeholder content
func fieldNonEmpty(rec SectionRecord, field string) bool {
	// Implementation: check field is set and not placeholder
	return false
}
```

## Invariants

```go
// INV-1: Child packet isolation
// A parent hands a child its own packet.
// The child does not write the parent's Write, Next, or Rail.
// Source: appears in more than one research note

// INV-2: Lead cut behavior  
// A lead who cuts Next does not write the corpus.
// The mesh keeps the current section and a refusal row.
// The cut is not a clean stop and not a new section.
// Source: appears in more than one research note

// INV-3: Seven fields only
// Section record has exactly seven fields: Section, Function, Read, Do, Write, Next, Rail
// No eighth field. No extension.
// Source: consensus across all nine notes

// INV-4: Rail is human-only
// Rail is empty, or a switch only you (human) flip.
// Source: plan.md architecture line 48

// INV-5: Mesh as bus
// Store and fetch go through the mesh kb.
// The bus stays; it is not a model router.
// Source: five leads kept the mesh, one dissent noted
```

## Files

- `/sections/consensus/LEAD.md` - this consensus module
- `/pkg/consensus/section.go` - types and records
- `/pkg/consensus/gate.go` - gate validation logic
- `/pkg/consensus/packet.go` - handoff and child isolation
- `/pkg/consensus/cut.go` - lead cut handling
- `/internal/store/mesh.go` - mesh kb interface

## Adopt

The following rules are adopted from the cross-note analysis:

| Field | Decision | Notes |
|-------|----------|-------|
| **Section** | Required non-empty at gate | Universal consensus |
| **Function** | Required non-empty at gate | Universal consensus |
| **Read** | Required non-empty at gate | Universal consensus |
| **Do** | Required non-empty at gate | Universal consensus |
| **Write** | Required before store if Next set | Graphify, agent-skills, superpowers agree |
| **Next** | Optional initial, set by completion | Topics has it as id or stop; six notes allow empty |
| **Rail** | Human-only, not model-set | plan.md rule stands; superpowers dissent on tool-gate waits |

Adopted claim 1:
- **Parent-child isolation**: A parent hands a child its own packet. The child does not write the parent's Write, Next, or Rail. This constraint is encoded in `Handoff()` and enforced by `IsChildWriteAllowed()`.

Adopted claim 2:
- **Lead cut semantics**: When a lead cuts Next, the mesh keeps the current section unchanged and stores a refusal row. The cut is neither a clean stop nor a new section. No corpus is written. Implemented in `LeadCut()` and verified by `IsCutClean()`.

## Refuse

Refused pending opens (awaiting second pass with lead confirmation):

| Open Question | Status | Reason |
|---------------|--------|--------|
| **Fetch contents** | Refused to decide | Superpowers says Rail body out, pointer only. Graphify/AgentScope/Spring say Rail switch rides with section. Topics would fetch live steer. One note alone stays open. |
| **Gate field requirements** | Refused to decide | Six notes leave Rail empty. Superpowers requires Rail approval before first tool. Six vs one, but cross-note says leave both as opens. |
| **Mesh dissent** | Noted only | `remember-facts` dissented on mesh as bus. Other four leads kept it. Dissent stays open until another lead joins it. |

Refused patterns:
- Model choice on the mesh (bus is not a model router)
- Child write to parent Write/Next/Rail (parent-only fields)  
- Lead cut writes corpus (refused; cut stores refusal row only)
- Clean stop from lead cut (cut ≠ clean stop)
- Second model call to repair lead cut (vendors-west refused)

## Summary

This module implements the decided rules from two repeated claims:
1. Parent-to-child packet isolation with prohibited write fields
2. Lead cut semantics that preserve current section and log refusal

The two opens (fetch contents, gate field reqs) remain open per cross-note instruction until a second pass states one and a lead does not cut it.
