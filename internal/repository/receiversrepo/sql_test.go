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

func TestCustomerRepository(t *testing.T) {
	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(filepath.Join("../../..", "db/migrations", "000001_init_schema.up.sql.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})
	mappedPort, err := pgContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatal(err)
	}

	connStr := "postgres://postgres:postgres@localhost:" + mappedPort.Port() + "/test-db?sslmode=disable"

	conn, err := db.ConnectPostgres(connStr)
	store := conn.DB()
	if err != nil {
		fmt.Println("Failed to Connected Database")
		panic(err)
	}
	assert.NoError(t, err)

	customerRepo := NewReceiverRepository(store)
	assert.NoError(t, err)

	err = customerRepo.Create(*entity.NewReceiver(uuid.New(), "Henry", "CPF", "henry@gmail.com", "henry@gmail.com", "123456789", "Validado"))
	assert.NoError(t, err)

	assert.NoError(t, err)

	response, err := customerRepo.GetAll(entity.Filter{
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
