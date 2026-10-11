# Controller plan

The minimal controller is `cmd/denctl`. It is a Go binary and one static page. Listen address defaults to `127.0.0.1:8793` so it does not share a port with the Python sketch on `8792`. Do not grow the sketch. Expansion is a slot in `internal/console/state.go`, not a second framework.

## Want

The start before this page exists is `plan/start.md`. Onboarding reports which local services are present. It does not reveal secrets. The page, when built, shows the same present or missing state.

The person running the harness can see every section branch, every service, every guest, and every schedule, on one page, and can tell a seat to start or stop only by a button that asks first. The page is local. It is not the Foxcore panel and it is not a public site.

## Landscape

| Thing | Already true | Use it |
| --- | --- | --- |
| fox | SvelteKit, Cloudflare Workers, the company panel | Visual language only. Do not deploy the controller there |
| meshctl | Go, loopback hub | Same language for the controller API |
| den | Markdown plan, section branches, `tcpsystems/den` | The data the page reads |
| Sketch | `console/server.py` and `console/index.html` | Throw away after the real binary exists |
| Embedder | `127.0.0.1:8091` | A probe, not a dependency of the page |
| smolvm | Guests on the `harness` network | A probe |

## Language and framework

The API is Go. One static binary, no runtime package manager, the same toolchain as meshctl. It binds to `127.0.0.1` only.

The page is Svelte, compiled to static files and served by that binary. No Node process in production. Svelte is the fox panel's language, so the controller can share type and spacing habits without becoming the panel.

Do not use React, Next, or a dashboard kit. Do not use LangGraph, Deep Agents, or OmniRoute to render the page.

## Plugins and extensions

There is no plugin market. An extension is one Go file that implements a probe.

```
type Probe interface {
    Name() string
    Read(ctx context.Context) (Status, error)
}
```

Shipped probes: git branches, mesh health, embedder, OpenViking, OmniRoute health, smolvm guests, schedules. A new probe is a new file and a test. It is not a config flag that loads remote code.

Schedules are read from a local file `schedules.json` when one exists. An empty list is an honest empty state, not a fake calendar.

## Deployment

On this workstation, a user systemd unit runs the binary. The unit does not start at boot until the checks below pass. The listen address is `127.0.0.1:8792`. There is no Caddy site, no public DNS, and no copy on tcp.systems.

A fresh Debian host gets the same binary from `deploy/debian-fresh.sh` only after KVM and smolvm are in. The controller is not part of that first script.

## Usage

1. Open `http://127.0.0.1:8792/`.
2. Read the four counts: branches, section branches, drafts, branches ahead of `develop`.
3. Read the service pills. Down is red. Up is not a claim that the service is correct.
4. Read the branch table. Drafts that are still the base commit are hidden. Section branches that are ahead of `develop` stay visible.
5. Reload fetches `/api/state` again. The page also fetches every 15 seconds.
6. A later packet may add Start and Stop. Those buttons do not exist in the first build. They require a confirm step and a refusal if the target is not a den seat.

## Packet for the orchestrator

- Section: controller
- Function: serve a loopback page of den state
- Read: this file, `DEVELOP.md`, the git repo
- Do: one seat for the Go API, one seat for the Svelte page, one seat for probes, one seat for tests
- Write: `cmd/denctl` and `web/`
- Next: the check seat
- Rail: empty until a human asks for Start and Stop

## Seats

| Seat | Skill | Done when |
| --- | --- | --- |
| API | Go section | `GET /api/state` returns the counts and does not listen on any other address |
| Page | web section | The page has one job, readable type, a reload control, and no decorative framework |
| Probes | tools reference | Each probe has a timeout and a down state |
| Checks | optimise reference | Tests cover empty git, a down probe, and a branch with nothing ahead |

## Checks

1. The binary refuses to bind anything but `127.0.0.1`.
2. A down probe does not fail the whole page.
3. Branch names are text, not HTML. The page inserts them as text nodes.
4. No secret from the environment is included in `/api/state`.
5. The sketch on port 8792 is stopped before the binary is installed, so two servers do not share the port.

## Not in this build

Start and stop of seats. OmniRoute routing. A login form. A public URL. Charts of model cost. Editing `plan/PLAN.md` from the page.
