package services

import (
	"errors"
	"strings"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TrainingMasterService struct {
	db *gorm.DB
}

func newTrainingMasterService(db *gorm.DB) *TrainingMasterService {
	return &TrainingMasterService{db: db}
}

func (service TrainingMasterService) GetTrainingMaster() ([]responses.TrainingMasterResponse, *dto.ApiError) {
	trainings := []models.TrainingMaster{}

	err := service.db.
		Order("created_at asc").
		Find(&trainings).Error

	if err != nil {
		zap.L().Error("error query training list", zap.Error(err))
		return []responses.TrainingMasterResponse{}, dto.InternalError(err)
	}

	results := []responses.TrainingMasterResponse{}

	for _, training := range trainings {
		results = append(results, responses.TrainingMasterResponse{
			Id:       training.Base.Uuid.String(),
			Name:     training.Name,
			Code:     training.Code,
			Jenjang:  training.Jenjang,
			SkillId:  training.SkillsId.String(),
			Level:    training.Level,
			Type:     training.Type,
			Mode:     training.Mode,
			Provider: strings.Split(training.Provider, ","),
			Silabus:  training.Silabus,
		})
	}

	return results, nil
}

func (service TrainingMasterService) AddTrainingMaster(req requests.AddTrainingMasterRequest) *dto.ApiError {
	training := models.TrainingMaster{
		Name:     req.Name,
		Code:     req.Code,
		Jenjang:  req.Jenjang,
		SkillsId:  req.SkillId,
		Level:    req.Level,
		Type:     req.Type,
		Mode:     req.Mode,
		Provider: strings.Join(req.Provider, ","),
		Silabus:  req.Silabus,
	}

	err := service.db.Create(&training).Error
	if err != nil {
		zap.L().Error("error create training", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service TrainingMasterService) UpdateTrainingMaster(req requests.UpdateTrainingMasterRequest) *dto.ApiError {
	training := models.TrainingMaster{}

	err := service.db.
		Where("uuid = ?", req.TrainingId).
		First(&training).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("training master not found", zap.String("uuid", req.TrainingId.String()))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error get training master", zap.Error(err))
		return dto.InternalError(err)
	}

	training.Name = req.Name
	training.Code = req.Code
	training.Jenjang = req.Jenjang
	training.SkillsId = req.SkillId
	training.Level = req.Level
	training.Type = req.Type
	training.Mode = req.Mode
	training.Provider = strings.Join(req.Provider, ",")
	training.Silabus = req.Silabus

	err = service.db.Save(&training).Error
	if err != nil {
		zap.L().Error("error update training", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service TrainingMasterService) DeleteTrainingMaster(trainingId string) *dto.ApiError {
	training := models.TrainingMaster{}

	err := service.db.
		Where("uuid = ?", trainingId).
		First(&training).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Warn("training master not found", zap.String("uuid", trainingId))
			return dto.NotFoundError(err)
		}
	
		if err != nil {
			zap.L().Error("error get training master", zap.Error(err))
			return dto.InternalError(err)
		}

	err = service.db.Delete(&training).Error
	if err != nil {
		zap.L().Error("error delete training", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

