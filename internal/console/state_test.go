package console

import "testing"

func TestSlotsAreOneList(t *testing.T) {
	slots := Slots()
	if len(slots) < 10 {
		t.Fatalf("slots %d", len(slots))
	}
	seen := map[string]bool{}
	for _, slot := range slots {
		if seen[slot.ID] {
			t.Fatalf("duplicate %s", slot.ID)
		}
		seen[slot.ID] = true
		if slot.Status != "this" && slot.Status != "later" {
			t.Fatalf("%s status %s", slot.ID, slot.Status)
		}
	}
	if !seen["go"] || !seen["svelte"] || !seen["sse"] || !seen["dry"] {
		t.Fatal("missing this-build slots")
	}
	for _, id := range []string{"react", "shadcn", "cloudflare", "hetzner", "scaleblade"} {
		if !seen[id] {
			t.Fatalf("missing later slot %s", id)
		}
	}
}

func TestAtoi(t *testing.T) {
	if atoi("12") != 12 || atoi("x") != 0 {
		t.Fatal("atoi")
	}
}
