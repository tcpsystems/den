# Skill adoption — capability to source matrix

Written 2026-10-12. Expansion, not a replacement. Each row is a capability the harness needs, a source that already ships it, and what to keep or leave. A source is a suggestion until two notes agree and a page read confirms the licence and the file shape.

## Capabilities and their established sources

| Capability | Established source | Keep | Leave |
| --- | --- | --- | --- |
| Profiling / observability | ECC `performance-optimization`, `trace-analysis` under `affaan-m/ECC/skills` | One measurement per call, record the slow field. | Vendor APM in the packet. |
| Evaluating / debating | `Picrew/awesome-agent-harness` ADRs and review loops | Two-claim note before a plan change. | Model vote as truth. |
| Contributing / collaborating | `kygura/void-agent` (fork of `badlogic/pi-mono`) headless child-agent orchestration | Child is an opaque task, parent merges. | Shared parent transcript. |
| Enhancing / planning | ECC `planning-*` skills | One plan file, one section, one check. | Plan that names files not in the tree. |
| Designing | Anthropic design language via Claude Marketplace plugins; `figma/mcp-server-guide` | Design tokens and component context only when the packet names a UI. | Design in every packet. |
| Testing | ECC `test-driven-development`, `verification-loop` | Red before green; local pass or fail. | Vendor SWE-bench as a result here. |
| Pen testing | ECC `security-review` + `semgrep` soft rules | AST/RegEx rules on CPU; report the file and line. | Exploit steps or flood generation. |
| Security enhancing | `cisco-ai-defense/securebert2`, CodeBERT detectors | Small score/label model, local only. | A chat model as the guardrail. |
| Best practice enforcing | `ai-boost/awesome-harness-engineering` catalog + `Picrew` shared language | One rule per skill, fail closed. | Importing the whole catalog. |
| Language guidelines | ECC per-language skills; Effective Go; TS handbook; Rust book | The file picks the practice. | A language server for a file the packet did not name. |

## Figma MCP

- `figma/mcp-server-guide` is the vendor guide. `GLips/Figma-Context-MCP` is a third-party context server. `github.com/topics/figma-mcp` lists the rest.
- Adopt: design context only when a UI change is on the packet. The server turns a selection into tokens, spacing, and colours.
- Leave: importing the Figma canvas into a coding worker that never touches UI.

## Void agents

- `kygura/void-agent` is a TypeScript monorepo: terminal coding agent, provider-neutral runtime, multi-provider LLM API, headless child-agent orchestration.
- Adopt: provider-neutral runtime and headless child orchestration.
- Leave: its terminal UI as the denctl replacement, and any Slack automation not named by a packet.

## ECC

- `affaan-m/ECC` is the source. `ecc.tools` is the index. Skills, instincts, memory, security, research-first development.
- Adopt one file at a time by description, exactly as `develop-skill` already does.
- Leave: loading the directory. The ECC operating system is an inspiration, not a second brain.

## How adoption happens

1. One skill, one description, one source URL. No multi-file clone.
2. `develop-skill` appends one index row and one short SKILL.md, unless an installed skill already matches.
3. A page read confirms the licence before any body text is copied in.
4. One local benchmark row proves the skill helps before it is promoted.
5. Two notes agree before the skill moves toward the base.

## Commit

Imperative subject. Body names the file and the rule. Author `kivox <petar@petarmarkov.co.uk>`.
