# Coordination

Written 2026-10-12. Research draft. This does not change `plan/BASE.md`. One host, one daemon, seats that stop after one packet.

## The rule already written

`plan/BASE.md` Loop says: "Two notes must state a claim before it is decided. Evidence first." `PROMPT.md` keeps the same gate on the base file: "Closed. Ascend only after two notes agree." `plan/foundation.md` says one packet is one turn and the turn ends. `plan/stack.md` says the daemon is the only process that reads git. `cmd/den/main.go` starts one controller for one session on `den.sock`.

## What Raft is

The Raft site says: "A consensus algorithm is used to agree on the commands in the servers' logs." URL: https://raft.github.io/. Raft replicates a log of state-machine commands so a majority of servers can apply the same command at the same index after a minority fails. A planning note is not that command. Two notes that quote the same file are evidence, not a replicated log entry.

## Mechanism

Keep the two-note rule. Do not add Raft, and do not add a recorded vote. The daemon appends one row per note: packet id, seat id, claim, source path or URL, and the quoted line. A claim is decided only when two rows name the same claim and each row cites a source the other row also cites. The later reader checks the two rows against those paths.

Refuse a model vote as truth. Refuse a count, a percent, or a latency that has no cited row. Refuse a consensus algorithm whose failure mode is a lost server or a partitioned peer. This daemon has one process. A dead process is a stop, not a quorum.

## Assessment

adopt. Change this only if a second den process on another host must apply the same command after the first host is gone.
