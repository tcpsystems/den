# Operator plane

Read from the working tree of `tcpsystems/den` on 2026-10-12. Sources: `cmd/denctl` (`main.go`, `listen.go`, `web/index.html`), `cmd/dentui/main.go`, `cmd/denapi` (`main.go`, `openapi.json`), `internal/console` (`state.go`). The TUI enable path calls `internal/initd` (`Load`, `Enable`, `Save`). That package is not an HTTP surface.

`cmd/denapi` is present on disk and untracked. `cmd/denctl/main.go` and `cmd/dentui/main.go` are dirty against HEAD `10bfe8a`. HEAD still has `denctl` listen on `127.0.0.1:8793` as well as the socket. The working tree does not. This map follows the working tree. It does not invent a route that is not registered in these files.

## Socket

`cmd/denctl` is the daemon. `main` takes `-sock` (default `defaultSock()`) and `-repo` (default `.`). `serve` builds one `http.ServeMux` and returns `listenUnix(sock, mux)`. There is no TCP listener in this tree.

`defaultSock` (`listen.go`) is `$HOME/.local/share/harness/den.sock`. If the home directory cannot be read it is `/tmp/den.sock`.

`listenUnix` creates the parent directory at mode `0700`, removes any existing socket file, `net.Listen("unix", path)`, `chmod 0600`, then `http.Server.Serve` with `ReadHeaderTimeout` 5s. The same mux is the only handler. The daemon log line is `denctl sock %s`.

`publicAddr.Error` still returns `denctl listens on 127.0.0.1 only`. Nothing in `serve` returns that error. Loopback enforcement moved to `denapi`.

## What the daemon serves

Registered on the socket mux only:

| Method | Path | Handler | Body |
| --- | --- | --- | --- |
| GET | `/api/state` | `writeState` | JSON from `stateJSON` → `console.Collect` |
| GET | `/api/events` | inline in `serve` | `text/event-stream` |
| GET | `/` | `http.FileServer` of embedded `web/` | `web/index.html` |

No other path is registered. There is no POST, PUT, PATCH, or DELETE. `writeState` answers 500 with `err.Error()` if `stateJSON` fails. `stateJSON` gives `Collect` an 8s timeout.

`/api/events` sets `Content-Type: text/event-stream`, `Cache-Control: no-cache`, `Connection: keep-alive`. It writes `retry: 3000`, then a frame only when the JSON changes: `id`, `event: state`, `data` = the same payload as `/api/state`. A 200ms ticker polls. A 15s ticker writes the comment `: keepalive`. The handler returns when the request context ends or `stateJSON` fails. A missing `http.Flusher` returns 500 `stream unsupported`.

`console.Collect` fills `State`: `Repo`, `Services`, `Counts`, `Branches` (`readBranches`), `Slots` (`Slots()`), `Projects` (`Projects`). Counts are `develop`, `drafts`, `archive`, `ahead`. `readBranches` runs `git for-each-ref refs/heads` and `git rev-list --count develop..ref`. Kind is `line` for `develop` and `main`, `archive` under `archive/`, otherwise `draft`. `probe` GETs four health URLs with a 2s timeout and records `up` or `down`: mesh `http://127.0.0.1:8787/api/health`, headscale `http://127.0.0.1:8080/health`, llama `http://127.0.0.1:8091/health`, omniroute `https://ai.tcp.systems/api/health`. `Projects` is a static nest list plus `present`/`missing` from the home directory. Slot notes name the API, the TUI, and SSE. `Slots` does not register a route.

## TUI

`cmd/dentui` dials the socket. It does not listen. `main` takes `-sock` (same default path, not via `defaultSock`) and `-repo`. `client` uses an `http.Transport` whose `DialContext` ignores the address and dials `unix`. `readEvents` loops `GET http://den/api/events`. The host name is only a label for the dialer.

The scanner keeps `event:` and `data:` until a blank line. A line starting with `:` is skipped. Only `event: state` with a JSON body that unmarshals into `console.State` becomes `stateMsg`. Any other event name is dropped. On scan error or a closed body it sends `errMsg` and sleeps 3s, then dials again.

`Update` paints on `stateMsg` (`status` `live`) and on `errMsg` (`status` `reconnecting`). Keys:

| Key | Function | Effect |
| --- | --- | --- |
| `q`, `ctrl+c` | `tea.Quit` | leave the program |
| `w` | `enable("web")` | local board only |
| `p` | `enable("pages")` | local board only |
| `f` | `enable("files")` | local board only |
| `s` | `enable("seats")` | local board only |

`enable` calls `initd.Enable` then `initd.Save`. `Enable` sets a part to `on` only when its note is `loopback page` or `this screen`. In `Default` that is `web` and `tui`. `pages`, `files`, and `seats` stay off and the notice is `stays off. this build has no process for it`. `Save` writes `repo/.den/init.json` at directory mode `0700`. That file is not an HTTP call. `Load` runs once in `newModel`. The screen lists parts, services, nest rows (`Org`/`Name`/`Status`/`Role`), and up to 20 branches that are not `archive` and not a zero-ahead `draft`.

The TUI never calls `/api/state`. It never writes to the socket. Turning `web` on does not start `denapi`.

## Web API

`cmd/denapi` `main` takes `-addr` default `127.0.0.1:8793` and `-sock` default `$HOME/.local/share/harness/den.sock` (no `/tmp/den.sock` fallback). `net.SplitHostPort` must yield `127.0.0.1` or `::1`. Any other host calls `log.Fatal("denapi listens on 127.0.0.1 only")`.

The mux has two registrations:

| Path | What it does |
| --- | --- |
| `/openapi.json` | writes the embedded `openapi.json` as `application/json` |
| `/` | `httputil.ReverseProxy` to the daemon socket |

The proxy `Director` sets `http://den`. `DialContext` dials `unix` at `-sock` and ignores the request host. There is no path allowlist, so a request that is not `/openapi.json` is forwarded as-is. The daemon then answers only the three routes above. A forwarded POST is not implemented on the daemon; Go's default mux rejects a method it did not register.

`openapi.json` documents GET `/api/state` (`getState`) and GET `/api/events` (`streamState`). It does not document `/openapi.json` or `/`.

## Web is observability only

Yes, for what these binaries implement. The page and the published API only read `console.Collect` and stream it.

`web/index.html` `paint` fills counts, services, slots, projects, and branches from one JSON object. `start` uses `EventSource("/api/events")` and listens for the event name `state`. It inserts branch name, subject, and SHA as text. There is no form, no button, and no `fetch` except the event stream. The page does not send a command.

`denapi` adds no control route. It serves the spec and proxies the daemon. The daemon's writes are git reads (`readBranches`, `run`) and outbound health GETs (`probe`). Nothing in this plane starts a seat, edits a branch, or changes a service.

The only operator write is local and off the socket: dentui `enable` → `initd.Enable` / `initd.Save`. `Enable` will mark `web` on. It does not listen, and it does not spawn `denapi`.
