package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type Receiver struct {
	Uuid       uuid.UUID
	Name       string
	CpfCnpj    string
	PixKeyType string
	PixKey     string
	Email      string
	Status     string
}

func NewReceiver(Uuid uuid.UUID, Name, PixKeyType, PixKey, Email, CpfCnpj, Status string) *Receiver {
	return &Receiver{
		Uuid:       Uuid,
		Name:       Name,
		CpfCnpj:    CpfCnpj,
		PixKeyType: PixKeyType,
		PixKey:     PixKey,
		Email:      Email,
		Status:     Status,
	}
}

func (c *Receiver) Validate() error {

	if c.Name == "" {
		return errors.New("name is required")
	}
	if validateCPFCNPJ(c.CpfCnpj) == false {
		return errors.New("cpf/cpnj invalid")
	}

	if validateEmail(c.Email) == false {
		return errors.New("email is required")
	}

	if validatePixKeyType(c.PixKeyType) == false {
		return errors.New("invalid pix key type")
	}
	if validatePixKey(c.PixKey, c.PixKeyType) == false {
		return errors.New("invalid pix key")
	}

	return nil
}

func (r *Receiver) ValidateUpdate() error {
	if r.Status == "Validado" {
		if r.Name != "" {
			return errors.New("cannot update the name when status is 'Validado'")
		}
		if r.CpfCnpj != "" {
			return errors.New("cannot update the CPF/CNPJ when status is 'Validado'")
		}
		if r.PixKeyType != "" {
			return errors.New("cannot update the PixKeyType when status is 'Validado'")
		}
		if r.PixKey != "" {
			return errors.New("cannot update the PixKey when status is 'Validado'")
		}
	}
	return nil
}
func validateCPFCNPJ(cpfCnpj string) bool {
	var valid bool
	if len(cpfCnpj) == 11 {
		valid = validateCPF(cpfCnpj)
	} else if len(cpfCnpj) == 14 {
		valid = validateCNPJ(cpfCnpj)
	}
	return valid
}

func validatePixKeyType(pixKeyType string) bool {
	allowedTypes := map[string]bool{
		"CPF":             true,
		"CNPJ":            true,
		"EMAIL":           true,
		"TELEFONE":        true,
		"CHAVE_ALEATORIA": true,
	}
	return allowedTypes[pixKeyType]
}

func validatePixKey(pixKey, pixKeyType string) bool {
	switch pixKeyType {
	case "CPF":
		return validateCPF(pixKey)
	case "CNPJ":
		return validateCNPJ(pixKey)
	case "EMAIL":
		return validateEmail(pixKey)
	case "TELEFONE":
		return validatePhone(pixKey)
	case "CHAVE_ALEATORIA":
		return validateRandomKey(pixKey)
	default:
		return false
	}
}

func validateCPF(cpf string) bool {
	pattern := regexp.MustCompile(`^[0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2}$`)
	return pattern.MatchString(cpf)
}

func validateCNPJ(cnpj string) bool {
	pattern := regexp.MustCompile(`^[0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2}$`)
	return pattern.MatchString(cnpj)
}

func validateEmail(email string) bool {
	pattern := regexp.MustCompile(`^[a-z0-9+_.-]+@[a-z0-9.-]+$`)
	return pattern.MatchString(email)
}

func validatePhone(telefone string) bool {
	pattern := regexp.MustCompile(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`)
	return pattern.MatchString(telefone)
}

func validateRandomKey(randomKey string) bool {
	pattern := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	return pattern.MatchString(randomKey)
}
