# Critical Assessment of the Optimized Harness

This is a roast and feasibility check of the preceding expansion notes. It splits claims into what is grounded, what is feasible but unmeasured, and what should be cut.

## 1. Grounded and worth keeping

- **One packet, one skill, one argv call, then stop.** This is the only part of the plan that survives cost pressure. It is minimal, measurable, and matches the local score row.
- **Local embedder as Qwen3-Embedding-8B Q6_K on the RX 6600.** The size and 4096-dimension match the live index. Ollama as an alternate path is documented in the research report.
- **A2A, MCP, Protobuf/buf as named layers.** These are real protocols. Keeping them as separate layers is sound; the risk is wiring all of them at once.
- **AST/RegEx soft rules over an LLM for security.** This is the right call for cost. Semgrep-style rules run on CPU and are deterministic.

## 2. Feasible but unmeasured

- **CPU-bound classifier for skill probability.** A logistic classifier or FastText can work, but no measurement on this machine has been taken. The accuracy on arbitrary task descriptions is a guess.
- **Relativity graph joins.** The graph is a good mental model, but it is not implemented and no traversal cost has been benchmarked.
- **vLLM and Ollama "auto-configure".** vLLM on this host is heavy and Linux/ROCm-constrained. Ollama is the realistic local fallback. Auto-selecting between them needs a detection step, not a config file.
- **Embedding-based loop detection with a similarity threshold.** Repeated outputs can look similar while still making progress, and different outputs can be the same loop. A threshold is a heuristic, not a guarantee.

## 3. Should be cut or deferred

- **Full AT Protocol with DIDs, relays, and signed repositories.** For a single-host local harness this is overkill. Keep decentralized identity as a later goal; do not block the foundation on it.
- **"Unrestricted lead" as a property.** Every lead still runs inside path jail. The correct wording is "wide reasoning scope, same boundary." The human is the only unrestricted actor.
- **A watchdog "trained model" before a rules file exists.** Train only after the first 100 local rows are collected. Until then a fixed rule list is cheaper and inspectable.
- **Loading the 3B offline model on the 1050 Ti.** This is already out of the live packet. It stays out.

## 4. Pattern matching

Across all expansion notes the same three patterns repeat:

1. **A file plus a local check beats a process.** Skill, profile, benchmark, niche — all land as files.
2. **The packet is the unit.** Child packets are opaque tasks. No shared graph resume.
3. **Security is a deny-by-default boundary.** Path jail, argv metacharacters, and the rail are the only doors.

The future work is to make the relativity graph a concrete join on those files, and to add the CPU classifier only after a benchmark row proves it beats description matching.
