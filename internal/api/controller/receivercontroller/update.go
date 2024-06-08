package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/transferapi/internal/dto"
	"github.com/moura95/transferapi/pkg/ginx"
	httpRes "github.com/moura95/transferapi/pkg/response"
)

func (r *Receiver) update(ctx *gin.Context) {
	var req dto.UpdateReceiverInputDto
	r.logger.Info("Update Receiver")

	err := ginx.ParseJSON(ctx, &req)
	if err != nil {
		r.logger.Info("Bad Request %s", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	param := ctx.Param("uuid")
	uid, err := uuid.Parse(param)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, httpRes.ErrorResponse("uuid invalid"))
		return
	}

	err = r.service.Update(uid, req.Name, req.PixKeyType, req.PixKey, req.Email, req.CpfCnpj)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, httpRes.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
