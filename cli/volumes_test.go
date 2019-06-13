package main

import (
    "../testing/tools"
    "testing"
    "os/exec"
)

func VolumesTest(t *testing.T) {
    cmd := exec.Command(TestCmd, "volumes")
    out := tools.GetCmdStdOut(cmd)
    if out != VolumesHelp {
        t.Fatalf("Expected servers help message, Got: \"%v\".", out)
    }
}
