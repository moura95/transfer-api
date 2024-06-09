package receiverrepo

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/moura95/transferapi/db"
	"github.com/moura95/transferapi/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	pgContainer testcontainers.Container
	connStr     string
)

func setupPostgresContainer() (func(), error) {
	ctx := context.Background()

	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(filepath.Join("../../..", "db/migrations", "000001_init_schema.up.sql.sql")),
		postgres.WithInitScripts(filepath.Join("../../..", "db/migrations", "000002_seed_schema.up.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return nil, err
	}

	connStr = "postgres://postgres:postgres@localhost:" + mappedPort.Port() + "/test-db?sslmode=disable"

	return func() {
		if err := container.Terminate(ctx); err != nil {
			fmt.Printf("failed to terminate pgContainer: %s", err)
		}
	}, nil
}

func TestMain(m *testing.M) {
	cleanup, err := setupPostgresContainer()
	if err != nil {
		panic(fmt.Sprintf("Failed to set up PostgreSQL container: %s", err))
	}
	defer cleanup()

	m.Run()
}

func TestReceiverRepository_CreateAndGetAllWithFilter(t *testing.T) {
	conn, err := db.ConnectPostgres(connStr)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	store := conn.DB()

	receiverRepo := NewReceiverRepository(store)
	err = receiverRepo.Create(*entity.NewReceiver(uuid.New(), "Henry", "CPF", "henry@gmail.com", "henry@gmail.com", "123456789", "Validado"))
	assert.NoError(t, err)

	response, err := receiverRepo.GetAll(entity.Filter{
		Status:     "Validado",
		Name:       "Henry",
		PixKeyType: "CPF",
	})
	assert.NoError(t, err)
	assert.NotNil(t, response.Receivers)
	assert.Len(t, response.Receivers, 1)
	receiver := response.Receivers[0]

	assert.Equal(t, "Henry", receiver.Name)
	assert.Equal(t, "CPF", receiver.PixKeyType)
	assert.Equal(t, "henry@gmail.com", receiver.PixKey)
	assert.Equal(t, "henry@gmail.com", receiver.Email)
	assert.Equal(t, "Validado", receiver.Status)

}

func TestReceiverRepository_GetAll(t *testing.T) {
	conn, err := db.ConnectPostgres(connStr)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	store := conn.DB()

	receiverRepo := NewReceiverRepository(store)
	response, err := receiverRepo.GetAll(entity.Filter{
		Limit: 100,
	})
	assert.NoError(t, err)
	assert.NotNil(t, response.Receivers)
	assert.Len(t, response.Receivers, 64)
	// Receiver 1
	assert.Equal(t, "Ana Costa", response.Receivers[0].Name)
	assert.Equal(t, "55566677778", response.Receivers[0].CpfCnpj)
	assert.Equal(t, "CNPJ", response.Receivers[0].PixKeyType)
	assert.Equal(t, "12345678000199", response.Receivers[0].PixKey)
	assert.Equal(t, "ana.costan@example.com", response.Receivers[0].Email)
	assert.Equal(t, "Rascunho", response.Receivers[0].Status)

	// Receiver 10
	assert.Equal(t, "Carla Rocha", response.Receivers[10].Name)
	assert.Equal(t, "11137655577", response.Receivers[10].CpfCnpj)
	assert.Equal(t, "CHAVE_ALEATORIA", response.Receivers[10].PixKeyType)
	assert.Equal(t, "9b1f3c9e-3e35-47b8-9299-5c7a3d1b5c90", response.Receivers[10].PixKey)
	assert.Equal(t, "carlinha.rocha@example.com", response.Receivers[10].Email)
	assert.Equal(t, "Rascunho", response.Receivers[10].Status)
	// Receiver 20
	assert.Equal(t, "Gabriela Santos", response.Receivers[20].Name)
	assert.Equal(t, "44415566678", response.Receivers[20].CpfCnpj)
	assert.Equal(t, "CHAVE_ALEATORIA", response.Receivers[20].PixKeyType)
	assert.Equal(t, "5a8f9e2a-9eaf-4f6a-a15c-24b5eae1d451", response.Receivers[20].PixKey)
	assert.Equal(t, "gabrielllla.santos@example.com", response.Receivers[20].Email)
	assert.Equal(t, "Rascunho", response.Receivers[0].Status)

	// Receiver 32
	assert.Equal(t, "Juliana Almeida", response.Receivers[32].Name)
	assert.Equal(t, "10233344452", response.Receivers[32].CpfCnpj)
	assert.Equal(t, "TELEFONE", response.Receivers[32].PixKeyType)
	assert.Equal(t, "+5521999997777", response.Receivers[32].PixKey)
	assert.Equal(t, "juliaaaanna.almeida@example.com", response.Receivers[32].Email)
	assert.Equal(t, "Rascunho", response.Receivers[32].Status)
}

func TestReceiverRepository_UpdateReceiverValidadoSuccess(t *testing.T) {
	conn, err := db.ConnectPostgres(connStr)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	store := conn.DB()

	receiverRepo := NewReceiverRepository(store)
	uid := uuid.MustParse("4c57ae7f-0eb0-4bfe-9c62-d9a87880ea61")

	beforeReceiver, err := receiverRepo.GetByID(uid)
	assert.NoError(t, err)
	assert.NotNil(t, beforeReceiver)
	assert.Equal(t, "jo3ao.silva@example.com", beforeReceiver.Email)
	assert.Equal(t, "Validado", beforeReceiver.Status)

	receiver := entity.NewReceiver(uid, "", "", "", "henryupdateemail@gmail.com", "", "")
	err = receiver.ValidateUpdate()
	assert.NoError(t, err)
	err = receiverRepo.Update(uid, receiver)
	assert.NoError(t, err)

	newReceiver, err := receiverRepo.GetByID(uid)
	assert.Equal(t, "henryupdateemail@gmail.com", newReceiver.Email)
}

func TestReceiverRepository_UpdateReceiverValidadoFailed(t *testing.T) {
	conn, err := db.ConnectPostgres(connStr)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	store := conn.DB()

	receiverRepo := NewReceiverRepository(store)
	uid := uuid.MustParse("1bc4ab0d-1384-4343-881f-12513f1510f2")

	beforeReceiver, err := receiverRepo.GetByID(uid)
	assert.NoError(t, err)
	assert.NotNil(t, beforeReceiver)
	assert.Equal(t, "p3edro.santos@example.com", beforeReceiver.Email)
	assert.Equal(t, "Validado", beforeReceiver.Status)
	receiver := entity.NewReceiver(uid, "", "", "", "joaopedroupdate@gmail.com", "02651133399", "Validado")
	err = receiver.ValidateUpdate()
	assert.Equal(t, "cannot update the CPF/CNPJ when status is Validado", err.Error())

}

func TestReceiverRepository_HardDeleteReceiver(t *testing.T) {
	conn, err := db.ConnectPostgres(connStr)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	store := conn.DB()

	receiverRepo := NewReceiverRepository(store)
	uid := uuid.MustParse("4c57ae7f-0eb0-4bfe-9c62-d9a87880ea61")

	err = receiverRepo.HardDelete(uid)
	assert.NoError(t, err)
}

func TestReceiverRepository_BulkDelete(t *testing.T) {
	conn, err := db.ConnectPostgres(connStr)
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	store := conn.DB()

	receiverRepo := NewReceiverRepository(store)
	filters := entity.NewFilter("", "", "", "", 100, 1)
	responseGetAll, err := receiverRepo.GetAll(filters)
	assert.NoError(t, err)
	assert.Equal(t, 63, len(responseGetAll.Receivers))

	// delete 15 uuids

	var uuids []string

	for i := 0; i < 15; i++ {
		uuids = append(uuids, responseGetAll.Receivers[i].Uuid.String())
	}
	err = receiverRepo.BulkDelete(uuids)
	assert.NoError(t, err)

	// get all
	responseGetAll, err = receiverRepo.GetAll(filters)
	assert.NoError(t, err)
	assert.Equal(t, 48, len(responseGetAll.Receivers))

}
