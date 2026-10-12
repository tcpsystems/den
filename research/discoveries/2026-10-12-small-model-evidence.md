# Small-model evidence and corrections

Status: partial research; no model installed, trained, or benchmarked locally in this workflow.

## Corrections to earlier session claims

- Absolute claims of under 300 MB RAM, millisecond inference, and sub-millisecond static analysis remain unmeasured. Parameter count does not establish resident memory or latency.
- A similarity threshold of 0.95 across three turns has no validated operating point here. Repetition alone cannot establish lack of progress.
- A fixed minimum of 100 training rows has no demonstrated justification. Dataset needs depend on coverage, class balance, leakage, and uncertainty.
- Deterministic execution does not establish complete security coverage or absence of false positives.
- The proposed hardware baseline is a test target, not a measured deployment requirement.

## Bounded next evaluation (proposal)

Separate skill selection, code-security signals, and progress monitoring into distinct datasets and evaluations. Start with existing rules and a cheap lexical baseline; compare a linear classifier or fastText, then an encoder only where a measured gap warrants it. Split by project/task family to reduce leakage. Report precision/recall, abstention, false interventions, p50/p95 latency, peak resident memory, and total operational cost. Measure progress using test outcomes, accepted artifacts, unresolved failures, and tool results. Run learned suggestions in shadow mode; retain explicit authorization checks independently. No live permission expansion or autonomous training is approved by this report.

## Provenance and limitations

The following is the workflow report verbatim, including its sources, rejected claims, and coverage gaps. Source verification is attributed to that workflow; this archival pass did not independently reopen the papers. Search authentication failures limited coverage. The report contains inconsistent NEXUS/Jaccard statements in its uncertainty and rejection sections; those statements must not be adopted without reopening the source.

---

# Research result

**Status: Partial**

Small CPU-oriented classifiers and deterministic checks are credible components for these harness tasks, but the evidence is strongest for text classification—not comprehensive security detection or similarity-based progress assessment. FastText has published multicore-CPU throughput results, and compact encoders offer measured size–accuracy trade-offs; neither establishes end-to-end performance on the specified machine. The packet contains no matched benchmark across logistic regression, FastText, small encoders, AST rules, and RegEx on a 4–8 vCPU, 8–16 GB host, so absolute latency, RAM requirements, and the benefit of its 8 GB GPU remain unverified.

### Security checks: useful signals, not verified coverage

Rule-based static analysis does not provide reliable vulnerability coverage merely by avoiding an LLM. On one real-world Node.js vulnerability dataset, the three best-performing JavaScript analyzers combined detected up to **57.6%** of known vulnerabilities at **0.11% precision**. That result is specific to the evaluated tools and dataset, not a universal accuracy estimate for AST or RegEx checks. [S1]

There is evidence for combining rules with a lightweight learned score: deterministic safety rules, argument inspection, and calibrated logistic regression achieved **0.949 F1** and **64.06% four-class intervention accuracy**, improving intervention selection by **27.3 percentage points** over rules alone. However, this was a **128-instance synthetic runtime-safety benchmark**, not source-code vulnerability detection or bad-practice detection. [S5]

The packet therefore does not establish whether logistic regression, FastText, or small encoders outperform AST/RegEx for code security. Its direct learned-model-versus-static-analysis comparison concerns **LLM prompting on partial code**, which cannot substantiate a non-LLM replacement. [S2]

### Skill routing: the strongest case for lightweight classification

Intent classification provides a relevant—but indirect—benchmark for mapping short descriptions to skills. The published comparison below favors BERT over FastText for in-scope accuracy, but supplies no corresponding logistic-regression or rule-based routing baseline. Actual skill-routing accuracy remains unmeasured here: SkillsBench contains **87 tasks across eight domains**, with curated skills and deterministic verifiers, but evaluates whether skills help agents rather than whether a classifier chooses the right skill. [S3]

