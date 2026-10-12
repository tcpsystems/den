# Prompt cache

## Question

How should den keep a prompt cacheable?

## Sources

- Lumer et al., Don't Break the Cache: An Evaluation of Prompt Caching for Long-Horizon Agentic Tasks, arXiv:2601.06007 (9 Jan 2026). https://arxiv.org/abs/2601.06007 and https://arxiv.org/html/2601.06007v1
- Anthropic, Prompt caching (Claude Platform Docs), fetched 2026-10-12. https://platform.claude.com/docs/en/build-with-claude/prompt-caching

## Finding

A cache hit is a byte-identical prefix, not a similar prompt. Anthropic checks the prefix up to the cache breakpoint and walks backward at most 20 blocks for an earlier write. The lookup order is `tools`, then `system`, then `messages`. A change in any block at or before the breakpoint changes the hash, so the prior entry misses.

Lumer et al. compared no cache, full context, system prompt only, and exclude-tool-results on DeepResearchBench (40 sessions per model per mode, 10,000-token system prompts). Best-mode cost cuts were 79.6% (GPT-5.2, exclude tool results), 78.5% (Claude Sonnet 4.5, system prompt), 45.9% (GPT-4o, system prompt), and 41.4% (Gemini 2.5 Pro, system prompt). Within a model, the three caching modes differed by about 2–4 points on cost. System-prompt-only and exclude-tool-results beat full-context caching on time to first token. GPT-4o full-context TTFT was 8.8% worse than no cache. The paper's rule: cache only the stable prefix (instructions, tool definitions, persona). Put timestamps, session ids, and user-specific text after that prefix. A changed tool list invalidates the prefix.

For den, the cached prefix is identity, one skill body, and one tool schema, in that order, and it stays byte-identical for the life of the seat. The seven fields come after the breakpoint: Section, Function, Read, Do, Write, Next, Rail. Rail stays empty unless the human flips it. Do not put the catalog, a second skill, or a second tool list in front. A tool-list edit or a model change is a new prefix, not a hit on the old one. Mark the breakpoint at the end of the tool schema, not on the packet.

## What den should not copy

Do not copy the paper's 10,000-token research prompt or full-context caching of tool results. Those results are session-specific and can add cache-write cost without a later read. Do not copy Anthropic automatic caching that walks the breakpoint forward over a growing history. Den's reusable prefix stops before the packet. Do not pad a short prefix to a vendor minimum. Do not treat another vendor's prices or minimums as den's.

## One plan line

Keep identity, one skill, and one tool schema as a fixed prefix, and put the seven-field packet after the breakpoint.

## Unproven

- No number in this note was measured on den. The 79.6%, 78.5%, 45.9%, 41.4%, 2–4 point, and −8.8% TTFT figures are the paper's, on models and prices as of early January 2026.
- Anthropic's current page prices cache writes and hits per model (for example Sonnet 5.5: $2 / $2.50 / $0.10 per million for input, 5-minute write, and hit). Those are not den's prices. Paper Table 3 prices are historical and are not copied here as current.
- Anthropic documents cache invalidation as tool, system, model-switch, image, or thinking changes. It does not state a separate rule that "any model change" is universal across OpenAI and Google. That cross-vendor claim is unproven here.
- Minimum cacheable length depends on the model. The fetched Anthropic page lists 512, 1,024, 2,048, and 4,096 tokens for different Claude models. The paper's January 2026 minimums (1,024 for GPT-4o, GPT-5.2, and Claude Sonnet 4.5; 4,096 for Gemini 2.5 Pro) are not current den limits.
- Whether den's identity-plus-one-skill-plus-one-schema prefix is long enough to cache on the seat's model is unproven. A hit is confirmed only when the response reports cache read tokens.
- Default Anthropic TTL is 5 minutes from the start of the request that writes or reads the entry, refreshed on use. A 1-hour TTL exists. Whether a den seat turn fits inside 5 minutes is unproven.
- The paper's own limit: results are one benchmark, one agent framework, and four models. They do not show that system-prompt caching is best for every den seat.

## Types

Prefix bytes (identity, one skill, one tool schema). Packet (seven fields). Cache breakpoint (one mark at the end of the schema).

## Functions

Build the prefix once per seat. Append the packet after it. Treat a tool-list or model change as a new prefix.

## Invariants

The prefix does not contain the packet, a timestamp, a session id, a second skill, or a second tool list. Rail is empty unless the human flips it.

## File path

`docs/wave/prompt-cache.md`

## Adopt or refuse

Draft. Adopt the prefix order and the seven-field suffix. Refuse an eighth field. Refuse full-context caching of tool results.
