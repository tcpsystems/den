# Benchmarks

Den measures itself. A vendor score is not a score for this machine.

## Local tasks

A benchmark is a set of local tasks with a known pass or fail. The harness records the result, the tool call, and whether OmniRoute was named. That row is the score.

## Rejected as proof here

The 70.58 MTEB figure is the unquantized Qwen3-Embedding-8B. The 91.4 LoCoMo figure is Hindsight's own bench, and that project says not to trust LoCoMo. SWE-bench rates published for other SDKs were run on other models and other machines. None of those justify a change here.

## What the rows train

A pass keeps the description match. A fail drops it. An offline pass may use the 3B file on the GTX 1050 Ti while that card is idle. That pass is not in the live packet.
