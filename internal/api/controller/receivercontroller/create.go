package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moura95/transferapi/internal/dto"
	"github.com/moura95/transferapi/pkg/ginx"
)

// @Summary Add a new receiver
// @Description Add a new receiver with the provided information
// @Tags receiver
// @Accept json
// @Produce json
// @Param receiver body dto.CreateReceiverInputDto true "Receiver"
// @Success 201 {object} object{message=string} "Ok"
// @Failure 400 {object} object{error=string}
// @Router /receiver [post]
func (r *Receiver) create(ctx *gin.Context) {
	var req dto.CreateReceiverInputDto
	r.logger.Info("Create Receiver")

	err := ginx.ParseJSON(ctx, &req)
	if err != nil {
		r.logger.Info("Bad Request %s", err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = r.service.Create(req.Name, req.PixKeyType, req.PixKey, req.Email, req.CpfCnpj)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, ginx.ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, ginx.SuccessResponse("Ok"))
}
