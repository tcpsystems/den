# Skill index

One row, one job. A seat loads the path in the last column and nothing else from this page. A row that says "source" is not installed. `develop-skill` adds that one skill in its own packet.

GitHub has no skill directory in this repo. Use `gh`. Do not invent a token file.

## Already a den skill

| Domain | Attach |
| --- | --- |
| This harness | `skills/use-harness/SKILL.md` |
| Pick one skill | `skills/select-skill/SKILL.md` |
| Add one skill | `skills/develop-skill/SKILL.md` |
| Onboard this machine | `skills/onboard/SKILL.md` |
| Deep reading of one path | `skills/indepth/SKILL.md` and `skills/seek/SKILL.md` |
| Research, eval, cut a prompt | `skills/use-harness/references/optimise.md` |
| Prompt shape | `skills/improve-prompt/SKILL.md` then `skills/optimise-prompt/SKILL.md` |
| Steps | `skills/breakdown-steps/SKILL.md` |
| Tokens | `skills/token-counts/SKILL.md` and `skills/token-window/SKILL.md` |
| OmniRoute metric names | `skills/omniroute-metrics/SKILL.md` |
| OmniRoute call shape | `skills/use-harness/references/omniroute.md` |
| One tool call | `skills/use-tools/SKILL.md` |
| MCP or CLI | `skills/use-harness/references/tools.md` |
| Memory and embeddings | `skills/use-harness/references/embed.md` and `skills/use-harness/references/memory.md` |
| Guest and sandbox | `skills/use-harness/references/workers.md` |
| Argv, path jail, Rail | `sections/security` |
| Public page read | `cmd/denfetch` (OmniRoute search, then Firecrawl). Not a skill body. |

## Already installed outside the repo

The packet names the path. The body stays there.

| Domain | Attach |
| --- | --- |
| Git and GitHub | `gh` on the host. No skill file. |
| Cloudflare Agents SDK | `/home/kivox/.grok/skills/agents-sdk/SKILL.md` |
| Cloudflare products | `/home/kivox/.grok/skills/cloudflare/SKILL.md` |
| Workers | `/home/kivox/.grok/skills/workers-best-practices/SKILL.md` |
| Wrangler | `/home/kivox/.grok/skills/wrangler/SKILL.md` |
| Durable Objects | `/home/kivox/.grok/skills/durable-objects/SKILL.md` |
| Vercel AI SDK | `/home/kivox/.grok/installed-plugins/vercel-plugin-8fed268a/skills/ai-sdk/SKILL.md` |
| Web performance | `/home/kivox/.grok/skills/web-perf/SKILL.md` |
| Web design review | `/home/kivox/.agents/skills/web-design-guidelines/SKILL.md` |
| Writing | `/home/kivox/.agents/skills/writing-guidelines/SKILL.md` |
| React | `/home/kivox/.agents/skills/vercel-react-best-practices/SKILL.md` |
| Svelte | source: https://svelte.dev/docs/kit/introduction — not installed in this repo |

## One source, not yet a den skill

| Domain | Source | Do not copy |
| --- | --- | --- |
| Go | https://go.dev/doc/effective_go | A style catalog in the packet |
| TypeScript | https://www.typescriptlang.org/docs/handbook/intro.html | A second typechecker |
| JavaScript | the file's own runtime, Node or a Worker | A new bundler |
| Rust | https://doc.rust-lang.org/book/ | A second controller beside `denctl` |
| C | https://en.cppreference.com/w/c | A new runtime |
| C++ | https://en.cppreference.com/w/cpp | A new runtime |
| Linux | `man` on the host | A pasted man page |
| Chrome and Chromium | https://developer.chrome.com/docs/chromium | A disguised browser |
| Headless page | Firecrawl through `denfetch` | A second crawler, or a client that hides its name |
| Firefox | https://firefox-source-docs.mozilla.org/ | A crawl that pretends to be a person |
| Workstation sandbox | `skills/use-harness/references/workers.md` and `scripts/stand-memory.sh` | A second memory product |
| Harness runtime | `cmd/denctl`, `cmd/dentui`, `cmd/denfetch` | `ToolLoopAgent`, a 20-step loop |
| MCP servers | the operator config outside this repo. Hub unit `mcp-hub.service` | A printed bearer token |
| GitHub | https://cli.github.com/manual/ | A token in the vault or the packet |
| Code review | Anthropic code-review doc, already cited on `wave/20261012-review-loop` | A pasted skill pack |
| AppArmor, WireGuard | https://apparmor.net/ and https://www.wireguard.com/ | An exploit or a bypass |
| DNS | https://www.rfc-editor.org/info/rfc1035 | Publishing a zone |
| eBPF | https://ebpf.io/what-is-ebpf/ | Loading a program |
| Podman | https://docs.podman.io/ | A shared writable volume |
| OpenViking | https://github.com/volcengine/OpenViking | A replacement store |
| Qwen embedder | https://huggingface.co/Qwen/Qwen3-Embedding-8B-GGUF | A second model on the RX 6600 |

A language row is a handbook the seat opens while editing that language. It is not a reason to rewrite `denctl`.
