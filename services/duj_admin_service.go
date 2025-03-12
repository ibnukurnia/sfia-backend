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

type DujAdminService struct {
	db *gorm.DB
}

func newDujAdminService(db *gorm.DB) *DujAdminService {
	return &DujAdminService{db: db}
}

func (service DujAdminService) GetAdminDuj() ([]responses.AdminDujResponse, *dto.ApiError) {
	type QueryResult struct {
		ID             string  `gorm:"column:id"`
		JobDescription string  `gorm:"column:job_description"`
		RoleID         *string `gorm:"column:role_id"`         
		RoleName       *string `gorm:"column:role_name"`       
		SkillID        *string `gorm:"column:skill_id"`        
		SkillName      *string `gorm:"column:skill_name"`      
		CreatedAt      string  `gorm:"column:created_at"` 
	}

	var results []QueryResult

	err := service.db.Table("duj").
		Select("duj.uuid AS id, duj.job_description as job_description, roles.uuid AS role_id, roles.name AS role_name, skills.uuid AS skill_id, skills.name AS skill_name, duj.created_at").
		Joins("LEFT JOIN duj_roles ON duj.uuid = duj_roles.duj_id AND duj_roles.deleted_at is null").
		Joins("LEFT JOIN roles ON duj_roles.role_id = roles.uuid AND roles.deleted_at is null").
		Joins("LEFT JOIN duj_skills ON duj.uuid = duj_skills.duj_id and duj_skills.deleted_at is null").
		Joins("LEFT JOIN skills ON duj_skills.skill_id = skills.uuid and skills.deleted_at is null").
		Where("duj.deleted_at IS NULL").
		Order("duj.created_at asc").
		Find(&results).Error

	if err != nil {
		zap.L().Error("error querying admin duj", zap.Error(err))
		return []responses.AdminDujResponse{}, dto.InternalError(err)
	}

	dujMap := make(map[string]*responses.AdminDujResponse)
	var dujOrder []string

	for _, result := range results {
		if _, exists := dujMap[result.ID]; !exists {
			dujMap[result.ID] = &responses.AdminDujResponse{
				Id:             result.ID,
				JobDescription: result.JobDescription,
				Roles:          []responses.RoleItem{},
				Skillset:       []responses.SkillItem{},
			}
			dujOrder = append(dujOrder, result.ID)
		}

		if result.RoleID != nil && result.RoleName != nil {
			roleExists := false
			for _, role := range dujMap[result.ID].Roles {
				if role.Id == *result.RoleID {
					roleExists = true
					break
				}
			}
			if !roleExists {
				dujMap[result.ID].Roles = append(dujMap[result.ID].Roles, responses.RoleItem{
					Id:   *result.RoleID,
					Name: *result.RoleName,
				})
			}
		}

		if result.SkillID != nil && result.SkillName != nil {
			skillExists := false
			for _, skill := range dujMap[result.ID].Skillset {
				if skill.Id == *result.SkillID {
					skillExists = true
					break
				}
			}
			if !skillExists {
				dujMap[result.ID].Skillset = append(dujMap[result.ID].Skillset, responses.SkillItem{
					Id:   *result.SkillID,
					Name: *result.SkillName,
				})
			}
		}
	}

	adminDujResponses := []responses.AdminDujResponse{}
	for _, id := range dujOrder {
		adminDujResponses = append(adminDujResponses, *dujMap[id])
	}

	return adminDujResponses, nil
}

func (service DujAdminService) AddAdminDuj(req requests.AddAdminDujRequest) *dto.ApiError {
	duj := models.DujAdmin{
		JobDescription: req.JobDescription,
	}

	if err := service.db.Create(&duj).Error; err != nil {
		zap.L().Error("error creating duj", zap.Error(err))
		return dto.InternalError(err)
	}

	for _, role := range req.Roles {
		dujRole := models.DujRole{
			DujID: duj.Base.Uuid,
			RoleID: role,
		}

		if err := service.db.Create(&dujRole).Error; err != nil {
			zap.L().Error("error creating duj role", zap.Error(err))
			return dto.InternalError(err)
		}
	}

	for _, skill := range req.Skills {
		dujSkill := models.DujSkill{
			DujID: duj.Base.Uuid,
			SkillID: skill,
		}

		if err := service.db.Create(&dujSkill).Error; err != nil {
			zap.L().Error("error creating duj skill", zap.Error(err))
			return dto.InternalError(err)
		}
	}

	return nil
}

func (service DujAdminService) UpdateAdminDuj(req requests.UpdateAdminDujRequest) *dto.ApiError {
	duj := models.DujAdmin{}

	err := service.db.Where("uuid = ?", req.DujId).First(&duj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("duj not found", zap.String("uuid", req.DujId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query duj", zap.Error(err))
		return dto.InternalError(err)
	}

	duj.JobDescription = req.JobDescription

	if err := service.db.Save(&duj).Error; err != nil {
		zap.L().Error("error updating duj", zap.Error(err))
		return dto.InternalError(err)
	}

	if err := service.db.Where("duj_id = ?", req.DujId).Delete(&models.DujRole{}).Error; err != nil {
		zap.L().Error("error deleting duj role", zap.Error(err))
		return dto.InternalError(err)
	}

	if err := service.db.Where("duj_id = ?", req.DujId).Delete(&models.DujSkill{}).Error; err != nil {
		zap.L().Error("error deleting duj skill", zap.Error(err))
		return dto.InternalError(err)
	}

	for _, role := range req.Roles {
		dujRole := models.DujRole{
			DujID: duj.Base.Uuid,
			RoleID: role,
		}

		if err := service.db.Create(&dujRole).Error; err != nil {
			zap.L().Error("error creating duj role", zap.Error(err))
			return dto.InternalError(err)
		}
	}

	for _, skill := range req.Skills {
		dujSkill := models.DujSkill{
			DujID: duj.Base.Uuid,
			SkillID: skill,
		}

		if err := service.db.Create(&dujSkill).Error; err != nil {
			zap.L().Error("error creating duj skill", zap.Error(err))
			return dto.InternalError(err)
		}
	}

	return nil
}

func (service DujAdminService) DeleteAdminDuj(dujId string) *dto.ApiError {
	duj := models.DujAdmin{}

	err := service.db.Where("uuid = ?", dujId).First(&duj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Warn("duj not found", zap.String("uuid", dujId))
		return dto.NotFoundError(err)
	}

	if err != nil {
		zap.L().Error("error query duj", zap.Error(err))
		return dto.InternalError(err)
	}

	if err := service.db.Delete(&duj).Error; err != nil {
		zap.L().Error("error deleting duj", zap.Error(err))
		return dto.InternalError(err)
	}

	return nil
}