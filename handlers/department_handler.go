package handlers

import (
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type departmentHandler struct {
	deparmentService *services.DepartmentService
}

func (handler *departmentHandler) GetDepartments(ctx *gin.Context) {
	departments, apiErr := handler.deparmentService.GetDepartments()

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, departments, "Success Get Departments", 200)
}

func (handler *departmentHandler) GetDepartmentTeams(ctx *gin.Context) {
	departmentId := ctx.Param("id")

	teams, apiErr := handler.deparmentService.GetDepartmentTeams(departmentId)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, teams, "Success Get Department Teams", 200)
}

func (handler *departmentHandler) GetDepartmentUnits(ctx *gin.Context) {
	departmentId := ctx.Param("id")

	units, apiErr := handler.deparmentService.GetDepartmentUnits(departmentId)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, units, "Success Get Department Roles", 200)
}
