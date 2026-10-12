# A2A and planning

## Question

What should den take from Google A2A (Agent2Agent, April 2025) and from multi-agent workflow engines for planning? Settled: a child gets its own seven-field packet and does not inherit the parent transcript. Spring AI `shareState` default is refused. A pause ends the packet.

## Sources

- Google, "Announcing the Agent2Agent Protocol (A2A)", 9 April 2025. https://developers.googleblog.com/en/a2a-a-new-era-of-agent-interoperability/ (Wayback capture of the same day: https://web.archive.org/web/20250410002008/https://developers.googleblog.com/en/a2a-a-new-era-of-agent-interoperability/)
- Habler, Huang, Narajala, Kulkarni, "Building A Secure Agentic AI Application Leveraging Google's A2A Protocol", arXiv:2504.16902, submitted 22 April 2025. https://arxiv.org/abs/2504.16902
- LangGraph subgraphs and interrupts, read 12 October 2026. https://docs.langchain.com/oss/python/langgraph/use-subgraphs and https://docs.langchain.com/oss/python/langgraph/interrupts
- Spring AI Alibaba `A2aRemoteAgent` on `main`, read 12 October 2026. https://github.com/alibaba/spring-ai-alibaba/blob/main/spring-ai-alibaba-agent-framework/src/main/java/com/alibaba/cloud/ai/graph/agent/a2a/A2aRemoteAgent.java

## Finding

A2A is a client agent handing a remote agent a task. The remote agent is opaque: no shared memory, tools, or context by default. The task is the unit of work and has a lifecycle. It may finish at once, or stay open while both sides sync status. The April announcement names the capabilities discovery, task lifecycle, collaboration, and modality negotiation. It does not name the state strings. The paper, citing the early spec, lists submitted, working, input-required, completed, failed, and canceled. Input-required is the long task waiting on the user. Completed, failed, and canceled are the ends.

That matches a child packet. The parent is the client. The child is the remote seat. The seven fields are the task, not a copy of the parent transcript. Artifacts come back as the child's Write. The parent does not read the child's tools.

Workflow engines do the other thing. LangGraph, when a subgraph is added as a node, shares state keys, so the child reads and writes the parent's channels. A wrapper node is required to keep the schemas apart. `interrupt()` pauses the same graph: the checkpointer saves the state and the same `thread_id` resumes it. Spring AI Alibaba's current `A2aRemoteAgent` field is `private boolean shareState = true`, so a nested remote agent shares state unless the caller turns it off. That default is the one den refuses.

## What den should not copy

Do not open an A2A task between den seats. Seats already talk by mesh messages and child packets. Do not copy Agent Cards, `/.well-known/agent.json`, JSON-RPC, or SSE. Do not share parent channels the way a LangGraph subgraph node does. Do not resume a paused graph. A pause ends the packet. The next work is a new seven-field packet. Do not take `shareState = true`.

## One plan line

A child packet is an opaque task: its own seven fields, no parent transcript, and a pause is a terminal state, not a checkpoint.

## Unproven

Partner counts, the "later this year" production date, and any throughput or latency figure were not treated as evidence. The current A2A state enum was not opened, so the six names above are the paper's, not a claim about today's spec. No test of Spring AI or LangGraph was run here.
