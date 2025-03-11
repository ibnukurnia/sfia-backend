package services

import (
	"errors"
	"fmt"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TresholdService struct {
	db *gorm.DB
}

func newTresholdService(db *gorm.DB) *ParameterService {
	return &ParameterService{db: db}
}

func (service ParameterService) GetTresholdList() (responses.TresholdResponse, *dto.ApiError) {
	tresholds := []models.Treshold{}

	err := service.db.
		Order("created_at asc").
		Find(&tresholds).Error

	if err != nil {
		zap.L().Error("error query treshold list", zap.Error(err))
		return responses.TresholdResponse{}, dto.InternalError(err)
	}

	results := responses.TresholdResponse{
		RoleLevel:  []responses.TresholdItem{},
		SkillLevel: []responses.TresholdItem{},
	}

	for _, treshold := range tresholds {
		item := responses.TresholdItem{
			Id:          treshold.Uuid.String(),
			Name:        treshold.Name,
			TresholdFrom: treshold.TresholdFrom,
			TresholdTo:   treshold.TresholdTo,
			Color:       treshold.Color,
		}
		fmt.Print(item)

		if treshold.Category == "role" {
			results.RoleLevel = append(results.RoleLevel, item)
		} else if treshold.Category == "skill" {
			results.SkillLevel = append(results.SkillLevel, item)
		}
	}

	return results, nil
}

func (service ParameterService) AddTreshold(req requests.AddTreshold) *dto.ApiError {
	treshold := models.Treshold{
		Name:        req.Name,
		Category:    req.Category,
		TresholdFrom: req.TresholdFrom,
		TresholdTo:   req.TresholdTo,
		Color:       req.Color,
	}

	err := service.db.Create(&treshold).Error

	if err != nil {
		zap.L().Error("error create treshold", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service ParameterService) UpdateTreshold(req requests.UpdateTreshold) *dto.ApiError {
	treshold := models.Treshold{}

	err := service.db.Where("uuid = ?", req.TresholdId).First(&treshold).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("treshold not found", zap.String("uuid", req.TresholdId.String()))
		return dto.NotFoundError(err)
	}
	if err != nil {
		zap.L().Error("error query treshold", zap.Error(err))
		return dto.InternalError(err)
	}

	treshold.Name = req.Name
	treshold.Category = req.Category
	treshold.TresholdFrom = req.TresholdFrom
	treshold.TresholdTo = req.TresholdTo
	treshold.Color = req.Color

	err = service.db.Save(&treshold).Error

	if err != nil {
		zap.L().Error("error update treshold", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service ParameterService) DeleteTreshold(tresholdId string) *dto.ApiError {
	treshold := models.Treshold{}

	err := service.db.Where("uuid = ?", tresholdId).First(&treshold).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("treshold not found", zap.String("uuid", tresholdId))
		return dto.NotFoundError(err)
	}
	if err != nil {
		zap.L().Error("error query treshold", zap.Error(err))
		return dto.InternalError(err)
	}

	err = service.db.Delete(&treshold).Error

	if err != nil {
		zap.L().Error("error delete treshold", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}