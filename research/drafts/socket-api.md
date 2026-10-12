# Socket API

Draft. The daemon socket owns the runtime. HTTP is the same handlers, not a second model.

## Read

`https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/REST` returned page-not-found. That guide is UNVERIFIED. Method and status claims below cite the pages that did open: [HTTP request methods](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Methods), [200 OK](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/200), [201 Created](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/201), [409 Conflict](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/409), and [Using server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events).

## What the socket is

`cmd/denctl/listen.go` `listenUnix` serves `http.Handler` on `unix` at `~/.local/share/harness/den.sock` (mode `0600`). `cmd/denctl/main.go` `serve` registers `/` (embedded `web/index.html`), `GET`-shaped `/api/state`, and `/api/events`. `cmd/denapi/main.go` refuses any host other than `127.0.0.1` or `::1`, default `127.0.0.1:8793`, and reverse-proxies `/` onto that socket. `/openapi.json` is the only denapi-local route (`cmd/denapi/openapi.json`).

## Resources

Names are the daemon runtime, not a parallel catalog.

| Resource | Meaning on the socket | Today |
| --- | --- | --- |
| session | One daemon process and the repo it reads (`State.Repo`) | folded into `/api/state` |
| team | The nest: `State.Projects` (org, name, path, role, status) | folded into `/api/state` |
| seat | One expansion row: `Slot` (`id`, `status`, `note`) | folded into `/api/state` |
| packet | One seven-field handoff the daemon would store | not served |
| grant | A named allow on a seat (no delete grant is already a rule, not an HTTP scheme) | not served |
| event stream | One-way state pushes | `GET /api/events` |

## Methods and status

MDN: GET retrieves and has no request content; PUT replaces the target and is idempotent; DELETE is idempotent; POST submits and may change state and is not idempotent. 200 means the request succeeded and, for GET, includes the resource. 201 means one or more resources were created. 409 means the request conflicts with the target's current state.

A client implements only these.

| Call | Result |
| --- | --- |
| `GET /api/state` | `200` `application/json`. Body is `console.State`: repo, services, counts, branches, slots, projects. |
| `GET /api/events` | `200` `text/event-stream`. See stream. |
| `GET /` | `200` embedded page. |
| `PUT /api/sessions/{id}` | Create or replace that session. Same body again is the same session. `200` if it already existed, `201` if this call created it. |
| `PUT /api/teams/{id}` | Same for a team. `200` or `201`. |
| `PUT /api/teams/{team}/seats/{id}` | Same for a seat. `200` or `201`. |
| `PUT /api/packets/{id}` | Same for a packet. Body is the seven fields. `200` or `201`. |
| `PUT /api/teams/{team}/seats/{id}/grants/{name}` | Same for a grant. `200` or `201`. |
| `GET` on any of those paths | `200` and that JSON object, once the daemon stores it. |
| `DELETE` on any of those paths | `204` when the id is gone, including when it was already gone. |

Creates are PUT to a client-chosen id. POST is not a create. A second identical PUT is a no-op success (`200`), not a second object. MDN's `Idempotency-Key` header is a different mechanism and is not used here.

Status a client must handle: `200`, `201`, `204`, `400` (fields are not the seven, or the id is empty), `404` (GET or a parent team/seat is missing), `405` (POST, PATCH, or a verb the path does not own), `409` (this body disagrees with the object that id already has), `500` (handler already returns this from `writeState` and from a missing `http.Flusher`).

## Error body

`application/json`. One object. `rule` is a short token. `id` is the target when there is one. No essay.

```json
{"rule":"packet-fields","id":"p1"}
```

Rule tokens this draft uses: `packet-fields`, `id-empty`, `not-found`, `method`, `conflict`, `stream-unsupported`. `http.Error` in `writeState` and the events handler is plain text today (`stream unsupported`, or `err.Error()`). The JSON body is the shape new handlers should write. It does not describe the current `500` text.

## Socket-only and HTTP-only

The event stream stays on the socket. `serve` writes `Content-Type: text/event-stream`, `Cache-Control: no-cache`, `Connection: keep-alive`, then `retry: 3000`, then `id`, `event: state`, and `data` ending in a blank line. A comment line (`: keepalive`) is the 15s idle tick. MDN: the response MIME type is `text/event-stream`; events are text separated by a pair of newlines; a line starting with `:` is a comment; `retry` is the reconnection time in milliseconds; the connection is one-way, client to server is not a stream. `EventSource` is a browser client. The daemon does not open one.

Streaming events are not reimplemented in denapi. The proxy dials `unix` and copies the response, so a browser on loopback reads the socket's stream.

What HTTP adds is TCP. `denapi` `ListenAndServe` on `127.0.0.1:8793` (or `::1`) because a browser cannot dial the Unix socket (`listen.go` says the same). The host check is `host != "127.0.0.1" && host != "::1"` then `log.Fatal`. `publicAddr.Error` in denctl is the string `denctl listens on 127.0.0.1 only`. denapi also serves `GET /openapi.json` from the embedded spec. That spec is not a socket resource.

No auth scheme, no version prefix, no GraphQL. Loopback plus socket mode `0600` is the current bound.

## Assessment

Already served on the socket, and therefore on loopback through the proxy: `GET /api/state` (session, team, and seat as fields of one JSON object), `GET /api/events` (the event stream), and `GET /` (the embedded page). denapi adds only the TCP listen and `GET /openapi.json`.

New, not in the mux: packet and grant as their own resources; PUT-by-id for session, team, seat, packet, and grant; DELETE; the `rule` JSON error body. Those wait until the daemon owns the bytes. The HTTP wrapper does not invent them ahead of the socket.

## Types

- `State`, `Slot`, `Project`, `Branch` in `internal/console/state.go`.
- Error object: `rule` string, optional `id` string.
- Ids: session, team, seat, packet, grant name. Client-chosen. PUT target.

## Functions

- `listenUnix`, `defaultSock` — socket listen.
- `serve`, `writeState`, `stateJSON` — page, state, event stream.
- `denapi.main` — loopback reverse proxy and OpenAPI file.
- Not implemented: PUT and DELETE handlers for the five resources.

## Invariants

- One resource model. If the socket does not store it, HTTP does not.
- Creates are idempotent PUT. A repeat with the same body is `200`.
- The event stream is SSE from denctl. denapi only forwards it.
- denapi binds `127.0.0.1` or `::1` only.
- An error body names a rule. It does not explain.

## File path

`research/drafts/socket-api.md`

## Adopt or refuse

- Section: socket API, this draft.
- Function: project the daemon socket. Do not add a second model.
- Read: `cmd/denctl/main.go`, `cmd/denctl/listen.go`, `cmd/denapi/main.go`, `cmd/denapi/openapi.json`.
- Do: serve state, events, and the page as they are. Add packet and grant only in denctl.
- Write: this file only.
- Next: a denctl PUT for packet, with the `rule` body, before any denapi route.
- Rail: empty.
