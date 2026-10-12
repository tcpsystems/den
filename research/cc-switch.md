# cc-switch

One finished note for `farion1231/cc-switch`. It is a suggestion. It does not enter the harness plan until a second finished note states the same scaffolding claim and no lead cuts it. This file is not a second plan.

Pages opened: the repository README, `CONTRIBUTING.md`, `session-manager.md`, and the English user manual for providers, MCP, prompts, skills, sessions, routing, failover, configuration files, the interface, and settings. Manual version stated on the manual index: v4.0.4, last updated 2026-10-07. Quota bars, auth failures, and missing model ids are ignored.

## What it is

CC Switch is a local desktop manager (Tauri, Rust, React, SQLite) for several coding tools. Its job is to keep one store of providers, MCP servers, prompts, skills, and projects, then write a narrow projection into each tool's own config. The database on the computer is the single source of truth. Each tool's live config is a projection. Device files stay on that computer.

`CONTRIBUTING.md` draws three paths out of the backend: the database, atomic live-config writers (temp file plus rename), and a separate local-routing process. Routing is not the same path as writing a live config. The mesh analog is the database plus the live writer plus the fetch of what was stored. The routing process looks up a target and can move to another provider. That part chooses a model. It is not the bus.

Switch-mode tools keep one active provider. Coexist-mode tools can hold several providers in one live file and do not backfill on a switch. A provider switch replaces key fields only: endpoint, key, model name, API protocol, and a few compatibility options that belong to that provider. Plugins, hooks, permissions, MCP, comments, and settings the user added stay in the tool file.

## Best practices

The seven section names stay Section, Function, Read, Do, Write, Next, and Rail. CC Switch never adds an eighth stored field for display text. Notes, MCP description, tags, homepage, and docs may exist in the editor and are not written into the app config.

### Which fields must be non-empty

Before a live write or the first tool, Section and Write are the fields that must be present. Function, Next, and Rail may be empty.

MCP is the clearest form. Name is required, may not contain spaces, and cannot be renamed in place. That is Section. Transport is required. A local server also requires Command, and Command is the program only (`npx`, `uvx`); arguments are a separate optional list. An HTTP or SSE server requires a URL instead of Command. Arguments, environment, working directory, and headers may be empty. Description and links may be empty, and the manual says they are shown only in CC Switch and are not written to app configs. That is Function, and it is allowed to be empty. Enabling a cell updates the database, then writes that one app's file. That file is Write. If the app directory does not exist, the toggle does not error and the file is not written, except one coexist-mode app that creates its directory. Read must name an existing directory or the live write is skipped. The database row can still remember the switch. Next is not required: a server may be stored with every app switch off. Rail may be empty.

The first model-list call is the same gate in another shape. Fetch Models runs only after the API key and the endpoint URL are filled. A connectivity check is not that call. The README and the provider cards say the check only tests that the address is reachable and does not send a model request. A Claude provider requires the API key. A Codex provider requires `model_provider`, `model`, and `base_url`. Notes stay optional. Gemini's model and Claude's base URL may be empty. The record does not grow a field to hold them when they are absent.

A prompt requires a name. Its description is optional and is shown only in CC Switch. The markdown body is what an enable writes, into exactly one target file.

A live write also refuses a Read that does not parse. The configuration-file manual says CC Switch writes a file only if it can read it. A broken file produces "Nothing was written, so your configuration is unchanged", and every file stays as it was. Commands validate before the service layer. The write itself is temp file plus rename, so a failed replace does not leave a half file.

Do is non-empty only for a local first tool: the program name. It may be empty when the transport is a URL, and argument lines may be empty either way. There is no shell in that field.

### When a lead cuts NEXT

What is stored: the provider's already-saved key fields, and nothing from the temporary next choice. The edit manual says a model picked in the tool, such as `/model`, is not saved back onto the provider. After a switch away and back, the file returns to the value stored on the provider. Shared settings are not copied onto that provider. The same page says the old backfill is gone: shared settings already live in the tool file, and they are not stored on any one provider. Settings the editor did not change are not submitted, so they cannot overwrite what the tool wrote meanwhile. Values already stored on the provider are not deleted when they no longer take effect on a switch.

