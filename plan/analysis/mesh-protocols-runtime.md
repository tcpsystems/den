# Agent-Mesh Protocols and Agnostic Runtime

This expands the harness foundation to ensure universal compatibility, highly optimized communication, and data-driven hardware scaling.

## Protocol Hardening (Agent-Mesh)
- **Modern Communication Standards**: 
  - **A2A**: Governs peer-to-peer agent collaboration and delegation. Reuses HTTP, JSON-RPC 2.0, and Server-Sent Events, using Protocol Buffer messages for its canonical data model.
  - **MCP**: Connects agents to tools and data via JSON-RPC 2.0. Standardizes the "how-to" for capabilities.
  - **Protobuf/buf**: Supplies a compact, evolvable wire format with strict breaking-change detection.
  - **AT Protocol**: Contributes decentralized identity (DIDs) and synchronized signed repositories over HTTP (XRPC) and WebSockets.
- **Design Goals**: Optimal, efficient, measured, benchmarked, hardened, and security-first with no UX/UI pain points or knowledge gaps.

## Fully Agnostic Runtime Daemon
- **Universal Compatibility**: The daemon must operate strictly agnostic to the model, agent, LLM, provider, OS, environment, network, firewall, region, project, product, and industry.
- **Resource Mastery**: Built for Docker compatibility, extreme resource/token efficiency, and graceful fallback. It must adjust and excel regardless of the host environment.

## Local Compute and Hardware Baselines
- **Local Runtimes**: The harness can automatically configure local embedding engines like **Ollama** (lighter on-device path, exposing an `/api/embed` API) and **vLLM** (high-throughput, exposing an OpenAI-compatible `/v1/embeddings` API).
- **Data-Driven Hardware Scaling**: 
  - Realistic local baseline derived from hardware surveys: a **4–8 vCPU harness with 8–16 GB RAM** plus a discrete **~8 GB GPU** to hold a full 8B embedding model (like Qwen3-Embedding-8B).
  - This perfectly aligns with the modal hardware specs: 8-core CPU, 8-16 GB VRAM, 16-32 GB system RAM.

## Continuous Exploration
- **Analytics and Curiosity**: The harness relies on analytics-heavy, inspiration-driven exploration. Looping is halted, but deliberate, curiosity-driven exploration expands the harness's capabilities over time.
