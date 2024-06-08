package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moura95/transferapi/pkg/errors"
	httpRes "github.com/moura95/transferapi/pkg/response"
)

func (r *Receiver) hardDelete(ctx *gin.Context) {

	r.logger.Info("Delete UUID Receiver")

	param := ctx.Param("uuid")
	uid, err := uuid.Parse(param)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, httpRes.ErrorResponse("uuid invalid"))
		return
	}

	err = r.service.Delete(uid)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, httpRes.ErrorResponse(errors.FailedToDelete("Receiver")))
		return
	}

	ctx.JSON(http.StatusOK, httpRes.SuccessResponse("Ok"))
}