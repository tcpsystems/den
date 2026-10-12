let meta = #{
    name: "harness-advance",
    description: "One harness slice: read the gap, commit it with use-git, then check the commit",
    when_to_use: "Advance /home/kivox/Projects/harness-plan by one verified local commit",
    phases: [
        #{ title: "Survey", detail: "one reader" },
        #{ title: "Slice", detail: "one writer" },
        #{ title: "Check", detail: "one reviewer" },
    ],
};

let gap_schema = #{
    "type": "object",
    "required": ["gap", "paths", "subject"],
    "properties": #{
        "gap": #{ "type": "string" },
        "paths": #{ "type": "array", "maxItems": 4, "items": #{ "type": "string" } },
        "subject": #{ "type": "string" },
    },
};

let slice_schema = #{
    "type": "object",
    "required": ["committed", "subject", "test"],
    "properties": #{
        "committed": #{ "type": "boolean" },
        "subject": #{ "type": "string" },
        "test": #{ "type": "string" },
    },
};

let check_schema = #{
    "type": "object",
    "required": ["ok", "reason"],
    "properties": #{
        "ok": #{ "type": "boolean" },
        "reason": #{ "type": "string" },
    },
};

let repo = "/home/kivox/Projects/harness-plan";

phase("Survey");
let survey_prompt = "Read " + repo + "/skills/use-git/SKILL.md, " + repo + "/plan/READY.md, and " + repo + "/plan/protocols.md. Run git -C " + repo + " status -sb and git -C " + repo + " log --oneline -8. Name one gap that blocks using the harness from inside itself. Return paths that exist or that the writer should create under plan/ or skills/, and a conventional commit subject. Do not edit.";
let survey = agent(survey_prompt, #{
    label: "survey",
    capability_mode: "read-only",
    output_schema: gap_schema,
});
let gap = "bring the lead profile online";
let paths_line = "skills/use-git/SKILL.md";
if survey != () && survey.success && survey.output != () && survey.output.gap != () && survey.output.gap != "" {
    gap = survey.output.gap;
    if survey.output.subject != () && survey.output.subject != "" {
        paths_line = survey.output.subject;
    }
}

phase("Slice");
let slice_prompt = "Repo " + repo + " branch wave/20261012-skill-adoption. Read skills/use-git/SKILL.md and follow it. Close this one gap: " + gap + ". Suggested subject: " + paths_line + ". Touch at most four paths under plan/ or skills/. Run go test ./... in the repo. If it fails, do not commit; set committed false and put the first failure in test. If it passes, commit only those paths. Do not push. Do not commit secrets.";
let slice = agent(slice_prompt, #{
    label: "slice",
    capability_mode: "execute",
    output_schema: slice_schema,
});
if slice == () || !slice.success || slice.output.committed != true {
    complete(#{ summary: "slice did not commit", ok: false });
}

phase("Check");
let check_prompt = "In " + repo + " run git log -1 --oneline and git status --porcelain. Set ok true only if HEAD subject matches a conventional commit and the committed paths are clean. Reason must quote the log line. Do not edit.";
let check = agent(check_prompt, #{
    label: "check",
    capability_mode: "read-only",
    output_schema: check_schema,
});
if check == () || !check.success || check.output.ok != true {
    complete(#{ summary: "check failed", ok: false });
}
complete(#{ summary: slice.output.subject, ok: true });
