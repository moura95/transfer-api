package receivercontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/moura95/transferapi/internal/services/receiverservice"
	"go.uber.org/zap"
)

type IReceiver interface {
	SetupReceiverRoute(routers *gin.RouterGroup)
}

type Receiver struct {
	service *receiverservice.Service
	logger  *zap.SugaredLogger
}

func NewReceiverRouter(s *receiverservice.Service, log *zap.SugaredLogger) *Receiver {
	return &Receiver{
		service: s,
		logger:  log,
	}
}

func (r *Receiver) SetupReceiverRoute(routers *gin.RouterGroup) {
	routers.GET("/receiver", r.list)

}
