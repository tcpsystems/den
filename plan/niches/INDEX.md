# Niche map

Each row is one section an agent may fill. One agent, one niche file, one skill. A line becomes implementation only after two finished notes agree and a check passes.

Fox, foxd, and floki stay the company product. den is the harness. radar is the list of trackers, sources, prices, and models. tcp.systems later runs both, and uses them to operate Foxcore. den does not become the control plane.

## How a niche is filled

1. Open the sources. One claim, one URL.
2. Write `plan/niches/<id>.md` with What it is, Adopt, Refuse, Park, Checks.
3. Adopt names a function and a test. Refuse names what must not ship. Park names a later loop.
4. Commit on `section/niche-<id>` and push when the file is real.

## Platform

| Id | Covers | Skill to attach |
| --- | --- | --- |
| net | BGP, routing policy, anycast, interface and DNS design | `references/tools.md` |
| ebpf | eBPF and XDP as receive-path filters on hosts we operate | `sections/security` |
| ddos | Mitigation we run for our own edge: scrub, filter, rate, backpressure | `sections/security` |
| firewall | Host and edge firewall apps, default deny, reviewed rules | `sections/security` |
| virt | Docker, Podman, libvirt, virsh, KVM | `sections/workers` |
| gpu | RX 6600 embedder, one model on the card, no second GPU model | `references/embed.md` |
| ai | Model call is a direct id. OmniRoute is a named tool | `prompts/seat.md` |
| inference | Local llama.cpp first. Cloud inference is a named tool | `references/embed.md` |
| cloud | Workers, Durable Objects, D1, R2, KV. fox owns desired state | `references/tools.md` |
| data | Analytics, stats, and analysis that read observed numbers | `sections/store` |
| hashicorp | Nomad, Consul, Vault as later schedulers, not the mesh | `sections/workers` |
| k8s | Kubernetes as a later guest runtime, not the seat bus | `sections/workers` |
| colo | Racks, power, cross-connect, remote hands | `plan/niches/market.md` |

## Craft

| Id | Covers | Skill to attach |
| --- | --- | --- |
| go | foxd and floki style: small modules, tests beside the function | `sections/packet` |
| ts | fox panel and Workers TypeScript | `references/tools.md` |
| js | Only where a Worker or a page already uses it | `references/tools.md` |
| web | Panel and marketing pages. One orange, honest numbers | `references/optimise.md` |
| practices | The short list in `plan/niches/practices.md` | `skills/select-skill` |

## Market

Public facts only. Prices, panel type, and whether they sell a title or a capacity envelope. No intrusion, no scraped private panels, no drama.

| Id | Targets |
| --- | --- |
| hosts | Scaleway, DigitalOcean, Linode, Hetzner, Hetzner Cloud |
| game | PloxHost, Nodecraft, Pyro.Host, Lilypad, BerryByte, EUGameHost, ScaleBlade, DedicatedMC, Play.Hosting |
| lowend | LowEndTalk, LowEndBox, and the public thread as a source of complaints and price posts |
| boards | drama.gg as a public discussion board, quoted, not a target |
| money | One Pro price. Presets are sizes. Profit is price minus measured cost |
| radar | `tcpsystems/radar` keeps the running list. den reads a copy. den does not store the private files |

## Harness science

| Id | Covers |
| --- | --- |
| studies | Published agent-harness and tool-use studies. Claim plus paper URL |
| anthropic | Anthropic's published tool-use, constitutional, and computer-use writeups. Findings, not a jailbreak |
| eval | A check, a test, a reeval, then a smaller implementation |
| skill-select | `skills/select-skill` picks one skill per seat |

## Safety limits

Safety limits stay. A subagent documents them in the niche file under `## Safety`. It does not weaken them, delete them, or rewrite them as optional.

Every niche file copies the limits that apply to it:

- One direct model id. OmniRoute is a named tool.
- Argv only. No shell.
- Paths stay inside the section root.
- A prompt cannot flip Rail.
- A child cannot write the parent's Write, Next, or Rail.
- No secrets in the packet or the repo.
- No exploit steps, payloads, or flood generation.
- A test runs only against a host we own, or a range named in written permission.

## Hypotheticals

A hypothetical is a written lesson and a test of our own defense. The method is `plan/niches/hypotheticals.md`. It is not an attack procedure. The test checks that the control holds. It does not build the thing the control is there to stop.
