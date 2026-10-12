# Memory scopes

The live note is written on core. The backup is a git tree. Restoring the git tree rebuilds the files. It does not rebuild OpenViking vectors until the embedder runs again.

| Scope | Holds | Live path | Backup path |
| --- | --- | --- | --- |
| global | Principles, concepts, analysis frameworks | OpenViking `viking://resources/global/` | `memory/global/` |
| thread | One conversation | OpenViking user of that thread | `memory/thread/<id>/` |
| group | A set of threads that share a job | OpenViking account for the group | `memory/group/<id>/` |
| session | One seat turn's working set | `memory/fast/` and `memory/recent/` | same names |
| team | den, fox, foxd, or floki | One OpenViking account per product | `memory/team/<product>/` |

A scope does not read another scope's account. A note that must cross products is committed in git.
