# Tooling TUI

## Question

What is the smallest next step for denctl, denapi, and dentui so one packet can be submitted and one result read?

## Sources

- `cmd/denctl/main.go` — flags `-addr` (default `127.0.0.1:8793`), `-sock`, `-repo`. Mux serves `/`, `GET /api/state`, and `GET /api/events`. Any host other than `127.0.0.1` or `::1` is refused.
- `cmd/denctl/listen.go` — Unix socket `~/.local/share/harness/den.sock`, mode `0600`, same handler as the loopback page.
- `cmd/denctl/web/index.html` — static page. No packet form and no submit control.
- `cmd/dentui/main.go` — flag `-sock` (same default path). Dials that socket, `GET http://den/api/events`, paints a frame on `event: state`. Keys handled: `q`, `ctrl+c`.
- `cmd/denfetch/main.go` — `denfetch <https-url>` reads one public page through Firecrawl. `denfetch -read N <query>` searches OmniRoute, then reads pages (`-read` defaults to 3, help text says max 4). `OMNIROUTE_API_KEY` and `FIRECRAWL_API_KEY` come from the environment and are not printed.
- `internal/fetch/research.go` — search URL is `https://ai.tcp.systems/v1/search`.
- `internal/console/state.go` — `State` is repo, services, counts, branches, slots, projects. Slot `go` is "API binary". No packet field.
- `scripts/stand-memory.sh` — starts `llama-embed.service`, `docker compose -p openviking` from the existing compose file, and `mcp-hub.service` when that user unit exists. Checks `http://127.0.0.1:8091/health` and `http://127.0.0.1:1933/health`. Does not print keys. Refuses the embedder while the display pin is still the RX 6600.
- `plan.md` section 3 — a packet is Section, Function, Read, Do, Write, Next, Rail. No other fields.
- No `cmd/denapi` directory.

## Finding

denapi is not a separate command here. The API is the two GET handlers on `cmd/denctl`. dentui only reads that event stream. Nothing in those files accepts a packet body or returns a packet result. denfetch already searches OmniRoute and reads one public page; that result is page markdown, not a stored packet. `scripts/stand-memory.sh` only starts llama-embed, OpenViking, and mcp-hub. Channels and a webapp stay after one local round-trip.

## What den should not copy

Do not copy the OmniRoute search POST or the Firecrawl page read as the packet bus. Do not add a second API binary beside `cmd/denctl`. Do not copy a channel product or a webapp in front of a submit that does not exist. Do not install a second memory product; `scripts/stand-memory.sh` already refuses that.

## One plan line

On the mux and socket `cmd/denctl` already serves, accept one seven-field packet and return one result, and teach `cmd/dentui` to send that packet and show that result on the event stream it already reads.

## Unproven

This note did not run denctl, dentui, denfetch, or `scripts/stand-memory.sh`, and it did not call OmniRoute or Firecrawl. It does not show that `GET /api/events` can carry a packet result without a new handler.
