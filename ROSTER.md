# Roster

Researchers are r1–r10. Grunts are rNg1–rNg5. Each has its own git branch and worktree.

- `r1` Packet schema and validator — grok-4.7
  - `r1g1` omniroute-fast — MCP JSON-RPC message
  - `r1g2` omniroute-coding-fast — A2A task object
  - `r1g3` omniroute-cheap — Read ~/Projects/agent-mesh for task submit, claim, and kb put
  - `r1g4` omniroute-glm — Read ~/Projects/deer-flow only as a comparison of run state
  - `r1g5` omniroute-minimax — Write a Go struct for the seven fields plus the validation function

- `r2` Turn loop and four exits — omniroute-claude-opus
  - `r2g1` omniroute-gemma — Vercel AI SDK tool loop
  - `r2g2` omniroute-pro-fast — OpenAI Agents SDK run loop
  - `r2g3` omniroute-mimo — Claude tool_use stop
  - `r2g4` omniroute-llama — Four exits after the model: final message, one tool call, halt, empty
  - `r2g5` omniroute-thrifty — The seat does not plan inside the call

- `r3` Append-only store — grok-4.6
  - `r3g1` omniroute-fast — Mesh kb put and kb query
  - `r3g2` omniroute-coding-fast — Refusal row
  - `r3g3` omniroute-cheap — LangGraph checkpointer versus store
  - `r3g4` omniroute-glm — Single writer
  - `r3g5` omniroute-minimax — Write the store interface: put, query, append refusal

- `r4` Gate before tool and before store — omniroute-gemini
  - `r4g1` omniroute-gemma — The check finishes before the tool starts
  - `r4g2` omniroute-pro-fast — Throw before persist
  - `r4g3` omniroute-mimo — Fail closed
  - `r4g4` omniroute-llama — Do holds the tool's required arguments or the tool does not run
  - `r4g5` omniroute-thrifty — Rail at the gate

- `r5` Child handoff — omniroute-gpt
  - `r5g1` omniroute-fast — The parent hands the child one packet
  - `r5g2` omniroute-coding-fast — The child cannot write the parent's Write, Next, or Rail
  - `r5g3` omniroute-cheap — A handoff that transfers the conversation is refused
  - `r5g4` omniroute-glm — One child
  - `r5g5` omniroute-minimax — Write the handoff function signature

- `r6` Skill loader — omniroute-claude-sonnet
  - `r6g1` omniroute-gemma — Catalog is name and description
  - `r6g2` omniroute-pro-fast — Do not inject the whole catalog into the packet
  - `r6g3` omniroute-mimo — The procedure does not name a model
  - `r6g4` omniroute-llama — Supporting files load one level deep, only when a step reaches them
  - `r6g5` omniroute-thrifty — Gotchas, examples, and verification are a quality bar on the body, not an eighth field

- `r7` Security controls — omniroute-deepseek
  - `r7g1` omniroute-fast — CLI is argv only
  - `r7g2` omniroute-coding-fast — Read and Write are paths inside the section root
  - `r7g3` omniroute-cheap — Redact secrets before a store write
  - `r7g4` omniroute-glm — A prompt cannot flip Rail
  - `r7g5` omniroute-minimax — A prose prohibition is not a control

- `r8` Fetch size and speed — omniroute-qwen
  - `r8g1` omniroute-gemma — The next fetch is the section, not the ledger bodies
  - `r8g2` omniroute-pro-fast — One skill body per step
  - `r8g3` omniroute-mimo — The tool list is the tools this section named, not the host catalog
  - `r8g4` omniroute-llama — Stream the model call
  - `r8g5` omniroute-thrifty — Concurrency is many seats on the mesh

- `r9` Steer without killing the run — grok-4.5
  - `r9g1` omniroute-fast — A steer line enters the live step
  - `r9g2` omniroute-coding-fast — The side record is written by the human
  - `r9g3` omniroute-cheap — Rail is a switch on the section or it is empty
  - `r9g4` omniroute-glm — The spectator can inject a line mid-run without a new section and without a halt
  - `r9g5` omniroute-minimax — Pick the software rule for the next fetch

- `r10` Git consensus of section edits — omniroute-kimi
  - `r10g1` omniroute-gemma — One file per worker
  - `r10g2` omniroute-pro-fast — A branch commits only its files
  - `r10g3` omniroute-mimo — Two finished notes and no lead cut
  - `r10g4` omniroute-llama — Conflict rule when two branches edit the same heading
  - `r10g5` omniroute-thrifty — The fetch open and the gate open stay open until the predicate passes

