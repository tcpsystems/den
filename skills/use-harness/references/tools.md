# Tools

Read this only when the task calls a named tool.

Call the tool whose description matches the step. One call, then the result.

| Tool | Use |
| --- | --- |
| `meshctl suite check` | Which local repos and MCP tools this machine can see |
| OmniRoute MCP or `https://ai.tcp.systems/v1` | A named model, combo, batch, or A2A call. Not the default seat model |
| Hetzner CLI `hcloud --context foxcore` | Lab and staging hosts. Keep the disk unless the human says otherwise |
| Cloudflare MCP `https://mcp.cloudflare.com/mcp` | Account work after OAuth. Docs MCP is separate |
| Local embedder `:8091` | Vectors |

Do not put a secret in the packet. Do not let a tool pick the next model.