Who stores it: the switcher, into the provider row in the database. The lead does not write the corpus. Plugins, hooks, permissions, MCP, and the user's own environment stay in the tool file. The Grok writer is the sharp case. On switch-away it deletes only the table it wrote last time. The table name is recorded in `live-state.json`, not taken from whatever `default` points at now. A model the user changed inside Grok is left alone, and the previous provider's table is not left behind as if it were the user's.

The prompt panel is the same cut with a library step. If the target file holds text the library does not have, that text is saved as an "Original prompt" entry before the new prompt overwrites the file. The README's shorter line is the same order: existing file content is saved back to the library first. The store holds the previous text. The lead does not invent a corpus document for the discarded next.

What the cut leaves when the lead does not write the corpus: the tool's global settings stay, the temporary model is absent from the provider row, and a failed parse leaves every file unchanged. Deleting a prompt preset changes only the library entry and does not modify the target file. An enabled preset cannot be deleted until it is disabled. Disabling empties the target file (Pi deletes `AGENTS.md`) and leaves the preset in the library. Leaving the editor with unsaved changes discards them only after an explicit discard. They never become the stored next.

Routing is not a stored next. While routing is on, "Route here" changes the forwarding target for later requests and leaves the direct provider labeled Direct. Leaving routing, or quitting the app, writes the config again from that direct provider. The routing manual says this write does not depend on a backup taken when routing started. The first time any config is rewritten, the original is copied once to `~/.cc-switch/backups/live-first-write/`. That backup is not the next section.

`CONTRIBUTING.md` still says a switch-mode change backfills the current live config into the current provider, then writes the new one, with Claude Desktop excepted. The v4 edit manual says the new version does not do that whole-file backfill. Follow the manual.

A project switch, from the README and the interface page, saves the current provider, MCP, skills, and prompts back onto the previous project, then loads the other project. Claude Desktop's project snapshot keeps only the provider. The switcher stores that snapshot. It is not the lead writing the tool corpus. The project switcher is off until a human turns on "Show project switcher".

### What a parent hands a child

The parent is the database row: a universal provider, one prompt, one MCP server, or the skill master. The child is one app's projection.

A universal provider stores a name, an API key, and an endpoint, plus the list of apps to sync. Save stores the parent and does not touch the children. Save and Sync overwrites the linked provider in each selected app. Deleting the parent deletes those linked providers. The packet that moves is the key fields, not the child's plugins, hooks, or MCP.

Copy to other apps, on the prompt panel, is the packet a child may actually hold. Each chosen app receives its own copy. The copy is not enabled. Later edits do not cross. The child therefore does not receive the parent's enabled Rail, does not become the parent's current prompt (Next), and does not write back into the parent's entry (Write). Enabling still overwrites only that app's one target file.

MCP hands the child only the servers whose switch for that app is on, and only the fields that belong in the app file. Display text stays on the parent. A server the user added in the tool and never imported is left untouched. The same name edited in the app is overwritten by the database version on sync. Two apps are the exception on a bulk sync: servers that are not enabled for them are not deleted from their `mcp.json` unless the switch is turned off or the row is deleted.

Skills keep a master copy, default `~/.cc-switch/skills/`, and project it into each app by symlink, or by copy if the link fails. The user can force symlink only or copy only. Import from this computer copies into the master and leaves the originals untouched. A ZIP whose directory name is already taken is skipped, not overwritten. Update replaces the local directory from the repository; local edits survive only in the backup taken first. Uninstall backs up to `~/.cc-switch/skill-backups/`, then removes the app copies, the master, and the database row.

The child may not touch three parent fields. Write is the database row and the skill master. Next is the direct provider that routing will restore, and any other app's target. Rail is the device switch: settings, the project-switcher toggle, the failover confirmation, and the keep-official-login option. A symlink the child can write through would let the child edit Write. That is not the handoff.

### Steer, side record, and section Rail

They are three objects.

The steer line is `live-state.json`. It records whether each switch-mode app is Direct, Routing, or Aggregation, and what was last written. Routing applies to subsequent requests. Clicking a mode tab only views that page. The button that names the consequence ("Start routing", "Back to direct") is the flip. Quitting writes the direct provider back first. The in-flight tool is not cancelled by viewing a tab. Restarting a Codex daemon is a separate interrupt and is not this steer line.

