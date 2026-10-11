# den

Local seats for the tcp.systems agentic system. Each seat gets one packet, one skill, and a microVM. They still talk on the mesh.

This is not the Foxcore control plane. `fox`, `foxd`, and `floki` stay in `foxcore-cloud`. The deprecated base is [`foxcore-cloud/agent-mesh`](https://github.com/foxcore-cloud/agent-mesh) at `c985423`. That line is archived. New work lands on `develop`.

`main` on this repo is the archive pointer. Do not add features to it.

Trackers, sources, pricing, and models live in [`tcpsystems/radar`](https://github.com/tcpsystems/radar), moved from `foxcore-cloud/tcp-systems`.

## Read first

`plan/PLAN.md` is the plan. `DEVELOP.md` is how a section is committed and pushed.

## What a seat already has

- A direct model id. OmniRoute is a tool, not the default hop.
- Fast and recent memory on the machine. The OmniRoute host keeps the backup.
- Embeddings at `127.0.0.1:8091` when the RX 6600 embedder is installed.
- A smolvm guest on the `harness` network.
- One route skill, `skills/use-harness`.
