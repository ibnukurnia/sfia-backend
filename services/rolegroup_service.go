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

type RoleGroupService struct {
	db *gorm.DB
}

func newRoleGroupService(db *gorm.DB) *RoleGroupService {
	return &RoleGroupService{
		db: db,
	}
}

func (service *RoleGroupService) GetRoleGroup() ([]responses.RoleGroupResponse, *dto.ApiError) {
	groupRoles := []models.RoleGroup{}

	err := service.db.
		Preload("Roles").
		Find(&groupRoles).
		Order("role_groups.created_at ASC").
		Error

	if err != nil {
		zap.L().Error("error querying roles", zap.Error(err))
		return nil, dto.InternalError(err)
	}

	response := []responses.RoleGroupResponse{}

	for _, group := range groupRoles {
		g := responses.RoleGroupResponse{
			Id: group.Uuid.String(),
			Name: group.Name,
		}
		response = append(response, g)
	}

	return response, nil
}

func (service *RoleGroupService) AddRoleGroup(req requests.AddRoleGroupRequest) (*dto.ApiError) {
	roleGroup := models.RoleGroup{
		Name: req.Name,
	}

	if err := service.db.Create(&roleGroup).Error; err != nil {
		zap.L().Error("error creating role group", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *RoleGroupService) UpdateRoleGroup(req requests.UpdateRoleGroupRequest) ( *dto.ApiError) {
	roleGroup := models.RoleGroup{}

	err := service.db.Where("uuid = ?", req.RoleGroupId).First(&roleGroup).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("role group not found", zap.String("uuid", req.RoleGroupId))
		return dto.NotFoundError(err)
	}
	
	if err != nil {
		zap.L().Error("error querying role group", zap.Error(err))

		return dto.InternalError(err)
	}

	roleGroup.Name = req.Name

	if err = service.db.Save(&roleGroup).Error; err != nil {
		zap.L().Error("error updating role group", zap.Error(err))
		return  dto.InternalError(err)
	}

	return nil
}

func (service *RoleGroupService) DeleteRoleGroup(roleGroupId string) (*dto.ApiError) {
	roleGroup := models.RoleGroup{}

	err := service.db.Where("uuid = ?", roleGroupId).First(&roleGroup).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("role group not found", zap.String("uuid", roleGroupId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error querying role group", zap.Error(err))
		return dto.InternalError(err)
	}

	if err = service.db.Delete(&roleGroup).Error; err != nil {
		zap.L().Error("error deleting role group", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}
