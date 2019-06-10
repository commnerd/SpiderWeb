package services

import (
    "../models"
)

type transmitter struct {}

type Transmitter interface{
    Transmit(models.Message) error
}

func NewTransmitter() Transmitter {
    return &transmitter{}
}

func (this *transmitter) Transmit(m models.Message) error {
    return nil
}
