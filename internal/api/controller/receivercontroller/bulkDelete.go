package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moura95/transferapi/internal/dto"
	"github.com/moura95/transferapi/pkg/errors"
	"github.com/moura95/transferapi/pkg/ginx"
	httpRes "github.com/moura95/transferapi/pkg/response"
)

func (r *Receiver) bulkDelete(ctx *gin.Context) {
	var req dto.DeleteReceiverInputDto
	r.logger.Info("Bulk Delete UUIDs Receiver")

	err := ginx.ParseJSON(ctx, &req)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, httpRes.ErrorResponse(err.Error()))
		return
	}

	err = r.service.BulkDelete(req.Uuids)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, httpRes.ErrorResponse(errors.FailedToDelete("Receiver")))
		return
	}

	ctx.JSON(http.StatusOK, httpRes.SuccessResponse("Ok"))
}
