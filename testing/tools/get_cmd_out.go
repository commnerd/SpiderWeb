package tools

import (
    "os/exec"
    "fmt"
)

func GetCmdStdOut(cmd *exec.Cmd) string {
    return GetFunctionStdOut(func() {
        out, _ := cmd.Output()
        fmt.Print(string(out))
    })
}
