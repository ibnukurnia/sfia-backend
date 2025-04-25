package handlers

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type userAdminHandler struct {
	userAdminService *services.UserAdminService
}

func (handler *userAdminHandler) GetUserAdmin(ctx *gin.Context) {
	var req requests.UserAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(err))
		return
	}

	result, err := handler.userAdminService.GetUserAdmin(req)
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, result, "success get user admin list", 200)
}

func (handler *userAdminHandler) UpdateUserRole(ctx *gin.Context) {
	request := requests.UpdateUserAdminRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.userAdminService.UpdateRoleAccess(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update user role", 200)
}

func (handler *userAdminHandler) DeleteUserAdmin(ctx *gin.Context) {
	userID := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(userID, "required,uuid"); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("invalid user id")))
		return
	}

	if err := handler.userAdminService.DeleteUserAdmin(userID); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete user", 200)
}
