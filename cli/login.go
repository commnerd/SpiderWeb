package main

import (
    "golang.org/x/crypto/ssh/terminal"
    "strings"
    "syscall"
    "bufio"
    "log"
    "fmt"
    "os"
)

const (
    // EmailPrompt : Query text for username
    EmailPrompt = "Email: "
    // PasswordPrompt : Query for password
    PasswordPrompt = "Password: "
)

// LoginHelp : Help texts
const LoginHelp = `Usage: %v login

Authenticate to the network
`

// LoginCommand : Login subcommand
func (c cli) LoginCommand() {
    showHelp := false
    for _, val := range os.Args[1:] {
        if val == "--help" || val == "-help" {
            showHelp = true
        }
    }

    if showHelp {
        fmt.Printf(LoginHelp, cmdString)
        os.Exit(0)
    }

    fmt.Print(EmailPrompt)
    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')
    if err != nil {
        log.Panic(err)
    }
    // convert CRLF to LF
    text = strings.Replace(text, "\n", "", -1)

    fmt.Print(PasswordPrompt)
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    password := string(bytePassword)
    fmt.Println() // it's necessary to add a new line after user's input
    fmt.Printf("Your password has leaked, it is '%s'", password)
}