| Method | Published accuracy or quality evidence | Published speed evidence | Size or memory evidence |
|---|---|---|---|
| Logistic regression + rules | Runtime-safety result above; no isolated logistic skill-routing result | — | — |
| FastText | **89.0%** full-setting CLINC150 in-scope accuracy, versus **96.9%** for BERT. [S4] | Over **1 billion words** trained in **under 10 minutes** on a “standard multicore CPU”; **500,000 sentences** classified among **312,000 classes** in **under one minute**. [S9] | — |
| FastText.zip | Slightly inferior accuracy to uncompressed FastText | — | Typically **two orders of magnitude less memory**, using product quantization. [S10] |
| DistilBERT | Retains **97%** of BERT’s reported language-understanding capability | **60% faster** than BERT | **40% smaller** than BERT. [S11] |
| TinyBERT4 | More than **96.8%** of BERT-base’s GLUE performance | **9.4× faster** inference | **7.5× smaller** than BERT-base. [S12] |
| MobileBERT | — | **5.5× faster** than BERT-base; **62 ms** inference on a Pixel 4 | **4.3× smaller** than BERT-base. [S13] |
| FlashText / keyword rules | No routing-accuracy result | **O(N)** versus **O(M×N)** for the paper’s dictionary-matching RegEx comparison | — [S14] |
| AST / general RegEx checks | No matched skill-routing result | — | — |

These results are not interchangeable: CLINC150 measures intent classification, GLUE measures broader language understanding, and compression ratios do not specify deployed RAM consumption. FastText’s aggregate throughput is not single-request latency; the phone measurement is not a server-CPU benchmark. Likewise, the keyword-matching complexity comparison is not a universal bound for RegEx engines or patterns.

### Loop detection: repetition is not demonstrated lack of progress

The packet provides **no accuracy, latency, or memory benchmark for output-similarity loop detection**, whether implemented with lexical rules, logistic regression, FastText, or encoder embeddings. It therefore does not establish a similarity threshold or demonstrate that repeated-looking output reliably identifies stalled work.

The available evidence instead highlights the importance of measured outcomes. In a progress-mirage testbed, an agent claimed improvement in **all 54 cycles**, although **56%** had an outcome delta at or below zero; trusting those claims reduced the best deployed state by **19%**. [S6] Even a judge given the full artifact, change diff, and verdict history accepted cycles of which **44% were regressions**, while rejecting **38% of real improvements**. [S7]

Production monitoring also has demonstrated value, but not specifically through similarity detection: across **more than 10,000 coding-agent trajectories**, about **30%** contained specification drift, reasoning problems, or tool-call failures. Single interventions resolved **90% of misbehaviors requiring one intervention**, not all misbehaviors. [S8]

### Assessment for the target machine

**FastText has the clearest published CPU evidence; compact encoders have promising relative efficiency evidence; logistic regression has a relevant hybrid-safety result.** None supplies a hardware-matched comparison for all three tasks. The supported conclusion is narrower than a universal winner: lightweight classification is credible for routing, rule-based security checks have demonstrated coverage and precision limitations, and output similarity remains an unvalidated proxy for progress in this evidence set.

