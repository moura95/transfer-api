package receiverrepo

import (
	"github.com/google/uuid"
	"github.com/moura95/transferapi/internal/dto"
	"github.com/moura95/transferapi/internal/entity"
)

type ReceiverRepositoryMemory struct {
	receivers []entity.Receiver
}

func NewMemoryReceiverRepository() *ReceiverRepositoryMemory {
	return &ReceiverRepositoryMemory{
		receivers: []entity.Receiver{
			{
				Uuid: uuid.MustParse("61a218e4-7908-45d7-88bf-6226b53ab321"),
				Cpf:  "Receiver 1",
			},
			{
				Uuid: uuid.MustParse("ef9da75e-949f-4780-92b5-eda71618fc6c"),
				Cpf:  "Receiver 2",
			},
		},
	}
}

func (r *ReceiverRepositoryMemory) GetAll() ([]entity.Receiver, error) {
	return r.receivers, nil
}

func (r *ReceiverRepositoryMemory) Create(receiver dto.CreateReceiverInputDto) (*entity.Receiver, error) {
	receiverInstance := entity.NewReceiver(receiver.Name)
	r.receivers = append(r.receivers, *receiverInstance)
	return receiverInstance, nil
}

func (r *ReceiverRepositoryMemory) GetByID(u uuid.UUID) (*entity.Receiver, error) {
	return nil, nil

}
