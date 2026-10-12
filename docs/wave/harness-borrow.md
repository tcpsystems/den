# Harness borrow

## Question

What one idea is worth borrowing from each of the Vercel AI SDK, the Cloudflare Agents SDK, Hermes, OpenCode, and Reasonix, and what is already refused?

## Sources

Search was OmniRoute at `https://ai.tcp.systems/v1/search`, through TinyFish. One public page per project, read with TinyFish fetch. No second crawler. Traffic was not disguised.

- Vercel AI SDK: https://ai-sdk.dev/docs/reference/ai-sdk-core/tool-loop-agent
- Cloudflare Agents SDK: https://developers.cloudflare.com/agents/runtime/lifecycle/agent-class/
- Hermes: https://hermes-agent.nousresearch.com/docs/developer-guide/agent-loop
- OpenCode: https://opencode.ai/docs/agents/
- Reasonix: https://github.com/esengine/DeepSeek-Reasonix/blob/studio/docs/GUIDE.md

## Finding

One idea each. None of these adds a field, a store, a router, or a policy language.

**Vercel AI SDK — `toolOrder`.** `ToolLoopAgent` can send tools in a fixed order. Unlisted tools follow, sorted alphabetically. The default loop is `isStepCount(20)`. Do not take that cap. Do not take `prepareStep` as a per-step prompt rewrite: swapping the tool set between steps is how a stable prefix gets invalidated. Borrow only a fixed tool order so the packet's tool list is byte-stable before the packet body.

**Cloudflare Agents SDK — one alarm, many named callbacks.** A Durable Object allows one `alarm()`. The Agent class stores many schedules and, when that alarm rings, calls one method by name with a payload. Borrow the shape, not the store: one seat wake-up, and the packet names the method. Do not take Durable Object SQL, `this.sql`, or `cf_agents_schedules` as a second store.

**Hermes — flush memory before the summary.** On compression, Hermes writes memory to disk first, then summarizes the middle turns and keeps the last N messages intact. A tool-call pair is never split. Borrow the order only: if a turn must be shortened, the keep is written before the summary, and a tool result stays with its call. Do not take `MEMORY.md`, the 500-iteration budget, or a fallback model chain.

**OpenCode — a read-only child for lookup.** Explore is a subagent that cannot modify files. The parent asks it to find files or answer a question about the tree. Borrow that as a child packet whose write set is empty. Do not take OpenCode permission keys (`allow`, `ask`, `deny` on `file edits` and `bash`) as a second policy language. The child still returns one report.

**Reasonix — a stable prefix, then the packet.** The planner system prompt is one stable string. Only a small per-turn instruction is appended on the user turn, so the prefix cache survives. The executor and the planner are separate sessions for the same reason. Borrow the order: the skill body and the tool list are the prefix; the packet is the suffix. Do not export traces to `api.openai.com`. Reasonix telemetry is a different pipe and stays out.

## What den should not copy

Already refused, and still refused after these five pages:

- `ToolLoopAgent`'s default 20-step loop
- the Vercel AI Gateway as a second router
- Durable Object SQL as a second store
- Hermes extra memory products
- OpenCode permission keys as a second policy language
- exporting traces to `api.openai.com`

Cache rule that stays: a stable prefix, then the packet.

## One plan line

Borrow five shapes only: a fixed tool order, one named wake-up, flush-before-summary, a read-only child packet, and a stable prefix before the packet.

## Unproven

These pages do not show that a fixed tool order or a stable prefix changes token cost on the mesh. They do not show that a read-only child packet is cheaper than the parent doing the lookup. They do not show that one named wake-up is enough when two seats are due at once. No page was timed, and no packet was run.
