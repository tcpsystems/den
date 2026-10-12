# Structural gaps: using the harness from inside itself

Read from `internal/*` and `cmd/*` at this worktree HEAD. A seat opened by `denorch` can edit its owned paths and commit. It cannot staff itself, load one skill body, sign its own lease, call a model, or reveal a secret through the runtime. `denctl` stays an observer.

What already holds, and is not a gap:

- `inject.One` returns one catalog path (`internal/inject/inject.go`).
- `grant.Sign` signs a contract; `grant.Verify` checks the pinned parent key (`internal/grant/grant.go`).
- `denorch open` adds one worktree per seat on `wave--seat` (`internal/orch/wave.go`, `cmd/denorch/main.go`).
- `denctl` listens on `~/.local/share/harness/den.sock` and on `127.0.0.1` only (`cmd/denctl/listen.go`, `cmd/denctl/main.go`).
- The web page reads `/api/state` and `/api/events`. It does not start a seat (`cmd/denctl/web/index.html`).

## 1. No self-staffing command

- Evidence: `internal/catalog/catalog.go` (`Lead`, `Bind`); `cmd/dencat/main.go`.
- Exists: `catalog.Lead` is the profile a session would bind (`den-1`, builder, `auto/coding`, skill `use-harness`). `catalog.Bind` checks that those ids are in the index. `dencat` prints the index as JSON.
- Missing: no `cmd/` binary calls `Lead` or `Bind`. Opening a wave does not attach the lead profile to a seat.
- Smallest next file: `cmd/dencat/main.go` — a `bind` subcommand that prints `catalog.Lead()` after `catalog.Bind`.

## 2. The one skill path is never loaded

- Evidence: `internal/inject/inject.go`; `cmd/` has no import of `inject`.
- Exists: `inject.One` returns one path, or `skills/use-harness/SKILL.md` with `false`. Builtin lists two skills (`use-harness`, `use-git`).
- Missing: nothing reads that path into a context packet. `One` does not call `grant.Verify` (it trusts a `parentOK` bool). `dencat` does not call `catalog.Load`, so `index.json` is unused. The other skill directories under `skills/` are not catalog rows.
- Smallest next file: `cmd/dencat/main.go` — a `skill` subcommand that calls `inject.One` and prints the path only.

## 3. A seat cannot obtain a signed contract

- Evidence: `internal/grant/grant.go`; `cmd/denmcp/main.go`.
- Exists: `grant.Sign` and `grant.Verify`. `denmcp spawn` reads a grant JSON the operator already wrote and a raw parent public key.
- Missing: no command generates a key, calls `Sign`, or writes a grant. A seat has no parent key and no way to ask for `skill:`, `mcp:`, or `secret:` scope.
- Smallest next file: `cmd/denmcp/main.go` — reject `spawn` unless the grant file was produced by `grant.Sign` under the pinned parent key. Do not put signing in `denctl`.

## 4. No model call inside the runtime

- Evidence: `internal/catalog/catalog.go` (model rows `auto/coding`, `auto/coding:cheap`); `internal/fetch/research.go`; `cmd/denfetch/main.go`.
- Exists: catalog names OmniRoute routes. `fetch.Search` and `fetch.Scrape` are HTTP clients for public search and page read. Keys come from the process environment.
- Missing: no chat-completions client, no session, no tool loop. A seat cannot ask the lead route to continue the harness. `denfetch` does not consult a grant or the vault.
- Smallest next file: `internal/fetch/research.go` — keep `Search` as the public search client. Do not add a model call to `cmd/denctl`.

## 5. Secrets are not held by the daemon

- Evidence: `internal/vault/vault.go`; `cmd/denfetch/main.go`; `cmd/denmcp/main.go`.
- Exists: `vault.Reveal` returns one value when `grant.Verify` allows `secret:<name>`. `vault.Map` is an in-memory map. `vault.Redact` strips values in text.
- Missing: no process loads a store. `Reveal` is called only from `internal/vault/vault_test.go`. `denfetch` reads `OMNIROUTE_API_KEY` and `FIRECRAWL_API_KEY` from the environment. `denmcp` does not accept secret values, and it also does not reveal them.
- Smallest next file: `internal/vault/vault.go` — a file-backed `Store` whose `Get` is only reached through `Reveal`. Wire it from `cmd/denfetch/main.go`, not from the web handler.

