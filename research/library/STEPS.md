# Library steps

The parent job is one research library under `research/library/`. A step below is split only where the ask hid two actions.

1. Create `research/library/` and write the reconcile rule in `RULES.md`.
2. Reconcile the card with `plan/` and with `research/library/cards/`. The card does not become a second store.
3. Record every source URL or path on the card.
4. Fetch the source into a scratch directory outside the repo.
5. Parse it. Drop a changelog, a license text, and a block that is a generated list with no procedure.
6. Write the parsed card into the repo.
7. Mark that card read-only.
8. An agent who must retire a card sets `status: deprecated` on it.
9. A person may delete a card. An agent may not.

OmniRoute and DeerFlow use steps 4–7 with the same drop rules.

10. Ingest `pterodactyl/panel`.
11. Ingest `pterodactyl/wings`.
12. Ingest `pelican-dev/panel`.
13. Ingest `pelican-dev/wings`.

14. Parse a file into headings and one-line claims.
15. Compress that parse into the card.
16. Where the file describes a procedure, add pseudocode under the claim.

17. Index the theme list in `INDEX.md`.
18. Index the document seeds in `llm-intel/docs.go`. Leave the price sources out.
