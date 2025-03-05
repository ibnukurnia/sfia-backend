package handlers

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type roleHandler struct {
	roleService *services.RoleService
}

func (handler *roleHandler) GetRoles(ctx *gin.Context) {
	roles, err := handler.roleService.GetRoles()
	if err != nil {
		responses.ResponseError(ctx, err)

		return
	}

	responses.WriteApiResponse(ctx, roles, "success get roles", 200)
}

func (handler *roleHandler) GetRoleList(ctx *gin.Context) {
	roles, err := handler.roleService.GetRoleList()
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, roles, "success get roles", 200)
}

func (handler *roleHandler) AddRole(ctx *gin.Context) {
	request := requests.AddRoleRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.roleService.AddRole(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add role", 200)
}

func (handler *roleHandler) UpdateRole(ctx *gin.Context) {
	request := requests.UpdateRoleRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.roleService.UpdateRole(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update role", 200)
}

func (handler *roleHandler) DeleteRole(ctx *gin.Context) {
	roleID := ctx.Param("id")
	
	if err := requests.NewValidationRaw().Var(roleID, "required,uuid"); err != nil {
		err := dto.BadRequestError(errors.New("role id is required"))
		responses.ResponseError(ctx, err)
		return
	}

	if err := handler.roleService.DeleteRole(roleID); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete role", 200)
}