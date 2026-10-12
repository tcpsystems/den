package main

import "testing"

func TestOpenRequiresWaveAndBranch(t *testing.T) {
	err := cmdOpen([]string{"-repo", t.TempDir()})
	if err == nil || err.Error() != "open needs -wave -branch -seats" {
		t.Fatalf("%v", err)
	}
}
