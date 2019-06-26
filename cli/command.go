// The cli package provides a tool to interface with SpiderWeb at the command
// line
package main

import (
	"path/filepath"
	"strings"
	"reflect"
	"fmt"
	"os"
)

const (
	// SubCmdError : Subcommand error text
	SubCmdError string = "Subcommand expected"

	// Help : Help text
	Help string = `Usage: %v COMMAND

A command-line tool for interacting with SpiderWeb

Commands:
  login     Send username/password to the network for authentication
  server    Find a server using various filters
  servers   List servers owned by you
  volumes   List volumes owned by you
`
)

var cmdString = fmt.Sprintf(filepath.Base(os.Args[0]))

type cli struct{}

// HelpCommand : Run the help command
func HelpCommand() {
	fmt.Printf(Help, cmdString)
}

// Main is the enterypoint for the command line tool
func main() {
	if len(os.Args) <= 1 {
		fmt.Println(SubCmdError)
		HelpCommand()
		os.Exit(1)
	}

	runCommand()
}

// runCommand : Map and execute the appropriate command
func runCommand() {
	cmd := strings.Title(os.Args[1]) + "Command"
	c := cli{}
	target, ok := reflect.TypeOf(c).MethodByName(cmd)
	if ok {
	    target.Func.Call([]reflect.Value{reflect.ValueOf(c)})
	}
}
