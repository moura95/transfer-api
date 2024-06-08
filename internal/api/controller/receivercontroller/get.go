package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/transferapi/internal/dto"
	"github.com/moura95/transferapi/pkg/errors"
	httpRes "github.com/moura95/transferapi/pkg/response"
)

func (r *Receiver) get(ctx *gin.Context) {

	r.logger.Info("Get By UUID Receivers")

	param := ctx.Param("uuid")
	uid, err := uuid.Parse(param)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, httpRes.ErrorResponse("uuid invalid"))
		return
	}

	receiver, err := r.service.GetByID(uid)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, httpRes.ErrorResponse(errors.FailedToGet("Receiver")))
		return
	}

	response := dto.ReceiverOutputDto{
		Uuid:       receiver.Uuid,
		Name:       receiver.Name,
		CpfCnpj:    receiver.CpfCnpj,
		PixKeyType: receiver.PixKeyType,
		PixKey:     receiver.PixKey,
		Email:      receiver.Email,
		Status:     receiver.Status,
	}

	ctx.JSON(http.StatusOK, httpRes.SuccessResponse(response))
}
