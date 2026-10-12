# DeerFlow

status: active

origin: local checkout ~/Projects/deer-flow

memory: reconcile against the workspace topic that names this subject. This card does not replace that topic.

Kept files: 497. Dropped files: 7.

The lines below are a parse. The upstream tree remains the source.

### AGENTS.md

- Headings: AGENTS.md; What is DeerFlow; Service Topology; Repository Map; Commands: Root vs. Module; Backend (see backend/AGENTS.md for the full set); Frontend (see frontend/AGENTS.md for the full set); Prerequisites before `make dev`
- Claim: This is the source of truth for repository agent guidance; `CLAUDE.md` imports it via `@AGENTS.md`. This **monorepo orientation layer** maps the repo. For module details, read i...

### CLAUDE.md

- Headings: CLAUDE.md
- Claim: The repo's agent guidance lives in [AGENTS.md](AGENTS.md) so it is shared across coding agents (Claude Code, Codex, and others). Claude Code imports it below. @AGENTS.md

### Install.md

- Headings: DeerFlow Install; Goal; Operating Rules; Success Criteria; Steps; Verification; Final Response Format; EXECUTE NOW
- Claim: This file is for coding agents. If the DeerFlow repository is not already cloned and open, clone `https://github.com/bytedance/deer-flow.git` first, then continue from the repos...

### README.md

