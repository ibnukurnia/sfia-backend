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

type DepartmentService struct {
	db *gorm.DB
}

func newDeparmentService(db *gorm.DB) *DepartmentService {
	return &DepartmentService{
		db: db,
	}
}

func (service *DepartmentService) GetDepartments(req requests.GetDepartmentsRequest) ([]responses.DepartmentResponse, *dto.ApiError) {
	departments := []models.Department{}

	query := service.db

	if req.Search != "" {
		query = query.Where("name ILIKE ?", "%"+req.Search+"%")
	}

	err := query.Order("created_at asc").
		Find(&departments).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("error querying department", zap.Error(err))
		return nil, dto.InternalError(err)
	}

	response := []responses.DepartmentResponse{}

	for _, department := range departments {
		d := responses.DepartmentResponse{
			Id:   department.Uuid.String(),
			Name: department.Name,
		}

		response = append(response, d)
	}

	return response, nil
}

func (service *DepartmentService) AddDepartment(req requests.AddDepartment) *dto.ApiError {
	department := models.Department{
		Name: req.Name,
	}

	err := service.db.Create(&department).Error

	if err != nil {
		zap.L().Error("error creating department", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *DepartmentService) UpdateDepartment(req requests.UpdateDepartment) *dto.ApiError {
	department := models.Department{}

	err := service.db.
		Where("uuid = ?", req.DepartmendId).
		First(&department).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("department not found", zap.String("uuid", req.DepartmendId.String()))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error get department", zap.Error(err))
		return dto.InternalError(err)
	}

	department.Name = req.Name

	err = service.db.Save(&department).Error

	if err != nil {
		zap.L().Error("error updating department", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *DepartmentService) DeleteDepartment(departmentId string) *dto.ApiError {
	department := models.Department{}

	err := service.db.
		Where("uuid = ?", departmentId).
		First(&department).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("department not found", zap.String("uuid", departmentId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query department", zap.Error(err))
		return dto.InternalError(err)
	}

	err = service.db.Delete(&department).Error

	if err != nil {
		zap.L().Error("error delete department", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *DepartmentService) GetDepartmentTeams(departmentId string) ([]responses.DepartmentTeamResponse, *dto.ApiError) {
	teams := []models.DepartmentTeam{}

	err := service.db.
		Where("department_id = ?", departmentId).
		Order("created_at asc").
		Find(&teams).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("error querying department team", zap.Error(err))
		return nil, dto.InternalError(err)
	}

	response := []responses.DepartmentTeamResponse{}

	for _, team := range teams {
		response = append(response, responses.DepartmentTeamResponse{
			Id:           team.Uuid.String(),
			DepartmentId: team.DepartmentId.String(),
			Name:         team.Name,
		})
	}

	return response, nil
}

func (service *DepartmentService) AddDepartmentTeam(req requests.AddDepartmentTeam) *dto.ApiError {
	team := models.DepartmentTeam{
		Name:         req.Name,
		DepartmentId: req.DepartmendId,
	}

	err := service.db.Create(&team).Error

	if err != nil {
		zap.L().Error("error creating department team", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *DepartmentService) UpdateDepartmentTeam(req requests.UpdateDepartmentTeam) *dto.ApiError {
	team := models.DepartmentTeam{}

	err := service.db.
		Where("uuid = ?", req.DepartmendId).
		First(&team).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("department team not found", zap.String("uuid", req.DepartmendId.String()))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query update department team", zap.Error(err))
		return dto.InternalError(err)
	}

	team.Name = req.Name

	err = service.db.Save(&team).Error

	if err != nil {
		zap.L().Error("error updating department team", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *DepartmentService) DeleteDepartmentTeam(teamId string) *dto.ApiError {
	team := models.DepartmentTeam{}

	err := service.db.
		Where("uuid = ?", teamId).
		First(&team).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("department team not found", zap.String("uuid", teamId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query delete department team", zap.Error(err))
		return dto.InternalError(err)
	}

	err = service.db.Delete(&team).Error

	if err != nil {
		zap.L().Error("error delete department team", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *DepartmentService) GetDepartmentUnits(departmentId string) ([]responses.DepartmentRoleResponse, *dto.ApiError) {
	units := []models.DepartmentUnit{}

	err := service.db.
		Where("department_id = ?", departmentId).
		Order("created_at asc").
		Find(&units).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, dto.InternalError(err)
	}

	response := []responses.DepartmentRoleResponse{}

	for _, unit := range units {
		response = append(response, responses.DepartmentRoleResponse{
			Id:           unit.Uuid.String(),
			DepartmentId: unit.DepartmentId.String(),
			Name:         unit.Name,
		})
	}

	return response, nil
}

func (service *DepartmentService) AddDepartmentUnit(req requests.AddDepartmentUnit) *dto.ApiError {
	unit := models.DepartmentUnit{
		Name:         req.Name,
		DepartmentId: req.DepartmendId,
	}

	err := service.db.Create(&unit).Error

	if err != nil {
		zap.L().Error("error creating department unit", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *DepartmentService) UpdateDepartmentUnit(req requests.UpdateDepartmentUnit) *dto.ApiError {
	unit := models.DepartmentUnit{}

	err := service.db.
		Where("uuid = ?", req.DepartmentUnitId).
		First(&unit).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("department unit not found", zap.String("uuid", req.DepartmentUnitId.String()))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query update department unit", zap.Error(err))
		return dto.InternalError(err)
	}

	unit.Name = req.Name

	err = service.db.Save(&unit).Error

	if err != nil {
		zap.L().Error("error updating department unit", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *DepartmentService) DeleteDepartmentUnit(unitId string) *dto.ApiError {
	unit := models.DepartmentUnit{}

	err := service.db.
		Where("uuid = ?", unitId).
		First(&unit).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("department unit not found", zap.String("uuid", unitId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query delete department unit", zap.Error(err))
		return dto.InternalError(err)
	}

	err = service.db.Delete(&unit).Error

	if err != nil {
		zap.L().Error("error delete department unit", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}
