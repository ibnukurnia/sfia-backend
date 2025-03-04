package services

import (
	"sv-sfia/dto"
	"sv-sfia/dto/responses"
	"sv-sfia/models"

	"github.com/google/uuid"
	"github.com/gookit/goutil/dump"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RoleService struct {
	db *gorm.DB
}

func newRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		db: db,
	}
}

func (service *RoleService) GetRoles() ([]responses.GroupedRolesResponse, *dto.ApiError) {
	groupRoles := []models.RoleGroup{}

	err := service.db.
		Preload("Roles").
		Find(&groupRoles).
		Error

	if err != nil {
		zap.L().Error("error querying roles", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	response := []responses.GroupedRolesResponse{}

	for _, group := range groupRoles {
		g := responses.GroupedRolesResponse{
			Name: group.Name,
		}

		roles := []responses.RoleResponse{}

		for _, role := range group.Roles {
			roles = append(roles, responses.RoleResponse{Name: role.Name, Id: role.Uuid.String()})
		}

		g.Roles = roles
		response = append(response, g)
	}

	return response, nil
}

func (service RoleService) GetParticipantRoles(participantId uuid.UUID) (dto.ParticipantRoleIdsDto, *dto.ApiError) {
	roles := models.ParticipantRole{}

	err := service.db.Where("participant_id = ?", participantId).
		First(&roles).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return dto.ParticipantRoleIdsDto{}, dto.InternalError(err)
	}

	return dto.ParticipantRoleIdsDto{
		MainRoleId:      roles.MainRoleId,
		SecondaryRoleId: roles.SecondaryRoleId,
		InterestRoleId:  roles.InterestRoleId,
	}, nil
}

func (service RoleService) GetRoleTraining(participantId uuid.UUID) (*responses.ParticipantTraingResponse, *dto.ApiError) {
	roles, apiErr := service.GetParticipantRoles(participantId)
	if apiErr != nil {
		return nil, apiErr
	}

	roleIds := []uuid.UUID{roles.MainRoleId}

	if roles.SecondaryRoleId != nil {
		roleIds = append(roleIds, *roles.SecondaryRoleId)
	}

	if roles.InterestRoleId != nil {
		roleIds = append(roleIds, *roles.InterestRoleId)
	}

	participantRoles := []models.Role{}

	err := service.db.Where("uuid in ?", roleIds).
		Preload("Trainings").
		Find(&participantRoles).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("error querying role trainings", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	trainings := []models.ParticipantTraining{}
	err = service.db.Where("participant_id = ?", participantId).
		Find(&trainings).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("error querying role trainings", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	currentTraining := map[string]models.ParticipantTraining{}

	for _, training := range trainings {
		currentTraining[training.TrainingId.String()] = training
	}

	roleResponse := []responses.RoleTrainingResponse{}

	for _, role := range participantRoles {
		trainings := []responses.TrainingResponse{}

		for _, training := range role.Trainings {
			r := responses.TrainingResponse{
				Id:   training.Uuid.String(),
				Name: training.Name,
			}

			if cT, exist := currentTraining[training.Uuid.String()]; exist {
				r.Selected = true
				r.NeedSertification = cT.NeedCertification
			}

			trainings = append(trainings, r)
		}

		roleResponse = append(roleResponse, responses.RoleTrainingResponse{
			Name:      role.Name,
			Trainings: trainings,
		})
	}

	dump.P(roleResponse)

	return &responses.ParticipantTraingResponse{Roles: roleResponse}, nil
}
