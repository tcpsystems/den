# Skill router

## Question

Should den train a model to pick skills, or keep description match? Also cover auto-find and auto-inject.

## Sources

- Agent Skills specification: https://agentskills.io/specification
- Learned routing, SkillOrchestra: https://arxiv.org/html/2602.19672v1
- Auto-find and auto-inject, OpenHands: https://docs.openhands.dev/sdk/guides/skill

## Finding

Keep description match. Do not train a skill picker.

The spec requires `name` and `description`. The description says what the skill does and when to use it, and it should carry keywords. Progressive disclosure is three steps: `name` and `description` for every skill at startup, the `SKILL.md` body only after activation, and `scripts/`, `references/`, or `assets/` only when a step needs them. Once the agent activates a skill, it loads that whole file. That is one body.

OpenHands lists Agent Skills descriptions in `<available_skills>`. With no trigger, the model calls `invoke_skill()`. With keyword triggers, a match auto-injects the body into the user message. `load_project_skills` auto-finds `AGENTS.md`, `CLAUDE.md`, and `GEMINI.md` at the workspace root and loads them as project context. A legacy skill with no trigger puts its full body in the system prompt on every turn.

SkillOrchestra does not pick a `SKILL.md`. It learns a Skill Handbook from traces: an LLM names skills, Beta profiles estimate each agent's success and cost, and the orchestrator routes to an agent. Optional retrieval adds the k nearest skills in an embedding space. A cosine ranker always has a neighbor, so it cannot say nothing matches. Den does not use one to pick skills.

A repeated task reuses a skill only on an exact ledger hit. Qwen3-Embedding-8B Q6_K is the memory embedder, not the skill picker.

## What den should not copy

- A trained picker, an RL router, or SkillOrchestra's handbook, skill discoverer, and agent-competence profiles.
- Cosine or k-nearest skill retrieval. It always returns a neighbor.
- Keyword auto-inject of a skill body, and auto-find that loads whole repo files into the packet.
- A second router stacked on description match.

## One plan line

Match one skill by description, load that one body, and reuse it on a later task only when the ledger hit is exact.

## Unproven

Den has not measured these. The paper reports up to 22.5% over RL orchestrators, and about 700× and 300× lower learning cost than Router-R1 and ToolOrchestra. The spec's startup metadata size (about 100 tokens) and body cap (under 5000 tokens recommended) are guidance, not a den count. OpenHands does not publish a false-inject rate for keyword auto-inject.
