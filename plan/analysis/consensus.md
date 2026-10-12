# Consensus

One record per decision. A claim needs a file or a command that was run.

## 2026-10-12 — ship one self-use agent

Claim: the harness could not staff itself. `Builtin` named roles, tools, models, one niche, and one memory, and no identity, skill, or profile. `Bind` also did not check identity.

Evidence: `internal/catalog/catalog.go` `Builtin` and `Lead`. `go test ./...` exited 0, including `TestLeadBindsWithoutRequiredMCP`.

Accept.
