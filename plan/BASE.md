# Base

This is the architectural base. Later work extends this file. It does not reopen it.

## What this is

den is the harness. The user talks to one operator. That operator writes a packet and starts seats. Seats do not invent the want.

The projects sit in the nest on this machine. They are not one GitHub repository.

| Role | Path | Remote |
| --- | --- | --- |
| Harness | `~/Projects/harness-plan` on `develop` | `tcpsystems/den` |
| Trackers, prices, models | `~/Projects/tcp-systems` | `tcpsystems/radar`, private |
| Company panel | `~/Projects/org/foxcore-cloud/fox` | `foxcore-cloud/fox` |
| Node agent | `~/Projects/org/foxcore-cloud/foxd` | `foxcore-cloud/foxd` |
| Hypervisor agent | `~/Projects/org/foxcore-cloud/floki` | `foxcore-cloud/floki` |
| Doctrine | `~/Projects/org/foxcore-cloud/ideology` | `foxcore-cloud/ideology` |
| Old harness | `~/Projects/agent-mesh` | `foxcore-cloud/agent-mesh`, archived at `c985423` |

fox on Cloudflare is the edge brain. foxd and floki are outbound agents. den does not write their desired state. Hetzner is lab and staging. Checkout is one Pro price. The tcp price desk is research, not that price.

## Packet

Seven fields only: Section, Function, Read, Do, Write, Next, Rail. Rail changes only from the human side record. A child packet does not carry the parent's Write or Rail. One tool call, then the result. A failed tool is a refusal row, not a new plan.

The model id is in the packet. OmniRoute is a tool. A combo, a batch, or A2A runs only when `plan/omniroute.md` says yes. Embeddings are `127.0.0.1:8091` on this machine. From `tcp`, the same embedder is `http://core.ts.tcp.systems:8091` on Headscale only. If it is down, the seat waits.

## Memory

Fast and recent are local files. OpenViking at `127.0.0.1:1933` is the long store. A backup may copy fast and recent to the OmniRoute host after the note exists. That host is not the read path. A started agent keeps soul, persona, identity, context, and memory paths for its life. It does not rewrite that record.

## Operator

The TUI and `denctl` are two views of one operator. `denctl` listens on `127.0.0.1:8793`. `GET /api/events` is a server-sent event stream: `event: state`, an id, `retry: 3000`, and a comment keepalive. A `state` frame is sent only when the JSON changes. The page does not start or stop a seat. Start and Stop are a later packet.

## Loop

Worker, branch, analyse, consider, assess, evaluate, fix, design, implement, evaluate, feedback. Repeat until the change is the smallest one that holds. Then one skill, and stop. A section seat that is unsure gets one explore or one deep-dive crutch. The crutch returns sources. It does not decide. Two notes must state a claim before it is decided. Evidence first. `User context` is a labeled block after that.

## Safety

Copy these. Do not drop them. Argv only. Paths stay inside the section root. A prompt cannot flip Rail. A child cannot write the parent's Write, Next, or Rail. No secret in the packet, the repo, or `/api/state`. A match runs a disallowed action against a control we own. The spectator writes the premise and the refusal. No exploit steps, payloads, or flood generation.

## Done on develop

The plan notes. `cmd/denctl` and its tests. The route skill, the picker, onboarding, and the match runner. The nest list on `/api/state`.

## Done on section branches, not merged

`section/packet`, `section/store`, and `section/steer` each have Go and a passing `go test`. They wait for a second note before merge.

## Not the base

Twenty-five `r*` branches have one unreviewed commit. They are drafts. Eight worktrees are dirty and uncommitted. Several seats died on HTTP 401 from the OmniRoute route, or on a provider serialization error, or they repeated the assignment. Those are not decisions.

Slots marked `later` are names. They are not installed products. A Git forge, an Obsidian server, Sentry, D1, Workers as this listen address, QUIC, and WebSocket are not built.

## How to extend

Add a section branch from `develop`. Change one directory. Include a check. Push the branch. Do not commit to `main`. Do not merge an `r*` draft in a batch. Do not add a slot without a function and a test.
