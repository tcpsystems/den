# Hot reload and skill injection

## Config and skills

Den hot-reloads config and skills because the path jail reads the file when the turn starts. A skill body is a file. The harness reads `SKILL.md` when the description matches. It does not keep that body in memory. A change on disk applies on the next turn.

## Auto-injection

The human may name a skill. If the human does not, the harness picks one. The prompt is the query and the description is the match. `skills/select-skill` already does this. It is a lexical match, not a cosine rank over the catalog, and it does not embed the catalog on port 8091.

## Webapp

The web UI is an editor and a watch window. It writes plan text on disk. It reads the same socket `dentui` reads. It does not run the model, does not merge, and does not start a seat.
