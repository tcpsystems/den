---
name: use-git
description: Record one local change with git. Use before a seat commits, branches, rebases, or reads history. Does not push.
---

# Use git

Read this before `git add` or `git commit`. One change, one commit. Stop after the commit. Do not push.

## What the words mean

A commit is one snapshot and its parent. A branch is a movable name for one commit. `HEAD` is the commit that is checked out. The index is the snapshot the next commit will store. The work tree is the files on disk. `git status --porcelain` is the check. Empty output means clean.

`git help` on this host groups the commands as: start (`init`, `clone`), change (`add`, `restore`, `rm`), examine (`status`, `diff`, `log`, `show`, `grep`), history (`branch`, `commit`, `merge`, `rebase`, `switch`, `tag`), and collaborate (`fetch`, `pull`, `push`). This skill uses the first four groups. Collaborate stays with the operator.

## This repo

- Path: `/home/kivox/Projects/harness-plan`
- Branch for shared work: `wave/20261012-skill-adoption`
- Author: `kivox <petar@petarmarkov.co.uk>`
- Subject: `type(scope): subject`
- Types: `feat` `fix` `chore` `refactor` `docs` `test` `perf` `style`
- A seat branch is `wave/<date>-<name>--<seat>`. Two worktrees cannot share one branch.
- The wave branch receives merges. A seat does not commit on it while another seat is writing.
- Do not commit on `main`.
- Do not commit `.env`, keys, tokens, or `.den/orch/`.
- Name every path. Do not use `git add -A` or `git commit -a`.

## Commands

| Job | Command |
| --- | --- |
| Where and how dirty | `git status -sb` |
| The change you will record | `git diff -- path` |
| Recent history | `git log --oneline -15` |
| Record | `git add -- path` then `git commit -m "type(scope): subject"` |
| New line of work | `git switch -c wave/YYYYMMDD-name` |
| Replay a seat onto a newer base | `git rebase --onto <new> <old> <branch>` |
| One commit | `git show --stat HEAD` |

## GitHub

`gh` talks to GitHub. `git` stores the local history. Use `gh pr create`, `gh repo view`, and `gh api` only when the operator asked for a remote action. Do not print a token. Do not push from this skill.

## Check

`git status --porcelain` is empty for the paths you named. `git log -1 --oneline` shows your subject. The branch is not `main`.
