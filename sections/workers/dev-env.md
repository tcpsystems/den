# Worker dev environment

The seat's model call is a direct model id. OmniRoute is a tool, not that default. Its combos, batch API, A2A card, and MCP are used when a seat names them. A combo adds a hop, and that hop is latency.

## Embeddings

The harness embeds through the local server on port 8091. On this machine that is Qwen3-Embedding-8B Q6_K on the RX 6600, Vulkan, context 8192. The card is full. Do not load a second GPU model beside it. A fresh machine without that card leaves the embedder unset until one is installed. It does not fall through to a cloud embedder.

## Isolation

Each worker's dev environment is a smolvm microVM: its own kernel, via KVM on Linux. The VM has a network so it can reach the mesh and the other workers. The filesystem, packages, and processes stay inside the guest. The society is the mesh. The isolation is the VM.

smolvm is https://github.com/smol-machines/smolvm. Install on Linux is the upstream script into `~/.smolvm`. A machine is `smolvm machine create --net --name <seat>`. Networking is off unless `--net` is set. Workers need `--net`.

Vendor SDKs are not this boundary. Vercel Sandbox, Cloudflare Sandbox, Anthropic, DeepSeek, DeerFlow, Qwen, Alibaba, Kimi, MiniMax, and Mistral are clients or hosted sandboxes. Use one of them only when it replaces the microVM with a stronger dev environment, not as a second bus and not as the model picker.

## Fresh Debian

`deploy/debian-fresh.sh` is the bootstrap. It expects a new Debian with sudo, installs KVM, installs smolvm for the user, and leaves the mesh and the embedder as later steps. It does not join this Fedora host's disk or the 1 TB volume.
