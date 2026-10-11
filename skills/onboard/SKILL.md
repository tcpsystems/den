---
name: onboard
description: Check which local services exist before a project starts. Report present or missing. Never read a secret into the packet.
---

# Onboard

Read `plan/start.md`. Check each row. Write one line per service: name, present or missing, and the path or command that proved it.

Do not open a key file. Do not copy a token. Do not pass a secret to a model. If a check would require printing file contents, skip it and mark the service present by path only.

Notion and Obsidian stay missing until the user names a token path or a vault. Do not search the home directory for secrets to fill those rows.

When the note is written, stop. The orchestrator starts the next seat. This skill does not.
