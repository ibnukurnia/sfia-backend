package handlers

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type tresholdHandler struct {
	tresholdService *services.TresholdService
}

func (handler *tresholdHandler) GetTresholdList(ctx *gin.Context) {
	tresholds, err := handler.tresholdService.GetTresholdList()
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, tresholds, "success get treshold list", 200)
}

func (handler *tresholdHandler) AddTreshold(ctx *gin.Context) {
	request := requests.AddTreshold{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.tresholdService.AddTreshold(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add treshold", 200)
}

func (handler *tresholdHandler) UpdateTreshold(ctx *gin.Context) {
	request := requests.UpdateTreshold{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.tresholdService.UpdateTreshold(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update treshold", 200)
}

func (handler *tresholdHandler) DeleteTreshold(ctx *gin.Context) {
	tresholdID := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(tresholdID, "required,uuid"); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("invalid treshold id")))
		return
	}

	if err := handler.tresholdService.DeleteTreshold(tresholdID); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete treshold", 200)
}

func (handler tresholdHandler) GetSkillLevelTreshold(ctx *gin.Context) {
	treshold, err := handler.tresholdService.GetSkillLevelTreshold()

	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, treshold, "success get treshold skills", 200)
}
