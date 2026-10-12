# Context Relativity: Linking Niches, Profiles, and Traits

To ensure smooth routing and execution, the agent-mesh must understand how discrete pieces of knowledge link together. We achieve this by mapping context memory traits by "relativity" (semantic proximity and dependency graphs), rather than just flat keyword matching.

## 1. The Relativity Graph

Knowledge in the harness is not isolated. It operates on a Graph linking:
- **Profiles (Agents)**: E.g., Lead Orchestrator, Database Architect, Frontend Scribe.
- **Niche Skills**: E.g., `stripe-billing`, `polar-sh-api`, `postgres-migrations`.
- **Project Context**: Current stack, previous architectural decisions, current directory state.
- **Memory Traits**: Long-term lessons learned (e.g., "Last time we used Edge runtime, we hit a bundle limit").

### Edge Weights (How things connect)
When an agent requests knowledge, the router algorithms navigate this graph using weighted connections:
1. **Direct Dependency**: A skill explicitly requires another (e.g., `nextjs-app-router` heavily links to `vercel-react-best-practices`).
2. **Semantic Proximity**: The local embedding model (`Qwen3-Embedding-8B`) calculates how close the current task is to stored Memory Traits.
3. **Environment State**: The router prioritizes skills that match the detected workspace language (Go, TS, Rust).

## 2. Highly Aware Routing Algorithms

The router doesn't just send tasks; it ensures "everything plays along smoothly" by checking the Relativity Graph before dispatch:

- **Context Pre-Fetching**: If the router delegates a task requiring Stripe implementation, it automatically pre-loads the `stripe-billing` niche knowledge and the `security-review` ECC agent. 
- **Conflict Avoidance**: If the graph detects a conflict (e.g., a trait saying "Do not use Prisma" alongside a task mentioning Prisma), the router intercepts the task and routes it to the Lead Orchestrator to resolve the discrepancy before a coding worker touches it.

## 3. Optimizing the Niche Explorer

The Niche Explorer/Loader is a CPU-bound classifier combined with the embedding engine. 
- **Relativity Joins**: When a trait is retrieved, the loader performs a "relativity join," fetching the primary document AND its immediate highest-weighted graph neighbors (e.g., fetching a specific API endpoint doc, plus the global authentication standard for that API).
- **Token Efficiency**: These joins are then condensed by the `auto/coding:cheap` route into a single, cohesive brief before being passed to the worker agent, guaranteeing high-functioning memory without bloated contexts.
