# Repo split

The harness is `tcpsystems/den`. It consumes tools, it does not own them.

| Repo | What | How den calls it |
| --- | --- | --- |
| `tcpsystems/den` | The harness, `denctl`, `denapi`. | It is the process. |
| `tcpsystems/mcp-hetzner` | Hetzner Cloud MCP server. | MCP stdio. |
| `tcpsystems/denfetch` | The crawler and page reader. | Argv. |
| `tcpsystems/sandbox` | The workstation jail. | Argv or API. |
| `foxcore-cloud/agent-mesh` | The WebSocket communication protocol. | Loopback WebSocket. |

## Protocol

`agent-mesh` is the communication layer. Den seats talk to each other over `127.0.0.1:8787`.
The foundation is a WebSocket hub. The proof is that `meshd` already runs the claim/run/evidence loop on this machine without HTTP polling overhead. A seat claims a job, posts the result, and leaves. It does not poll.

## Tooling

The harness develops the tools by editing their checkouts. It does not pull their code into `den`.
