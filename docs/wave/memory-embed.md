# Memory embed

Written 2026-10-12. Stay on Qwen3-Embedding-8B Q6_K. Do not switch.

## Question

Is Qwen3-Embedding-8B Q6_K the right model for session memory and global knowledge, or should den switch?

## Sources

- Official GGUF card: https://huggingface.co/Qwen/Qwen3-Embedding-8B-GGUF
- Official file listing: https://huggingface.co/api/models/Qwen/Qwen3-Embedding-8B-GGUF/tree/main
- Paper page: https://huggingface.co/papers/2506.05176 (arXiv:2506.05176)
- Paper HTML: https://arxiv.org/html/2506.05176

## Finding

Keep the live file. The official tree lists `Qwen3-Embedding-8B-Q6_K.gguf` at 6,214,476,096 bytes. That is the live file size. The same listing is Q4_K_M 4,676,804,928 bytes (4.68 GB), Q5_K_M 5,422,342,464 bytes (5.42 GB), Q8_0 8,047,105,824 bytes (8.05 GB), and f16 15,141,156,384 bytes (15.1 GB). The card's June 2025 line is the unquantized 8B at 70.58 on MTEB Multilingual. The paper says the same 70.58, and that Qwen3-Embedding-8B is 8B, 36 layers, 32K context, 4096 dimensions, with MRL and instruction awareness. GGUF file_type 18 is Q6_K. Quantizing does not change the 4096 width.

Q8_0 is 8.05 GB before the KV cache. The card is 8 GB and the server context is 8192, so Q8_0 does not fit. Q5_K_M would fit and free about 0.8 GB. That is not a reason to switch: the live Q6_K already fits, it is the official file, and no source here measures a retrieval gain from dropping a quant on this machine. A smaller embedding model would shrink the index and the quality together. The paper's 70.58 is not a license to load a second GPU model.

Session memory and global knowledge are the same embedder. OpenViking is the long store. GitHub is the backup, not the read path. One skill is selected by description. The custom piece is the exact function-line ledger.

## What den should not copy

Do not copy a second vector store, a second embedder, or a quant leaderboard. Do not treat 70.58 as a score for Q6_K, Q5_K_M, or Q8_0. Do not load Q8_0 or f16 beside the live model. Do not raise server context toward the card's 32K or the GGUF metadata's 40960 on this 8 GB card. Do not read memory from GitHub. Do not add a skill router beyond description match.

## One plan line

Keep Qwen3-Embedding-8B Q6_K at 127.0.0.1:8091, context 8192, one vector per text, OpenViking as the long store, and the function-line ledger as the only custom index.

## Unproven

- Any MTEB or retrieval number for this Q6_K file. 70.58 is the unquantized 8B as of June 2025.
- That Q6_K beats Q5_K_M for den's session or global notes. No local measurement.
- That 4096 is the right width for the ledger versus a shorter MRL cut. The paper allows shorter vectors; den has not tested one.
- Fit of Q8_0 if context were cut below 8192. Not tested, and not the plan.
- Recall or latency of this file on the RX 6600. Unseen.
