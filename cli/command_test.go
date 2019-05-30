package main

import (
    tools "github.com/commnerd/SpiderWeb/testing/tools"
    "testing"
    "os"
)

func TestEmptyHelpMessage(t *testing.T) {
    got := tools.GetStdOut(main)
    if got != Help() {
        t.Errorf("Got: %s", got)
    }
}

func TestBadCommandHelpMessage(t *testing.T) {
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()

    os.Args = []string{"sw", "blah"}
    got := tools.GetStdOut(main)
    if got != Help() {
        t.Errorf("Got: %s", got)
    }
}

func TestHelloCommand(t *testing.T) {
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()

    os.Args = []string{"sw", "hello"}
    got := tools.GetStdOut(main)
    if got != Hello() {
        t.Errorf("Got: %s", got)
    }
}