## 6. The controller socket does not run a seat

- Evidence: `cmd/denctl/main.go`; `cmd/denctl/listen.go`; `cmd/dentui/main.go`; `internal/console/state.go`.
- Exists: Unix socket and loopback TCP serve the same mux. Routes are `/`, `/api/state`, `/api/events`. `console.Collect` reads git refs, counts, and project rows. `dentui` dials the socket and renders that stream.
- Missing: no request path opens a wave, commits, injects a skill, verifies a grant, or spawns MCP. A seat cannot drive the harness by dialing `den.sock`.
- Smallest next file: leave `cmd/denctl/main.go` as observability. Seat actions stay in `cmd/denorch/main.go` and `cmd/denmcp/main.go`.

## 7. Catalog tools have no jail, and one listed tool is absent

- Evidence: `internal/catalog/catalog.go` tool rows; `internal/catalog/terms.go`; `cmd/` listing.
- Exists: terms say a tool is "one argv the runtime runs inside the path jail" (`denread`, `denfetch`, `git`). Builtin points `denread` at `cmd/denread`. `denfetch` and `denorch` exist. `internal/orch` runs `git` and the named check argv.
- Missing: `cmd/denread` is not in the tree. No package restricts argv to a seat's `Owns`. `orch.Commit` runs git in the seat worktree and refuses foreign porcelain paths; that is not a jail around `denfetch` or an arbitrary tool.
- Smallest next file: `cmd/denread/main.go` — print one file only when its path is under a seat `Owns` prefix. Do not invent a plugin for it.

## 8. MCP spawn does not speak MCP

- Evidence: `internal/mcp/spawn.go`; `mcp/catalog.json`; `cmd/denmcp/main.go`.
- Exists: `mcp.Spawn` starts one `docker run` per granted `mcp:<id>`, in parallel. `denmcp list` prints the JSON catalog. The grant check is real.
- Missing: every row in `mcp/catalog.json` has an empty image, so `startOne` returns "no image set". Builtin MCP ids `omniroute` and `tinyfish` are not in that file. Nothing dials a container, lists tools, or forwards a seat call. Spawn is not reachable from `denctl` or from `inject`.
- Smallest next file: `mcp/catalog.json` — set an image only on a row that already exists, or add the existing builtin ids `omniroute` and `tinyfish` with empty images left empty. Do not start them from the web.

## 9. The orchestrator does not apply the role it stored

- Evidence: `internal/orch/commit.go` (`CheckMessage`, `roleType`); `internal/orch/wave.go` (`Commit`); `cmd/denorch/main.go`.
- Exists: `CheckMessage(role, subject)` maps researcher and architect to `docs`, builder to `feat`, reviewer to `test`. `Open` stores the role on the seat. Seat branches and owned-path checks work. Merge does not push.
- Missing: `Commit` calls `CheckMessage("", message)`, so the role is ignored. `denorch open` accepts any role string. `Take` merges a seat that has never committed. No command looks up the seat profile before the commit.
- Smallest next file: `internal/orch/wave.go` — pass the seat's stored `Role` into `CheckMessage`.

## 10. Memory is a catalog row, not a store

- Evidence: `internal/catalog/terms.go`; `internal/catalog/catalog.go` (`KindMemory`, id `session-ledger`).
- Exists: the vocabulary says memory is "facts the seat may read" and "session notes on disk." The lead profile names `session-ledger`.
- Missing: no `internal/` package reads or writes a session ledger. A seat cannot leave a note the next seat loads. OpenViking is not a backend in this tree.
- Smallest next file: `internal/catalog/catalog.go` — stop at the index. A note writer is not in `internal/` yet, so the next edit is a path field on the existing `session-ledger` row, not a new package.

`denctl` does not gain routes for any of the above. The page stays a view of `console.State`.
