# Library rules

This directory is the harness research library. Memory is reconciled from these cards. A card is not a second copy of a memory topic.

| Path | Holds |
| --- | --- |
| `RULES.md` | This file. |
| `STEPS.md` | The split steps. |
| `INDEX.md` | One row per subject. |
| `cards/<id>/CARD.md` | Parsed claims, source list, status. |
| `cards/<id>/sources.jsonl` | One source per line: path or URL, bytes, sha256, keep or drop. |

`cards/` is one list. A card is not filed under a model, an SDK, a TUI, a platform, a language, or a provider. The subject name is the id. The directory above it stays the same for every subject.

## Read and write

After a card is parsed, an agent treats `CARD.md` and `sources.jsonl` as read-only. The next pass writes a new card only when the human asks for a refresh. The old card stays.

An agent retires a card by setting `status: deprecated` and writing one line, `reason`. An agent does not delete a card, a source line, or a raw file. A person may delete.

## What a parse keeps

English prose that states a procedure, an API, a type, a provider, an example, a notice, or a decision.

## What a parse drops

A changelog. A license or notice block that is only a grant of rights. A generated index with no procedure. A translation pair of the same page. A file under a runtime user directory.

The upstream file stays the source of truth. The card stores the claim and the pointer, not a second full copy of the upstream text.

## Reconcile

Read `cards/<id>/CARD.md`. Reconcile it with `plan/` and with the other cards. When they disagree, the card says so and does not overwrite `plan/`. This repository has no agent home. Do not read or write a path outside the repo to reconcile a card.
