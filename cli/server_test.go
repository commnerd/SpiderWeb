package main

import (
    "github.com/commnerd/SpiderWeb/testing/tools"
    "testing"
    "os/exec"
)

func TestServer(t *testing.T) {
    cmd := exec.Command(TestCmd, "server")
    out := tools.GetCmdStdOut(cmd)
    if out == "Servers" {
        return
    }
    t.Fatalf("Expected \"Servers\", Got: \"" + out + "\".")
}