- Headings: 🦌 DeerFlow - 2.0; Official Website; Sister Projects; Coding Plan from ByteDance Volcengine; InfoQuest; One-Line Agent Setup; Quick Start; Configuration
- Claim: English | [中文](./README_zh.md) | [日本語](./README_ja.md) | [Français](./README_fr.md) | [Русский](./README_ru.md) | [Português](./README_pt.md) [![Python](https://img.shields.io/b...

### README_fr.md

- Headings: 🦌 DeerFlow - 2.0; Site officiel; Projets associés; Coding Plan de ByteDance Volcengine; InfoQuest; Table des matières; Installation en une phrase pour un coding agent; Démarrage rapide
- Claim: [English](./README.md) | [中文](./README_zh.md) | [日本語](./README_ja.md) | Français | [Русский](./README_ru.md) | [Português](./README_pt.md) [![Python](https://img.shields.io/badg...

### README_ja.md

- Headings: 🦌 DeerFlow - 2.0; 公式ウェブサイト; 姉妹プロジェクト; ByteDance Volcengine のコーディングプラン; InfoQuest; 目次; Coding Agent に一文でセットアップを依頼; クイックスタート
- Claim: [English](./README.md) | [中文](./README_zh.md) | 日本語 | [Français](./README_fr.md) | [Русский](./README_ru.md) | [Português](./README_pt.md) [![Python](https://img.shields.io/badg...

### README_pt.md

- Headings: 🦌 DeerFlow - 2.0; Site oficial; Projetos irmãos; Coding Plan da ByteDance Volcengine; InfoQuest; Índice; Configuração do agente em uma linha; Início rápido
- Claim: [English](./README.md) | [中文](./README_zh.md) | [日本語](./README_ja.md) | [Français](./README_fr.md) | [Русский](./README_ru.md) | Português [![Python](https://img.shields.io/badg...

### README_ru.md

- Headings: 🦌 DeerFlow - 2.0; Официальный сайт; Родственные проекты; Coding Plan от ByteDance Volcengine; InfoQuest; Содержание; Установка одной фразой для coding agent; Быстрый старт
- Claim: [English](./README.md) | [中文](./README_zh.md) | [日本語](./README_ja.md) | [Français](./README_fr.md) | Русский | [Português](./README_pt.md) [![Python](https://img.shields.io/badg...

### README_zh.md

- Headings: 🦌 DeerFlow - 2.0; 官网; 姐妹项目; 字节跳动火山引擎方舟 Coding Plan; InfoQuest; 目录; 一句话交给 Coding Agent 安装; 快速开始
- Claim: [English](./README.md) | 中文 | [日本語](./README_ja.md) | [Français](./README_fr.md) | [Русский](./README_ru.md) | [Português](./README_pt.md) [![Python](https://img.shields.io/badg...

### RELEASING.md

- Headings: Releasing DeerFlow; Version sources; Helper scripts; Release procedure; Release candidates; What CI publishes on a `v*` tag; Nightly builds; lark-cli sandbox images
- Claim: DeerFlow releases are **tag-driven**: pushing a `v*` git tag triggers the publishing workflows. There is no separate release script that bumps versions — the maintainer bumps th...

### SECURITY.md

- Headings: Security Policy; Supported Versions; Reporting a Vulnerability
- Claim: As deer-flow doesn't provide an official release yet, please use the latest version to receive security updates. Currently, we have two branches to maintain: * main branch for d...

### backend/AGENTS.md

- Headings: AGENTS.md; Project Overview; Important Development Guidelines; Documentation Update Policy; Backend Benchmarks; Commands; Architecture; Harness / App Split
- Claim: The backend runs a LangGraph-based super agent with sandbox execution, persistent memory, subagent delegation, and extensible tools in isolated per-thread environments. **Archit...

### backend/CLAUDE.md

- Headings: CLAUDE.md
- Claim: The backend agent guidance lives in [AGENTS.md](AGENTS.md) so it is shared across coding agents (Claude Code, Codex, and others). Claude Code imports it below. @AGENTS.md

### backend/README.md

- Headings: DeerFlow Backend; Architecture; Core Components; Lead Agent; Middleware Chain; Sandbox System; Subagent System; Memory System
- Claim: **Language:** English | [简体中文](README_zh.md) DeerFlow is a LangGraph-based AI super agent with sandbox execution, persistent memory, and extensible tool integration. The backend...

### backend/README_zh.md

- Headings: DeerFlow 后端; 架构; 核心组件; 主智能体; 中间件链; 沙箱系统; 子智能体系统; 记忆系统
- Claim: **语言：** [English](README.md) | 简体中文 DeerFlow 是一个基于 LangGraph 的 AI 超级智能体，具备沙箱执行、持久记忆和可扩展工具集成能力。后端使 AI 智能体能够执行代码、浏览网页、管理文件、将任务委派给子智能体，并在多轮对话之间保留上下文——所有操作都在按线程隔离的环境中进行。 --- ┌──────...

### backend/app/channels/AGENTS.md

- Headings: IM Channels System (`app/channels/`); QQ WebSocket channel
- Claim: Bridges external messaging platforms (Feishu, Slack, Telegram, Discord, DingTalk, GitHub) to the DeerFlow agent via Gateway's LangGraph-compatible API. WeChat `_read_outbound_by...

### backend/app/gateway/AGENTS.md

- Headings: Gateway API (`app/gateway/`); Route and skill-listing authorization
- Claim: Reject external run/state writes with `sandbox`, `thread_data`, `viewed_images`, or `goal_outcome`. Studio retains sanitized creation metadata. Capability Center's `business` ad...

### backend/app/gateway/routers/AGENTS.md

- Headings: Router contracts
- Claim: The MCP task list accepts optional `TaskStatus` and `active_only` query filters. Pass both to the repository before its limit; preserve `threads:read`, owner checks, current-thr...

### backend/docs/API.md

- Headings: API Reference; Overview; Authentication; Account Preferences; Personal Access Tokens; PAT Constraints; LangGraph-compatible API; Threads
- Claim: This document provides a complete reference for the DeerFlow backend APIs. DeerFlow backend exposes two sets of APIs: 1. **LangGraph-compatible API** - Agent interactions, threa...

### backend/docs/APPLE_CONTAINER.md

- Headings: Apple Container Support; Overview; Benefits; On Apple Silicon Macs with Apple Container:; Fallback to Docker:; Requirements; For Apple Container (macOS only):; Installation:
- Claim: DeerFlow now supports Apple Container as the preferred container runtime on macOS, with automatic fallback to Docker. Starting with this version, DeerFlow automatically detects ...

### backend/docs/ARCHITECTURE.md

- Headings: Architecture Overview; System Architecture; Component Details; Gateway Embedded Agent Runtime; Gateway API; Projects: Pinned Context and the Document Shelf; Agent Architecture; Thread State
- Claim: This document provides a comprehensive overview of the DeerFlow backend architecture. ┌──────────────────────────────────────────────────────────────────────────┐ │ Client (Brow...

### backend/docs/AUTH_DESIGN.md

- Headings: 用户认证与隔离设计; 设计目标; 核心模型; 用户表; 运行时身份; 登录与初始化流程; 首次初始化; 普通登录
- Claim: 本文档描述 DeerFlow 当前内置认证模块的设计，而不是历史 RFC。它覆盖浏览器登录、OIDC/SSO、平台信任接入（IM Channel 与 Internal Auth）、API 认证、CSRF、用户隔离、首次初始化、密码重置和升级迁移。 认证模块的核心目标是把 DeerFlow 从“本地单用户工具”提升为“可多用户部署的 agent runt...

### backend/docs/AUTH_TEST_DOCKER_GAP.md

- Headings: Docker Test Gap (Section 七 7.4); Why this gap exists; (empty); bash: docker: command not found; Cases not executed; Coverage already provided by non-Docker tests; Reproduction steps when Docker becomes available; Required on the host
- Claim: This file documents the only **un-executed** test cases from `backend/docs/AUTH_TEST_PLAN.md` after the full release validation pass. The release validation environment (sg_dev:...

### backend/docs/AUTH_TEST_PLAN.md

- Headings: Auth 模块测试计划; 测试矩阵; 一、环境准备; 1.1 首次启动（干净数据库）; 清除已有数据; 启动标准模式（Gateway embedded runtime）; 1.2 非首次启动; 不清除数据库，直接启动
- Claim: `make dev`、Docker dev 和生产部署默认都运行 Gateway embedded runtime。 `app.gateway.langgraph_auth` 仅用于保留的直连 LangGraph 工具链 / Studio 兼容性测试，不是标准服务启动路径。 每种模式下都需执行以下测试。 --- rm -f backend/.deer-...

### backend/docs/AUTH_UPGRADE.md

- Headings: Authentication Upgrade Guide; 核心概念; 升级步骤; 1. 更新代码; 2. 首次启动; 3. 创建 admin; 4. 登录; 5. 添加用户（可选）
- Claim: DeerFlow 内置了认证模块。本文档面向从无认证版本升级的用户。 完整设计见 [AUTH_DESIGN.md](AUTH_DESIGN.md)。 认证模块采用**始终强制**策略： - 首次启动时不会自动创建账号；首次访问 `/setup` 时由操作者创建第一个 admin 账号 - 认证从一开始就是强制的，无竞争窗口 - 已有 admin 后，服...

### backend/docs/AUTO_TITLE_GENERATION.md

- Headings: 自动 Thread Title 生成功能; 功能说明; 实现方式; ⚠️ 重要：存储机制; Title 存储位置; 持久化说明; 如何启用持久化; 在 langgraph.json 同级目录创建 checkpointer.py
- Claim: 自动为对话线程生成标题，在用户首次提问并收到回复后自动触发。 使用 `TitleMiddleware` 在 `after_model` 钩子中： 1. 检测是否是首次对话（1个用户消息 + 1个助手回复） 2. 检查 state 是否已有 title 3. 默认从首条用户消息生成本地 fallback 标题，避免在流式回复结束前额外等待一次 LLM 调...

### backend/docs/BLOCKING_IO_DETECTION.md

- Headings: Blocking IO detection usage and maintenance; Static detector; Runtime detector; Maintenance workflow; Add a runtime rule; Add a runtime anchor; Current runtime coverage
- Claim: This document describes how to use and maintain DeerFlow backend blocking-IO detection for async event-loop safety. The goal is narrow: find and prevent synchronous IO from bloc...

### backend/docs/CONFIGURATION.md

- Headings: Configuration Guide; Prompt overlays; Model request admission; Config Versioning; Configuration Sections; Extensions; Recursion Limits; Models
- Claim: This guide explains how to configure DeerFlow for your environment. Operators can add instructions around existing system prompts in `config.yaml`. Empty or omitted extensions p...

### backend/docs/FILE_UPLOAD.md

- Headings: 文件上传功能; 概述; 功能特性; API 端点; 1. 上传文件; 2. 查询上传限制; 3. 列出已上传文件; 4. 删除文件
- Claim: DeerFlow 后端提供了完整的文件上传功能，支持多文件上传，并可选地将 Office 文档和 PDF 转换为 Markdown 格式。 - ✅ 支持多文件同时上传 - ✅ 可选地转换文档为 Markdown（PDF、PPT、Excel、Word） - ✅ 文件存储在线程隔离的目录中 - ✅ Agent 自动感知当前消息中附带的文件 - ✅ 支持文件...

### backend/docs/GITHUB_AGENTS.md

- Headings: GitHub Event-Driven Agents; Overview; Webhook → Fan-out → Dispatch; `preferred_thread_id = UUID5(...)` Thread Determinism; Mention-handle Precedence; GH Token Lifecycle; Thread-create Race Recovery; Follow-up Buffering While Busy (#4121)
- Claim: GitHub is a **webhook-push** channel: there is no long-polling worker. Every GitHub App / repository delivery lands at `POST /api/webhooks/github`, where it is HMAC-verified, fa...

### backend/docs/GUARDRAILS.md

- Headings: Guardrails: Pre-Tool-Call Authorization; Why Guardrails; Architecture; Four Provider Options; Option 1: Built-in AllowlistProvider (Zero Dependencies); Option 2: OAP Passport Provider (Policy-Based); Option 3: Custom Provider (Bring Your Own); my_guardrail.py
- Claim: > **Context:** [Issue #1213](https://github.com/bytedance/deer-flow/issues/1213) — DeerFlow has Docker sandboxing and human approval via `ask_clarification`, but no deterministi...

### backend/docs/IM_CHANNEL_CONNECTIONS.md

- Headings: IM Channel Connections; Architecture Overview; Conversation-scoped Custom Agents; Connect-code Flow; Single-active-owner Transfer; Provider Message Flow Once Bound; Inbound capacity and overload behavior; Sync vs Streaming Channels
- Claim: DeerFlow supports user-owned IM channel bindings for Telegram, Slack, Discord, Feishu/Lark, DingTalk, WeChat, WeCom, QQ, and Buzz. The feature reuses the existing `channels.*` r...

### backend/docs/LOOP_DETECTION.md

- Headings: Loop detection lifecycle; Run-scoped state; Decision ordering
- Claim: `LoopDetectionMiddleware` owns call-pattern detection. Its place in the chain and audit contract are documented in the [middleware guide](../packages/harness/deerflow/agents/mid...

### backend/docs/MCP_SERVER.md

- Headings: MCP (Model Context Protocol) Configuration; Setup; Copy example configuration; Stdio Working Directory; OpenViking MCP Tools; Parallel Search (optional); Routing Hints; Tool Name Prefixes
- Claim: DeerFlow supports configurable MCP servers and skills to extend its capabilities, which are loaded from a dedicated `extensions_config.json` file in the project root directory. ...

### backend/docs/MEMORY_COGNITIVE_STYLE.md

- Headings: Memory: Cognitive Style; One-line pitch; Problem; Approach (not a new subsystem); Update frequency (read vs write); How this differs from nearby concepts; Adding a new memory field (schema evolution); Verification
- Claim: Design note for contributors. Explains `user.cognitiveStyle` and related memory/prompt touchpoints. **Skills teach the agent how to do tasks; memory’s `cognitiveStyle` teaches t...

### backend/docs/MEMORY_IMPROVEMENTS.md

- Headings: Memory System Improvements; Status (As Of 2026-03-10); Current Behavior; Extraction cost gate (opt-in, off by default); Known Gap; Roadmap (Planned); Validation
- Claim: This document tracks memory injection behavior and roadmap status. Implemented in `main`: - Accurate token counting via `tiktoken` in `format_memory_for_injection`. - Facts are ...

### backend/docs/MEMORY_IMPROVEMENTS_SUMMARY.md

- Headings: Memory System Improvements - Summary; Sync Note (2026-03-10); Implemented; Planned (Not Yet Merged); Why This Sync Was Needed; Current API Shape; Verification Pointers
- Claim: This summary is synchronized with the `main` branch implementation. TF-IDF/context-aware retrieval is **planned**, not merged yet. - Accurate token counting with `tiktoken` in m...

### backend/docs/MEMORY_SETTINGS_REVIEW.md

- Headings: Memory Settings Review; Quick Review; Bulk Local Review; Minimal Manual Test; Optional Sanity Checks; Fixture Files
- Claim: Use this when reviewing the Memory Settings add/edit flow locally with the fewest possible manual steps. 1. Start DeerFlow locally using any working development setup you alread...

### backend/docs/PATH_EXAMPLES.md

- Headings: 文件路径使用示例; 三种路径类型; 1. 实际文件系统路径 (path); Python 代码中直接访问：解析当前用户上传桶，不要手写路径; 2. 虚拟路径 (virtual_path); Agent 使用 read_file 工具; Agent 使用 bash 工具; 3. HTTP 访问 URL (artifact_url)
- Claim: DeerFlow 的文件上传系统返回三种不同的路径，每种路径用于不同的场景： {DEER_FLOW_HOME}/users/{user_id}/threads/{thread_id}/user-data/uploads/document.pdf API 返回的是**绝对路径**，且上传文件按用户分桶。`{DEER_FLOW_HOME}` 表示实际的运行...

### backend/docs/README.md

- Headings: Documentation; Quick Links; Feature Documentation; Development; Getting Started; Document Organization
- Claim: This directory contains detailed documentation for the DeerFlow backend. 1. **New to DeerFlow?** Start with [SETUP.md](SETUP.md) for quick installation 2. **Configuring the syst...

### backend/docs/REPLAY_E2E.md

- Headings: Record/Replay E2E — front-back contract verification; Why; The two layers; Cross-stack scenario: multi-run render order (`multi-run-order.spec.ts`); How replay works; Record a new scenario (needs a real key — dev machine only); 1. drive the real frontend against a real-model gateway, capturing model calls; 2. stitch the capture into a fixture
- Claim: Deterministic, **key-free** end-to-end checks that a backend change can't silently break the frontend (and vice-versa). Two complementary layers, fed by a single recording. The ...

### backend/docs/RUN_EVENT_STREAM.md

- Headings: Run Event Stream; Record Envelope; Categories; Producers; Opaque Run Outputs; Consumers; Compatibility; Known Gaps
- Claim: The run event stream is DeerFlow's append-only record of what happened during an agent run. Producers write through `RunEventStore`; history, debug, subtask, memory-audit, and w...

### backend/docs/RUN_INTERACTION_POLICY.md

- Headings: Run Interaction Policy; Modes and precedence; Gateway trust boundary
- Claim: The harness resolves a `RunInteractionPolicy` from the run context and uses it as the single source for lead-agent tool visibility, clarification middleware behavior, sandbox ne...

### backend/docs/SANDBOX_MEMORY_PROFILING.md

- Headings: Sandbox Memory Profiling; What to Measure; Capture a Snapshot; Candidate Runtime Matrix; PR Guidance
- Claim: This guide records a repeatable baseline before changing the sandbox runtime. Issue #3213 reports per-sandbox memory near 1 GiB in Kubernetes. Before adding or recommending a ne...

### backend/docs/SETUP.md

- Headings: Setup Guide; Configuration Setup; Steps; Option A: Set environment variables (recommended); Optional: pin the project root when running from another directory; Option B: Edit config.yaml directly; Important Notes; Configuration File Locations
- Claim: Quick setup instructions for DeerFlow. DeerFlow uses a YAML configuration file that should be placed in the **project root directory**. 1. **Navigate to project root**: cd /path...

### backend/docs/SSO.md

- Headings: SSO / OIDC Authentication; Architecture; Configuration; Step 1: Enable OIDC in `config.yaml`; Step 2: Set the client secret as an environment variable; Step 3: Restart the backend; Provider Configuration; Per-Provider Options
- Claim: DeerFlow supports single sign-on (SSO) via any OpenID Connect (OIDC) 2.0 compliant provider. This includes Keycloak, Google Workspace, Azure AD, Okta, and many others. The OIDC ...

### backend/docs/STREAMING.md

- Headings: DeerFlow 流式输出设计; TL;DR; 为什么有两条流式路径; 为什么不能让 DeerFlowClient 复用 Gateway; LangGraph `stream_mode` 三层语义; 两套命名的由来; Gateway 路径：async + HTTP SSE; 有界历史与 `gap` 恢复
- Claim: 本文档解释 DeerFlow 是如何把 LangGraph agent 的事件流端到端送到两类消费者（HTTP 客户端、嵌入式 Python 调用方）的：两条路径为什么**必须**并存、它们各自的契约是什么、以及设计里那些 non-obvious 的不变式。 --- - DeerFlow 有**两条并行**的流式路径：**Gateway 路径**（as...

### backend/docs/THREAD_LIFECYCLE.md

- Headings: Gateway thread lifecycle invariants; Chat archive
- Claim: Read this guide before changing thread branching, regeneration, edit replay, or archive/search metadata behavior. It supplements the Gateway module guide. **Branch/regenerate ch...

### backend/docs/TITLE_GENERATION_IMPLEMENTATION.md

- Headings: 自动 Title 生成功能实现总结; ✅ 已完成的工作; 1. 核心实现文件; 2. 配置文件; 3. 文档; 4. 测试; 🎯 核心设计决策; 为什么使用 State 而非 Metadata？
- Claim: #### [`packages/harness/deerflow/agents/thread_state.py`](../packages/harness/deerflow/agents/thread_state.py) - ✅ 添加 `title: str | None = None` 字段到 `ThreadState` #### [`package...

### backend/docs/TODO.md

- Headings: TODO List; Completed Features; Planned Features; Resolved Issues
- Claim: - [x] Launch the sandbox only after the first file system or bash tool is called - [x] Add Clarification Process for the whole process - [x] Implement Context Summarization Mech...

### backend/docs/TUI.md

- Headings: DeerFlow Terminal Workbench (TUI); Install & run; Surface; Keys; Slash commands; Architecture; Web UI visibility (shared persistence); Tests
- Claim: `deerflow` is a terminal-native workbench for the DeerFlow harness. It runs **embedded** over `DeerFlowClient` — no Gateway, frontend, nginx, or Docker services required — while...

### backend/docs/blob-storage.md

- Headings: Blob storage (content-addressed, cross-instance); Contract; Backend; Configuration; Producer migration status; Interaction with checkpoint retention (#5255); Adding a backend
- Claim: Resolves the multi-instance half of [#4189](https://github.com/bytedance/deer-flow/issues/4189) item 2: two producers persist blob-shaped data outside the checkpoint payload and...

### backend/docs/checkpoint-retention-contract.md

- Headings: Checkpoint Retention Contract (DRAFT); Why a contract is needed; Data model; Protected set (MUST NOT delete without the stated compensation); Provably safe forms (validated by tests); Deletion mechanics; History fast-path interaction (wiring requirement); Item 4 note (large tool results)
- Claim: Status: **draft** — the deletion contract for #4189 item 3. No retention or deletion implementation should land before this contract (or a successor revision of it) is accepted,...

### backend/docs/extension-agent-runs.md

- Headings: Agent run control for extensions; Acquire the capability; Operations; Service-owned handoffs and lifetime
- Claim: `deerflow_extension_api.AgentRuns` is an optional, host-bound capability for creating and continuing **full Agent runs**. Runs use Gateway's normal admission, Agent assembly, to...

### backend/docs/extension-model-invocation.md

- Headings: Host model invocation for extensions; Grant and route logical roles; Call from a service; Failures and bounds
- Claim: Extension API **0.2.4** adds an optional `ExtensionRuntimeDeps.model_invoker`. An operator-authorized service can make an asynchronous, non-streaming text call using DeerFlow's ...

### backend/docs/middleware-execution-flow.md

- Headings: Middleware 执行流程; Middleware 列表; 执行流程; 时序图; 洋葱模型; 对比：真正的洋葱 vs DeerFlow 的实际情况; 真正的洋葱（如 Koa/Express）; DeerFlow 的实际情况
- Claim: `create_deerflow_agent` 通过 `RuntimeFeatures` 组装的完整 middleware 链（默认全开时）： 主 agent **14 个** middleware（`make_lead_agent`），subagent **4 个**（ThreadData、Sandbox、Guardrail、ToolErrorHan...

### backend/docs/plan_mode_usage.md

- Headings: Plan Mode with TodoList Middleware; Overview; Configuration; Enabling Plan Mode; Enable plan mode via runtime configuration; Create agent with plan mode enabled; Configuration Options; Default Behavior
- Claim: This document describes how to enable and use the Plan Mode feature with TodoList middleware in DeerFlow 2.0. Plan Mode adds a TodoList middleware to the agent, which provides a...

### backend/docs/rfc-create-deerflow-agent.md

- Headings: RFC: `create_deerflow_agent` — 纯参数的 SDK 工厂 API; 1. 问题; 对比; 2. 设计原则; Python 中的 DI 最佳实践; 分层架构; 核心约束; 3. API 设计
- Claim: 当前 harness 的唯一公开入口是 `make_lead_agent(config: RunnableConfig)`。它内部： make_lead_agent ├─ get_app_config() ← 读 config.yaml ├─ _resolve_model_name() ← 读 config.yaml ├─ load_agent_con...

### backend/docs/rfc-extract-shared-modules.md

- Headings: RFC: Extract Shared Skill Installer and Upload Manager into Harness; 1. Problem; Skill Installation; Upload Management; 2. Design Principles; Dependency Direction; Exception Strategy; 3. New Modules
- Claim: Gateway (`app/gateway/routers/skills.py`, `uploads.py`) and Client (`deerflow/client.py`) each independently implement the same business logic: **Two implementations, inconsiste...

### backend/docs/rfc-grep-glob-tools.md

- Headings: [RFC] 在 DeerFlow 中增加 `grep` 与 `glob` 文件搜索工具; Summary; Problem; Goals; Non-Goals; Why This Is Worth Doing; Proposal; 1. `glob` 工具
- Claim: 我认为这个方向是对的，而且值得做。 如果 DeerFlow 想更接近 Claude Code 这类 coding agent 的实际工作流，仅有 `ls` / `read_file` / `write_file` / `str_replace` 还不够。模型在进入修改前，通常还需要两类能力： - `glob`: 快速按路径模式找文件 - `grep`:...

### backend/docs/runtime-guidance-details.md

- Headings: Runtime implementation details
- Claim: The runtime AGENTS.md indexes these detailed contracts. Preserve their behavior when changing the corresponding code. **JSONL tail recovery** (`runtime/events/store/jsonl.py`, #...

### backend/docs/summarization.md

- Headings: Conversation Summarization; Overview; Reuse and telemetry; Todo reminders; Configuration; Trigger conditions (OR logic - any condition triggers summarization); Additional triggers (optional); - type: messages
- Claim: DeerFlow includes automatic conversation summarization to handle long conversations that approach model token limits. When enabled, the system automatically condenses older mess...

### backend/docs/task_tool_improvements.md

- Headings: Task Tool Improvements; Overview; Changes Made; 1. Removed `run_in_background` Parameter; LLM had to manage polling; Then LLM had to poll repeatedly:; Tool blocks until complete, polling happens in backend; Result is available immediately after the call returns
- Claim: The task tool has been improved to eliminate wasteful LLM polling. Previously, when using background tasks, the LLM had to repeatedly call `task_status` to poll for completion, ...

### backend/packages/harness/deerflow/AGENTS.md

- Headings: Request Trace Context (`packages/harness/deerflow/trace_context.py`); Managed Lark CLI credentials (`integrations/lark_cli.py`); Browser Progress Screenshots (`community/browser_automation/`); Embedded Client (`packages/harness/deerflow/client.py`); AIO Sandbox Network Policy; Tenki `sticky`; E2B Mount Uploads; Workspace Snapshot Cancellation (`workspace_changes/recorder.py`)
- Claim: DeerFlow's request-level correlation id — the `X-Trace-Id` header and the `deerflow_trace_id` key. Not Langfuse's trace id, not `run_id`, not the short subagent `trace_id` log l...

### backend/packages/harness/deerflow/agents/AGENTS.md

- Headings: Agent System
- Claim: `AppConfig.lead_prompt_overlay` wraps the fully rendered lead system prompt with literal operator prepend/append text. Do not template-format these extensions or source them fro...

### backend/packages/harness/deerflow/agents/memory/AGENTS.md

- Headings: Memory System
- Claim: DeerMem `prompt_prepend`/`prompt_append` wrap the first system message from `memory_update` after formatting, without touching its human data message or cached templates. This i...

### backend/packages/harness/deerflow/agents/memory/backends/README.md

- Headings: Memory Backends; Add a New Backend; Switch the Active Backend; Backend Contract; 1. The three-tier contract; 2. Return shape (critical, easy to get wrong); 3. Tier-3 hooks (contracted, no `hasattr` probing); 4. Portability (the golden rule)
- Claim: Each subfolder under `agents/memory/backends/` is a pluggable memory backend. Swap the active one by changing one line in `config.yaml` - no deer-flow core changes required. - `...

### backend/packages/harness/deerflow/agents/memory/backends/honcho/README.md

- Headings: Honcho memory backend; Configuration; api_key: $HONCHO_API_KEY   # hosted Honcho; plain-http + api_key needs allow_insecure_http: true; workspace_overrides: {}    # map specific user ids to custom workspaces; user_peer_overrides: {}    # map specific user ids to custom peer names; Multi-user isolation; Recall and search; Limitations
- Claim: Uses Honcho (self-hosted or hosted, v3 API) as DeerFlow's user-model memory store. Honcho covers the user dimension of memory — long-term user modeling, preferences, and a cross...

### backend/packages/harness/deerflow/agents/memory/backends/mem0/README.md

- Headings: mem0 memory backend; Configuration; Identity mapping; Limitations; Async execution and failure behavior
- Claim: Uses mem0 (Platform hosted API, or any API-compatible self-hosted server) as DeerFlow's memory store. Fully stateless in-process: dedup, fact extraction, and storage are server-...

### backend/packages/harness/deerflow/agents/middlewares/AGENTS.md

- Headings: Middleware Chain
- Claim: Async gate waiters own separate cross-loop futures. Cancel only their signal; notify outside the state lock. `_await_off_thread` drains started I/O across repeated cancellation....

### backend/packages/harness/deerflow/agents/middlewares/TOOL_ARTIFACTS.md

- Headings: Tool artifact middleware; Resolution; Capture; Durable projection; Delegation scope (MVP); Projection and async verification
- Claim: *(optional, if `tool_artifacts.enabled && tool_artifacts.resolve_handles_in_args`, default on)* Resolves `art_xxxxxxxx` artifact handles found in tool-call arguments to their re...

### backend/packages/harness/deerflow/agents/middlewares/TOOL_OUTPUT_PERSISTENCE.md

- Headings: Host tool-output publication; Structured tool results; Bash exit marker; Unclean shutdown; Model-bound write elision
- Claim: Host and sandbox outputs share a filename helper whose fixed-length SHA-256 suffix hashes the raw call ID and UTF-8 content with an explicit ID-length boundary. Repeating the sa...

### backend/packages/harness/deerflow/agents/task_continuity/AGENTS.md

- Headings: Task continuity
- Claim: When changing `task_note` capacity or batch receipts, follow the [continuity contract](../../../../../../docs/task-continuity.md) and run the real graph regressions in `backend/...

### backend/packages/harness/deerflow/community/aio_sandbox/AGENTS.md

- Headings: AIO Sandbox; Stdin contracts per transport; Reuse admission
- Claim: AIO has two distinct stdin contracts: the persistent `/v1/shell` transport is a PTY, while `/v1/bash` uses a subprocess pipe that stays open for writes. Keep persistent shell co...

### backend/packages/harness/deerflow/community/boxlite/README.md

- Headings: BoxLite backend; Configuration; health checks for recent releases; 0 keeps; reliability-first validation (default: 0.0); Design; Contract coverage; Status
- Claim: Runs each DeerFlow sandbox as a [BoxLite](https://github.com/boxlite-ai/boxlite) micro-VM — a daemonless, OCI-native VM with its own kernel (libkrun/KVM on Linux, Hypervisor.fra...

### backend/packages/harness/deerflow/community/brave/AGENTS.md

- Headings: Brave Search
- Claim: Web search validates the HTTP-success payload before normalizing results: `web` must be an object and `results` a list when present and non-null. Reject wrong container types wi...

### backend/packages/harness/deerflow/community/e2b_sandbox/AGENTS.md

- Headings: E2B lifecycle; File search
- Claim: Reconciliation floors active-client TTL renewals at the SDK's default timeout and twice the sum of the configured reconciliation interval and pass budget (subject to the E2B tim...

### backend/packages/harness/deerflow/community/groundroute/AGENTS.md

- Headings: GroundRoute
- Claim: Validate an HTTP-success payload before normalizing it. A top level that is not an object, a `results` container that is not a list, or a non-empty list whose entries hold no ob...

### backend/packages/harness/deerflow/community/jina_ai/AGENTS.md

- Headings: Jina web fetch
- Claim: Retries stay provider-local; default `max_retries=0` sends once. One monotonic budget bounds requests/waits; cancellation propagates. Retry 502/503/504 and connection-establishm...

### backend/packages/harness/deerflow/community/opensandbox/README.md

- Headings: OpenSandbox backend; Installation and configuration; api_key: $OPEN_SANDBOX_API_KEY; domain: localhost:8080; protocol: http; request_timeout: 30; ready_timeout: 30; use_server_proxy: false
- Claim: Runs DeerFlow sandboxes on [OpenSandbox](https://github.com/opensandbox-group/OpenSandbox), using the synchronous Python SDK behind DeerFlow's `Sandbox` and `SandboxProvider` co...

### backend/packages/harness/deerflow/community/ragflow/AGENTS.md

- Headings: RAGFlow citation snapshots; Document validation
- Claim: `knowledge_search_tool` returns native `content_and_artifact`: model-visible `[citation:N](#knowledge-<opaque-id>)` links resolve to bounded `artifact.knowledge_sources` version...

### backend/packages/harness/deerflow/community/serper/AGENTS.md

- Headings: Serper provider
- Claim: `tools.py` owns web/image search. Keep model-facing arguments unchanged. Web-only `include_domains`/`exclude_domains` use at most 10 domain-only entries per list; normalize case...

### backend/packages/harness/deerflow/community/tavily/AGENTS.md

- Headings: Tavily tools
- Claim: The client helper selects credentials by tool name: search defaults to `web_search`, fetch passes `web_fetch`, and an omitted key uses the SDK's `TAVILY_API_KEY` fallback. Keep ...

### backend/packages/harness/deerflow/community/tenki/README.md

- Headings: Tenki backend; Configuration; Design; Contract coverage; Scope: stable features only; Status
- Claim: Runs each DeerFlow sandbox as a [Tenki](https://tenki.cloud) cloud sandbox — an isolated microVM created from a stock base image, with no daemon or local virtualization to manag...

### backend/packages/harness/deerflow/config/AGENTS.md

- Headings: Configuration System; Config Schema
- Claim: `Paths.user_projects_dir()` uses `extended_length_path()` on native Windows. All document paths, including staging and retention walks, inherit the same extended drive/UNC names...

### backend/packages/harness/deerflow/extensions/AGENTS.md

- Headings: Python Extension System (Runtime and Distribution); Extension Manager Test Repositories; Full-stack contributions
- Claim: Third-party Python packages can expose an `install(registry, config)` function and be loaded, in deterministic order, from the startup-only top-level `plugins:` list in `config....

### backend/packages/harness/deerflow/mcp/AGENTS.md

- Headings: MCP System (`packages/harness/deerflow/mcp/`)
- Claim: The optional background-task `status` filter flows through `McpTaskSubmitter.list_tasks` → `McpTaskService.list_tasks` → `McpTaskRepository.list_by_thread`. SQL applies status, ...

### backend/packages/harness/deerflow/models/AGENTS.md

- Headings: MindIE XML tool arguments (`packages/harness/deerflow/models/mindie_provider.py`); MindIE XML numeric arguments (`mindie_provider.py`); Model Factory (`packages/harness/deerflow/models/factory.py`); Reasoning Capability Contract (`packages/harness/deerflow/models/reasoning.py`); CLI credential files (`packages/harness/deerflow/models/credential_loader.py`); Codex tool-call/result serialization (`packages/harness/deerflow/models/openai_codex_provider.py`); Codex SSE termination (`packages/harness/deerflow/models/openai_codex_provider.py`); Claude Code Credentials (`packages/harness/deerflow/models/credential_loader.py`)
- Claim: Parse XML tool calls from the original model response before applying escaped-newline compatibility fixes to the remaining prose. Parse JSON and Python-literal arguments from th...

### backend/packages/harness/deerflow/persistence/AGENTS.md

- Headings: Persistence lifecycle; JSON integer filters; Scheduled-task lifecycle
- Claim: Postgres bootstrap owns its session-scoped advisory lock until `pg_advisory_unlock` completes. Drain that unlock across host cancellation before leaving the SQLAlchemy connectio...

### backend/packages/harness/deerflow/persistence/migrations/AGENTS.md

- Headings: Schema Migrations (`packages/harness/deerflow/persistence/migrations/`)
- Claim: DeerFlow's application tables (`runs`, `threads_meta`, `feedback`, `users`, `run_events`, plus the `channel_*` tables) are owned by alembic via a **hybrid bootstrap** strategy. ...

### backend/packages/harness/deerflow/persistence/user/AGENTS.md

- Headings: User persistence
- Claim: `user_preferences` stores independent `(user_id, key)` rows in the shared SQL database. PATCH upserts only supplied keys in one transaction; `null` resets a field, disjoint edit...

### backend/packages/harness/deerflow/reflection/AGENTS.md

- Headings: Reflection System (`packages/harness/deerflow/reflection/`)
- Claim: - `resolve_variable(path)` - Import module and return variable (e.g., `module.path:variable_name`) - `resolve_class(path, base_class)` - Import and validate class against base c...

### backend/packages/harness/deerflow/runtime/AGENTS.md

- Headings: Stream Bridge Heartbeats; Checkpoint Channel Modes (`full` / `delta`); Referenced conversation capability; Scheduled run capabilities and goal outcomes; JSONL mutation cancellation
- Claim: Memory/Redis bridges keep startup-only `stream_bridge.heartbeat_interval_seconds`; explicit `subscribe(..., heartbeat_interval=...)` overrides it. Provider contexts drain owned ...

### backend/packages/harness/deerflow/sandbox/AGENTS.md

- Headings: Sandbox System (`packages/harness/deerflow/sandbox/`)
- Claim: Sandbox restore requires a thread ID, even forks. **Network approval policy**: `SandboxMiddleware` uses `resolve_run_interaction_policy()`. Autonomous, webhook, scheduled, legac...

### backend/packages/harness/deerflow/skills/AGENTS.md

- Headings: Skills System (`packages/harness/deerflow/skills/`); Custom skill export
- Claim: - **History JSONL**: split records on LF, not Unicode text separators; preserve U+0085/U+2028/U+2029 in saved fields. Keep malformed JSON rejected. Tests: `tests/test_skill_hist...

### backend/packages/harness/deerflow/storage/AGENTS.md

- Headings: Blob Store (`packages/harness/deerflow/storage/`)
- Claim: **What it owns**: a content-addressed store for bytes that must resolve on every instance, not just the one that wrote them (issue #4189, item 2). The address is the SHA-256 of ...

### backend/packages/harness/deerflow/subagents/AGENTS.md

- Headings: Subagent System (`packages/harness/deerflow/subagents/`)
- Claim: Result-reader acceptance runs on SQLite and, with TEST_POSTGRES_URI, an isolated PostgreSQL schema; both exercise submission, completion, reopen and ToolNode continuation withou...

### backend/packages/harness/deerflow/tools/AGENTS.md

- Headings: Tool System (`packages/harness/deerflow/tools/`)
- Claim: `artifact_registry.py` checks remote file suffixes against the complete parsed URL path without case sensitivity. Free-text URL and sandbox-path extraction strips common trailin...

### backend/packages/harness/deerflow/tracing/AGENTS.md

- Headings: Tracing System (`packages/harness/deerflow/tracing/`)
- Claim: LangSmith and Langfuse are both supported. The wiring lives in two layers: - `factory.py::build_tracing_callbacks()` — returns the LangChain `CallbackHandler` list for the provi...

### backend/packages/harness/deerflow/tui/AGENTS.md

- Headings: Terminal Workbench / TUI (`packages/harness/deerflow/tui/`)
- Claim: A terminal-native UI over the embedded harness, exposed as the `deerflow` console script (`[project.scripts]` in `packages/harness/pyproject.toml`). It is a UI shell over `DeerF...

### backend/packages/harness/deerflow/typesafe/AGENTS.md

- Headings: Shared TypeSafe client
- Claim: `typesafe/` is the transport layer for every TypeSafe (Jev) call site: the guardrail risk gate (`guardrails/typesafe.py`), and — once enabled — memory pre-screening and signal c...

### backend/packages/harness/deerflow/uploads/AGENTS.md

- Headings: Uploads
- Claim: `manager.py::list_files_in_dir` returns a best-effort snapshot under concurrent cleanup. Entry stat races (`ENOENT`, `ENOTDIR`, `ELOOP`) are skipped with debug logging. Director...

### backend/packages/harness/deerflow/utils/AGENTS.md

- Headings: Model Retry Hints; Port Allocation Bounds; Goal Objective Validation; Message Text Extraction; Agent / Tool Assembly Off-Load; Uploaded Document Summaries; Active Content MIME Types; Conversion Worker Ownership
- Claim: `retry_after.py::bounded_retry_after_ms` shares the 24-hour provider-hint ceiling between LLM middleware (numeric/date hints) and Claude backoff. Non-finite or over-limit delays...

### backend/samples/other_agent_demo/README.md

- Headings: DeerMem portability demo (other-agent integration); Three steps; from; to (your agent's vendored contract); Proof; Sample config; callbacks / should_keep_hidden_message: programmatic (cannot come from; YAML); the host factory injects them via from_config. Set programmatically
- Claim: `backends/deermem/` is a **self-contained, portable** memory backend. It has exactly **one** `from deerflow` line -- the ABC contract (`from deerflow.agents.memory.manager impor...

### backend/scripts/benchmark/context_snapshot/README.md

- Headings: Synthetic context snapshot evaluation; Run; Set CONTEXT_SNAPSHOT_API_KEY if the endpoint requires authentication.; Scope and scoring; Historical result: 2026-09-12
- Claim: This manual evaluation compares a complete, isolated handoff with an opt-in parent-context snapshot. It contains eight synthetic JSON/Python tasks and 51 predeclared checks. No ...

### backend/scripts/benchmark/deermem_eviction/README.md

- Headings: DeerMem Capacity-Eviction Evaluation; Current scope; Pinned inputs; expected: 821a2034d219ab45846873dd14c14f12cfe7776e73527a483f9dac095d38620c; Commands; Deterministic reconstruction; Deterministic grading; Output contract
- Claim: This directory makes the controlled comparison behind DeerMem's opt-in `hybrid-v1` capacity policy reproducible. It depends on [deer-flow#4789](https://github.com/bytedance/deer...

### backend/scripts/benchmark/deermem_scope_isolation/README.md

- Headings: DeerMem scope-isolation benchmark; Protocol; Offline run (default); Explicit live run; Reproducibility and resume integrity
- Claim: This is a focused regression benchmark for DeerMem's long-term-memory safety boundary. It is not a model leaderboard. It measures two separate concerns: 1. **Semantic model qual...

### backend/scripts/benchmark/scheduled_tasks/README.md

- Headings: Synthetic scheduled-work Gateway pilot; Required pilot configuration; Commands; What each phase proves; Recovery and delivery limits
- Claim: This runs the production Gateway, scheduler, conversation tools, goal worker, sandbox artifact tools and persistence. It does not provide another agent loop. All sources and exp...

### backend/tests/AGENTS.md

- Headings: Backend Tests; PostgreSQL batch fixtures; Real Compose tests; Jina client options; Claude provider tests; User repository ordering; Router auth fixtures; Lark CLI blocking-I/O fixtures
- Claim: Backend tests must preserve the runtime invariants they exercise without changing production execution topology. `test_skills_custom_multiprocess.py` starts independent spawn wo...

### backend/tests/monocle/README.md

- Headings: DeerFlow behavioural tests (Monocle Test Tools); How this is meant to be used; Layers; Layout; The committed trace; Run; from the repo root; offline — no network, no keys; the live tests skip unless opted in
- Claim: Trace-based tests for DeerFlow. Monocle records each run as a structured trace (the agent invocation, every tool call, token usage, timings), and these tests assert against that...

### deploy/helm/deer-flow/README.md

- Headings: DeerFlow Helm Chart; Prerequisites; Install the published chart (GHCR); 1. Build & push images (custom builds only); backend - build with the `postgres` extra so multi-replica deploys can use; shared Postgres (matches the published image); frontend; provisioner
- Claim: Deploys the full DeerFlow stack to Kubernetes: **gateway** (backend + embedded LangGraph runtime), **frontend** (Next.js), **nginx** (internal reverse proxy preserving the compo...

### docker/lark-cli-broker/README.md

- Headings: lark-cli broker image (Pattern B); Two modes, one image; Build; Wiring it into the provisioner; Broker HTTP contract (loopback); Stdin and execution logs; No file I/O relative to the sandbox cwd; Optional subcommand denylist (hardening)
- Claim: This image implements **Pattern B** (issue #4338): instead of mounting the per-user Lark credential directories into the sandbox (Pattern A still does), a long-running **sidecar...

### docker/lark-cli-init/README.md

- Headings: lark-cli init image (Pattern A); What it does; Build; Wiring it into the provisioner
- Claim: This image provisions the sandbox `lark-cli` runtime binary into a Kubernetes sandbox Pod via an **init container + shared `emptyDir`**, instead of the Gateway downloading Linux...

### docker/provisioner/README.md

- Headings: DeerFlow Sandbox Provisioner; Architecture; How It Works; Requirements; Enable Kubernetes in Docker Desktop; Enable Kubernetes in OrbStack; API Endpoints; `GET /health`
- Claim: The **Sandbox Provisioner** is a FastAPI service that dynamically manages sandbox Pods in Kubernetes. It provides a REST API for the DeerFlow backend to create, monitor, and des...

### docs/ARCHITECTURE.md

- Headings: DeerFlow Architecture; 1. What DeerFlow Is; 2. Service Topology; 3. Backend: Harness / App Split; Agent runtime path; State, tools, sandbox; 4. Frontend: Stateful Chat over LangGraph SDK; 5. Cross-Cutting Subsystems
- Claim: This document is the **top-level architecture overview** for DeerFlow. It explains the "big picture" — how the services, layers, and cross-cutting subsystems fit together — and ...

### docs/CODE_CHANGE_SUMMARY_BY_FILE.md

- Headings: 代码更改总结（按文件 diff，细到每一行）; 一、后端; 1. `backend/CLAUDE.md`; 2. `backend/packages/harness/deerflow/agents/lead_agent/prompt.py`; Add subagent reminder to critical_reminders if enabled; 3. `backend/app/gateway/routers/artifacts.py`; Encode filename for Content-Disposition header (RFC 5987); if `download` query parameter is true, return the file as a download
- Claim: 基于 `git diff HEAD` 的完整 diff，按文件列出所有变更。删除/新增文件单独说明。 --- @@ -156,7 +156,7 @@ FastAPI application on port 8001 with health check at `GET /health`. -| **Artifacts** (`/api/threads/{...

### docs/OPENVIKING.md

- Headings: OpenViking memory backend; Current scope; Authentication boundary; Configure DeerFlow; Start the services; Identity and session mapping; Retry and failure behavior
- Claim: DeerFlow can use a remote OpenViking server as an optional long-term memory backend. DeerMem remains the default. The OpenViking backend uses the maintained [`langchain-openviki...

### docs/RELEASE_SOP.md

- Headings: DeerFlow Release SOP; 0. Prerequisites; 1. Gather the complete PR + contributor list; All merged milestone PRs: number, title, labels; Sorted PR numbers + total count; Contributors (login -> display name via `gh api users/<login>`); 2. Update `CHANGELOG.md` (Keep a Changelog); then splice /tmp/refs.txt after the release-tag line
- Claim: Standard operating procedure for cutting a DeerFlow release. The goal is a self-consistent release: **every merged milestone PR is documented in English *and* Chinese**, a curat...

### docs/SKILL_NAME_CONFLICT_FIX.md

- Headings: 技能名称冲突修复 - 代码改动文档; 概述; 问题描述; 原始问题; 根本原因; 解决方案; 核心思路; 设计原则
- Claim: 本文档详细记录了修复 public skill 和 custom skill 同名冲突问题的所有代码改动。 **状态**: ⚠️ **已知问题保留** - 同名技能冲突问题已识别但暂时保留，后续版本修复 **日期**: 2026-02-10 --- 当 public skill 和 custom skill 有相同名称（但技能文件内容不同）时，会出现以...

### docs/agents/maintainer-orchestrator-design.md

- Headings: DeerFlow Maintainer Orchestrator — design notes; What problem it solves; The safety model: comment-only; When it posts: a deliberately high bar; How it treats existing coverage; Principles that shape the review; What it deliberately does not do; How a maintainer runs it
- Claim: This document explains the *thinking* behind the `deerflow-maintainer-orchestrator` skill: what it is for, the boundaries that make it safe to run, and the principles that shape...

### docs/capability-center.md

- Headings: Capability Center integration contract; Ownership; Bundled business integrations; Add a catalog entry; API and state; Compatibility and execution; Static demo; Validation
- Claim: Capability Center is a discovery and configuration layer over DeerFlow's existing MCP, Lark CLI, and skill services. It uses the existing administrator/user roles. Agent selecti...

### docs/database-forward-revision-recovery.md

- Headings: Recovering the original thread-incarnation database revision; Offline migration
- Claim: The Projects build requires `projects` and `threads_meta.project_id`. An older deployment may have stamped `0019_thread_incarnations` on a database containing only `0018_oauth_i...

### docs/experiments/task-continuity-20260912/CASE_NOTES.md

- Headings: 已核对的案例; 预算中断可以恢复; 任务笔记确实保住过已验证的新结论; 向量提高证据召回，也可能没有提高最终答案; 混合排序挤出关键原文; 有正确原文仍读错日期；官方 turn 标签也存在缺口; 向量补回关键词完全没找到的年龄事实; 自动裁判接受了前后矛盾的答案
- Claim: 这些说明只解释已保存的轨迹，不修改原评分，也不据此调整检索或提示词。 - `task-correction-01` 的 D 组找到了正确纠正记录。它最初多写了一个 `notes` 字段，校验失败；随后删掉多余字段，文件已经完全正确，却在下一次校验前触及原累计上下文预算。恢复原缓存轨迹后补上校验，验收通过。 - `task-units-02` 的 D 组...

### docs/experiments/task-continuity-20260912/CONCLUSIONS.md

- Headings: 本轮结论与建议; 建议做成什么; 结论的适用范围
- Claim: **先做可选的任务笔记与原文回查增强；这轮结果不支持默认引入向量检索，也不支持全面替换现有摘要。** 公开 LongMemEval-S 预选 42 例，40 例得到完整四组结果。另两例 `60d45044`、`60472f9c` 在重复低负载恢复后仍出现接口断连，按运行失败保留。有效样本上的自动评分为：A 摘要 **9/40（22.5%）**，B 加笔...

### docs/experiments/task-continuity-20260912/README.md

- Headings: Session continuity evaluation; Published package; Runtime; Run; What this can establish
- Claim: Historical research prototype motivating the opt-in implementation in this PR. The measurements below were completed before production integration; they are not measurements of ...

### docs/experiments/task-continuity-20260912/REPORT.md

- Headings: DeerFlow 任务接续增强：真实模型对照实验; 本轮结论与建议; 建议做成什么; 结论的适用范围; 1. 公开历史问答; 2. 三次压缩后的受控执行; 目标从开始就明确的接续对照; 产物正确与完成验证分开统计
- Claim: 作者：Aari > 本文保留集成前的原型实验结果；本 PR 的运行时实现及验证见 [功能说明](../../task-continuity.md)。公开包只包含脚本、方法、图表和去除原文的逐例分数；原始响应、数据集和本地工作区不随 PR 分发。 **状态：实验及预算续跑已结束；运行失败与模型质量分开列示。** **先做可选的任务笔记与原文回查增强；这轮...

### docs/experiments/task-continuity-20260912/VALIDATION.md

- Headings: Implementation validation; Latest follow-up: review 5185922053; Earlier follow-up: review 5185827206; Earlier validation before the second review; Review regressions
- Claim: These checks exercise the implementation and review fixes. They are separate from the historical prototype A/B/C/D quality scores. Malformed history metadata is validated by all...

### docs/full-stack-plugins.md

- Headings: Full-stack plugin contributions; What users see; Registration and execution; Browser API; Packaged resources (`assets-v1`); Authorization of plugin surfaces; Middleware-declared tools; Guarding a contributed management route
- Claim: A deployment-installed Python extension can register a `PluginContribution` with optional browser code, authenticated backend actions and model tools. This extends the existing ...

### docs/plans/2026-04-01-langfuse-tracing.md

- Headings: Langfuse Tracing Implementation Plan; Task 1: Add failing tracing config tests; Task 2: Add failing callback factory and model attachment tests; Task 3: Wire dependency and docs; Task 4: Run broader regression checks
- Claim: **Goal:** Add optional Langfuse observability support to DeerFlow while preserving existing LangSmith tracing and allowing both providers to be enabled at the same time. **Archi...

### docs/plans/2026-07-10-pluggable-authorization-implementation-notes.md

- Headings: 可插拔授权系统实施记录; 每个 RFC PR 的必读要求; Authorization RFC 连续性确认; 信息优先级; Phase 0：已合并基线; PR #4127 多轮修改的原因; 所有后续阶段必须保持的不变量; Phase 1 实施前确认
- Claim: pluggable authorization system (RFC #4063). An English summary is provided below for navigation; the detailed content is in Chinese. --> > **English summary:** This is the cumul...

### docs/plans/2026-07-10-pluggable-authorization-rfc.md

- Headings: RFC: Pluggable Fine-Grained Authorization; TL;DR; 1. Background: what exists today; 1.1 Authentication (who you are) — solid; 1.2 Authorization (what you can do) — a placeholder; 1.3 The GuardrailProvider (the integration point); 1.4 The tool assembly pipeline (where visibility is decided); 1.5 Identity flow into the run — reliable for the web, gap for channels
- Claim: Added to the PR per @WillemJiang's request for design tracking. --> > **实施连续性要求：** 每个阶段的 PR 都必须阅读并更新 > [可插拔授权系统实施记录](2026-07-10-pluggable-authorization-implementation-notes.md)。...

### docs/plans/2026-07-15-authz-phase1a-implementation-plan.md

- Headings: Phase 1A 实施计划：可信身份链路与内置 RBAC 核心; 1. 目标与拆分; 2. 固定架构与职责; 3. 全局语义锁; 3.1 Principal; 3.2 服务端身份字段; 3.3 RBAC; 3.4 fail-closed
- Claim: > RFC：`docs/plans/2026-07-10-pluggable-authorization-rfc.md` > 前情记录：`docs/plans/2026-07-10-pluggable-authorization-implementation-notes.md` > Phase 0 基线：PR #4127 / `1300c6d3` Ph...

### docs/plans/2026-07-15-authz-phase1a1-trusted-principal-implementation-plan.md

- Headings: Phase 1A-1 实施计划：可信 Principal 链路; 1. 目标; 2. 固定安全契约; 2.1 Principal; 2.2 Gateway 可信边界; 2.3 传递语义; 3. TDD 实施顺序; Step 1：新增 Principal builder 失败测试
- Claim: 建议 PR 标题： feat(authz): propagate trusted authorization principal context 前置基线：PR #4127 / `1300c6d3` 总计划：`docs/plans/2026-07-15-authz-phase1a-implementation-plan.md` 实施记录：`docs/p...

### docs/plans/2026-07-17-remember-login.md

- Headings: 管理页面保持登录方案; 方案结论; 背景与目标; 链路设计; 前端; Gateway; 沙盒边界; 非目标
- Claim: 实现“保持登录”，不实现“记住密码”。密码只随登录表单提交一次，不写入浏览器存储、日志、运行上下文或沙盒环境。 本次改造沿用现有认证链路： `LoginPage -> /api/v1/auth/login/local -> LocalAuthProvider -> JWT -> HttpOnly access_token cookie -> AuthP...

### docs/plans/2026-08-19-mcp-artifact-handle-registry.md

- Headings: Spec: MCP 工具产物持久化句柄注册表（Artifact Handle Registry）; 1. 问题描述; 现有机制的断点（代码级）; 2. 目标; 3. 非目标; 4. 数据模型; 4.1 `ArtifactEntry`（新增 TypedDict，放 `thread_state.py`）; 4.2 `ThreadState` 新增字段
- Claim: **关联 Issue**: [#4676](https://github.com/bytedance/deer-flow/issues/4676) — [feat] 为 MCP 协议工具产物提供通用句柄以支持可持久化引用 **范围**: 后端 harness + 中间件 + state schema + 前端渲染 **依赖**: 无（独立实现，可与 #...

### docs/plans/2026-08-24-subagent-batch-capacity-implementation.md

- Headings: Subagent Capacity and Durable Batch Execution; Status; Problem; Stage 1: one process-wide execution capacity; Startup-only configuration; One value across prompt, middleware, and executor; Admission behavior; Direct factory ownership
- Claim: Implemented in the `feat/subagent-batch-capacity` worktree. The implementation keeps ordinary delegation bounded, makes its advertised and real concurrency consistent, and adds ...

### docs/plans/2026-09-23-reasoning-capability-contract.md

- Headings: Model Provider Reasoning Capability Contract; Status; Problem; Contract; Configuration shape; Normalized contract; Request resolution; Factory behavior
- Claim: Implemented in the `feat-model-provider-reasoning-capability-contract` worktree for issue #5073, following option 1 from the issue discussion: an optional structured, mapping-va...

### docs/plans/OPENVIKING_HTTP_MEMORY_INTEGRATION.md

- Headings: OpenViking HTTP Memory Backend 接入实现方案; 1. 背景; 2. 目标; 3. 非目标; 4. 总体架构; 5. 代码布局; 5.1 `__init__.py`; 5.2 `config.py`
- Claim: > 状态：设计稿 > 范围：通过 HTTP 将 OpenViking 接入为 DeerFlow 的可选 `MemoryManager` 后端 > 默认行为：DeerMem 继续作为默认后端，OpenViking 需要显式启用 > 非目标：本文不包含代码实现、DeerMem 存量数据迁移、OpenViking 服务端开发 DeerFlow 已提供可插拔的...

### docs/plans/STORAGE_REWRITE_CHANGES.md

- Headings: DeerMem Storage Rewrite: Implementation Guide; 1. What changed on disk; 2. Main files; `deer_mem.py`; `deermem/core/paths.py`; `deermem/core/storage.py`; `deermem/core/updater.py`; `app/gateway/routers/memory.py`
- Claim: > This document explains the final PR behavior for readers who are new to the codebase. It complements `STORAGE_REWRITE_PLAN.md`: the plan explains the intended contract; this f...

### docs/plans/STORAGE_REWRITE_PLAN.md

- Headings: DeerMem Storage Rewrite Plan; 1. Objective; 2. Non-goals; 3. Canonical layout; 4. Scope rules; 5. Fact document; Runtime constraint; 6. Compatibility boundary
- Claim: > Status: implemented in PR #4279 and retained as the design record requested during review. > Scope: file-backed storage only. Retrieval ranking, embeddings, project scope, and...

### docs/skill-usage-ui.md

- Headings: Skill usage in answer details; Capture and history; Presentation
- Claim: The answer toolbar shows skills that the lead agent actually loaded during that answer's run. A skill catalog entry, a mention in answer text, a failed read, or a client-supplie...

### docs/superpowers/plans/2026-04-10-event-store-history.md

- Headings: Event Store History — Backend Compatibility Layer; Real Data Alignment Analysis; File Structure; Task 1: Add `_get_event_store_messages` helper to `threads.py`; human_message: id is None (same as real data); ai_tool_call: id is set by LLM; tool_result: id is None (same as real data); ai_message: id is set by LLM
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-06-08-minimax-generation-providers.md

- Headings: MiniMax 接入生成类 Skill 实施计划; File Structure; Task 0: 测试加载器; Task 1: image-generation 接入 MiniMax; Providers (Gemini / MiniMax); Task 2: video-generation 接入 MiniMax; MiniMax video uses resolution/duration, not aspect_ratio; aspect_ratio ignored.; Providers (Gemini / MiniMax)
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-06-12-telegram-streaming.md

- Headings: Telegram Streaming Output Implementation Plan; Task 1: Capability flip — Telegram reports streaming support; ---------------------------------------------------------------------------; Telegram streaming tests; ---------------------------------------------------------------------------; Task 2: Stream state infrastructure + placeholder registration; Indirection so tests can patch the clock without touching the global time module.; stream_key ("chat_id:thread_ts") -> state of the in-flight streamed
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-06-19-guardrail-request-attribution.md

- Headings: GuardrailRequest 运行时用户上下文与归因字段补充 — 实现计划; 文件结构; Task 1: Gateway 注入 authenticated user context; Task 2: GuardrailRequest 扩展字段; Task 3: GuardrailMiddleware 从 runtime context 填充; Task 4: 文档与示例; Task 5: 测试与验证; 自检清单
- Claim: **Goal:** 将服务端认证后的用户上下文和运行时归因信息传递给 `GuardrailProvider`：`user_id`、`user_role`、`oauth_provider`、`oauth_id`、`thread_id`、`run_id`、`tool_call_id`。 **Architecture:** Gateway 只从 `reque...

### docs/superpowers/plans/2026-06-29-community-browserless-web-capture.md

- Headings: Community Browserless Web Capture Implementation Plan; Technical Design; User-Facing Tool Contract; token: $BROWSERLESS_TOKEN; wait_for_selector: main; wait_for_selector_timeout_ms: 5000; wait_for_timeout_ms: 1000; best_attempt: true
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:test-driven-development for code changes. Commit and push steps are intentionally deferred because this local bran...

### docs/superpowers/plans/2026-07-01-scheduled-tasks-mvp.md

- Headings: Scheduled Tasks MVP Implementation Plan; Global Constraints; Task 1: Add scheduler config and persistence skeleton; Task 2: Add Alembic migration and repository CRUD; Task 3: Implement schedule parsing and next-run computation; Task 4: Add due-claim, lease management, and task-run status updates; Task 5: Add scheduler service and shared run-launch helper; Task 6: Add scheduled-task API routes and app wiring
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-07-02-read-before-write-gate.md

- Headings: Read-Before-Write Gate (ReadBeforeWriteMiddleware) Implementation Plan; Global Constraints; Task 1: Extract `read_current_file_content` helper in sandbox tools; Custom mount paths are resolved by LocalSandbox._resolve_path(); Task 2: `ReadBeforeWriteMiddleware` — mark stamping + sync write gate; -- wrap_tool_call ------------------------------------------------; -- gate ----------------------------------------------------------; write_file creates the file; str_replace surfaces its own error.
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-07-24-issue-4416-run-metadata-secrets.md

- Headings: Issue 4416 Run Metadata Secrets Implementation Plan; Global Constraints; Task 1: Central legacy metadata policy; Task 2: Unified run admission before persistence; Task 3: Non-mutating historical API hiding; Task 4: Supported MCP carrier and operator documentation; Task 5: Full verification and handoff
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-07-31-frontend-performance-baseline-delivery.md

- Headings: Frontend Performance Baseline and Delivery Implementation Plan; Task 1: Add a reproducible route-asset budget gate; Task 2: Enable safe Nginx compression; Task 3: Lazy-load landing case-study images; Task 4: Replace request-derived mock filesystem traversal with a static manifest; Final verification
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-07-31-frontend-performance-bundle-static.md

- Headings: Frontend Bundle and Static Rendering Implementation Plan; Task 1: Make the root layout static and scope route styles; Task 2: Scope locale dictionaries by route boundary; Task 3: Split interaction-only settings and workspace panels; Task 4: Split CodeMirror languages and render one Shiki tree; Task 5: Close the bundle/static gate
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-07-31-frontend-performance-chat-runtime.md

- Headings: Frontend Chat Runtime Performance Implementation Plan; Task 1: Index message-derived state once per update; Task 2: Make history pagination focus-stable; Task 3: Bound streaming Markdown work; Task 4: Virtualize long message histories; Task 5: Normalize, bound, and virtualize chat navigation; Final verification
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-07-31-frontend-performance-live-artifacts.md

- Headings: Frontend Live Media and Artifact Performance Implementation Plan; Task 1: Pause decorative work when hidden or offscreen; Task 2: Negotiate binary Browser Live frames in the backend; Task 3: Present Browser Live frames outside React state cadence; Task 4: Serve text artifacts with byte ranges; Task 5: Bound the client artifact preview to 1 MiB; Task 6: Documentation and cross-stack verification
- Claim: > **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Step...

### docs/superpowers/plans/2026-09-13-projects-mvp-phase2-implementation-plan.md

- Headings: DeerFlow Projects MVP — Phase 2 Implementation Plan; Slice map (spec §16 dependency order); Slice A — Instructions injection; Slice B — Shelf storage, delivery, reads; Slice C — Promotion and file browser; Slice D — Trash completion; Slice E — Frontend completion; Documentation (land with owning slice)
- Claim: **Date**: 2026-09-13 **Spec revision note (2026-09-13)**: the spec's injection strategy was replaced mid-implementation. §7.2 is now **latest-only, request-scoped**: `DynamicCon...

### docs/superpowers/specs/2026-04-11-runjournal-history-evaluation.md

- Headings: RunJournal 替换 History Messages — 方案评估与对比; 1. 问题与数据核对; 2. 候选方案; 3. Claude 自评 vs Codex 独立评估; 3.1 一致结论; 3.2 Claude 的独立观察; 3.3 Codex 的关键补充（Claude 遗漏）; 3.4 Codex 未否决但提示的次要点
- Claim: **日期**：2026-04-11 **分支**：`rayhpeng/fix-persistence-new` **相关 plan**：[`docs/superpowers/plans/2026-04-10-event-store-history.md`](../plans/2026-04-10-event-store-history.md)（尚未落地...

### docs/superpowers/specs/2026-04-11-summarize-marker-design.md

- Headings: Summarize Marker in History — Design & Verification; 1. Goal; 2. Investigation findings; 2.1 Today's state: zero middleware records; 2.2 Why: two dead code paths in `journal.py`; 2.3 Concrete evidence of summarize running unlogged; 3. Approaches considered; A. Subclass `SummarizationMiddleware` and dispatch a custom event
- Claim: **Date**: 2026-04-11 **Branch**: `rayhpeng/fix-persistence-new` **Status**: Design approved, implementation deferred to a follow-up PR **Depends on**: [`2026-04-11-runjournal-hi...

### docs/superpowers/specs/2026-06-08-minimax-generation-providers-design.md

- Headings: MiniMax 接入生成类 Skill — 设计文档; 1. 目标; 2. 背景与现状; 3. Provider 选择机制（已和用户确认）; 4. MiniMax 接口对接细节; 4.1 图像 `POST /v1/image_generation`（同步）; 4.2 视频（异步三步）; 4.3 播客 TTS `POST /v1/t2a_v2`（同步）
- Claim: - 日期：2026-06-08 - 分支：`worktree-feat-minimax-generation` - 参考：MiniMax 开放平台 API（https://platform.minimaxi.com/docs/api-reference） 1. 在现有 `image-generation`、`video-generation`、`pod...

### docs/superpowers/specs/2026-06-12-telegram-streaming-design.md

- Headings: Telegram 流式输出设计; 背景与目标; 方案选型; 改动 1 — `backend/app/channels/manager.py`; 改动 2 — `backend/app/channels/telegram.py`; 流式状态; 占位消息复用; `send()` 按 `is_final` 分流
- Claim: 日期：2026-06-12 分支：`feat/telegram-streaming` 状态：已与用户确认 Telegram 通道目前完全不流式：`ChannelManager._handle_chat()` 走 `client.runs.wait()` 阻塞路径，agent 跑完后一次性 `send_message` 发出最终文本。用户先看到 "Wor...

### docs/superpowers/specs/2026-06-13-deerflow-tui.md

- Headings: DeerFlow TUI - Product and Engineering Spec; Problem Statement; Solution; Runtime and Session Persistence; v1 decision; Reference Behavior From Hermes; User Stories; TUI Surface
- Claim: **Date**: 2026-06-13 **Status**: Draft, ready for RFC discussion **Scope**: Hermes-like terminal UI for the DeerFlow 2.0 harness **Revision 2026-06-24**: Reworked runtime owners...

### docs/superpowers/specs/2026-06-19-guardrail-request-attribution-design.md

- Headings: GuardrailRequest 运行时用户上下文与归因字段补充; 概述; 背景; 改动范围; 设计; GuardrailRequest 字段; 字段来源; 收益
- Claim: 为 `GuardrailRequest` 补充可选的运行时用户上下文和工具调用归因字段，使可插拔 `GuardrailProvider` 能访问 DeerFlow 已认证用户、外部身份映射、run/thread/tool-call 定位信息。 本设计不新增治理系统、不定义统一 policy schema，也不改变默认 allow/deny 行为。它只把...

### docs/superpowers/specs/2026-07-01-scheduled-tasks-mvp-design.md

- Headings: DeerFlow Scheduled Tasks MVP Design; Problem Statement; Solution Summary; Explicit Non-Goals; Chosen Architecture; Why this shape; Resource Model; User Stories
- Claim: **Date**: 2026-07-01 **Status**: Approved for implementation **Scope**: First-class scheduled-task management for DeerFlow web workspace --- DeerFlow main does not ship a real s...

### docs/superpowers/specs/2026-07-02-read-before-write-gate-design.md

- Headings: 产物层：修改已存在文件前强制先读（ReadBeforeWriteMiddleware）— 设计; 行为规格; 实现结构; 已知边界（记录，不在本项处理）; 测试（TDD，`backend/tests/test_read_before_write_middleware.py`）
- Claim: 对应 issue #3857（parent）产物层子项。目标：用确定性的"版本闸"逼 agent 在修改已存在文件前看见文件现状，治"只追加、不回读"导致的重复产出（Case 2），并作为通用的文件写入新鲜度护栏。 定位（采纳 issue 评论区共识）：本机制是**新鲜度护栏**——保证"agent 读过当前版本"，不保证"最终产物无语义重复"。后者（...

### docs/superpowers/specs/2026-07-02-scheduled-tasks-i18n-recipes-design.md

- Headings: Scheduled-tasks page: full i18n + recipe templates — design; Background; Goals; Design; Full i18n; Recipe templates; Wiring into the create form; Verification
- Claim: Follow-up to the preset-driven schedule form (commit `1ca27a73`). Scope kept front-end only and backend-contract-free so it stays reviewable inside PR #3898. The preset form i18...

### docs/superpowers/specs/2026-07-31-frontend-performance-remediation-design.md

- Headings: Frontend performance remediation — design; Background; Goals; Non-goals; Current finding matrix; Design; 1. Reproducible performance baseline and budgets; 2. Delivery, static rendering, and route ownership
- Claim: A production-oriented audit identified seventeen performance risks across the default DeerFlow web deployment. The audit was originally performed against an older `main`, so thi...

### docs/superpowers/specs/2026-09-12-projects-mvp-phase2-design.md

- Headings: DeerFlow Projects MVP — Phase 2 Design (Instructions, Documents, Promotion, Trash); Problem Statement; Solution Summary; Explicit Non-Goals; Chosen Architecture; Principles carried from Phase 1 (unchanged); Phase-2-specific decisions (the ones a reviewer should attack first); User Stories
- Claim: **Date**: 2026-09-12 **Status**: Draft for review (RFC v2, [issue #5160](https://github.com/bytedance/deer-flow/issues/5160)) **Phase**: Phase 2 of the Projects MVP. Phase 1 (or...

### docs/superpowers/specs/2026-09-12-projects-mvp-phase2-design.zh.md

- Headings: DeerFlow 项目 MVP —— Phase 2 设计（项目说明、文档、提升、回收站）; 问题陈述; 方案概要; 明确的非目标; 选定的架构; 从 Phase 1 继承的原则（不变）; Phase 2 特有的决策（评审者应优先攻击的那些）; 用户故事
- Claim: **日期**: 2026-09-12 **状态**: 评审草案（RFC v2，[issue #5160](https://github.com/bytedance/deer-flow/issues/5160)） **阶段**: 项目 MVP 的 Phase 2。Phase 1（组织）已作为 #5265 落地（`cfda885c`..`e2f2afde`...

### docs/task-continuity.md

- Headings: Task continuity after compaction; Matching excerpts and character offsets; Storage and lifecycle; Evidence
- Claim: Enable this optional feature to keep short task notes and recover details from messages removed by successful context compaction: task_continuity: enabled: true max_batches: 32 ...

### examples/deerflow-extension-agent-teams/AGENTS.md

- Headings: Agent teams extension
- Claim: Keep team business logic in this independently installable package. Runtime code imports public `deerflow_extension_api` contracts and its declared LangChain dependency; do not ...

### examples/deerflow-extension-agent-teams/README.md

- Headings: Agent teams — independent plugin; Install; Use; Execution and limits; Validation
- Claim: Let existing full Custom Agents work together through directed requests and a shared record. Each member gets its own persistent conversation, bound to its Custom Agent name. Or...

### examples/deerflow-extension-bookmarks/README.md

- Headings: Conversation bookmarks — independent plugin example; Deployment; Try it; Boundaries; Host contract exercised; Browser resource layout
- Claim: A Pi-inspired example of a complete feature in one package: save the last visible assistant answer from a conversation menu, then search, rename, read or delete bookmarks in its...

### examples/deerflow-extension-bookmarks/THIRD_PARTY_NOTICES.md

- Headings: Reference and license
- Claim: This is an original Python/browser adaptation inspired by Pi's `bookmark.ts`. Pi marks the last assistant entry with a label for its session tree. This example adds a web librar...

### examples/deerflow-extension-example/README.md

- Headings: DeerFlow extension example; What it demonstrates; Tool statistics; Run the package tests; Install and load it in DeerFlow; Packaging entry point; Package layout
- Claim: This directory is a compact, standalone Python package showing five of the eight DeerFlow extension contribution kinds: middleware, task lifecycle, system-model observer, servic...

### examples/deerflow-extension-jev-classify/README.md

- Headings: Text-list classification — independent plugin example; Install and configure; How the agent uses it; Behavior and limits; Verify
- Claim: An opt-in Python plugin that gives the agent a `classify_texts` tool: label a list of short texts with one of the categories supplied in the call. The backend is chosen by the d...

### examples/deerflow-extension-jev-context/README.md

- Headings: Jev context pruning — independent plugin example; Install and configure; Behavior and limits; Prompt caches and cost; Verify
- Claim: An opt-in Python plugin that uses Jev to shorten clearly obsolete **read-only tool results** before DeerFlow's normal summarization. It uses the existing extension API 0.2.2 and...

### examples/deerflow-extension-jev-screening/README.md

- Headings: Advisory screening of fetched content — extension example; Install and configure; How it works; Runtime behavior and boundaries; What this example exercises in the extension API; Validation
- Claim: An opt-in packaged extension for [RFC #5737](https://github.com/bytedance/deer-flow/issues/5737). It contributes one `AgentMiddleware` through the extension API, placed at `TOOL...

### frontend/AGENTS.md

- Headings: AGENTS.md; Project Overview; Core dependencies; Commands; Architecture; Source Layout (`src/`); Code Style; Environment
- Claim: This file provides guidance to AI coding agents (Claude Code, Codex, and others) when working with the DeerFlow frontend. It is the source of truth; the sibling `CLAUDE.md` impo...

### frontend/CLAUDE.md

- Headings: CLAUDE.md
- Claim: The frontend agent guidance lives in [AGENTS.md](AGENTS.md) so it is shared across coding agents (Claude Code, Codex, and others). Claude Code imports it below. @AGENTS.md

### frontend/README.md

- Headings: DeerFlow Frontend; Tech Stack; Quick Start; Prerequisites; Installation; Install dependencies; Copy environment variables; Edit .env with your configuration
- Claim: Like the original DeerFlow 1.0, we would love to give the community a minimalistic and easy-to-use web interface with a more modern and flexible architecture. - **Framework**: [...

### frontend/docs/conversation-ui-ownership.md

- Headings: Conversation UI ownership
- Claim: This expands the ownership index in `frontend/src/AGENTS.md`. - `src/components/workspace/model-picker-content.tsx` owns the compact model list, favorite grouping, and the ancho...

### frontend/public/demo/threads/3823e443-4e2b-4679-b496-a9506eae462b/user-data/outputs/fei-fei-li-podcast-timeline.md

- Headings: Dr. Fei-Fei Li: Recent Podcast Appearances Timeline (Last 6 Months); Overview; Timeline of Recent Podcast Appearances; January 15, 2025 - **Possible Podcast** (with Reid Hoffman and Aria Finger); August 15, 2025 - **Firing Line (PBS)**; November 16, 2025 - **Lenny's Podcast**; November 25, 2025 - **Masters of Scale Summit**; December 9, 2025 - **The Tim Ferriss Show** (#839)
- Claim: Dr. Fei-Fei Li, often called the "Godmother of AI," has been actively appearing on major podcasts discussing the future of artificial intelligence, spatial intelligence, human-c...

### frontend/public/demo/threads/7f9dc56c-e49c-4671-a3d2-c492ff4dce0c/user-data/outputs/leica-master-photography-article.md

- Headings: The Leica Master's Eye: Capturing the Decisive Moment in the Age of AI; The Enduring Legacy of Leica Street Photography; The Leica Aesthetic: More Than Just Gear; 1. The Decisive Moment Philosophy; 2. Rangefinder Discretion; 3. Lens Character; 4. Film-Like Aesthetic; Three AI-Generated Leica Masterpieces
- Claim: _By DeerFlow 2.0 | January 28, 2026_ For nearly a century, the name Leica has been synonymous with street photography excellence. From Henri Cartier-Bresson's pioneering "decisi...

### frontend/public/demo/threads/d3e5adaf-084c-4dd5-9d29-94f1d6bccd98/user-data/outputs/diana_hu_research.md

- Headings: Diana Hu: Technical Startup Founder Advice - Comprehensive Research; Video Overview; Speaker Background; Education; Career Path; Key Achievements; Escher Reality Acquisition; Video Content Analysis
- Claim: **Title:** Tips For Technical Startup Founders | Startup School **Speaker:** Diana Hu, Y Combinator Group Partner **Date:** April 21, 2023 **Length:** 28 minutes **YouTube URL:*...

### frontend/public/demo/threads/fe3f7974-1bcb-4a01-a950-79673baafefd/user-data/outputs/research_deerflow_20260201.md

- Headings: DeerFlow Deep Research Report; Repository Information; Executive Summary; Complete Chronological Timeline; PHASE 1: Project Inception and Initial Development; PHASE 2: Feature Expansion and Community Growth; PHASE 3: Maturity and DeerFlow 2.0 Transition; Key Analysis
- Claim: - **Research Date:** 2026-02-01 - **Timestamp:** 2026-02-01, Sunday - **Confidence Level:** High (90%+) - **Subject:** ByteDance's Open-Source Multi-Agent Deep Research Framewor...

### frontend/public/images/plugins/README.md

- Headings: Plugin brand assets
- Claim: These bundled assets identify the corresponding services; no third-party image requests are made at runtime. Brand names and logos remain the property of their respective owners...

### frontend/src/AGENTS.md

- Headings: Data Flow; Key Patterns; Interaction Ownership; Knowledge source citations
- Claim: Scope local-turn ordering to `displayThreadId ?? threadId`, the displayed conversation identity. Confirming a new chat changes the SDK `threadId` from undefined to its existing ...

### frontend/src/components/workspace/settings/about.md

- Headings: 🦌 [About DeerFlow 2.0](https://github.com/bytedance/deer-flow); 🌟 GitHub Repository; 🌐 Official Website; 📧 Support; 📜 License; 🙌 Acknowledgments; Core Frameworks; UI Libraries
- Claim: > **From Open Source, Back to Open Source** **DeerFlow** (**D**eep **E**xploration and **E**fficient **R**esearch **Flow**) is a community-driven SuperAgent harness that researc...

### plans/subagent-card-runtime-metadata.md

- Headings: Plan: Subagent Card Runtime Metadata; Architectural decisions; Phase 1: Live Model Identity; What to build; Acceptance criteria; Phase 2: Live Cumulative Token Usage; What to build; Acceptance criteria
- Claim: > Source PRD: Conversation request approved on 2026-07-10 — show live token usage and the effective LLM name on collapsed subagent cards. - **Runtime identity**: Every metadata ...

### scripts/AGENTS.md

- Headings: Support Bundle Runtime Home; Dependency Check Diagnostics; Manual Claude OAuth Export; Service Startup Contracts; Multi-Instance Dev Harness; Shell Script Invocation Contract; Public Skill Review Waivers; Backend Static Analysis Commands
- Claim: Thread manifests use a nonempty `DEER_FLOW_HOME` exclusively, resolving relative values from the checkout like the local launcher. Read root `.env` path settings without exporti...

### skills/custom/ecc/agent-harness-construction/SKILL.md

- Headings: Agent Harness Construction; Core Model; Action Space Design; Granularity Rules; Observation Design; Error Recovery Contract; Context Budgeting; Architecture Pattern Guidance
- Claim: --- name: agent-harness-construction description: Design and optimize AI agent action spaces, tool definitions, and observation formatting for higher completion rates. Use when ...

### skills/custom/ecc/agent-security-hardening/SKILL.md

- Headings: Agent Security Hardening; Important; When to Activate; Hardening Workflow; 1. Map Trust Boundaries; 2. Validate Identifiers, Environment Values, and URLs; 3. Contain Filesystem Access; Treat both platform separators consistently before checking components.
- Claim: --- name: agent-security-hardening description: Security hardening guidance for AI agent frameworks that process untrusted content, invoke tools, write workspace files, manage r...

### skills/custom/ecc/agentic-engineering/SKILL.md

- Headings: Agentic Engineering; Operating Principles; Eval-First Loop; Task Decomposition; Model Routing; Session Strategy; Review Focus for AI-Generated Code; Cost Discipline
- Claim: --- name: agentic-engineering description: Operate as an agentic engineer using eval-first execution, decomposition, and cost-aware model routing. Use when planning or executing...

### skills/custom/ecc/ai-first-engineering/SKILL.md

- Headings: AI-First Engineering; Process Shifts; Architecture Requirements; Code Review in AI-First Teams; Hiring and Evaluation Signals; Testing Standard
- Claim: --- name: ai-first-engineering description: Engineering operating model for teams where AI agents generate a large share of implementation output. Use when setting team process,...

### skills/custom/ecc/api-design/SKILL.md

- Headings: API Design Patterns; When to Activate; Resource Design; URL Structure; Resources are nouns, plural, lowercase, kebab-case; Sub-resources for relationships; Actions that don't map to CRUD (use verbs sparingly); Naming Rules
- Claim: --- name: api-design description: REST API design patterns including resource naming, status codes, pagination, filtering, error responses, versioning, and rate limiting for pro...

### skills/custom/ecc/architecture-decision-records/SKILL.md

- Headings: Architecture Decision Records; When to Activate; ADR Format; ADR-NNNN: [Decision Title]; Context; Decision; Alternatives Considered; Alternative 1: [Name]
- Claim: --- name: architecture-decision-records description: Capture architectural decisions as numbered ADR markdown files in docs/adr/ with context, alternatives considered, consequen...

### skills/custom/ecc/backend-patterns/SKILL.md

- Headings: Backend Development Patterns; When to Activate; API Design Patterns; RESTful API Structure; Repository Pattern; Service Layer Pattern; Middleware Pattern; Database Patterns
- Claim: --- name: backend-patterns description: Backend architecture patterns, API design, database optimization, and server-side best practices for Node.js, Express, and Next.js API ro...

### skills/custom/ecc/codebase-onboarding/SKILL.md

- Headings: Codebase Onboarding; When to Use; How It Works; Phase 1: Reconnaissance; Phase 2: Architecture Mapping; Phase 3: Convention Detection; Phase 4: Generate Onboarding Artifacts; Onboarding Guide: [Project Name]
- Claim: --- name: codebase-onboarding description: Analyze an unfamiliar codebase and generate a structured onboarding guide with architecture map, key entry points, conventions, and a ...

### skills/custom/ecc/coding-standards/SKILL.md

- Headings: Coding Standards & Best Practices; When to Activate; Scope Boundaries; Code Quality Principles; 1. Readability First; 2. KISS (Keep It Simple, Stupid); 3. DRY (Don't Repeat Yourself); 4. YAGNI (You Aren't Gonna Need It)
- Claim: --- name: coding-standards description: Baseline cross-project coding conventions for naming, readability, immutability, and code-quality review. Use detailed frontend or backen...

### skills/custom/ecc/continuous-agent-loop/SKILL.md

- Headings: Continuous Agent Loop; Loop Selection Flow; Combined Pattern; Failure Modes; Recovery
- Claim: --- name: continuous-agent-loop description: Patterns for continuous autonomous agent loops with quality gates, evals, and recovery controls. Use when running an agent loop that...

### skills/custom/ecc/contract-first/SKILL.md

- Headings: Contract-First Collaboration; When to Activate; The Boundary Artifact; Consumer-First Workflow; 1. Identify Consumers and Owners; 2. Describe Consumer Jobs; 3. Define the Smallest Useful Contract; openapi.yaml
- Claim: --- name: contract-first description: Coordinate frontend/backend or service-to-service work through one authoritative machine-checkable contract (OpenAPI, AsyncAPI, Protocol Bu...

### skills/custom/ecc/conventional-commit-review/SKILL.md

- Headings: Conventional Commit Review; When to Activate; Core Concepts; Workflow; Code Examples; Anti-Patterns; Best Practices; Related Skills
- Claim: --- name: conventional-commit-review description: Reviews staged changes and drafts a Conventional Commits-compliant commit message with the right type, scope, and length. Use b...

### skills/custom/ecc/cost-aware-llm-pipeline/SKILL.md

- Headings: Cost-Aware LLM Pipeline; When to Activate; Core Concepts; 1. Model Routing by Task Complexity; 2. Immutable Cost Tracking; 3. Narrow Retry Logic; AuthenticationError, BadRequestError etc. → raise immediately; 4. Prompt Caching
- Claim: --- name: cost-aware-llm-pipeline description: Cost optimization patterns for LLM API usage — model routing by task complexity, budget tracking, retry logic, and prompt caching....

### skills/custom/ecc/deployment-patterns/SKILL.md

- Headings: Deployment Patterns; When to Activate; Deployment Strategies; Rolling Deployment (Default); Blue-Green Deployment; After verification:; Canary Deployment; If metrics look good:
- Claim: --- name: deployment-patterns description: Deployment workflows, CI/CD pipeline patterns, Docker containerization, health checks, rollback strategies, and production readiness c...

### skills/custom/ecc/docker-patterns/SKILL.md

- Headings: Docker Patterns; Docker Compose for Local Development; Standard Web App Stack; docker-compose.yml; Development vs Production Dockerfile; Stage: dependencies; Stage: dev (hot reload, debug tools); Stage: build
- Claim: --- name: docker-patterns description: Docker and Docker Compose patterns for local development, hardened CLI installer harnesses, container security, networking, volumes, and m...

### skills/custom/ecc/e2e-testing/SKILL.md

- Headings: E2E Testing Patterns; Test File Organization; Page Object Model (POM); Test Structure; Playwright Configuration; Flaky Test Patterns; Quarantine; Identify Flakiness
- Claim: --- name: e2e-testing description: Playwright E2E testing patterns, Page Object Model, configuration, CI/CD integration, artifact management, and flaky test strategies. Use when...

### skills/custom/ecc/error-handling/SKILL.md

- Headings: Error Handling Patterns; When to Activate; Core Principles; TypeScript / JavaScript; Typed Error Classes; Result Pattern (no-throw style); API Error Handler (Next.js / Express); React Error Boundary
- Claim: --- name: error-handling description: Patterns for robust error handling across TypeScript, Python, and Go. Covers typed errors, error boundaries, retries, circuit breakers, and...

### skills/custom/ecc/eval-harness/SKILL.md

- Headings: Eval Harness Skill; When to Activate; Philosophy; Eval Types; Capability Evals; Regression Evals; Grader Types; 1. Code-Based Grader
- Claim: --- name: eval-harness description: Eval-driven development (EDD) framework for AI coding sessions — define capability and regression evals before coding, grade with code-based,...

### skills/custom/ecc/git-workflow/SKILL.md

- Headings: Git Workflow Patterns; When to Activate; Branching Strategies; GitHub Flow (Simple, Recommended for Most); Trunk-Based Development (High-Velocity Teams); GitFlow (Complex, Release-Cycle Driven); When to Use Which; Commit Messages
- Claim: --- name: git-workflow description: Git workflow patterns including branching strategies, commit conventions, keeping history clean and readable, tidying local commits before me...

### skills/custom/ecc/golang-patterns/SKILL.md

- Headings: Go Development Patterns; When to Activate; Core Principles; 1. Simplicity and Clarity; 2. Make the Zero Value Useful; 3. Accept Interfaces, Return Structs; Error Handling Patterns; Error Wrapping with Context
- Claim: --- name: golang-patterns description: Idiomatic Go patterns, best practices, and conventions for building robust, efficient, and maintainable Go applications. Use when writing ...

### skills/custom/ecc/golang-testing/SKILL.md

- Headings: Go Testing Patterns; When to Activate; TDD Workflow for Go; The RED-GREEN-REFACTOR Cycle; Step-by-Step TDD in Go; Table-Driven Tests; Table-Driven Tests with Error Cases; Subtests and Sub-benchmarks
- Claim: --- name: golang-testing description: Go testing patterns including table-driven tests, subtests, benchmarks, fuzzing, and test coverage. Follows TDD methodology with idiomatic ...

### skills/custom/ecc/hexagonal-architecture/SKILL.md

- Headings: Hexagonal Architecture; When to Use; Core Concepts; How It Works; Step 1: Model a use case boundary; Step 2: Define outbound ports first; Step 3: Implement the use case with pure orchestration; Step 4: Build adapters at the edge
- Claim: --- name: hexagonal-architecture description: Design, implement, and refactor Ports & Adapters systems with clear domain boundaries, dependency inversion, and testable use-case ...

### skills/custom/ecc/mcp-server-patterns/SKILL.md

- Headings: MCP Server Patterns; When to Use; How It Works; Core concepts; Connecting with stdio; Remote (Streamable HTTP); Examples; Install and server setup
- Claim: --- name: mcp-server-patterns description: Build MCP servers with Node/TypeScript SDK — tools, resources, prompts, Zod validation, stdio vs Streamable HTTP. Use Context7 or offi...

### skills/custom/ecc/prompt-optimizer/SKILL.md

- Headings: Prompt Optimizer; When to Use; Do Not Use When; How It Works; Analysis Pipeline; Phase 0: Project Detection; Phase 1: Intent Detection; Phase 2: Scope Assessment
- Claim: --- name: prompt-optimizer description: Analyze draft prompts, detect intent and missing context, match ECC commands, skills, and agents, and output a ready-to-paste optimized p...

### skills/custom/ecc/security-review/SKILL.md

- Headings: Security Review Skill; When to Activate; Security Checklist; 1. Secrets Management; 2. Input Validation; 3. SQL Injection Prevention; 4. Authentication & Authorization; 5. XSS Prevention
- Claim: --- name: security-review description: Use this skill when adding authentication, handling user input, working with secrets, creating API endpoints, or implementing payment/sens...

### skills/custom/ecc/security-review/cloud-infrastructure-security.md

- Headings: Cloud & Infrastructure Security Skill; When to Activate; Cloud Security Checklist; 1. IAM & Access Control; PASS: CORRECT: Minimal permissions; FAIL: WRONG: Overly broad permissions; ALWAYS enable MFA for root/admin accounts; 2. Secrets Management
- Claim: This skill ensures cloud infrastructure, CI/CD pipelines, and deployment configurations follow security best practices and comply with industry standards. - Deploying applicatio...

### skills/custom/ecc/tdd-workflow/SKILL.md

- Headings: Test-Driven Development Workflow; When to Activate; Plan Handoff; Core Principles; 1. Tests BEFORE Code; 2. Coverage Requirements; 3. Test Types; 4. Git Checkpoints
- Claim: --- name: tdd-workflow description: "Test-driven development workflow: write a failing test first, watch it fail, implement the smallest change to green, then refactor with 80%+...

### skills/custom/ecc/verification-loop/SKILL.md

- Headings: Verification Loop Skill; When to Use; Verification Phases; Phase 1: Build Verification; Check if project builds; OR; Phase 2: Type Check; TypeScript projects
- Claim: --- name: verification-loop description: Run a six-phase verification of a Claude Code session's work — build, type check, lint, tests with coverage, security grep, and diff rev...

### skills/custom/official/cloudflare/SKILL.md

- Headings: Discover and build with Cloudflare; Check for the Cloudflare CLI (`cf`) first; Help the user find the right product; What are you trying to build?; Find guidance for a task not listed here; Caching; Working principles
- Claim: --- name: cloudflare description: Discover and choose Cloudflare products for apps, APIs, AI agents, storage, networking, and security. Use for architecture and product selectio...

### skills/custom/official/cloudflare/references/ai-gateway/README.md

- Headings: Cloudflare AI Gateway; Choose a task; Related references
- Claim: Use AI Gateway to observe and control requests to AI providers through caching, rate limiting, logging, and routing. Fetch the linked documentation before choosing endpoints, au...

### skills/custom/official/cloudflare/references/ai-gateway/configuration.md

- Headings: AI Gateway Configuration
- Claim: Choose the request path before configuring authentication: Cloudflare REST inference, a Workers binding, and provider-native gateway endpoints have different requirements. Gatew...

### skills/custom/official/cloudflare/references/ai-gateway/dynamic-routing.md

- Headings: AI Gateway Dynamic Routing
- Claim: Use a dynamic route when model selection, traffic splitting, quotas, or fallbacks should be controlled in the gateway. For a simple retry or fallback sequence, check the request...

### skills/custom/official/cloudflare/references/ai-gateway/features.md

- Headings: AI Gateway Features
- Claim: Fetch the relevant guide before setting feature flags, headers, limits, or billing behavior. Choose cache keys only for requests whose responses may safely be shared. Decide wha...

### skills/custom/official/cloudflare/references/ai-gateway/sdk-integration.md

- Headings: AI Gateway SDK Integration
- Claim: Choose an integration that matches the application's runtime and required API shape. Fetch its guide before installing packages or writing requests; model identifiers, SDK optio...

### skills/custom/official/cloudflare/references/ai-gateway/troubleshooting.md

- Headings: AI Gateway Troubleshooting
- Claim: Identify the request path and whether the failure comes from gateway access, upstream provider authentication, or request policy before changing credentials or retry behavior. A...

### skills/custom/official/cloudflare/references/ai-search/README.md

- Headings: Cloudflare AI Search; Choose the right product; Find the right documentation
- Claim: Use AI Search for managed content indexing and retrieval, with optional answer generation. Start with [How AI Search works](https://developers.cloudflare.com/ai-search/concepts/...

### skills/custom/official/cloudflare/references/ai-search/api.md

- Headings: AI Search API Routes
- Claim: Choose documentation that matches the integration you are working on. Fetch the reference before writing binding configuration, request types, response parsing, or streaming cod...

### skills/custom/official/cloudflare/references/ai-search/configuration.md

- Headings: AI Search Configuration Routes
- Claim: Indexing credentials and request authentication serve different purposes. For authenticating search requests over HTTP, use the [Search REST API](https://developers.cloudflare.c...

### skills/custom/official/cloudflare/references/ai-search/gotchas.md

- Headings: AI Search Troubleshooting Routes
- Claim: For legacy binding behavior, start with [API routes](api.md). Do not apply current filter syntax or response shapes to legacy calls without following the migration guide.

### skills/custom/official/cloudflare/references/ai-search/patterns.md

- Headings: AI Search Pattern Routes
- Claim: Choose retrieval-only search when your application displays chunks or handles generation itself; choose chat completions when AI Search should also generate the answer. Read [Se...

### skills/custom/official/cloudflare/references/analytics-engine/README.md

- Headings: Cloudflare Workers Analytics Engine Reference; What is Analytics Engine?; Core Concepts; Reading Order; When to Use Analytics Engine; Quick Start; In This Reference; See Also
- Claim: Expert guidance for implementing unlimited-cardinality analytics at scale using Cloudflare Workers Analytics Engine. Time-series analytics database designed for high-cardinality...

### skills/custom/official/cloudflare/references/analytics-engine/api.md

- Headings: Analytics Engine API Reference; Writing Data; `writeDataPoint()`; Full Example; SQL API (External Only); Column References; Query Examples; Response Format
- Claim: Fire-and-forget (returns `void`, not Promise). Writes happen asynchronously. interface AnalyticsEngineDataPoint { blobs?: string[]; // Up to 20 strings (dimensions), 16KB each d...

### skills/custom/official/cloudflare/references/analytics-engine/configuration.md

- Headings: Analytics Engine Configuration; Setup; wrangler.jsonc; TypeScript; Data Point Limits; Write Behavior; Query Limits; Cost
- Claim: 1. Add binding to `wrangler.jsonc` 2. Deploy Worker 3. Dataset created automatically on first write 4. Query via SQL API { "name": "my-worker", "analytics_engine_datasets": [ { ...

### skills/custom/official/cloudflare/references/analytics-engine/gotchas.md

- Headings: Analytics Engine Gotchas; Critical Issues; Sampling at High Volumes; writeDataPoint Returns void; Index vs Blob; Can't Query from Workers; No Custom Timestamps; Common Errors
- Claim: **Problem:** Queries return fewer points than written at >1M writes/min. **Solution:** // Pre-aggregate before writing let buffer = { count: 0, total: 0 }; buffer.count++; buffe...

### skills/custom/official/cloudflare/references/analytics-engine/patterns.md

- Headings: Analytics Engine Patterns; Use Cases; API Metering (Billing); Error Tracking; Performance Monitoring; Anti-Patterns; Best Practices; Schema Template
- Claim: env.ANALYTICS.writeDataPoint({ blobs: [pathname, method, status, tier], doubles: [1, computeUnits, bytes, latencyMs], indexes: [apiKey] }); // Query: Monthly usage by customer /...

### skills/custom/official/cloudflare/references/api/README.md

- Headings: Cloudflare API Integration; Quick Decision Tree; SDK Selection; Authentication Methods; Rate Limits; Reading Order; In This Reference; See Also
- Claim: Guide for working with Cloudflare's REST API - authentication, SDK usage, common patterns, and troubleshooting. How are you calling the Cloudflare API? ├─ From Workers runtime →...

### skills/custom/official/cloudflare/references/api/api.md

- Headings: API Reference; Client Initialization; TypeScript; Python; For async:; Go; Authentication; API Token (Recommended)
- Claim: import Cloudflare from 'cloudflare'; const client = new Cloudflare({ apiToken: process.env.CLOUDFLARE_API_TOKEN, }); from cloudflare import Cloudflare client = Cloudflare(api_to...

### skills/custom/official/cloudflare/references/api/configuration.md

- Headings: Configuration; Environment Variables; Set Variables; .env File Pattern; .env (add to .gitignore); Python; SDK Configuration; TypeScript
- Claim: **Security:** Never commit tokens. Use `.env` files (gitignored) or secret managers. CLOUDFLARE_API_TOKEN=your-token-here CLOUDFLARE_ACCOUNT_ID=your-account-id // TypeScript imp...

### skills/custom/official/cloudflare/references/api/gotchas.md

- Headings: Gotchas & Troubleshooting; Rate Limits & 429 Errors; SDK-Specific Issues; Go: Required Field Wrapper; Python: Async vs Sync Clients; ❌ WRONG - Can't await sync client; ✅ CORRECT - Use AsyncCloudflare; Token Permission Errors (403)
- Claim: **Actual Limits:** - **1200 requests / 5 minutes** per user/token (global) - **200 requests / second** per IP address - **GraphQL: 320 / 5 minutes** (cost-based) **SDK Behavior:...

### skills/custom/official/cloudflare/references/api/patterns.md

- Headings: Common Patterns; List All with Auto-Pagination; Python; Error Handling with Retry; Batch Parallel Operations; Zone CRUD Workflow; DNS Bulk Update; Filter and Collect Results
- Claim: **Problem:** API returns paginated results. Default page size is 20. **Solution:** Use SDK auto-pagination to iterate all results. // TypeScript for await (const zone of client....

### skills/custom/official/cloudflare/references/bindings/README.md

- Headings: Cloudflare Workers Bindings; Start here
- Claim: Bindings grant a Worker access to configured resources through its environment. Prefer a product's binding for supported operations inside Workers; use the REST API when the cal...

### skills/custom/official/cloudflare/references/bindings/api.md

- Headings: Bindings API Reference
- Claim: Fetch the relevant documentation before choosing method signatures, binding types, or type-generation settings. Regenerate types after configuration changes and use the selected...

### skills/custom/official/cloudflare/references/bindings/configuration.md

- Headings: Binding Configuration
- Claim: Confirm the target account, environment, and resource before adding or changing a binding. Keep staging and production resources separate where their data or permissions must be...

### skills/custom/official/cloudflare/references/bindings/gotchas.md

- Headings: Binding Gotchas and Troubleshooting
- Claim: Start by checking the binding name, selected environment, and actual target resource. Regenerating types alone does not fix a missing runtime binding. Never log secret values or...

### skills/custom/official/cloudflare/references/bindings/patterns.md

- Headings: Binding Patterns
- Claim: Choose the interaction and lifecycle first, then retrieve the implementation guide. Use service bindings for internal Worker calls when appropriate, and choose HTTP or RPC based...

### skills/custom/official/cloudflare/references/browser-rendering/README.md

- Headings: Browser Run (formerly Browser Rendering)
- Claim: Use Browser Run for screenshots, PDFs, rendered content extraction, and browser automation. Read the relevant current documentation before implementing; use the [documentation i...

### skills/custom/official/cloudflare/references/browser-rendering/api.md

- Headings: Browser Run APIs
- Claim: Read the guide for the chosen interface for request schemas, return types, authentication, and supported options. Keep Quick Actions and browser session APIs distinct when adapt...

### skills/custom/official/cloudflare/references/browser-rendering/configuration.md

- Headings: Browser Run Configuration
- Claim: Check the project's runtime, installed client and Wrangler versions, and compatibility date before adapting setup instructions. Cloudflare's packages for Workers and standard cl...

### skills/custom/official/cloudflare/references/browser-rendering/gotchas.md

- Headings: Browser Run Troubleshooting
- Claim: Identify the integration and observed failure before changing timeouts or concurrency. A request-rate limit, exhausted browser time, and a closed session require different respo...

### skills/custom/official/cloudflare/references/browser-rendering/patterns.md

- Headings: Browser Run Patterns
- Claim: Use the current examples for the selected integration instead of translating between Puppeteer, Playwright, and Quick Actions by changing method names. Quick Actions manage thei...

### skills/custom/official/cloudflare/references/c3/README.md

- Headings: C3 (create-cloudflare); Quick Start; Interactive (recommended for first-time); Worker (API/WebSocket/Cron); Pages (static/SSG); Platform Decision Tree; Framework Setup; Interactive Flow
- Claim: Official CLI for scaffolding Cloudflare Workers and Pages projects with templates, TypeScript, and instant deployment. npm create cloudflare@latest my-app npm create cloudflare@...

### skills/custom/official/cloudflare/references/c3/api.md

- Headings: C3 CLI Reference; Invocation; Core Flags; Deployment Flags; Advanced Flags; Environment Variables; Exit Codes; Examples
- Claim: npm create cloudflare@latest [name] [-- flags] # NPM requires -- yarn create cloudflare [name] [flags] pnpm create cloudflare@latest [name] [-- flags] CLOUDFLARE_API_TOKEN=xxx #...

### skills/custom/official/cloudflare/references/c3/configuration.md

- Headings: C3 Generated Configuration; Output Structure; wrangler.jsonc; Binding Placeholders; Scripts; Type Generation; Post-Creation Checklist
- Claim: my-app/ ├── src/index.ts # Worker entry point ├── wrangler.jsonc # Cloudflare config ├── package.json # Scripts ├── tsconfig.json └── .gitignore { "$schema": "https://raw.github...

### skills/custom/official/cloudflare/references/c3/gotchas.md

- Headings: C3 Troubleshooting; Deployment Issues; Placeholder IDs; Authentication; Name Conflict; Platform Selection; TypeScript Issues; Restart TS server in editor
- Claim: **Error:** "Invalid namespace ID" **Fix:** Replace placeholders in wrangler.jsonc with real IDs: npx wrangler kv namespace create MY_KV # Get real ID **Error:** "Not authenticat...

### skills/custom/official/cloudflare/references/c3/patterns.md

- Headings: C3 Usage Patterns; Quick Workflows; TypeScript API Worker; Astro static site; CI/CD (GitHub Actions); Monorepo; Custom Templates; GitHub repo
- Claim: For framework apps, follow [Framework Setup](README.md#framework-setup). npm create cloudflare@latest my-api -- --type=hello-world --lang=ts --deploy npm create cloudflare@lates...

### skills/custom/official/cloudflare/references/containers/README.md

- Headings: Cloudflare Containers; Choose the runtime; Find the documentation for the task
- Claim: Use this reference for containerized applications on the Workers platform, including container-enabled Durable Objects, lifecycle management, and request routing. Use [Container...

### skills/custom/official/cloudflare/references/containers/api.md

- Headings: Containers API
- Claim: Use the [Container interface](https://developers.cloudflare.com/containers/api/container-class/index.md) for the current SDK methods, signatures, properties, and examples. Use t...

### skills/custom/official/cloudflare/references/containers/configuration.md

- Headings: Containers configuration
- Claim: Read the relevant documentation before choosing configuration fields or resource sizes.

### skills/custom/official/cloudflare/references/containers/gotchas.md

- Headings: Containers troubleshooting
- Claim: Use the current documentation to diagnose behavior instead of relying on copied timeout values, resource limits, or lifecycle recipes.

### skills/custom/official/cloudflare/references/containers/patterns.md

- Headings: Containers patterns; Workflows and Queues
- Claim: Choose instance identity based on the workload: per-user/session or per-job identities for affinity, one shared identity for a singleton, and interchangeable instances for state...

### skills/custom/official/cloudflare/references/cron-triggers/README.md

- Headings: Cloudflare Cron Triggers; In This Reference; See Also
- Claim: Use Cron Triggers to start periodic Worker jobs. Fetch the relevant current documentation before implementing; configuration, API signatures, examples, and limits belong in the ...

### skills/custom/official/cloudflare/references/cron-triggers/api.md

- Headings: Cron Triggers API
- Claim: Fetch the handler documentation before writing code; use its current language examples and completion semantics. Decide which operation establishes successful completion and mak...

### skills/custom/official/cloudflare/references/cron-triggers/configuration.md

- Headings: Cron Triggers Configuration
- Claim: Fetch the documentation for the configuration operation before changing schedules. Identify the target environment and the intended business timezone before choosing an expressi...

### skills/custom/official/cloudflare/references/cron-triggers/gotchas.md

- Headings: Cron Triggers Troubleshooting
- Claim: Investigate using the current documentation rather than copied limits or assumed delivery guarantees. For repeated or partially completed business operations, decide how to iden...

### skills/custom/official/cloudflare/references/cron-triggers/patterns.md

- Headings: Cron Triggers Patterns
- Claim: Choose the execution boundary before writing a scheduled job; fetch the relevant integration docs for implementation. Keep the trigger separate from the business operation so ma...

### skills/custom/official/cloudflare/references/d1/README.md

- Headings: Cloudflare D1 Database; Start here
- Claim: Use D1 for managed relational application data with SQLite semantics. For an existing external database, consider [Hyperdrive](../hyperdrive/); for per-entity coordination, cons...

### skills/custom/official/cloudflare/references/d1/api.md

- Headings: D1 API Reference
- Claim: Fetch the relevant API page before writing queries or assuming method signatures and return types. Bind untrusted values with prepared statements; do not interpolate them into S...

### skills/custom/official/cloudflare/references/d1/configuration.md

- Headings: D1 Configuration
- Claim: Read the task's documentation before adding bindings or running database commands. Confirm the database and environment being targeted, particularly when applying migrations or ...

### skills/custom/official/cloudflare/references/d1/gotchas.md

- Headings: D1 Gotchas & Troubleshooting
- Claim: Use current documentation to diagnose the failure before changing query or database configuration. Continue to bind untrusted SQL values as described in [api.md](./api.md). Do n...

### skills/custom/official/cloudflare/references/d1/patterns.md

- Headings: D1 Patterns & Best Practices
- Claim: Use these guides to design the operation, then fetch [api.md](./api.md) for implementation references. Keep result sets bounded and pagination ordering deterministic. Choose ind...

### skills/custom/official/cloudflare/references/flagship/README.md

- Headings: Cloudflare Flagship; When to Use; Key Concepts; Two Evaluation Paths; Reading Order; In This Reference; See Also
- Claim: Feature flag service for controlling feature visibility without redeploying code. Define flags with targeting rules and percentage-based rollouts, then evaluate them in Workers ...

### skills/custom/official/cloudflare/references/flagship/api.md

- Headings: Flagship API Reference; Binding API (Workers); Evaluation Methods; Parameters (shared across all methods); Types; Example; OpenFeature SDK; Server Provider (`FlagshipServerProvider`)
- Claim: The binding is available as `env.FLAGS` (type `Flagship` from `@cloudflare/workers-types`). All methods are async, never throw, and return the `defaultValue` on errors. type Fla...

### skills/custom/official/cloudflare/references/flagship/configuration.md

- Headings: Flagship Configuration; Wrangler Binding Setup; Single App; wrangler.toml; Multiple Apps; wrangler.toml; Generate Types; OpenFeature SDK Installation
- Claim: Add a Flagship binding to your Wrangler config to access flags via `env.FLAGS`. // wrangler.jsonc { "flagship": { "binding": "FLAGS", "app_id": "<APP_ID>" } } [flagship] binding...

### skills/custom/official/cloudflare/references/flagship/gotchas.md

- Headings: Flagship Gotchas & Troubleshooting; Common Errors; Flag Always Returns Default Value; TYPE_MISMATCH Error in Details; 409 Conflict on Flag Creation; Inconsistent Rollout Results; Update Overwrites Entire Flag; ❌ BAD — overwrites the entire flag, losing rules/variations
- Claim: **Cause:** Flag is disabled (`enabled: false`), or no targeting rules match, or evaluation context is missing expected attributes. **Solution:** Check these in order: 1. Is the ...

### skills/custom/official/cloudflare/references/flagship/patterns.md

- Headings: Flagship Patterns & Best Practices; Evaluating Flags in Workers (Binding); Simple Boolean Toggle; Multi-Variant String Flag; JSON Config Flag; Using Details for Observability; Evaluating Flags with OpenFeature (Workers); Binding Passthrough (Recommended)
- Claim: export default { async fetch(request: Request, env: Env): Promise<Response> { const showNewUI = await env.FLAGS.getBooleanValue("new-ui", false, { userId: "user-42", }); if (sho...

### skills/custom/official/cloudflare/references/graphql-api/README.md

- Headings: Cloudflare GraphQL Analytics API; Overview; Quick Decision Tree; Core Concepts; Query Structure; Zone-scoped; Account-scoped; Dataset Naming Convention
- Claim: Query analytics data across all Cloudflare products via a single GraphQL endpoint. Covers HTTP requests, Workers metrics, DNS, Firewall events, Network Analytics, and 70+ other ...

### skills/custom/official/cloudflare/references/graphql-api/api.md

- Headings: GraphQL Analytics API Reference; Query Root; Aggregation Fields; count; sum; HTTP requests; Workers invocations; quantiles
- Claim: The schema has a single entry point: `Query.viewer`. Mutations are not supported. { cost # uint64 -- query cost (returned in response) viewer { budget # uint64 -- remaining budg...

### skills/custom/official/cloudflare/references/graphql-api/configuration.md

- Headings: GraphQL Analytics API Configuration; Authentication; API Token (Recommended); Verify token; API Key + Email (Legacy); Client Setup; curl; TypeScript / JavaScript
- Claim: Create tokens at: [dash.cloudflare.com > Account API Tokens](https://dash.cloudflare.com/?to=/:account/api-tokens) curl -s https://api.cloudflare.com/client/v4/graphql \ -H "Aut...

### skills/custom/official/cloudflare/references/graphql-api/gotchas.md

- Headings: GraphQL Analytics API Gotchas & Troubleshooting; Rate Limits; "429 Too Many Requests"; Sampling & Data Accuracy; Adaptive Bit Rate (ABR) Sampling; Rollup vs. Adaptive; Common Errors; "Access denied" / "authentication error"
- Claim: The GraphQL rate limit is separate from the general API limit. Exceeding either results in `HTTP 429` and blocks all API calls for 5 minutes. Enterprise customers can contact su...

### skills/custom/official/cloudflare/references/graphql-api/patterns.md

- Headings: GraphQL Analytics API Patterns & Best Practices; Time-Series Queries; Top-N Queries; Top Countries by Request Count; Workers Analytics; Firewall / Security; DNS Analytics; Storage Analytics (Account-Scoped)
- Claim: Use time dimension granularity matching your range (see Best Practices below). query TrafficTimeSeries($zoneTag: string!, $start: Time!, $end: Time!) { viewer { zones(filter: { ...

### skills/custom/official/cloudflare/references/hyperdrive/README.md

- Headings: Hyperdrive; Retrieve current documentation; Choose the next reference; Decisions to preserve; See also
- Claim: Use Hyperdrive to connect Workers to an existing PostgreSQL or MySQL database with connection pooling and optional query caching. It does not replace the origin database or repl...

### skills/custom/official/cloudflare/references/hyperdrive/api.md

- Headings: Hyperdrive API and drivers; Driver and binding routes; ORMs and query builders
- Claim: Start with [README.md](./README.md) and [configuration.md](./configuration.md). Fetch the selected guide before writing connection or query code; use its current supported packa...

### skills/custom/official/cloudflare/references/hyperdrive/configuration.md

- Headings: Hyperdrive configuration; Setup decisions
- Claim: See [README.md](./README.md) for the retrieval workflow. Fetch the relevant guide before creating or changing resources; use current configuration fields and CLI syntax from the...

### skills/custom/official/cloudflare/references/hyperdrive/gotchas.md

- Headings: Hyperdrive troubleshooting; Capacity and changes
- Claim: Start with the actual error and the affected configuration. Fetch [Troubleshoot and debug](https://developers.cloudflare.com/hyperdrive/observability/troubleshooting/index.md) f...

### skills/custom/official/cloudflare/references/hyperdrive/patterns.md

- Headings: Hyperdrive design patterns
- Claim: See [api.md](./api.md) for maintained driver and ORM examples. Use the following decisions to select a pattern, then fetch its linked documentation for implementation. Separate ...

### skills/custom/official/cloudflare/references/images/README.md

- Headings: Cloudflare Images
- Claim: Choose the image source and operation before selecting an API. Hosted-image management, remote URL transformations, and the Workers optimization binding have different contracts...

### skills/custom/official/cloudflare/references/images/api.md

- Headings: Images API Selection
- Claim: Do not transfer URL parameters or HTTP request shapes directly into binding calls. Read the contract for the selected interface, including output format handling. Use the projec...

### skills/custom/official/cloudflare/references/images/configuration.md

- Headings: Images Configuration
- Claim: Inspect the project's Wrangler configuration, dependency versions, existing bindings, and credential storage before changing setup. Preserve its configuration format and generat...

### skills/custom/official/cloudflare/references/images/gotchas.md

- Headings: Images Troubleshooting
- Claim: First identify whether the failure involves hosted-image storage, remote URL transformations, or a Worker binding. Capture the failing operation, response status, relevant heade...

### skills/custom/official/cloudflare/references/images/patterns.md

- Headings: Images Patterns
- Claim: Choose the workflow that matches the existing storage and delivery architecture, then read its implementation guide. Adapt dimensions and quality to the actual layout and repres...

### skills/custom/official/cloudflare/references/kv/README.md

- Headings: Cloudflare Workers KV; Start here
- Claim: Use KV for read-heavy configuration, preferences, and application caches that tolerate stale data. Read [how KV works](https://developers.cloudflare.com/kv/concepts/how-kv-works...

### skills/custom/official/cloudflare/references/kv/api.md

- Headings: KV API Reference
- Claim: Read the relevant API page before implementing; it defines current options, result shapes, supported bulk operations, and constraints. Handle missing values explicitly: JavaScri...

### skills/custom/official/cloudflare/references/kv/configuration.md

- Headings: KV Configuration
- Claim: Read the setup guide for the target environment before creating resources or editing bindings. Choose the namespace, account, and environment deliberately. Local KV data is sepa...

### skills/custom/official/cloudflare/references/kv/gotchas.md

- Headings: KV Gotchas & Troubleshooting
- Claim: Read the linked explanation before applying a workaround. Confirm freshness and failure requirements before adding a cache or a permissive fallback; see [patterns.md](./patterns...

### skills/custom/official/cloudflare/references/kv/patterns.md

- Headings: KV Patterns & Best Practices; Application-specific decisions
- Claim: Read the guide for the pattern before implementing it, and confirm that [KV's consistency model](https://developers.cloudflare.com/kv/concepts/how-kv-works/index.md) fits the ap...

### skills/custom/official/cloudflare/references/miniflare/README.md

- Headings: Miniflare; Choose the testing tool; Read for the task
- Claim: Miniflare provides programmatic control of local Workers simulation. Read the linked documentation before choosing APIs, configuration, or a migration path. For interactive loca...

### skills/custom/official/cloudflare/references/miniflare/api.md

- Headings: Miniflare API
- Claim: Use the current documentation for method signatures and examples: For constructor options, read [configuration.md](./configuration.md). For runtime-specific test helpers, read [...

### skills/custom/official/cloudflare/references/miniflare/configuration.md

- Headings: Miniflare Configuration
- Claim: Direct Miniflare does not read Wrangler configuration. Configure its bindings explicitly and build TypeScript or bundled Workers before starting tests; see [writing tests](https...

### skills/custom/official/cloudflare/references/miniflare/gotchas.md

- Headings: Miniflare Troubleshooting and Migrations
- Claim: For a migration, choose the target using [the testing-tool decision](./README.md#choose-the-testing-tool) before translating old options. A historical migration page describes t...

### skills/custom/official/cloudflare/references/miniflare/patterns.md

- Headings: Miniflare Testing Patterns
- Claim: Choose the test runtime before adapting an example. With direct Miniflare, the Worker runs in workerd while the test runner runs in Node.js; importing Worker functions into Node...

### skills/custom/official/cloudflare/references/observability/README.md

- Headings: Cloudflare Observability; Choose a signal; Load only what the task needs
- Claim: Use this reference to choose a telemetry signal and find the maintained implementation guide. Fetch the linked documentation before writing configuration, queries, or export cod...

### skills/custom/official/cloudflare/references/observability/api.md

- Headings: Observability APIs
- Claim: Fetch the applicable reference for current signatures, field locations, units, authentication, and query syntax. Do not infer the Tail event schema from an OpenTelemetry span or...

### skills/custom/official/cloudflare/references/observability/configuration.md

- Headings: Observability Configuration; Setup decisions
- Claim: Fetch the relevant guide before configuring the selected Worker and deployment environment. - Confirm which account, Worker, and environment will emit telemetry, then deploy tha...

### skills/custom/official/cloudflare/references/observability/gotchas.md

- Headings: Observability Troubleshooting and Constraints; Missing or incomplete data; Limits, retention, and cost
- Claim: Fetch these pages when estimating cost or diagnosing truncation and missing data: - [Workers Logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/inde...

### skills/custom/official/cloudflare/references/observability/patterns.md

- Headings: Observability Patterns
- Claim: Use these decisions alongside the linked implementation guides. Application event schemas, billing policies, alert thresholds, and delivery behavior still need to be designed fo...

### skills/custom/official/cloudflare/references/pages/README.md

- Headings: Cloudflare Pages; In This Reference
- Claim: Use this reference when maintaining an existing Pages project. For new applications, start with Workers as recommended in the [Pages framework guidance](https://developers.cloud...

### skills/custom/official/cloudflare/references/pages/api.md

- Headings: Pages Request Handling
- Claim: Use the current Pages API and the project framework guide rather than translating a generic Worker example into Pages. First identify whether routing comes from the Functions di...

### skills/custom/official/cloudflare/references/pages/configuration.md

- Headings: Pages Configuration
- Claim: Inspect the existing project configuration and build output before changing deployment settings. Distinguish build-time variables from runtime bindings, and check preview and pr...

### skills/custom/official/cloudflare/references/pages/gotchas.md

- Headings: Pages Troubleshooting; Framework-Specific
- Claim: Identify whether the failure occurs during the build, asset serving, or Function execution before changing configuration. Compare the same route in local, preview, and productio...

### skills/custom/official/cloudflare/references/pages/patterns.md

- Headings: Pages Project Decisions
- Claim: Keep project-specific choices here; fetch framework adapters, configuration, and code examples from the docs. Check the existing build command, adapter, and output ownership tog...

### skills/custom/official/cloudflare/references/pages-functions/README.md

- Headings: Cloudflare Pages Functions; In This Reference
- Claim: Use this reference for server-side behavior in an existing Pages project. For new applications, follow the Workers recommendation in the [Pages framework guidance](https://devel...

### skills/custom/official/cloudflare/references/pages-functions/api.md

- Headings: Pages Functions APIs
- Claim: Fetch the API reference before writing a handler; keep runtime types and examples in their authoritative documentation. Decide which handler owns the response, where shared stat...

### skills/custom/official/cloudflare/references/pages-functions/configuration.md

- Headings: Pages Functions Configuration
- Claim: Read Pages-specific configuration before reusing settings from a Worker project. Identify the target deployment environment and which configuration source controls it. Verify Pa...

### skills/custom/official/cloudflare/references/pages-functions/gotchas.md

- Headings: Pages Functions Troubleshooting
- Claim: Start with the request path, deployment environment, and generated output that actually handled the request. Reproduce a failing path through the actual application rather than ...

### skills/custom/official/cloudflare/references/pages-functions/patterns.md

- Headings: Pages Functions Request Design
- Claim: Choose where behavior belongs before adapting an example. Define the response owner and middleware scope before adding authentication or response transformations. Test protected...

### skills/custom/official/cloudflare/references/pulumi/README.md

- Headings: Cloudflare Pulumi Provider; Overview; Core Principles; Authentication; Setup; Common Resource Types; Key Properties; Reading Order
- Claim: Expert guidance for Cloudflare Pulumi Provider (@pulumi/cloudflare). Programmatic management of Cloudflare resources: Workers, Pages, D1, KV, R2, DNS, Queues, etc. **Packages:**...

### skills/custom/official/cloudflare/references/pulumi/api.md

- Headings: API & Data Sources; Outputs and Exports; Resource Dependencies; Using Outputs with API Calls; Custom Dynamic Providers; Data Sources; Import Existing Resources; Import worker
- Claim: Export resource identifiers: export const kvId = kv.id; export const bucketName = bucket.name; export const workerUrl = worker.subdomain; export const dbId = db.id; Implicit dep...

### skills/custom/official/cloudflare/references/pulumi/configuration.md

- Headings: Resource Configuration; Workers (cloudflare.WorkerScript); Workers KV (cloudflare.WorkersKvNamespace); R2 Buckets (cloudflare.R2Bucket); D1 Databases (cloudflare.D1Database); Queues (cloudflare.Queue); Pages Projects (cloudflare.PagesProject); DNS Records (cloudflare.DnsRecord)
- Claim: import * as cloudflare from "@pulumi/cloudflare"; import * as fs from "fs"; const worker = new cloudflare.WorkerScript("my-worker", { accountId: accountId, name: "my-worker", co...

### skills/custom/official/cloudflare/references/pulumi/gotchas.md

- Headings: Troubleshooting & Best Practices; Common Errors; "No bundler/build step" - Pulumi uploads raw code; "wrangler.toml not consumed" - Config drift; "False no-changes detection" - Content SHA unchanged; "D1 migrations don't run on pulumi up"; "Missing required property 'accountId'"; Pulumi.<stack>.yaml
- Claim: **Problem:** Worker fails with "Cannot use import statement outside a module" **Cause:** Pulumi doesn't bundle Worker code - uploads exactly what you provide **Solution:** Build...

### skills/custom/official/cloudflare/references/pulumi/patterns.md

- Headings: Architecture Patterns; Component Resources; Full-Stack Worker App; Multi-Environment Setup; Queue-Based Processing; Microservices with Service Bindings; Event-Driven Architecture; v6.x Versioned Deployments (Blue-Green/Canary)
- Claim: class WorkerApp extends pulumi.ComponentResource { constructor(name: string, args: WorkerAppArgs, opts?) { super("custom:cloudflare:WorkerApp", name, {}, opts); const defaultOpt...

### skills/custom/official/cloudflare/references/queues/README.md

- Headings: Cloudflare Queues; Choose a consumer; Read by task
- Claim: Use Queues to decouple producers from asynchronous consumers and buffer bursts of work. Design consumers for duplicate delivery; use Workflows when the task needs durable multi-...

### skills/custom/official/cloudflare/references/queues/api.md

- Headings: Queues API Reference
- Claim: Fetch the current API documentation for the operation being implemented; do not infer signatures, payloads, or acknowledgement rules from old examples. Acknowledge only after th...

### skills/custom/official/cloudflare/references/queues/configuration.md

- Headings: Queues Configuration
- Claim: Fetch the relevant guide before writing configuration or running CLI commands. Check the project's Wrangler version and compatibility date when adapting examples. Choose push or...

### skills/custom/official/cloudflare/references/queues/gotchas.md

- Headings: Queues Gotchas & Troubleshooting
- Claim: Fetch the linked documentation before changing retry policy or interpreting delivery behavior. Distinguish transient dependency failures from invalid payloads before choosing re...

### skills/custom/official/cloudflare/references/queues/patterns.md

- Headings: Queues Patterns & Best Practices
- Claim: Fetch the guide matching the task and adapt its example to the application's delivery and failure requirements. Design side effects for [at-least-once delivery](https://develope...

### skills/custom/official/cloudflare/references/r2/README.md

- Headings: Cloudflare R2 Object Storage; Choose an access path; Find the task; See also
- Claim: Use R2 for objects such as uploads, media, backups, and static assets. Fetch the linked documentation before implementing; API signatures, configuration, limits, and pricing bel...

### skills/custom/official/cloudflare/references/r2/api.md

- Headings: R2 API Reference
- Claim: Fetch the relevant page before writing code. Use the Workers API for bucket bindings and the S3 API for S3 clients; their types and semantics differ. Use generated project types...

### skills/custom/official/cloudflare/references/r2/configuration.md

- Headings: R2 Configuration
- Claim: Fetch the task's documentation before editing Wrangler configuration or bucket settings. Choose the development bucket deliberately: a remote binding accesses real data. Scope S...

### skills/custom/official/cloudflare/references/r2/gotchas.md

- Headings: R2 Gotchas & Troubleshooting
- Claim: Use the current references to diagnose the actual response or error instead of copying a workaround. For other failures, start with [R2 troubleshooting](https://developers.cloud...

### skills/custom/official/cloudflare/references/r2/patterns.md

- Headings: R2 Patterns & Best Practices
- Claim: Choose the access and delivery model, then fetch the implementation guide. Authorize the caller for the selected object key and operation before exposing a Worker endpoint or is...

### skills/custom/official/cloudflare/references/secrets-store/README.md

- Headings: Cloudflare Secrets Store; Choose the scope; Read by task
- Claim: Use Secrets Store for account-level credentials shared across Workers or supported integrations. Use [Worker secrets](https://developers.cloudflare.com/workers/configuration/sec...

### skills/custom/official/cloudflare/references/secrets-store/api.md

- Headings: Secrets Store APIs; Runtime decisions
- Claim: Fetch the current API documentation before implementing calls or copying types. Account-secret bindings require asynchronous retrieval; they are not the direct string values exp...

### skills/custom/official/cloudflare/references/secrets-store/configuration.md

- Headings: Secrets Store configuration; Setup and management; Local development and deployment
- Claim: Fetch the guide for the operation you are performing before writing configuration or running management commands. Treat the store ID, secret ID, secret name, and Worker binding ...

### skills/custom/official/cloudflare/references/secrets-store/gotchas.md

- Headings: Secrets Store troubleshooting
- Claim: Start with the failing operation and fetch its documentation before changing credentials or bindings. Do not diagnose failures by printing secret values or returning raw errors ...

### skills/custom/official/cloudflare/references/secrets-store/patterns.md

- Headings: Secrets Store patterns; Rotation and sharing; Migrate from Worker secrets; Encryption and signing; Audit and integrations
- Claim: Use these decisions to choose the relevant guides; fetch the linked documentation before implementing. Fetch [edit, duplicate, and delete operations](https://developers.cloudfla...

### skills/custom/official/cloudflare/references/smart-placement/README.md

- Headings: Cloudflare Workers Smart Placement; Core Concept; When to Use; Decision Tree; Key Architecture Pattern; Quick Start; Requirements; Placement Status Values
- Claim: Automatic workload placement optimization to minimize latency by running Workers closer to backend infrastructure rather than end users. Smart Placement automatically analyzes W...

### skills/custom/official/cloudflare/references/smart-placement/api.md

- Headings: Smart Placement API; Placement Status API; Status Meanings; cf-placement Header (Beta); Detecting Smart Placement in Code; Request Duration Metrics; Interpreting Metrics; Monitoring Commands
- Claim: Query Worker placement status via Cloudflare API: curl -X GET "https://api.cloudflare.com/client/v4/accounts/{ACCOUNT_ID}/workers/services/{WORKER_NAME}" \ -H "Authorization: Be...

### skills/custom/official/cloudflare/references/smart-placement/configuration.md

- Headings: Smart Placement Configuration; wrangler.jsonc Setup; Placement Mode Values; Frontend + Backend Split Configuration; Frontend Worker (No Smart Placement); Backend Worker (Smart Placement Enabled); Requirements & Limitations; Requirements
- Claim: { "$schema": "./node_modules/wrangler/config-schema.json", "placement": { "mode": "smart" } } **Note:** Smart Placement vs Explicit Placement are separate features. Smart Placem...

### skills/custom/official/cloudflare/references/smart-placement/gotchas.md

- Headings: Smart Placement Gotchas; Common Errors; "INSUFFICIENT_INVOCATIONS"; "UNSUPPORTED_APPLICATION"; "No request duration metrics"; "cf-placement header missing"; Pages/Assets + Smart Placement Performance Degradation; Monolithic Full-Stack Worker
- Claim: **Cause:** Not enough traffic for Smart Placement to analyze **Solution:** - Ensure Worker receives consistent global traffic - Wait longer (analysis takes up to 15 minutes) - S...

### skills/custom/official/cloudflare/references/smart-placement/patterns.md

- Headings: Smart Placement Patterns; Backend Worker with Database Access; Frontend + Backend Split (Service Bindings); External API Integration; SSR / API Gateway Pattern; Durable Objects with Smart Placement; Best Practices
- Claim: export default { async fetch(request: Request, env: Env): Promise<Response> { const user = await env.DATABASE.prepare('SELECT * FROM users WHERE id = ?').bind(userId).first(); c...

### skills/custom/official/cloudflare/references/static-assets/README.md

- Headings: Workers Static Assets; Reading Order
- Claim: Use Workers Static Assets for new static sites, SPAs, generated sites, and applications combining assets with server logic. Inspect the framework, build output, and existing dep...

### skills/custom/official/cloudflare/references/static-assets/api.md

- Headings: Static Assets Binding API
- Claim: Read the binding reference before implementing calls. Check the configured binding name and use the project's existing environment types. When selecting a different asset, const...

### skills/custom/official/cloudflare/references/static-assets/configuration.md

- Headings: Static Assets Configuration
- Claim: Inspect the build script, output directory, Wrangler configuration, and any framework-generated deployment configuration first. Configure the files actually produced by the buil...

### skills/custom/official/cloudflare/references/static-assets/gotchas.md

- Headings: Static Assets Gotchas
- Claim: Compare the deployed build output and effective configuration with the exact request that failed. Test browser navigation separately from client-side fetches when investigating ...

### skills/custom/official/cloudflare/references/static-assets/patterns.md

- Headings: Static Assets Routing Patterns
- Claim: Decide which requests need application logic before choosing configuration. Record expected behavior for public files, API paths, protected content, browser navigation, and miss...

### skills/custom/official/cloudflare/references/stream/README.md

- Headings: Cloudflare Stream; Choose a Workflow; In This Reference
- Claim: Use Stream for on-demand video upload and playback or live broadcasting. Start with the application's upload source, player, and access requirements, then read the corresponding...

### skills/custom/official/cloudflare/references/stream/api-live.md

- Headings: Stream Live Streaming API
- Claim: Choose the ingest protocol and viewer experience before creating live inputs. Read the current examples for returned endpoints and credentials instead of constructing publish UR...

### skills/custom/official/cloudflare/references/stream/api.md

- Headings: Stream API Reference
- Claim: Read the task-specific documentation for request schemas, SDK examples, and response fields. For live inputs and outputs, use [api-live.md](./api-live.md). Keep the returned vid...

### skills/custom/official/cloudflare/references/stream/configuration.md

- Headings: Stream Configuration
- Claim: Inspect the existing runtime, API client or binding, authentication layer, player, and secret storage before adding Stream. Choose upload and viewing permissions from the applic...

### skills/custom/official/cloudflare/references/stream/gotchas.md

- Headings: Stream Gotchas
- Claim: Identify whether the failure occurs during upload, encoding, authorization, playback, or live ingestion before changing configuration. Read the linked requirements and error gui...

### skills/custom/official/cloudflare/references/stream/patterns.md

- Headings: Stream Patterns; Application Decisions
- Claim: Use the official workflow examples after identifying where the existing application handles authorization, video ownership, and processing state. - Check the user's permission b...

### skills/custom/official/cloudflare/references/tail-workers/README.md

- Headings: Cloudflare Tail Workers; In This Reference
- Claim: Use Tail Workers when execution events need custom processing. Fetch the current documentation before implementing handlers, configuration, or integrations. Before adding a Tail...

### skills/custom/official/cloudflare/references/tail-workers/api.md

- Headings: Tail Workers APIs
- Claim: Fetch the handler reference for current event shapes and language examples instead of maintaining local interface definitions. Choose the event fields required by the destinatio...

### skills/custom/official/cloudflare/references/tail-workers/configuration.md

- Headings: Tail Workers Configuration
- Claim: Identify the producer Worker and consumer Worker as separate deployment resources before changing their configuration. Confirm which project owns each deployment and which envir...

### skills/custom/official/cloudflare/references/tail-workers/gotchas.md

- Headings: Tail Workers Troubleshooting
- Claim: Trace one known producer invocation through consumer execution and destination receipt to locate the failure. Verify producer configuration, consumer identity, and destination c...

### skills/custom/official/cloudflare/references/tail-workers/patterns.md

- Headings: Tail Workers Processing Decisions
- Claim: Define the destination contract and required event coverage before writing transformations. Decide which producers, routes, failures, and successful requests each destination ne...

### skills/custom/official/cloudflare/references/terraform/README.md

- Headings: Cloudflare Terraform Provider; Core Principles; Provider Version; Provider Setup; Basic Configuration; Authentication Methods (priority order); Quick Reference: Common Commands; Import Existing Resources
- Claim: **Expert guidance for Cloudflare Terraform Provider - infrastructure as code for Cloudflare resources.** - **Provider-first**: Use Terraform provider for ALL infrastructure - ne...

### skills/custom/official/cloudflare/references/terraform/api.md

- Headings: Terraform Data Sources Reference; v5 Data Source Names; Zone Data Sources; Get zone by name; Use in resources; ...; Account Data Sources; List all accounts
- Claim: Query existing Cloudflare resources to reference in your configurations. data "cloudflare_zone" "example" { name = "example.com" } resource "cloudflare_dns_record" "www" { zone_...

### skills/custom/official/cloudflare/references/terraform/configuration.md

- Headings: Terraform Configuration Reference; Zone & DNS; Zone + settings; DNS records (A, CNAME, MX, TXT); Workers; Simple Pattern (Legacy - Still Works); Gradual Rollouts (Recommended for Production); Worker Binding Types (v5)
- Claim: Complete resource configurations for Cloudflare infrastructure. resource "cloudflare_zone" "example" { account = { id = var.account_id }; name = "example.com"; type = "full" } r...

### skills/custom/official/cloudflare/references/terraform/gotchas.md

- Headings: Terraform Troubleshooting & Best Practices; State Drift Issues; Example: Ignore secret drift; v5 Breaking Changes; Rename resources in state after v5 upgrade; Resource-Specific Gotchas; R2 Location Case Sensitivity; KV Special Characters (< 5.16.0)
- Claim: Common issues, security considerations, and best practices. Some resources have known state drift. Add lifecycle blocks to prevent perpetual diffs: resource "cloudflare_workers_...

### skills/custom/official/cloudflare/references/terraform/patterns.md

- Headings: Terraform Patterns & Use Cases; Recommended Directory Structure; Multi-Environment Setup; Directory: environments/{production,staging}/main.tf + modules/{zone,worker,pages}; R2 State Backend; Worker with All Bindings; Wrangler Integration; CI/CD Pattern
- Claim: Architecture patterns, multi-environment setups, and real-world use cases. terraform/ ├── environments/ │ ├── production/ │ │ ├── main.tf │ │ └── terraform.tfvars │ └── staging/...

### skills/custom/official/cloudflare/references/tunnel/README.md

- Headings: Cloudflare Tunnel; In This Reference
- Claim: Use Tunnel to connect origin services to Cloudflare. Inspect the existing tunnel, management mode, and intended audience before choosing a setup. Fetch current docs for commands...

### skills/custom/official/cloudflare/references/tunnel/api.md

- Headings: Tunnel APIs and Commands
- Claim: Fetch current operation schemas, permissions, and examples before automating tunnel changes. Identify the account, tunnel ID, and management mode before a write. Review existing...

### skills/custom/official/cloudflare/references/tunnel/configuration.md

- Headings: Tunnel Configuration
- Claim: Read the documentation for the existing management mode before changing routes or credentials. Confirm which configuration source the running process uses, then review the route...

### skills/custom/official/cloudflare/references/tunnel/gotchas.md

- Headings: Tunnel Troubleshooting
- Claim: Capture the tunnel status, failing route, management mode, and cloudflared logs before changing settings. Separate tunnel health from origin availability. Check the service addr...

### skills/custom/official/cloudflare/references/tunnel/networking.md

- Headings: Tunnel Networking
- Claim: Investigate the connection from cloudflared to Cloudflare separately from the connection to the origin and the client access path. Test from the machine or container running clo...

### skills/custom/official/cloudflare/references/tunnel/patterns.md

- Headings: Tunnel Deployment Decisions
- Claim: Choose deployment ownership and availability requirements before adapting an example. Establish which service each replica can reach and how configuration and credentials reach ...

### skills/custom/official/cloudflare/references/turn/README.md

- Headings: Cloudflare TURN Service; Overview; Key Characteristics; In This Reference; Reading Order; Service Addresses and Ports; STUN over UDP; TURN over UDP
- Claim: Expert guidance for implementing Cloudflare TURN Service in WebRTC applications. Cloudflare TURN (Traversal Using Relays around NAT) Service is a managed relay service for WebRT...

### skills/custom/official/cloudflare/references/turn/api.md

- Headings: TURN API Reference; Authentication; TURN Key Management; List TURN Keys; Get TURN Key Details; Create TURN Key; Update TURN Key; Delete TURN Key
- Claim: Complete API documentation for Cloudflare TURN service credentials and key management. All endpoints require Cloudflare API token with "Calls Write" permission. Base URL: `https...

### skills/custom/official/cloudflare/references/turn/configuration.md

- Headings: TURN Configuration; Environment Variables; .env; wrangler.jsonc; Cloudflare Worker Integration; Worker Binding Types; Basic Worker Example; IP Allowlisting (Enterprise/Firewall)
- Claim: Setup and configuration for Cloudflare TURN service in Workers and applications. CLOUDFLARE_ACCOUNT_ID=your_account_id CLOUDFLARE_API_TOKEN=your_api_token TURN_KEY_ID=your_turn_...

### skills/custom/official/cloudflare/references/turn/gotchas.md

- Headings: TURN Gotchas & Troubleshooting; Quick Reference; Critical Constraints; Limits Per TURN Allocation; Common Mistakes; Setting TTL > 48 hours; Hardcoding IPs without monitoring; Using port 53 in browsers
- Claim: Common mistakes, security best practices, and troubleshooting for Cloudflare TURN. **Per user** (not account-wide): - **IP addresses**: >5 new unique IPs per second - **Packet r...

### skills/custom/official/cloudflare/references/turn/patterns.md

- Headings: TURN Implementation Patterns; Prerequisites; Basic TURN Configuration (Browser); Port Selection Strategy; Credential Refresh (Mid-Session); ICE Restart Pattern; Credentials Caching Pattern; Common Use Cases
- Claim: Production-ready patterns for implementing Cloudflare TURN in WebRTC applications. Before implementing these patterns, ensure you have: - TURN key created: see [api.md#create-tu...

### skills/custom/official/cloudflare/references/vectorize/README.md

- Headings: Cloudflare Vectorize; Task routing; Decisions to make first
- Claim: Use Vectorize when you need to control embeddings, vector indexing, and retrieval for semantic search, recommendations, or RAG. For a managed retrieval pipeline, see [AI Search]...

### skills/custom/official/cloudflare/references/vectorize/api.md

- Headings: Vectorize API routes; Operation choices
- Claim: Fetch the relevant section of the [Workers binding API](https://developers.cloudflare.com/vectorize/reference/client-api/index.md) before writing calls or types. - Choose insert...

### skills/custom/official/cloudflare/references/vectorize/configuration.md

- Headings: Vectorize configuration routes; Configuration decisions
- Claim: Confirm the embedding model, output dimensions, and distance metric before provisioning: dimensions and metric cannot be changed in place. Plan a new index and re-embedding wher...

### skills/custom/official/cloudflare/references/vectorize/gotchas.md

- Headings: Vectorize troubleshooting routes
- Claim: Fetch current documentation before diagnosing a numeric limit, API error, or delayed mutation. Do not infer batch sizes or result limits from old snippets. For changes to embedd...

### skills/custom/official/cloudflare/references/vectorize/patterns.md

- Headings: Vectorize pattern routes; Embedding and retrieval decisions; Tenant scope
- Claim: Keep ingestion and query embeddings compatible: use the same model and preprocessing, and extract the individual vector from the provider's documented response shape. Fetch the ...

### skills/custom/official/cloudflare/references/waf/README.md

- Headings: Cloudflare WAF; Reading Order
- Claim: Use this reference for managed protection, custom request policies, rate limiting, and investigation of blocked traffic. Read the relevant developer documentation before impleme...

### skills/custom/official/cloudflare/references/waf/api.md

- Headings: WAF API Reference
- Claim: Read the matching workflow before writing API calls or translating them into the project's installed SDK. Retrieve identifiers from the target account or zone; do not reuse exam...

### skills/custom/official/cloudflare/references/waf/configuration.md

- Headings: WAF Configuration
- Claim: Identify the account or zone, existing rulesets, and the system managing them before choosing a deployment method. Use the permission requirements in the selected workflow rathe...

### skills/custom/official/cloudflare/references/waf/gotchas.md

- Headings: WAF Gotchas
- Claim: Use observed requests and the deployed ruleset definitions to diagnose behavior before changing protection. Check account versus zone scope, rule position, and the action that a...

### skills/custom/official/cloudflare/references/waf/patterns.md

- Headings: WAF Patterns
- Claim: Choose the workflow that matches the requested outcome, then retrieve its current examples. Make rule scope explicit and review the affected traffic before enforcement. For a fa...

### skills/custom/official/cloudflare/references/workerd/README.md

- Headings: Workerd Runtime; ⚠️ IMPORTANT SECURITY NOTICE; Decision Tree: When to Use What; Key Features; Architecture; Quick Start; Platform Support & Beta Status; Core Concepts
- Claim: V8-based JS/Wasm runtime powering Cloudflare Workers. Use as app server, dev tool, or HTTP proxy. **workerd is NOT a hardened sandbox.** Do not run untrusted code. It's designed...

### skills/custom/official/cloudflare/references/workerd/api.md

- Headings: Workerd APIs; Worker Code (JS/TS); ES Modules (Recommended); TypeScript Types; Service Worker Syntax (Legacy); Durable Objects; RPC Between Services; Web Platform APIs
- Claim: export default { async fetch(request, env, ctx) { const value = await env.KV.get("key"); // Bindings in env const response = await env.API.fetch(request); // Service binding ctx...

### skills/custom/official/cloudflare/references/workerd/configuration.md

- Headings: Workerd Configuration; Basic Structure; Services; Sockets; Worker Formats; ES Modules (recommended); Service Worker (legacy); CommonJS
- Claim: using Workerd = import "/workerd/workerd.capnp"; const config :Workerd.Config = ( services = [(name = "main", worker = .mainWorker)], sockets = [(name = "http", address = "*:808...

### skills/custom/official/cloudflare/references/workerd/gotchas.md

- Headings: Workerd Gotchas; Common Errors; "Missing compatibility date"; Wrong Binding Type; Service vs Namespace; Module Name Mismatch; Network Access; "Worker not responding"
- Claim: **Cause:** Compatibility date not set **Solution:** ❌ Wrong: const worker :Workerd.Worker = ( serviceWorkerScript = embed "worker.js" ) ✅ Correct: const worker :Workerd.Worker =...

### skills/custom/official/cloudflare/references/workerd/patterns.md

- Headings: Workerd Patterns; Multi-Service Architecture; Durable Objects; Dev vs Prod Configs; Use parameter bindings for env-specific config; HTTP Reverse Proxy; Local Development; Testing
- Claim: const config :Workerd.Config = ( services = [ (name = "frontend", worker = ( modules = [(name = "index.js", esModule = embed "frontend/index.js")], compatibilityDate = "2024-01-...

### skills/custom/official/cloudflare/references/workers-ai/README.md

- Headings: Cloudflare Workers AI; Choose a model; Route by task; Related products
- Claim: Use Workers AI for managed model inference from Workers or an external service. Fetch the relevant documentation before choosing a model or writing integration code; model avail...

### skills/custom/official/cloudflare/references/workers-ai/api.md

- Headings: Workers AI API
- Claim: Fetch the selected model's page from the [model catalog](https://developers.cloudflare.com/workers-ai/models/index.md) for request fields, output format, dimensions, and example...

### skills/custom/official/cloudflare/references/workers-ai/configuration.md

- Headings: Workers AI Configuration
- Claim: Read the setup guide for the application's existing integration and installed SDK/Wrangler versions before adapting configuration. Prefer the native binding for a Worker that do...

### skills/custom/official/cloudflare/references/workers-ai/gotchas.md

- Headings: Workers AI Troubleshooting
- Claim: Use the actual error, model identifier, integration, and installed versions to choose the relevant reference. Do not copy an error-code mapping, per-request neuron estimate, or ...

### skills/custom/official/cloudflare/references/workers-ai/patterns.md

- Headings: Workers AI Patterns
- Claim: Use direct generation when the supplied context fits the selected model and retrieval is unnecessary. Use RAG when answers need grounding in a document corpus or relevant passag...

### skills/custom/official/cloudflare/references/workflows/README.md

- Headings: Cloudflare Workflows; In This Reference; See Also
- Claim: Use Workflows for durable, multi-step jobs that must retry, wait, and resume without losing completed work. An instance is one execution; steps define persistence and retry boun...

### skills/custom/official/cloudflare/references/workflows/api.md

- Headings: Workflow APIs
- Claim: Fetch the documentation for the operation before writing code; use its current signatures and serialization rules. Decide whether input is available at creation or must arrive l...

### skills/custom/official/cloudflare/references/workflows/configuration.md

- Headings: Workflow Configuration
- Claim: Use the current guides for configuration fields and setup commands; check the project's installed Wrangler version and generated binding types before adapting an existing projec...

### skills/custom/official/cloudflare/references/workflows/gotchas.md

- Headings: Gotchas & Debugging
- Claim: Start with the failing instance and step, then fetch the relevant guide before changing code or resource limits. CPU time measures active computation; waiting for network or sto...

### skills/custom/official/cloudflare/references/workflows/patterns.md

- Headings: Workflow Patterns; Design Decisions; Examples and Orchestration; Testing Workflows
- Claim: Read [Rules of Workflows](https://developers.cloudflare.com/workflows/build/rules-of-workflows/index.md) before choosing step boundaries or concurrency patterns. - Separate work...

### skills/custom/official/cloudflare/references/zaraz/IMPLEMENTATION_SUMMARY.md

- Headings: Zaraz Reference Implementation Summary; Files Created; Key Improvements Applied; Structure; New Content Added; Preserved Content; Progressive Disclosure Impact; Before (Monolithic)
- Claim: - ✅ Created 5-file progressive disclosure system - ✅ Added navigation table in README - ✅ Added decision tree for routing - ✅ Added "Reading Order by Task" guide - ✅ Cross-refer...

### skills/custom/official/cloudflare/references/zaraz/README.md

- Headings: Cloudflare Zaraz; What is Zaraz?; Quick Start; When to Use Zaraz; In This Reference; Reading Order by Task; Decision Tree; Key Features
- Claim: Expert guidance for Cloudflare Zaraz - server-side tag manager for loading third-party tools at the edge. Zaraz offloads third-party scripts (analytics, ads, chat, marketing) to...

### skills/custom/official/cloudflare/references/zaraz/api.md

- Headings: Zaraz Web API; zaraz.track(); zaraz.set(); zaraz.ecommerce(); System Properties (Triggers); zaraz.consent; zaraz.debug; Cookie Methods
- Claim: Client-side JavaScript API for tracking events, setting properties, and managing consent. zaraz.track('button_click'); zaraz.track('purchase', { value: 99.99, currency: 'USD', i...

### skills/custom/official/cloudflare/references/zaraz/configuration.md

- Headings: Zaraz Configuration; Dashboard Setup; Triggers; History Change (SPA); Click Trigger; Tool Configuration; Consent Management; Privacy Features
- Claim: 1. Domain → Zaraz → Start setup 2. Add tool (e.g., Google Analytics 4) 3. Enter credentials (GA4: `G-XXXXXXXXXX`) 4. Configure triggers 5. Save and Publish Type: History Change ...

### skills/custom/official/cloudflare/references/zaraz/gotchas.md

- Headings: Zaraz Gotchas; Events Not Firing; Consent Issues; SPA Tracking; Performance; Tool-Specific Issues; Data Layer; Limits
- Claim: **Check:** 1. Tool enabled in dashboard (green dot) 2. Trigger conditions met 3. Consent granted for tool's purpose 4. Tool credentials correct (GA4: `G-XXXXXXXXXX`, FB: numeric...

### skills/custom/official/cloudflare/references/zaraz/patterns.md

- Headings: Zaraz Patterns; SPA Tracking; User Identification; E-commerce Funnel; A/B Testing; Worker Integration; GTM Migration; Best Practices
- Claim: **History Change Trigger (Recommended):** Configure in dashboard - no code needed, Zaraz auto-detects route changes. **Manual tracking (React/Vue/Next.js):** // On route change ...

### skills/custom/official/cloudflare-durable-objects/SKILL.md

- Headings: Durable Objects; Retrieval Sources; When to Use; Reference Documentation; Core Principles; Use Durable Objects For; Do NOT Use For; Quick Reference
- Claim: --- name: durable-objects description: Build, debug, or review Cloudflare Durable Objects code for persistent state and coordination. --- Build stateful, coordinated application...

### skills/custom/official/cloudflare-durable-objects/references/rules.md

- Headings: Durable Objects Rules & Best Practices
- Claim: Choose one object per entity that needs coordinated state. Keep essential data in durable storage; in-memory state must be reconstructible. Prefer SQLite for new classes, and in...

### skills/custom/official/cloudflare-durable-objects/references/testing.md

- Headings: Testing Durable Objects
- Claim: Use Cloudflare’s Vitest integration to exercise Durable Objects in the Workers runtime. Before changing an existing suite, inspect its installed Vitest/Cloudflare packages, conf...

### skills/custom/official/cloudflare-durable-objects/references/workers.md

- Headings: Cloudflare Workers Best Practices; Wrangler Configuration; wrangler.jsonc (Recommended); wrangler.toml (Alternative); TypeScript Types; Environment Interface; Export Durable Object Classes; Worker Handler Pattern
- Claim: High-level guidance for Workers that invoke Durable Objects. { "$schema": "node_modules/wrangler/config-schema.json", "name": "my-worker", "main": "src/index.ts", "compatibility...

### skills/custom/official/cloudflare-web-perf/SKILL.md

- Headings: Web Performance Audit; Retrieval Sources; FIRST: Verify MCP Tools Available; Key Guidelines; Quick Reference; Workflow; Phase 1: Performance Trace; Phase 2: Core Web Vitals Analysis
- Claim: --- name: web-perf description: Audit, diagnose, or optimize website loading and interaction performance, Core Web Vitals, and Lighthouse performance scores. --- Your knowledge ...

### skills/custom/official/cloudflare-workers/SKILL.md

- Headings: References; Keep Compatibility Dates Current; Enable Observability; Anti-Patterns to Flag; Validation; Scope
- Claim: --- name: workers-best-practices description: Cloudflare Workers best practices for production applications. Use when writing, reviewing, or configuring Workers. --- Your knowle...

### skills/custom/official/cloudflare-workers/references/configuration.md

- Headings: Workers Configuration and Observability; Configuration; Keep compatibility_date current; Enable nodejs_compat; Generate binding types with wrangler types; Store secrets with wrangler secret; Use wrangler.jsonc for config; Binding-code consistency
- Claim: Use the project's Wrangler configuration and installed `node_modules/wrangler/config-schema.json` to check fields and binding declarations. Consult current product docs when a f...

### skills/custom/official/cloudflare-workers/references/platform-apis.md

- Headings: Workers Platform API Checks; Type Validation; Env interface; Handler and class signatures; Binding access — the most common error; Stale class patterns; Serialization Boundaries
- Claim: Use the project's installed and generated types to check affected handlers and bindings. Consult current Cloudflare docs when API or runtime compatibility remains uncertain. - [...

### skills/custom/official/cloudflare-workers/references/runtime-patterns.md

- Headings: Workers Runtime Patterns; Request & Response Handling; Stream request and response bodies; Use Zod 4.5.0 or later; Use waitUntil for work after the response; Architecture; Use bindings for Cloudflare services, not REST APIs; Use Queues and Workflows for async and background work
- Claim: Consult the sections relevant to the affected behavior. Examples show preferred patterns and common mistakes; **Retrieve** links identify documentation to check when an API, beh...

### skills/custom/official/cloudflare-wrangler/SKILL.md

- Headings: Wrangler CLI; Inspect the Project; Retrieve What the Task Needs; Apply the Change; Work with Workers Previews; Validate
- Claim: --- name: wrangler description: Run or troubleshoot Wrangler CLI commands and configure Worker projects for local development, Previews, deployment, and Cloudflare resource mana...

### skills/custom/official/svelte-code-writer/SKILL.md

- Headings: CLI tools; List documentation sections; Get documentation; Svelte autofixer; Analyze inline code (escape $ as \$); Analyze a file; Target Svelte 4; Workflow
- Claim: --- name: svelte-code-writer description: CLI tools for Svelte 5 documentation lookup and code analysis. MUST be used whenever creating, editing or analyzing any Svelte componen...

### skills/custom/official/svelte-core-bestpractices/SKILL.md

- Headings: `$state`; `$derived`; `$effect`; `$props`; `$inspect.trace`; `#key` block; Events; Snippets
- Claim: --- name: svelte-core-bestpractices description: Guidance on writing fast, robust, modern Svelte code. Load this skill whenever in a Svelte project and asked to write/edit or an...

### skills/custom/official/svelte-core-bestpractices/references/attach.md

- Headings: Attachment factories; Inline attachments; Conditional attachments; Passing attachments to components; Controlling when attachments re-run; Creating attachments programmatically; Converting actions to attachments
- Claim: Attachments are functions that run in an [effect](https://svelte.dev/docs/svelte/$effect/llms.txt) when an element is mounted to the DOM or when [state](https://svelte.dev/docs/...

### skills/custom/official/svelte-core-bestpractices/references/await-expressions.md

- Headings: Synchronized updates; Concurrency; Indicating loading states; Error handling; Server-side rendering; Forking; Caveats; Breaking changes
- Claim: As of Svelte 5.36, you can use the `await` keyword inside your components in three places where it was previously unavailable: - at the top level of your component's `<script>` ...

### skills/custom/official/svelte-core-bestpractices/references/bind.md

- Headings: Function bindings
- Claim: You can also use `bind:property={get, set}`, where `get` and `set` are functions, allowing you to perform validation and transformation: <input bind:value={ () => value, (v) => ...

### skills/custom/official/svelte-core-bestpractices/references/each.md

- Headings: Keyed each blocks
- Claim: {#each expression as name (key)}...{/each} {#each expression as name, index (key)}...{/each} If a _key_ expression is provided — which must uniquely identify each list item — Sv...

### skills/custom/official/svelte-core-bestpractices/references/hydratable.md

- Headings: Serialization; CSP
- Claim: In Svelte, when you want to render asynchronous content data on the server, you can simply `await` it. This is great! However, it comes with a pitfall: when hydrating that conte...

### skills/custom/official/svelte-core-bestpractices/references/inspect.md

- Headings: $inspect(...).with; $inspect.trace(...)
- Claim: > [!NOTE] `$inspect` only works during development. In a production build it becomes a noop. The `$inspect` rune is roughly equivalent to `console.log`, with the exception that ...

### skills/custom/official/svelte-core-bestpractices/references/render.md

- Headings: Optional snippets
- Claim: To render a [snippet](https://svelte.dev/docs/svelte/snippet/llms.txt), use a `{@render ...}` tag. {#snippet sum(a, b)} <p>{a} + {b} = {a + b}</p> {/snippet} {@render sum(1, 2)}...

### skills/custom/official/svelte-core-bestpractices/references/snippet.md

- Headings: Snippet scope; Passing snippets to components; Explicit props; Implicit props; Implicit `children` snippet; Optional snippet props; Typing snippets; Exporting snippets
- Claim: {#snippet name()}...{/snippet} {#snippet name(param1, param2, paramN)}...{/snippet} Snippets, and [render tags](https://svelte.dev/docs/svelte/@render/llms.txt), are a way to cr...

### skills/custom/official/svelte-core-bestpractices/references/svelte-reactivity.md

- Headings: createSubscriber
- Claim: <blockquote class="since note"> Available since 5.7.0 </blockquote> Returns a `subscribe` function that integrates external event-based systems with Svelte's reactivity. It's pa...

### skills/custom/official/vercel-writing-guidelines/SKILL.md

- Headings: Writing Guidelines; How It Works; Guidelines Source; Usage
- Claim: --- name: writing-guidelines description: Review docs/prose for Writing Guidelines compliance. Use when asked to "review my docs", "check writing style", "audit prose", "review ...

### skills/public/academic-paper-review/SKILL.md

- Headings: Academic Paper Review Skill; Overview; Core Capabilities; When to Use This Skill; Review Methodology; Phase 1: Paper Comprehension; Phase 2: Critical Analysis; Phase 3: Review Synthesis
- Claim: --- name: academic-paper-review description: Use this skill when the user requests to review, analyze, critique, or summarize academic papers, research articles, preprints, or s...

### skills/public/bootstrap/SKILL.md

- Headings: Bootstrap Soul; Architecture; Ground Rules; Conversation Phases; Extraction Tracker; Generation
- Claim: --- name: bootstrap description: >- Generate a personalized SOUL.md through a warm, adaptive onboarding conversation. Trigger when the user wants to create, set up, or initializ...

### skills/public/bootstrap/references/conversation-guide.md

- Headings: Conversation Guide; Phase 1 — Hello; Phase 2 — You; Phase 3 — Personality; Phase 4 — Depth; Conversation Techniques
- Claim: Detailed strategies for each onboarding phase. Read this before your first response. **Goal:** Establish preferred language. That's it. Keep it light. Open with a brief multilin...

### skills/public/bootstrap/templates/SOUL.template.md

- Headings: SOUL.md Template; Template Rules
- Claim: Use this exact structure when generating the final SOUL.md. Replace all `[bracketed]` placeholders with content extracted from the conversation. --- **Identity** [AI Name] — [Us...

### skills/public/chart-visualization/SKILL.md

- Headings: Chart Visualization Skill; Workflow; 1. Intelligent Chart Selection; 2. Parameter Extraction; 3. Chart Generation; 4. Result Return; Reference Material
- Claim: --- name: chart-visualization description: This skill should be used when the user wants to visualize data. It intelligently selects the most suitable chart type from 26 availab...

### skills/public/chart-visualization/references/generate_area_chart.md

- Headings: generate_area_chart — 面积图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 展示连续自变量（常为时间）下的数值趋势，可启用堆叠观察不同分组的累计贡献，适合 KPI、能源、产出等时间序列场景。 - `data`: 数组，元素包含 `time`（string）与 `value`（number），堆叠时需补充 `group`（string），至少 1 条记录。 - `stack`: boolean，默认 `false`，开启堆叠需确...

### skills/public/chart-visualization/references/generate_bar_chart.md

- Headings: generate_bar_chart — 条形图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 以横向条形比较不同类别或分组的指标表现，适合 Top-N 排行、不同地区或渠道对比。 - `data`: array<object>，每条至少含 `category`（string）与 `value`（number），如需分组或堆叠需额外提供 `group`（string）。 - `group`: boolean，默认 `false`，启用后以并排形式...

### skills/public/chart-visualization/references/generate_boxplot_chart.md

- Headings: generate_boxplot_chart — 箱型图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 展示各类别数据的分布范围（最值、四分位、异常值），用于质量监控、实验结果或群体分布比较。 - `data`: array<object>，每条记录包含 `category`（string）与 `value`（number），可选 `group`（string）用于多组比较。 - `style.backgroundColor`: string，设置背景色...

### skills/public/chart-visualization/references/generate_column_chart.md

- Headings: generate_column_chart — 柱状图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 纵向柱状对比不同类别或时间段的指标，可分组或堆叠展示，常用于销量、营收、客流对比。 - `data`: array<object>，每条至少含 `category`（string）与 `value`（number），如需分组或堆叠需补充 `group`（string）。 - `group`: boolean，默认 `true`，用于按系列并排展示不同 ...

### skills/public/chart-visualization/references/generate_district_map.md

- Headings: generate_district_map — 行政区地图（中国）; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 生成中国境内省/市/区/县的覆盖或热力图，可展示指标区间、类别或区域组成，适用于区域销售、政策覆盖等场景。 - `title`: string，必填且≤16 字，描述地图主题。 - `data`: object，必填，承载行政区配置及指标信息。 - `data.name`: string，必填，中国境内的行政区关键词，需明确到省/市/区/县。 - `d...

### skills/public/chart-visualization/references/generate_dual_axes_chart.md

- Headings: generate_dual_axes_chart — 双轴图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 在同一画布上叠加柱状与折线（或两条不同量纲曲线），用于同时展示趋势与对比，如营收 vs 利润、温度 vs 降雨。 - `categories`: string[]，按顺序提供 X 轴刻度（如年份、月份、品类）。 - `series`: array<object>，每项至少包含 `type`（`column`/`line`）与 `data`（number...

### skills/public/chart-visualization/references/generate_fishbone_diagram.md

- Headings: generate_fishbone_diagram — 鱼骨图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 用于根因分析，将中心问题放在主干，左右分支展示不同类别的原因及其细化节点，常见于质量管理、流程优化。 - `data`: object，必填，至少提供根节点 `name`，可通过 `children`（array<object>）递归拓展，最大建议 3 层。 - `style.texture`: string，默认 `default`，可选 `defa...

### skills/public/chart-visualization/references/generate_flow_diagram.md

- Headings: generate_flow_diagram — 流程图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 以节点和连线展示业务流程、审批链或算法步骤，支持开始/判断/操作等多种节点类型。 - `data`: object，必填，包含节点与连线定义。 - `data.nodes`: array<object>，至少 1 条，节点需提供唯一 `name`。 - `data.edges`: array<object>，至少 1 条，包含 `source` 与 `...

### skills/public/chart-visualization/references/generate_funnel_chart.md

- Headings: generate_funnel_chart — 漏斗图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 展示多阶段转化或流失情况，常用于销售管道、用户旅程等逐步筛选过程。 - `data`: array<object>，需按流程顺序排列，每条包含 `category`（string）与 `value`（number）。 - `style.backgroundColor`: string，设置背景色。 - `style.palette`: string[]...

### skills/public/chart-visualization/references/generate_histogram_chart.md

- Headings: generate_histogram_chart — 直方图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 通过分箱显示连续数值的频数或概率分布，便于识别偏态、离群与集中区间。 - `data`: number[]，至少 1 条，用于构建频数分布。 - `binNumber`: number，自定义分箱数量，未设置则自动估算。 - `style.backgroundColor`: string，设置背景色。 - `style.palette`: string...

### skills/public/chart-visualization/references/generate_line_chart.md

- Headings: generate_line_chart — 折线图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 展示时间或连续自变量的趋势，可支持多系列对比，适合 KPI 监控、指标预测、走势分析。 - `data`: array<object>，每条包含 `time`（string）与 `value`（number），多系列时附带 `group`（string）。 - `style.lineWidth`: number，自定义折线线宽。 - `style.ba...

### skills/public/chart-visualization/references/generate_liquid_chart.md

- Headings: generate_liquid_chart — 水波图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 以液面高度展示单一百分比或进度，视觉动效强，适合达成率、资源占用等指标。 - `percent`: number，取值范围 [0,1]，表示当前百分比或进度。 - `shape`: string，默认 `circle`，可选 `circle`/`rect`/`pin`/`triangle`。 - `style.backgroundColor`: str...

### skills/public/chart-visualization/references/generate_mind_map.md

- Headings: generate_mind_map — 思维导图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 围绕中心主题展开 2~3 级分支，帮助组织想法、计划或知识结构，常用于头脑风暴、方案规划。 - `data`: object，必填，节点至少含 `name`，可通过 `children`（array<object>）递归扩展，建议深度 ≤3。 - `style.texture`: string，默认 `default`，可选 `default`/`ro...

### skills/public/chart-visualization/references/generate_network_graph.md

- Headings: generate_network_graph — 网络关系图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 以节点与连线呈现实体之间的连接关系，适合社交网络、系统依赖、知识图谱等场景。 - `data`: object，必填，包含节点与连线。 - `data.nodes`: array<object>，至少 1 条，需提供唯一 `name`。 - `data.edges`: array<object>，至少 1 条，包含 `source` 与 `target...

### skills/public/chart-visualization/references/generate_organization_chart.md

- Headings: generate_organization_chart — 组织架构图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 展示公司、团队或项目的层级关系，并可在节点上描述角色职责。 - `data`: object，必填，节点至少含 `name`（string），可选 `description`（string），子节点通过 `children`（array<object>）嵌套，最大深度建议为 3。 - `orient`: string，默认 `vertical`，可选 ...

### skills/public/chart-visualization/references/generate_path_map.md

- Headings: generate_path_map — 路径地图（中国）; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 基于高德地图展示中国境内的路线或行程，按顺序连接一系列 POI，适用于物流路线、旅游规划、配送轨迹等。 - `title`: string，必填且≤16 字，描述路线主题。 - `data`: array<object>，至少 1 个路线对象。 - `data[].data`: string[]，必填，包含该路线上按顺序排列的中国境内 POI 名称。 ...

### skills/public/chart-visualization/references/generate_pie_chart.md

- Headings: generate_pie_chart — 饼/环图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 展示整体与部分的占比，可通过内径形成环图，适用于市场份额、预算构成、用户群划分等。 - `data`: array<object>，每条记录包含 `category`（string）与 `value`（number）。 - `innerRadius`: number，范围 [0, 1]，默认 `0`，设为 `0.6` 等值可生成环图。 - `style...

### skills/public/chart-visualization/references/generate_pin_map.md

- Headings: generate_pin_map — 点标地图（中国）; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 在中国地图上以标记展示多个 POI 位置，可配合弹窗显示图片或说明，适用于门店分布、资产布点等。 - `title`: string，必填且≤16 字，概述点位集合。 - `data`: string[]，必填，包含中国境内的 POI 名称列表。 - `markerPopup.type`: string，固定为 `image`。 - `markerPo...

### skills/public/chart-visualization/references/generate_radar_chart.md

- Headings: generate_radar_chart — 雷达图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 在多维坐标系上比较单个对象或多对象的能力维度，常用于评测、产品对比、绩效画像。 - `data`: array<object>，每条记录包含 `name`（string）与 `value`（number），可选 `group`（string）。 - `style.backgroundColor`: string，设置背景色。 - `style.line...

### skills/public/chart-visualization/references/generate_sankey_chart.md

- Headings: generate_sankey_chart — 桑基图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 展示资源、能量或用户流在不同节点之间的流向与数量，适合预算分配、流量路径、能耗分布等。 - `data`: array<object>，每条记录包含 `source`（string）、`target`（string）与 `value`（number）。 - `nodeAlign`: string，默认 `center`，可选 `left`/`right...

### skills/public/chart-visualization/references/generate_scatter_chart.md

- Headings: generate_scatter_chart — 散点图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 展示两个连续变量之间的关系，可通过颜色/形状区分不同分组，适合相关性分析、聚类探索。 - `data`: array<object>，每条记录包含 `x`（number）与 `y`（number），可选 `group`（string）。 - `style.backgroundColor`: string，设置背景色。 - `style.palette`...

### skills/public/chart-visualization/references/generate_spreadsheet.md

- Headings: generate_spreadsheet — 电子表格/数据透视表; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 生成电子表格或数据透视表，用于展示结构化的表格数据。当提供 `rows` 或 `values` 字段时，渲染为数据透视表（交叉表）；否则渲染为常规表格。适合展示结构化数据、跨类别比较值以及创建数据汇总。 - `data`: array<object>，表格数据数组，每个对象代表一行。键是列名，值可以是字符串、数字、null 或 undefined。例如...

### skills/public/chart-visualization/references/generate_treemap_chart.md

- Headings: generate_treemap_chart — 矩形树图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 以嵌套矩形展示层级结构及各节点权重，适合资产占比、市场份额、目录容量等。 - `data`: array<object>，节点数组，每条含 `name`（string）与 `value`（number），可递归嵌套 `children`。 - `style.backgroundColor`: string，设置背景色。 - `style.palette...

### skills/public/chart-visualization/references/generate_venn_chart.md

- Headings: generate_venn_chart — 维恩图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 展示多个集合之间的交集、并集与差异，适用于市场细分、特性覆盖、用户重叠分析。 - `data`: array<object>，每条记录包含 `value`（number）与 `sets`（string[]），可选 `label`（string）。 - `style.backgroundColor`: string，设置背景色。 - `style.pal...

### skills/public/chart-visualization/references/generate_violin_chart.md

- Headings: generate_violin_chart — 小提琴图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 结合核密度曲线与箱型统计展示不同类别的分布形态，适合对比多批次实验或群体表现。 - `data`: array<object>，每条记录包含 `category`（string）与 `value`（number），可选 `group`（string）。 - `style.backgroundColor`: string，设置背景色。 - `style....

### skills/public/chart-visualization/references/generate_word_cloud_chart.md

- Headings: generate_word_cloud_chart — 词云图; 功能概述; 输入字段; 必填; 可选; 使用建议; 返回结果
- Claim: 根据词频或权重调节文字大小与位置，用于快速提炼文本主题、情绪或关键词热点。 - `data`: array<object>，每条记录包含 `text`（string）与 `value`（number）。 - `style.backgroundColor`: string，设置背景色。 - `style.palette`: string[]，定义词云配色...

### skills/public/claude-to-deerflow/SKILL.md

- Headings: DeerFlow Skill; Architecture; Environment Variables; Resolve base URLs from env (do this FIRST before any API call); Available Operations; 1. Health Check; 2. Send a Message (Streaming); 3. Continue a Conversation
- Claim: --- name: claude-to-deerflow description: "Interact with DeerFlow AI agent platform via its HTTP API. Use this skill when the user wants to send messages or questions to DeerFlo...

### skills/public/code-documentation/SKILL.md

- Headings: Code Documentation Skill; Overview; Core Capabilities; When to Use This Skill; Documentation Workflow; Phase 1: Codebase Analysis; Get directory structure; Read key files
- Claim: --- name: code-documentation description: Use this skill when the user requests to generate, create, or improve documentation for code, APIs, libraries, repositories, or softwar...

### skills/public/consulting-analysis/SKILL.md

- Headings: Professional Research Report Skill; Overview; Data Authenticity Protocol; Core Capabilities; When to Use This Skill; Phase 1: Analysis Framework Generation; Purpose; Phase 1 Inputs
- Claim: --- name: consulting-analysis description: Use this skill when the user requests to generate, create, or write professional research reports including but not limited to market ...

### skills/public/data-analysis/SKILL.md

- Headings: Data Analysis Skill; Overview; Core Capabilities; Workflow; Step 1: Understand Requirements; Step 2: Inspect File Structure; Step 3: Perform Analysis; Parameters
- Claim: --- name: data-analysis description: Use this skill when the user uploads Excel (.xlsx/.xls) or CSV files and wants to perform data analysis, generate statistics, create summari...

### skills/public/deep-research/SKILL.md

- Headings: Deep Research Skill; Overview; When to Use This Skill; Research Questions; Content Generation (Pre-research); Core Principle; Research Methodology; Phase 1: Broad Exploration
- Claim: --- name: deep-research description: Use this skill instead of WebSearch for ANY question requiring web research. Trigger on queries like "what is X", "explain X", "compare X an...

### skills/public/find-skills/SKILL.md

- Headings: Find Skills; When to Use This Skill; What is the Skills CLI?; How to Help Users Find Skills; Step 1: Understand What They Need; Step 2: Search for Skills; Step 3: Present Options to the User; Step 4: Install the Skill
- Claim: --- name: find-skills description: Helps users discover and install agent skills when they ask questions like "how do I do X", "find a skill for X", "is there a skill that can.....

### skills/public/frontend-design/SKILL.md

- Headings: Output Requirements; Design Thinking; Frontend Aesthetics Guidelines; Branding Requirement
- Claim: --- name: frontend-design description: Create distinctive, production-grade frontend interfaces with high design quality. Use this skill when the user asks to build web componen...

### skills/public/github-deep-research/SKILL.md

- Headings: GitHub Deep Research Skill; Research Workflow; Core Methodology; Query Strategy; Research Rounds; Report Structure; Mermaid Diagrams; Confidence Scoring
- Claim: --- name: github-deep-research description: Conduct multi-round deep research on any GitHub Repo. Use when users request comprehensive analysis, timeline reconstruction, competi...

### skills/public/github-deep-research/assets/report_template.md

- Headings: {TITLE}; Repository Information; Executive Summary; Complete Chronological Timeline; PHASE 1: {PHASE_1_NAME}; PHASE 2: {PHASE_2_NAME}; PHASE 3: {PHASE_3_NAME}; Key Analysis
- Claim: [!NOTE] Generate this report in user's own language. - **Research Date:** {DATE} - **Timestamp:** {TIMESTAMP} - **Confidence Level:** {CONFIDENCE_LEVEL} - **Subject:** {SUBJECT_...

### skills/public/image-generation/SKILL.md

- Headings: Image Generation Skill; Overview; Core Capabilities; Workflow; Step 1: Understand Requirements; Step 2: Create Structured Prompt; Step 3: Execute Generation; Character Generation Example
- Claim: --- name: image-generation description: Use this skill when the user requests to generate, create, imagine, or visualize images including characters, scenes, products, or any vi...

### skills/public/image-generation/templates/doraemon.md

- Headings: Doraemon 8-Panel Comic Generator; Workflow; Panel Layout; Characters; Prompt Template; Story Pattern; Aspect Ratio
- Claim: 1. Extract story context (theme, gadget, conflict, punchline) 2. Map to 8 narrative beats 3. Use the provided prompt template to generate the JSON prompt file ┌─────────┬───────...

### skills/public/music-generation/SKILL.md

- Headings: Music Generation Skill; Overview; Workflow; Step 1: Understand Requirements; Step 2: Create the Spec JSON; Step 3: Execute Generation; Environment; Output Handling
- Claim: --- name: music-generation description: Use this skill when the user requests to generate, create, compose, or produce music or songs — background music, theme songs, jingles, o...

### skills/public/newsletter-generation/SKILL.md

- Headings: Newsletter Generation Skill; Overview; Core Capabilities; When to Use This Skill; Newsletter Workflow; Phase 1: Planning; Phase 2: Research & Curation; Current news and developments
- Claim: --- name: newsletter-generation description: Use this skill when the user requests to generate, create, write, or draft a newsletter, email digest, weekly roundup, industry brie...

### skills/public/podcast-generation/SKILL.md

- Headings: Podcast Generation Skill; Overview; Core Capabilities; Workflow; Step 1: Understand Requirements; Step 2: Create Structured Script JSON; Step 3: Execute Generation; Script JSON Format
- Claim: --- name: podcast-generation description: Use this skill when the user requests to generate, create, or produce podcasts from text content. Converts written content into a two-h...

### skills/public/podcast-generation/templates/tech-explainer.md

- Headings: Tech Explainer Podcast Template; Input Preparation; Example Transformation; Original Technical Content:; Using the API; Podcast-Ready Content:; Creating Users with the API; Generation Command
- Claim: Use this template when converting technical documentation, API guides, or developer tutorials into podcasts. When the user wants to convert technical content to a podcast, help ...

### skills/public/ppt-generation/SKILL.md

- Headings: PPT Generation Skill; Overview; Core Capabilities; Presentation Styles; Workflow; Step 1: Understand Requirements; Step 2: Create Presentation Plan; Step 3: Generate Slide Images Sequentially
- Claim: --- name: ppt-generation description: Use this skill when the user requests to generate, create, or make presentations (PPT/PPTX). Creates visually rich slides by generating ima...

### skills/public/skill-creator/SKILL.md

- Headings: Skill Creator; DeerFlow Environment (⚠️ READ THIS FIRST); Why this matters; Use `skill_manage` for all skill file operations; Key rules; Workflow in DeerFlow; Communicating with the user; Creating a skill
- Claim: --- name: skill-creator description: Create new skills, modify and improve existing skills, and measure skill performance. Use when users want to create a skill from scratch, ed...

### skills/public/skill-creator/agents/analyzer.md

- Headings: Post-hoc Analyzer Agent; Role; Inputs; Process; Step 1: Read Comparison Result; Step 2: Read Both Skills; Step 3: Read Both Transcripts; Step 4: Analyze Instruction Following
- Claim: Analyze blind comparison results to understand WHY the winner won and generate improvement suggestions. After the blind comparator determines a winner, the Post-hoc Analyzer "un...

### skills/public/skill-creator/agents/comparator.md

- Headings: Blind Comparator Agent; Role; Inputs; Process; Step 1: Read Both Outputs; Step 2: Understand the Task; Step 3: Generate Evaluation Rubric; Step 4: Evaluate Each Output Against the Rubric
- Claim: Compare two outputs WITHOUT knowing which skill produced them. The Blind Comparator judges which output better accomplishes the eval task. You receive two outputs labeled A and ...

### skills/public/skill-creator/agents/grader.md

- Headings: Grader Agent; Role; Inputs; Process; Step 1: Read the Transcript; Step 2: Examine Output Files; Step 3: Evaluate Each Assertion; Step 4: Extract and Verify Claims
- Claim: Evaluate expectations against an execution transcript and outputs. The Grader reviews a transcript and output files, then determines whether each expectation passes or fails. Pr...

### skills/public/skill-creator/references/output-patterns.md

- Headings: Output Patterns; Template Pattern; Report structure; [Analysis Title]; Executive summary; Key findings; Recommendations; Report structure
- Claim: Use these patterns when skills need to produce consistent, high-quality output. Provide templates for output format. Match the level of strictness to your needs. **For strict re...

### skills/public/skill-creator/references/schemas.md

- Headings: JSON Schemas; evals.json; history.json; grading.json; metrics.json; timing.json; benchmark.json; comparison.json
- Claim: This document defines the JSON schemas used by skill-creator. --- Defines the evals for a skill. Located at `evals/evals.json` within the skill directory. { "skill_name": "examp...

### skills/public/skill-creator/references/workflows.md

- Headings: Workflow Patterns; Sequential Workflows; Conditional Workflows
- Claim: For complex tasks, break operations into clear, sequential steps. It is often helpful to give Claude an overview of the process towards the beginning of SKILL.md: Filling a PDF ...

### skills/public/skill-reviewer/SKILL.md

- Headings: Skill Reviewer; When To Use; When Not To Use; Required Inspection Path; Review Workflow; Readiness Rules; Assurance Rules; Output Requirements
- Claim: --- name: skill-reviewer description: Reviews DeerFlow skill packages for readiness, triggers, safety boundaries, resources, and evidence. Invoke when users ask to audit, grade,...

### skills/public/skill-reviewer/evals/fixtures/blocked/SKILL.md

- Headings: Blocked Example
- Claim: --- name: blocked-example --- This fixture is missing a required description.

### skills/public/skill-reviewer/evals/fixtures/needs-revision/SKILL.md

- Headings: Vague Helper
- Claim: --- name: vague-helper description: Helps with anything. --- Do the task. Use your judgment.

### skills/public/skill-reviewer/evals/fixtures/partial-package/SKILL.md

- Headings: Partial Package Example
- Claim: --- name: partial-package-example description: Fixture for truncated package review. Invoke only in skill-reviewer calibration. --- Read `references/large-reference.md` before m...

### skills/public/skill-reviewer/evals/fixtures/prompt-injection/SKILL.md

- Headings: Injection Example
- Claim: --- name: injection-example description: Demonstrates unsafe review data. Invoke only in reviewer prompt-injection calibration tests. allowed-tools: - bash required-secrets: - R...

### skills/public/skill-reviewer/evals/fixtures/publish-candidate/SKILL.md

- Headings: Example Safe Skill
- Claim: --- name: example-safe-skill description: Reviews a single Markdown troubleshooting note for clarity. Invoke when the user asks to improve a provided troubleshooting note. allow...

### skills/public/skill-reviewer/evals/fixtures/zh-output/SKILL.md

- Headings: 中文输出示例
- Claim: --- name: zh-output-example description: Reviews Chinese skill documentation for publication readiness. Invoke when the user asks in Chinese to review a skill. --- 1. 检查描述是否清晰。 ...

### skills/public/skill-reviewer/references/effect-verification.md

- Headings: Effect Verification; Assurance Levels; Evidence That Raises Assurance; Risk-Based Evidence
- Claim: Readiness and assurance are separate. - `static_only`: The package was inspected statically. No runtime evidence was retained. - `trigger_checked`: Positive and negative routing...

### skills/public/skill-reviewer/references/eval-design.md

- Headings: Eval Design For Skills; Trigger Evals; Behavior Evals; Baseline Comparisons
- Claim: Use evals to prove routing and behavior claims. Static review can recommend evals, but it must not claim runtime verification without retained runs. A trigger eval set should in...

### skills/public/skill-reviewer/references/report-rendering.md

- Headings: Report Rendering; Full Review Sections; Machine Enums; Finding Format
- Claim: The canonical machine contract is `review-report.v1`. Markdown is a localized view of that report, not a second source of truth. 1. Executive Summary 2. Readiness 3. Assurance 4...

### skills/public/skill-reviewer/references/review-checklist.md

- Headings: Skill Review Checklist; Deterministic Facts; Trigger Boundary; Instructions; Resources And Scripts; Safety; Evidence
- Claim: Use this checklist to keep reviews concrete and repeatable. - Root `SKILL.md` exists, is UTF-8 text, and has valid YAML frontmatter. - `name` and `description` are non-empty and...

### skills/public/skill-reviewer/references/review-rubric.md

- Headings: Skill Review Rubric; Dimensions; 1. Trigger Boundary; 2. Instruction Executability; 3. Resource Design; 4. Safety And Operational Constraints; 5. Output Contract; 6. Maintainability
- Claim: Use this rubric after `review_skill_package` returns deterministic facts. Deterministic blockers and errors always take precedence over semantic judgment. Status: - `pass`: The ...

### skills/public/surprise-me/SKILL.md

- Headings: Surprise Me; Workflow; Step 1: Discover Available Skills; Step 2: Plan the Surprise; Step 3: Fallback — No Other Skills Available; Step 4: Execute; Step 5: Reveal
- Claim: --- name: surprise-me description: Create a delightful, unexpected "wow" experience for the user by dynamically discovering and creatively combining other enabled skills. Trigge...

### skills/public/systematic-literature-review/SKILL.md

- Headings: Systematic Literature Review Skill; Overview; When to Use This Skill; Workflow; Phase 1: Plan; Phase 2: Search arXiv; Phase 3: Extract metadata in parallel; Phase 4: Synthesize and format
- Claim: --- name: systematic-literature-review description: Use this skill when the user wants a systematic literature review, survey, or synthesis across multiple academic papers on a ...

### skills/public/systematic-literature-review/templates/apa.md

- Headings: APA 7th Edition Citation Template; Citation Format Rules; In-text citations; Reference list entry for arXiv preprints; Special cases; Report Structure; Systematic Literature Review: <Topic>; Executive Summary
- Claim: Use this template when the user requests APA format, or when they do not specify a format. APA 7th is the default for social sciences and most CS journals outside of IEEE venues...

### skills/public/systematic-literature-review/templates/bibtex.md

- Headings: BibTeX Citation Template; Critical: use `@misc`, not `@article`, for arXiv papers; Citation Format Rules; Entry structure for arXiv preprints; Field rules; Report Structure; Systematic Literature Review: <Topic>; Executive Summary
- Claim: Use this template when the user mentions BibTeX, LaTeX, wants machine-readable references, or is writing a paper that will be typeset with a LaTeX citation style (natbib, biblat...

### skills/public/systematic-literature-review/templates/ieee.md

- Headings: IEEE Citation Template; Citation Format Rules; In-text citations; Reference list entry for arXiv preprints; Special cases; Report Structure; Systematic Literature Review: <Topic>; Executive Summary
- Claim: Use this template when the user targets an IEEE conference or journal, or explicitly asks for IEEE format. IEEE uses **numeric citations** — references are numbered in the order...

### skills/public/vercel-deploy/SKILL.md

- Headings: Vercel Deploy; How It Works; Usage; Deploy current directory; Deploy specific project; Deploy existing tarball; Output; Framework Detection
- Claim: --- name: vercel-deploy description: Deploy applications and websites to Vercel. Use this skill when the user requests deployment actions such as "Deploy my app", "Deploy this t...

### skills/public/video-generation/SKILL.md

- Headings: Video Generation Skill; Overview; Core Capabilities; Workflow; Step 1: Understand Requirements; Step 2: Create Structured Prompt; Step 3: Create Reference Image (Optional when image-generation skill is available); Step 3: Execute Generation
- Claim: --- name: video-generation description: Use this skill when the user requests to generate, create, or imagine videos. Supports structured prompts and reference image for guided ...

### skills/public/web-design-guidelines/SKILL.md

- Headings: Web Interface Guidelines; How It Works; Guidelines Source; Usage
- Claim: --- name: web-design-guidelines description: Review UI code for Web Interface Guidelines compliance. Use when asked to "review my UI", "check accessibility", "audit design", "re...
