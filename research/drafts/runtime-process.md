# Runtime process

Research draft. This does not change `plan/BASE.md`. It is not wired into `denctl`. No implementation code.

The runtime is one process: the daemon already named `denctl`. Agents, teams, and the local mesh live inside that process. The TUI (`dentui`) only dials the daemon socket, displays state, and sends commands. It does not run agents. There is no second runtime binary for agents. `denapi` stays a loopback reverse proxy of the same socket (`127.0.0.1:8793`). It is a door, not a runtime. `meshd` on `127.0.0.1:8787` is a separate process today; this shape does not start it and does not grow a second mesh process. `denctl` only probes that address (`internal/console/state.go`).

## Words

One scheme. Four words. "Swarm" is rejected: it would mean the same thing as team.

| Word | Means | Does not mean |
| --- | --- | --- |
| Process | The daemon. One OS process, `denctl`. It owns the socket, the board, every team, every seat, and the local mesh tables. | A seat. A team. A tool the daemon execs. `den`, `dentui`, or `denapi`. |
| Seat | One agent turn inside the process. One packet, one skill, one argv call, then stop. | A goroutine farm, a model, or a second binary. |
| Team | A named set of seats inside the process. One parent seat, zero or more child seats. Created by a daemon command. | A swarm. A department outside the process. A lead model. |
| Mesh | The tables inside the process: section records, the current section, and the rules between them. Store and fetch stay in-process. | `meshd`. A model router. A second store. |

`plan.md` already says the mesh is the memory and the bus, and that it is not a model router. `plan/BASE.md` already says one operator writes a packet and starts seats. This draft keeps those meanings and puts both inside `denctl`.

## What the tree already does

Verified by reading this checkout (`module github.com/tcpsystems/den`):

- `cmd/den` loads `repo/.den/init.json` through `internal/initd`, saves it, starts `denctl -repo -sock`, then runs `dentui -sock -repo`. `.den/` is gitignored.
- The socket default is `~/.local/share/harness/den.sock`. `listenUnix` removes that path, then listens. Mode `0600`.
- `denctl` serves `/`, `/api/state`, and `/api/events` (SSE: `event: state`, id, `retry: 3000`, comment keepalive). It does not spawn agents. Comment on `defaultSock` still says the same handlers are also served on `127.0.0.1`. `serve` returns only `listenUnix`. `errPublic` ("denctl listens on 127.0.0.1 only") is unused. UNVERIFIED whether an older binary still bound that port. The source in this tree does not.
- `dentui` is a Bubble Tea client of that socket. Keys `w` / `p` / `f` / `s` call `initd.Enable` for `web`, `pages`, `files`, `seats`. `Present` is true only for `controller` and `tui`, so those four keys cannot start anything. The notice is "stays off. this build has no process for it".
- `denapi` reverse-proxies the socket to `127.0.0.1:8793` and refuses any other host. `cmd/den` does not start it.
- `console.Collect` probes `http://127.0.0.1:8787/api/health` and labels the result `mesh`. Nothing in `cmd/` starts `meshd`.
- A seat packet in `plan/BASE.md` and `plan.md` §3 has seven fields: Section, Function, Read, Do, Write, Next, Rail. `sections/packet/LEAD.md` drafts the same seven as `SectionPacket`. That file says it is a draft, not a decided line.
- `sections/security/match.py` defines `path_jail` (realpath must stay under the section root) and `run_argv` (non-empty argv; refuse `| & ; < >` or a newline in any argument). `set_rail` refuses unless `from_human` is true. No Go file imports or calls this module.

`plan/BASE.md` says `denctl` listens on `127.0.0.1:8793`. The working tree does not. The listener is the unix socket. `8793` is `denapi`. This draft follows the tree. `plan/controller.md` still describes the old TCP listen and a page that asks before start/stop. That page is not the runtime.

## Shape

```
den                         one command, session parent
  denctl                    the process (runtime)
    socket                  ~/.local/share/harness/den.sock
    board                   repo/.den/init.json
    mesh                    in-process tables
    teams                   in-process
      seats                 in-process, one packet each
    exec                    one argv, then stop (denfetch, denread, or a granted binary)
  dentui                    client only
denapi                      optional door, not started by den, not a runtime
```

Init stays `den`. That command still starts the daemon and then the TUI. Expansion is a command on the daemon socket, not a new process. The exceptions are tools the daemon execs: one argv vector, one call, then the child exits. `denfetch` and `denread` are tools of that kind. They are not agents and not a second runtime. A model call is a direct model id named by the packet. OmniRoute is a tool a grant names, not a path inside the daemon. UNVERIFIED: which binary, if any, the packet would exec for a model call. This tree has no model client.