The human rail side record is `settings.json`, plus the first-write backup, the Codex login stash, and the confirmations. `settings.json` holds language, directory overrides, backup policy, and each app's direct-provider pointer. It is not cloud-synced. The configuration manual says not to edit `live-state.json` by hand. Unsaved-changes prompts, "the provider in use cannot be deleted", and the failover dialog "I understand, enable" are this side record. A seat packet does not write it.

The section Rail is the switch stored on the row itself, or it is empty. An MCP or skill cell stores that app's enabled bit on the database row, then projects it. A prompt's Enabled badge is that row's switch, one per app. A provider row's key fields are not that pointer. The pointer sits in `settings.json`. So the section Rail is not the steer line and not the direct pointer.

The next fetch may include the section row: key fields, and an enabled bit only when that bit is stored on the row. It may not include `live-state.json`, `settings.json`, the login stash, or `backups/live-first-write/`. Those stay on this computer and are outside cloud sync. A session resume fetches a different packet again: the resume command, and the project directory when the tool has one. The session-manager requirements note allows a missing directory and still runs resume. The shipped manual disables Resume when no resume command exists. Neither fetch loads the provider key or the device files. Codex history is per provider, so a switched provider may be unable to continue the old session. That is a missed resume, not a reason to put the other two objects on the fetch.

## Adopt

Before a live store write or the first tool, Section and Write must be non-empty, Read must name an existing app directory or the write is skipped, and Do must name the program when the tool is local, while Function, Next, and Rail may be empty and no field past those seven is stored.

When a lead cuts NEXT, the store keeps that provider's saved key fields and does not write the temporary model or the shared corpus into the provider, so if the lead does not write the corpus the cut leaves the tool's global settings in place and leaves no stored next.

A parent hands a child an independent copy of the prompt that is not enabled, and the child may not touch the parent's Write, may not replace the parent's Next, and may not turn the parent's Rail on.

The steer line, the human rail side record, and the section Rail are three objects, and the next fetch may include the provider row's key fields but not the live-state steer line or the settings.json side record.

A same-name skill or prompt target is skipped or kept as its own copy, not overwritten in place. A config that does not parse is not written. The previous prompt text is saved in the library before a new enable overwrites the file. The switcher saves a project snapshot back onto the previous project and does not treat that snapshot as the tool corpus.

## Refuse

Do not adopt `CONTRIBUTING.md`'s switch backfill. The v4 editor does not read the whole live file back into the provider. Doing so would store the shared corpus on one section.

Do not put failover on this loop. With the switch on, a failed request moves to the next provider in the queue, and the tray becomes read-only. That is the bus choosing a model. Aggregation also sends a request to whichever provider owns the model name picked in the client, and it does not fail over. The mesh stores the chosen direct id and fetches it. It does not pick the next id.

Do not let a child write through the skill master. Symlink-first, and copy-if-the-link-fails, can make the app directory and the master the same bytes. The handoff is the independent copy that is not enabled.

Do not delete from a seat. The shipped session manager deletes session files from disk and skips the Trash. Skill uninstall removes the master, the app copies, and the database row after a backup. Both are delete grants. The older session requirements note said the transcript stays local and read-only. Keep that, not the delete button.

Do not treat a second memory file on a tool, or a workspace editor, as the section store. The database is the store. The tool file is the projection.

Do not treat the connectivity check, a quota bar, or a missing model id as a section field or as a successful first tool.

## Park

Park local routing's format conversion as a later loop. It rewrites the live endpoint to `127.0.0.1` and the key to `PROXY_MANAGED`, and the real key stays in the database. That is useful only after a packet already names one direct model. It is not a step on this packet, and it must not grow a failover queue.

Park symlink projection as a later read-only loop. The master stays the Write. The app directory may point at it only when the child cannot write through the link. Until that is true, the child gets a copy.

Park the tray and the mode tabs as a later spectator. Viewing a tab changes nothing. A steer line still enters only through the button that states the consequence, the in-flight request finishes, and neither the tray nor `live-state.json` is the section store.

Park the session index's optional directory and summary. The requirements note says the index must hold session id, provider, and last active time, and that the directory, summary, and file path may be empty. The shipped resume still needs a resume command. That index is a fetch aid, not an eighth section field.
