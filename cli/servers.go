package main

import "fmt"

const ServersHelp string = `Usage: %v servers [ARGS]...

List servers owned by you

Args:
  --help, -h Print this message
`

func (c cli) ServersCommand() {
    fmt.Printf(ServersHelp, cmdString)
}
