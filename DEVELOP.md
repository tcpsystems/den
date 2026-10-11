# Develop

`main` is archived. Sections live on `develop`. An agent adds to one section, commits locally, and pushes that section branch when the diff is ready to compare.

The `r*` branches are the section drafts moved with this repo. They are not `main`.

## Sections

| Path | Owns |
| --- | --- |
| `sections/packet` | The seven fields and the validator |
| `sections/loop` | One call, then a result. Four exits |
| `sections/store` | Append-only store and memory layers |
| `sections/gate` | Check before the tool. Fail closed |
| `sections/child` | One packet in, a report back |
| `sections/skills` | Catalog, then one body |
| `sections/security` | Argv, path jail, Rail locked to the human |
| `sections/efficiency` | Fetch size and the local embedder |
| `sections/steer` | A line enters the live step |
| `sections/consensus` | Two notes before a line is decided |
| `sections/workers` | smolvm guests on the `harness` network |
| `plan` | The plan and the source notes |
| `skills/use-harness` | The route-first skill |
| `deploy` | Fresh Debian bootstrap |
| `prompts` | The injected seat prompt |
| `scripts` | Backup and scoring tools |

## Commit

Work in a branch named `section/<name>`, cut from `develop`. Commit only files in that section. Do not commit secrets. Do not commit to `main`.

Push the section branch when it is ready to be compared:

```
git push -u origin section/<name>
```

A review reads the diff against `develop`. It does not land on `main`.
