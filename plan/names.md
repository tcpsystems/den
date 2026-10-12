# Names

One grammar. A name is a row here before it is a DNS record, a Worker, a binary, or a branch.

## Places

| Token | Owner | What it is |
| --- | --- | --- |
| tcp | host `tcp.systems` | The system. Hetzner `cax41`, `23.88.44.32`. Orchestrates Foxcore. |
| core | this workstation | Tailnet `100.64.0.2`. Embeddings and the local desk live here. |
| den | `tcpsystems/den` | The harness. Binaries `denctl`, `denapi`, `dentui`. Checkout `~/Projects/harness-plan`. |
| dan | later public app | Spoken name for the den frontend. Binaries stay `den` until a rename packet. |
| radar | `tcpsystems/radar` | Private catalog. Trackers, sources, prices, models. Checkout `~/Projects/tcp-systems`. |
| fox | `foxcore-cloud/fox` | Company panel. Cloudflare Worker. Desired state. |
| foxd | `foxcore-cloud/foxd` | Outbound node agent. |
| floki | `foxcore-cloud/floki` | Outbound hypervisor agent. |
| deer | local desk on core | DeerFlow. Published through the tailnet name below. |

fox, foxd, and floki stay the company products. den holds rules and skills. radar holds the price desk. tcp runs the doors.

## Domain grammar

| Shape | Meaning |
| --- | --- |
| `<label>.tcp.systems` | One public HTTPS door. Caddy on tcp. Upstream is loopback on that host. |
| `www.tcp.systems` | Alias. Redirects to the apex. |
| `<node>.ts.tcp.systems` | Headscale MagicDNS. Tailnet only. Nodes: `core`, `tcp`. |
| `<service>.ts.tcp.systems` | Public TLS name whose upstream is a tailnet address. Caddy on tcp terminates it. |

A new label needs a row in the door table, one upstream, and a Caddy site. The socket stays on the machine. Auth sits on `denapi` before a door leaves the machine.

## Doors

Live Caddy on tcp, read 2026-10-11. Apex redirects to chat.

| Door | Upstream | Role |
| --- | --- | --- |
| `tcp.systems` | redirect to `https://chat.tcp.systems` | Apex. |
| `chat.tcp.systems` | `127.0.0.1:3080` | Open WebUI. |
| `ai.tcp.systems` | `127.0.0.1:20128` | OmniRoute. |
| `vault.tcp.systems` | `127.0.0.1:8222` | Vaultwarden. |
| `hs.tcp.systems` | `127.0.0.1:8080` | Headscale. |
| `mail.tcp.systems` | maddy | Mail. |
| `rd.tcp.systems` | RustDesk hbbs/hbbr | Remote desktop. |
| `deer.ts.tcp.systems` | `100.64.0.2:2026` | Public door to the desk on core. |
| `core.ts.tcp.systems:8091` | `127.0.0.1:8091` on core | Embeddings. Tailnet HTTP only. Callers use this hostname. |
| `tcp.ts.tcp.systems:20128` | `127.0.0.1:20128` | Tailnet door to loopback OmniRoute. |

Catalog nginx in radar still proxies the apex to the price desk on `127.0.0.1:2096`. That file is the desk's vhost set. Live TLS is the Caddy table above. The price desk's public name is a later `domain` row when the live apex and the desk are the same site.

`deer.tcp.systems` still has a public A record to `23.88.44.32`. Caddy has no site for it. DeerFlow is on core.

`dan.tcp.systems` is the later Workers frontend. It calls the published den API.

## Layers

```
door        dns label
  edge      Caddy now, Worker when this file names one
    publisher   denapi  OpenAPI  127.0.0.1:8793
      daemon    denctl  ~/.local/share/harness/den.sock
        harness skills, rules, impl flow, eval flow
          project   tcp, den, fox, foxd, floki
            host    core or tcp
```

`dentui` dials the socket when it is local. Browsers and later monitors use the publisher.

## Cloudflare config names

The Worker `name` is the app token. The public URL is the product domain. A tcp door stays on Caddy until this file gives it a Worker.

| Kind | Pattern | Live example |
| --- | --- | --- |
| Worker `name` | app token | `foxcore`. Later `dan`. |
| `database_name` | product token | D1 `fox` |
| binding | role | `DB`, `NODE_SESSION`, `NODE_JOBS`, `ASSETS` |
| queue | `<product>-<job>` | `fox-node-jobs` |
| dead letter | `<queue>-dlq` | `fox-node-jobs-dlq` |
| public URL | product domain | `https://foxcore.cloud` |
| secret | name only in wrangler | value from `wrangler secret` |

fox config: `foxcore-cloud/fox/wrangler.jsonc`. dan config, when it exists: `cloudflare/dan/wrangler.jsonc` in the den repo, route `dan.tcp.systems`, upstream the published API. No socket in a Worker. No secret in the config file.

## Files

| Path | Holds |
| --- | --- |
| `plan/names.md` | This grammar. |
| `plan/stack.md` | Process layout. |
| `plan/architecture.md` | How a niche ships. |
| `plan/BASE.md` | Closed base. |
| `plan/niches/<id>.md` | One niche. |
| `plan/analysis/<subject>.md` | One analyse note. |
| `section/<name>` | One extend branch. |

`develop` integrates. `section/<name>` is one change. An `r*` branch is an unreviewed draft.

## Verbs

The packet Function is one verb. One verb, one output.

| Verb | Does | Writes |
| --- | --- | --- |
| seek | Find one file or URL. | A citation. |
| indepth | Read that one path. | A passage and its source. |
| analyse | State the evidence. | `plan/analysis/<subject>.md` |
| assess | Say whether it is present, and the limit. | One row. |
| eval | Run the named check. | Pass or fail. |
| test | Be that check. | `*_test.go` or the match runner. |
| domain | Add a name. | A row in this file, first. |
| run | Run one owned command. | A log line. |
| extend | Open one section branch. | `section/<name>` |
| visualise | Draw the current names. | One diagram, or the state stream. |
| expand | Add one nest row that has a function. | `plan/nest.md` |
| improve | Make the function smaller. The same test passes. | The section branch. |
| debate | Require a second note. | The note. A claim waits for it. |
| suggest | Offer one option. | A labeled suggestion. |
| ascend | Move a debated claim into the base. | `plan/BASE.md`, after two agreeing notes. |

## Subjects

These subjects use the verbs. They are not extra verbs.

| Subject | Verbs | Scope |
| --- | --- | --- |
| Research | indepth, analyse | One path or one URL. |
| Expansion | expand, extend | One nest row, or one section branch. |
| Ecology | assess | What the host can feed: cores, RAM, disk, the RX 6600, the tailnet. |
| Economy | analyse | One unit in the note. Token USD per 1M, GPU USD per hour, or host EUR per month. Figures come from radar. Foxcore's price is a separate note. |
| OmniRoute | assess, improve | One combo or one setting named in `plan/omniroute.md`. The fixed list is `harness-winner`. |
| Multi-agent | assess, improve | One orchestrator, one seat, one skill, one crutch. Fan-out is 4 readers and 1 skeptic. |
| Visual | visualise | The layer diagram above, or `GET /api/events`. |
| Debate | debate, suggest, ascend | Evidence, then a block labeled `User context`. Ascend only after debate agrees. |
