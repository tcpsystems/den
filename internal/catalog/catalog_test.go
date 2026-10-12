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
	if len(idx.ByKind(KindIdentity)) != 1 || len(idx.ByKind(KindProfile)) != 1 {
		t.Fatal("self-use agent missing")
	}
}

func TestLeadBindsWithoutRequiredMCP(t *testing.T) {
	if err := Bind(Builtin(), Lead()); err != nil {
		t.Fatal(err)
	}
}

func TestBindRequiresCatalog(t *testing.T) {
	idx := Builtin()
	p := Lead()
	p.MCP = []string{"omniroute"}
	if err := Bind(idx, p); err != nil {
		t.Fatal(err)
	}
	p.Identity = "missing"
	if err := Bind(idx, p); err == nil {
		t.Fatal("want missing identity")
	}
	p = Lead()
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
