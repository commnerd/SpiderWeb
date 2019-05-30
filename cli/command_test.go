package main

import (
    tools "github.com/commnerd/SpiderWeb/testing/tools"
    "testing"
    "os/exec"
    "os"
)

func TestEmptyHelpMessage(t *testing.T) {
    got := tools.GetStdOut(main)
    if got != GetHelp() {
        t.Errorf("Got: %s", got)
    }
}

func TestHelper(t *testing.T) {
    cmd := exec.Command(os.Args[0], "help")
    err := cmd.Run()
    if e, ok := err.(*exec.ExitError); ok && !e.Success() {
        return
    }
    t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestCrasher(t *testing.T) {
    cmd := exec.Command(os.Args[0], "blah")
    err := cmd.Run()
    if e, ok := err.(*exec.ExitError); ok && !e.Success() {
        return
    }
    t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestEmptyCommand(t *testing.T) {
    cmd := exec.Command(os.Args[0], "")
    err := cmd.Run()
    if e, ok := err.(*exec.ExitError); ok && !e.Success() {
        return
    }
    t.Fatalf("process ran with err %v, want exit status 1", err)
}
