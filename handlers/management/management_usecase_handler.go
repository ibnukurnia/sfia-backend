package management

import (
	"github.com/gin-gonic/gin"
	"sv-sfia/dto/responses"
	response_management "sv-sfia/dto/responses/management"
)

type ManagementUseCaseHandler struct{}

func (handler *ManagementUseCaseHandler) GetPersebaranTipeRole(ctx *gin.Context) {
	res := response_management.DummyPersebaranRole()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetPersebaranLevelRole(ctx *gin.Context) {
	res := response_management.DummyPersebaranRole()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetPersebaranSkill(ctx *gin.Context) {
	res := response_management.DummyPersebaranSkill()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetRelevansiTahunChart(ctx *gin.Context) {
	res := response_management.DummyRelevansiTahunLevelChartResponse()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetRelevansiTahunTables(ctx *gin.Context) {
	res := response_management.DummyRelevansiTahunLevelChartResponse()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetUnMasteredSkillChart(ctx *gin.Context) {
	res := response_management.DummyUnMasteredSkillResponse()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetUnMasteredSkillTable(ctx *gin.Context) {
	res := response_management.DummyUnMasteredSkillTablesReponse()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetYearLevelRoleMapping(ctx *gin.Context) {
	res := response_management.DummyDataRoleMappingResponse()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetYearLevelRoleDistribution(ctx *gin.Context) {
	res := response_management.DummyDataYearLevelDistribution()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetSkillRequirement(ctx *gin.Context) {
	res := response_management.DummyDataSkillRequirements()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetCrossRoleChartRecommendation(ctx *gin.Context) {
	res := response_management.DummyDataCrossRoleRecommendations()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetCrossRoleTableRecommendation(ctx *gin.Context) {
	res := response_management.DummyDataCrossRoleRecommendationTable()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetCrossSkillChartRecommendation(ctx *gin.Context) {
	res := response_management.DummyDataCrossSkillRecommendationChart()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *ManagementUseCaseHandler) GetCrossSkillTableRecommendation(ctx *gin.Context) {
	res := response_management.DummyDataCrossSkillRecommendationTable()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}
