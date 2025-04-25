package handlers

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type corporateTitleHandler struct {
	corporateTitleService *services.CorporateTitleService
}

func (handler *corporateTitleHandler) GetCorporateTitles(ctx *gin.Context) {
	request := requests.GetCorporateTitlesRequest{}

	if err := ctx.ShouldBindQuery(&request); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(err))
		return
	}

	corporateTitles, err := handler.corporateTitleService.GetCorporateTitles(request)
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, corporateTitles, "success get corporateTitles", 200)
}

func (handler *corporateTitleHandler) AddCorporateTitle(ctx *gin.Context) {
	request := requests.AddCorporateTitleRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.corporateTitleService.AddCorporateTitle(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add corporateTitle", 200)
}

func (handler *corporateTitleHandler) UpdateCorporateTitle(ctx *gin.Context) {
	request := requests.UpdateCorporateTitleRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.corporateTitleService.UpdateCorporateTitle(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update corporateTitle", 200)
}

func (handler *corporateTitleHandler) DeleteCorporateTitle(ctx *gin.Context) {
	corporateTitleID := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(corporateTitleID, "required,uuid"); err != nil {
		err := dto.BadRequestError(errors.New("corporateTitle id is required"))
		responses.ResponseError(ctx, err)
		return
	}

	if err := handler.corporateTitleService.DeleteCorporateTitle(corporateTitleID); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete corporateTitle", 200)
}
