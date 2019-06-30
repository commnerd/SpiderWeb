package flags

import (
    "os"
)

var flagMap map[string]int
var term bool

func init() {
    term = false
    flagMap := *new(map[string]int)
    for i, arg := range os.Args[1:] {
        if isValidFlag(arg) {
            flagMap[getFlagLabel(arg)] = i
        }
    }
}

func MapShortToLong(s string, l string) {

}

func isValidFlag(str string) bool {
    if str[0] != '-' {
        return false
    }
    if str[1] == '-' {
        if len(str) == 3 {
            return false
        }
        if inCharRange(str[2], 'a', 'z') {
            return true
        }
        if inCharRange(str[2], 'A', 'Z') {
            return true
        }
    }
    if len(str) > 2 {
        return false
    }
    if inCharRange(str[1], 'a', 'z') {
        return true
    }
    if inCharRange(str[1], 'A', 'Z') {
        return true
    }

    return false
}

func getFlagLabel(str string) string {
    if isValidFlag(str) {
        return shaveDashes(str)
    }
    return str
}

func shaveDashes(str string) string {
    for str[0] == '-' {
        str = str[1:]
    }
    return str
}

func inCharRange(t byte, start byte, end byte) bool {
    if t >= start && t <= end {
        return true
    }
    return false
}
