# Initial state

Written 2026-10-12. Expansion of the locked foundation and of `plan/analysis/structure.md`. Not a second plan. `plan/BASE.md` changes only when two finished notes agree and no lead cuts them.

## Opinion

The harness starts small and stays small until a file earns its place. The initial state is one operator, one packet, one profile, one skill, one channel, one argv call, then a stop.

Profiles, skills, channels, MCP servers, and benchmarks are files. Adding one is a packet. Starting a second runtime for it is a leave.

The code that ships is the smallest change that finishes the job, checked once, in the language the file already uses. A shorter diff that drops a check at a trust boundary is unfinished.

## Profile

A profile is a text file. It names the seat's job and nothing else. It does not name a model, a temperature, a tool catalog, or a second memory.

The initial profile is the operator. Its job is to write the seven fields and start one seat. The seat's job is the Function line. When that line is done, refused, or cut, the seat stops.

| Profile | Job | Skill it may load |
| --- | --- | --- |
| Operator | Write the packet. Start one seat. | `skills/use-harness/SKILL.md` |
| Reader | Quote one page or one local file. | `skills/indepth/SKILL.md` |
| Critic | Break one claim against `plan/BASE.md`. | `skills/use-harness/SKILL.md` |
| Scribe | Write one note: adopt, leave, or cut. | `skills/use-harness/SKILL.md` |
| Builder | One argv change inside the section root. | the one skill the packet names |
| Reviewer | Read the diff and the callers. Write must-fix or should-fix. | the review skill below, one of them |

A new profile is one row in this table plus a file. It does not inherit another profile's transcript.

## Initiation

1. The human names the job in one sentence.
2. The operator writes Section, Function, Read, Do, Write, Next. Rail stays empty unless the human side already flipped it.
3. `select-skill` returns one path. No match returns `use-harness` only. A human-named skill wins.
4. The seat reads that skill at turn start. The body is not cached.
5. The seat reads the Read paths, makes one argv call, writes the row, and returns.

That is initiation. A channel, a server, and a benchmark are later packets. They are not part of boot.

## Token budget

The packet carries identity, one skill body, the Read excerpts, and the task. The catalog stays out. A second skill is a second packet.

Counts use `skills/token-counts/SKILL.md` and `skills/token-window/SKILL.md` when the packet names them. A count is a measurement. It is not a new field on the packet.

ECC `agent-harness-construction` says to keep the system prompt minimal and to load guidance on demand. That sentence is the one this note keeps. The rest of that file (hybrid ReAct, `next_actions`, pass@3) is a leave. den already has one argv call and a four-way exit.

## Review

Two review skills exist. The packet names one.

| Skill | Where | What the seat keeps |
| --- | --- | --- |
| Ponytail review | `DietrichGebert/ponytail` `skills/ponytail-review/SKILL.md`, MIT. Not copied into this tree. | Order: correct, safe, holds at the load this repo names, tested, fast, lean. A finding needs a concrete case. Groups: must-fix, should-fix, nice. |
| Code review | Anthropic review doc already cited on `wave/20261012-review-loop`. | One diff, the callers, a block or a pass. |
| ECC `security-review` | `~/Projects/deer-flow/skills/custom/ecc/security-review/SKILL.md` | Checklist only when the packet adds auth, secrets, or an endpoint. |
| ECC `verification-loop` | same tree, `verification-loop/SKILL.md` | Build, types, lint, test, secret grep, diff. Report pass or fail. |
| ECC `golang-testing` | same tree | Table tests for Go. One bench or fuzz only when the packet names it. |
| ECC `conventional-commit-review` | same tree | Subject is imperative. Body names the file and the rule. |

Ponytail itself (`skills/ponytail/SKILL.md`) is the build rule: the smallest complete change, deletion before addition, no new dependency for a few lines, one check on new risky logic. Never cut a check at a trust boundary. Levels stay `full` unless the human says `lite` or `ultra`.

The seat loads one of these files. It does not load the ECC directory.

## Language

The file picks the practice.

