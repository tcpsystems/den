# OmniRoute

status: active

origin: container diegosouzapw/omniroute 3.8.51 path /app

memory: reconcile against the workspace topic that names this subject. This card does not replace that topic.

Kept files: 211. Dropped files: 1.

The lines below are a parse. The upstream tree remains the source.

### app/docs/DEVELOPER-ENVIRONMENT.md

- Headings: Developer environment notes; .env postinstall behavior; Recommended local workflow; Scanner notes; If you find a leak
- Claim: This page explains the project's local `.env` behavior and how to handle environment files and secrets when developing OmniRoute. The project may generate a local `.env` file du...

### app/docs/DEVIN_CLAUDE_BRIDGE.md

- Headings: Devin Claude Bridge; Architecture; Isolation and threat model; First-time setup and normal use; Validation commands; Updating pinned tools; Diagnosis and cleanup; Limits
- Claim: `devin-cli-agentic` lets the real Claude Code runtime use OmniRoute's local Anthropic Messages endpoint while the official Devin CLI supplies model responses over ACP stdio. It ...

### app/docs/OMNIROUTE_ALLOCATION_HANDOFF.md

- Headings: OmniRoute Allocation Handoff
- Claim: Allocation is not provider quota. Quota pools define which API keys may consume a provider pool and how hard, soft, or burst policies apply. Provider quota is external capacity ...

### app/docs/OMNIROUTE_PROVIDER_FAILOVER.md

- Headings: OmniRoute Provider Failover
- Claim: Failures are classified before retry decisions are made. Transient failures such as timeouts, network errors, rate limits, and provider 5xx responses may fail over. Authenticati...

### app/docs/OMNIROUTE_QUOTA_TELEMETRY.md

- Headings: OmniRoute Quota Telemetry; Truthful states
- Claim: OmniRoute separates provider quota telemetry from Ghostlight accounting. - `healthy` means a source reported usable remaining capacity. - `approaching_limit` means a source repo...

### app/docs/OMNIROUTE_ROUTING_POLICY.md

- Headings: OmniRoute Routing Policy
- Claim: Routing preserves the existing capability and combo selection logic, then applies allocation, health, circuit, quota, latency, reliability, model preference, and cost preference...

### app/docs/README.md

- Headings: OmniRoute Documentation; For Non-Tech Users; getting-started/; guides/; For Tech Users; architecture/; reference/; frameworks/
- Claim: --- title: "OmniRoute Documentation" version: 3.8.40 lastUpdated: 2026-06-28 --- Navigable index of the OmniRoute documentation set. Topics are grouped by intent so you can find...

### app/docs/architecture/ADAPTIVE_ROUTING.md

- Headings: Adaptive Routing: Routing Events, Quality Feedback & Explainability; 1. Architectural context; What was already there (audited, not duplicated); What was missing (added here); 2. Routing Events (feedback foundation); Measured overhead (honest comparison); 3. Quality Signal (feedback-driven provider state); Confidence / sample awareness
- Claim: --- title: "Adaptive Routing: Routing Events, Quality Feedback & Explainability" version: 3.8.50 lastUpdated: 2026-08-20 --- This document describes the feedback-driven adaptive...

### app/docs/architecture/ARCHITECTURE.md

