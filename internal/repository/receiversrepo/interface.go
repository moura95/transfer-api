package receiverrepo

import (
	"github.com/google/uuid"
	"github.com/moura95/transferapi/internal/entity"
)

type IReceiverRepository interface {
	GetAll() ([]entity.Receiver, error)
	Create(receiver entity.Receiver) error
	GetByID(uuid.UUID) (*entity.Receiver, error)
	Update(uuid.UUID, *entity.Receiver) error
}