## Sources
- [S1] "Study of JavaScript Static Analysis Tools for Vulnerability Detection in Node.js Packages (Brito et al., 2023)" — "https://arxiv.org/abs/2301.05097" (independently checked against "Study of JavaScript Static Analysis Tools for Vulnerability Detection in Node.js Packages" — "https://arxiv.org/abs/2301.05097")
- [S2] "Can LLM Prompting Serve as a Proxy for Static Analysis in Vulnerability Detection (Ceka et al., 2024)" — "https://arxiv.org/abs/2412.12039" (independently checked against "Can LLM Prompting Serve as a Proxy for Static Analysis in Vulnerability Detection" — "https://arxiv.org/abs/2412.12039")
- [S3] "SkillsBench: Benchmarking How Well Agent Skills Work" — "https://arxiv.org/abs/2602.12670" (independently checked against "SkillsBench: Benchmarking How Well Agent Skills Work Across Diverse Tasks" — "https://arxiv.org/pdf/2602.12670")
- [S4] "An Evaluation Dataset for Intent Classification and Out-of-Scope Prediction (CLINC150, HTML)" — "https://arxiv.org/html/1909.02027v1" (independently checked against "An Evaluation Dataset for Intent Classification and Out-of-Scope Prediction" — "https://arxiv.org/html/1909.02027v1")
- [S5] "NEXUS: Structured Runtime Safety for Tool-Using LLM Agents" — "https://arxiv.org/html/2607.19356v1"
- [S6] "When Do Agent Loops Mistake Stagnation for Progress?" — "https://arxiv.org/pdf/2607.25152v2" (independently checked against "When Do Agent Loops Mistake Stagnation for Progress? Self-Evaluation Bias and Externally Grounded Verification in Long-Running Autonomous LLM Agent Loops" — "https://arxiv.org/pdf/2607.25152v2")
- [S7] "When Do Agent Loops Mistake Stagnation for Progress?" — "https://arxiv.org/pdf/2607.25152v2" (independently checked against "When Do Agent Loops Mistake Stagnation for Progress? Self-Evaluation Bias and Externally Grounded Verification in Long-Running Autonomous LLM Agent Loops" — "https://arxiv.org/pdf/2607.25152v2, page 1, abstract; https://arxiv.org/abs/2607.25152")
- [S8] "Wink: Recovering from Misbehaviors in Coding Agents" — "https://arxiv.org/pdf/2602.17037v2"
- [S9] "Bag of Tricks for Efficient Text Classification (fastText), arXiv:1607.01759" — "https://arxiv.org/abs/1607.01759" (independently checked against "Bag of Tricks for Efficient Text Classification" — "https://arxiv.org/abs/1607.01759")
- [S10] "FastText.zip: Compressing text classification models, arXiv:1612.03651" — "https://arxiv.org/abs/1612.03651" (independently checked against "FastText.zip: Compressing text classification models" — "https://arxiv.org/abs/1612.03651, abstract; https://fasttext.cc/docs/en/references.html, reference 3")
- [S11] "DistilBERT, a distilled version of BERT: smaller, faster, cheaper and lighter, arXiv:1910.01108" — "https://arxiv.org/abs/1910.01108" (independently checked against "DistilBERT, a distilled version of BERT: smaller, faster, cheaper and lighter" — "https://arxiv.org/abs/1910.01108")
- [S12] "TinyBERT: Distilling BERT for Natural Language Understanding, arXiv:1909.10351" — "https://arxiv.org/abs/1909.10351" (independently checked against "TinyBERT: Distilling BERT for Natural Language Understanding" — "https://arxiv.org/html/1909.10351v5, abstract and Introduction; https://aclanthology.org/2020.findings-emnlp.372/")
- [S13] "MobileBERT: a Compact Task-Agnostic BERT for Resource-Limited Devices, arXiv:2004.02984" — "https://arxiv.org/abs/2004.02984" (independently checked against "MobileBERT: a Compact Task-Agnostic BERT for Resource-Limited Devices" — "https://aclanthology.org/2020.acl-main.195/")
- [S14] "Replace or Retrieve Keywords In Documents at Scale (FlashText), arXiv:1711.00046" — "https://arxiv.org/abs/1711.00046" (independently checked against "Replace or Retrieve Keywords In Documents at Scale" — "https://arxiv.org/abs/1711.00046, abstract; https://github.com/vi3k6i5/flashtext")

