package main

import (
    "../testing/tools"
    "testing"
    "os/exec"
    "fmt"
)

func TestServer(t *testing.T) {
    cmd := exec.Command(cmdString, "server")
    serverHelp := fmt.Sprintf(ServerHelp, cmdString)
    out := tools.GetCmdStdOut(cmd)
    if out != serverHelp {
        t.Fatalf("\nExpected: \"%v\",\nGot: \"%v\".", serverHelp, out)
    }
}
