package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/transferapi/pkg/errors"
	"github.com/moura95/transferapi/pkg/ginx"
)

// @Summary delete a receiver by UUID
// @Description delete with the given ID
// @Tags receiver
// @Accept json
// @Produce json
// @Param id path int true "UUID"
// @Success 200 {object} dto.ReceiverOutputDto
// @Failure 404 {object} object{error=string}
// @Router /receiver/{id} [delete]
func (r *Receiver) hardDelete(ctx *gin.Context) {

	r.logger.Info("Delete UUID Receiver")

	param := ctx.Param("uuid")
	uid, err := uuid.Parse(param)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, ginx.ErrorResponse("uuid invalid"))
		return
	}

	err = r.service.Delete(uid)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, ginx.ErrorResponse(errors.FailedToDelete("Receiver")))
		return
	}

	ctx.JSON(http.StatusOK, ginx.SuccessResponse("Ok"))
}
