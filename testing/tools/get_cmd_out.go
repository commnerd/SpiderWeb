package tools

import (
    "os/exec"
    "fmt"
)

// GetCmdStdOut : Get the Stdout from command run
func GetCmdStdOut(cmd *exec.Cmd) string {
    return GetFunctionStdOut(func() {
        out, _ := cmd.Output()
        fmt.Print(string(out))
    })
}
