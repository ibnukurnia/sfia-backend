package handlers

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type toolsHandler struct {
	toolsService *services.ToolsMasterService
}

func (handler *toolsHandler) GetToolsList(ctx *gin.Context) {
	tools, err := handler.toolsService.GetToolList()
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, tools, "success get tools list", 200)
}

func (handler *toolsHandler) AddTool(ctx *gin.Context) {
	request := requests.AddToolsRequest{}	

	fileHeader, err := ctx.FormFile("file")

	if fileHeader == nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("file is required")))
		return
	}

	if err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("file error")))
		return
	}

	
	request.Name = ctx.PostForm("name")
	if err := requests.NewValidationRaw().Struct(request); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("required name")))
		return
	}

	if err := handler.toolsService.AddTools(request, fileHeader); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add tools", 200)
}

func (handler *toolsHandler) UpdateTools(ctx *gin.Context) {
	request := requests.UpdateToolsRequest{}

	fileHeader, err := ctx.FormFile("file")

	if fileHeader == nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("file is required")))
		return
	}

	if err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("file error")))
		return
	}

	
	request.Name = ctx.PostForm("name")
	request.ToolsId = ctx.PostForm("id")

	if err := requests.NewValidationRaw().Struct(request); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("required name")))
		return
	}

	if err := handler.toolsService.UpdateTools(request, fileHeader); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update tools", 200)
}

func (handler *toolsHandler) DeleteTools(ctx *gin.Context) {
	trainingId := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(trainingId, "required,uuid"); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("invalid training id")))
		return
	}

	if err := handler.toolsService.DeleteTools(trainingId); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete tools", 200)
}