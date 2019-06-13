package main

import (
    "../testing/tools"
    "testing"
    "os/exec"
    "os"
)

func ServersTest(t *testing.T) {
    cmd := exec.Command(os.Args[0], "servers")
    out := tools.GetCmdStdOut(cmd)
    if out != ServersHelp {
        t.Fatalf("Expected servers help message, Got: \"%v\".", out)
    }
}
