// The cli package provides a tool to interface with SpiderWeb at the command
// line
package main

import (
	"fmt"
	"os"
)

// Get the help string for the top level cli
func Help() string {
	return `Usage: sw COMMAND ARGS

A command-line tool for interacting with SpiderWeb
`
}

func Hello() string {
	return "Hello World"
}

// Main is the enterypoint for the command line tool
func main() {
	if len(os.Args) > 1 {
		runCommand()
		return
	}
	fmt.Print(Help())
}

func runCommand() {
	switch(os.Args[1]) {
	case "hello":
		fmt.Print(Hello())
	default:
		fmt.Print(Help())
	}
}
