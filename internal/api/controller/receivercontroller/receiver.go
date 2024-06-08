package receivercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// @Summary Get a receiver by UUID
// @Description Get details of a receiver with the given UUID
// @Tags receiver
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 {object} dto.ReceiverOutputDto
// @Failure 404 {object} object{error=string}
// @Router /receiver/{uuid} [get]
func (r *Receiver) get(ctx *gin.Context) {

	r.logger.Info("Get By UUID Receivers")

	param := ctx.Param("uuid")
	uid, err := uuid.Parse(param)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, ginx.ErrorResponse("uuid invalid"))
		return
	}

	receiver, err := r.service.GetByID(uid)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, ginx.ErrorResponse(errors.FailedToGet("Receiver")))
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

	ctx.JSON(http.StatusOK, ginx.SuccessResponse(response))
}

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

// @Summary Update a receiver
// @Description Update a receiver with the given UUID
// @Tags receiver
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Param receiver body dto.UpdateReceiverInputDto true "Receiver"
// @Success 204
// @Failure 400 {object} object{error=string}
// @Failure 404 {object} object{error=string}
// @Router /receiver/{uuid} [patch]
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

// @Summary Bulk Delete receiver
// @Description Bulk delete array of uuids receivers
// @Tags receiver
// @Accept json
// @Produce json
// @Param receiver body dto.DeleteReceiverInputDto true "Receiver"
// @Success 200 {object} object{message=string} "Ok"
// @Failure 400 {object} object{error=string}
// @Failure 404 {object} object{error=string}
// @Router /bulk-delete/receiver [put]
func (r *Receiver) bulkDelete(ctx *gin.Context) {
	var req dto.DeleteReceiverInputDto
	r.logger.Info("Bulk Delete UUIDs Receiver")

	err := ginx.ParseJSON(ctx, &req)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, ginx.ErrorResponse(err.Error()))
		return
	}

	err = r.service.BulkDelete(req.Uuids)
	if err != nil {
		r.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, ginx.ErrorResponse(errors.FailedToDelete("Receiver")))
		return
	}

	ctx.JSON(http.StatusOK, ginx.SuccessResponse("Ok"))
}

// @Summary delete a receiver by UUID
// @Description delete with the given ID
// @Tags receiver
// @Accept json
// @Produce json
// @Param uuid path string true "UUID"
// @Success 200 {object} dto.ReceiverOutputDto
// @Failure 404 {object} object{error=string}
// @Router /receiver/{uuid} [delete]
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
