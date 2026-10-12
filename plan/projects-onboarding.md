# Projects onboarding

Fixed order: radar first (private, read-only), then foxcore-cloud doctrine
and product repos — ideology, skills, fox, foxd, floki. Den capabilities
named here: read-only observation, skills indexing, market research
workers, docs, code build-out.

## 1. radar

State: `~/Projects/tcp-systems` (`tcpsystems/radar`, private, `main`). README calls it a 2026-10-05 backup; repo is actively committed. No `.den/context.md` yet.
Needs first: read-only observation, then market research workers.
Tasks: write a read-only `.den/context.md` for `llm-intel/`, `tcp-passkey/`, `nginx/`, `systemd/llm-intel-desk.service` (no secrets, skip `private/`/`worker.env`) — assign one worker per group (GPU capacity; VPS and bare metal; web/container platforms) per `plan/market-research.md` — pull rows (provider, offer, region, price, billing unit, date, source URL) from `config/sources.d/`; a row becomes a Radar offer only when two workers agree on price and unit, and never writes into the model's `Price` record.
Done when: `.den/context.md` committed with no secret paths named; at least one row per group has two-worker agreement and shows as an accepted Radar offer, unaccepted rows stay visible as gaps; no account, trial, or server was created and no runtime loop changed.

## 2. ideology

State: `~/Projects/org/foxcore-cloud/ideology` (`main`), doctrine only. Committed `.den/context.md` exists. `CURRENT_STATE.md` maps each surface to its file and test.
Needs first: skills indexing, then docs.
Tasks: read `AGENTS.md`, `IDEOLOGY.md`, `CODEMAP.md` in that order — match each `CURRENT_STATE.md` surface to a skill already in `foxcore-cloud/skills` or a repo's own `.agents/skills` — record any surface with no matching skill as a gap, do not invent one.
Done when: every `CURRENT_STATE.md` row has a matched skill or a recorded gap; `.den/context.md` still matches the repo.

## 3. skills

State: `~/Projects/org/foxcore-cloud/skills` (`main`), 35 shared skills separate from `fox`'s own `.agents/skills`. Committed `.den/context.md` exists. `release-timeline.md` names the v1 gate order.
Needs first: docs, then skills indexing.
Tasks: confirm `release-timeline.md`'s gate order (`pnpm check` in `fox/`, `go test ./...` in `foxd/` and `floki/`) still matches those repos — index the 35 skills by niche/department for the `ideology` matching pass — record any skill the gate order names but the repo lacks.
Done when: the gate order is verified against the three repos' real check commands; the skill index is committed with no orphaned skill id.

## 4. fox

State: `~/Projects/org/foxcore-cloud/fox` (`develop`), SvelteKit panel on Cloudflare Workers/D1/Durable Objects, the customer-facing control plane. Own `.agents/skills/`. Committed `.den/context.md` exists.
Needs first: code build-out.
Tasks: run `pnpm check` and fix only files a seat may change per `.den/context.md` — run `pnpm test` — pick one `docs/` runbook (operators, instance-manager, billing) and confirm it matches current `src/routes` behavior.
Done when: `pnpm check` and `pnpm test` exit 0; the chosen runbook has no stale claim against `src/`.

## 5. foxd

State: `~/Projects/org/foxcore-cloud/foxd` (`develop`), Go node agent, rootless Podman plus crun, outbound-only. Committed `.den/context.md` exists.
Needs first: code build-out.
Tasks: run `go test ./...` (podman-dependent tests skip here as expected) — build with `CGO_ENABLED=0 go build -trimpath -o foxd ./cmd/foxd` and discard the binary — check the enroll/session/desired handshake in `internal/panel` still matches `.den/context.md`.
Done when: `go test ./...` passes (skips allowed, failures not); the build command exits 0.

## 6. floki

State: `~/Projects/org/foxcore-cloud/floki` (`main`), Go libvirt/KVM isolator for the `vps` kind, no listen address. `virsh` is not installed here, so `floki check` exits 1. Committed `.den/context.md` exists.
Needs first: code build-out, then docs.
Tasks: run `go test ./...` in `internal/virt` — run `floki check` and confirm it still exits 1 with "virt: virsh not installed" on this machine — note in `.den/context.md` that the missing `virsh` is expected here, not a defect.
Done when: `go test ./...` passes; `floki check`'s exit code and message match what `.den/context.md` documents.
