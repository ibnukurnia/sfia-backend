package handlers

import (
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/services"
	"sv-sfia/utils"

	"github.com/gin-gonic/gin"
)

type participantHandler struct {
	participantService *services.ParticipantService
	roleService        *services.RoleService
	skillService       *services.SkillService
	toolService        *services.ToolService
	trainingService    *services.TrainingService
}

func (handler participantHandler) Login(ctx *gin.Context) {
	request := requests.LoginRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	res, err := handler.participantService.Login(request)
	if err != nil {
		responses.ResponseError(ctx, err)

		return
	}

	responses.WriteApiResponse(ctx, res, "Loged in", 200)
}

func (handler participantHandler) Register(ctx *gin.Context) {
	request := requests.RegisterRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	res, err := handler.participantService.Register(request)
	if err != nil {
		responses.ResponseError(ctx, err)

		return
	}

	responses.WriteApiResponse(ctx, res, "Success Register", 201)
}

func (handler participantHandler) Profile(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	res, err := handler.participantService.GetPersonalInformation(participantId)
	if err != nil {
		responses.ResponseError(ctx, err)

		return
	}

	responses.WriteApiResponse(ctx, res, "success get profile", 200)
}

func (handler participantHandler) StorePersonalInformation(ctx *gin.Context) {
	request := requests.PersonalInformationRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	participantId := utils.GetUserIdFromContext(ctx)

	err := handler.participantService.StorePersonalInformation(participantId, request)
	if err != nil {
		responses.ResponseError(ctx, err)

		return
	}

	responses.WriteApiResponse(ctx, nil, "ok", 200)
}

func (handler participantHandler) CreateParticipantRole(ctx *gin.Context) {
	request := requests.CreateParticipantRoleRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	userId := utils.GetUserIdFromContext(ctx)

	apiErr := handler.participantService.StoreParticipantRole(userId, request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success creating participant role", 200)
}

func (handler participantHandler) GetParticipantRoleSkills(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	roles, apiErr := handler.roleService.GetParticipantRoles(participantId)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	skills, apiErr := handler.skillService.FindSkillByRoleIds(participantId, roles)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, skills, "success get role skills", 200)
}

func (handler participantHandler) StoreParticipantSkill(ctx *gin.Context) {
	request := requests.StoreParticipantSkillRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	userId := utils.GetUserIdFromContext(ctx)

	apiErr := handler.participantService.StoreParticipantSkill(userId, request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success creating participant skill", 200)
}

func (handler participantHandler) CreateParticipantTool(ctx *gin.Context) {
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

	apiErr = handler.toolService.CreateParticipantTool(participantId, assessmentId, request)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success create participant tools", 201)
}

func (handler participantHandler) GetParticipantTool(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	tools, apiErr := handler.toolService.GetParticipantTool(participantId)

	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, tools, "success get participant tools", 200)
}

func (handler participantHandler) GetParticipantRoleTraining(ctx *gin.Context) {
	participantId := utils.GetUserIdFromContext(ctx)

	trainings, apiErr := handler.roleService.GetRoleTraining(participantId)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, trainings, "success get role trainings", 200)
}

func (handler participantHandler) CreateParticipantTraining(ctx *gin.Context) {
	request := requests.CreateParticipantTrainingRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	participantId := utils.GetUserIdFromContext(ctx)
	assessmentId, apiErr := utils.ParseUUid(ctx.Param("id"))
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	apiErr = handler.trainingService.CreateParticipantTraining(participantId, assessmentId, request)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success storing trainings", 201)
}

func (handler participantHandler) CreateParticipantUpdatedTraining(ctx *gin.Context) {
	request := requests.CreateParticipantUpdatedTrainingRequest{}

	if !requests.NewValidation(ctx).Validate(&request) {
		return
	}

	assessmentId, apiErr := utils.ParseUUid(ctx.Param("id"))
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	apiErr = handler.trainingService.CreateParticipantUpdatedTraining(assessmentId, request)
	if apiErr != nil {
		responses.ResponseError(ctx, apiErr)
		return
	}

	responses.WriteApiResponse(ctx, nil, "success storing trainings", 201)
}
