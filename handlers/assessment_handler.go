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
	toolService       *services.ToolService
}

func (handler assessmentHandler) ListAssessment(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	res := handler.assessmentService.ListAssessment(participantId)

	responses.WriteApiResponse(ctx, res, "success get assessments", 200)
}

func (handler assessmentHandler) AssessmentStatus(ctx *gin.Context) {
	assessmentId := ctx.Param("id")

	id, err := utils.ParseUUid(assessmentId)
	if err != nil {
		responses.ResponseError(ctx, err)

		return
	}

	assessment, err := handler.assessmentService.GetAssessmentById(id)
	if err != nil {
		responses.ResponseError(ctx, err)

		return
	}

	responses.WriteApiResponse(ctx, responses.AssessmentResponse{
		Id:     assessment.Uuid.String(),
		Status: assessment.Status,
	}, "success get assessments", 200)
}

func (handler assessmentHandler) SfiaResult(ctx *gin.Context) {
	assessmentId, apiErr := utils.ParseUUid(ctx.Param("id"))
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	res, apiErr := handler.assessmentService.SfiaResult(assessmentId)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, res, "success get assessment result", 200)
}

func (handler assessmentHandler) CreateNewAssessment(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	request := requests.CreateAssessmentRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	res, err := handler.assessmentService.CreateNewAssessment(participantId, request)
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, res, "success creating new assessments", 201)
}

func (handler assessmentHandler) GetListSkills(ctx *gin.Context) {
	assessmentId := ctx.Param("assessment_id")

	skills, err := handler.skillService.GetParticipantSkills(assessmentId)
	if err != nil {
		responses.ResponseError(ctx, err)
	}

	responses.WriteApiResponse(ctx, skills, "success get skills", 200)
}

func (handler assessmentHandler) GetSelfAssessments(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	res := handler.assessmentService.GetSelfAssessment(participantId)

	responses.WriteApiResponse(ctx, res, "success get self assessments", 200)
}

func (handler assessmentHandler) SaveSelfAssessmentAnswer(ctx *gin.Context) {
	assessmentId := ctx.Param("id")

	request := requests.SelfAssessmentRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	err := handler.assessmentService.StoreSelfAssessment(assessmentId, request)
	if err != nil {
		responses.ResponseError(ctx, err)
		return
	}

	responses.WriteApiResponse(ctx, nil, "self assessment saved", 201)
}

func (handler assessmentHandler) GetDujAssesments(ctx *gin.Context) {
	assessmentId, apiErr := utils.ParseUUid(ctx.Param("id"))
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	participantId := utils.GetUserIdFromContext(ctx)

	dujs, apiErr := handler.dujService.GetParticipantDuj(participantId, assessmentId)

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

	assessmentId, apiErr := utils.ParseUUid(ctx.Param("id"))
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	apiErr = handler.dujService.StoreParticipantDuj(participantId, assessmentId, request)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "duj assessment saved", 201)
}

func (handler assessmentHandler) GetToolAssessment(ctx *gin.Context) {
	assessmentId, apiErr := utils.ParseUUid(ctx.Param("id"))
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	tools, apiErr := handler.toolService.GetToolsAssessment(assessmentId)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, gin.H{"tools": tools}, "success get tools", 200)
}

func (handler assessmentHandler) SaveToolAssessmentAnswers(ctx *gin.Context) {
	request := requests.CreateParticipantToolRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	participantId := utils.GetUserIdFromContext(ctx)

	assessmentId, apiErr := utils.ParseUUid(ctx.Param("id"))
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	apiErr = handler.assessmentService.StoreToolAssessment(participantId, assessmentId, request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success create participant tools", 201)
}

func (handler assessmentHandler) Resume(ctx *gin.Context) {
	assessmentId, apiErr := utils.ParseUUid(ctx.Param("id"))
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	resume, apiErr := handler.assessmentService.Resume(assessmentId)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, resume, "success get assessment resume", 200)
}
