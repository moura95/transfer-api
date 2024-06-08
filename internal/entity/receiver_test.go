package entity

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestNewReceiverValid(t *testing.T) {
	receiver := NewReceiver(uuid.New(), "Joao Silva", "CPF", "12345678921", "joao.silva@example.com", "12345678901", "Validado")
	err := receiver.Validate()
	if err != nil {
		t.Error("Expected no error, got:", err)
	}
	fmt.Println(receiver)
}

func TestNewReceiverInvalidEmail(t *testing.T) {
	receiver := NewReceiver(uuid.New(), "Joao Silva", "CPF", "12345678921", "invalid", "12345678901", "Validado")
	err := receiver.Validate()
	if err.Error() != "email invalid" {
		t.Errorf("Expected error 'email invalid', got '%s'", err.Error())
	}
}

func TestNewReceiverInvalidCPF(t *testing.T) {
	receiver := NewReceiver(uuid.New(), "Joao Silva", "CPF", "12345678921", "joao.silva@example.com", "invalid", "Validado")
	err := receiver.Validate()
	if err.Error() != "cpf/cpnj invalid" {
		t.Errorf("Expected error 'cpf/cpnj invalid', got '%s'", err.Error())
	}
}

func TestNewReceiverInvalidName(t *testing.T) {
	receiver := NewReceiver(uuid.New(), "", "CPF", "12345678921", "joao.silva@example.com", "12345678901", "Validado")
	err := receiver.Validate()
	if err.Error() != "name is required" {
		t.Errorf("Expected error 'name is required', got '%s'", err.Error())
	}
}

func TestNewReceiverInvalidPixKeyType(t *testing.T) {
	receiver := NewReceiver(uuid.New(), "Joao", "invalid", "12345678921", "joao.silva@example.com", "12345678901", "Validado")
	err := receiver.Validate()
	if err.Error() != "invalid pix key type" {
		t.Errorf("Expected error 'invalid pix key type', got '%s'", err.Error())
	}
}
