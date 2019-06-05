// The cli package provides a tool to interface with SpiderWeb at the command
// line
package main

import (
	"strings"
	"reflect"
	"fmt"
	"os"
)

const Help string = `Usage: sw COMMAND ARGS

A command-line tool for interacting with SpiderWeb
`

type cli struct{}

// Run the help command
func HelpCommand() {
	fmt.Print(Help)
}

// Main is the enterypoint for the command line tool
func main() {
	if len(os.Args) > 1 {
		runCommand()
		return
	}
	HelpCommand()
	os.Exit(1)
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
