package dto

import "github.com/google/uuid"

type CreateReceiverInputDto struct {
	Name       string `json:"name"`
	Cpf        string `json:"cpf"`
	PixKeyType string `json:"pix_key_type"`
	PixKey     string `json:"pix_key"`
	Email      string `json:"email"`
	Status     string `json:"status"`
}

type UpdateReceiverInputDto struct {
	Name       string `json:"name"`
	Cpf        string `json:"cpf"`
	PixKeyType string `json:"pix_key_type"`
	PixKey     string `json:"pix_key"`
	Email      string `json:"email"`
}

type ReceiverOutputDto struct {
	Uuid       uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	Cpf        string    `json:"cpf"`
	PixKeyType string    `json:"pix_key_type"`
	PixKey     string    `json:"pix_key"`
	Email      string    `json:"email"`
	Status     string    `json:"status"`
}

type ReceiverListOutputDto struct {
	Receivers []ReceiverOutputDto `json:"receivers"`
}

type DeleteReceiverInputDto struct {
	Uuids []uuid.UUID `json:"uuids"`
}
