package handlers

import (
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
