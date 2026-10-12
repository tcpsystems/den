# Review: expansion wave

Written 2026-10-12. Read only. Not merged.

## Read

| Branch | Tip | Parent | File |
| --- | --- | --- | --- |
| `wave/20261012-repo-split` | `744ff2b` | `0d05e2f` | `plan/analysis/repo-split.md` |
| `wave/20261012-hot-reload` | `4ac2bfc` | `0d05e2f` | `plan/analysis/hot-reload.md` |
| `wave/20261012-benchmarks` | `d86f1b7` | `0d05e2f` | `plan/analysis/benchmarks.md` |
| `wave/20261012-sdk-lessons` | `b5e1eb8` | `d86f1b7` | `plan/analysis/sdk-lessons.md` |

`plan/BASE.md` is the same blob on `develop` and on all four tips (`f120388`). None of them edit it.

Checked against the locked base and the rules the earlier wave already kept: seven packet fields, one argv call, one skill by description, local files plus OpenViking, embedder `127.0.0.1:8091` (Qwen3-Embedding-8B Q6_K), agent-mesh archived at `c985423`, foxd and floki outbound, Hetzner lab only.

## Verdict

No conflict with `plan/BASE.md`. Do not merge yet. `sdk-lessons` sits on `benchmarks`, so those two are one line, not two.

## Per branch

**repo-split.** Agrees. den stays `tcpsystems/den` and consumes tools instead of owning them. `mcp-hetzner` over MCP stdio and "edit the checkout, do not import the code" match the nest. Loopback `127.0.0.1:8787` matches the hub the base already names.

Two lines need a later note, not a base edit. `tcpsystems/denfetch` and `tcpsystems/sandbox` are not in the nest table. "Argv or API" for the jail is wider than argv-only. Calling archived agent-mesh the communication layer is a use of the old hub, not a second brain, as long as den does not write foxd or floki desired state.

**hot-reload.** Agrees. Reading `SKILL.md` at turn start is the file jail, not a resident catalog. Auto-pick by description, and not a cosine rank on port 8091, matches `select-skill` and leaves the embedder for memory. The web UI as an editor and a watch window matches the socket spectator: it does not run the model, merge, or start a seat.

**benchmarks.** Agrees. A local pass or fail row is not a vendor score. Rejecting the unquantized 70.58 MTEB number, Hindsight's LoCoMo 91.4, and other-SDK SWE-bench numbers keeps the locked 8B Q6_K embedder. A pass that keeps the description match does not install a learned picker.

One constraint: an offline pass on a 3B file on the GTX 1050 Ti must stay out of the live packet. The base waits when `127.0.0.1:8091` is down. It does not fall over to another model.

**sdk-lessons.** Agrees, and it is the cleanest of the four. Each row keeps a rule already in the foundation and names what not to import: one tool call, not a 20-step loop; one skill by description; argv validated before it runs, with a jail failure as a refusal rather than a retry; a child that does not see the parent thread; a stable prefix, then the packet; flush the note before the summary. The closer refuses an AST executor, a Wasm runtime, and an agent bus that copies parent state.

The commit also carries `plan/analysis/benchmarks.md` because the branch was cut from `benchmarks`. Review that file on `wave/20261012-benchmarks`.

## Merge order, when a second note exists

1. `wave/20261012-repo-split`
2. `wave/20261012-hot-reload`
3. `wave/20261012-benchmarks`, then fast-forward `wave/20261012-sdk-lessons`

Not done here.
