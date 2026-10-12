# OmniRoute

Read this only when the packet names OmniRoute. `plan/omniroute.md` says whether the call is allowed. This page says how that call is shaped. It does not add a default hop.

The gateway is `https://ai.tcp.systems`. A check on 2026-10-12 returned `GET /api/health` 200 with `status=ok` and a timestamp. That is the only live body this note records.

## Call

| Job | Method and path | When |
| --- | --- | --- |
| Health | `GET /api/health` | A seat checks the gateway is up. No key. |
| Models | `GET /v1/models` | The packet named a model id and the seat must confirm the id is listed. |
| Chat | `POST /v1/chat/completions` | The packet named a model id or a combo. |
| Search | `POST /v1/search` | A public-page search. `denfetch` is the client. |
| MCP | `POST /api/mcp/stream` | A gateway tool, and the local suite check shows it. |
| A2A | `https://ai.tcp.systems/api/a2a` | The other party is a remote agent card, not a den seat. |

The key is `OMNIROUTE_API_KEY` in the environment. Do not print it, do not write it into the packet, and do not copy it into the vault.

## Prefix

Identity, one skill body, and one tool schema stay a byte-identical prefix. The seven-field packet comes after that prefix. Changing the tool list or the model starts a new prefix. That rule is the prompt-cache note, not a second cache product.

## What OmniRoute already scores

These are OmniRoute's own numbers. A seat may read them when the packet names OmniRoute. A seat does not post a rating back.

| Read | Tool | What it is |
| --- | --- | --- |
| Health | `omniroute_get_health` | Uptime, memory, circuit breakers, rate limits, cache stats |
| Cache | `omniroute_cache_stats` | Prompt-cache tokens and the semantic-cache hit rate |
| Pool | `omniroute_pool_status` | Session counts, request totals, provider success rate |
| Route | `omniroute_explain_route` | Why one request id landed on a provider |

A read on 2026-10-12: prompt cache had seen 6,911 requests, 2,857 of them with cache control, and 58,226,295 cached input tokens. Semantic cache had 4 hits and 378 misses. Those figures belong to that read. They are not a harness score.

There is no MCP tool that accepts a success score, a response rating, a strength, or a failure from den. Do not invent that post. OmniRoute would treat it as a second control loop.

## Harness score

The harness keeps its own row, next to the packet, not on the gateway.

| Field | Value |
| --- | --- |
| Goal | The Function |
| Result | done, refused, no-progress, or cut |
| Accomplished | The Write path, or empty |
| Failure | The control name, or empty |
| Strength | The field that changed |
| Downfall | The field that did not change |

One row per turn. Append it. Do not send it to OmniRoute. A later seat reads it locally. Priority is the order in `plan/omniroute.md`: a named model, then a combo only when the packet allows it. The row does not change that order.

## Refuse

- Embeddings. Vectors are `127.0.0.1:8091`.
- Memory reads. Fast, recent, and OpenViking stay local.
- A combo when the packet already named one model id.
- The Vercel AI Gateway, a provider fallback, or a cost sort. Those are a second router.
- A 20-step tool loop. One call, then the result.
- Exporting the prompt or the packet to a vendor trace endpoint.

## Combos named in the plan

`harness-winner` only when the packet says the call may move among the scored models. `cost-clock` only for a cheap call that may follow the DeepSeek peak clock. A direct model id beats both.

Batch is for completions that are already written and need no tool in the middle. Compression runs only when the leader already cut the prompt and the packet still asks for a smaller body.
