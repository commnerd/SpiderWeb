package main

import (
    "flag"
    "fmt"
)

const ServerHelp string = `Usage: sw server [ARGS]...

Find a server using various filters

Args:
  --help, -h Print this message
  --name, -n Name of the server
`

func (c cli) ServerCommand() {

    flag.NewFlagSet("server", flag.ExitOnError)
    fmt.Print(ServerHelp)
}
