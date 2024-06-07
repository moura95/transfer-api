package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Receiver struct {
	Uuid uuid.UUID
	Cpf  string
}

func NewReceiver(cpf string) *Receiver {
	return &Receiver{
		Cpf: cpf,
	}
}

func (c *Receiver) Validate() error {

	if c.Cpf == "" {
		return errors.New("cpf is required")
	}
	return nil
}
