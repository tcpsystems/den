# Ready

Check an item only when a file or a passing test shows it.

- [x] protocols — `plan/protocols.md`
- [x] agent definitions — `catalog.Lead` binds on `Builtin` (`TestLeadBindsWithoutRequiredMCP`)
- [x] signed grants — `internal/grant` `TestSignVerifyScopeAndExpiry`
- [x] MCP catalog with many providers and no required provider — `mcp/catalog.json` (7 rows, images unset); `TestSpawnManyProvidersInParallel` starts only granted ids
- [x] denorch seat branches — `internal/orch` `Open` adds one worktree and branch per seat
- [x] init command — `cmd/den` writes `.den/init.json`; `TestSaveAndEnable`
- [x] go test — `go test ./...` exited 0 on 2026-10-12
