let meta = #{
    name: "seat-holdup",
    description: "Explain why seats stalled, then check embeddings, LangGraph structures, and the 32k context suggestion",
    phases: [
        #{ title: "Research", detail: "four readers" },
        #{ title: "Verify", detail: "one skeptic" },
        #{ title: "Report" },
    ],
};

let note_schema = #{
    "type": "object",
    "required": ["claims"],
    "properties": #{
        "claims": #{
            "type": "array",
            "maxItems": 6,
            "items": #{
                "type": "object",
                "required": ["claim", "evidence"],
                "properties": #{
                    "claim": #{ "type": "string" },
                    "evidence": #{ "type": "string" },
                },
            },
        },
    },
};

let verdict_schema = #{
    "type": "object",
    "required": ["supported", "reason"],
    "properties": #{
        "supported": #{ "type": "boolean" },
        "reason": #{ "type": "string" },
        "evidence": #{ "type": "string" },
    },
};

let repo = "/home/kivox/Projects/harness-plan";
let topics = [
    "Seat holdup. Read plan/BASE.md and run git commands in " + repo + ". Count r* branches not in develop, name section branches with Go tests, and list why seats stalled (401, empty replies, unpushed drafts). Use the shell and read_file. Do not invent a count.",
    "Embeddings. Read plan/nest.md and plan/structures.md. Confirm where Qwen3-Embedding-8B listens locally and how tcp reaches it on Headscale. Say what must stay loopback.",
    "LangChain and LangGraph structures. Read plan/structures.md. Open the public READMEs for langchain-ai/langgraph and langchain-ai/deepagents. State which structures den may use and which must not become the bus.",
    "Context optimisation. The user suggests production seats stay sharper under about 32000 input tokens, including identity, skills, and rules, and that focus slips in a long chat. This is a suggestion, not a cap. Find whether a primary source supports or limits that hint. If you cannot open a source, say so.",
];

phase("Research");
let jobs = [];
let ti = 0;
for topic in topics {
    let prompt = "";
    prompt += "Read the den repo at " + repo + " before you answer. Use read_file or the shell. ";
    prompt += topic;
    prompt += " Return at most 6 claims. Each claim is one sentence. evidence is a file path or a URL you opened. An empty list is valid only after you looked.";
    jobs.push(#{
        prompt: prompt,
        label: "read-" + ti.to_string(),
        capability_mode: "read-only",
        output_schema: note_schema,
        phase: "Research",
    });
    ti += 1;
}
let results = parallel(jobs);

let claims = [];
for r in results {
    if r != () && r.success && r.output != () && r.output.claims != () {
        for c in r.output.claims {
            if c != () && c.claim != () && c.evidence != () && claims.len() < 12 {
                claims.push(c);
            }
        }
    }
}
log("claims " + claims.len().to_string());
if claims.len() == 0 {
    complete(#{ summary: "No sourced claims.", confirmed: [] });
}

phase("Verify");
let packed = json_encode(claims);
let vprompt = "";
vprompt += "Check each claim against the den repo or the cited URL. Set supported=true only when you opened the evidence. Otherwise supported=false. ";
vprompt += "The 32000-token hint is a suggestion, not a hard cap. Reject a claim that states it as a universal law without a source. ";
vprompt += packed;
let verdict = agent(vprompt, #{
    label: "skeptic",
    capability_mode: "read-only",
    output_schema: verdict_schema,
    phase: "Verify",
});

phase("Report");
let kept = [];
if verdict != () && verdict.success && verdict.output != () && verdict.output.supported == true {
    kept = claims;
}
let report = "# Seat holdup\n\n";
if kept.len() == 0 {
    report += "The skeptic did not confirm the bundle. Treat every claim below as unverified.\n\n";
    for c in claims {
        report += "- " + c.claim + " (" + c.evidence + ")\n";
    }
} else {
    report += "Confirmed as a bundle. The 32000-token figure stays a suggestion.\n\n";
    for c in kept {
        report += "- " + c.claim + " (" + c.evidence + ")\n";
    }
}
report += "\nOptional orchestrator parameters stay optional. Safety limits are not optional.\n";
let path = write_scratch_file("seat-holdup.md", report);
complete(#{ path: path, report: report, confirmed: kept.len() });
