package services

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ParameterService struct {
	db *gorm.DB
}

func newParameterService(db *gorm.DB) *ParameterService {
	return &ParameterService{db: db}
}

func (service ParameterService) GetParameterList() (responses.ParameterResponse, *dto.ApiError) {
	parameterScore := []models.ParameterScore{}
	parameterDifficulty := []models.ParameterDifficulty{}

	err := service.db.
		Order("created_at asc").
		Find(&parameterScore).Error

	if err != nil {
		zap.L().Error("error query parameter list", zap.Error(err))
		return responses.ParameterResponse{}, dto.InternalError(err)
	}

	err = service.db.
		Order("created_at asc").
		Find(&parameterDifficulty).Error

	if err != nil {
		zap.L().Error("error query parameter list", zap.Error(err))
		return responses.ParameterResponse{}, dto.InternalError(err)
	}

	results := responses.ParameterResponse{
		ParameterScore:  []responses.ParameterScoreItem{},
		ParameterDifficulty: []responses.ParameterDifficultyItem{},
	}

	for _, param := range parameterScore {
		item := responses.ParameterScoreItem{
			Id:          param.Uuid.String(),
			Name:        param.Name,
			Description: param.Description,
		}
		
		results.ParameterScore = append(results.ParameterScore, item)
	}

	for _, param := range parameterDifficulty {
		item := responses.ParameterDifficultyItem{
			Id:          param.Uuid.String(),
			Description: param.Description,
		}

		results.ParameterDifficulty = append(results.ParameterDifficulty, item)
	}

	return results, nil
}

func (service ParameterService) AddParameterScore(req requests.AddParameterScore) *dto.ApiError {
	parameter := models.ParameterScore{
		Name:        req.Name,
		Description: req.Description,
	}

	err := service.db.Create(&parameter).Error

	if err != nil {
		zap.L().Error("error creating parameter score", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service ParameterService) AddParameterDifficulty(req requests.AddParameterDifficulty) *dto.ApiError {
	parameter := models.ParameterDifficulty{
		Description: req.Description,
	}

	err := service.db.Create(&parameter).Error

	if err != nil {
		zap.L().Error("error creating parameter difficulty", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service ParameterService) UpdateParameterScore(req requests.UpdateParameterScore) *dto.ApiError {
	parameter := models.ParameterScore{}

	err := service.db.
		Where("uuid = ?", req.ParameterScoreId).
		First(&parameter).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("parameter score not found", zap.String("uuid", req.ParameterScoreId.String()))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error get parameter score", zap.Error(err))
		return dto.InternalError(err)
	}

	parameter.Name = req.Name
	parameter.Description = req.Description

	err = service.db.Save(&parameter).Error

	if err != nil {
		zap.L().Error("error updating parameter score", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service ParameterService) UpdateParameterDifficulty(req requests.UpdateParameterDifficulty) *dto.ApiError {
	parameter := models.ParameterDifficulty{}

	err := service.db.
		Where("uuid = ?", req.ParameterDifficultyId).
		First(&parameter).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("parameter difficulty not found", zap.String("uuid", req.ParameterDifficultyId.String()))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error get parameter difficulty", zap.Error(err))
		return dto.InternalError(err)
	}

	parameter.Description = req.Description

	err = service.db.Save(&parameter).Error

	if err != nil {
		zap.L().Error("error updating parameter difficulty", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service ParameterService) DeleteParameterScore(parameterScoreId string) *dto.ApiError {
	parameter := models.ParameterScore{}

	err := service.db.
		Where("uuid = ?", parameterScoreId).
		First(&parameter).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("parameter score not found", zap.String("uuid", parameterScoreId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query parameter score", zap.Error(err))
		return dto.InternalError(err)
	}

	err = service.db.Delete(&parameter).Error

	if err != nil {
		zap.L().Error("error delete parameter score", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service ParameterService) DeleteParameterDifficulty(parameterDifficultyId string) *dto.ApiError {
	parameter := models.ParameterDifficulty{}

	err := service.db.
		Where("uuid = ?", parameterDifficultyId).
		First(&parameter).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("parameter difficulty not found", zap.String("uuid", parameterDifficultyId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query parameter difficulty", zap.Error(err))
		return dto.InternalError(err)
	}

	err = service.db.Delete(&parameter).Error

	if err != nil {
		zap.L().Error("error delete parameter difficulty", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}