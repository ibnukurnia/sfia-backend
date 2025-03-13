package handlers

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"

	"github.com/gin-gonic/gin"
)

type trainingMasterHandler struct {
	trainingMasterService *services.TrainingMasterService
}

func (handler *trainingMasterHandler) GetTrainingMaster(ctx *gin.Context) {
	trainings, err := handler.trainingMasterService.GetTrainingMaster()
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, trainings, "success get training master list", 200)
}

func (handler *trainingMasterHandler) AddTrainingMaster(ctx *gin.Context) {
	request := requests.AddTrainingMasterRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.trainingMasterService.AddTrainingMaster(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success add training master", 200)
}

func (handler *trainingMasterHandler) UpdateTrainingMaster(ctx *gin.Context) {
	request := requests.UpdateTrainingMasterRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	if err := handler.trainingMasterService.UpdateTrainingMaster(request); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success update training master", 200)
}

func (handler *trainingMasterHandler) DeleteTrainingMaster(ctx *gin.Context) {
	trainingId := ctx.Param("id")

	if err := requests.NewValidationRaw().Var(trainingId, "required,uuid"); err != nil {
		responses.ResponseError(ctx, dto.BadRequestError(errors.New("invalid training id")))
		return
	}

	if err := handler.trainingMasterService.DeleteTrainingMaster(trainingId); err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success delete training master", 200)
}