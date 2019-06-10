package models

import (
    "testing"
)

func TestCreation(t *testing.T) {
    m := NewFromSchema("https://schema.org/Thing")

    if len(m.Context) == 0 || len(m.Graph) == 0 {
        t.Errorf("Schema turned up empty")
    }
}
