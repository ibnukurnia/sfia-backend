package handlers

import (
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type toolHandler struct {
	toolService *services.ToolService
}

func (handler toolHandler) GetTools(ctx *gin.Context) {
	res, err := handler.toolService.GetTools()
	if err != nil {
		responses.ResponseError(ctx, err)

		return
	}

	responses.WriteApiResponse(ctx, res, "success get tools", 200)
}
