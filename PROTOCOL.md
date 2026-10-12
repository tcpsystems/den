# Protocol

This repo is the software design of the harness in `plan.md`. It is not a second plan. It is not a running service.

## What to write

Write types, functions, invariants, file paths, and refusal rules. A note that only retells a vendor is not done. End every note with:

- Types
- Functions
- Invariants
- File path
- Adopt or refuse, against the seven fields already in `plan.md`: Section, Function, Read, Do, Write, Next, Rail

## What not to do

- Do not edit `plan.md` in this repo. Proposals go in your own file.
- Do not delete. Do not rewrite a decided line. Decided text is Architecture through Seat expansion in `plan.md`.
- Do not build a server, a container, or a second bus.
- Do not choose a model inside the mesh. The mesh is bus, store, and fetch.
- Do not mention quotas, auth failures, or the retired project name.
- A sentence is decided only when two finished notes state it and no lead cuts it. Mark your lines as drafts.

## Git

You are on your own branch in your own worktree. Commit only the file named in your assignment.

```
git add <your file>
git commit -m "<id>: <one line>"
```

That commit is then landed on local `main` by `scripts/land-on-main.sh`. The land runs only when the diff is exactly your file. It refuses any other path, including `plan.md` and another worker's file. A conflict aborts. Nothing is pushed.

Do not edit `plan.md`. A later sweep reads the landed diffs and reconciles regions into `plan.md`.

## Nested workers

This host cannot let a researcher start its own sessions. The five grunt branches for each researcher are already open. Grunts write their own files. Researchers write the section lead. A later merge reads the branches.

## Cluster fan-out

A cluster task spec is a single plain text spec, not a vendor format.

Four or more free models answer that one spec in parallel over OmniRoute
(`https://ai.tcp.systems/v1`). Each call is an independent HTTPS request, not
a subagent.

Each worker's raw output lands in `/tmp/harness-wave/taskN/`, one file per
worker (`workerM.txt`).

A consensus merge step picks exactly one implementation from the worker
outputs and applies it to the repo.

Workers never commit. Only the orchestrator reviews, tests, and commits the
chosen implementation.

A failing model id (4xx, 429, or an empty response) is skipped mid-run and
replaced from a verified pool of free models, so the fan-out still reaches
its target worker count.

## Opens already on the table

- Next fetch: Superpowers keeps steer, ledger, and Rail bodies out, pointer only. Graphify, AgentScope, and Spring let the Rail switch ride with the section.
- Gate: six notes leave Rail empty. Superpowers waits for human approval before the first implementation tool.
Pick one software rule if your assignment asks. Say which notes you followed. Do not pretend both won.
