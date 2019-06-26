package main

import (
    "../testing/tools"
    "testing"
    "os/exec"
    "fmt"
)

func TestServers(t *testing.T) {
    cmd := exec.Command(execString, "servers")
    serversHelp := fmt.Sprintf(ServersHelp, cmdString)
    out := tools.GetCmdStdOut(cmd)
    if out != serversHelp {
        t.Fatalf("\nExpected: \"%v\",\nGot: \"%v\".", serversHelp, out)    }
    }
