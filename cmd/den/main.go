// den starts the controller and the TUI as one session.
//
//	den
//	den init
//
// The controller stays up while the TUI is open. Other parts stay off
// until the operator turns them on from the TUI.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/tcpsystems/den/internal/initd"
)

func main() {
	repo, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	home, _ := os.UserHomeDir()
	sock := filepath.Join(home, ".local", "share", "harness", "den.sock")
	board, err := initd.Load(repo, sock)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	board.Repo = repo
	board.Sock = sock
	if err := initd.Save(board); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ctl := exec.Command("denctl", "-repo", repo, "-sock", sock)
	ctl.Stdout = os.Stdout
	ctl.Stderr = os.Stderr
	if err := ctl.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "denctl is not on PATH")
		os.Exit(1)
	}
	defer func() { _ = ctl.Process.Kill() }()
	tui := exec.Command("dentui", "-sock", sock, "-repo", repo)
	tui.Stdin = os.Stdin
	tui.Stdout = os.Stdout
	tui.Stderr = os.Stderr
	if err := tui.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
