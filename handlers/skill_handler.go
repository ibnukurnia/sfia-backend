package handlers

import (
	"fmt"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type skillHandler struct {
	skillService *services.SkillService
	roleSevice   *services.RoleService
}

func (handler skillHandler) GetSkills(ctx *gin.Context) {
	roleIds := ctx.QueryArray("role_id[]")

	if len(roleIds) < 1 {
		responses.ResponseError(ctx, &dto.ApiError{
			Err:          fmt.Errorf("role_id is required"),
			ErrorMessage: "role_id is required",
			Typ:          dto.ErrorBadData,
		})
		return
	}

	res, apiErr := handler.roleSevice.GetRoleSkills(roleIds)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, res, "success", 200)
}

func (handler skillHandler) GetSkillsetList(ctx *gin.Context) {
	res, apiErr := handler.skillService.GetSkillsetList()
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, res, "success", 200)
}

func (handler skillHandler) AddSkillset(ctx *gin.Context) {
	request := requests.AddSkillsetRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.skillService.AddSkillSet(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add skillset", 200)
}

func (handler skillHandler) UpdateSkillset(ctx *gin.Context) {
	request := requests.UpdateSkillsetRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.skillService.UpdateSkillSet(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update skillset", 200)
}

func (handler skillHandler) DeleteSkillset(ctx *gin.Context) {
	skillsetId := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(skillsetId, "required,uuid"); err != nil {
		err := dto.BadRequestError(err)
		responses.ResponseError(ctx, err)
		return
	}

	if err := handler.skillService.DeleteSkillSet(skillsetId); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete skillset", 200)
}
