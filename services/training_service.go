package services

import (
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/models"
	"time"

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

func (service TrainingService) CreateParticipantTraining(participantId, assessmentId uuid.UUID, request requests.CreateParticipantTrainingRequest) *dto.ApiError {
	trainings := []models.ParticipantTraining{}

	for _, training := range request.Trainings {
		t := models.Training{}

		if training.Id != nil {
			id, err := uuid.Parse(*training.Id)
			if err != nil {
				zap.L().Error("error parsing uuid", zap.Error(err))

				return dto.InternalError(err)
			}

			t.Uuid = id
		} else {
			roleId, err := uuid.Parse(*training.RoleId)
			if err != nil {
				zap.L().Error("error parsing uuid", zap.Error(err))

				return dto.InternalError(err)
			}

			t.RoleId = roleId
			t.Name = training.Name

			service.db.Create(&t)
		}

		trainings = append(trainings, models.ParticipantTraining{
			TrainingId:        t.Uuid,
			Name:              "",
			IsNeeded:          training.IsNeeded,
			Priority:          training.Priority,
			ParticipantId:     participantId,
			NeedCertification: training.NeedCertification,
			AssessmentId:      assessmentId,
		})
	}

	tx := service.db.Begin()

	if len(trainings) > 0 {
		err := tx.Create(&trainings).Error

		if err != nil {
			zap.L().Error("error saving participant training", zap.Error(err))

			tx.Rollback()

			return dto.InternalError(err)
		}
	}

	if err := tx.Model(models.Assessment{}).
		Where("uuid = ?", assessmentId).
		Update("status", models.DONE).
		Error; err != nil {

		zap.L().Error("error saving participant training", zap.Error(err))

		tx.Rollback()

		return dto.InternalError(err)
	}

	tx.Commit()

	return nil
}

func (service TrainingService) CreateParticipantUpdatedTraining(assessmentId uuid.UUID, request requests.CreateParticipantUpdatedTrainingRequest) *dto.ApiError {
	trainings := []models.UpdatedTraining{}
	layout := "2006-01-02"

	for _, training := range request.Trainings {
		startDate, err := time.Parse(layout, training.StartDate)
		if err != nil {
			zap.L().Error("error parsing date: ", zap.Error(err))

			return dto.InternalError(err)
		}

		endDate, err := time.Parse(layout, training.EndDate)
		if err != nil {
			zap.L().Error("error parsing date: ", zap.Error(err))

			return dto.InternalError(err)
		}

		trainings = append(trainings, models.UpdatedTraining{
			Name:             training.Name,
			HasCertification: training.HasCertification,
			GetCertification: training.GetCertification,
			Implementation:   models.TrainingImplementation(training.Implementation),
			Location:         models.TrainingLocation(training.Location),
			Provider:         training.Provider,
			StartDate:        startDate,
			EndDate:          endDate,
			AssessmentId:     assessmentId,
		})
	}

	tx := service.db.Begin()

	if len(trainings) > 0 {
		err := tx.Create(&trainings).Error

		if err != nil {
			zap.L().Error("error saving participant training", zap.Error(err))

			tx.Rollback()

			return dto.InternalError(err)
		}
	}

	err := tx.Model(models.Assessment{}).Where("uuid = ?", assessmentId).Update("status", models.UPDATETRANING).Error
	if err != nil {
		zap.L().Error("error saving participant training", zap.Error(err))

		tx.Rollback()

		return dto.InternalError(err)
	}

	tx.Commit()

	return nil
}
