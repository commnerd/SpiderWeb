package tools

import (
    "bytes"
    "io"
    "os"
)

// SetFunctionStdIn : Pass Stdout to Stdin
func SetFunctionStdIn(f func()) string {
    old := os.Stdin // keep backup of the real stdin
    r, w, _ := os.Pipe()
    os.Stdin = w

    f()

    inC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        inC <- buf.String()
    }()

    // back to normal state
    w.Close()
    os.Stdin = old // restoring the real stdin
    in := <-inC

    return in
}
