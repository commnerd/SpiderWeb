package main

import (
    "path/filepath"
    "io/ioutil"
    "testing"
    "os/exec"
    "log"
    "fmt"
    "os"
)

// The command executable for the test suite
var TestCmd string = fmt.Sprintf(filepath.Base(os.Args[0]))

// Set up the test suite
func setup() error {
    err := exec.Command("go", "get", "-d").Run()
    if err != nil {
        return err
    }
    err = exec.Command("go", "build", "-o", TestCmd).Run()
    if err != nil {
        return err
    }
    return nil
}

// Tear down the test suite
func teardown() {
    if _, err := os.Stat(TestCmd); err == nil {
        os.Remove(TestCmd)
    }
}

// Prepare the test suite
func TestMain(m *testing.M) {
    err := setup()

    if err == nil {
        m.Run()
    }

    teardown()
}

// Test to ensure we get help output on blank command call
func TestEmptySubcommand(t *testing.T) {
    cmd := exec.Command(TestCmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
        log.Fatal(err)
	}

	if err = cmd.Start(); err != nil {
        log.Fatal(err)
	}

	out, err := ioutil.ReadAll(stdout)
    if err != nil {
        log.Fatal(err)
	}

    help := fmt.Sprintf(Help, TestCmd)
    if string(out) != fmt.Sprintf("%s\n%s", SubCmdError, help) {
        t.Errorf("Expecting help text as output to stderr")
        fmt.Println(string(out))
    }

	if err := cmd.Wait(); err == nil {
        t.Errorf("Expecting error status 1, got 0")
	}
}
