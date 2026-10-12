# OmniRoute ops

## Question

How should den call OmniRoute (`https://ai.tcp.systems/v1`) without making it the default path?

## Sources

- `plan/omniroute.md` — provisions, and the calculated-use rules.
- `skills/use-harness/SKILL.md` — the model call is a direct model id; OmniRoute is a tool you name.
- `cmd/denfetch/main.go` and `internal/fetch/research.go` — public-page search.
- `internal/fetch/page.go` — a page read is Firecrawl, not OmniRoute.
- `skills/use-harness/references/tools.md` — call OmniRoute only when `plan/omniroute.md` says yes.
- `skills/use-harness/references/memory.md` — fast and recent stay local.
- `plan/standup.md` — health check `curl -fsS https://ai.tcp.systems/api/health`.

## Finding

Settled: OmniRoute runs only when the packet names it. It is not the seat's model hop. If the packet does not name a provision, the provision stays off.

Search for public pages is `POST https://ai.tcp.systems/v1/search` via `cmd/denfetch`. The body is JSON `query` and `max_results`. The bearer is `OMNIROUTE_API_KEY` from the environment. It is not printed.

`denfetch` takes one argument. A URL is a page read. Any other argument is a search. `-read N` reads that many public pages after the search (max 4). With no `-read`, it reads 3. A private host is dropped before a page is opened.

```
denfetch https://example.com/page
denfetch -read 3 'agent harness'
```

Calculated use, in order:

1. If the packet names a model id, call that id. Stop.
2. If the packet names `harness-winner`, call the combo. A model that fails three times gets two more tries. Five failures remove it from the combo. Survivors keep their local score as weight.
3. If the next step is a tool, do not batch and do not hand the turn to A2A.
4. If the other party is a den seat, send a mesh packet. Do not open an A2A task.
5. If the step is memory or an embedding, stay on this machine.

A2A is only the remote card at `https://ai.tcp.systems/api/a2a`. Loading an OmniRoute agent skill does not turn OmniRoute on.

## What den should not copy

- Do not set `https://ai.tcp.systems/v1` as the default model hop.
- Do not use a combo when the packet names one model id.
- Do not batch while the seat is in a tool loop.
- Do not open A2A toward a den seat. Seats talk on the mesh.
- Do not use MCP when a local CLI already does the step. Gateway MCP only if `meshctl suite check` shows it.
- Do not rewrite or compress every call. Only after the leader already cut the prompt and the packet still asks for a smaller body.
- Do not embed through OmniRoute. Vectors are `127.0.0.1:8091`.
- Do not read memory through OmniRoute. Fast and recent are local. OpenViking is `127.0.0.1:1933`.
- Do not score models inside the loop. Scores come from local probes in `scripts/score-harness-winner.py`.
- Do not copy persistence before a local note exists. `scripts/backup-fast-memory.sh` runs after the note is written.
- Do not put a secret in the packet. Do not let a tool pick the next model.

## One plan line

OmniRoute is a tool the packet names; otherwise it stays off, and public search is only `POST https://ai.tcp.systems/v1/search` through `denfetch`.

## Unproven

- This note did not call `https://ai.tcp.systems/v1/search` and did not send a key.
- The live search JSON shape is unread here. The client accepts `results`, `data`, or `organic`.
- Combo retry weights and the health endpoint were not exercised in this pass.
