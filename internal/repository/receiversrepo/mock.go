package receiverrepo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/moura95/transferapi/internal/entity"
)

type ReceiverRepositoryMock struct {
	receivers []entity.Receiver
}

func NewMockReceiverRepository() *ReceiverRepositoryMock {
	return &ReceiverRepositoryMock{
		receivers: []entity.Receiver{
			{
				Uuid:       uuid.MustParse("4c57ae7f-0eb0-4bfe-9c62-d9a87880ea61"),
				Name:       "João Silva",
				CpfCnpj:    "12345678921",
				Email:      "jo3ao.silva@example.com",
				Status:     "Validado",
				PixKeyType: "CPF",
				PixKey:     "12345678921",
			},
			{
				Uuid:       uuid.MustParse("66cfbbed-e3f8-4f2a-935d-665a368a915e"),
				Name:       "Maria Oliveira",
				CpfCnpj:    "98765433100",
				Email:      "maria.o1liveira@example.com",
				Status:     "Validado",
				PixKeyType: "EMAIL",
				PixKey:     "maria.oliveira@example.com",
			},
			{
				Uuid:       uuid.MustParse("1bc4ab0d-1384-4343-881f-12513f1510f2"),
				Name:       "Pedro Santos",
				CpfCnpj:    "11122233344",
				Email:      "p3edro.santos@example.com",
				Status:     "Validado",
				PixKeyType: "TELEFONE",
				PixKey:     "+5511999998888",
			},

			{
				Uuid:       uuid.MustParse("796bb798-a29e-4271-9e70-f5c065374257"),
				Name:       "Ana Costa",
				CpfCnpj:    "55566677788",
				Email:      "ana.cost4a@example.com",
				Status:     "Rascunho",
				PixKeyType: "CNPJ",
				PixKey:     "12345678000199",
			},
			{
				Uuid:       uuid.MustParse("3ee27437-fdb6-48ce-85f4-0b16e046c82a"),
				Name:       "Lucas Lima",
				CpfCnpj:    "44455566677",
				Email:      "lucas.lima@example.com",
				Status:     "Validado",
				PixKeyType: "CHAVE_ALEATORIA",
				PixKey:     "5a8f9e2a-9eaf-4f6a-a15c-24b5eae1d452",
			},
		},
	}
}

func (r *ReceiverRepositoryMock) GetAll(filters entity.Filter) (*GetAllResponse, error) {
	var response GetAllResponse
	defaultLimit := 10

	if filters.Limit <= 0 {
		filters.Limit = defaultLimit
	}

	if filters.Page < 1 {
		filters.Page = 1
	}

	offset := (filters.Page - 1) * filters.Limit
	response.TotalRecords = len(r.receivers)
	response.TotalPages = (response.TotalRecords + filters.Limit - 1) / filters.Limit
	response.CurrentPage = filters.Page

	if offset > response.TotalRecords {
		return &response, errors.New("page number out of range")
	}

	end := offset + filters.Limit
	if end > response.TotalRecords {
		end = response.TotalRecords
	}

	response.Receivers = r.receivers[offset:end]

	return &response, nil
}

func (r *ReceiverRepositoryMock) Create(receiver entity.Receiver) error {
	r.receivers = append(r.receivers, receiver)
	return nil
}

func (r *ReceiverRepositoryMock) GetByID(u uuid.UUID) (*entity.Receiver, error) {
	for _, rec := range r.receivers {
		if rec.Uuid == u {
			return &rec, nil
		}
	}
	return nil, errors.New("not Found")

}

func (r ReceiverRepositoryMock) Update(uid uuid.UUID, receiver *entity.Receiver) error {
	for _, rec := range r.receivers {
		if rec.Uuid == uid {
			if rec.Status == "Validado" {
				rec.Email = receiver.Email
				return nil
			} else {
				rec.Name = receiver.Name
				rec.Email = receiver.Email
				rec.CpfCnpj = receiver.CpfCnpj
				rec.PixKey = receiver.PixKey
				rec.PixKeyType = receiver.PixKeyType
				return nil
			}

		}
	}

	return errors.New("Failed to updated")
}

func (r ReceiverRepositoryMock) HardDelete(uid uuid.UUID) error {
	for i, rec := range r.receivers {
		if rec.Uuid == uid {
			r.receivers = append(r.receivers[:i], r.receivers[i+1:]...)
			return nil
		}
	}
	return errors.New("receiver not found")
}

func (r *ReceiverRepositoryMock) BulkDelete(strings []string) error {

	panic("implement me")
}
