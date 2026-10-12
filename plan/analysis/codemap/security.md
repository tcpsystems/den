# Codemap: grant, vault, inject

Seat c4. Sources: `internal/grant`, `internal/vault`, `internal/inject`, `sections/security`. `sections/security/LEAD.md` and `r7g1.md`–`r7g5.md` are empty stubs. The live controls are the Go packages plus `sections/security/match.py`.

## Signed contract fields

`grant.Contract` is the signed payload. JSON names:

| Field | JSON | Role |
| --- | --- | --- |
| ID | `id` | Non-empty. Not re-checked after the signature. |
| Seat | `seat` | Non-empty. `Verify` compares it to the caller seat. |
| Parent | `parent` | Non-empty string inside the signature. `Verify` does not match it to a caller. The parent is the pinned ed25519 key. |
| Reason | `reason` | Non-empty. `TestContractRejectsEmptyReason` refuses a sign with no reason. |
| Scope | `scope` | Non-empty string list. `Canonical` sorts a copy before `json.Marshal`. Exact match later, no wildcards. |
| NotBefore | `nbf` | Unix seconds. |
| Expires | `exp` | Unix seconds. `Canonical` requires `exp > nbf`. |

`grant.Signed` is `contract` plus `sig` (standard base64 of the ed25519 signature). There is no public-key field. The comment on `Signed` says verify uses the pinned parent key, and `Verify` does that. `Sign` keeps the private key with the caller. The package comment says a contract never holds a secret value. Nothing in `Canonical` checks that.

## What Verify checks

`grant.Verify(parent, signed, seat, scope, now)`:

1. Rebuilds canonical bytes. Empty id, seat, parent, or reason, an empty scope, or `exp <= nbf` fails here.
2. Decodes `sig` as standard base64.
3. `ed25519.Verify` against the pinned public key and those bytes. A body edited after signing fails this, because the signature is over the rebuilt JSON, not the raw file bytes.
4. `Contract.Seat` must equal the caller seat.
5. `scope` must equal one scope string (`hasScope`).
6. `now.Unix() < nbf` is "not active yet". `now.Unix() >= exp` is "expired".

`TestSignVerifyScopeAndExpiry` covers a good `mcp:exa` check, the wrong seat `r9`, the wrong scope `mcp:other`, a time two hours past expiry, and a signature from another key. `TestContractRejectsEmptyReason` covers the empty reason at `Sign`. No test names the not-yet-active clock or a mutated field after `Sign`.

Callers that actually invoke `Verify`: `vault.Reveal` (`secret:<name>`) and `mcp.startOne` (`mcp:<id>`). `TestSpawnManyProvidersInParallel` shows `browserbase` is not started when the grant only names `mcp:exa` and `mcp:firecrawl`.

## What Reveal requires

`vault.Reveal(parent, store, signed, seat, name, now)`:

1. `grant.Verify(..., seat, "secret:"+name, now)` must pass.
2. `store.Get(name)` must return a non-empty value. `Map` is a plain `map[string]string`.

`TestRevealRequiresGrant` signs `secret:SEARCH_TOKEN` for seat `r1`, expects the stored value, and expects an error for name `OTHER`. It does not cover a bad signature, the wrong seat, expiry, or a missing store row.

`Redact` is separate. It rewrites `NAME=value` lines whose name contains `key`, `token`, `secret`, `password`, `passwd`, or `credential` (case folded) to `NAME=[redacted]`. `TestRedactKeepsNames` keeps `PLAIN=visible` and hides `OMNIROUTE_API_KEY`. Nothing calls `Redact` on the way into a prompt. The vault comment that a prompt, a log, and a catalog see names only is not a hook.

## What inject.One returns

`inject.One(idx, signed, seat, description, parentOK) (string, bool)`.

- `parentOK == false` or `description == ""` returns `DefaultPath` (`skills/use-harness/SKILL.md`) and `false`. `TestOneRefusesEmptyDescription` checks the empty description.
- Otherwise it walks `signed.Contract.Scope`. The first `skill:<id>` whose contract seat equals `seat`, and whose id is a `catalog.KindSkill` entry with a non-empty path, returns that path and `true`. `TestOneReturnsGrantedSkill` uses an unsigned `grant.Signed` with seat `den-1` and `skill:use-git` and expects `skills/use-git/SKILL.md`, `true`.
- Any other input returns `DefaultPath` and `false`. `TestOneDefaultsWhenUngranted` uses scope `mcp:exa`. `TestOneRefusesWrongSeat` uses seat `other` against caller `den-1`.

`description` is only a non-empty flag. It is not matched to a skill summary. The `contains(scope, "skill:"+id)` check repeats the prefix test in `skillID`. `One` does not call `grant.Verify`, does not read `sig`, and does not look at `nbf` or `exp`.

## Default-deny that is still only on paper

Enforced deny is `grant.Verify`: wrong key, wrong seat, scope string absent, clock outside `[nbf, exp)`, and the canonical empty-field checks. `Reveal` and `mcp.Spawn` use it.

Still paper:

- `inject.One`'s comment says an unsigned or expired grant returns `DefaultPath`. The function never checks the signature or the clock. `TestOneReturnsGrantedSkill` passes a `Signed` with no `Sig`. `parentOK` is a bool the caller invents.
- A failed `One` still returns `skills/use-harness/SKILL.md`. That is a default skill, not an empty deny.
- `Parent` is signed and then ignored. Any holder of the parent key can put any parent name in the contract.
- "The contract never contains a secret value" is a package comment. A scope string could be the value.
- Vault "names only" depends on someone calling `Redact`. `Reveal` will return the raw value once `Verify` passes.
- `sections/security` is not wired to `Verify`. `match.py` refuses three local cases, and `match_test.py` records them: `test_path_outside_root_is_refused_and_written_up` (`PathJail`), `test_shell_string_is_refused` (`RunArgv`), `test_packet_cannot_flip_rail` (`Rail`). `test_plain_argv_is_allowed` shows a clean argv is returned. `LEAD.md` has no invariants. The niche note that firewall apps are default deny is not implemented in this tree.
