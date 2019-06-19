package services

import (
    "../models"
    "net/http"
)

type transmitter struct {
    client *http.Client
}

type Transmitter interface{
    Transmit(models.Message) error
}

func NewTransmitter() Transmitter {        
    return &transmitter{
        client: &http.Client{}
    }
}

func (this *transmitter) Transmit(m models.Message) error {
    return nil
}
