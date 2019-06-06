package main

import (
    "../testing/tools"
    "testing"
    "os/exec"
)

func ServersTest(t *testing.T) {
    cmd := exec.Command(TestCmd, "servers")
    out := tools.GetCmdStdOut(cmd)
    if out != ServersHelp {
        t.Fatalf("Expected servers help message, Got: \"" + out + "\".")
    }
}
