package management

import (
	"github.com/gin-gonic/gin"
	"sv-sfia/dto/responses"
	response_management "sv-sfia/dto/responses/management"
)

type RoleAndSkillManagement struct {
}

func (handler *RoleAndSkillManagement) GetChartData(ctx *gin.Context) {
	res := response_management.DummyChartData()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *RoleAndSkillManagement) GetCountKomposisiData(ctx *gin.Context) {
	res := response_management.DummyCountKomposisiData()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}

func (handler *RoleAndSkillManagement) GetSampleData(ctx *gin.Context) {
	res := response_management.DummySampleData()
	responses.WriteApiResponse(ctx, res, "success get persebaran role", 200)
}