## Coverage and uncertainty
- "Question 1 uncertainty: No single benchmark was found that pits logistic regression + FastText + small encoder transformers against rule-based AST/RegEx on one unified 'common code security flaws and bad practices' task with accuracy, latency, and memory reported together."
- "Question 1 uncertainty: Memory (RAM) consumption is not reported in any of the located ML-vs-rule-based comparisons; Akter et al. report only execution time, and Semgrep reports only cost."
- "Question 1 uncertainty: Latency/throughput is not benchmarked head-to-head between rule-based SAST and ML classifiers anywhere found; Akter et al. measure ML execution time without a rule-based baseline."
- "Question 1 uncertainty: Logistic regression is not confirmed as a named, numerically-reported baseline in the vulnerability-detection studies located; it appears in code-smell ML studies (Fontana 2016 / Di Nucci 2018) but its exact measured accuracy there was not verified."
- "Question 1 uncertainty: The ML baselines in Semgrep's and Ceka et al.'s comparisons are LLMs, not 'small encoder transformers', so neither directly answers the small-transformer-vs-rules question."
- "Question 1 uncertainty: Semgrep's '3.5x more true positives' figure was verified only via a search-engine snippet quoting the page, not the page body itself (the SPA truncated the fetched content)."
- "Question 2 uncertainty: The exact trio (logistic + FastText + small encoder transformer vs. regex) could appear in a non-academic source such as a blog post, tutorial, or GitHub repo that my searches did not surface; I could not verify it exists."
- "Question 2 uncertainty: 'Small encoder transformer' is not a standard label in any of the located benchmarks; the closest analogs are CLINC150's BERT/CNN or SkillRouter's dense encoder baselines, neither of which is framed as the question describes."
- "Question 2 uncertainty: None of the located works report an 'error rate' in the classification sense for skill routing; SkillRouter uses retrieval metrics (Hit@1/R@10) and CLINC150 uses in-scope accuracy plus out-of-scope recall."
- "Question 2 uncertainty: FastText in CLINC150 refers to the classifier of Joulin et al. (a shallow n-gram averaging model), which may differ from how the question's 'FastText' baseline is intended."
- "Question 2 uncertainty: SkillRouter's baseline set (BM25 + dense encoders + rerankers) overlaps only partially with the question's 'small encoder transformers' and omits the logistic/rule-based components entirely."
- "Question 3 uncertainty: No inspected source reports a head-to-head measurement of logistic regression, FastText, and an encoder against rule-based or heuristic similarity thresholds specifically for detecting agent loops or lack of progress from output similarity."
- "Question 3 uncertainty: NEXUS’s logistic-versus-rules comparison is for multi-class safety intervention, not a benchmark of output-similarity loop detectors; its loop signal is one Jaccard rule rather than a learned similarity model."
- "Question 3 uncertainty: The progress-mirage and Wink measurements concern self-evaluation and production recovery, not a comparison of learned text-similarity models with fixed similarity thresholds."
- "Question 4 uncertainty: No inspected source reports wall-clock latency (ms) or runtime memory (RSS) for a rule-based classifier benchmarked head-to-head against logistic regression, FastText, and a small encoder transformer on the same corpus and hardware."
- "Question 4 uncertainty: No inspected source benchmarks the specific 4-8 vCPU / 8-16 GB / single-8GB-GPU profile: fastText only says 'standard multicore CPU', and the small-transformer latency figures come from phones (iPhone 7 Plus in DistilBERT, Pixel 4 in MobileBERT)."
- "Question 4 uncertainty: Logistic-classifier (e.g., LIBLINEAR/scikit-learn) latency and memory figures were not found in the inspected sources; the fastText paper cites linear classifiers only as accuracy baselines, not with timing/memory numbers."
- "Question 4 uncertainty: The single 8 GB GPU is effectively irrelevant to these CPU-bound models: no inspected source reports GPU inference latency or memory for logistic regression, FastText, or small encoders at this scale."
- "Question 4 uncertainty: Memory is reported almost entirely as parameter counts or relative ratios (40%, 7.5x, 4.3x, 'two orders of magnitude') rather than absolute RAM in GB during inference."
- "Question 4 uncertainty: The web_search tool returned HTTP 401 for all queries in this session, so these absence statements rest on the primary sources I fetched directly and are not an exhaustive literature sweep."
- "Claim claim-1 was excluded by verification: The paper reports accuracy and total execution time, with no memory-use metric or direct rule-based SAST benchmark in the reported experiments. However, the exact claim labels execution time as latency, which the evidence does not establish: the reported times are whole-model execution durations, not measured per-request inference latency.."
- "Claim claim-3 was excluded by verification: The DOI metadata confirms the paper's title, authors, SANER venue, and 2018 publication, but accessible evidence did not expose the abstract or results needed to verify the specific multi-smell finding. The IEEE page returned no usable content and attempted institutional sources were unavailable.."
- "Claim claim-4 was excluded by verification: The comparison is against a tuned guided LLM prompt, not tuned regex rules. This materially contradicts the claim.."
- "Claim claim-5 was excluded by verification: The cited file does not exist at the supplied path. The analysis directory contains skill-adoption.md, and repository searches did not locate the quoted statements. Therefore the claimed repository evidence could not be verified.."
- "Claim claim-6 was excluded by verification: The benchmark construction and baseline descriptions are supported, but the exact statement additionally calls SkillRouter the “main published” skill-routing benchmark. The accessible evidence does not establish that field-wide primacy.."
- "Claim claim-7 was excluded by verification: The numerical comparison is supported, but the exact metric list is not: the accessible v5 text lists Hit@1, MRR@10, Recall@K, and FC@10, without nDCG@10. The entire compound claim therefore fails exact verification.."
- "Claim claim-9 was excluded by verification: The dataset size and intent/OOS framing are supported, but the claim's classifier enumeration is incomplete. The paper explicitly evaluates a linear SVM and commercial platforms in addition to FastText, CNN, MLP, and BERT. The exact statement cannot be accepted as the paper's model list.."
- "Claim claim-11 was excluded by verification: The cited SkillRouter paper cannot establish a literature-wide negative or verify the asserted history of searches. Independent web search attempts failed with an authentication error. No accessible evidence directly substantiated the exact claim that no matching published benchmark or result was found.."
- "Claim claim-13 was excluded by verification: The accessible cited paper does not contain the claimed Jaccard loop-detection rule or the asserted Table 3 detector row. Its Table 3 instead summarizes evaluation across synthetic, R-Judge, and indirect-prompt-injection settings.."
- "Claim claim-17 was excluded by verification: Opening the exact supplied query produced a DuckDuckGo bot challenge, not a results page stating that no results were found. The claimed search outcome is consequently unverified.."
