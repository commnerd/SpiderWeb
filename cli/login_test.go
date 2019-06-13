package main

import (
    "../testing/tools"
    "testing"
    "os/exec"
    "fmt"
)

func TestLogin(t *testing.T) {
    cmd := exec.Command(cmdString, "login")
    loginHelp := fmt.Sprintf(LoginHelp, cmdString)
    out := tools.GetCmdStdOut(cmd)
    if out != loginHelp {
        t.Fatalf("\nExpected: \"%v\",\nGot: \"%v\".", loginHelp, out)
    }
}

func TestLoginHelp(t *testing.T) {
    cmd := exec.Command(cmdString, "login")
    loginHelp := fmt.Sprintf(LoginHelp, cmdString)
    out := tools.GetCmdStdOut(cmd)
    if out != loginHelp {
        t.Fatalf("\nExpected: \"%v\",\nGot: \"%v\".", loginHelp, out)
    }
}
