package handlers

import (
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type departmentHandler struct {
	deparmentService *services.DepartmentService
}

func (handler *departmentHandler) GetDepartments(ctx *gin.Context) {
	request := requests.GetDepartmentsRequest{}

	if err := ctx.ShouldBindQuery(&request); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(err))
		return
	}

	departments, apiErr := handler.deparmentService.GetDepartments(request)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, departments, "Success Get Departments", 200)
}

func (handler *departmentHandler) AddDepartment(ctx *gin.Context) {
	request := requests.AddDepartment{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	apiErr := handler.deparmentService.AddDepartment(request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "Success Add Department", 200)
}

func (handler *departmentHandler) UpdateDepartment(ctx *gin.Context) {
	request := requests.UpdateDepartment{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	apiErr := handler.deparmentService.UpdateDepartment(request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "Success Update Department", 200)
}

func (handler *departmentHandler) DeleteDepartment(ctx *gin.Context) {
	departmentId := ctx.Param("id")

	apiErr := handler.deparmentService.DeleteDepartment(departmentId)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "Success Delete Department", 200)
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

func (handler *departmentHandler) AddDepartmentTeam(ctx *gin.Context) {
	id := ctx.Param("id")
	request := requests.AddDepartmentTeam{}
	request.DepartmendId = uuid.MustParse(id)

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	apiErr := handler.deparmentService.AddDepartmentTeam(request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "Success Add Department Team", 200)
}

func (handler *departmentHandler) UpdateDepartmentTeam(ctx *gin.Context) {
	id := ctx.Param("id")
	request := requests.UpdateDepartmentTeam{}
	request.DepartmendId = uuid.MustParse(id)

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	apiErr := handler.deparmentService.UpdateDepartmentTeam(request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "Success Update Department Team", 200)
}

func (handler *departmentHandler) DeleteDepartmentTeam(ctx *gin.Context) {
	teamId := ctx.Param("teamId")

	apiErr := handler.deparmentService.DeleteDepartmentTeam(teamId)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "Success Delete Department Team", 200)
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

func (handler *departmentHandler) AddDepartmentUnit(ctx *gin.Context) {
	id := ctx.Param("id")
	request := requests.AddDepartmentUnit{}
	request.DepartmendId = uuid.MustParse(id)

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	apiErr := handler.deparmentService.AddDepartmentUnit(request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "Success Add Department Unit", 200)
}

func (handler *departmentHandler) UpdateDepartmentUnit(ctx *gin.Context) {
	id := ctx.Param("id")
	request := requests.UpdateDepartmentUnit{}
	request.DepartmendId = uuid.MustParse(id)

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	apiErr := handler.deparmentService.UpdateDepartmentUnit(request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "Success Update Department Unit", 200)
}

func (handler *departmentHandler) DeleteDepartmentUnit(ctx *gin.Context) {
	unitId := ctx.Param("unitId")

	apiErr := handler.deparmentService.DeleteDepartmentUnit(unitId)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "Success Delete Department Unit", 200)
}
