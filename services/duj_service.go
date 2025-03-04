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

func (service DujService) GetParticipantDuj(participantId uuid.UUID) (responses.DujAssessmentResponse, *dto.ApiError) {
	answers := []models.DujAnswer{}

	isExist := service.db.Where("participant_id = ?", participantId).
		Find(&answers).RowsAffected > 0

	if isExist {
		return responses.NewDujAssessmentResponseCurrentAnswer(answers), nil
	}

	participantDepartment := models.ParticipantDepartment{}

	err := service.db.Where("participant_id = ?", participantId).
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

func (service DujService) StoreParticipantDuj(participantId uuid.UUID, req requests.DujAssessmentRequest) {
	answers := []models.DujAnswer{}

	for _, job := range req.Jobs {

		answer := models.DujAnswer{
			ParticipantId: participantId,
			Job:           job.Name,
			Detail:        "",
			CurrentJob:    job.CurrentJob,
			HaveTrouble:   job.HaveTrouble,
			TroubleCause:  job.TroubleCause,
		}

		if job.Id != nil {
			id, _ := utils.ParseUUid(*job.Id)

			service.db.Model(&models.DujAnswer{}).Where("uuid = ?", id).
				Updates(&answer)

			continue
		}

		answers = append(answers, answer)
	}

	if len(answers) > 0 {
		service.db.Create(&answers)
	}
}
