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

type UserAdminService struct {
	db *gorm.DB
}

func newUserAdminService(db *gorm.DB) *UserAdminService {
	return &UserAdminService{db: db}
}

func (service UserAdminService) GetUserAdmin() ([]responses.UserAdminResponse, *dto.ApiError) {
	

	var results []responses.UserAdminResponse

	err := service.db.Table("participants p").
		Select("p.uuid, p.name, p.pn, p.role_access, roles.name as role_name, corporations.name as corporation_title").
		Joins("LEFT JOIN participant_roles pr ON pr.participant_id = p.uuid AND pr.deleted_at is null").
		Joins("LEFT JOIN roles ON roles.uuid = pr.main_role_id AND roles.deleted_at is null").
		Joins("LEFT JOIN corporations ON p.corporation_id = corporations.uuid AND corporations.deleted_at is null").
		Where("p.deleted_at IS NULL").
		Order("p.created_at asc").
		Find(&results).Error

	if err != nil {
		zap.L().Error("error querying user admin", zap.Error(err))
		return []responses.UserAdminResponse{}, dto.InternalError(err)
	}


	return results, nil
}

func (service UserAdminService) UpdateRoleAccess(req requests.UpdateUserAdminRequest) *dto.ApiError {
	user := models.Participant{}

	err := service.db.Where("uuid = ?", req.Id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("participants not found", zap.String("uuid", req.Id.String()))
		return dto.NotFoundError(err)
	}
	if err != nil {
		zap.L().Error("error query participants", zap.Error(err))
		return dto.InternalError(err)
	}

	user.RoleAccess = req.Role

	err = service.db.Save(&user).Error
	if err != nil {
		zap.L().Error("error updating role access", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service UserAdminService) DeleteUserAdmin(userId string) *dto.ApiError {
	user := models.Participant{}

	err := service.db.
		Where("uuid = ?", userId).
		First(&user).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("participants not found", zap.String("uuid", userId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query participants", zap.Error(err))
		return dto.InternalError(err)
	}

	err = service.db.Delete(&user).Error

	if err != nil {
		zap.L().Error("error delete participants", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}