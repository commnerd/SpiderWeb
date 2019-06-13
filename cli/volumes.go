package main

import "fmt"

const VolumesHelp string = `Usage: %v servers [ARGS]...

List servers owned by you

Args:
  --help, -h Print this message
`

func (c cli) VolumesCommand() {
    fmt.Printf("Volumes")
}
