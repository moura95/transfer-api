package receiverrepo

import (
	"github.com/google/uuid"
	"github.com/moura95/transferapi/internal/entity"
)

type IReceiverRepository interface {
	GetAll(filters entity.Filter) (response *GetAllResponse, err error)
	Create(receiver entity.Receiver) error
	GetByID(uuid.UUID) (*entity.Receiver, error)
	Update(uuid.UUID, *entity.Receiver) error
	HardDelete(uid uuid.UUID) error
	BulkDelete([]string) error
}
