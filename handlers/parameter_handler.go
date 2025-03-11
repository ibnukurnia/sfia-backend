package handlers

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type parameterHandler struct {
	parameterService *services.ParameterService
}

func (handler *parameterHandler) GetParameterList(ctx *gin.Context) {
	tresholds, err := handler.parameterService.GetParameterList()
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, tresholds, "success get parameter list", 200)
}

func (handler *parameterHandler) AddParameterScore(ctx *gin.Context) {
	request := requests.AddParameterScore{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.parameterService.AddParameterScore(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add parameter", 200)
}

func (handler *parameterHandler) AddParameterDifficulty(ctx *gin.Context) {
	request := requests.AddParameterDifficulty{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.parameterService.AddParameterDifficulty(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add parameter", 200)
}

func (handler *parameterHandler) UpdateParameterScore(ctx *gin.Context) {
	request := requests.UpdateParameterScore{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.parameterService.UpdateParameterScore(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update parameter", 200)
}

func (handler *parameterHandler) UpdateParameterDifficulty(ctx *gin.Context) {
	request := requests.UpdateParameterDifficulty{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.parameterService.UpdateParameterDifficulty(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update parameter", 200)
}

func (handler *parameterHandler) DeleteParameterScore(ctx *gin.Context) {
	parameterID := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(parameterID, "required,uuid"); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("invalid parameter id")))
		return
	}

	if err := handler.parameterService.DeleteParameterScore(parameterID); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete parameter", 200)
}

func (handler *parameterHandler) DeleteParameterDifficulty(ctx *gin.Context) {
	parameterID := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(parameterID, "required,uuid"); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("invalid parameter id")))
		return
	}

	if err := handler.parameterService.DeleteParameterDifficulty(parameterID); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete parameter", 200)
}