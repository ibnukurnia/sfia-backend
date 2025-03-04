package services

import (
	"fmt"
	"sv-sfia/dto"
	"sv-sfia/dto/responses"
	"sv-sfia/models"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SkillService struct {
	db *gorm.DB
}

func newSkillService(db *gorm.DB) *SkillService {
	return &SkillService{db: db}
}

func (service SkillService) FindSkillByRoleIds(participantId uuid.UUID, roleDto dto.ParticipantRoleIdsDto) ([]responses.SkillResponse, *dto.ApiError) {
	roleIds := []uuid.UUID{roleDto.MainRoleId}

	if roleDto.SecondaryRoleId != nil {
		roleIds = append(roleIds, *roleDto.SecondaryRoleId)
	}

	if roleDto.InterestRoleId != nil {
		roleIds = append(roleIds, *roleDto.InterestRoleId)
	}

	skillSets := []models.RoleSkill{}

	err := service.db.Distinct("skill_id").
		Preload("Skill").
		Where("role_id in ?", roleIds).
		Find(&skillSets).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("error query role skills", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	participantSkills := []models.ParticipantSkill{}

	err = service.db.
		Where("participant_id = ?", participantId).
		Where("is_mastered = ?", true).
		Find(&participantSkills).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		zap.L().Error("error query role skills", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	existingSkills := map[string]int{}

	for _, skill := range participantSkills {
		existingSkills[skill.SkillId.String()] = int(skill.UsedFor)
	}

	res := []responses.SkillResponse{}

	for _, skillSet := range skillSets {
		skillId := skillSet.Skill.Uuid.String()
		usedFor, exist := existingSkills[skillId]

		res = append(res, responses.SkillResponse{
			Id:         skillId,
			Skill:      fmt.Sprintf("%s (%s)", skillSet.Skill.Name, skillSet.Skill.Code),
			IsMastered: exist,
			UsedFor:    usedFor,
		})
	}

	return res, nil
}

func (service SkillService) GetParticipantSkills(participantId uuid.UUID) ([]string, *dto.ApiError) {
	skillIds := []uuid.UUID{}

	err := service.db.
		Model(&models.ParticipantSkill{}).
		Where("participant_id = ?", participantId).
		Where("is_mastered = ?", true).
		Select("skill_id").
		Find(&skillIds).Error

	if err != nil {
		zap.L().Error("error query participant skills", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	results := []string{}

	for _, skill := range skillIds {
		results = append(results, skill.String())
	}

	return results, nil
}
