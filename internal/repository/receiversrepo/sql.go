package receiverrepo

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/moura95/transferapi/internal/entity"
	"go.uber.org/zap"
)

type ReceiverRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewReceiverRepository(db *sqlx.DB, log *zap.SugaredLogger) IReceiverRepository {
	return &ReceiverRepository{db: db, logger: log}
}

func (r ReceiverRepository) GetAll() ([]entity.Receiver, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReceiverRepository) Create(receiver entity.Receiver) error {
	//TODO implement me
	panic("implement me")
}

func (r ReceiverRepository) GetByID(u uuid.UUID) (*entity.Receiver, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReceiverRepository) Update(u uuid.UUID, receiver *entity.Receiver) error {
	//TODO implement me
	panic("implement me")
}

type ReceiverModel struct {
	Uuid      uuid.UUID `db:"uuid"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	TaxID     string    `db:"tax_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"update_at"`
}
