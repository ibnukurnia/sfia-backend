package management

import (
	"github.com/gin-gonic/gin"
	"sv-sfia/dto/responses"
	response_management "sv-sfia/dto/responses/management"
)

type ManagementResumeHandler struct {
}

func (handler *ManagementResumeHandler) GetRoleDataManagement(ctx *gin.Context) {
	res := response_management.DummyRoleDataResponse()
	responses.WriteApiResponse(ctx, res, "success get role data", 200)
}

func (handler *ManagementResumeHandler) GetSkillExistingRoleDataManagement(ctx *gin.Context) {
	res := response_management.DummySkillExistingRoleDataResponse()
	responses.WriteApiResponse(ctx, res, "success get skill role data", 200)
}

func (handler *ManagementResumeHandler) GetKesulitanDUJDataManagement(ctx *gin.Context) {
	res := response_management.DummyKesulitanDUJDataResponse()
	responses.WriteApiResponse(ctx, res, "success get skill role data", 200)
}
