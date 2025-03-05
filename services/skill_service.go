package services

import (
	"errors"
	"fmt"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
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

func (service SkillService) GetSkillsetList() ([]responses.SkillsetResponse, *dto.ApiError) {
	skills := []responses.SkillsetResponse{}

	err := service.db.Table("role_skills rs").
		Select("r.uuid as role_id, r.name as role_name, s.uuid as skill_id, s.name as skill_name, s.description as skill_description").
		Joins("LEFT JOIN roles r ON r.uuid = rs.role_id AND r.deleted_at IS NULL").
		Joins("JOIN skills s ON s.uuid = rs.skill_id").
		Where("rs.deleted_at IS NULL").
		Order("s.created_at asc").
		Find(&skills).Error

	if err != nil {
		zap.L().Error("error query skill", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	return skills, nil
}

func (service SkillService) AddSkillSet(req requests.AddSkillsetRequest) *dto.ApiError {
	tx := service.db.Begin()

	skill := models.Skill{
		Name: req.SkillsetName,
		Description: req.SkillsetDescription,
	}
	
	if err := tx.Create(&skill).Error; err != nil {
		tx.Rollback()
		zap.L().Error("error insert skill", zap.Error(tx.Error))
		return dto.InternalError(tx.Error)
	}

	role := models.Role{}

	err := tx.Where("uuid = ?", req.RoleId).First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		zap.L().Warn("role not found", zap.String("uuid", req.RoleId.String()))
		return dto.NotFoundError(err)
	}
	
	if err != nil {
		tx.Rollback()
		zap.L().Error("error querying role", zap.Error(err))
		return dto.InternalError(err)
	}

	roleSkill := models.RoleSkill{
		RoleId: req.RoleId,
		SkillId: skill.Uuid,
	}

	if err := tx.Create(&roleSkill).Error; err != nil {
		tx.Rollback()
		zap.L().Error("error insert role skill", zap.Error(tx.Error))
		return dto.InternalError(err)
	}

	if err := tx.Commit().Error; err != nil {
		zap.L().Error("error commit query", zap.Error(tx.Error))
		return dto.InternalError(err)
	}

	return nil
}

func (service SkillService) UpdateSkillSet(req requests.UpdateSkillsetRequest) *dto.ApiError {
	tx := service.db.Begin()

	skill := models.Skill{}

	err := tx.Where("uuid = ?", req.SkillsetId).First(&skill).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		zap.L().Warn("skill not found", zap.String("uuid", req.SkillsetId.String()))
		return dto.NotFoundError(err)
	}
	if err != nil {
		tx.Rollback()
		zap.L().Error("error querying skill", zap.Error(err))
		return dto.InternalError(err)
	}

	err = tx.Where("uuid = ?", req.RoleId).First(&models.Role{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		zap.L().Warn("role not found", zap.String("uuid", req.SkillsetId.String()))
		return dto.NotFoundError(err)
	}
	if err != nil {
		tx.Rollback()
		zap.L().Error("error querying skill", zap.Error(err))
		return dto.InternalError(err)
	}

	skill.Name = req.SkillsetName
	skill.Description = req.SkillsetDescription

	if err = tx.Save(&skill).Error; err != nil {
		tx.Rollback()
		zap.L().Error("error updating skill", zap.Error(err))
		return dto.InternalError(err)
	}

	roleSkill := models.RoleSkill{}

	err = tx.Where("skill_id = ?", req.SkillsetId).First(&roleSkill).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		zap.L().Warn("skillset not found")
		return dto.NotFoundError(err)
	}

	if err != nil {
		tx.Rollback()
		zap.L().Error("error querying role skill", zap.Error(err))

		return dto.InternalError(err)
	}

	roleSkill.RoleId = req.RoleId

	if err = tx.Save(&roleSkill).Error; err != nil {
		tx.Rollback()
		zap.L().Error("error updating role skill", zap.Error(err))
		return dto.InternalError(err)
	}

	if err := tx.Commit().Error; err != nil {
		zap.L().Error("error commit query", zap.Error(tx.Error))
		return dto.InternalError(err)
	}

	return nil
}

func (service SkillService) DeleteSkillSet(skillId string) *dto.ApiError {
	tx := service.db.Begin()

	skill := models.Skill{}

	err := tx.Where("uuid = ?", skillId).First(&skill).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		zap.L().Warn("skill not found", zap.String("uuid", skillId))
		return dto.NotFoundError(err)
	}
	if err != nil {
		tx.Rollback()
		zap.L().Error("error querying skill", zap.Error(err))

		return dto.InternalError(err)
	}

	if err = tx.Delete(&skill).Error; err != nil {
		tx.Rollback()
		zap.L().Error("error deleting skill", zap.Error(err))
		return dto.InternalError(err)
	}

	roleSkill := models.RoleSkill{}

	err = tx.Where("skill_id = ?", skillId).First(&roleSkill).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		zap.L().Warn("skillset not found")
		return dto.NotFoundError(err)
	}
	if err != nil {
		tx.Rollback()
		zap.L().Error("error querying role skill", zap.Error(err))

		return dto.InternalError(err)
	}

	if err = tx.Delete(&roleSkill).Error; err != nil {
		tx.Rollback()
		zap.L().Error("error deleting role skill", zap.Error(err))
		return dto.InternalError(err)
	}

	if err := tx.Commit().Error; err != nil {
		zap.L().Error("error commit query", zap.Error(tx.Error))
		return dto.InternalError(err)
	}

	return nil
}