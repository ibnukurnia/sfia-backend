package handlers

import (
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService *services.AuthService
}

func (handler authHandler) Register(ctx *gin.Context) {
	request := requests.RegisterRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	res, apiErr := handler.authService.Register(request)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)

		return
	}

	responses.WriteApiResponse(ctx, res, "success create new user", 201)
}

func (handler authHandler) Login(ctx *gin.Context) {
	request := requests.LoginRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	res, apiErr := handler.authService.Login(request)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)

		return
	}

	responses.WriteApiResponse(ctx, res, "success login", 200)
}
