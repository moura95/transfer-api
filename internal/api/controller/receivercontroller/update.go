package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/transferapi/internal/dto"
	"github.com/moura95/transferapi/pkg/ginx"
)

// @Summary Update a receiver
// @Description Update a receiver with the given ID
// @Tags receiver
// @Accept json
// @Produce json
// @Param id path int true "UUID"
// @Param receiver body dto.UpdateReceiverInputDto true "Receiver"
// @Success 204
// @Failure 400 {object} object{error=string}
// @Failure 404 {object} object{error=string}
// @Router /receiver/{id} [patch]
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
		ctx.JSON(http.StatusBadRequest, ginx.ErrorResponse("uuid invalid"))
		return
	}

	err = r.service.Update(uid, req.Name, req.PixKeyType, req.PixKey, req.Email, req.CpfCnpj)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, ginx.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
