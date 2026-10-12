# Gateway expansion — one tailored gateway layer

Written 2026-10-12. Expansion of the foundation, not a replacement. `plan/BASE.md` is unchanged until two notes agree.

## The idea, condensed

OmniRoute is one of many inference gateways. The goal is a single gateway tailored to the harness, not a second copy of OmniRoute run beside it. The harness sends one call. The gateway decides which account, which model, which route, and what it costs, then returns one result and one score.

## What it absorbs

- **Many providers, one door.** Grok CLI, Claude Code, Gemini, Cursor CLI, OpenCode, Kimi, AGY, and others already authenticate through their own CLI tools. Those tokens become provider accounts behind one gateway. A provider without a usable CLI token is a leave until a page confirms an endpoint.
- **OAuth from CLI tools.** Where a CLI stores an OAuth or API credential, the gateway reads the same place the CLI reads. It never prints the credential. Where the CLI does not expose a token, the gateway does not scrape a browser or a keychain.
- **Learning routing.** A routing layer that keeps a local ledger per route: cost per token, quality per call, latency, and pass/fail. A route improves when a measured row says it did. A route is demoted only by a local measurement, never by a vendor headline.
- **Value for token and value for quality.** Two scores, kept separate. Cost-per-token and quality-per-call are different axes. A route is cheap and slow, or expensive and reliable. The packet names which axis matters.

## Inspiration, not import

| Source | What to borrow | What to leave |
| --- | --- | --- |
| Google ADK | Pause on one tool call; the packet is the unit. | Nested session transcript, shared event list. |
| Claude Agent SDK | Permission mode as a gate, skills as files. | Permission mode as the jail. |
| Vercel AI SDK | Tool order, one finished result stays in history. | 20-step ToolLoopAgent, prepareStep rewrite. |
| Hermes | Flush memory before a summary. | Extra memory products, catalog in the prompt. |
| DeepSeek harness | Log argv before run, fail closed. | Unenforced jail claims. |
| Reasonix | Stable prefix then packet. | Trace export to a third party. |
| Aider | Repo map for the named file. | Whole-repo map in the packet. |
| Qoder | Quiet, precise, local-first edits. | Its payment or auth surface. |
| DeerFlow | Debounced per-conversation writes, regex tool budget. | Its sandbox as the harness sandbox. |
| LangChain / LangGraph | Checkpoint idea. | Checkpoint resume inside a call. |
| OpenHands SDK | One condensation, one write. | 60s abort as failure, context replay. |
| smolagents | One JSON tool call. | Code-action executor, Wasm executor. |

## The gateway profile

A gateway is not a seat. It is a route function with a ledger. It has no packet and no skill.

| Field | What it holds |
| --- | --- |
| Provider | The account the token authorises. |
| Model | The model id as the provider names it. |
| Cost | Measured cost of the last call, not a catalogue price. |
| Quality | Pass or fail against the local check. |
| Latency | Time to first token and total. |
| Ledger | A local row per route. Never posted to the provider. |

The gateway picks a route by the ledger. A route with no local row is a candidate, not a winner.

## Constraints the gateway must hold

- The gateway is one process on the host. It does not add a second router beside OmniRoute. It is the harness router.
- One call, one result, one score. No fan-out by default.
- A route that returns a 403, an empty body, or a repaired JSON answer is a failed row. No retry, no repair.
- The gateway does not post scores back to a provider, and does not export traces to a third party.
- The gateway reads only credentials the user has already installed. It does not generate, rotate, or exfiltrate tokens.
- Cost optimisation means fewer tokens and a cheaper correct route. It does not mean a model that is free but wrong.

## What changes in the foundation

Nothing changes in the seven-field packet. The gateway is invoked by a packet that names the route. A packet that does not name a route uses the current default. The gateway is an argv-callable route function, not an eighth packet field.

## What must be verified before this is built

1. Which CLIs on this host expose a usable token without a browser. Each is a page read, not a guess.
2. The protocol the gateway speaks. OpenAI-compatible `/v1/chat/completions` for most; the CLI-specific surface for Grok and Cursor.
3. The licence of each provider's endpoint. A CLI token is for the CLI. Using it through a gateway is a separate check.
4. One local benchmark row per route before the route is allowed to be picked.
5. The ADK/A2A/MCP protocol layers the gateway must announce, and which it must not.

## Commit

One sentence in the imperative. Body names the file and the rule. Author `kivox <petar@petarmarkov.co.uk>`.
