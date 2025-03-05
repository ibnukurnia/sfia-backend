package handlers

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type roleGroupHandler struct {
	roleGroupService *services.RoleGroupService
}

func (handler *roleGroupHandler) GetRoleGroup(ctx *gin.Context) {
	roles, err := handler.roleGroupService.GetRoleGroup()
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, roles, "success get roles", 200)
}

func (handler *roleGroupHandler) AddRoleGroup(ctx *gin.Context) {
	request := requests.AddRoleGroupRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.roleGroupService.AddRoleGroup(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add role group", 200)
}

func (handler *roleGroupHandler) UpdateRoleGroup(ctx *gin.Context) {
	request := requests.UpdateRoleGroupRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.roleGroupService.UpdateRoleGroup(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update role group", 200)
}

func (handler *roleGroupHandler) DeleteRoleGroup(ctx *gin.Context) {
	roleGorupID := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(roleGorupID, "required,uuid"); err != nil {
		err := dto.BadRequestError(errors.New("role group id is required"))
		responses.ResponseError(ctx, err)
		return
	}

	if err := handler.roleGroupService.DeleteRoleGroup(roleGorupID); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete role group", 200)
}
