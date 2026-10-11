# Fetch size and speed

Draft. Researcher `r8` (omniroute-qwen) owns this file.

## Module

Efficiency section lead. Implements the "pointer only" rule for Rail and Steer in the next fetch, as decided by Superpowers and supported by Graphify, AgentScope, and Spring. The section record travels with its own data; external state (Steer, Ledger, Rail bodies) is excluded from the fetch payload, referenced by pointer if needed.

## Types

```go
// SectionRecord represents the immutable core of a task unit.
type SectionRecord struct {
    Section  string   // Stable id. One job.
    Function string   // One sentence describing the purpose.
    Read     []string // The only files. Read-only paths.
    Do       []string // Steps, in order.
    Write    string   // One mesh kb key for output.
    Next     string   // The next section id, or "stop".
    Rail     string   // Empty, or a switch only the human flips.
}

// FetchPayload is what the mesh sends to a model seat.
type FetchPayload struct {
    Record SectionRecord
    // Steer, Ledger, and Rail bodies are NOT included.
    // Only pointers or references if required by specific tool contracts.
}
```

## Functions

```go
// FetchSection retrieves the current section from the mesh kb.
// It enforces the invariant that no external state bodies are embedded.
func FetchSection(mesh KB, sectionID string) (FetchPayload, error)

// ValidateSection checks that all required fields are present and non-placeholder.
// Rail must be empty unless explicitly set by human approval.
func ValidateSection(rec SectionRecord) error

// StoreSection writes the completed record back to the mesh kb.
func StoreSection(mesh KB, rec SectionRecord) error
```

## Invariants

1. **No Vendor Essay**: Content must be software symbols (types, functions), not narrative.
2. **Pointer Only Rule**: The `FetchPayload` never contains Steer, Ledger, or Rail bodies. These remain outside the section packet.
3. **Rail Gate**: `Rail` is empty at the gate. It is only populated by a human switch, not by model inference.
4. **Immutable Core**: `Section`, `Function`, `Read`, `Do`, `Write`, `Next`, and `Rail` are the only fields. No eighth field is added.
5. **Refusal**: If a note suggests embedding model choice or training loops, it is refused.

## Files

- `sections/efficiency/LEAD.md` (this file)
- `plan.md` (reference only, do not edit)
- `research/cross-note.md` (reference for consensus on fetch structure)

## Adopt

- Adopts the Superpowers consensus: Steer, Ledger, and Rail bodies are out of the fetch payload.
- Adopts the seven-field section record structure from `plan.md`.
- Adopts the refusal to expand the section record beyond these seven fields.

## Refuse

- Refuses embedding model selection logic in the mesh.
- Refuses adding new fields to the section record (e.g., "Approval", "Context").
- Refuses vendor-specific implementation details (e.g., OpenViking embeddings, deerflow-memory).

</content>