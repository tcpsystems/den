# Graphify

One source. Suggestion only. This note does not decide a line in the harness plan, and it does not change that file. A line moves only when a second finished note states the same scaffolding claim and no lead cuts it.

Pack: `Graphify-Labs/graphify`, tree `v8`, plus the docs pages opened with it. The local engine is the open-source graph builder. The PyPI package name on those pages is `graphifyy`. The command is `graphify`. `/graphify .` is an assistant instruction, not a shell command. This note does not install it and does not build a graph.

The mesh stays the bus, the store, and the fetch. It does not choose a model. Graphify does not become that bus. Marketing recall numbers and token-spend claims are ignored.

## What it is

Graphify turns a folder of code, docs, PDFs, images, and video into files under `graphify-out/`. The skill's own description says it produces "three outputs: interactive HTML, GraphRAG-ready JSON, and a plain-language GRAPH_REPORT.md." The machine-readable file is `graph.json`. The human-readable file is `GRAPH_REPORT.md`. A browser view is `graph.html` when generated. Docs for outputs also list `manifest.json`, `cache/`, `converted/`, and `cost.json`. Those are sidecars, not the section record.

The code path is local and deterministic. `ARCHITECTURE.md` on `v8` says the pipeline is `detect → extract → build → cluster → analyze helpers → report.generate → export`. Code is parsed with tree-sitter. The docs page "How Graphify works" says a local code parse does not send files to a model. Documents, papers, images, and transcripts are a separate semantic path, and that path may send its input to a provider the operator selected. `INFERRED` does not always mean a model created the edge. Static resolution can infer it.

Every edge is supposed to carry a relationship and a confidence label. The architecture file defines `EXTRACTED` as explicit in the source, `INFERRED` as a deduction, and `AMBIGUOUS` as uncertain and flagged for human review in the report. The docs page "Edges and provenance" says the same three labels, and it says a path is evidence in the analyzed graph, not proof that a request will execute that path.

`validate.py` on `v8` runs before `build()` consumes an extraction. It requires `nodes` and `edges`. A node must have `id`, `label`, `file_type`, and `source_file`. An edge must have `source`, `target`, `relation`, `confidence`, and `source_file`. `file_type` must be one of `code`, `document`, `paper`, `image`, `rationale`, `concept`. Confidence must be one of the three labels. `assert_valid` raises and lists every error. The section record cannot grow to hold `file_type`, `relation`, or `confidence`. Those keys stay outside it.

The skill file `graphify/skill.md` is the packet the assistant follows. Before anything else it checks whether `graphify-out/graph.json` exists. If it exists and the user asked a natural-language question, and the user did not ask for a rebuild, the skill skips extraction steps 1–5 and runs `graphify query` immediately. If the user passed no path, the skill uses `.` and does not ask. If the user asked for help and nothing else, it prints Usage and runs no command. Otherwise it follows the numbered steps in order and does not skip them.

A store write is later, and it is guarded. After `build_from_json`, the skill checks node count before any write. Zero nodes is an error and `SystemExit`. It does not replace a good `graph.json`, `GRAPH_REPORT.md`, or analysis sidecar. `to_json` then returns false and writes nothing when the new graph is smaller than the existing one. The report and the analysis sidecar are written only when the graph file was actually written, "so they never describe a graph that graph.json doesn't contain." A later re-export says: if the shrink guard still refuses, surface the message and do not force past it. The human re-run is `--force`. The docs page "Keep the graph current" says the same thing: inspect the warning, then explicitly allow a smaller graph, and keep a copy of the old graph if you need to compare. Hooks can refresh in the background. A query immediately after a commit can still see the old graph. A pull or a merge needs an explicit `graphify update`.

Semantic work can be handed to children. The parent loads `references/extraction-spec.md` only when a chunk holds a doc, paper, or image. A pure-code corpus never reads that file. Each child receives that prompt verbatim, with only `FILE_LIST`, `CHUNK_NUM`, `TOTAL_CHUNKS`, `DEEP_MODE`, and an absolute `CHUNK_PATH` substituted. The prompt says to write JSON at that exact path and nowhere else. The parent, not the child, merges chunks into `.graphify_semantic.json` and later into `graph.json`. The skill tells the parent not to write an empty semantic file over a cache hit, because that drops every cached node. A missing chunk is a warning, not a silent success. One bad chunk is skipped and does not abort the run. If more than half the chunks are missing or invalid, the parent stops and tells the user. The skill wants those children to be `general-purpose` agents with Write and Bash, because a read-only child cannot write the chunk file. That grant is wider than this harness.

