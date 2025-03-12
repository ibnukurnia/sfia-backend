package handlers

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type dujAdminHandler struct {
	dujAdminService *services.DujAdminService
}

func (handler *dujAdminHandler) GetDujAdminList(ctx *gin.Context) {
	tresholds, err := handler.dujAdminService.GetAdminDuj()
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, tresholds, "success get duj list", 200)
}

func (handler *dujAdminHandler) AddDujAdmin(ctx *gin.Context) {
	request := requests.AddAdminDujRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.dujAdminService.AddAdminDuj(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add duj", 200)
}

func (handler *dujAdminHandler) UpdateDujAdmin(ctx *gin.Context) {
	request := requests.UpdateAdminDujRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.dujAdminService.UpdateAdminDuj(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update duj", 200)
}

func (handler *dujAdminHandler) DeleteDujAdmin(ctx *gin.Context) {
	dujID := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(dujID, "required,uuid"); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("invalid duj id")))
		return
	}

	if err := handler.dujAdminService.DeleteAdminDuj(dujID); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete duj", 200)
}