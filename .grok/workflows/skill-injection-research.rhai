let meta = #{
    name: "skill-injection-research",
    description: "Persistent research on bringing skill injection, models, and the lead agent online",
    when_to_use: "Append one sourced note on how den loads one skill and one model, then commit it",
    phases: [
        #{ title: "Read", detail: "three readers" },
        #{ title: "Verify", detail: "one skeptic" },
        #{ title: "Record", detail: "one writer" },
    ],
};

let note_schema = #{
    "type": "object",
    "required": ["claims"],
    "properties": #{
        "claims": #{
            "type": "array",
            "maxItems": 5,
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

let record_schema = #{
    "type": "object",
    "required": ["path", "subject"],
    "properties": #{
        "path": #{ "type": "string" },
        "subject": #{ "type": "string" },
    },
};

let repo = "/home/kivox/Projects/harness-plan";

phase("Read");
let jobs = [
    #{
        prompt: "Read " + repo + "/research/drafts/boundaries.md and " + repo + "/sections/skills if that directory exists. Use grep for injector and skill pointer. Claim only what a file says about how one skill body is chosen. Evidence is path:line or a quote under 20 words. Do not edit.",
        label: "read:injector",
        capability_mode: "read-only",
        output_schema: note_schema,
    },
    #{
        prompt: "Read " + repo + "/internal/catalog/catalog.go function Lead and Builtin. Run go test ./internal/catalog/ in that repo. Claim what the lead profile binds today: identity, role, model id, skill id. Evidence is the function or the test name. Do not edit.",
        label: "read:lead",
        capability_mode: "read-only",
        output_schema: note_schema,
    },
    #{
        prompt: "Read " + repo + "/skills/use-harness/references/omniroute.md and " + repo + "/plan/omniroute.md if present. Claim the model id the lead uses and whether a grant is required to call OmniRoute. Evidence is a file path. Do not edit. Do not call the network.",
        label: "read:model",
        capability_mode: "read-only",
        output_schema: note_schema,
    },
];
let notes = parallel(jobs);

phase("Verify");
let packed = json_encode(notes);
let verify_prompt = "Check these research claims against the files they name under " + repo + ". If the claims are empty, read internal/catalog/catalog.go Lead yourself and judge that. Set supported true only if you opened a file and the claim matches. Claims: " + packed;
let verdict = agent(verify_prompt, #{
    label: "verify",
    capability_mode: "read-only",
    output_schema: verdict_schema,
});
if verdict == () || !verdict.success || verdict.output.supported != true {
    complete(#{ summary: "claims not supported", ok: false });
}

phase("Record");
let record_prompt = "Read " + repo + "/skills/use-git/SKILL.md and follow it. Append one section to " + repo + "/research/drafts/skill-injection.md titled with today's work, not a new file. Include only claims the verifier accepted. End with a bring-up list of at most six commands or file reads that start the lead profile, load use-harness, and name auto/coding. Say what is not wired if denctl does not call the injector. Commit only that file with docs(research): and do not push. Return the path and the subject.";
let recorded = agent(record_prompt, #{
    label: "record",
    capability_mode: "execute",
    output_schema: record_schema,
});
if recorded == () || !recorded.success || recorded.output.path == () {
    complete(#{ summary: "record failed", ok: false });
}
complete(#{ summary: recorded.output.subject, path: recorded.output.path, ok: true });