Work memory is a further file, not the graph. The docs page "Record useful findings" says it is "an optional layer over graph exploration, not a replacement for current source evidence." `graphify save-result` stores a question, an answer, node names, and an outcome of `useful`, `dead_end`, or `corrected`. `graphify reflect` writes lessons under `graphify-out/reflections/LESSONS.md` and can mark nodes preferred, tentative, or contested. Query and explain can show those hints, including a prompt to re-verify after the source changes. That overlay is not `graph.json`.

Graph file paths are constrained. `ARCHITECTURE.md` says `validate_graph_path()` must resolve inside `graphify-out/`. Labels are sanitized. URLs are http or https only. The graph is not the source tree.

PR commands inspect open pull requests and graph impact. The pull-request guide says merging and changing code remain separate actions. They are not a section write.

## Best practices

Map Graphify onto the seven fields only. The fields stay Section, Function, Read, Do, Write, Next, and Rail. Rail stays what the plan already says: empty, or a switch only the human flips. Do not reuse Rail as a confidence label.

The four holes, from the pages opened:

**Before a tool or a store write.** The skill's first gate is not a schema check. A help invocation stores nothing and calls nothing. A question against an existing `graph.json` is a query, and steps 1–5 do not run. The question is the non-empty Function. The path is Read, and a missing path becomes `.` instead of an empty Read or a question back to the user. The query does not need a new Section, a filled Do, a Write, a Next, or a Rail. The prior graph must already be stored, or the skill does not take that fast path. A store write is the other gate. The skill will not write until the ordered steps have produced a graph, the destination is `graphify-out/graph.json`, and the graph has nodes. An empty graph writes nothing. A smaller graph writes nothing, and the report is withheld with it. Next and Rail are not inputs to either gate. `validate.py` is stricter than the seven fields: it also demands `file_type` and edge `confidence` before `build()`. Those two have no home in the section record, so they do not become extra required fields on this packet.

**When a lead cuts the next store.** The cutter is the export guard, not a corpus edit. `to_json` writes nothing. The process prints `ERROR: refused to shrink graphify-out/graph.json` and exits. The previous `graph.json` remains the stored record. `GRAPH_REPORT.md` is not rewritten to describe a graph that was not stored. The source files are not the write target. The skill tells the agent not to force past the guard. `--force` is a later human re-run, which is the rail, not the lead. A help stop is the same shape one step earlier: print the usage text, run no command, store nothing. A missing child chunk is a printed warning. It is not a silent skip, not an empty overwrite of a cached semantic file, and not a clean stop. This source does not write a new section for the cut. It leaves the previous graph and the refusal text. A `dead_end` lesson is a separate optional command. It is not what the shrink guard stores.

**What a parent hands a child.** The child packet is the extraction prompt plus five substitutions: the file list, the chunk number, the chunk count, the deep-mode flag, and one absolute chunk path. The child writes only that path. The parent's `graph.json` is Write, and it is not in the child prompt. The merge and the following build are Next, and the child does not run them. `--force` and the shrink decision are Rail, and the child is not given them. The prompt does tell the child to mark uncertain edges `AMBIGUOUS` and not to omit them. That mark is a confidence label inside the fragment. It is not permission to flip the parent's Rail, and it does not fit the seven fields. The skill's demand that the child have Bash and a general Write grant is not adopted. One chunk key is the whole write.

**Steer, side record, and Rail.** They are three objects. The steer line is the live invocation: the question, `--update`, or `--force` typed for this run. The skill must not apply `--force` by itself when the guard refuses, and a background hook does not make the next query see the new graph. The human side record is the lesson file and, separately, `GRAPH_REPORT.md`. The report is written only after the graph write succeeds. Lessons are an optional overlay and "not a replacement for current source evidence." The section Rail field is the human switch already stored on the section. Graphify's fetch of `graph.json` returns the audit trail because that trail is stored on the edges. This harness does not have that trail on the section, so a section fetch includes Rail only when the stored section already has it. It does not include the live steer. It does not include a report that was withheld. It does not include lesson hints.

