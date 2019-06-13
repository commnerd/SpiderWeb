package main

import (
    "../testing/tools"
    "testing"
    "os/exec"
    "fmt"
)

func TestVolumes(t *testing.T) {
    cmd := exec.Command(cmdString, "volumes")
    volumesHelp := fmt.Sprintf(VolumesHelp, cmdString)
    out := tools.GetCmdStdOut(cmd)
    if out != volumesHelp {
        t.Fatalf("\nExpected: \"%v\",\nGot: \"%v\".", volumesHelp, out)
    }
}
