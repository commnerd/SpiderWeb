// The cli package provides a tool to interface with SpiderWeb at the command
// line
package main

import (
	"strings"
	"reflect"
	"fmt"
	"os"
)

const (
	SubCmdError string = "Subcommand expected"
	Help string = `Usage: sw COMMAND

A command-line tool for interacting with SpiderWeb

Commands:
  server    Find a server using various filters
  servers   List servers owned by you
`
)

type cli struct{}

// Run the help command
func HelpCommand() {
	fmt.Print(Help)
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

// Map and execute the appropriate command
func runCommand() {
	cmd := strings.Title(os.Args[1]) + "Command"
	c := cli{}
	target, ok := reflect.TypeOf(c).MethodByName(cmd)
	if ok {
	    target.Func.Call([]reflect.Value{reflect.ValueOf(c)})
	}
}
