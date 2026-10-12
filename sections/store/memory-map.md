# Memory map

Fast is the packet on hand. Keep is long-term memory. A note moves between them. Nothing is deleted. A later note supersedes.

| Layer | Holds | Lives |
| --- | --- | --- |
| fast | The current section, the one skill body, the tool result just returned | The seat packet. Not embedded. |
| running | The task that is claimed and not finished | Mesh task row |
| useful | A result a later seat can reuse | OpenViking, tag `useful` |
| stale | Superseded but kept | OpenViking, tag `stale` |
| keep | Doctrine, contracts, and scored decisions | OpenViking, tag `keep` |
| skills | Name and description catalog. One body loads into fast | Skill files, not the packet |
| teaching | A worked example a seat should copy | OpenViking, tag `teaching` |
| suggestions | One-source ideas that are not decided | Plan suggestions and OpenViking tag `suggestions` |
| analysis | Source notes, including the OmniRoute docs condensation | OpenViking, tag `analysis` |

The OmniRoute docs condensation is `viking://resources/omniroute/omniroute.xml`. The changelog directory was left out. The file on disk is `/opt/omniroute.xml` on tcp.systems.

OpenViking is the keep store already running here. DeerMem is the settings editor and is not this store. A second vector database is not the store. The local embedder on port 8091 is the vector step for OpenViking.
