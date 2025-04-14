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

type DujService struct {
	db *gorm.DB
}

func newDujService(db *gorm.DB) *DujService {
	return &DujService{
		db: db,
	}
}

func (service DujService) GetParticipantDuj(participantId uuid.UUID, assessmentId uuid.UUID) (responses.DujAssessmentResponse, *dto.ApiError) {

	participantDepartment := models.ParticipantDepartment{}

	err := service.db.Where("assessment_id = ?", assessmentId).
		Find(&participantDepartment).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("participant department not found", zap.Error(err))

			return responses.DujAssessmentResponse{}, &dto.ApiError{
				Err:          err,
				ErrorMessage: "Belum mengisi departemen",
				Typ:          dto.ErrorForbidden,
			}
		}

		zap.L().Error("error find participant department", zap.Error(err))

		return responses.DujAssessmentResponse{}, &dto.ApiError{
			Err:          err,
			ErrorMessage: "internal server error",
			Typ:          dto.ErrorInternal,
		}
	}

	dujs := []models.Duj{}

	err = service.db.Where("department_unit_id = ?", participantDepartment.DepartmentUnitId).
		Find(&dujs).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("error find duj", zap.Error(err))

		return responses.DujAssessmentResponse{}, &dto.ApiError{
			Err:          err,
			ErrorMessage: "internal server error",
			Typ:          dto.ErrorInternal,
		}
	}

	return responses.NewDujAssessmentResponse(dujs), nil
}

func (service DujService) StoreParticipantDuj(participantId, assessmentId uuid.UUID, req requests.DujAssessmentRequest) *dto.ApiError {
	answers := []models.DujAnswer{}

	for _, job := range req.Jobs {
		jobId, err := utils.ParseUUid(job.Id)
		if err != nil {
			return err
		}

		answer := models.DujAnswer{
			ParticipantId: participantId,
			AssessmentId:  assessmentId,
			JobId:         jobId,
			CurrentJob:    job.CurrentJob,
			HaveTrouble:   job.HaveTrouble,
			TroubleCause:  job.TroubleCause,
		}

		answers = append(answers, answer)
	}

	tx := service.db.Begin()

	if len(answers) > 0 {
		if err := tx.Create(&answers).Error; err != nil {
			tx.Rollback()

			zap.L().Error("error save duj answers: ", zap.Error(err))

			return dto.InternalError(err)
		}

		if err := tx.Model(models.Assessment{}).Where("uuid = ?", assessmentId).
			Update("status", models.TOOL).Error; err != nil {
			tx.Rollback()

			zap.L().Error("error save duj answers: ", zap.Error(err))

			return dto.InternalError(err)
		}
	}

	tx.Commit()

	return nil
}