## Adopt

Suggestion. Before the first tool, Function and Read must be non-empty. Function is the question. Read is the path, and an omitted path is `.`, not a blank and not a question back to the user. A query tool also requires the already stored graph. Do, Write, Next, and Rail may be empty for that read. Before a store write, Section, Function, Read, Do, and Write must be non-empty: the graph has node identities, the steps have already run, and the destination is the one store key inside the store root. An empty graph or a smaller graph writes nothing. Next and Rail may still be empty. Do not add `file_type` or `confidence` to pass Graphify's validator.

Suggestion. When a lead cuts Next, the mesh stores no new section and does not write the corpus. The previous section stays. The report for the rejected graph is not written. The lead does not pass the human force switch. The record of the cut is the refusal itself, not a new section and not a clean stop. A help invocation that runs no command is the same cut with nothing stored at all.

Suggestion. A parent hands a child only that child's file list, chunk numbers, deep-mode flag, and one absolute write path, on the verbatim extraction prompt. The child writes only that path. The child does not touch the parent's Write, the parent's Next, or the parent's Rail.

Suggestion. The steer line, the human side record, and the section Rail field are three objects. The next fetch includes the stored section, so it includes Rail only when that section already has one. It does not include the live steer. It does not include a report or a lesson that was not stored on that section.

The local code parse is the only Graphify habit that fits a later read: deterministic, no model, and not a vector store. It is not adopted as this loop. The honesty rule that does fit now is narrower: do not invent a relationship, and do not describe a graph that was not stored.

## Refuse

Refuse Graphify as the mesh, the memory model, or a model chooser. The skill will use a host model, or Gemini when a key is already set, for semantic extraction. This harness does not let the graph step pick a model, and it does not read provider keys to decide one.

Refuse the semantic path on this packet. Docs, papers, images, video, transcripts, and hyperedges can leave the machine. The code parse does not. Do not mix those paths into one store write.

Refuse growing the section record. `file_type`, `relation`, `confidence`, `confidence_score`, community names, god nodes, and hyperedges do not become fields. Refuse writing `EXTRACTED`, `INFERRED`, or `AMBIGUOUS` into Rail. Rail remains the human switch.

Refuse a child with a shell or an open write grant. The skill says a read-only child silently drops the chunk, and it answers by giving the child Bash and Write. This harness keeps CLI as argv only. The child gets one chunk key, not a shell.

Refuse an empty overwrite and a shrink. Do not write an empty semantic file over cached nodes. Do not replace a larger stored graph with a smaller one. Do not force past that guard from a seat. Do not update the human report to describe a graph the store refused.

Refuse treating a path, an inferred edge, or a PR impact note as Next, as proof, or as a merge. The provenance page says a path does not prove a run. The pull-request page says merge and code change stay separate.

Refuse Graphify Cloud, hosted MCP, formal verification, wiki generation, and always-on memory as this loop. Benchmarks and token-spend claims are not scaffolding.

Refuse a lesson file as the section store. `save-result` and `LESSONS.md` are not `kb.put` of the section. A seat packet does not write that side record.

## Park

Park a later local graph read until a packet already exists. If it is ever used, it is a fetch aid inside a Read the section already named. It is not a second store, not a model router, and not the loop. Code only. No semantic backend. No embeddings. The mesh still fetches and stores the section.

Park confidence labels, lesson hints, preferred or tentative or contested marks, and community reports as an overlay after commit. Query and explain in Graphify may show lesson hints. This packet does not. The plan already keeps a training or harness-edit loop off the live packet. Treat reflect the same way.

Park `--force` on the human rail. The seat surfaces the refusal. Only the human re-runs a replacement, and only under the commit gate that already covers a widening write. A background hook is not a steer and not a fetch of the new graph.

Park PR inspection, HTML export, wiki articles, and watch mode. They are views and maintenance. They are not Section, Function, Read, Do, Write, Next, or Rail.

Park the extraction-spec rubric that tells a child to emit `AMBIGUOUS` instead of omitting an uncertain edge. The instinct not to drop uncertainty is sound. The field it wants does not exist on this record. Until a later loop has a place for it outside the seven fields, the child does not invent one.
