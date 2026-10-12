# Codemap — skills versus catalog

Seat c5. Wave `codemap-20261012`. Sources read in this worktree only.

## What the catalog binds

`internal/catalog/catalog.go` `Lead()` is the profile a session binds. Its skill list is one id: `use-harness`. The matching `Builtin()` profile row says the same: "den-1, builder, harness-runtime, auto/coding, use-harness, session-ledger."

```go
func Lead() Profile {
	return Profile{
		ID: "lead", Identity: "den-1", Role: "builder", Niche: "harness-runtime",
		Model: "auto/coding", Skills: []string{"use-harness"}, Memory: "session-ledger",
		Avatar: "den",
	}
}
```

`Builtin()` indexes two `KindSkill` rows and nothing else under that kind:

| ID | Path | Bound by Lead? |
| --- | --- | --- |
| `use-harness` | `skills/use-harness/SKILL.md` | yes, the only skill on `Lead().Skills` |
| `use-git` | `skills/use-git/SKILL.md` | no. Indexed, not on the lead profile |

`internal/inject/inject.go` is the only Go that returns a skill path. `DefaultPath` is `skills/use-harness/SKILL.md`. `One` returns a catalog path only when a grant scope is `skill:<id>`, that id is a `KindSkill` row, and the row has a non-empty path. No match returns the default. The catalog test `Bind(Builtin(), Lead())` checks that lead's skill id exists in `Builtin`.

`internal/catalog/terms.go` defines a skill as "one markdown playbook loaded when its description matches the step. Not a tool." Go does not read those markdown bodies. It stores the path.

## Front matter of the four named skills

| File | `name` | First job |
| --- | --- | --- |
| `skills/use-harness/SKILL.md` | `use-harness` | Route a seat to one reference. Load this file, then one reference. Do not load the catalog into the packet. |
| `skills/use-git/SKILL.md` | `use-git` | Record one local change. One commit. Does not push. |
| `skills/select-skill/SKILL.md` | `select-skill` | Pick one skill path for one subagent before a seat starts. |
| `skills/develop-skill/SKILL.md` | `develop-skill` | Grow the skill set by one. Discover, assess, append. Do not paste a catalog into the packet. |

`use-harness` routes memory, embed, workers, tools, and optimise to files under `skills/use-harness/references/`. It points skill selection at `../select-skill/SKILL.md` and niches at `plan/niches/INDEX.md`. Those are markdown links. No Go function follows them.

`select-skill` is the markdown table that names one path per task. Its last rule: if nothing matches, attach `skills/use-harness/SKILL.md` only. That fallback matches `inject.DefaultPath` and `Lead().Skills`, but the table itself is not called by Go.

`develop-skill` tells a seat to append one row to `skills/select-skill/SKILL.md` and one short `SKILL.md`. It does not add a `KindSkill` row. A skill written that way stays off `Builtin()` until someone edits `catalog.go`.

## Called by Go versus markdown only

Go names a skill path in three places: the two `Builtin()` `KindSkill` rows, `inject.DefaultPath`, and `inject_test.go` (expects `skill:use-git` to resolve to `skills/use-git/SKILL.md`). Plugins (`internal/catalog/plugin.go`) can append more rows from `plugin.json`. This tree has no plugin that adds a skill.

Markdown only, not called by Go:

- `select-skill`, `develop-skill`
- `onboard`, `token-window`, `optimise-prompt`, `breakdown-steps`, `token-counts`, `omniroute-metrics`, `challenge-claim`, `expand` (named in the select-skill table)
- every other directory under `skills/`: `alternative-tooling`, `assess-tooling`, `best-tooling`, `codebase`, `data-analysis`, `evaluation`, `improve-breakdown`, `improve-prompt`, `improvement`, `indepth`, `iterate`, `maintain`, `rule-tokens`, `seek`, `skill-tokens`, `use-tools`
- the five `use-harness` references (`memory.md`, `embed.md`, `workers.md`, `tools.md`, `optimise.md`)

`use-git` is the middle case. It is a `Builtin` skill path, so a grant `skill:use-git` makes `inject.One` return it. `Lead()` does not bind it. A lead session that never names that grant still gets `use-harness`.
