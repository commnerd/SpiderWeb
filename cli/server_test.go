package main

import (
    "../testing/tools"
    "testing"
    "os/exec"
)

func ServerTest(t *testing.T) {
    cmd := exec.Command(TestCmd, "server")
    out := tools.GetCmdStdOut(cmd)
    if out != ServerHelp {
        t.Fatalf("Expected help message, Got: \"" + out + "\".")
    }
}
