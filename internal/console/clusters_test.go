package console

import (
	"os"
	"path/filepath"
	"testing"
)

func TestClustersWithWorkersAndSummary(t *testing.T) {
	root := t.TempDir()
	taskDir := filepath.Join(root, "taskA")
	if err := os.Mkdir(taskDir, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"worker1.txt", "worker2.txt"} {
		if err := os.WriteFile(filepath.Join(taskDir, name), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	summary := "consensus note here\nsecond line\n"
	if err := os.WriteFile(filepath.Join(taskDir, "SUMMARY.md"), []byte(summary), 0o644); err != nil {
		t.Fatal(err)
	}

	clusters := Clusters(root)
	if len(clusters) != 1 {
		t.Fatalf("clusters %d", len(clusters))
	}
	c := clusters[0]
	if c.Task != "taskA" || c.Workers != 2 || !c.Consensus || c.Note != "consensus note here" {
		t.Fatalf("taskA %+v", c)
	}
}

func TestClustersEmptyTaskIsRunning(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "taskB"), 0o755); err != nil {
		t.Fatal(err)
	}

	clusters := Clusters(root)
	if len(clusters) != 1 {
		t.Fatalf("clusters %d", len(clusters))
	}
	c := clusters[0]
	if c.Task != "taskB" || c.Workers != 0 || c.Consensus || c.Note != "running" {
		t.Fatalf("taskB %+v", c)
	}
}

func TestClustersMissingRootIsEmpty(t *testing.T) {
	root := filepath.Join(t.TempDir(), "does-not-exist")

	clusters := Clusters(root)
	if len(clusters) != 0 {
		t.Fatalf("clusters %d, want 0", len(clusters))
	}
}
