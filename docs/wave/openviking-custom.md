# OpenViking or a custom context database

## Question

Should den keep OpenViking (volcengine/OpenViking, local `127.0.0.1:1933`) or build a custom context database? Memory must be rebuildable from git. GitHub is the backup of global, thread, group, session, team, and selector memory (tool, niche, skill, model, section). The live read path stays on the machine. A frequently used skill is promoted to global memory as a principle, not by copying the skill body into the packet.

## Sources

- README, fetched 2026-10-12: https://github.com/volcengine/OpenViking and https://raw.githubusercontent.com/volcengine/OpenViking/main/README.md
- Path locks and crash recovery, fetched 2026-10-12: https://docs.openviking.ai/en/concepts/09-transaction
- Critical issue, open, fetched 2026-10-12: https://github.com/volcengine/OpenViking/issues/3508

## Finding

Keep OpenViking as the local long store. Do not build a second context database. Do not treat `viking://` as the backup.

Proven from the README: OpenViking is one virtual filesystem for resources, memories, and skills under `viking://`. Directory abstracts (L0) and overviews (L1) sit in front of full content (L2). A session commit archives the conversation and extracts memories as Markdown. The CLI can run against a local server. License is Apache-2.0.

Proven from the transaction page: filesystem content is the source of truth and VectorDB is a derived index. An index can generally be rebuilt from retained source data. Lost source data depends on backups. Path locks plus a persistent queue are not an atomic transaction across VikingFS, VectorDB, and the queue. The stated rule is to miss a search hit rather than return a bad one.

That matches den: live read stays on this machine at `127.0.0.1:1933`; GitHub holds the Markdown scopes. A custom database would reimplement directory scope, layered summaries, and a derived index that the docs already separate from source files.

## What den should not copy

- Do not put skill bodies in `viking://user/.../skills/` or in the packet. OpenViking stores the skill itself. den stores a principle in global memory and loads one skill body from the catalog.
- Do not adopt `session.commit` extraction as "memory happened." Issue #3508 (opened 2026-07-24, still open): with a llama.cpp OpenAI-compatible backend, a JSON list of memories is truncated to the first element, `memory_diff.json` shows `"adds": []`, and the task still reports `Processed: 1` with `Errors: 0`.
- Do not treat VectorDB, AGFS, or the queue as the git backup. Export Markdown for global, thread, group, session, team, and selector scopes. Rebuild the local index from those files.

## One plan line

OpenViking at `127.0.0.1:1933` is the live long read; git Markdown is the backup and the rebuild source; a used skill promotes a principle, not its body.

## Unproven

- No fetched page shows a command that rebuilds VectorDB from a git checkout of those six scopes.
- No fetched page shows OpenViking writing or reading a git remote. "Rebuildable from git" is a den export, not an OpenViking feature.
- Issue #3508 was not re-run here. Whether current `main` still truncates list-shaped memory JSON is unproven in this note.
- Whether a crash can leave VikingFS and VectorDB disagreeing after queue recovery is stated as a design limit, not measured on this host.
