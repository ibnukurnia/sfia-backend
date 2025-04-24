package management

import (
	"github.com/gin-gonic/gin"
	"sv-sfia/dto/responses"
	response_management "sv-sfia/dto/responses/management"
)

type ManagementTalentHandler struct {
}

func (handler *ManagementTalentHandler) GetStatusTalent(ctx *gin.Context) {
	res := response_management.DummyStatusTalentResponse()
	responses.WriteApiResponse(ctx, res, "success get status talent", 200)
}

func (handler *ManagementTalentHandler) GetDepartmentTalent(ctx *gin.Context) {
	res := response_management.DummyDepartementTalentResponse()
	responses.WriteApiResponse(ctx, res, "success get department data talent", 200)
}

func (handler *ManagementTalentHandler) GetFunctionTalent(ctx *gin.Context) {
	res := response_management.DummyFunctionResponse()
	responses.WriteApiResponse(ctx, res, "success get department data talent", 200)
}

func (handler *ManagementTalentHandler) GetTeamTalent(ctx *gin.Context) {
	res := response_management.DummyTeamResponse()
	responses.WriteApiResponse(ctx, res, "success get department data talent", 200)
}

func (handler *ManagementTalentHandler) GetCorporateTalent(ctx *gin.Context) {
	res := response_management.DummyCorporateTitleResponse()
	responses.WriteApiResponse(ctx, res, "success get department data talent", 200)
}

func (handler *ManagementTalentHandler) GetSpecializationTalent(ctx *gin.Context) {
	res := response_management.DummySpecializationResponse()
	responses.WriteApiResponse(ctx, res, "success get specialization data talent", 200)
}

func (handler *ManagementTalentHandler) GetYearOfExperienceTalent(ctx *gin.Context) {
	res := response_management.DummyYearOfExperienceResponse()
	responses.WriteApiResponse(ctx, res, "success get YoE data talent", 200)

}
