# Operator

The operator is the surface the user talks to. It is the main orchestrator. The TUI and the web UI are the lowest consumers of `denctl`. They are the harness framework. They are not two bosses.

The daemon is `cmd/denctl`. Checks, queues, and internal tooling belong there. It presents a Unix socket at `~/.local/share/harness/den.sock`, mode 0600. The TUI dials that socket. The browser cannot dial a Unix socket, so the same handlers are also on `127.0.0.1:8793` for the web UI only. Neither client holds the git reads or the probes.

The TUI is `cmd/dentui`, built with Charmbracelet Bubble Tea and Lip Gloss. It is the admin task manager: services, the nest, and branches ahead of `develop`. It reads the socket's event stream and paints a frame only when a `state` event arrives. Quit with `q`.

The web UI is the same data with web limits: loopback only, branch names inserted as text, no cookies, no secret in the payload. Plugins come later as local directories. Charmbracelet Crush, Soft Serve, and Fantasy are not this harness. Soft Serve is a Git server, and that stays off this host.

The user says the want once. The operator writes it into the packet, starts the seats, and shows what they are doing. A big leader's jobs live here when a person is at the keyboard: prompt cut, landscape, one skill, and a tiebreak only when a lower tier is not decided.

## What sticks to a started agent

An agent is not a blank call. The operator writes one record before the first tool call. The record stays for the life of that agent. The agent reads it at start and after every resume. It is aware of it. It does not rewrite it.

| Field | Sticks as | Example |
| --- | --- | --- |
| Soul | The purpose it must not contradict | Implement the packet. Do not drop a safety limit. |
| Persona | How it speaks | Short sentences. Evidence, then a labeled user-context block. |
| Identity | A stable id | `seat/store` on branch `section/store` |
| Context | The packet | Section, Function, Read, Do, Write, Next, Rail |
| Memory | Paths, not a copy of every note | fast and recent for this id, plus OpenViking for the long note |

Fast and recent files for that id live under `~/.local/share/harness/memory/agents/<id>/`. OpenViking keeps the long note tagged with the same id. The OmniRoute host gets a backup and is not the read path.

A child gets its own record. It may read the parent's purpose. It may not write the parent's Write, Next, or Rail.

## Plugins

The dashboard and the suite grow by local plugins. A plugin is a directory under `plugins/<name>/` with one `plugin.json` and one probe or one panel. The operator reads that directory. It does not download a plugin. It does not run remote code because a config line named it.

`internal/console/state.go` slots are the list of names. A plugin may later fill a slot that is marked `later`. Until then the slot stays a name.

## Talk

The user talks to the operator. The operator talks to seats by packets. Seats talk to each other by mesh messages and by child packets. They do not open an A2A task for a den seat.

Start and Stop on the page are still not built. When they are, the button asks first, and it only targets a den seat id from the record above.
