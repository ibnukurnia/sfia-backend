package services

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"
	"sv-sfia/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserAdminService struct {
	db *gorm.DB
}

func newUserAdminService(db *gorm.DB) *UserAdminService {
	return &UserAdminService{db: db}
}

func (service UserAdminService) GetUserAdmin(req requests.UserAdminRequest) (responses.UserAdminPaginatedResponse, *dto.ApiError) {
	var results []responses.UserAdminResponse
	var totalRecords int64

	query := service.db.Table("participants p").
		Select("p.uuid, p.name, p.pn, p.role_access, roles.name as role_name, corporations.name as corporation_title, d.name as department").
		Joins("LEFT JOIN participant_roles pr ON pr.participant_id = p.uuid AND pr.deleted_at is null").
		Joins("LEFT JOIN roles ON roles.uuid = pr.main_role_id AND roles.deleted_at is null").
		Joins("LEFT JOIN corporations ON p.corporation_id = corporations.uuid AND corporations.deleted_at is null").
		Joins("LEFT JOIN participant_departments pd ON pd.participant_id = p.uuid AND pd.deleted_at is null").
		Joins("LEFT JOIN departments d ON d.uuid = pd.department_id AND d.deleted_at is null").
		Where("p.deleted_at IS NULL")

	if req.Search != "" {
		searchQuery := "%" + req.Search + "%"
		query = query.Where("p.name LIKE ? OR p.pn LIKE ? OR roles.name LIKE ? OR corporations.name LIKE ? OR d.name LIKE ?",
			searchQuery, searchQuery, searchQuery, searchQuery, searchQuery)
	}

	// Apply corporate_id filter
	if len(req.CorporateIDs) > 0 {
		query = query.Where("p.corporation_id IN ?", req.CorporateIDs)
	}

	// Apply role_id filter
	if len(req.RoleIDs) > 0 {
		query = query.Where("pr.main_role_id IN ?", req.RoleIDs)
	}

	// Apply department_id filter
	if len(req.DepartmentIDs) > 0 {
		query = query.Where("pd.department_id IN ?", req.DepartmentIDs)
	}

	// Get total records
	if err := query.Count(&totalRecords).Error; err != nil {
		zap.L().Error("error counting user admin records", zap.Error(err))
		return responses.UserAdminPaginatedResponse{}, dto.InternalError(err)
	}

	// Calculate pagination
	offset := (req.Page - 1) * req.Limit

	// Get paginated results
	err := query.Order("p.created_at asc").
		Offset(offset).
		Limit(req.Limit).
		Find(&results).Error

	if err != nil {
		zap.L().Error("error querying user admin", zap.Error(err))
		return responses.UserAdminPaginatedResponse{}, dto.InternalError(err)
	}

	response := responses.UserAdminPaginatedResponse{
		Data: results,
	}
	response.Paginator = utils.NewPaginator(req.Page, req.Limit, int(totalRecords))

	return response, nil
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
