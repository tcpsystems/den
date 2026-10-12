# Observability seats and named tools

## Question

Which of Sentry, New Relic, Cloudflare MCP, Context7, DeepWiki, and Devin should den treat as a named tool, and which are a second store or a second harness? A tool is one name, one schema, one call. The packet is not a span. Prompts are not exported. Local notes may say one argv call ran.

## Sources

Fetched 2026-10-12 with TinyFish. Marketing lines below are marked unproven.

- Sentry MCP: https://mcp.sentry.dev/
- Context7: https://context7.com/docs/overview
- DeepWiki MCP: https://docs.devin.ai/work-with-devin/deepwiki-mcp

## Finding

A named tool is one name, one schema, one call, then the result. It does not pick the next seat and it does not replace the packet.

| Name | Class | Why |
| --- | --- | --- |
| Sentry MCP | Named tool, remote | One server, `https://mcp.sentry.dev/mcp`. The fetched page lists separate tools: `whoami`, `find_organizations`, `find_projects`, `search_events`, `analyze_issue_with_seer`. One of those names is one call. The Sentry issue store stays outside the packet. |
| Cloudflare MCP | Named tool, remote | Already one endpoint, `https://mcp.cloudflare.com/mcp`, after OAuth. Docs MCP stays a different name. Not fetched this note. |
| Context7 | Named tool into a second store | The overview names two tools, `resolve-library-id` and `query-docs`. Each call is one name and one schema. The library index is a second store. Its text may be read. It is not memory and it is not the packet. |
| DeepWiki | Named tool into a second store | The docs page names `read_wiki_structure`, `read_wiki_contents`, and `ask_question` on `https://mcp.deepwiki.com/mcp`. Each is one call against an indexed repo. That index is a second store. |
| New Relic | Not a named tool here | No page fetched. Do not invent a tool name. An APM account would be a second store of spans, not the packet. |
| Devin | Second harness | The DeepWiki page is "Ask Devin about any GitHub repository" and tells you to paste the MCP block into your Devin session. Devin is another coding agent with its own session. It is not one argv call and it does not sit in this loop. |

Sentry also documents a second product on the same host: MCP monitoring (span, request, session, connection, execution, notification, prompt-render). That product traces someone else's MCP server. It is not how den records a call.

## What den should not copy

- Do not turn the packet into a span, a session, or a prompt-render. Prompts are not exported.
- Do not send the section record, Rail, or a secret to Sentry, New Relic, Cloudflare, Context7, or DeepWiki.
- Do not add Context7's library index, DeepWiki's repo index, or a New Relic account as a second memory store.
- Do not let Seer, an Ask call, or Devin choose Next or rewrite the packet.
- Do not collapse a server's tool list into one name. `search_events` is not `analyze_issue_with_seer`.
- A local note may say that one argv call ran. It does not paste the prompt or the remote body.

## One plan line

Name Sentry and Cloudflare as remote tools, call Context7 or DeepWiki only as one read against a foreign index, and leave New Relic and Devin off the loop.

## Unproven

- Sentry's page says Seer "embeds the Sentry context your agents need — stack traces, breadcrumbs, traces, linked issues — into the prompt." Export behavior was not tested. Treat "embeds into the prompt" as marketing until one bounded call shows the returned fields.
- The same page says span tracking "helps you fix MCP servers faster" and "analyze patterns and identify systematic issues across all your MCP servers." No trace was opened. Those outcomes are unproven.
- Context7 says it "pulls up-to-date, version-specific documentation and code examples straight from the source." Freshness and completeness were not checked.
- DeepWiki says `ask_question` returns a "context-grounded" answer with relevant citations. Citation quality was not checked.
- New Relic was not fetched. Any claim that it is or is not an MCP tool is unproven here.
