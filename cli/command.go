// The cli package provides a tool to interface with SpiderWeb at the command
// line
package main

import (
	"strings"
	"reflect"
	"fmt"
	"os"
)

type cli struct{}

// Get the help string for the top level cli
func GetHelp() string {
	return `Usage: sw COMMAND ARGS

A command-line tool for interacting with SpiderWeb
`
}

// Run the help command
func HelpCommand() {
	fmt.Print(GetHelp())
}

// Main is the enterypoint for the command line tool
func main() {
	if len(os.Args) > 1 {
		runCommand()
		return
	}
	HelpCommand()
}

// Map and execute the appropriate command
func runCommand() {
	cmd := strings.Title(os.Args[1]) + "Command"
	c := cli{}
	target, ok := reflect.TypeOf(c).MethodByName(cmd)
	if ok {
	    target.Func.Call([]reflect.Value{reflect.ValueOf(c)})
	}
	HelpCommand()
	os.Exit(1)
}
