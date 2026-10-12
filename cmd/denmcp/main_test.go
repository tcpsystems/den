package main

import "testing"

func TestSpawnRequiresGrantFlags(t *testing.T) {
	err := cmdSpawn([]string{"-ids", "exa"})
	if err == nil || err.Error() != "spawn needs -ids -grant -parent-key -seat" {
		t.Fatalf("%v", err)
	}
}
