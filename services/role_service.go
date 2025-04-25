package services

import (
	"errors"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"
	"sv-sfia/utils"

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

func (service *RoleService) GetRoleSkills(roleIds []string) ([]responses.SkillResponse, *dto.ApiError) {
	skills := []models.RoleSkill{}

	roleUUids := []uuid.UUID{}

	for _, roleId := range roleIds {
		id, err := utils.ParseUUid(roleId)
		if err != nil {
			return nil, err
		}

		roleUUids = append(roleUUids, id)
	}

	err := service.db.Distinct("skill_id").
		Preload("Skill").
		Where("role_id in ?", roleUUids).
		Find(&skills).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("error query participant skills", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	res := []responses.SkillResponse{}

	for _, skill := range skills {
		res = append(res, responses.SkillResponse{
			Id:    skill.SkillId.String(),
			Skill: skill.Skill.Name,
		})
	}

	return res, nil
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
			Id:        role.Uuid.String(),
			Name:      role.Name,
			Trainings: trainings,
		})
	}

	dump.P(roleResponse)

	return &responses.ParticipantTraingResponse{Roles: roleResponse}, nil
}

func (service *RoleService) GetRoleList(req requests.GetRoleListRequest) ([]responses.RoleListResponse, *dto.ApiError) {
	roles := []responses.RoleListResponse{}

	query := service.db.Table("roles").
		Select("roles.uuid, roles.name, role_groups.uuid as group_id, role_groups.name AS group_name").
		Joins("LEFT JOIN role_groups ON roles.group_id = role_groups.uuid AND role_groups.deleted_at IS NULL").
		Where("roles.deleted_at IS NULL")

	if req.Search != "" {
		query = query.Where("roles.name ILIKE ?", "%"+req.Search+"%")
	}

	err := query.Order("roles.created_at ASC").
		Find(&roles).Error

	if err != nil {
		zap.L().Error("error querying roles", zap.Error(err))
		return nil, dto.InternalError(err)
	}

	return roles, nil
}

func (service *RoleService) AddRole(req requests.AddRoleRequest) *dto.ApiError {
	roleGroup := models.RoleGroup{}

	err := service.db.Where("uuid = ?", req.RoleGroupId).First(&roleGroup).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("role group not found", zap.String("uuid", req.RoleGroupId.String()))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error querying role group", zap.Error(err))
		return dto.InternalError(err)
	}

	role := models.Role{
		Name:    req.Name,
		GroupId: req.RoleGroupId,
	}

	if err := service.db.Create(&role).Error; err != nil {
		zap.L().Error("error creating role", zap.Error(err))

		return dto.InternalError(err)
	}

	return nil
}

func (service *RoleService) UpdateRole(req requests.UpdateRoleRequest) *dto.ApiError {
	role := models.Role{}

	err := service.db.Where("uuid = ?", req.RoleId).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("role not found", zap.String("uuid", req.RoleId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error querying role", zap.Error(err))
		return dto.InternalError(err)
	}

	roleGroup := models.RoleGroup{}

	err = service.db.Where("uuid = ?", req.RoleGroupId).First(&roleGroup).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("role group not found", zap.String("uuid", req.RoleGroupId.String()))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error querying role group", zap.Error(err))
		return dto.InternalError(err)
	}

	role.Name = req.Name
	role.GroupId = req.RoleGroupId

	if err = service.db.Save(&role).Error; err != nil {
		zap.L().Error("error updating role", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}

func (service *RoleService) DeleteRole(roleId string) *dto.ApiError {
	role := models.Role{}

	err := service.db.Where("uuid = ?", roleId).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("role not found", zap.String("uuid", roleId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error querying role", zap.Error(err))
		return dto.InternalError(err)
	}

	if err = service.db.Delete(&role).Error; err != nil {
		zap.L().Error("error deleting role", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}
