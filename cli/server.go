package main

import (
    "flag"
    "fmt"
)

const ServerHelp string = `Usage: %v server [ARGS]...

Find a server using various filters

Args:
  --help, -h Print this message
  --name, -n Name of the server
`

func (c cli) ServerCommand() {

    flag.NewFlagSet("server", flag.ExitOnError)
    fmt.Printf(ServerHelp, cmdString)
}
