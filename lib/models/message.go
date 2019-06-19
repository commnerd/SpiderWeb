package models

import (
    "net/http"
)

// The groundwork for all node communications
type Message interface {
    BuildRequest() *http.Request
}
