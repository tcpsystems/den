# Condensed & Optimized Harness Initialization

This document synthesizes the user's expansive requirements into a highly optimized, token-efficient, and actionable core prompt and architecture analysis, prepared prior to execution.

## 1. The Optimized Lead Orchestrator Prompt

**Role:** Lead Orchestrator (Unrestricted reasoning scope; restricted domain access)  
**Objective:** Orchestrate a hierarchical, multi-agent mesh to produce secure, DRY, and performant systems based on rigorous web standards (Mozilla, Vercel, Anthropic).  

**Capabilities & Protocols:**
- **A2A / MCP / AT Protocol:** Coordinate worker agents via A2A over Protobuf; connect to tools via MCP; maintain identity via AT Protocol.
- **Memory & Embeddings:** Query the local `Qwen3-Embedding-8B` (4096-d) via `127.0.0.1:8091` for exact-match ledger hits and semantic context. 
- **Routing:** Route complex reasoning to `omniroute-best-reasoning`, coding to `omniroute-coding:cheap` or `deerflow-lead`.

**Guardrails & Middleware (The "Soft Rules"):**
- **Permission System:** Default deny. You cannot read/write outside your jailed domain. Escalate to the parent orchestrator to request boundary adjustments.
- **Dependency Veto:** Reject unmaintained, insecure, or low-reputation packages automatically unless explicitly overridden by the User (Highest Privilege).
- **Anti-Loop:** If the programmatic watchdog detects looping (no state change across 3 turns) or a repeated failed fix, execution halts for human/critic intervention.

---

## 2. Deep Analysis: Watchdogs, Detectors, and Niche Loaders

To achieve the requested "models trained for skill probability, watchdog overseeing, and loop detection" without exhausting the baseline hardware (4-8 vCPU, 8-16 GB RAM, 8GB GPU), we must strictly separate **Generative AI** from **Classification/Heuristics**.

### The Programmatic Rule Engine (Soft Rules)
- **Implementation:** Abstract Syntax Tree (AST) parsing (e.g., Semgrep) and RegEx heuristics.
- **Target:** Instantly flags known bad practices (e.g., `eval()`, `chmod 777`, unchecked errors, hardcoded secrets, performative agreement phrasing).
- **Merits:** Zero token cost, sub-millisecond execution on CPU. 100% deterministic.

### The Watchdog & Skill Probability Transformers
- **Model Type:** Small encoder-only Transformers (e.g., MiniLM-L6-v2, DeBERTa-v3-xsmall) or CPU-bound logistic classifiers (FastText).
- **Function:** 
  - *Skill Probability:* Embeds the agent's task description and outputs a probability matrix for which `SKILL.md` to inject.
  - *Loop/Flaw Detector:* Computes cosine similarity between consecutive agent thoughts/actions. If similarity > 0.95 over 3 turns, it triggers the anti-loop middleware.
- **Merits:** Under 100M parameters. Consumes <300MB RAM. Runs instantly on the 4-8 vCPU baseline, leaving the 8GB GPU entirely dedicated to the `Qwen3-Embedding-8B` model.

### The Niche Loader / Memory Explorer
- **Implementation:** The existing OpenViking integration backed by the 8B Qwen embedder. 
- **Function:** When an agent requests niche knowledge (e.g., "Polar.sh specifics", "Stripe API"), the Niche Loader queries the vector ledger. Instead of feeding raw docs, a fast `auto/coding:cheap` route condenses the exact answer and returns it via A2A to the coding worker.
- **Merits:** Keeps worker contexts razor-thin (optimizing tokens/cache). Coding workers only receive actionable technical specs, not encyclopedic fluff.
