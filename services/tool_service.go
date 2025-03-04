package services

import (
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"

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
			Tool:  tool.Tool,
			Level: tool.Level,
		})
	}

	return res, nil
}

func (service ToolService) CreateParticipantTool(participantId uuid.UUID, request requests.CreateParticipantToolRequest) *dto.ApiError {
	err := service.db.Unscoped().
		Where("participant_id = ?", participantId).
		Delete(&models.ParticipantTool{}).Error

	if err != nil {
		zap.L().Error("error saving participant tools", zap.Error(err))

		return dto.InternalError(err)
	}

	tools := []models.ParticipantTool{}

	for _, tool := range request.Tools {
		tools = append(tools, models.ParticipantTool{
			ParticipantId: participantId,
			Tool:          tool.Tool,
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
