# Repo split

The harness is `tcpsystems/den`. It consumes tools that already sit in the nest. It does not own them, and it does not invent a GitHub name for a directory that is still inside `den`.

The nest table in `plan/nest.md` is the list. A row below is a real checkout. A tool that has no row stays a path inside `den` until a human adds the row.

| Place | What | How den calls it |
| --- | --- | --- |
| `tcpsystems/den` | The harness and `denctl`. | It is the process. |
| `tcpsystems/mcp-hetzner` | Hetzner Cloud MCP server. | MCP stdio. |
| `cmd/denfetch` inside `den` | The public page reader. | One argv vector. |
| `sections/security/match.py` inside `den` | The workstation jail. | One argv vector. `path_jail` then `run_argv`. |
| `foxcore-cloud/agent-mesh` | The loopback seat bus. | WebSocket on `127.0.0.1:8787`. |

`denapi` is not a separate repository and not a separate command. Its handlers live with `denctl`.

There is no `tcpsystems/denfetch` repository and no `tcpsystems/sandbox` repository. Creating either is a later human decision, after two notes agree. Until that row exists, the page reader and the jail stay files in this tree.

## Protocol

Seats talk over `127.0.0.1:8787`. The hub is `meshd`. A seat claims a job, posts the result, and leaves. It does not poll. Halt is `meshctl halt --confirm`, and a worker token cannot issue it.

The packet that crosses the socket is the seven fields: Section, Function, Read, Do, Write, Next, Rail. No eighth field. The child is a new packet. It does not inherit the parent transcript, Write, Next, or Rail.

## Tooling

The harness edits a checkout that the nest table names. It does not vendor that checkout into `den`. A call is one argv vector. `run_argv` refuses `|`, `&`, `;`, `<`, `>`, and a newline. An API beside argv is a second door, and this note does not open one.

## Commit

Commit subjects on these branches are one sentence in the imperative. The body names the file and the rule it holds. A seat does not merge and does not push.
