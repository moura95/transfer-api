package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moura95/transferapi/internal/dto"
	"github.com/moura95/transferapi/internal/services/receiverservice"
	"github.com/moura95/transferapi/pkg/errors"
	"github.com/moura95/transferapi/pkg/ginx"
)

type ListRequest struct {
	Status      string `form:"status"`
	Name        string `form:"name"`
	PixKeyType  string `form:"pix_key_type"`
	PixKeyValue string `form:"pix_key_value"`
	Limit       int    `form:"limit"`
	Page        int    `form:"page"`
}

// @Summary List all receivers
// @Description Get a list of all receivers
// @Tags receivers
// @Accept json
// @Produce json
// @Success 200 {array} dto.ReceiverOutputDto
// @Router /receiver [get]
func (r *Receiver) list(c *gin.Context) {
	r.logger.Info("List All Receivers")

	var filters ListRequest
	err := ginx.ParseQuery(c, &filters)
	if err != nil {
		r.logger.Error(err)
		c.JSON(http.StatusBadRequest, ginx.ErrorResponse(err.Error()))
		return
	}
	receivers, pageinfo, err := r.service.List(receiverservice.ListRequest{
		Status:      filters.Status,
		Name:        filters.Name,
		PixKeyType:  filters.PixKeyType,
		PixKeyValue: filters.PixKeyValue,
		Limit:       filters.Limit,
		Page:        filters.Page,
	})
	if err != nil {
		r.logger.Error(err)
		c.JSON(http.StatusInternalServerError, ginx.ErrorResponse(errors.FailedToList("Receivers")))
		return
	}

	var response []dto.ReceiverOutputDto
	for _, receiver := range receivers {
		response = append(response, dto.ReceiverOutputDto{
			Uuid:       receiver.Uuid,
			Name:       receiver.Name,
			CpfCnpj:    receiver.CpfCnpj,
			PixKeyType: receiver.PixKeyType,
			PixKey:     receiver.PixKey,
			Email:      receiver.Email,
			Status:     receiver.Status,
		})

	}

	c.JSON(http.StatusOK, ginx.SuccessResponseWithPageInfo(response, pageinfo))
}
