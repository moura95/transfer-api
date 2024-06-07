package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/moura95/transferapi/config"
	"github.com/moura95/transferapi/internal/api/controller/receivercontroller"
	receiverrepo "github.com/moura95/transferapi/internal/repository/receiversrepo"
	"github.com/moura95/transferapi/internal/services/receiverservice"
	"go.uber.org/zap"
)

func CreateRoutesV1(store *sqlx.DB, cfg *config.Config, router *gin.Engine, log *zap.SugaredLogger) {
	router.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	routes := router.Group("/")
	// Instance Receiver Repository Postgres
	receiverRepository := receiverrepo.NewReceiverRepository(store, log)
	// Instance Receiver Service with Postgres
	receiverService := receiverservice.NewReceiverService(receiverRepository, *cfg, log)

	// Init all Routers
	receivercontroller.NewReceiverRouter(receiverService, log).SetupReceiverRoute(routes)

}
