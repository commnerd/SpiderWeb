package main

import (
    "../testing/tools"
    "testing"
    "os/exec"
    "fmt"
)

func TestLogin(t *testing.T) {
    cmd := exec.Command(execString, "login")
    expected := "Email: "
    out := tools.GetCmdStdOut(cmd)
    if out != expected {
        t.Fatalf("\nExpected: \"%v\",\nGot: \"%v\".", expected, out)
    }
}

func TestLoginHelp(t *testing.T) {
    cmd := exec.Command(execString, "login", "--help")
    loginHelp := fmt.Sprintf(LoginHelp, cmdString)
    out := tools.GetCmdStdOut(cmd)
    if out != loginHelp {
        t.Fatalf("\nExpected: \"%v\",\nGot: \"%v\".", loginHelp, out)
    }
}
