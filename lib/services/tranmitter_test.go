package services

import (
    "../models"
    "testing"
)

func TestNewTransmitter(t *testing.T) {
    tr := NewTransmitter()
    if _, ok := tr.(Transmitter); !ok {
        t.Errorf("A transmitter was not instantiated.")
    }
}

func TestTransmitInitRequest(t *testing.T) {
    tr := NewTransmitter()
    err := tr.Transmit(models.NewMessage(models.InitRequest))
    if err != nil {
        t.Errorf("Something went wrong with InitMessage.")
    }
}

func TestTransmitLogin(t *testing.T) {
    tr := NewTransmitter()
    err := tr.Transmit(models.NewMessage(models.LoginRequest))
    if err != nil {
        t.Errorf("Something went wrong with LoginRequest.")
    }
}
