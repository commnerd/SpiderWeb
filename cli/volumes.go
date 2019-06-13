package main

import (
    "fmt"
)

const VolumesHelp string = `Usage: %v volumes [ARGS]...

List volumes owned by you

Args:
  --help, -h Print this message
`

func (c cli) VolumesCommand() {
    fmt.Printf(VolumesHelp, cmdString)
}