The local mesh moves in. The probe of `127.0.0.1:8787` becomes a report of the in-process mesh (up, because the daemon is up), not a health check of a foreign process. An outside `meshd` is ignored. This draft does not say to stop or delete that binary. UNVERIFIED: whether anything on this host still depends on `meshd` speaking the agent-mesh protocol. `plan/BASE.md` marks `~/Projects/agent-mesh` archived at `c985423`.

`Present` stops meaning "a separate OS process exists". It means "this daemon build can run that part". `seats` becomes present because seats are tables inside `denctl`, not a binary. `web`, `pages`, and `files` stay off until a daemon command can actually serve or exec them. Turning a part on is `POST` (or the existing method the socket grows) to the daemon, which flips the board and starts the in-process work. The TUI key only sends that command. It does not call `Enable` on its own copy and hope the daemon notices. Today the TUI writes `init.json` itself (`initd.Save` in `dentui`). That split is how a screen and a daemon can disagree. The daemon becomes the only writer of the board.

## Defect 1 — TUI failure skips the kill

Verified in `cmd/den/main.go`:

```go
defer func() { _ = ctl.Process.Kill() }()
if err := tui.Run(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
}
```

`os.Exit` does not run deferred calls. A non-zero `dentui` exit leaves `denctl` running and the socket held.

Fix in this shape, still inside `den`, not a new supervisor:

1. On any TUI error, kill `denctl` first, then exit. Do not use `os.Exit` on that path. Return an exit code from `main` only after the kill, or call `Kill` in the error branch before `os.Exit`.
2. The kill stays a session teardown. The daemon is the runtime, but this command is "controller while the TUI is open" (`cmd/den` comment). A later "daemon stays up after the TUI" would be a different command, not this defect.
3. If the kill fails, print the error and the pid. Do not unlink the socket from `den`. The daemon owns the socket. UNVERIFIED: `Kill` on a already-reaped pid. The parent still has not called `Wait`, so the pid should still be the child.

Killing the daemon is safe for seats only because seats are inside it. There is no agent binary left behind. In-flight argv children must be in the daemon's process group, or the daemon must reap them in a signal handler before it dies. Otherwise a tool outlives the runtime. UNVERIFIED: `denctl` today sets no process group and has no signal handler. `cmd/den` does not set one either.

## Defect 2 — a second `den` unlinks the socket and orphans the first

Verified in `cmd/denctl/listen.go`: `os.Remove(path)` then `net.Listen("unix", path)`. The remove does not check who, if anyone, is listening. The first `denctl` keeps its listener with no name. Clients dial the new socket. The first process is orphaned.

Fix in this shape:

1. Before remove, dial the socket. If the dial succeeds and `/` or `/api/state` answers, refuse to start. Print the existing socket path. Do not kill the other daemon from `listenUnix`.
2. Remove the file only when the dial fails because nothing is listening (stale file). Then listen.
3. On daemon shutdown, unlink only if the path still refers to this listener. A second starter that lost the race must not be the one that deletes the live socket.
4. Hold a lock file next to the socket (`den.sock.lock`, same directory, `0600`) for the race where two processes both see a stale file. The lock is taken before remove. UNVERIFIED: `flock` behavior on this host's filesystem. The directory is under the user home, which is local disk on this machine; that is not a proof that the lock is reliable on every mount.
5. `den` itself checks the same way before `ctl.Start`. A second `den` exits without starting a second `denctl` and without starting a second TUI against a stolen socket.

The TUI does not get a "take over the daemon" key. Attach is dialing the live socket. Replace is a human stopping the first `den` (quit the TUI, which kills the daemon once defect 1 is fixed) and starting again.

## Team, seat, packet

Multi-agent and multi-team stay inside the one process.

Create a team. The TUI sends a command: team name, parent section id. The daemon allocates a team row: id, name, parent seat id (empty until the parent starts), state `idle`. No process is spawned. A second team is a second row. Teams do not share seats. A seat has one team id.

Give a seat a packet. The operator command names the team and the seven fields. The daemon checks the packet before the seat runs:

- Exactly those fields: Section, Function, Read, Do, Write, Next, Rail. No extras.
- Rail is empty, or a switch copied from the human side record. A packet from a seat cannot set Rail. `match.py` `set_rail` is the rule; the daemon must apply it in Go. Calling Python is not required. Wiring means the same refusal, in the process, on every packet. The Python file stays the spec until that Go check exists. It is not imported today.
- A child packet does not carry the parent's Write or Rail (`plan/BASE.md`).
- Read paths pass `path_jail` against that section's root. The section root is the section directory the packet names. UNVERIFIED: the on-disk root for a section id. `sections/<id>/` exists for the plan sections. A runtime section id may not be one of those directories.
- The one skill is one path, or none. The daemon does not load a catalog into the packet.
- The one call is one argv vector that passes `run_argv`. The daemon execs it, captures the result, and does not start another call for that seat.

