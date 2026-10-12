# Release 0.1.0

Awareness: this tag would mark the libraries and the operator binaries. It would not mark a running multi-agent session. Do not tag until the checks below are run on the commit you tag.

## What a user can run

| Command | Does | Does not |
| --- | --- | --- |
| `dencat` | Prints the catalog. `-terms` prints the vocabulary. | Bind a profile. |
| `denorch open` | Creates `wave--seat` worktrees. | Push. |
| `denorch commit` | Commits paths the seat owns, conventional subject. | Commit another seat's files. |
| `denorch take` | Runs `go test ./...`, then merges the seat onto the wave. | Push. |
| `denmcp list` | Prints `mcp/catalog.json`. | Pull an image. |
| `denmcp spawn` | Starts granted images in parallel. | Start a provider with an empty image. |
| `denctl` | Serves loopback state and the page. | Start a seat or call a model. |
| `dentui` | Dials the socket. | Staff a seat. |
| `den` | Writes `.den/init.json`. Starts `denctl` and `dentui` when they are on `PATH`. | Onboard a repo or an MCP server. |
| `denread` / `denfetch` | Read a local file or a public page. | Hold a credential. |

## Check before a tag

```bash
cd /home/kivox/Projects/harness-plan
go test ./...
git status --porcelain
```

`go test ./...` must exit 0. `git status --porcelain` must be empty, or the tag must name the commit and not the dirty tree. Do not tag a tree that contains `.env`, a key, or `.den/`.

## Version

`VERSION` is `0.1.0`. The next library change that keeps these commands is `0.1.x`. A change that removes a command or a grant field is `0.2.0`.

## Not a release artifact

Plan notes, codemap files, and seat research branches describe the tree. They are not the product. Foxcore repos are not in this release. SearXNG stays outside this repo.
