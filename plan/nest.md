# Nest

The harness is a nest. The projects sit in it on this machine. They are not one GitHub repository.

| Place in the nest | Local path | GitHub |
| --- | --- | --- |
| Harness | `~/Projects/harness-plan` (`develop`) | `tcpsystems/den` |
| Price and tracker desk | `~/Projects/tcp-systems` | `tcpsystems/radar` (private) |
| Company panel | `~/Projects/org/foxcore-cloud/fox` | `foxcore-cloud/fox` |
| Node agent | `~/Projects/org/foxcore-cloud/foxd` | `foxcore-cloud/foxd` |
| Hypervisor agent | `~/Projects/org/foxcore-cloud/floki` | `foxcore-cloud/floki` |
| Doctrine | `~/Projects/org/foxcore-cloud/ideology` | `foxcore-cloud/ideology` |
| Old harness | `~/Projects/agent-mesh` | `foxcore-cloud/agent-mesh`, archived |

tcp.systems the host is up: Caddy, Headscale, maddy, and Tailscale are active. OmniRoute, Open WebUI, SearXNG, Vaultwarden, and RustDesk containers are up. `https://tcp.systems` answers, `https://ai.tcp.systems/api/health` is 200, and `https://chat.tcp.systems` is 200.

## Loop

A worker takes a branch. Then, in order: analyse, consider, assess, evaluate, fix or improve, design, implement, evaluate again, feedback. Repeat until the change is the smallest one that holds. Then assess the workflow, inject one skill, and stop.

The skill, the guard, and the tool belong to the project. The project sits in the organisation. The organisation sits in the harness. The harness sits on the host. The host keeps the skills, the guards, the tools, the ideology, the UI, and the same way of working.

## Complete enough to push

| Part | State |
| --- | --- |
| `develop` plan line | pushed |
| `cmd/denctl` live stream | pushed, loopback `127.0.0.1:8793` |
| `section/packet` | pushed, tests passed |
| `section/store` | pushed, tests passed |
| `section/steer` | pushed, tests passed |
| `section/niches` | pushed, the plan notes |
| `r*` drafts | one commit each on 23 branches, not merged |
| Controller Start/Stop | not built |
| Sentry, D1, Workers, QUIC, WebSocket | named slots, not wired |

## Stateless

Each event recomputes state from git and the probes. The server keeps the last payload only to skip a duplicate frame. It does not keep a session.
