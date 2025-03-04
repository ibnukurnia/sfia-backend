package handlers

import (
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
