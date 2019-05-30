package main

import (
    "testing"
    "os/exec"
    "os"
)

func TestServer(t *testing.T) {
    cmd := exec.Command(os.Args[0], "server")
    err := cmd.Run()
    if e, ok := err.(*exec.ExitError); ok && !e.Success() {
        return
    }
    t.Fatalf("process ran with err %v, want exit status 1", err)
}
