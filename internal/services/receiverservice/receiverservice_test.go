package receiverservice

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/moura95/transferapi/internal/entity"
	receiverrepo "github.com/moura95/transferapi/internal/repository/receiversrepo"
	"github.com/stretchr/testify/assert"
)

func NewReceiverServiceTest(repo receiverrepo.IReceiverRepository) *Service {
	return &Service{
		repository: repo,
	}
}

func TestCreateReceiver(t *testing.T) {
	mockRepo := receiverrepo.NewMemoryReceiverRepository()
	service := NewReceiverServiceTest(mockRepo)

	receiverInstance := entity.NewReceiver(uuid.New(), "Recebedor 1", "CPF", "12345678921", "driver1@example.com", "12345678901", statusRascunho)
	err := receiverInstance.Validate()
	if err != nil {
		fmt.Println(err)
	}

	err = service.repository.Create(*receiverInstance)
	if err != nil {
		t.Error("Failed to created")
	}
	assert.NoError(t, err)

}

func TestGetAll(t *testing.T) {
	mockRepo := receiverrepo.NewMemoryReceiverRepository()
	service := NewReceiverServiceTest(mockRepo)
	filters := map[string]string{
		"limit": "10",
	}

	receivers, _, err := service.List(filters)
	if err != nil {
		t.Error("Failed to list")
	}
	assert.NoError(t, err)
	assert.Equal(t, receivers[0].Name, "João Silva")
	assert.Equal(t, receivers[0].CpfCnpj, "12345678921")
	assert.Equal(t, receivers[0].Email, "jo3ao.silva@example.com")
	assert.Equal(t, receivers[0].PixKeyType, "CPF")
	assert.Equal(t, receivers[0].PixKey, "12345678921")
	assert.Equal(t, receivers[0].Status, "Validado")

}

func TestGetByID(t *testing.T) {
	mockRepo := receiverrepo.NewMemoryReceiverRepository()
	service := NewReceiverServiceTest(mockRepo)

	uid := uuid.MustParse("66cfbbed-e3f8-4f2a-935d-665a368a915e")
	receiver, err := service.GetByID(uid)
	if err != nil {
		t.Error("Failed to find")
	}
	assert.NoError(t, err)
	assert.Equal(t, receiver.Name, "Maria Oliveira")
	assert.Equal(t, receiver.CpfCnpj, "98765433100")
	assert.Equal(t, receiver.Email, "maria.o1liveira@example.com")
	assert.Equal(t, receiver.PixKeyType, "EMAIL")
	assert.Equal(t, receiver.PixKey, "maria.oliveira@example.com")
	assert.Equal(t, receiver.Status, "Validado")

}

func TestUpdateReceiverRascunho(t *testing.T) {
	mockRepo := receiverrepo.NewMemoryReceiverRepository()
	service := NewReceiverServiceTest(mockRepo)

	uid := uuid.MustParse("796bb798-a29e-4271-9e70-f5c065374257")
	err := service.Update(uid, "Maria Update", "EMAIL", "maria.oliveira@example.com", "maria.o1liveira@example.com", "98765433100")
	if err != nil {
		t.Error("Failed to update")
	}
	assert.NoError(t, err)

}

func TestUpdateReceiverValidatorSuccess(t *testing.T) {
	mockRepo := receiverrepo.NewMemoryReceiverRepository()
	service := NewReceiverServiceTest(mockRepo)

	uid := uuid.MustParse("3ee27437-fdb6-48ce-85f4-0b16e046c82a")
	err := service.Update(uid, "", "", "", "maria.o1liveira@example.com", "")
	if err != nil {
		t.Error("Failed to update")
	}
	assert.NoError(t, err)

}

func TestUpdateReceiverValidatorFailed(t *testing.T) {
	mockRepo := receiverrepo.NewMemoryReceiverRepository()
	service := NewReceiverServiceTest(mockRepo)

	uid := uuid.MustParse("3ee27437-fdb6-48ce-85f4-0b16e046c82a")
	err := service.Update(uid, "Maria Update", "", "", "", "")
	assert.Equal(t, err.Error(), "cannot update the name when status is Validado")

}

func TestHardDelete(t *testing.T) {
	mockRepo := receiverrepo.NewMemoryReceiverRepository()
	service := NewReceiverServiceTest(mockRepo)

	uid := uuid.MustParse("3ee27437-fdb6-48ce-85f4-0b16e046c82a")
	err := service.Delete(uid)
	assert.NoError(t, err)

}
