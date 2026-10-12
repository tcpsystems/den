# Start

The controller plan is the ground. Before any project seat runs, one onboarding seat checks what this machine already has. It reports present or missing. It does not print a secret, copy a secret into the repo, or pass a secret to a model.

## Order

1. Wants. The big leader writes the user want. If it is missing, stop.
2. Onboard. Run the checks below. Write `~/.local/share/harness/memory/fast/onboard.md` with present or missing only.
3. Memory. Fast and recent are local. OpenViking at `127.0.0.1:1933` is the long store. The backup to the OmniRoute host runs after the note is written, and the note contains no secrets.
4. Analysis. The leader reads the onboard note and the landscape. Evidence first. `User context` is a later labeled block.
5. Seats. Only then does an orchestrator start. Each seat gets one skill from `skills/select-skill`.

## Checks

A check is "the file or service exists." The value inside is not read into the packet.

| Service | Local proof | This machine |
| --- | --- | --- |
| GitHub | `gh` on `PATH` and `~/.config/gh/hosts.yml` | present |
| Hetzner | `hcloud` and `~/.config/hcloud/cli.toml`, token file `~/.config/mcphub/secrets.env` | present |
| OmniRoute | `~/.config/omniroute/api.key` | present |
| Cloudflare | An MCP server named in the operator config. That file is outside this repo. OAuth on first use. No API token file is assumed | configured, not a key file |
| Tailscale | `tailscale` on `PATH`, status succeeds | present |
| Headscale | `headscale` active on `tcp` | present on that host, not a local token file |
| Notion | a token file the user names | missing until named |
| Obsidian | a vault directory that contains `.obsidian` | missing until a vault is named |

More services are the same shape. Add a row only when a path or a command is known. Do not invent a token location.

## Skill

`skills/onboard/SKILL.md` is the only skill for this seat. It refuses to `cat` a key file. The controller, when it is built, may show the same present or missing pills. It still does not show the secret.

## Not the start

Deploying fox. Creating Stripe prices. Turning OmniRoute into the default model hop. Printing `secrets.env`.
