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
	SubCmdError string = "Subcommand expected"
	Help string = `Usage: %v COMMAND

A command-line tool for interacting with SpiderWeb

Commands:
  server    Find a server using various filters
  servers   List servers owned by you
`
)

type cli struct{}

// Run the help command
func HelpCommand() {
	fmt.Printf(Help, fmt.Sprintf(filepath.Base(os.Args[0])))
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
