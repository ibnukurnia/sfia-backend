package handlers

import (
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"
	"sv-sfia/utils"

	"github.com/gin-gonic/gin"
)

type assessmentHandler struct {
	assessmentService *services.AssessmentService
	sfiaService       *services.SfiaService
	skillService      *services.SkillService
	dujService        *services.DujService
}

func (handler assessmentHandler) CreateNewAssessment(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	res, err := handler.assessmentService.CreateNewAssessment(participantId)
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, res, "success creating new assessments", 201)
}

func (handler assessmentHandler) GetSelfAssessments(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	res := handler.assessmentService.GetSelfAssessment(participantId)

	responses.WriteApiResponse(ctx, res, "success get self assessments", 200)
}

func (handler assessmentHandler) SaveSelfAssessmentAnswer(ctx *gin.Context) {
	request := requests.SelfAssessmentRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	participantId := utils.GetUserIdFromContext(ctx)

	err := handler.assessmentService.StoreSelfAssessment(participantId, request)
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "self assessment saved", 201)
}

func (handler assessmentHandler) GetDujAssesments(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	dujs, apiErr := handler.dujService.GetParticipantDuj(participantId)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, dujs, "success get duj assessments", 200)
}

func (handler assessmentHandler) SaveDujAnswer(ctx *gin.Context) {
	request := requests.DujAssessmentRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	participantId := utils.GetUserIdFromContext(ctx)

	handler.dujService.StoreParticipantDuj(participantId, request)
	// if apiErr != nil {
	// 	responses.ResponseError(ctx, apiErr)
	// 	return
	// }

	responses.WriteApiResponse(ctx, nil, "duj assessment saved", 201)
}

func (handler assessmentHandler) GetToolAssessment(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	tools := handler.assessmentService.GetToolAssessment(participantId)

	responses.WriteApiResponse(ctx, tools, "success get tools", 200)
}

func (handler assessmentHandler) SaveToolAssessmentAnswers(ctx *gin.Context) {
	request := requests.ToolAssessmentRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	participantId := utils.GetUserIdFromContext(ctx)

	apiErr := handler.assessmentService.StoreToolAssessment(participantId, request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "tools saved", 201)
}
