package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/transferapi/config"
	"github.com/moura95/transferapi/internal/api/controller/receivercontroller"
	"github.com/moura95/transferapi/internal/dto"
	receiverrepo "github.com/moura95/transferapi/internal/repository/receiversrepo"
	"github.com/moura95/transferapi/internal/services/receiverservice"
	"go.uber.org/zap"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestGetAllReceivers(t *testing.T) {
	router := setupRouter()

	routes := router.Group("/")
	mockRepo := receiverrepo.NewMockReceiverRepository()
	service := receiverservice.NewReceiverService(mockRepo, config.Config{}, zap.S())
	receivercontroller.NewReceiverRouter(service, zap.S()).SetupReceiverRoute(routes)

	expectedReceivers := []dto.ReceiverOutputDto{
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
	}

	// Prepare the request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/receiver?limit=10", nil)
	router.ServeHTTP(w, req)

	// Check the response code
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Unmarshal the response
	var actualReceivers []dto.ReceiverOutputDto
	if err := json.Unmarshal(w.Body.Bytes(), &actualReceivers); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Verify the response
	if !reflect.DeepEqual(expectedReceivers, actualReceivers) {
		t.Errorf("Expected %v, but got %v", expectedReceivers, actualReceivers)
	}
}
