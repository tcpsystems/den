# Goals

A goal is done only when a test or a running command shows it. A plan sentence is not done.

## 0.1.0 — this tree

- One catalog of identities, roles, niches, profiles, skills, tools, MCP servers, and models. `internal/catalog`. `TestLeadBindsWithoutRequiredMCP`.
- One signed grant. `internal/grant`. `TestSignVerifyScopeAndExpiry`.
- One secret store that shows names and reveals a value only for a matching grant. `internal/vault`.
- One skill path per grant. `internal/inject`. `TestOneReturnsGrantedSkill`.
- Many MCP providers, none required. `mcp/catalog.json`. `denmcp` starts only a granted image.
- One wave, one branch per seat, local commits, no push. `internal/orch`.
- A loopback controller and a TUI. `denctl`, `dentui`. The web page reads state. It does not start a seat.
- `go test ./...` exits 0.

## Not in 0.1.0

- `denctl` does not call `inject.One`, `catalog.Lead`, or `grant.Verify`.
- No model call is made by the harness binaries.
- OpenViking is a separate process on `127.0.0.1:1933`. This repo does not start it.
- A trained skill model does not exist. The injector does not rank.
- `den` writes `.den/init.json` and starts `denctl` and `dentui` if those binaries are on `PATH`. It does not onboard a codebase or an MCP server.

## Next release, 0.2.0

1. `den init` loads the board, prints the lead profile, and stops for questions the board cannot answer.
2. After those answers, the same command binds the lead profile and writes the skill path `inject.One` returns.
3. A seat message stays in this process. No second protocol until that mailbox has a test.
4. OpenViking recall stays a client setting in the DeerFlow checkout. This repo records the URL and does not raise a cap it does not own.
