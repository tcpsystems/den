# Stand-up

den is the harness. tcp.systems is the host that already runs the gateway. This note is how to get the harness up. It is not a second control plane and it is not a chat product yet.

Assessed from the tree on `develop`, listening ports, and `skills/codebase` plus `skills/evaluation`. A draft pass also read the dirty worktrees: `r10g2` and `r6` have real commits, `r4g4`, `r6g4`, and `r8g4` are placeholders and were not committed.

## What is already running

| Service | Where | Check |
| --- | --- | --- |
| Mesh hub | `127.0.0.1:8787` | `curl -fsS http://127.0.0.1:8787/api/health` |
| Embeddings | `127.0.0.1:8091` and Headscale `http://core.ts.tcp.systems:8091` | `curl -fsS http://127.0.0.1:8091/v1/models` |
| OpenViking | `127.0.0.1:1933` | `curl -fsS http://127.0.0.1:1933/health` |
| OmniRoute, chat, Caddy, Headscale, maddy | `tcp.systems` | `curl -fsS https://ai.tcp.systems/api/health` |
| denctl | `127.0.0.1:8793` when started | `curl -fsS http://127.0.0.1:8793/api/state` |

denctl was not listening at assessment time. Start it. Do not deploy it onto the OmniRoute host. That host is the gateway. The GPU and the git nest are here.

## Get started

On this machine, from the den checkout:

```
go run ./cmd/denctl -addr 127.0.0.1:8793 -repo .
```

Open `http://127.0.0.1:8793/`. The page is the operator view. It reads `GET /api/events` (server-sent events). It lists the nest: den, radar, fox, foxd, floki, ideology, and the archived agent-mesh. It does not start a seat.

The check is two curls: `/api/state` returns 200, and `/api/events` sends `event: state`. A second run is a reeval, not a new project.

## What the first build does not have

Chat. Multi-agent start. Docker. Those are the next packet, not a reason to add React, SvelteKit, or a public Worker.

## Next packet: chat on the same API

One input on the existing page. `POST /api/packets` appends a packet file under `~/.local/share/harness/memory/agents/operator/`. The body is the seven fields. Rail stays empty. The reply is the file path. No model call in this packet, so a missing OmniRoute key cannot 401 the page.

A later packet may hand that file to a seat. The seat reads `plan/operator.md` and keeps soul, persona, identity, context, and memory. Parameters other than the seven fields are optional. Safety limits are not optional.

## Docker

Docker is the third packet. `denctl` refuses to listen on anything but loopback, so a container that binds `0.0.0.0` is rejected. Do not weaken that check to make compose easy. When Docker is added, the process still listens on `127.0.0.1` inside a network that is not published, or the publish is host loopback only and the check is updated in the same commit as a test.

Do not put this compose on tcp.systems. OmniRoute, Open WebUI, and the mail stack stay there.

## Multi-agent

Multi-agent is already the branch rule. One orchestrator, many `section/<name>` branches, one skill each. `section/packet`, `section/store`, and `section/steer` have Go and a passing test. They are not merged. The `r*` drafts stay unreviewed. Placeholders are not commits.

## Expand later

Slots in `internal/console/state.go` are names. A plugin is a directory `plugins/<name>/`. Fill one slot with a function and a test. Do not add a framework to fill a name.
