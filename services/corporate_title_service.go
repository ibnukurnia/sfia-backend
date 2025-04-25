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

type CorporateTitleService struct {
	db *gorm.DB
}

func newCorporateTitleService(db *gorm.DB) *CorporateTitleService {
	return &CorporateTitleService{
		db: db,
	}
}

func (service *CorporateTitleService) GetCorporateTitles(req requests.GetCorporateTitlesRequest) ([]responses.CorporateTitleListResponse, *dto.ApiError) {
	corporations := []responses.CorporateTitleListResponse{}

	query := service.db.Table("corporations").
		Select("corporations.uuid, corporations.name").
		Where("corporations.deleted_at IS NULL")

	if req.Search != "" {
		query = query.Where("corporations.name ILIKE ?", "%"+req.Search+"%")
	}

	err := query.Order("corporations.created_at ASC").
		Find(&corporations).Error

	if err != nil {
		zap.L().Error("error querying corporations", zap.Error(err))
		return nil, dto.InternalError(err)
	}

	return corporations, nil
}

func (service *CorporateTitleService) AddCorporateTitle(req requests.AddCorporateTitleRequest) *dto.ApiError {

	corporateTitle := models.Corporations{
		Name: req.Name,
	}

	if err := service.db.Create(&corporateTitle).Error; err != nil {
		zap.L().Error("error creating corporaet title", zap.Error(err))

		return dto.InternalError(err)
	}

	return nil
}

func (service *CorporateTitleService) UpdateCorporateTitle(req requests.UpdateCorporateTitleRequest) *dto.ApiError {
	corporateTitle := models.Corporations{}

	err := service.db.Where("uuid = ?", req.CorporateTitleId).First(&corporateTitle).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("corporateTitle not found", zap.String("uuid", req.CorporateTitleId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error querying corporateTitle", zap.Error(err))
		return dto.InternalError(err)
	}

	corporateTitle.Name = req.Name

	if err = service.db.Save(&corporateTitle).Error; err != nil {
		zap.L().Error("error updating corporateTitle", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *CorporateTitleService) DeleteCorporateTitle(corporateTitleId string) *dto.ApiError {
	corporateTitle := models.Corporations{}

	err := service.db.Where("uuid = ?", corporateTitleId).First(&corporateTitle).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("corporateTitle not found", zap.String("uuid", corporateTitleId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error querying corporateTitle", zap.Error(err))
		return dto.InternalError(err)
	}

	if err = service.db.Delete(&corporateTitle).Error; err != nil {
		zap.L().Error("error deleting corporateTitle", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}
