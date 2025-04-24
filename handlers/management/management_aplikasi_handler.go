package management

import (
	"github.com/gin-gonic/gin"
	"sv-sfia/dto/responses"
	response_management "sv-sfia/dto/responses/management"
)

type ManagementAplikasiHandler struct {
}

func (handler *ManagementAplikasiHandler) GetManagementAplikasi(ctx *gin.Context) {
	res := response_management.DummyManagementAplikasiResponse()
	responses.WriteApiResponse(ctx, res, "success get status talent", 200)
}