| File | Practice | Check |
| --- | --- | --- |
| `*.go` | Effective Go. `gofmt`. Table tests. | `go test` on the package the packet names. |
| `*.py` | The stdlib. `match.py` stays until a test says move it. | The test file beside it. |
| `*.ts` / `*.tsx` | The handbook. fox and rea only. | The check that repo already runs. |
| `*.rs` | The book. Only when the packet names a Rust file. | `cargo test` on that crate, later. |

A language server starts only for a file the packet named. Go gets `gopls`. TypeScript gets the TypeScript server for fox or rea. Rust gets `rust-analyzer` for a named Rust file.

## Channels, MCP, access

A channel is one path a result can leave on. The initial channel is the seat's Write field and the mesh post on `127.0.0.1:8787`. A second channel is a later note.

| Door | Initial state | Expansion |
| --- | --- | --- |
| Mesh | Claim, run, evidence. Worker tokens cannot halt. | None until a note names one added event. |
| MCP | One stdio server the nest table names, after a scan note. | One server per packet. Schema for that one call only. |
| Page read | `denfetch` argv. Public `http` or `https`. | No second crawler. |
| Access | Path jail, then argv. Rail is human-side. | A yes does not add a path. |
| Secrets | Env on the host. Never a packet field. | No token file in the repo. |

An MCP recommendation is one server for the job already named. Trust is a checkout the human uses, or a vendor page the reader quoted. A directory listing is not trust. The benchmark is one call with a known result.

## Runtime

Four words. They are checks, not phases the seat may skip.

| Check | Ready when | If it fails |
| --- | --- | --- |
| Ready | `denctl` answers on its loopback port, or the note says it is down and the seat is git-only. Embedder `:8091` answers, or the seat waits. OpenViking `:1933` answers for a memory read. | Wait or refuse. Do not swap in another model or another store. |
| Optimise | The hot path is one compare of the Function line, then a file read on a hit. A miss embeds a short key, not the corpus. | A second cache, a ranker, or a catalog in the packet is a leave. |
| Expand | One file: a skill, a profile row, a benchmark task, or a niche pack. One check. One note. | A plugin process, a community catalog in the prompt, or a probability picker is a leave. |
| Learn | One local row: pass or fail, the skill path, the function line. Exact repeat reads the file and does not re-embed. | The row does not reorder models and is not posted upstream. |

Learning is the ledger. It is not a trainer in the packet.

## Inconsistencies

When two files disagree, the seat stops and writes the conflict. It does not pick a winner.

Order of authority, top wins:

1. `plan/BASE.md` on `develop`.
2. `plan/nest.md` for repo names.
3. A wave note that two later notes have not yet agreed.
4. An ECC file, a ponytail file, or a card under `research/library/`.

A card is a quote. A skill outside the repo is a procedure the packet must name. Neither overrides the base.

The fix for an inconsistency is one note that says which line holds and which line is cut. The seat does not edit `plan/BASE.md`.

## Flows

```
human sentence
  -> operator packet
    -> one profile
      -> one skill
        -> read
          -> one argv
            -> row
              -> stop
```

```
diff
  -> one review skill
    -> callers
      -> must-fix or pass
        -> human merges or sends it back
```

```
new skill or profile or benchmark
  -> one file
    -> one local check
      -> one note
        -> stop
```

A child packet starts a new copy of the first flow. It does not continue the parent.

## Leave out

Loading the ECC tree. Loading every ponytail skill. A profile that selects the model. A channel per product. An MCP scan that adds servers by itself. A runtime that retries, repairs JSON, or resumes inside the same packet. A learn step that trains on the live card. Vendor pass rates as a result on this machine.

## Usage

| Action | Notes |
| --- | --- |
| Start | Operator packet. Rail empty. |
| Build | Ponytail `full`. Smallest complete diff. One check if the logic is new. |
| Review | Name one review skill. Concrete case or no finding. |
| Add a profile | One row, one job, one skill path. |
| Add a skill | `develop-skill` adds one. The index is a path, not a prompt. |
| Add an MCP server | Scan note first: present, maintainer, one job. Then one call. |
| Ready check | Ports answer, or the seat waits. |
| Commit | Imperative subject. Body names the file and the rule. Author `kivox <petar@petarmarkov.co.uk>`. |
| Merge | Human. A seat does not merge and does not push. |
