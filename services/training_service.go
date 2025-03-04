package services

import (
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/models"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TrainingService struct {
	db *gorm.DB
}

func newTrainingService(db *gorm.DB) *TrainingService {
	return &TrainingService{db: db}
}

func (service TrainingService) CreateParticipantTraining(participantId uuid.UUID, request requests.CreateParticipantTrainingRequest) *dto.ApiError {
	// err := service.db.
	// 	Where("participant_id = ?", participantId).
	// 	Unscoped().
	// 	Delete(&models.ParticipantTraining{}).
	// 	Error

	// if err != nil {
	// 	zap.L().Error("error saving participant training", zap.Error(err))

	// 	return &dto.ApiError{
	// 		Typ:          dto.ErrorInternal,
	// 		Err:          err,
	// 		ErrorMessage: "internal server error",
	// 	}
	// }

	trainings := []models.ParticipantTraining{}

	for _, training := range request.Trainings {
		// id, err := uuid.Parse(training.Id)
		// if err != nil {
		// 	zap.L().Error("error parsing uuid", zap.Error(err))

		// 	return &dto.ApiError{
		// 		Typ:          dto.ErrorInternal,
		// 		Err:          err,
		// 		ErrorMessage: "internal server error",
		// 	}
		// }

		trainings = append(trainings, models.ParticipantTraining{
			// TrainingId:        id,
			Name:              training.Name,
			IsNeeded:          training.IsNeeded,
			Priority:          training.Priority,
			ParticipantId:     participantId,
			NeedCertification: training.NeedCertification,
		})
	}

	if len(trainings) > 0 {
		err := service.db.Create(&trainings).Error

		if err != nil {
			zap.L().Error("error saving participant training", zap.Error(err))

			return dto.InternalError(err)
		}
	}

	return nil
}
