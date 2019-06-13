package main

import (
    "flag"
    "fmt"
)

const LoginHelp string = `Usage: %v login

Authenticate to the network
`

func (c cli) LoginCommand() {
    flag.Bool("-help", true, "")
    flag.NewFlagSet("login", flag.ExitOnError)
    fmt.Printf(LoginHelp, cmdString)
}
