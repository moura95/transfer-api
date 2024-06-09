package receiverrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/moura95/transferapi/internal/entity"
)

type ReceiverRepository struct {
	db *sqlx.DB
}

func NewReceiverRepository(db *sqlx.DB) IReceiverRepository {
	return &ReceiverRepository{db: db}
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

type GetAllResponse struct {
	Receivers    []entity.Receiver
	TotalRecords int
	TotalPages   int
	CurrentPage  int
	Limit        int
}

func (r ReceiverRepository) GetAll(filters entity.Filter) (*GetAllResponse, error) {
	var response GetAllResponse
	defaultLimit := 10
	query := "SELECT uuid, name, pix_key_type, pix_key, email, cpf_cnpj, status FROM receivers WHERE 1 = 1 "

	if filters.Limit <= 0 {
		filters.Limit = defaultLimit
	}
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.Name != "" {
		query += " AND name ILIKE '%" + filters.Name + "%'"
	}
	if filters.Status != "" {
		query += " AND status ILIKE '" + filters.Status + "'"
	}
	if filters.PixKeyType != "" {
		query += " AND pix_key_type ILIKE '" + filters.PixKeyType + "'"
	}
	if filters.PixKeyValue != "" {
		query += " AND pix_key ILIKE '" + filters.PixKeyValue + "'"
	}
	// remove the last 'AND'
	query = strings.TrimSuffix(query, " AND")
	queryOrderBy := " ORDER BY name LIMIT $1 OFFSET $2 "
	query = query + queryOrderBy

	offset := (filters.Page - 1) * filters.Limit

	queryCount := "SELECT COUNT(*) FROM receivers"
	err := r.db.Get(&response.TotalRecords, queryCount)
	if err != nil {
		return nil, err
	}

	response.TotalPages = (response.TotalRecords + filters.Limit - 1) / filters.Limit
	response.CurrentPage = filters.Page
	response.Limit = filters.Limit

	var receiverModels []ReceiverModel
	if err := r.db.Select(&receiverModels, query, filters.Limit, offset); err != nil {
		return nil, err
	}

	for _, receiver := range receiverModels {
		response.Receivers = append(response.Receivers, *entity.ToEntity(receiver.Uuid, receiver.Name, receiver.PixKeyType, receiver.PixKey, receiver.Email, receiver.CpfCnpj, receiver.Status))
	}

	return &response, nil
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