Then the seat runs inside the process: one packet, one skill, one argv, stop. The seat row stores state (`queued`, `running`, `stopped`, `refused`), the packet, the tool result or the refusal row, and the team id. A failed tool is a refusal row, not a new plan (`plan/BASE.md`). Next is either stop or one section id the mesh fetches. The daemon does not let the model pick a different Next. A child seat is a new row in the same team, with its own packet, created by a daemon command from the parent seat's Next. The child does not inherit Write or Rail.

Two teams run as two sets of rows. The process is still one. Caps (how many seats run at once) are daemon settings, not extra binaries. UNVERIFIED: a numeric cap. None is in the tree. Do not invent one here; the first cut is one running seat per team.

## What the TUI shows

The TUI keeps the SSE stream. `/api/state` grows three read-only lists, filled by the daemon: teams (id, name, state), seats (id, team, section, state, one-line function), and the last refusal or tool result for each seat. The existing services, nest, and parts stay.

Keys that exist stay as command senders, not local board writers:

- `s` sends "seats on" to the daemon. The daemon marks the part present and on. The TUI learns that from the next state frame.
- A team command (name left as a daemon command, not a new key in this draft) sends the create. The TUI prints the new row when the state frame includes it.
- A seat line shows the seven fields only as a summary the daemon already stored. The TUI does not assemble a packet.
- Start and stop are commands to the daemon. `plan/BASE.md` says the page does not start or stop a seat and that start and stop are a later packet. This draft is that later step. The TUI still does not exec anything.

`dentui` exiting still only ends the view. Seats die only because `den` kills the daemon (defect 1). They do not move into the TUI.

## Commands the daemon grows

All of these are socket commands on the existing process. None of them is a new binary.

| Command | Effect |
| --- | --- |
| state (already `GET /api/state` and SSE) | Add teams, seats, mesh status. |
| enable part | Write the board. `seats` may turn on. `web`, `pages`, `files` stay refused until that part can run inside the daemon or as one exec. |
| create team | One team row. |
| post packet | One seat row in that team, then one skill and one argv, then stop. |
| stop seat | Mark the seat stopped and kill its argv child if it is still running. |
| stop team | Stop every seat in that team. The process stays up. |

`denapi` proxies these with the rest of the socket. It does not grow its own team table.

## What this cuts

- A second runtime binary for agents, workers, or teams.
- Starting `meshd` from `den`.
- The word swarm.
- The TUI writing `init.json` itself.
- `Present` as "an OS process we have not written".
- Shell strings. The call stays argv. `match.py` is the refusal list until the same check runs in the daemon.
- Any change to `plan/BASE.md`.

## Assessment

Load-bearing:

- The runtime is `denctl`, the process that already owns the socket. Seats, teams, and the mesh are rows in that process. Verified: the socket, the SSE state, and the lack of any agent spawn.
- The TUI is already a client. Keys already fail closed. Moving those keys to daemon commands matches the binary split that exists.
- Defect 1 is a real `os.Exit` skip of `Kill`. The fix is kill-then-exit in `cmd/den`.
- Defect 2 is a real unconditional `os.Remove` of the socket. The fix is dial-before-remove plus a lock, still in `denctl`.
- The packet has seven fields, and the jail and argv refusal exist only in `sections/security/match.py`. Go does not call them. Any seat that execs before that check is in-process is unfinished.

Guess:

- In-process mesh tables can replace the `8787` probe without a client that still needs `meshd`. The archive note supports that. It does not prove no client remains.
- One running seat per team is a stand-in cap. The tree has no number.
- Section root equals `sections/<id>/`. True for this repo's plan sections. Not shown for a seat section id.
- A lock file beside the socket is enough. Not tested on this host.
- Signal handling and process groups are absent, so a daemon kill might leave an argv child. Not verified with a running tool, because the daemon does not exec one today.
- `plan/BASE.md`'s "TUI and denctl are two views of one operator" still fits. Its listen address does not.

Cut:

- Do not add a worker binary, a team binary, or a mesh binary.
- Do not keep `meshd` in the start path.
- Do not let the TUI enable parts by writing the board locally.
- Do not treat `denapi` as the place seats run.
- Do not copy `match.py` into a second policy language. Call the same rules from the daemon, or port those two functions into Go and delete the unused path. Until one of those is true, say the jail is not in force.
- Do not reopen `plan/BASE.md` from this draft.
