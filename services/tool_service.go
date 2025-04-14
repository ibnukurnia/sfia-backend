package services

import (
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"
	"sv-sfia/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ToolService struct {
	db *gorm.DB
}

func newToolService(db *gorm.DB) *ToolService {
	return &ToolService{
		db: db,
	}
}

func (service ToolService) GetParticipantTool(participantId uuid.UUID) ([]responses.ParticipantToolResponse, *dto.ApiError) {
	tools := []models.ParticipantTool{}

	err := service.db.Where("participant_id = ?", participantId).
		Find(&tools).Error

	if err != nil {
		zap.L().Error("error querying participant tools", zap.Error(err))

		return nil, &dto.ApiError{
			ErrorMessage: "internal server error",
			Typ:          dto.ErrorInternal,
			Err:          err,
		}
	}

	res := []responses.ParticipantToolResponse{}

	for _, tool := range tools {
		res = append(res, responses.ParticipantToolResponse{
			Tool:  tool.Name,
			Level: tool.Level,
		})
	}

	return res, nil
}

func (service ToolService) CreateParticipantTool(participantId, assessmentId uuid.UUID, request requests.CreateParticipantToolRequest) *dto.ApiError {
	// err := service.db.Unscoped().
	// 	Where("participant_id = ?", participantId).
	// 	Delete(&models.ParticipantTool{}).Error

	// if err != nil {
	// 	zap.L().Error("error saving participant tools", zap.Error(err))

	// 	return dto.InternalError(err)
	// }

	tools := []models.ParticipantTool{}

	for _, tool := range request.Tools {
		t := models.Tool{}

		if tool.Id == nil {
			t.Name = tool.Name

			service.db.Create(&t)
		} else {
			toolId, err := utils.ParseUUid(*tool.Id)
			if err != nil {
				return err
			}

			t.Uuid = toolId
		}

		tools = append(tools, models.ParticipantTool{
			ParticipantId: participantId,
			AssessmentId:  assessmentId,
			ToolId:        t.Uuid,
			Name:          "",
			Level:         tool.Level,
		})
	}

	if len(tools) > 0 {
		err := service.db.Create(&tools).Error

		if err != nil {
			zap.L().Error("error saving participant tools", zap.Error(err))

			return &dto.ApiError{
				Typ:          dto.ErrorInternal,
				ErrorMessage: "internal server error",
				Err:          err,
			}
		}
	}

	return nil
}

func (service ToolService) GetTools() ([]responses.ToolResponse, *dto.ApiError) {
	tools := []models.Tool{}

	err := service.db.Find(&tools).Error
	if err != nil {
		zap.L().Error("error query tools: ", zap.Error(err))

		apiErr := dto.InternalError(err)

		if err == gorm.ErrRecordNotFound {
			apiErr = dto.NotFoundError(err)
		}

		return nil, apiErr
	}

	res := []responses.ToolResponse{}

	for _, tool := range tools {
		res = append(res, responses.ToolResponse{
			Id:   tool.Uuid.String(),
			Name: tool.Name,
		})
	}

	return res, nil
}

func (service ToolService) GetToolsAssessment(assessmentId uuid.UUID) ([]responses.ParticipantToolResponse, *dto.ApiError) {
	tools := []models.ParticipantTool{}

	err := service.db.Preload("Tool").Where("assessment_id = ?", assessmentId).
		Find(&tools).Error

	if err != nil {
		zap.L().Error("error querying participant tools", zap.Error(err))

		return nil, &dto.ApiError{
			ErrorMessage: "internal server error",
			Typ:          dto.ErrorInternal,
			Err:          err,
		}
	}

	res := []responses.ParticipantToolResponse{}

	for _, tool := range tools {
		res = append(res, responses.ParticipantToolResponse{
			Tool:     tool.Tool.Name,
			Level:    tool.Level,
			Evidence: tool.Evidence,
		})
	}

	return res, nil
}
