package flags

import (
    "testing"
    "os"
)

func init() {
    os.Args = []string{"cmd", "-a", "123", "--bravo", "something"}
}

func TestInCharRange(t *testing.T) {
    expected := true
    got := inCharRange('c', 'a', 'z')
    if got != expected {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestInCharRangeStartEdge(t *testing.T) {
    expected := true
    got := inCharRange('a', 'a', 'z')
    if got != expected {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestInCharRangeEndEdge(t *testing.T) {
    expected := true
    got := inCharRange('z', 'a', 'z')
    if got != expected {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestNotInCharRange(t *testing.T) {
    expected := false
    got := inCharRange('c', 'l', 'o')
    if got != expected {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestNotInCharRangeStartEdge(t *testing.T) {
    expected := false
    got := inCharRange('k', 'l', 'o')
    if got != expected {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestNotInCharRangeEndEdge(t *testing.T) {
    expected := false
    got := inCharRange('p', 'a', 'z')
    if got != expected {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestIsValidLongFlag(t *testing.T) {
    expected := true
    got := isValidFlag("--foo")
    if expected != got {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestNotIsValidLongFlag(t *testing.T) {
    expected := false
    got := isValidFlag("--f")
    if expected != got {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestIsValidShortFlag(t *testing.T) {
    expected := true
    got := isValidFlag("-f")
    if expected != got {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestNotIsValidShortFlag(t *testing.T) {
    expected := false
    got := isValidFlag("-foo")
    if expected != got {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestNonFlag(t *testing.T) {
    expected := false
    got := isValidFlag("foo")
    if expected != got {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestGetFlagLongLabel(t *testing.T) {
    expected := "foo"
    got := getFlagLabel("--foo")
    if expected != got {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestGetFlagShortLabel(t *testing.T) {
    expected := "f"
    got := getFlagLabel("-f")
    if expected != got {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestShaveDashes(t *testing.T) {
    expected := "foo"
    got := shaveDashes("--foo")
    if expected != got {
        t.Errorf("Expecting: \"%v\", Got: \"%v\".", expected, got)
    }
}

func TestDefaultArgs(t *testing.T) {
    if flagMap["a"] != 0 {
        t.Errorf("Expecting: \"123\", Got: \"%v\".", flagMap["a"])
    }
}
