package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moura95/transferapi/internal/dto"
	"github.com/moura95/transferapi/pkg/errors"
	"github.com/moura95/transferapi/pkg/ginx"
)

// @Summary List all receivers
// @Description Get a list of all receivers
// @Tags receivers
// @Accept json
// @Produce json
// @Success 200 {array} dto.ReceiverOutputDto
// @Router /receiver [get]
func (r *Receiver) list(c *gin.Context) {
	r.logger.Info("List All Receivers")

	filters := map[string]string{
		"status":        c.Query("status"),
		"name":          c.Query("name"),
		"pix_key_type":  c.Query("pix_key_type"),
		"pix_key_value": c.Query("pix_key_value"),
		"limit":         c.Query("limit"),
		"page":          c.Query("page"),
	}

	receivers, pageinfo, err := r.service.List(filters)

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
