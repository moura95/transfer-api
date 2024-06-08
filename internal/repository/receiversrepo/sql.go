package receiverrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

const (
	statusRascunho = "Rascunho"
	statusValidado = "Validado"
)

func NewReceiverRepository(db *sqlx.DB, log *zap.SugaredLogger) IReceiverRepository {
	return &ReceiverRepository{db: db, logger: log}
}

func (r ReceiverRepository) GetAll(filters map[string]string) ([]entity.Receiver, error) {
	defaultLimit := 10
	defaultOffset := 0

	limit, err := strconv.Atoi(filters["limit"])
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	offset, err := strconv.Atoi(filters["offset"])
	if err != nil || offset < 0 {
		offset = defaultOffset
	}

	query := "SELECT uuid, name, pix_key_type, pix_key, email, cpf_cnpj, status FROM receivers ORDER BY name LIMIT $1 OFFSET $2"
	var receivers []ReceiverModel
	if err := r.db.Select(&receivers, query, limit, offset); err != nil {
		return nil, err
	}

	var receiversEntity []entity.Receiver
	for _, receiver := range receivers {
		receiversEntity = append(receiversEntity, *entity.ToEntity(receiver.Uuid, receiver.Name, receiver.PixKeyType, receiver.PixKey, receiver.Email, receiver.CpfCnpj, receiver.Status))
	}
	return receiversEntity, nil
}

func (r ReceiverRepository) Create(receiver entity.Receiver) error {

	query := `
        INSERT INTO receivers (name, cpf_cnpj, pix_key_type, pix_key, email, status)
        VALUES ($1, $2, $3, $4, $5,$6)
    `
	args := []interface{}{
		receiver.Name,
		receiver.CpfCnpj,
		receiver.PixKeyType,
		receiver.PixKey,
		receiver.Email,
		receiver.Status,
	}
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r ReceiverRepository) GetByID(uid uuid.UUID) (*entity.Receiver, error) {
	var result ReceiverModel

	query := `SELECT * FROM receivers WHERE uuid = $1`

	err := r.db.Get(&result, query, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return entity.NewReceiver(result.Uuid, result.Name, result.PixKeyType, result.PixKey, result.Email, result.CpfCnpj, result.Status), nil
}

func (r ReceiverRepository) Update(uid uuid.UUID, receiver *entity.Receiver) error {
	query := `
        UPDATE receivers 
        SET 
            name = COALESCE($2, name),
            cpf_cnpj = COALESCE($3, cpf_cnpj),
            pix_key_type = COALESCE($4, pix_key_type),
            pix_key = COALESCE($5, pix_key),
            email = COALESCE($6, email),
            update_at = $7
        WHERE uuid = $1
    `

	args := []interface{}{
		uid,
		sql.NullString{String: receiver.Name, Valid: receiver.Name != ""},
		sql.NullString{String: receiver.CpfCnpj, Valid: receiver.CpfCnpj != ""},
		sql.NullString{String: receiver.PixKeyType, Valid: receiver.PixKeyType != ""},
		sql.NullString{String: receiver.PixKey, Valid: receiver.PixKey != ""},
		sql.NullString{String: receiver.Email, Valid: receiver.Email != ""},
		time.Now(),
	}

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r ReceiverRepository) HardDelete(uuid uuid.UUID) error {
	query := "DELETE FROM receivers WHERE uuid = :UUID"
	_, err := r.db.NamedExec(query, map[string]interface{}{"UUID": uuid})
	return err
}

func (r ReceiverRepository) BulkDelete(uuids []string) error {
	deleteUUIDs := make([]string, len(uuids))
	for i, uuid := range uuids {
		deleteUUIDs[i] = fmt.Sprintf("'%s'", uuid)
	}
	query := fmt.Sprintf("DELETE FROM receivers WHERE uuid IN (%s)", strings.Join(deleteUUIDs, ", "))
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

type ReceiverModel struct {
	Uuid       uuid.UUID `db:"uuid"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	CpfCnpj    string    `db:"cpf_cnpj"`
	PixKeyType string    `db:"pix_key_type"`
	PixKey     string    `db:"pix_key"`
	Status     string    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"update_at"`
}
