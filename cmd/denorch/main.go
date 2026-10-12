package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tcpsystems/den/internal/orch"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	var err error
	switch os.Args[1] {
	case "open":
		err = cmdOpen(os.Args[2:])
	case "commit":
		err = cmdCommit(os.Args[2:])
	case "take":
		err = cmdTake(os.Args[2:])
	case "show":
		err = cmdShow(os.Args[2:])
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "denorch open|commit|take|show")
}

func cmdOpen(args []string) error {
	fs := flag.NewFlagSet("open", flag.ContinueOnError)
	repo := fs.String("repo", ".", "repository")
	id := fs.String("wave", "", "wave id")
	branch := fs.String("branch", "", "shared branch")
	base := fs.String("base", "HEAD", "base revision")
	work := fs.String("worktrees", "", "worktree parent")
	seats := fs.String("seats", "", "id:role:path[,path] separated by semicolon")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *id == "" || *branch == "" || *seats == "" {
		return fmt.Errorf("open needs -wave -branch -seats")
	}
	w := orch.Wave{ID: *id, Branch: *branch, Base: *base, Repo: *repo, Worktrees: *work}
	for _, raw := range strings.Split(*seats, ";") {
		parts := strings.Split(raw, ":")
		if len(parts) != 3 {
			return fmt.Errorf("seat %q wants id:role:path", raw)
		}
		w.Seats = append(w.Seats, orch.Seat{
			ID: parts[0], Role: parts[1], Owns: strings.Split(parts[2], ","), Status: "new",
		})
	}
	out, err := orch.Open(w)
	if err != nil {
		return err
	}
	return printWave(out)
}

func cmdCommit(args []string) error {
	fs := flag.NewFlagSet("commit", flag.ContinueOnError)
	repo := fs.String("repo", ".", "repository")
	id := fs.String("wave", "", "wave id")
	seat := fs.String("seat", "", "seat id")
	msg := fs.String("m", "", "conventional commit subject")
	if err := fs.Parse(args); err != nil {
		return err
	}
	out, err := orch.Commit(*repo, *id, *seat, *msg)
	if err != nil {
		return err
	}
	return printWave(out)
}

func cmdTake(args []string) error {
	fs := flag.NewFlagSet("take", flag.ContinueOnError)
	repo := fs.String("repo", ".", "repository")
	id := fs.String("wave", "", "wave id")
	seat := fs.String("seat", "", "seat id")
	if err := fs.Parse(args); err != nil {
		return err
	}
	out, err := orch.Take(*repo, *id, *seat, orch.DefaultChecks())
	if err != nil {
		return err
	}
	return printWave(out)
}

func cmdShow(args []string) error {
	fs := flag.NewFlagSet("show", flag.ContinueOnError)
	repo := fs.String("repo", ".", "repository")
	id := fs.String("wave", "", "wave id")
	if err := fs.Parse(args); err != nil {
		return err
	}
	w, err := orch.Load(*repo, *id)
	if err != nil {
		return err
	}
	return printWave(w)
}

func printWave(w orch.Wave) error {
	b, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
