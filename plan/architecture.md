# Architecture

den is a seat harness. tcp.systems is the system it folds into. Foxcore is the company that system operates. The three do not share a brain.

```
seat packet
  -> one skill
  -> one tool call
  -> local fast memory
  -> smolvm guest on the harness network
fox / foxd / floki
  -> desired and observed, Cloudflare edge, no inbound agent port
radar
  -> trackers, sources, prices, models
```

A seat never writes fox desired state. A guest never picks the next model. radar is read. The private files in radar stay in radar.

Big leaders sit above the orchestrators. Their job is `plan/leaders.md`: wants, prompt cuts, the deep plan, the landscape, one skill, and a tiebreak only when a lower tier is not decided. Evidence is presented with no slant. A later block labeled `User context` applies the known want. That block is the bias, and it is named.

## Implement

Each niche ships in this order. Skip a step and the change is not done.

1. Note. `plan/niches/<id>.md` with Adopt, Refuse, Park, and a URL on every adopt line.
2. Function. One function in the matching `sections/` file. Seven fields only.
3. Check. A test that fails when the function accepts the refused case.
4. Eval. Run the test. Record pass or fail in the niche file.
5. Reeval. A second run after a fix. The first failure stays in the note.
6. Harden. The path jail, the argv rule, and the Rail rule still hold.
7. Improve. Delete nothing. A smaller function may supersede.

## What is now

| Piece | Where |
| --- | --- |
| Plan | `plan/PLAN.md` |
| Route skill | `skills/use-harness` |
| Skill picker | `skills/select-skill` |
| Niche index | `plan/niches/INDEX.md` |
| Worker VM | smolvm, network `harness` |
| Embeddings | `127.0.0.1:8091` on the RX 6600 |
| Fast memory | `~/.local/share/harness/memory` |
| Backup | `tcp:/opt/harness-memory` |
| Old harness | `foxcore-cloud/agent-mesh`, archived |

## What is next

Fill the niche files. Do not start a second control plane. Do not put OmniRoute on the default model hop. Do not load a second model on the RX 6600.
