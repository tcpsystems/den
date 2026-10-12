package catalog

import "testing"

func TestBuiltinKinds(t *testing.T) {
	idx := Builtin()
	if len(idx.ByKind(KindRole)) != 4 {
		t.Fatalf("roles: %d", len(idx.ByKind(KindRole)))
	}
	if len(idx.ByKind(KindTool)) < 3 {
		t.Fatal("tools")
	}
}

func TestBindRequiresCatalog(t *testing.T) {
	idx := Builtin()
	idx.Add(Entry{Kind: KindIdentity, ID: "den-1", Title: "Den", Summary: "Local operator seat."})
	idx.Add(Entry{Kind: KindSkill, ID: "use-harness", Title: "Use harness", Summary: "Route one reference.", Path: "skills/use-harness/SKILL.md"})
	p := Profile{
		ID: "lead", Identity: "den-1", Role: "builder", Niche: "harness-runtime",
		Model: "auto/coding", Skills: []string{"use-harness"}, MCP: []string{"omniroute"},
		Memory: "session-ledger",
	}
	if err := Bind(idx, p); err != nil {
		t.Fatal(err)
	}
	p.Skills = []string{"missing"}
	if err := Bind(idx, p); err == nil {
		t.Fatal("want missing skill")
	}
}

func TestAddRejectsUnknownKind(t *testing.T) {
	idx := Builtin()
	if err := idx.Add(Entry{Kind: "agent", ID: "x", Title: "X", Summary: "no"}); err == nil {
		t.Fatal("want unknown kind")
	}
}