- Headings: OmniRoute Architecture; Executive Summary; Reference Diagrams; Scope and Boundaries; In Scope; Out of Scope; Dashboard Surface (Current); High-Level System Context
- Claim: --- title: "OmniRoute Architecture" version: 3.8.40 lastUpdated: 2026-06-28 --- 🌐 **Languages:** 🇺🇸 [English](./ARCHITECTURE.md) | 🇪🇹 [አማርኛ](../i18n/am/docs/architecture/ARCHITE...

### app/docs/architecture/AUTHZ_GUIDE.md

- Headings: Authorization Guide; Two Auth Modes; 1. API Key (Bearer); 2. Dashboard Session (auth_token cookie); Route Classes; Pipeline; Policy contracts; Public Routes List
- Claim: --- title: "Authorization Guide" version: 3.8.40 lastUpdated: 2026-09-22 --- > **Source of truth:** `src/server/authz/`, `src/shared/constants/publicApiRoutes.ts`, `src/lib/api/...

### app/docs/architecture/CODEBASE_DOCUMENTATION.md

- Headings: OmniRoute Codebase Documentation; 1. Tech Stack; 2. Repository Layout; 3. `src/` — Next.js Application; 3.1 `src/app/` — App Router; 3.2 `src/lib/` — Core libraries; 3.3 `src/domain/` — Domain layer; 3.4 `src/server/` — Server-only
- Claim: --- title: "OmniRoute Codebase Documentation" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Version:** v3.8.51 > **Last updated:** 2026-06-28 > **Audience:** Engineers contrib...

### app/docs/architecture/DESIGN_SYSTEM.md

- Headings: OmniRoute — Design System & Visual Identity; 1. Purpose; 2. Principles; 3. Current state — what's already aligned vs. what's not; 3.1 Colors — already unified ✅; 3.2 Gaps — what the dashboard is missing; 3.3 Theming mechanics (so we don't break anything); 4. Part A — The graph-paper grid background (headline ask) — IMPLEMENTED (Phase 1)
- Claim: --- title: "Design System & Visual Identity" lastUpdated: 2026-07-11 --- > **Status:** reference — the standardization described here is **implemented** (phases 1–6: grid wallpa...

### app/docs/architecture/MONITORING_SECTIONS.md

- Headings: Monitoring & Costs — Navigation Structure; High-Level Navigation; Costs section (new, level 1); Monitoring section (reorganized); What changed from the old structure; Activity vs Audit Log; High-Level Actions allowlist; Adding a new high-level action
- Claim: --- title: "Monitoring & Costs — Navigation Structure" version: 3.8.40 lastUpdated: 2026-06-28 --- > Implemented in Group B (plan 16). See `src/shared/constants/sidebarVisibilit...

### app/docs/architecture/QUALITY_GATES.md

- Headings: Quality Gates Reference; Gate Inventory (~90 scripts); Release PR fast-path (`quality.yml`); Job: `lint`; Job: `quality-gate`; Job: `quality-extended`; Job: `docs-sync-strict`; Job: `i18n-ui-coverage`
- Claim: --- title: Quality Gates Reference --- This document is the authoritative reference for all CI quality gates in OmniRoute. It describes each gate, what it validates, which CI jo...

### app/docs/architecture/REPOSITORY_MAP.md

- Headings: Repository Map; Top-level tree; Root files; `src/` — Next.js application; `src/app/` — App Router (Next.js 16); `src/lib/` — Core libraries (~50 modules); `src/lib/db/` — Database (122 modules + 168 migrations); `src/domain/`
- Claim: --- title: "Repository Map" version: 3.8.40 lastUpdated: 2026-06-28 --- > **One-line description for every directory and root file.** > Last updated: 2026-06-28 — OmniRoute v3.8...

### app/docs/architecture/RESILIENCE_GUIDE.md

- Headings: Resilience Guide; 1. Provider Circuit Breaker; Opt-in global Provider Cooldown (window gate); 2. Connection Cooldown; Claude OAuth usage wall: lower-priority lane + session-limit reset; Session affinity (#7274); Exclusive managed session connection leases; 3. Model Lockout
- Claim: --- title: "Resilience Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- OmniRoute has three distinct but related resilience mechanisms. Each has a different scope and purpose....

### app/docs/architecture/ROUTER_BACKENDS.md

- Headings: Router Backends & Embedded Services — architecture contract (ADR); The core distinction — two orthogonal axes; The registry — single source of truth; Axis A — embedded services (supervised process side); Lifecycle route contract (`/api/services/<tool>/…`); Axis B — relay routing backend (dispatch side); Dashboard integration; Consequences
- Claim: --- title: "Router Backends & Embedded Services (ADR)" version: 3.8.43 lastUpdated: 2026-07-02 --- > **Status:** Accepted · **Context:** [#5670](https://github.com/diegosouzapw/...

### app/docs/architecture/admission-lanes.md

- Headings: Admission lanes (#9654) — two lane systems, what gates each, where each reports; 1. Byte-level process-wide admission (`chatBodyAdmission.ts`); 2. Adaptive runtime virtual lanes (`open-sse/services/admission`); 3. Fan-out probes — per-target admission for combo/fusion (#9654 Wave 2); Which one is showing in a dashboard; Why both exist; 4. One-process long `/v1/responses` (healthy-headroom)
- Claim: --- title: "Admission lanes — two lane systems, what gates each, where each reports" status: active lastUpdated: 2026-08-10 --- OmniRoute has **two** process-local lane systems ...

### app/docs/architecture/cluster-decisions.md

- Headings: Cluster Decisions — Optional Sidecar Profiles; TL;DR; Why this is conservative; The two profiles; `memory` — Qdrant Vector Memory Sidecar; `bifrost` — Bifrost Tier-1 Router Sidecar; What this PR explicitly does NOT do; 4-week rollout (if approved)
- Claim: --- title: "Cluster Decisions" version: 3.8.40 lastUpdated: 2026-06-28 --- **Status:** proposal (awaiting @diegosouzapw review) **Date:** 2026-06-20 **Refs:** [#3932](https://gi...

### app/docs/architecture/persistence-backend-boundary.md

- Headings: ADR: Pluggable persistence boundary; Context; Decision; Boundary rules; Portable repository surface; Backend capability surface; Transaction and migration model; Cross-backend conformance semantics
- Claim: --- title: "ADR: Pluggable persistence boundary" status: proposed lastUpdated: 2026-07-23 --- - **Status:** Proposed — requires maintainer approval before runtime work begins - ...

### app/docs/assets/flags/README.md

- Headings: Flag icons; Source & license
- Claim: --- title: "Flag icons" version: 3.8.50 lastUpdated: 2026-08-06 --- SVG country flags used by the language selector in the root `README.md`. They are served locally (instead of ...

### app/docs/changelog/fragments/10962.md

- Headings: (no heading)
- Claim: fix(catalog): expose only provider-routable GLM reasoning-effort tiers and remove unroutable ZCode aliases

### app/docs/changelog/fragments/14268.md

- Headings: (no heading)
- Claim: fix(cleanup): prune compression_engine_breakdown on retention schedule and usage resets (#14268)

### app/docs/changelog/fragments/15006.md

- Headings: (no heading)
- Claim: fix(providers): derive the new-connection default name from live connection names so adding after a delete no longer collides with (and overwrites) an existing connection (#15006)

### app/docs/changelog/fragments/7778.md

- Headings: (no heading)
- Claim: feat(resilience): atomically enforce cumulative global, provider, and account concurrency limits

### app/docs/comparison/OMNIROUTE_VS_ALTERNATIVES.md

- Headings: OmniRoute vs Alternatives; When to choose OmniRoute; When to choose LiteLLM; When to choose OpenRouter (SaaS); When to choose Portkey
- Claim: --- title: "OmniRoute vs Alternatives" version: 3.8.50 lastUpdated: 2026-08-02 --- Objective feature comparison vs popular open-source AI routers. > **Methodology**: Public repo...

### app/docs/compression/COMPRESSION_ENGINES.md

- Headings: Compression Engines; Modes; OmniGlyph compression profiles; Engine Registry; MCP description compression (related); Additional built-in engines; Caveman; RTK
- Claim: --- title: "Compression Engines" version: 3.8.40 lastUpdated: 2026-06-28 --- OmniRoute compression is built around engine contracts. A mode can run one engine directly (`caveman...

### app/docs/compression/COMPRESSION_GUIDE.md

- Headings: 🗜️ Prompt Compression Guide — OmniRoute; Overview; Compression Modes; Off; Lite Mode (~15% savings, <1ms latency); Standard Mode (~30% savings); Aggressive Mode (~50% savings); Ultra Mode (~75% savings)
- Claim: --- title: "🗜️ Prompt Compression Guide — OmniRoute" version: 3.8.40 lastUpdated: 2026-06-28 --- > Save 15-95% on eligible context automatically. For a quick overview, see the [...

### app/docs/compression/COMPRESSION_LANGUAGE_PACKS.md

- Headings: Compression Language Packs; Location; Language Detection; Config Shape; Adding a Language Pack; API; SHARED_BOUNDARIES (v3.8.0); Why this matters
- Claim: --- title: "Compression Language Packs" version: 3.8.40 lastUpdated: 2026-06-28 --- Caveman compression can load language-specific rule packs in addition to the built-in English...

### app/docs/compression/COMPRESSION_RULES_FORMAT.md

- Headings: Compression Rules Format; Caveman Rule Packs; Caveman Fields; RTK Filter Packs; RTK Fields; Safety Rules; Validation
- Claim: --- title: "Compression Rules Format" version: 3.8.40 lastUpdated: 2026-06-28 --- Compression rules are JSON files loaded at runtime. They are intentionally data-only so new lan...

### app/docs/compression/CONTEXT_EDITING.md

- Headings: Delegated Context Editing (Anthropic); What `clear_tool_uses` does; The per-combo enable toggle; Claude-only gating; The 400-fallback / relay coverage; `applied_edits` telemetry; Relationship to the local compression engines; See Also
- Claim: --- title: "Delegated Context Editing (Anthropic)" version: 3.8.40 lastUpdated: 2026-06-28 --- Delegated **Context Editing** is a Claude-only context-management feature. Unlike ...

### app/docs/compression/EXTENDING_COMPRESSION.md

- Headings: Extending the Compression Pipeline; Overview; Writing a Custom Compression Engine; The `CompressionEngine` Interface; Minimal Example: Whitespace Engine; Where to Place Custom Engines; Testing Your Engine; Creating Language Packs
- Claim: --- title: "Extending the Compression Pipeline" version: 3.8.44 lastUpdated: 2026-07-02 --- > **TL;DR**: OmniRoute's compression engine is **pluggable** — you can register custo...

### app/docs/compression/RTK_COMPRESSION.md

- Headings: RTK Compression; What It Compresses; Filter Resolution; RTK TOML schema v1 compatibility; Filter DSL; Line Deduplication (two layers); Line Grouping (`enableGrouping`); Code Comment Stripping (`stripCodeComments` / `preserveDocstrings`)
- Claim: --- title: "RTK Compression" version: 3.8.40 lastUpdated: 2026-06-28 --- RTK compression is OmniRoute's command-aware compression engine for terminal and tool output. It is desi...

### app/docs/diagrams/README.md

- Headings: Diagrams; Canonical diagrams; Hand-authored animated diagrams; How to update; Linking from a doc; Conventions
- Claim: --- title: "Diagrams" version: 3.8.49 lastUpdated: 2026-07-17 --- Mermaid sources (`.mmd`) and exported SVGs for OmniRoute v3.8.0 architecture flows. Not every diagram comes fro...

### app/docs/frameworks/A2A-SERVER.md

- Headings: OmniRoute A2A Server Documentation; Agent Discovery; Authentication; Enablement; JSON-RPC 2.0 Methods; `message/send` — Synchronous Execution; `message/stream` — SSE Streaming; `tasks/get` — Query Task Status
- Claim: --- title: "OmniRoute A2A Server Documentation" version: 3.8.40 lastUpdated: 2026-06-28 --- > Agent-to-Agent Protocol v0.3 — OmniRoute as an intelligent routing agent The A2A su...

### app/docs/frameworks/ACP.md

- Headings: ACP (Agent Client Protocol); What Is ACP?; Why Use ACP?; Supported CLI Agents; Custom Agents; Quick Start; Step 1: Install a CLI Agent; Example: Install Claude Code CLI
- Claim: --- title: ACP (Agent Client Protocol) --- > **TL;DR**: ACP lets OmniRoute spawn CLI agents (like Claude Code, Codex) as child processes instead of using HTTP APIs. This gives y...

### app/docs/frameworks/AGENT-SKILLS.md

- Headings: OmniRoute Agent Skills Catalog; Overview; Architecture; SKILL.md Format; Overview; Authentication; Endpoints; Custom Section (preserved across regeneration)
- Claim: --- title: "OmniRoute Agent Skills Catalog" version: 3.8.50 lastUpdated: 2026-08-02 --- > **Source of truth:** `src/lib/agentSkills/` (catalog, generator, parsers) + `skills/` d...

### app/docs/frameworks/AGENTBRIDGE.md

- Headings: AgentBridge; §1 Overview; What is AgentBridge?; Positioning vs. the market; §2 Architecture; 2.1 Components overview; 2.2 MITM server (`src/mitm/server.cjs`); 2.3 Handler base (`src/mitm/handlers/base.ts`)
- Claim: --- title: "AgentBridge" version: 3.8.40 lastUpdated: 2026-06-28 --- AgentBridge is OmniRoute's MITM (Man-in-the-Middle) proxy that intercepts HTTPS traffic from IDE AI agents a...

### app/docs/frameworks/AGENT_PROTOCOLS_GUIDE.md

- Headings: Agent Protocols Guide; TL;DR; Decision Tree; 1. A2A — Agent-to-Agent; When to use; Methods; Built-in skills (6); Deep dive
- Claim: --- title: "Agent Protocols Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source:** `src/lib/{a2a,acp,cloudAgent}/`, `src/app/api/{a2a,acp,cloud}/`, `src/app/api/v1/age...

### app/docs/frameworks/CLOUD_AGENT.md

- Headings: Cloud Agents; Supported Agents; Architecture; `CloudAgentBase` Interface; Domain Types; Database; REST API — Task Lifecycle; Create task
- Claim: --- title: "Cloud Agents" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/cloudAgent/` and `src/app/api/v1/agents/tasks/` > **Last updated:** 2026-06...

### app/docs/frameworks/EMBEDDED-SERVICES.md

- Headings: Embedded Services; 1. Overview; Why embedded services?; Key decisions (from design plan); 2. Architecture — 4 layers; Key source files; 3. Lifecycle state machine; State transitions
- Claim: --- title: "Embedded Services" description: "Reference for 9Router, CLIProxyAPI, Mux, Bifrost, and open-wa" --- > **Version:** v3.8.44 > **Last updated:** 2026-09-09 > **Audienc...

### app/docs/frameworks/EVALS.md

- Headings: Evaluations (Evals); Concepts; Suite; Case; Target; Scoring Rubrics; Database Schema; REST API
- Claim: --- title: "Evaluations (Evals)" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/evals/`, `src/lib/db/evals.ts`, `src/app/api/evals/` > **Last update...

### app/docs/frameworks/GAMIFICATION.md

- Headings: Gamification & Leaderboard System; Overview; Purpose; Scope; Design Principles; Architecture; High-Level Flow; Module Dependency Graph
- Claim: --- title: "Gamification & Leaderboard System" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/gamification/`, `src/lib/db/gamification.ts`, `src/app...

### app/docs/frameworks/LOCAL_CORPUS_CONTEXT.md

- Headings: Local Corpus Context Source; What it is; Configure the source; Connect an approved directory; Check configuration and index status; Disconnect without changing source files; MCP tools; Safety boundaries
- Claim: --- title: "Local Corpus Context Source" version: 3.8.49 lastUpdated: 2026-07-20 --- > **Source of truth:** `src/lib/localCorpus/index.ts` (bounded text index), > `src/lib/local...

### app/docs/frameworks/LOG-EXPORT.md

- Headings: Log export; 1. How it works; Payloads (prompts and completions); 2. Files; 3. REST API; 4. BigQuery destination; 5. Adding a destination; 6. Operating it
- Claim: --- title: "Log Export" version: 3.8.51 lastUpdated: 2026-08-29 --- Continuous, incremental export of OmniRoute call logs to an external analytics store. The Logs dashboard tab ...

### app/docs/frameworks/MCP-SERVER.md

- Headings: OmniRoute MCP Server Documentation; Installation; HTTP streamable transport (port 20130); Transports; Remote access (manage-scope bypass); Grant manage scope: open the dashboard API Keys page and toggle; "Management Access" on the key, or POST scopes:["manage"] when creating.; Then connect from a remote MCP client:
- Claim: --- title: "OmniRoute MCP Server Documentation" version: 3.8.51 lastUpdated: 2026-09-22 --- > Model Context Protocol server with 110 tools across routing, cache, compression, me...

### app/docs/frameworks/MEMORY.md

- Headings: Memory System; Architecture; Engine architecture (3-tier resolution); Embedding sources; Hybrid RRF (k=60); Backfill (lazy + reindex); Settings extension; Storage Layers
- Claim: --- title: "Memory System" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/memory/` and `src/app/api/memory/` > **Last updated:** 2026-06-28 — v3.8.4...

### app/docs/frameworks/NOTION_CONTEXT.md

- Headings: Notion Context Source; What it is; Setup; Configure via REST; Save + validate the integration token (POST validates by issuing a test search); Check connection status; Disconnect (clears the stored token); MCP tools (6)
- Claim: --- title: "Notion Context Source" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/notion/api.ts` (REST client), `src/lib/db/notion.ts` > (token pers...

### app/docs/frameworks/OBSIDIAN_CONTEXT.md

- Headings: Obsidian Context Source; What it is; Setup; Configuration keys (SQLite `key_value`, namespace `obsidian`); Configure via REST; Save + validate the Local REST API token (POST validates via a status check); Check connection status (returns connected, hasToken, baseUrl, vaultPath); Disconnect (clears the stored token)
- Claim: --- title: "Obsidian Context Source" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/obsidian/api.ts` (REST + sync client), > `src/lib/db/obsidian.ts...

### app/docs/frameworks/OPENCODE.md

- Headings: OpenCode Integration; Path 1 — CLI generator (no npm install); After installing OmniRoute (npm i -g @omniroute/cli or local clone); Path 2 — npm package `@omniroute/opencode-provider`; What the runtime actually does; Model catalog defaults; URL normalisation; Authentication modes
- Claim: --- title: "OpenCode Integration" version: 3.8.40 lastUpdated: 2027-07-27 --- > **Status:** Generally available. > **Audience:** Operators wiring OpenCode to an OmniRoute deploy...

### app/docs/frameworks/OPEN_SSE_ARCHITECTURE.md

- Headings: open-sse Architecture; Why a Separate Workspace Package?; Top-Level Structure; Module Counts; The Request Pipeline; Stage 1: Route (services/combo.ts); Stage 2: Translate (translator/); Stage 3: Execute (executors/)
- Claim: --- title: "open-sse Architecture" version: 3.8.40 lastUpdated: 2026-06-28 --- > **TL;DR**: `open-sse/` is the core streaming engine that powers every LLM request in OmniRoute. ...

### app/docs/frameworks/PLAYGROUND_STUDIO.md

- Headings: Playground Studio; Overview; Tabs; Chat Tab; Compare Tab; API Tab; Build Tab; Config Pane (Shared)
- Claim: --- title: "Playground Studio" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Feature:** Playground Studio — unified AI testing workspace for `/dashboard/playground`. > **Plans...

### app/docs/frameworks/PLUGINS.md

- Headings: OmniRoute CLI Plugin System; Quick start; Install a plugin from npm; Install a local plugin in development; List installed plugins; Scaffold a new plugin; Plugin anatomy; `package.json`
- Claim: --- title: "OmniRoute CLI Plugin System" version: 3.8.40 lastUpdated: 2026-06-28 --- Extend the `omniroute` CLI without modifying its core. Plugins follow the `omniroute-cmd-*` ...

### app/docs/frameworks/PLUGIN_MARKETPLACE.md

- Headings: Plugin Marketplace; How It Fits Together; Marketplace Catalog; REST API; List installed plugins; Install from a local path; Browse the marketplace; Update plugin config
- Claim: --- title: "Plugin Marketplace" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/plugins/` (`marketplace.ts`, `manager.ts`, `manifest.ts`, > `scanner....

### app/docs/frameworks/PLUGIN_SDK.md

- Headings: OmniRoute Plugin SDK; Quick Start; API Reference; `definePlugin(def: PluginDefinition): Plugin`; `blockRequest(response?): BlockingHookResult`; `modifyBody(body): PluginResult`; `addMetadata(metadata): PluginResult`; Plugin Context (`PluginContext`)
- Claim: --- title: "OmniRoute Plugin SDK" version: 3.8.40 lastUpdated: 2026-06-28 --- import { definePlugin } from "omniroute/plugins/sdk"; export default definePlugin({ name: "my-plugi...

### app/docs/frameworks/RADAR.md

- Headings: Radar Free-Model Catalog; Delivery status in v3.8.51; Public announcement reader; Flag: `RADAR_ENABLED` (default off); Data sync is a SEPARATE opt-in — the privacy promise; Access and safety rules shown before opt-in; Getting a supporter key; Recovering a lost supporter key
- Claim: --- title: "Radar Free-Model Catalog" version: 3.8.51 lastUpdated: 2026-09-01 --- > **Source of truth:** `src/lib/radar/`, `src/lib/db/radar.ts`, `src/app/api/radar/` > **Last u...

### app/docs/frameworks/SEARCH_TOOLS_STUDIO.md

- Headings: Search Tools Studio; Overview; Tabs; Search Tab; Scrape Tab; Compare Tab; Config Pane (Shared); SearchConceptCard
- Claim: --- title: "Search Tools Studio" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Feature:** Search Tools Studio — unified web tools workspace for `/dashboard/search-tools`. > **...

### app/docs/frameworks/SKILLS.md

- Headings: Skills Framework; Agent Skills vs Omni Skills; Concepts; Skill Sources; Skill Identity; Skill Mode; Status (executions); Registry Cache
- Claim: --- title: "Skills Framework" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/skills/` and `src/app/api/skills/` > **Last updated:** 2026-06-28 — v3....

### app/docs/frameworks/TRAFFIC_INSPECTOR.md

- Headings: Traffic Inspector; §1 Overview; What makes Traffic Inspector unique; Architecture in one paragraph; §2 Capture modes; Mode 1 — AgentBridge (default, always on); Mode 2 — Custom Hosts (DNS redirect); Mode 3 — HTTP_PROXY listener (port 8080)
- Claim: --- title: "Traffic Inspector" version: 3.8.40 lastUpdated: 2026-06-28 --- Traffic Inspector is OmniRoute's built-in HTTPS traffic debugger — a Charles Proxy / mitmweb / HTTP To...

### app/docs/frameworks/WEBHOOKS.md

- Headings: Webhooks; Supported Events; Architecture; HMAC Signing; Verifying on the receiver; Retry & Failure Policy; Database; REST API
- Claim: --- title: "Webhooks" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/webhookDispatcher.ts`, `src/lib/db/webhooks.ts`, `src/app/api/webhooks/` > **La...

### app/docs/getting-started/AUTO-COMBO-GUIDE.md

- Headings: Auto-Combo: Let OmniRoute Pick the Best AI for You; What It Does; Quick Start; Which "auto" Should I Use?; Examples; General chat — balanced; Code generation — quality-first; Quick answer — speed-first
- Claim: --- title: "Auto-Combo: Let OmniRoute Pick the Best AI for You" version: 3.8.50 lastUpdated: 2026-08-06 --- > **TL;DR**: Set your model to `auto` and OmniRoute automatically pic...

### app/docs/getting-started/FREE-TIERS-GUIDE.md

- Headings: Free Tiers Guide: Understand and Combine Free AI Access; What Are Free Tiers?; Representative Free-Access Providers; Recurring, Keyless, or Uncapped Access; Signup Grants and Provider-Specific Credits; Other Limited Access; How to Stack Free Tiers; Example: Broader Free-Tier Coverage
- Claim: > **TL;DR**: OmniRoute registers 357 provider IDs, with **152 provider-catalog entries marked `hasFree`**. The stricter audited free-model catalog covers **35 recurring pool key...

### app/docs/getting-started/PROVIDERS-GUIDE.md

- Headings: Providers Guide: Connect AI Models to OmniRoute; What Is a Provider?; Types of Providers; Web Cookie Providers; Quick Start: Connect Your First Provider; Optional first-run free-provider setup; Option A: Free Provider (No Credit Card); Option B: API Key Provider (Paid)
- Claim: --- title: "Providers Guide: Connect AI Models to OmniRoute" version: 3.8.50 lastUpdated: 2026-08-06 --- > **TL;DR**: A provider is a connection to an AI service (like OpenAI, A...

### app/docs/getting-started/QUICK-START.md

- Headings: Quick Start: Get OmniRoute Running in 3 Minutes; Step 1: Install OmniRoute; Option A: npm (Recommended); Option B: Docker; Option C: From Source; Step 2: Start OmniRoute; Step 3: Connect a Free Provider; Option A: Kiro (Free Claude — No Credit Card)
- Claim: --- title: "Quick Start: Get OmniRoute Running in 3 Minutes" version: 3.8.50 lastUpdated: 2026-08-06 --- > **TL;DR**: Install → Connect a free provider → Point your IDE to OmniR...

### app/docs/getting-started/SELF_HOST_GUIDE.md

- Headings: 🚀 Self-Host Guide — OmniRoute; Why this guide exists; Prerequisites; Step 1 — Configure (1 min); Step 2 — Start (1 min); Step 3 — Verify (30 s); process lifecycle + readiness; container health
- Claim: --- title: "🚀 Self-Host Guide — OmniRoute (零月费自托管 / zero-fee self-host)" version: 3.8.51 lastUpdated: 2026-09-27 --- > **TL;DR** — three commands, one local endpoint, zero month...

### app/docs/getting-started/WEB-COOKIE-GUIDE.md

- Headings: Web Cookie Providers; Before You Begin; Do NOT copy from Cookie Storage; Copy from a Live Request; General Setup; Provider Credential Formats; What Web Cookie Providers Can and Cannot Do; Supported
- Claim: --- title: "Getting Started — Web Cookie Providers" version: 3.8.40 lastUpdated: 2026-07-20 --- Web Cookie providers let OmniRoute use an AI service through your existing browse...

### app/docs/guides/ANTIGRAVITY-ONBOARDING.md

- Headings: OmniRoute Antigravity (Google One AI) Onboarding Guide; 1. Antigravity vs Antigravity CLI (agy); 2. Google One AI Pro: Quota System; The Two Layers; Plan Tiers; Gemini vs Non-Gemini Models; AI Credits (Overage); Key Details
- Claim: --- title: "Antigravity (Google One AI) — Onboarding with OmniRoute" version: 3.8.50 lastUpdated: 2026-07-31 --- > **What you get**: Access to Gemini 3.1 Pro, Gemini 3.7 Flash, ...

### app/docs/guides/CHAOS-MODE.md

- Headings: Chaos Mode; `auto/chaos`: parallel fan-out; Setup; Execution modes; API; `POST /api/chaos/run` — dashboard session; `POST /api/skills/collect/chaos` — API key; Troubleshooting
- Claim: --- title: "Chaos Mode" version: 3.8.51 lastUpdated: 2026-09-01 --- > **Dashboard:** **Chaos Mode** (sidebar) → `/dashboard/chaos` > **API:** `GET` / `PUT` `/api/chaos/config` ·...

### app/docs/guides/CLAUDE-CODE-CONFIGURATION.md

- Headings: Claude Code CLI — Configuration with OmniRoute; Quick start; Launch Claude Code against a local OmniRoute (auto-detects the active context); Against a remote OmniRoute (after `omniroute connect <host>`, this is automatic); Generate per-model profiles, then launch one; How Claude Code connects to a gateway; Discovery aliases — surface non-Claude models in the `/model` picker; Onboarding block on the dashboard
- Claim: --- title: "Claude Code CLI — Configuration with OmniRoute" version: 3.8.40 lastUpdated: 2026-07-24 --- Point the **Claude Code** CLI (`claude`) at OmniRoute — local or a remote...

### app/docs/guides/CLI-COMPANION.md

- Headings: CLI Companion; Server detection is not browser detection; Local or remote contexts; Implementation boundary
- Claim: --- title: "CLI Companion — copy-only terminal guidance" version: 3.8.51 lastUpdated: 2026-09-22 --- The CLI Code dashboard includes a copy-only companion for the executable CLI...

### app/docs/guides/CLI-INTEGRATIONS.md

- Headings: CLI Integrations; Master table; Pick from the active local or remote model catalog and configure the target.; Local usage; Codex: write a profile per matched model into ~/.codex/; Claude Code: write per-model profiles, then launch one; OpenCode: write the openai-compatible provider with all catalog models; Tools without auto-discovery need an explicit model:
- Claim: --- title: "CLI Integrations — point any coding CLI at OmniRoute" version: 3.8.50 lastUpdated: 2026-08-18 --- OmniRoute ships a family of `setup-*` commands that configure a cod...

### app/docs/guides/CLI-MANUAL-MODELS.md

- Headings: CLI model catalog and manual metadata; Manual model lifecycle; Test & Add: not implemented
- Claim: --- title: "CLI model catalog and manual metadata" version: "3.8.51" lastUpdated: "2026-09-21" --- `omniroute models` reads the selected server's public model catalog first, inc...

### app/docs/guides/CLI-SMOKE-TESTS.md

- Headings: Isolated CLI smoke tests
- Claim: --- title: Isolated CLI smoke tests description: Validate installed CLI protocol contracts with disposable local responses. version: "3.8.51" lastUpdated: "2026-09-22" --- Run t...

### app/docs/guides/CODEX-APP-SERVER-PROVIDER.md

- Headings: OpenAI Codex — App-Server provider (`codex-app-server`); 1. Architecture; 2. Bring it up; Start the stack WITH the codex app-server sidecar profile:; (podman: podman compose --profile base --profile codex-app-server up -d); 3. Connect + sign in; 4. Deployment scenarios; 5. Residential / UDP egress (operator extra, not shipped)
- Claim: --- title: "OpenAI Codex (App-Server) provider" version: 3.8.50 lastUpdated: 2026-08-22 --- OmniRoute exposes **two** ways to use OpenAI Codex: Because `codex-app-server` never ...

### app/docs/guides/CODEX-CLI-CONFIGURATION.md

- Headings: Codex CLI — Configuration with OmniRoute; Ready-to-paste config.toml; ~/.codex/config.toml; ~/.bashrc or ~/.zshrc — actual key value, never in config.toml; macOS: Codex bundled inside the ChatGPT app; Local unauthenticated OmniRoute: placeholder key is enough; `wire_api = "responses"` — why it works for all models; Context window and compaction
- Claim: --- title: "Codex CLI — Configuration with OmniRoute" version: 3.8.50 lastUpdated: 2026-08-18 --- Complete guide for using the Codex CLI pointed at OmniRoute as an OpenAI-compat...

### app/docs/guides/COST_TRACKING.md

- Headings: Cost & Spend Tracking; What it is (and what it is not); How costs are estimated; The pricing source; The cost formula; How spend is recorded; Dashboard: the Costs page; Related Costs sub-pages
- Claim: --- title: "Cost & Spend Tracking" version: 3.8.40 lastUpdated: 2026-06-28 --- How OmniRoute estimates, records, and reports the cost of every request — and why the dashboard nu...

### app/docs/guides/DOCKER_GUIDE.md

- Headings: 🐳 Docker Guide — OmniRoute; Quick Run; With Environment File; Copy and edit .env first; Docker Compose; Base profile (no CLI tools); CLI profile (Claude Code, Codex, OpenClaw built-in); Host profile (Linux-first; mounts host CLI binaries read-only)
- Claim: --- title: "🐳 Docker Guide — OmniRoute" version: 3.8.51 lastUpdated: 2026-09-18 --- > Complete Docker deployment reference. For a quick start, see the [README Docker section](.....

### app/docs/guides/ELECTRON_GUIDE.md

- Headings: Electron Desktop Guide; Architecture; Versions; Scripts (root `package.json`); Directory Layout; IPC Bridge (`preload.js`); Server Lifecycle; Zero-config Secret Bootstrap
- Claim: --- title: "Electron Desktop Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `electron/` workspace > **Last updated:** 2026-06-28 — v3.8.40 OmniRoute s...

### app/docs/guides/FEATURES.md

- Headings: OmniRoute — Dashboard Features Gallery; ✨ v3.8.0 Highlights; 🔌 Providers; 🎨 Combos; 📊 Analytics; 🏥 System Health; 🔧 Translator Playground; 🎮 Model Playground _(v2.0.9+)_
- Claim: --- title: "OmniRoute — Dashboard Features Gallery" version: 3.8.40 lastUpdated: 2026-06-28 --- 🌐 **Languages:** 🇺🇸 [English](./FEATURES.md) | 🇪🇹 [አማርኛ](../i18n/am/docs/guides/F...

### app/docs/guides/FREE_PROVIDER_RANKINGS.md

- Headings: Free Provider Rankings (Arena ELO); What It Is; How to Access; Dashboard page; API endpoint; How the Scores Work; Source: Arena AI leaderboard; Normalization (task-fit score)
- Claim: --- title: "Free Provider Rankings (Arena ELO)" version: 3.8.50 lastUpdated: 2026-08-02 --- > **TL;DR**: OmniRoute ranks its **free** providers by model quality using **Arena AI...

### app/docs/guides/I18N.md

- Headings: i18n — Internationalization Guide; Translation pipeline (recommended — v3.8.0); Run translations (incremental — only touches changed sources); Limit to one locale; Specific files (comma-separated, repo-relative paths); Force retranslate everything (expensive); Preview what would happen (no API calls, no writes); CI gate — exits non-zero if state is drifting (blocking in CI since 2026-09)
- Claim: --- title: "i18n — Internationalization Guide" version: 3.8.40 lastUpdated: 2026-09-02 --- OmniRoute supports **67 languages** with full dashboard UI translation, translated doc...

### app/docs/guides/KIRO_SETUP.md

- Headings: Kiro Setup Guide; Background: Why Kiro Accounts Can Conflict; How OmniRoute Solves This (v3.8.0+); Migration Note for Connections Created Before v3.8.0; Adding Two Kiro Accounts Side by Side; Prerequisites; Step 1: Import the first account; Step 2: Import the second account
- Claim: --- title: "Kiro Setup Guide" --- This guide covers adding Kiro (AWS-hosted AI coding assistant) accounts to OmniRoute, with a focus on running multiple accounts simultaneously ...

### app/docs/guides/LOCAL_OLLAMA_DOCKER.md

- Headings: Local Ollama + Docker OmniRoute; Architecture; Quick start; Verify Docker → Ollama connectivity; → 172.17.0.1  host.docker.internal; Troubleshooting: `Direct response did not start within 30000ms`; Root cause; Fix (v3.8.51+)
- Claim: --- title: "Local Ollama + Docker OmniRoute" version: 3.8.51 lastUpdated: 2026-09-29 --- Guide for running **OmniRoute in Docker** while **Ollama runs natively on the host** — t...

### app/docs/guides/MANAGEMENT-AUTH.md

- Headings: Management Authentication; Scope matrices; Access Token scopes (`oma_live_…`); API-key management scopes; How to create and revoke; Dashboard JWT session; CLI machine-id token; Scoped Access Token (`oma_live_…`)
- Claim: --- title: "Management Authentication" version: 3.8.50 lastUpdated: 2026-09-22 --- OmniRoute has **four credential families** that can authorize management routes. They are not ...

### app/docs/guides/OPENCODE-V2-PLUGIN.md

- Headings: OpenCode v2 plugin — install and configure; Requirements; Install; Credentials; The management token is a different key; Options; How the catalog stays fresh; Tool calling on Gemini
- Claim: --- title: "OpenCode v2 plugin — install and configure" version: 3.8.51 lastUpdated: 2026-09-06 --- `@omniroute/opencode-plugin-v2` puts your whole OmniRoute catalog — models, c...

### app/docs/guides/PWA_GUIDE.md

- Headings: Progressive Web App (PWA) Guide; What Is a PWA?; Installation; Android (Chrome); iOS (Safari); Desktop (Chrome / Edge); Features; Fullscreen Experience
- Claim: --- title: "Progressive Web App (PWA) Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- OmniRoute ships as a fully installable Progressive Web App. When you access the dashboar...

### app/docs/guides/REMOTE-MODE.md

- Headings: Remote Mode; How it works; Connecting; With the management password (bootstrap); Management password for http://192.168.0.15:20128: ********; ✔ Connected to http://192.168.0.15:20128 — context '192.168.0.15' (scope: admin); With a pre-generated token; Scopes
- Claim: --- title: "Remote Mode — Drive a remote OmniRoute from your laptop" version: 3.8.50 lastUpdated: 2026-08-18 --- Run the `omniroute` CLI on your laptop while OmniRoute itself ru...

### app/docs/guides/SETUP_GUIDE.md

- Headings: 📖 Setup Guide — OmniRoute; Install Methods; npm (recommended); pnpm; Arch Linux (AUR); From Source; Docker; Desktop App (Electron)
- Claim: --- title: "📖 Setup Guide — OmniRoute" version: 3.8.50 lastUpdated: 2026-08-18 --- > Complete setup reference for OmniRoute. For the quick version, see the [Quick Start in READM...

### app/docs/guides/TERMUX_GUIDE.md

- Headings: Termux Headless Setup; Prerequisites; Install; Run; Background Execution; Access From Other Devices; Data Directory; Limitations
- Claim: --- title: "Termux Headless Setup" version: 3.8.49 lastUpdated: 2026-07-25 --- OmniRoute can run as a headless server on Android through Termux. The Electron desktop app is not ...

### app/docs/guides/THINKING_BUDGET.md

- Headings: Thinking Budget; Modes; What `auto` removes; What this is **not**; Visible thinking (Codex / Responses clients); API examples; Read; Recommended for Codex / Desktop thinking visibility
- Claim: --- title: "Thinking Budget" version: 3.8.49 lastUpdated: 2026-08-12 --- > **Dashboard:** Settings → **AI** → Thinking Budget > **API:** `GET` / `PUT` `/api/settings/thinking-bu...

### app/docs/guides/TIERS.md

- Headings: OmniRoute Tiers — User Guide; Tier 1 — Subscription; Tier 2 — Cheap; Tier 3 — Free; Configuring tiers; Telemetry; Common patterns; Pure-free workload
- Claim: --- title: "OmniRoute Tiers — User Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- OmniRoute organizes the 352 supported providers into 3 economic tiers. Each request travels...

### app/docs/guides/TROUBLESHOOTING.md

- Headings: Troubleshooting; Quick Reference; Detailed Troubleshooting; Rate Limiting on Free Providers (429 / 400 / 401); npm install Warnings (ERESOLVE / peer / deprecated); Gemini Web and Playwright Chromium; Quick Fixes; Antivirus False Positives
- Claim: --- title: "Troubleshooting" version: 3.8.49 lastUpdated: 2026-07-15 --- > **For Users**: Looking for quick fixes? See the [Quick Reference](#quick-reference) below. 🌐 **Languag...

### app/docs/guides/UNINSTALL.md

- Headings: OmniRoute — Uninstall Guide; Quick Uninstall (v3.6.2+); Keep Your Data; Full Removal; Manual Uninstall; NPM Global Install; Remove the global package; (Optional) Remove data directory
- Claim: --- title: "OmniRoute — Uninstall Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- 🌐 **Languages:** 🇺🇸 [English](./UNINSTALL.md) | 🇪🇹 [አማርኛ](../i18n/am/docs/guides/UNINSTALL.m...

### app/docs/guides/USAGE_QUOTA_GUIDE.md

- Headings: Usage, Quota & Spend Tracking; Overview; What Gets Recorded; Where Tokens Come From; Cached Tokens; Cost Calculation; Pricing Sync; Manual trigger
- Claim: --- title: "Usage, Quota & Spend Tracking" version: 3.8.40 lastUpdated: 2026-06-28 --- > **TL;DR**: OmniRoute tracks every request's token usage, computes cost, enforces per-API...

### app/docs/guides/USER_GUIDE.md

- Headings: User Guide; 💰 Pricing at a Glance; 🎯 Use Cases; Case 1: "I have Claude Pro subscription"; Case 2: "I want zero cost"; Case 3: "I need 24/7 coding, no interruptions"; Case 4: "I want FREE AI in OpenClaw"; 📖 Provider Setup
- Claim: --- title: "User Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- 🌐 **Languages:** 🇺🇸 [English](./USER_GUIDE.md) | 🇪🇹 [አማርኛ](../i18n/am/docs/guides/USER_GUIDE.md) | 🇸🇦 [العربي...

### app/docs/guides/VSCODE-COPILOT.md

- Headings: VS Code Copilot Chat — OmniCopilot extension; Setup; What the picker will show; It asks for one id per model; It hides models that cannot chat; Providers you never configured; Dashboard inside a VS Code tab; the variable has to be present on the BUILD command
- Claim: --- title: "VS Code Copilot Chat — OmniCopilot extension" version: 3.8.50 lastUpdated: 2026-08-18 --- **OmniCopilot** puts every model your OmniRoute serves into the _native_ Gi...

### app/docs/ops/BRANCHING_MODEL.md

- Headings: Branching & Release Model; At a glance; Where should my PR target?; Release freeze (parallel cycles); Why both a branch and a tag?; Related docs
- Claim: --- title: "Branching & Release Model" lastUpdated: 2026-07-22 --- OmniRoute uses a **parallel-cycle** release model: a dedicated `release/vX.Y.Z` branch for the active cycle, `...

### app/docs/ops/BRANCH_PROTECTION_MAIN.md

- Headings: Branch protection — `main` (OpenSSF Scorecard: Branch-Protection)
- Claim: --- title: "Branch Protection — main" --- Owner action. Apply via Settings → Branches → Add rule, or: gh api -X PUT repos/diegosouzapw/OmniRoute/branches/main/protection \ --inp...

### app/docs/ops/CONTRIBUTION_GOLDEN_PATH.md

- Headings: Contribution Golden Path; The path every change follows; Golden paths by change type; Provider; Routing; UI / UX; i18n; CLI
- Claim: --- title: "Contribution Golden Path" --- Use this guide to choose the smallest reliable development loop for a pull request. It does not replace the area-specific architecture ...

### app/docs/ops/COVERAGE_PLAN.md

- Headings: Test Coverage Plan; Baseline; Rules; Current command set; Milestones; Priority hotspots; Execution checklist; Phase 1: 56.95% -> 60%
- Claim: --- title: "Test Coverage Plan" version: 3.8.40 lastUpdated: 2026-06-28 --- Last updated: 2026-06-28 > Status measured on 2026-05-13: lines 82.58%, statements 82.58%, functions ...

### app/docs/ops/DATABASE_GUIDE.md

- Headings: Database Schema & Operations Guide; Why SQLite?; WAL Journaling; Database Location; Domain Module Architecture; The 110 Top-Level DB Modules; Module Boundaries; Base Schema (17 tables)
- Claim: --- title: "Database Schema & Operations Guide" version: 3.8.50 lastUpdated: 2026-08-23 --- > **TL;DR**: OmniRoute uses **SQLite with WAL journaling** as its primary store, with...

### app/docs/ops/FLY_IO_DEPLOYMENT_GUIDE.md

- Headings: OmniRoute Fly.io Deployment Guide; 1. Deployment Goals; 2. Current Project Key Configuration; 3. Prerequisites; 3.1 Installing the Fly CLI; 3.2 Logging in to Your Fly Account; 3.3 Verifying Login Status; 4. First-Time Deployment of the Current Project
- Claim: --- title: "OmniRoute Fly.io Deployment Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- This document describes the actual deployment process for OmniRoute on Fly.io, coverin...

### app/docs/ops/HOMOLOGATION.md

- Headings: Homologation Suite (`npm run homolog`); What it covers; Prerequisites; How to run; Reading the report; Re-baselining when the UI changes legitimately; Roadmap (F2 / F3)
- Claim: --- title: "Homologation Suite (npm run homolog)" version: 3.8.49 lastUpdated: 2026-07-14 --- Real-environment E2E validation of the OmniRoute deploy running on the homologation...

### app/docs/ops/MERGE_TRAIN.md

- Headings: Merge Queue & Manual Merge-Train Runbook; Default path: the Mergify queue; Fallback: the manual merge-train; Tiering (why the queue is safe with fast-gates only)
- Claim: --- title: Merge Queue & Manual Merge-Train Runbook --- Since v3.8.49 (WS3.2/WS3.4 of the quality/velocity plan) the default merge path for reviewed PRs into `release/vX.Y.Z` is...

### app/docs/ops/MONITORING_GUIDE.md

- Headings: Monitoring & Observability Guide; Overview; Dashboard Pages; `/dashboard/health` (System Health); `/dashboard/providers` (Provider Health); `/dashboard/quota` (Quota Tracking); `/dashboard/combos` (Combo Health); Health Check API
- Claim: --- title: "Monitoring & Observability Guide" version: 3.8.50 lastUpdated: 2026-08-13 --- > **TL;DR**: OmniRoute ships with built-in health monitoring, provider autopilot, quota...

### app/docs/ops/PROXY_GUIDE.md

- Headings: 🌐 OmniRoute Proxy Guide; Why Use Proxies?; Architecture Overview; Key Components; 4-Level Proxy System; How Resolution Works; What Gets Proxied; Proxy Registry (CRUD)
- Claim: --- title: "🌐 OmniRoute Proxy Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Bypass geographic blocks, protect your identity, and route AI traffic through any proxy — wi...

### app/docs/ops/QUALITY_GATE_PLAYBOOK.md

- Headings: Quality-Gate System — Critical Assessment, Catalog and Replication Playbook; Part 1 — Verdict and Maturity Classification; Part 2 — Critical Assessment (strengths + honest weaknesses); Strengths (what's above average); Honest weaknesses (real gaps); Part 3 — Complete Catalog of Quality Checkpoints (portable); 1. Style & formatting (deterministic, fast); 2. Types
- Claim: --- title: "Quality Gate Playbook" --- > **What this document is.** A critical assessment of OmniRoute's quality-gate system, > compared to industry best practices, **plus** a c...

### app/docs/ops/REDIS_PRODUCTION_CONFIG.md

- Headings: Redis Production Configuration Guide; Overview; Current Configuration (Code Defaults); Key Namespacing; .env — all OmniRoute keys become omniroute:rl:*, omniroute:auth:*, omniroute:quota:*, omniroute:warmup:cb:*; Recommended Production Tuning; 1. Connection Pool / Client Options (ioredis `Redis` constructor); 2. Redis Server Configuration (`redis.conf`)
- Claim: --- title: "Redis Production Configuration Guide" version: 3.8.50 lastUpdated: 2026-08-06 --- Redis is an **optional, soft dependency** in OmniRoute — the application degrades g...

### app/docs/ops/RELEASE_CHECKLIST.md

- Headings: Release Checklist; TL;DR; 1. Bump version + generate CHANGELOG (skill); 2. Run quality gate locally; 3. Build & smoke; 4. Generate release (skill); 5. Deploy (skill); 6. Capture release evidences (skill)
- Claim: --- title: "Release Checklist" version: 3.8.51 lastUpdated: 2026-08-28 --- > **Last updated:** 2026-08-28 — v3.8.51 > Streamlined release flow that leverages Claude Code skills ...

### app/docs/ops/RELEASE_GREEN.md

- Headings: Release-Green: keeping the queue and release branch green; The problem this solves; The family (4 pieces) — and how each runs independently; PR-to-release advisory build; Solution C — `npm run check:release-green` (the engine); Solution A — `/green-prs` (the queue scan); Recommended cadence; Relationship to release
- Claim: --- title: "Release-Green — keeping the queue and release branch green" --- The **full gate** (`.github/workflows/ci.yml` — unit shards, vitest, ratchets, `package-artifact`, So...

### app/docs/ops/RUNNER_BOX.md

- Headings: Self-Hosted Runner Box Operations (.113 pool); Install the janitor (one-time, on the box); cron (root): every 30 min, log to /var/log/runner-janitor.log; Runner units: KillMode; Operating rules
- Claim: --- title: Self-Hosted Runner Box Operations --- The self-hosted pool (`self-hosted, omni-release` on all eight runners; `omni-build` on two) runs on the **.113** box. Measured ...

### app/docs/ops/SQLITE_RUNTIME.md

- Headings: SQLite Runtime Resolution; Why this complexity?; Magic-byte validation; Checking the active driver; Manual control; Skip postinstall warm-up (for fast CI installs); Force-reinstall runtime better-sqlite3; Check what driver is active
- Claim: --- title: "SQLite Runtime Resolution" --- OmniRoute resolves its SQLite driver at startup through a 5-step fallback chain: 1. **Bundled `better-sqlite3`** (via `dependencies` i...

### app/docs/ops/TUNNELS_GUIDE.md

- Headings: Tunnels Guide; Backends at a glance; 1. Cloudflare Tunnel (Quick Tunnel + Named Tunnel); Named tunnel setup (stable hostname); optional — overrides the hostname OmniRoute reports; otherwise read from the; config's first ingress rule:; export CLOUDFLARED_HOSTNAME="ai.example.com"; Enable / disable via REST
- Claim: --- title: "Tunnels Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/{cloudflaredTunnel,ngrokTunnel,tailscaleTunnel}.ts`, `src/app/api/tunnels/...

### app/docs/ops/VM_DEPLOYMENT_GUIDE.md

- Headings: OmniRoute — Deployment Guide on VM with Cloudflare; Prerequisites; 1. Configure the VM; 1.1 Create the instance; 1.2 Connect via SSH; 1.3 Update the system; 1.4 Install Docker; Install dependencies
- Claim: --- title: "OmniRoute — Deployment Guide on VM with Cloudflare" version: 3.8.40 lastUpdated: 2026-06-28 --- 🌐 **Languages:** 🇺🇸 [English](./VM_DEPLOYMENT_GUIDE.md) | 🇪🇹 [አማርኛ](....

### app/docs/plans/11690-anysearch-provider-integration.md

- Headings: AnySearch provider integration (PR #11690); 1. Service ground truth (official docs, cross-checked); 2. Touch set (~10 layers, mirrors the xquik footprint); 3. Decisions (all reached 2026-08-26); 4. References
- Claim: Template: xquik #11370 (merged in release/v3.8.51). Posture: `fallbackOnly`. Upstream proposal: issue diegosouzapw/OmniRoute#11637. - Base URL: `https://api.anysearch.com` (REST...

### app/docs/plans/7778-hierarchical-admission.md

- Headings: #7778 hierarchical admission cleanup plan
- Claim: 1. Lock the existing single-key semaphore contract and the new atomic multi-key contract with focused tests: no partial reservations, FIFO queueing, abort, timeout, queue-full, ...

### app/docs/providers/AGENTROUTER.md

- Headings: AgentRouter Setup Guide; Quick start — use the native `agentrouter` provider (recommended); Advanced: connecting via the Claude Code compatible provider type; Prerequisites; 1. Enable the CC-compatible provider type; 2. Create the provider in the dashboard; 3. Use it through a combo or directly; Wire image details
- Claim: --- title: "AgentRouter Setup Guide" version: 3.8.40 lastUpdated: 2026-06-28 --- [AgentRouter](https://agentrouter.org) is an Anthropic-compatible relay that resells Claude and ...

### app/docs/providers/ALIBABA-QWEN-PROVIDER-FAMILIES.md

- Headings: Alibaba and Qwen provider families; Decision; Endpoint matrix; Connection contract; Compatibility and rollout; Source of truth
- Claim: --- title: "Alibaba and Qwen provider families" description: "Regional provider design for Alibaba Model Studio and Qwen Cloud products" --- This document records the implementa...

### app/docs/providers/CHATGPT_WEB.md

- Headings: Providers — ChatGPT Web (Codex); Common provider retirement; Prerequisites; Dashboard setup; Models and combos; Security model; Displayless VPS and Docker; Interactive recovery
- Claim: --- title: "Providers — ChatGPT Web (Codex)" version: 3.8.51 lastUpdated: 2026-08-31 --- `chatgpt-web-codex` (alias `cgpt-codex`) bridges Codex Responses turns through an authen...

### app/docs/providers/CLAUDE_WEB.md

- Headings: Providers — Claude Web; `claude-web`; Model catalog; Credentials and organization resolution; Conversation operations; Tools and request payloads; Transport selection; SSE behavior
- Claim: --- title: "Providers — Claude Web" version: 3.8.49 lastUpdated: 2026-07-14 --- `claude-web` sends OpenAI-format chat requests through an authenticated `claude.ai` browser sessi...

### app/docs/providers/COPILOT-M365.md

- Headings: Providers — Microsoft 365 Copilot (BizChat); Prerequisites; Credential capture (WebSocket method); Credential capture (HAR import); Token lifetime and refresh; Automatic refresh with `refreshToken`; Tier selection; Security model
- Claim: --- title: "Providers — Microsoft 365 Copilot (BizChat)" version: 3.8.51 lastUpdated: 2026-09-11 --- `copilot-m365-web` (alias `m365copilot`) sends OpenAI-format chat requests t...

### app/docs/providers/CSV-IMPORT.md

- Headings: Import providers from a CSV or JSON file; CSV (positional); OmniRoute provider import (positional columns); JSON
- Claim: --- title: "Import providers from a CSV or JSON file" --- Dashboard → Providers → **Import from file** creates API-key connections from a CSV or JSON list. Each row can target a...

### app/docs/providers/CURSOR-API-KEY-AND-CLI.md

- Headings: Cursor API provider and the Cursor CLI passthrough; Why the key is exchanged; The `cursor-api` provider; Native model IDs and effort; Cursor CLI passthrough; Setup; Failure modes
- Claim: --- title: "Cursor API provider and the Cursor CLI passthrough" version: 3.8.51 lastUpdated: 2026-09-05 --- Two ways to put Cursor behind OmniRoute without an IDE session: 1. **...

### app/docs/providers/CURSOR-DOCKER.md

- Headings: Cursor Provider in Docker Environments; Why IDE / CLI Import Fails in Docker; Recommended: Login with Cursor; Model discovery; Live catalog is exclusive when synced; Helpers; Provider Limits (quota); Empty turns / out of usage
- Claim: --- title: "Cursor Provider in Docker Environments" version: 3.8.50 lastUpdated: 2026-08-17 --- When OmniRoute runs inside Docker, the legacy **Import from Cursor IDE** / `curso...

### app/docs/providers/CURSOR_IMAGE.md

- Headings: Cursor Image Generation; Why the Agent CLI; Access restriction (Hard Rules #15 + #17); Concurrency gate is module-level (single-instance limitation); Requirements; Example; LiteLLM
- Claim: --- title: "Cursor Image Generation" version: 3.8.49 lastUpdated: 2026-07-23 --- OmniRoute exposes Cursor plan **image generation** on `POST /v1/images/generations` through the ...

### app/docs/providers/ZED-DOCKER.md

- Headings: Zed IDE Integration in Docker Environments; Why Keychain Import Fails in Docker; Using the Manual Token Import Tab; Where Zed Stores API Keys on the Host; Volume-Mount Option (Advanced); docker-compose.yml snippet; Linux host; macOS host (uncomment instead)
- Claim: --- title: "Zed IDE Integration in Docker Environments" version: 3.8.40 lastUpdated: 2026-06-28 --- When OmniRoute runs inside Docker, the standard "Import from Zed Keychain" fl...

### app/docs/reference/API_REFERENCE.md

- Headings: API Reference; Chat Completions; Custom Headers; Exclusive Managed Session Leases; `x-omniroute-compression`; Embeddings; List all embedding models; Image Generation
- Claim: --- title: "API Reference" version: 3.8.51 lastUpdated: 2026-08-31 --- 🌐 **Languages:** 🇺🇸 [English](./API_REFERENCE.md) | 🇪🇹 [አማርኛ](../i18n/am/docs/reference/API_REFERENCE.md) ...

### app/docs/reference/CLI-TOOLS.md

- Headings: CLI Tools — OmniRoute; How It Works; Auto-configure with `setup-*`; Running these inside a container; Source of Truth; Capability tiers (cataloged × detectable × configurable × launchable); 1. CLI Code's Catalog (26 tools); 2. CLI Agents Catalog (10 tools)
- Claim: --- title: "CLI Tools — OmniRoute" version: 3.8.50 lastUpdated: 2026-08-23 --- Last updated: 2026-08-23 OmniRoute integrates with three categories of CLI tools spread across thr...

### app/docs/reference/EMBEDDINGS.md

- Headings: Embeddings client runbook; Working model ids; Broken / misleading ids; Native Gemini Embedding 2; Jina multimodal path; Jina / Memorix image object; Client notes; Hindsight 0.9.1
- Claim: --- title: "Embeddings client runbook" lastUpdated: 2026-08-17 --- Operator notes for `POST /v1/embeddings` when OmniRoute sits in front of Hindsight 0.9.1 (text-only `encode(li...

### app/docs/reference/ENVIRONMENT.md

- Headings: Environment Variables Reference; 1. Required Secrets; Generation Commands; Generate all four secrets at once:; 2. Storage & Database; Scenarios; 3. Network & Ports; Port Modes
- Claim: --- title: "Environment Variables Reference" version: 3.8.50 lastUpdated: 2026-08-18 --- > Complete reference for every environment variable recognized by OmniRoute. > For a qui...

### app/docs/reference/FEATURE_FLAGS.md

- Headings: Feature Flags; What Feature Flags Are; Resolution Order; Flag Catalog; Security (10); Network (22); Policies (5); Runtime (33)
- Claim: --- title: "Feature Flags" version: 3.8.51 lastUpdated: 2026-09-03 --- > Runtime toggles that change OmniRoute's behavior **without a redeploy**. > Every flag listed here is def...

### app/docs/reference/FREE_PROXIES_API.md

- Headings: Free Proxies API; List route — `GET /api/settings/free-proxies`; Query parameters; Response; Add-to-pool — `POST /api/settings/free-proxies/[id]/add-to-pool`; Sync — `POST /api/settings/free-proxies/sync`; Stats — `GET /api/settings/free-proxies/stats`
- Claim: --- title: "Free Proxies API" version: 3.8.43 lastUpdated: 2026-07-11 --- OmniRoute ships a curated pool of free proxies in the `free_proxies` table, synced from external provid...

### app/docs/reference/FREE_TIERS.md

- Headings: Free Tiers & Free-Token Budget; TL;DR — how much free inference does OmniRoute actually aggregate?; 2026-06-17 refresh — what changed since 2026-06-05; Two regimes — counting vs deciding; Methodology & caveats; Why our number is smaller than other aggregators'; ToS attention table; ⚠️ Caution — personal-use / proxy clauses worth checking (16)
- Claim: --- title: "Free Tiers & Free-Token Budget" version: 3.8.50 lastUpdated: 2026-09-03 --- > **For Users**: Looking for a simple guide? See the [Free Tiers Guide](../getting-starte...

### app/docs/reference/OPENCODE-SESSION-IDENTITY.md

- Headings: OpenCode conversation identity
- Claim: --- title: "OpenCode conversation identity" version: 3.8.51 lastUpdated: 2026-09-21 --- OpenCode and OpenCode Go reuse explicit conversation IDs across turns. The resolver check...

### app/docs/reference/PROVIDER_PLUGIN_MANIFEST.md

- Headings: Provider Plugin Manifest; Refreshing the Manifest; Goal; Contract; Capability Tags; Sidecar Use; Why Not Embed Providers Directly In Next
- Claim: --- title: "Provider Plugin Manifest" version: 3.8.42 lastUpdated: 2026-07-01 --- `open-sse/config/providerPluginManifest.ts` defines the JSON-safe provider plugin contract. `op...

### app/docs/reference/PROVIDER_REFERENCE.md

- Headings: Provider Reference; Categories; No-auth Providers (no key required) (10); OAuth Providers (25); Web Cookie Providers (33); API Key Providers (paid / paid-with-free-credits) (241); Local Providers (14); Search Providers (17)
- Claim: --- title: "Provider Reference" version: 3.8.51 lastUpdated: 2026-09-24 --- > **Auto-generated** from `src/shared/constants/providers.ts` — do not edit by hand. > Regenerate wit...

### app/docs/reference/RELAY_BACKEND_STRATEGY.md

- Headings: Relay Backend Strategy; Summary; Mode behavior; 9router vs CLIPROXYAPI today; Backend boundary contract; High-throughput guidance; Suggested baseline; Provider plugin contract
- Claim: --- title: "Relay Backend Strategy" version: 3.8.42 lastUpdated: 2026-06-30 --- OmniRoute now supports three relay modes for `/api/v1/relay/chat/completions`: - `ts`: Use the Ty...

### app/docs/reference/RELAY_TROUBLESHOOTING.md

- Headings: Relay Troubleshooting; How relay auth is stored; Failure mode 1 — undecryptable token after key rotation; Failure mode 2 — unrecoverable token (key rotated away); Failure mode 3 — relay reachable but unhealthy; API reference
- Claim: --- title: "Relay Troubleshooting" version: 3.8.43 lastUpdated: 2026-07-11 --- Relays (Vercel, Deno, Cloudflare) terminate the upstream connection on a serverless backend so Omn...

### app/docs/reference/REMOVED_PROVIDERS.md

- Headings: Providers removed at their operator's request; Policy; Register; Adding an entry
- Claim: --- title: "Removed Providers" version: 3.8.51 lastUpdated: 2026-09-03 --- Some services were integrated into OmniRoute and later removed because the people who run them asked f...

### app/docs/routing/AUTO-COMBO.md

- Headings: OmniRoute Auto-Combo Engine; Zero-Config Auto-Routing (`auto/` prefix); Quick Examples; Category × Tier Composition (`auto/<category>:<tier>`); Any IDE or CLI tool that supports OpenAI format; In your code/config, set model to:; Per-key candidate control (#7819, Level 1+2); Combo Names That Match a Real Model Id
- Claim: --- title: "OmniRoute Auto-Combo Engine" version: 3.8.40 lastUpdated: 2026-06-28 --- > **For Users**: Looking for a quick start? See the [Auto-Combo User Guide](../getting-start...

### app/docs/routing/DETERMINISTIC_ROUTING.md

- Headings: Deterministic routing strategies; Enable; providers.yaml; The five strategies; Pinned providers; Fallback on failure; Explainability; Failure contract (unchanged from RIC-738)
- Claim: --- title: "Deterministic Routing Strategies" --- OmniRoute's self-hosted gateway entry (`/v1/chat/completions`, RIC-738) routes to a single provider by default. When you config...

### app/docs/routing/MODEL_EXPOSURE_LIST.md

- Headings: Model Exposure Allow/Deny List; Why this exists; Settings; Two chokepoints, not one; What is NOT filtered
- Claim: --- title: "Model Exposure Allow/Deny List" version: 3.8.51 lastUpdated: 2026-08-29 --- > Opt-in, off by default (`modelVisibilityAllowlist`/`modelVisibilityDenylist` both empty...

### app/docs/routing/QUOTA_SHARE.md

- Headings: Quota Sharing Engine; Overview; Algorithm: Fair-Share Work-Conserving; Modes; Per-dimension decision; Cap absoluto; Multi-dimension check; Borrowing
- Claim: --- title: "Quota Sharing Engine" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Doc reference**: `docs/routing/QUOTA_SHARE.md` > Part of Group B (plans 16 + 22). --- The Quota...

### app/docs/routing/REASONING_REPLAY.md

- Headings: Reasoning Replay Cache; Why This Exists; Architecture; Storage — Hybrid Memory + SQLite; Database Schema; Provider / Model Detection; REST API; Operational Notes
- Claim: --- title: "Reasoning Replay Cache" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/db/reasoningCache.ts`, `open-sse/services/reasoningCache.ts` > **...

### app/docs/routing/REASONING_ROUTING.md

- Headings: Reasoning Routing; Management; Rule Resolution; Effort and Budget; Security and Transports; Persistence
- Claim: --- title: "Reasoning Routing" --- Reasoning routing rules extend the existing model and combo routing. When no active rule matches, the existing thinking, suffix, connection-de...

### app/docs/routing/SELF_HOSTED_OPENAI_ENTRY.md

- Headings: Self-hosted unified OpenAI-compatible entry; Enable; providers.yaml; apiKey: sk-...       # optional, runtime-only; Provider auto-route; via header; or via model prefix; Auth scaffold (D5 reserved)
- Claim: --- title: "Self-Hosted OpenAI-Compatible Entry" --- When enabled, OmniRoute's existing `/v1/chat/completions` (and the OpenAI-compatible contract it serves) becomes a **self-ho...

### app/docs/routing/STRICT_ZERO_COST.md

- Headings: STRICT_ZERO_COST; Why this exists, and why `hidePaidModels` alone isn't enough; Candidate classification; Connection safety (per-connection verification, never per-candidate); Quota caching (`open-sse/services/autoCombo/freeAccessQuota.ts`); ToS guard (independent of economic safety); Seeing what the guard excludes; Enabling
- Claim: --- title: "STRICT_ZERO_COST" version: 3.8.50 lastUpdated: 2026-08-20 --- > Opt-in, off by default (`settings.freeAccessPolicy !== "strict"` leaves every `auto/*` > candidate po...

### app/docs/routing/SUBSCRIPTION_LADDER.md

- Headings: Subscription-first routing; Why this exists; Billing is a connection fact, not a model fact; The rung model; `auto/subscription` — fail closed; Connection safety; `auto/thrifty` — escalate one rung at a time; Returning to the plan after a reset
- Claim: --- title: "Subscription-first routing" version: 3.8.50 lastUpdated: 2026-08-22 --- > Two new `auto/*` ids — `auto/subscription` and `auto/thrifty`. Both are opt-in by being > r...

### app/docs/security/AGENTROUTER_WAF.md

- Headings: agentrouter.org WAF (Web Application Firewall); Scope of the filter; Always-blocked patterns (case-insensitive); Almost-always-blocked patterns; Behavior under load; Mitigations already applied in OmniRoute; Guidance for prompts and tool output; Reporting the false positives upstream
- Claim: --- title: "agentrouter.org WAF (Web Application Firewall)" version: 3.8.50 lastUpdated: 2026-08-03 --- The `agentrouter` upstream gateway runs a keyword-based content filter on...

### app/docs/security/BAN_DETECTION.md

- Headings: Account-Ban / Banned-Keyword Detection; Built-in keywords; Not a ban: operator-actionable verification prompts; Detection flow; Scope — which providers are scanned; Custom banned keywords; Recovering a flagged connection; Probe isolation (model test-all)
- Claim: --- title: Account-Ban / Banned-Keyword Detection --- OmniRoute scans upstream error responses for signals that indicate a provider **account is permanently dead** (suspended / ...

### app/docs/security/CLI-CONFIGURATION.md

- Headings: CLI configuration previews and writes
- Claim: --- title: "CLI configuration previews and writes" --- The authenticated `/api/cli-tools/config` and `/api/cli-tools/apply` routes return non-cacheable, redacted previews. They ...

### app/docs/security/CLI_TOKEN.md

- Headings: CLI Machine-ID Token; Overview; How it works; Security properties; Default salt (random per install); Salt rotation; Persistent rotation (add to shell profile); Verify new token is in use
- Claim: --- title: "CLI Machine-ID Token" --- OmniRoute CLI commands authenticate against the local management API using a `HMAC-SHA256(machine-id, salt)` token sent via the `x-omnirout...

### app/docs/security/COMPLIANCE.md

- Headings: Compliance & Audit; What Gets Logged; Administrative audit events (`audit_log`); MCP tool calls (`mcp_tool_audit`); Request / usage logs; Storage Schema; Retention & Cleanup; `noLog` Opt-Out (per API key)
- Claim: --- title: "Compliance & Audit" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Source of truth:** `src/lib/compliance/`, `src/app/api/compliance/` > **Last updated:** 2026-06-2...

### app/docs/security/CORS.md

- Headings: CORS Configuration & Security; How an origin is resolved; Threat model — what `CORS_ALLOW_ALL=true` really exposes; One genuine exception — `/api/v1/agents/`; Production checklist; Development convenience — allow specific local origins; Vite (5173) + Next.js (3000) dev servers calling a local OmniRoute; API keys vs cookie sessions
- Claim: --- title: CORS Configuration & Security --- OmniRoute controls which **browser origins** may read cross-origin responses from a single, centralized allowlist. The model is **fa...

### app/docs/security/EGRESS_POLICY.md

- Headings: Egress IP Family Policy (IPv4/IPv6); What It Is; Why It Exists; The Three Values; How to Configure It; Dashboard; API; Create an IPv6-only proxy
- Claim: --- title: "Egress IP Family Policy (IPv4/IPv6)" version: 3.8.40 lastUpdated: 2026-06-28 --- > **Pin outbound traffic to a single IP family — `auto`, `ipv4`, or `ipv6` — per pro...

### app/docs/security/ERROR_SANITIZATION.md

- Headings: Error Message Sanitization; Why this exists; The mandatory pattern; 1. Building an error response (HTTP / API routes); 2. Custom error envelopes (rare); 3. Logging vs. responding; 4. Forbidden patterns; Coverage in CI
- Claim: --- title: "Error Message Sanitization" version: 3.8.51 lastUpdated: 2026-09-02 --- > **Source of truth:** `open-sse/utils/errorSanitization.ts`, > `open-sse/utils/errorPathReda...

### app/docs/security/GUARDRAILS.md

- Headings: Guardrails; Built-in Guardrails; Vision Bridge (`visionBridge.ts`) — Modality Bridge PR-1; Audio Bridge (`audioBridge.ts`) — Modality Bridge PR-3; Video Bridge (`videoBridge.ts`, `videoBridgePipeline.ts`); PII Masker (`piiMasker.ts`); Prompt Injection (`promptInjection.ts`); Credential Masker (`credentialMasker.ts`)
- Claim: --- title: "Guardrails" version: 3.8.51 lastUpdated: 2026-08-29 --- > **Source of truth:** `src/lib/guardrails/` > **Last updated:** 2026-08-29 — v3.8.51 (Video Bridge transcrip...

### app/docs/security/INFERENCE_AUTH_POSTURE.md

- Headings: Inference Auth Posture; Overview; How to actually check; Startup warning; Related
- Claim: --- title: "Inference Auth Posture" --- `GET /v1/models` and the inference endpoints are gated by **different settings**. The most natural probe an operator runs to answer "is m...

### app/docs/security/ROUTE_GUARD_TIERS.md

- Headings: Route Guard Tiers; Overview; Tiers; Tier 1 — LOCAL_ONLY; Tier 2 — ALWAYS_PROTECTED; Fresh-install bootstrap is loopback-only — by real peer, not `Host`; Tier 3 — MANAGEMENT (default); Evaluation order
- Claim: --- title: "Route Guard Tiers" --- All OmniRoute management API routes are classified into one of three protection tiers. Classification is static, defined in `src/server/authz/...

### app/docs/security/SUPPLY_CHAIN.md

- Headings: Supply-Chain Gates (Phase 8 · Block A); ⚠️ CVE Variance (blocking osv/Trivy gates); Known Accepted Risks; extract-zip 2.0.1 — GHSA-7pqw-9j4j-h8q3 / GHSA-jmr9-qjv8-65gv (#14482); Backlog: Scorecard advisory → blocking
- Claim: --- title: "Supply-Chain Gates" --- OmniRoute publishes npm + Docker artifacts. These gates provide provenance, inventory (SBOM) and CVE scanning, all OSS, plugged into release ...

### app/skills/cli-a2a/SKILL.md

- Headings: Overview; Quick install; Subcommands; `a2a`; `a2a status`; `a2a card`; `a2a skills`; `a2a invoke <skill>`
- Claim: --- name: cli-a2a description: Interact with the OmniRoute A2A server from the CLI. Send tasks, inspect skill execution history, and test the JSON-RPC 2.0 agent-to-agent protoco...

### app/skills/cli-backup-sync/SKILL.md

- Headings: Overview; Quick install; Subcommands; `backup`; `backup create`; `backup auto`; `backup enable`; `backup disable`
- Claim: --- name: cli-backup-sync description: Backup and restore OmniRoute data from the CLI. Trigger incremental snapshots, sync to cloud storage, manage backup schedules, and restore...

### app/skills/cli-batches/SKILL.md

- Headings: Overview; Quick install; Subcommands; `batches`; `batches list`; `batches get <batchId>`; `batches create`; `batches submit`
- Claim: --- name: cli-batches description: Submit and monitor batch inference jobs from the CLI. Upload and manage files for batch processing, retrieve results, and integrate batch pipe...

### app/skills/cli-chat/SKILL.md

- Headings: Overview; Quick install; Subcommands; `chat [prompt]`; `repl`; `stream [prompt]`
- Claim: --- name: cli-chat description: Send chat completions, stream responses, and start an interactive REPL session from the CLI. Supports all OmniRoute providers, combo routing, and...

### app/skills/cli-compression/SKILL.md

- Headings: Overview; Quick install; Subcommands; `compression`; `compression status`; `compression configure`; `compression engine`; `compression set <name>`
- Claim: --- name: cli-compression description: Configure and test prompt compression from the CLI. Manage RTK filters, Caveman rules, stacked compression modes, and preview compression ...

### app/skills/cli-contexts/SKILL.md

- Headings: Overview; Quick install; Subcommands; `context-eng`; `context-eng analytics`; `context-eng caveman`; `context-eng config`; `context-eng show`
- Claim: --- name: cli-contexts description: Manage context engineering configurations, RTK filter sets, and conversation sessions from the CLI. Apply context-relay settings and inspect ...

### app/skills/cli-cost-usage/SKILL.md

- Headings: Overview; Quick install; Subcommands; `cost`; `pricing`; `pricing sync`; `pricing list`; `pricing get <model>`
- Claim: --- name: cli-cost-usage description: View cost breakdowns, token usage, and call logs from the CLI. Filter by provider, model, or date range. Export usage reports and inspect p...

### app/skills/cli-eval/SKILL.md

- Headings: Overview; Quick install; Subcommands; `eval`; `eval suites`; `eval list`; `eval get <suiteId>`; `eval create`
- Claim: --- name: cli-eval description: Create and run evaluation suites, watch live benchmark progress, view scorecards, compare model performance, and integrate eval runs with CI work...

### app/skills/cli-health/SKILL.md

- Headings: Overview; Quick install; Subcommands; `doctor`; `health`; `health components`; `health watch`; `status`
- Claim: --- name: cli-health description: Check server health, component status, and live metrics from the CLI. Run `health`, `health components`, and `health watch` for a real-time das...

### app/skills/cli-keys/SKILL.md

- Headings: Overview; Quick install; Subcommands; `keys`; `keys add <provider> [apiKey]`; `keys list`; `keys remove <provider>`; `keys regenerate <id>`
- Claim: --- name: cli-keys description: Create, list, rotate, and revoke OmniRoute API keys from the CLI. Manage OAuth flows for provider authentication and inspect key scopes and expir...

### app/skills/cli-mcp/SKILL.md

- Headings: Overview; Quick install; Subcommands; `mcp`; `mcp status`; `mcp restart`; `mcp enable`; `mcp disable`
- Claim: --- name: cli-mcp description: Inspect the MCP server status, list registered tools and scopes, run tool invocations, and manage MCP audit logs from the CLI. --- Inspect the MCP...

### app/skills/cli-models/SKILL.md

- Headings: Overview; Quick install; Subcommands; `models [provider]`; `models manual <provider>`; `models add <provider> <model-id>`; `models edit <provider> <model-id>`; `models remove <provider> <model-id>`
- Claim: --- name: cli-models description: Query available AI models, list model aliases, and browse the full model catalog from the CLI. Filter by provider, search by capability, and re...

### app/skills/cli-plugins-skills/SKILL.md

- Headings: Overview; Quick install; Subcommands; `memory`; `memory search <query>`; `memory add`; `memory clear`; `memory list`
- Claim: --- name: cli-plugins-skills description: Manage Omni Skills (list, install, test, remove), plugins (create, configure), and persistent memory (search, add, clear) from the CLI....

### app/skills/cli-policy-audit/SKILL.md

- Headings: Overview; Quick install; Subcommands; `audit`; `audit tail`; `audit search <query>`; `audit export <file>`; `audit stats`
- Claim: --- name: cli-policy-audit description: Inspect audit logs, manage access policies, view telemetry data, and review request history from the CLI. Filter by event type, user, or ...

### app/skills/cli-providers/SKILL.md

- Headings: Overview; Quick install; Subcommands; `provider-cmd provider [subcommand]`; `providers`; `providers available`; `providers list`; `providers test <idOrName>`
- Claim: --- name: cli-providers description: "Manage provider connections from the CLI: list available/configured providers, add, test, test-all, validate, rotate API keys, and view per...

### app/skills/cli-resilience/SKILL.md

- Headings: Overview; Quick install; Subcommands; `quota`; `quota status`; `quota preview`; `quota ensure <json>`; `resilience`
- Claim: --- name: cli-resilience description: Inspect and manage circuit-breaker states, connection cooldowns, quota limits, and backoff levels from the CLI. Reset stuck providers and c...

### app/skills/cli-routing/SKILL.md

- Headings: Overview; Quick install; Subcommands; `suggest`; `combo combo`; `combo list`; `combo switch <name>`; `combo create <name>`
- Claim: --- name: cli-routing description: Create, list, update, and delete routing combos from the CLI. Test routing strategies, inspect combo metrics, and configure fallback chains in...

### app/skills/cli-serve/SKILL.md

- Headings: Overview; Quick install; Subcommands; `dashboard`; `restart`; `serve`; `stop`; Setup
- Claim: --- name: cli-serve description: Start, stop, and restart the OmniRoute server from the CLI. Manage daemon mode, port configuration, auto-recovery, system tray integration, and ...

### app/skills/cli-setup/SKILL.md

- Headings: Overview; Quick install; Subcommands; `autostart`; `autostart enable`; `autostart disable`; `autostart toggle`; `autostart status`
- Claim: --- name: cli-setup description: Run initial setup, configure global CLI settings, manage environment variables, check for updates, and configure autostart via the CLI setup and...

### app/skills/cli-skill-collector/SKILL.md

- Headings: Overview; Quick install; Subcommands; `autostart`; `autostart enable`; `autostart disable`; `autostart toggle`; `autostart status`
- Claim: --- name: cli-skill-collector description: "Detect installed CLI coding tools (Claude Code, Codex, Cursor, Copilot, Cline and more), search GitHub for matching agent skills, and...

### app/skills/cli-tunnel/SKILL.md

- Headings: Overview; Quick install; Subcommands; `tunnel`; `tunnel list`; `tunnel create [type]`; `tunnel stop <type>`; `tunnel status <type>`
- Claim: --- name: cli-tunnel description: Start and stop tunnel connections (ngrok, Cloudflare, custom) from the CLI. Inspect active tunnel URLs, configure authentication, and test exte...

### app/skills/config-codex-cli/SKILL.md

- Headings: Overview; Quick install; Subcommands; Long-running Codex tasks (#7287); Session affinity (default off); Stream idle timeout (default 10 minutes); Recommended multi-hour recipe; Defaults decision
- Claim: --- name: config-codex-cli description: Step-by-step agent workflow to configure the OpenAI Codex CLI on any machine (Linux, macOS, Windows) to use OmniRoute as an OpenAI-compat...

### app/skills/omni-agents-a2a/SKILL.md

- Headings: Overview; Authentication; Endpoints; Payloads; OmniRoute — A2A Protocol; Discovery; Available skills; Call example (JSON-RPC 2.0)
- Claim: --- name: omni-agents-a2a description: "Interact with OmniRoute via JSON-RPC 2.0 agent-to-agent protocol. 6 built-in A2A skills: smart-routing, quota-management, provider-discov...

### app/skills/omni-api-keys/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/keys; POST /api/keys; GET /api/keys/{id}; PATCH /api/keys/{id}; DELETE /api/keys/{id}
- Claim: --- name: omni-api-keys description: Create, list, rotate, and revoke OmniRoute API keys. Control per-key scopes, spending limits, and expiration. Keys gate access to all proxy ...

### app/skills/omni-auth/SKILL.md

- Headings: Overview; Authentication; Endpoints; POST /api/auth/login; POST /api/auth/logout; GET /api/auth/oidc/login; GET /api/auth/oidc/callback; GET /api/auth/csrf
- Claim: --- name: omni-auth description: Manage API key authentication and session tokens. Start here to authenticate requests via Bearer token, obtain session cookies, and configure lo...

### app/skills/omni-budget/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/rate-limit; POST /api/rate-limit; Payloads
- Claim: --- name: omni-budget description: Configure spending limits, token quotas, and rate-limit policies per API key or globally. Inspect current consumption and enforce cost control...

### app/skills/omni-cache/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/cache; DELETE /api/cache; GET /api/cache/stats; DELETE /api/cache/stats; GET /api/cache/entries
- Claim: --- name: omni-cache description: Manage the LLM response cache. View cache statistics, clear entries, configure TTL policies, and control semantic-similarity caching thresholds...

### app/skills/omni-cli-tools/SKILL.md

- Headings: Overview; Authentication; Endpoints; Payloads
- Claim: --- name: omni-cli-tools description: "Manage CLI tool integrations exposed via the API. List, configure, and invoke CLI tool plugins that extend OmniRoute's automation surface....

### app/skills/omni-cli-tools/references/endpoints.md

- Headings: Endpoint reference
- Claim: (structure only)

### app/skills/omni-combos-routing/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/combos; POST /api/combos; GET /api/combos/{id}; PUT /api/combos/{id}; PATCH /api/combos/{id}
- Claim: --- name: omni-combos-routing description: Create and manage routing combos with 19 strategies (priority, weighted, round-robin, Auto-combo, and more). Configure fallback chains...

### app/skills/omni-compression/SKILL.md

- Headings: Overview; Authentication; Endpoints; POST /api/compression/preview; GET /api/compression/language-packs; GET /api/compression/rules; POST /api/compression/compare; POST /api/compression/compare/verify
- Claim: --- name: omni-compression description: Configure RTK (command output), Caveman (prose), and stacked compression modes. Manage language packs, custom rules, and test prompt comp...

### app/skills/omni-context-rtk/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/context/rtk/config; PUT /api/context/rtk/config; GET /api/context/rtk/filters; POST /api/context/rtk/import; POST /api/context/rtk/test
- Claim: --- name: omni-context-rtk description: Configure RTK filters, context engineering rules, and context relay settings. Test compression with real prompt samples and manage contex...

### app/skills/omni-db-backups/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/system/env/repair; POST /api/system/env/repair; GET /api/system/version; POST /api/system/version; Payloads
- Claim: --- name: omni-db-backups description: Trigger system backups, restore from backup files, and manage the SQLite database lifecycle. Supports export, import, and incremental snap...

### app/skills/omni-github-skills/SKILL.md

- Headings: Overview; Authentication; Endpoints; Payloads
- Claim: --- name: omni-github-skills description: Search, score, scan, and import agent skills from GitHub repositories that contain SKILL.md, CLAUDE.md, .cursorrules, and similar agent...

### app/skills/omni-inference/SKILL.md

- Headings: Overview; Authentication; Endpoints; Payloads; Chat completions; Endpoints; Discover; OpenAI format example
- Claim: --- name: omni-inference description: "The core OpenAI-compatible inference endpoints: chat completions, embeddings, images, audio (TTS/STT), moderations, rerank, and the Respon...

### app/skills/omni-inference/references/endpoints.md

- Headings: Endpoint reference
- Claim: (structure only)

### app/skills/omni-mcp/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/mcp/audit; GET /api/mcp/audit/stats; GET /api/mcp/sse; POST /api/mcp/sse; GET /api/mcp/status
- Claim: --- name: omni-mcp description: "Connect to the OmniRoute MCP server (110 tools, 3 transports: SSE/stdio/HTTP). Covers routing, cache, compression, memory, skills, providers, an...

### app/skills/omni-models/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/v1/models; GET /api/models; POST /api/models/alias; GET /api/models/catalog; GET /api/models/openrouter-catalog
- Claim: --- name: omni-models description: Query available AI models across all configured providers. List models, resolve model aliases, and browse the full model catalog including pro...

### app/skills/omni-providers/SKILL.md

- Headings: Overview; Authentication; Endpoints; Payloads; Additional endpoints; POST /api/providers/{id}/refresh; GET /api/provider-metrics; GET /api/providers/health-matrix
- Claim: --- name: omni-providers description: "Manage provider connections, API keys, OAuth flows, and connection tests via the REST API. List, add, update, remove, and test AI provider...

### app/skills/omni-providers/references/endpoints.md

- Headings: Endpoint reference
- Claim: (structure only)

### app/skills/omni-proxies/SKILL.md

- Headings: Overview; Authentication; Endpoints; Payloads
- Claim: --- name: omni-proxies description: Configure HTTP/HTTPS/SOCKS proxies for upstream provider requests. Set per-provider or global proxy rules, test connectivity, and manage prox...

### app/skills/omni-resilience/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/monitoring/health; GET /api/monitoring/compression; GET /api/provider-metrics; Payloads; OmniRoute — Monitoring & Health
- Claim: --- name: omni-resilience description: Monitor provider health, circuit-breaker states, p50/p95/p99 latency metrics, and budget guard alerts. Inspect connection cooldowns and mo...

### app/skills/omni-settings/SKILL.md

- Headings: Overview; Authentication; Endpoints; Payloads; GET /api/settings/thinking-budget — behavior
- Claim: --- name: omni-settings description: "Read and update global application settings: system prompts, thinking budget, IP filters, payload rules, combo defaults, and require-login ...

### app/skills/omni-settings/references/endpoints.md

- Headings: Endpoint reference
- Claim: (structure only)

### app/skills/omni-sync-cloud/SKILL.md

- Headings: Overview; Authentication; Endpoints; POST /api/cloud/auth; PUT /api/cloud/credentials/update; POST /api/cloud/model/resolve; GET /api/cloud/models/alias; PUT /api/cloud/models/alias
- Claim: --- name: omni-sync-cloud description: Synchronise OmniRoute configuration, provider connections, and settings to/from cloud storage. Manage cloud worker authentication and remo...

### app/skills/omni-tunnels/SKILL.md

- Headings: Overview; Authentication; Endpoints; Payloads
- Claim: --- name: omni-tunnels description: Create and manage secure tunnels (ngrok, Cloudflare Tunnel, custom) to expose OmniRoute to the internet or share access with remote agents an...

### app/skills/omni-usage-logs/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/usage/analytics; GET /api/usage/call-logs; GET /api/usage/call-logs/{id}; GET /api/usage/{connectionId}; GET /api/usage/history
- Claim: --- name: omni-usage-logs description: Access detailed call logs and usage analytics. Filter by provider, model, time range, status, and cost. Export logs and aggregate token us...

### app/skills/omni-version-manager/SKILL.md

- Headings: Overview; Authentication; Endpoints; Payloads
- Claim: --- name: omni-version-manager description: Install, start, stop, restart, and update embedded services (9Router, CLIProxyAPI). Monitor service status, retrieve logs, and config...

### app/skills/omni-version-manager/references/endpoints.md

- Headings: Endpoint reference
- Claim: (structure only)

### app/skills/omni-webhooks/SKILL.md

- Headings: Overview; Authentication; Endpoints; GET /api/webhooks; POST /api/webhooks; GET /api/webhooks/{id}; PUT /api/webhooks/{id}; DELETE /api/webhooks/{id}
- Claim: --- name: omni-webhooks description: Register, list, test, and remove webhook endpoints. Configure event subscriptions (request.completed, request.failed, quota.exceeded, etc.) ...

### app/skills/ponytail/SKILL.md

- Headings: Ponytail; Persistence; The ladder; Rules; Output; Intensity; When NOT to be lazy; Boundaries
- Claim: --- name: ponytail description: "External agent skill (MIT, github.com/DietrichGebert/ponytail). Forces the laziest solution that actually works: question whether the task needs...
