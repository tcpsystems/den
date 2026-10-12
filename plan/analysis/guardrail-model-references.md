# Source-backed guardrail and routing model references

These are external, checkable starting points for the small action-assistance models. Each URL must be verified before it becomes a dependency. Claims here are search-snippet level; they need a page read before adoption.

## Security / bad-practice detectors

| Model / system | URL | What it claims |
| --- | --- | --- |
| OpenGuardrails-Text-2510 | https://huggingface.co/openguardrails/OpenGuardrails-Text-2510 | Configurable safety policy, low-latency deployment, LLM-based guardrail. |
| CodeBERT-vulnerability-detection | https://github.com/GeorgeVe/CodeBERT-vulnerability-detection | CodeBERT pretrained on 6 languages, fine-tuned for vulnerability detection. |
| codebert-base-finetuned-detect-insecure-code | https://huggingface.co/mrm8488/codebert-base-finetuned-detect-insecure-code | CodeBERT fine-tuned to detect insecure code. |
| cisco-ai-defense/securebert2 | https://github.com/cisco-ai-defense/securebert2 | Claims balanced detection with higher F1 and fewer false positives. |
| Neural-Vulnerability-Scanner | https://github.com/AmitAK1/Neural-Vulnerability-Scanner | CodeBERT 125M + LoRA on BigVul for C/C++ vulnerability detection. |
| Vulnerability Detection with Code Language Models | https://arxiv.org/html/2403.18624v2 | Survey of how far code LMs are from human-verified vulnerability fixes. |

## Skill / niche routing

| System | URL | What it claims |
| --- | --- | --- |
| aurelio-labs/semantic-router | https://github.com/aurelio-labs/semantic-router | Semantic vector-space routing instead of slow LLM tool decisions. |
| sentence-transformers | https://huggingface.co/sentence-transformers | Embedding library that underpins description-matching routers. |

## Loop / no-progress watchdogs

| System | URL | What it claims |
| --- | --- | --- |
| xz-dev/pi-reflect-watchdog | https://github.com/xz-dev/pi-reflect-watchdog | Root-loop and active-time thresholds trigger reflection and steering. |
| agent-loop-watchdog.md | https://github.com/dhrvrc/graph-architecture/blob/main/agents/agent-loop-watchdog.md | Heartbeat liveness, progress-rate analysis, stall detection. |
| antigravity progress watchdog | https://antigravitylab.net/en/articles/agents/antigravity-background-agent-stuck-watchdog-progress-heartbeat-design | Progress heartbeat to detect alive-but-stuck agents. |

## What to check before adopting any of these

1. Read the model card or README. Record the parameter count, the license, and the exact input/output format.
2. Confirm the model runs on the stated CPU baseline. A 125M CodeBERT checkpoint is plausible; a larger one is not.
3. Run one local sample through it and record pass or fail against the ground truth you already trust.
4. Check the license permits the deployment form (local, no public API).
5. Prefer a model that emits a score or label over one that emits chat. These are action-assistance models, not chat generators.
