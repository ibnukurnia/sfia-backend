package services

import (
	"sv-sfia/dto"
	"sv-sfia/dto/responses"
	"sv-sfia/models"

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

func (service *DepartmentService) GetDepartments() ([]responses.DepartmentResponse, *dto.ApiError) {
	departments := []models.Department{}

	err := service.db.
		Find(&departments).Error

	if err != nil && err != gorm.ErrRecordNotFound {
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

func (service *DepartmentService) GetDepartmentTeams(departmentId string) ([]responses.DepartmentTeamResponse, *dto.ApiError) {
	teams := []models.DepartmentTeam{}

	err := service.db.
		Where("department_id = ?", departmentId).
		Find(&teams).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, dto.InternalError(err)
	}

	response := []responses.DepartmentTeamResponse{}

	for _, team := range teams {
		response = append(response, responses.DepartmentTeamResponse{
			Id:   team.Uuid.String(),
			Name: team.Name,
		})
	}

	return response, nil
}

func (service *DepartmentService) GetDepartmentUnits(departmentId string) ([]responses.DepartmentRoleResponse, *dto.ApiError) {
	roles := []models.DepartmentUnit{}

	err := service.db.
		Where("department_id = ?", departmentId).
		Find(&roles).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, dto.InternalError(err)
	}

	response := []responses.DepartmentRoleResponse{}

	for _, team := range roles {
		response = append(response, responses.DepartmentRoleResponse{
			Id:   team.Uuid.String(),
			Name: team.Name,
		})
	}

	return response, nil
}
