package handlers

import (
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type skillHandler struct {
	skillService *services.SkillService
}

func (handler skillHandler) GetSkills(ctx *gin.Context) {
	// roleIds := []string{"48af133c-6898-4e1d-a03b-f7a35a2afb64", "4a774b40-0eed-4e87-9312-220a84f3eb82", "e3e8b7c1-e908-4cc0-a915-c0f6cb1ed5f3"}

	// res, apiErr := handler.skillService.FindSkillByRoleIds(dto.ParticipantRoleIdsDto{})
	// if apiErr != nil {
	// 	responses.ResponseError(ctx, apiErr)
	// 	return
	// }

	responses.WriteApiResponse(ctx, nil, "success", 200)
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