package services

import (
	"fmt"
	"math"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"
	"sv-sfia/utils"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AssessmentService struct {
	db             *gorm.DB
	skillThreshold models.SkillLevelThreshold
}

func newAssessmentService(db *gorm.DB) *AssessmentService {
	return &AssessmentService{
		db: db,
		skillThreshold: models.SkillLevelThreshold{
			Basic:        0,
			Intermediate: 3.0,
			Advance:      6.0,
		},
	}
}

func (service AssessmentService) CanAssignAssessment(participantId uuid.UUID) bool {
	year := time.Now().Year()

	return service.db.Model(&models.Assessment{}).
		Where("participant_id = ?", participantId).
		Where("year = ?", year).
		Find(&models.Assessment{}).
		RowsAffected < 1
}

func (service AssessmentService) GetAssessement(participantId uuid.UUID) *models.Assessment {
	year := time.Now().Year()

	assessment := models.Assessment{}

	service.db.
		Where("participant_id = ?", participantId).
		Where("year = ?", year).
		Find(&assessment)

	return &assessment
}

func (service AssessmentService) GetAssessmentById(assessmentId uuid.UUID) (*models.Assessment, *dto.ApiError) {
	assessment := models.Assessment{}

	errQuery := service.db.Where("uuid = ?", assessmentId).
		First(&assessment).
		Error

	if errQuery != nil {
		err := dto.InternalError(errQuery)

		if errQuery == gorm.ErrRecordNotFound {
			err = dto.NotFoundError(errQuery)
		}

		return nil, err
	}

	return &assessment, nil
}

func (service AssessmentService) ListAssessment(participantId uuid.UUID) []responses.ListAssessmentResponse {
	assessments := []models.Assessment{}

	service.db.Where("participant_id = ?", participantId).
		Find(&assessments)

	res := []responses.ListAssessmentResponse{}

	for _, assessment := range assessments {
		res = append(res, responses.ListAssessmentResponse{
			Id:     assessment.Uuid.String(),
			Year:   assessment.Year,
			Status: assessment.Status,
		})
	}

	return res
}

func (service AssessmentService) CreateNewAssessment(participantId uuid.UUID, request requests.CreateAssessmentRequest) (responses.AssessmentResponse, *dto.ApiError) {
	year := time.Now().Year()

	assessment := models.Assessment{
		Year:          uint16(year),
		ParticipantId: participantId,
		Status:        models.SFIA,
	}

	if !service.CanAssignAssessment(participantId) {
		return responses.AssessmentResponse{}, &dto.ApiError{
			Typ:          dto.ErrorBadData,
			Err:          fmt.Errorf("already have assessment this year"),
			ErrorMessage: "Sudah mengisi assessment tahun ini",
		}
	}

	tx := service.db.Begin()

	err := tx.Create(&assessment).Error
	if err != nil {
		zap.L().Error("error create new assessment", zap.Error(err))

		tx.Rollback()

		return responses.AssessmentResponse{}, dto.InternalError(err)
	}

	mainRoleId, errParse := utils.ParseUUid(request.Role.MainRoleId)
	if errParse != nil {
		zap.L().Error("error create new assessment", zap.Error(err))

		return responses.AssessmentResponse{}, errParse
	}

	roles := models.ParticipantRole{
		ParticipantId:   participantId,
		AssessmentId:    assessment.Uuid,
		MainRoleId:      mainRoleId,
		SecondaryRoleId: new(uuid.UUID),
		InterestRoleId:  new(uuid.UUID),
	}

	if request.Role.SecondaryRoleId != nil {
		secondaryRoleId, errParse := utils.ParseUUid(*request.Role.SecondaryRoleId)
		if errParse != nil {
			zap.L().Error("error create new assessment", zap.Error(err))

			return responses.AssessmentResponse{}, errParse
		}

		*roles.SecondaryRoleId = secondaryRoleId
	} else {
		roles.SecondaryRoleId = nil
	}

	if request.Role.InterestRoleId != nil {
		interestRoleId, errParse := utils.ParseUUid(*request.Role.InterestRoleId)
		if errParse != nil {
			zap.L().Error("error create new assessment", zap.Error(err))

			return responses.AssessmentResponse{}, errParse
		}

		*roles.InterestRoleId = interestRoleId
	} else {
		roles.InterestRoleId = nil
	}

	err = tx.Create(&roles).Error
	if err != nil {
		zap.L().Error("error create new assessment", zap.Error(err))

		tx.Rollback()

		return responses.AssessmentResponse{}, dto.InternalError(err)
	}

	departmentId, errParse := utils.ParseUUid(request.Department.DepartmentId)
	if errParse != nil {
		zap.L().Error("error create new assessment", zap.Error(err))

		return responses.AssessmentResponse{}, errParse
	}

	departmentTeamId, errParse := utils.ParseUUid(request.Department.DepartmentTeamId)
	if errParse != nil {
		zap.L().Error("error create new assessment", zap.Error(err))

		return responses.AssessmentResponse{}, errParse
	}

	departmentUnitId, errParse := utils.ParseUUid(request.Department.DepartmentRoleId)
	if errParse != nil {
		zap.L().Error("error create new assessment", zap.Error(err))

		return responses.AssessmentResponse{}, errParse
	}

	department := models.ParticipantDepartment{
		ParticipantId:    participantId,
		AssessmentId:     assessment.Uuid,
		DepartmentId:     departmentId,
		DepartmentTeamId: departmentTeamId,
		DepartmentUnitId: departmentUnitId,
	}

	err = tx.Create(&department).Error
	if err != nil {
		zap.L().Error("error create new assessment", zap.Error(err))

		tx.Rollback()

		return responses.AssessmentResponse{}, dto.InternalError(err)
	}

	skills := []models.ParticipantSkill{}

	for _, skill := range request.Skills {
		skillId, errParse := utils.ParseUUid(skill.Id)
		if errParse != nil {
			zap.L().Error("error create new assessment", zap.Error(err))

			return responses.AssessmentResponse{}, errParse
		}

		skills = append(skills, models.ParticipantSkill{
			ParticipantId: participantId,
			AssessmentId:  assessment.Uuid,
			SkillId:       skillId,
			UsedFor:       skill.UsedFor,
			IsMastered:    true,
		})
	}

	if len(skills) > 0 {
		err = tx.Create(&skills).Error

		if err != nil {
			zap.L().Error("error create new assessment", zap.Error(err))

			tx.Rollback()

			return responses.AssessmentResponse{}, dto.InternalError(err)
		}
	}

	tx.Commit()

	return responses.AssessmentResponse{
		Id: assessment.Uuid.String(),
	}, nil
}

func (service AssessmentService) GetSelfAssessment(participantId uuid.UUID) responses.SelfAssessmentResponse {
	participantRole := models.ParticipantRole{}

	// geting participant roles
	service.db.Where("participant_id = ?", participantId).
		Find(&participantRole)

	roleIds := []uuid.UUID{participantRole.MainRoleId}

	if participantRole.SecondaryRoleId != nil {
		roleIds = append(roleIds, *participantRole.SecondaryRoleId)
	}

	if participantRole.InterestRoleId != nil {
		roleIds = append(roleIds, *participantRole.InterestRoleId)
	}

	skillIds := []uuid.UUID{}

	// get skills
	service.db.
		Model(models.ParticipantSkill{}).
		Select("skill_id").
		Where("participant_id = ?", participantId).
		Where("is_mastered = ?", true).
		Find(&skillIds)

	skills := []models.Skill{}

	service.db.Where("uuid in ?", skillIds).
		Preload("RoleSkills", func(q *gorm.DB) *gorm.DB {
			return q.Where("role_id in ?", roleIds)
		}).
		Preload("RoleSkills.Role").
		Preload("SfiaQuestions").
		Find(&skills)

	answers := []models.SfiaAnswer{}

	service.db.Order("value asc").Find(&answers)

	return responses.NewSfiaResponse(skills, answers)
}

func (service AssessmentService) StoreSelfAssessment(assessmentId string, req requests.SelfAssessmentRequest) *dto.ApiError {
	answers := []models.SelfAssessmentAnswer{}

	assessmentUuid, err := utils.ParseUUid(assessmentId)
	if err != nil {
		return err
	}

	assessment, err := service.GetAssessmentById(assessmentUuid)
	if err != nil {
		return err
	}

	for _, answer := range req.Answers {
		questionId, err := utils.ParseUUid(answer.QuestionId)
		if err != nil {
			return err
		}

		answers = append(answers, models.SelfAssessmentAnswer{
			QuestionId:   questionId,
			Value:        answer.Value,
			Evidence:     answer.Evidence,
			AssessmentId: assessment.Uuid,
		})
	}

	assessment.Status = models.DUJ

	tx := service.db.Begin()

	if len(answers) > 0 {
		err := tx.Create(&answers).Error

		if err != nil {
			tx.Rollback()

			zap.L().Error("error save sfia answer: ", zap.Error(err))

			return dto.InternalError(err)
		}

	}

	if err := tx.Save(&assessment).Error; err != nil {
		tx.Rollback()

		zap.L().Error("error save sfia answer: ", zap.Error(err))

		return dto.InternalError(err)
	}

	tx.Commit()

	service.CalculateResult(assessment.Uuid)

	return nil
}

func (service AssessmentService) UpdateSelfAssessmentAnswer(answerId uuid.UUID, req requests.SelfAssessmentAnswer) *dto.ApiError {
	err := service.db.
		Model(&models.SelfAssessmentAnswer{}).
		Where("uuid = ?", answerId).Updates(&models.SelfAssessmentAnswer{
		Value:    req.Value,
		Evidence: req.Evidence,
	}).Error

	if err != nil {
		return &dto.ApiError{
			Typ:          dto.ErrorExec,
			ErrorMessage: "Gagal menyimpan jawaban",
			Err:          err,
		}
	}

	return nil
}

func (service AssessmentService) GetToolAssessment(participantId uuid.UUID) responses.ToolAssessmentResponse {
	tools := []models.ParticipantTool{}

	service.db.Where("participant_id = ?", participantId).
		Find(&tools)

	return responses.NewToolAssessmentResponse(tools)
}

func (service AssessmentService) StoreToolAssessment(participantId, assessmentId uuid.UUID, request requests.CreateParticipantToolRequest) *dto.ApiError {
	tools := []models.ParticipantTool{}

	for _, tool := range request.Tools {
		t := models.Tool{
			Url: "",
		}

		if tool.Id == nil {
			t.Name = tool.Name

			err := service.db.Create(&t).Error
			if err != nil {
				zap.L().Error("error create new tools:", zap.Error(err))

				return dto.InternalError(err)
			}
		} else {
			toolId, err := utils.ParseUUid(*tool.Id)
			if err != nil {
				return err
			}

			t.Uuid = toolId
		}

		tools = append(tools, models.ParticipantTool{
			ParticipantId: participantId,
			AssessmentId:  assessmentId,
			ToolId:        t.Uuid,
			Name:          "",
			Level:         tool.Level,
			Evidence:      tool.Evidence,
		})
	}

	if len(tools) > 0 {
		err := service.db.Create(&tools).Error

		if err != nil {
			zap.L().Error("error saving participant tools", zap.Error(err))

			return dto.InternalError(err)
		}
	}

	err := service.db.Model(&models.Assessment{}).Where("uuid = ?", assessmentId).
		Update("status", models.UPDATETRANING).Error
	if err != nil {
		zap.L().Error("error saving participant tools", zap.Error(err))

		return dto.InternalError(err)
	}

	return nil
}

func (service AssessmentService) CalculateResult(assessmentId uuid.UUID) {
	participantSkills := []models.ParticipantSkill{}

	service.db.Preload("Skill").
		Preload("Skill.SfiaQuestions").
		Preload("Skill.SfiaQuestions.ParticipantAnswer", func(tx *gorm.DB) *gorm.DB {
			return tx.Where("assessment_id = ?", assessmentId)
		}).
		Where("assessment_id = ?", assessmentId).Find(&participantSkills)

	skillScore := map[string]int8{}
	totalQuestion := map[string]int8{}

	for _, skill := range participantSkills {
		for _, question := range skill.Skill.SfiaQuestions {
			if value, exist := skillScore[skill.SkillId.String()]; !exist {
				skillScore[skill.SkillId.String()] = question.ParticipantAnswer.Value
			} else {
				skillScore[skill.SkillId.String()] = value + question.ParticipantAnswer.Value
			}

			totalQuestion[skill.SkillId.String()] += 1
		}
	}

	results := []models.SfiaResult{}

	for skillId, score := range skillScore {
		id, _ := utils.ParseUUid(skillId)

		finalScore := float32(math.Round(float64(float32(score)/float32(totalQuestion[skillId]))*100) / 100)

		results = append(results, models.SfiaResult{
			SkillId:      id,
			AssessmentId: assessmentId,
			Score:        finalScore,
		})
	}

	service.db.Create(results)
}

func (service AssessmentService) Resume(assessementId uuid.UUID) (*responses.AssessmentResumeResponse, *dto.ApiError) {
	service.db.First(&service.skillThreshold)

	assessment := models.Assessment{}

	err := service.db.Preload("Participant").
		Preload("ParticipantRole").
		Preload("ParticipantSkills").
		Preload("ParticipantRole.MainRole").
		Preload("ParticipantRole.MainRole.Group").
		Preload("ParticipantRole.SecondaryRole").
		Preload("ParticipantRole.SecondaryRole.Group").
		Preload("ParticipantRole.InterestRole").
		Preload("ParticipantRole.InterestRole.Group").
		Preload("ParticipantDepartment").
		Preload("ParticipantDepartment.Department").
		Preload("ParticipantDepartment.DepartmentUnit").
		Preload("ParticipantDepartment.DepartmentTeam").
		Preload("SfiaResults").
		Preload("SfiaResults.Skill").
		Where("uuid = ?", assessementId).First(&assessment).Error

	if err != nil {
		zap.L().Error("error query assessment", zap.Error(err))

		if err == gorm.ErrRecordNotFound {

			return nil, dto.NotFoundErrorWithMsg(err, "Assessment not found")
		}

		return nil, dto.InternalError(err)
	}

	skillIds := []uuid.UUID{}

	for _, result := range assessment.SfiaResults {
		skillIds = append(skillIds, result.SkillId)
	}

	skillResults := []responses.SkillResult{}

	for _, result := range assessment.SfiaResults {
		skillResults = append(skillResults, responses.SkillResult{
			Code:  result.Skill.Code,
			Score: result.Score,
			Level: service.skillThreshold.ToLevel(result.Score),
			Name:  result.Skill.Name,
		})
	}

	res := responses.AssessmentResumeResponse{
		GeneralInformation: responses.GeneralInformationResume{
			Name:           assessment.Participant.Name,
			Email:          assessment.Participant.Email,
			Pn:             assessment.Participant.Pn,
			Department:     assessment.ParticipantDepartment.Department.Name,
			DepartmentUnit: assessment.ParticipantDepartment.DepartmentUnit.Name,
			DepartmentTeam: assessment.ParticipantDepartment.DepartmentTeam.Name,
		},
		SkillResults: skillResults,
	}

	roleInformation := responses.RoleInformationResume{
		MainRole: responses.RoleResume{
			Name:  assessment.ParticipantRole.MainRole.Name,
			Group: assessment.ParticipantRole.MainRole.Group.Name,
			Level: service.roleLevelResult(assessment.ParticipantRole.MainRoleId, skillIds, skillResults),
		},
	}

	if assessment.ParticipantRole.SecondaryRole != nil {
		roleInformation.SecondaryRole = &responses.RoleResume{
			Name:  assessment.ParticipantRole.SecondaryRole.Name,
			Group: assessment.ParticipantRole.SecondaryRole.Group.Name,
			Level: service.roleLevelResult(*assessment.ParticipantRole.SecondaryRoleId, skillIds, skillResults),
		}
	}

	if assessment.ParticipantRole.InterestRole != nil {
		roleInformation.InterestRole = &responses.RoleResume{
			Name:  assessment.ParticipantRole.InterestRole.Name,
			Group: assessment.ParticipantRole.InterestRole.Group.Name,
			Level: service.roleLevelResult(*assessment.ParticipantRole.InterestRoleId, skillIds, skillResults),
		}
	}

	res.RoleInformation = roleInformation

	return &res, nil
}

func (service AssessmentService) SfiaResult(assessementId uuid.UUID) (*responses.SfiaRoleResultResponse, *dto.ApiError) {
	service.db.First(&service.skillThreshold)

	assessment := models.Assessment{}

	skillIds := []uuid.UUID{}

	err := service.db.Model(&models.ParticipantSkill{}).Select("skill_id").Where("assessment_id = ?", assessementId).Find(&skillIds).Error
	if err != nil {
		zap.L().Error("error", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	err = service.db.
		Preload("ParticipantRole").
		Preload("ParticipantRole.MainRole").
		Preload("ParticipantRole.MainRoleSkills", func(q *gorm.DB) *gorm.DB {
			return q.Where("skill_id in ?", skillIds)
		}).
		Preload("ParticipantRole.MainRole.Group").
		Preload("ParticipantRole.MainRoleSkills.Skill").
		Preload("ParticipantRole.SecondaryRole").
		Preload("ParticipantRole.SecondaryRoleSkills", func(q *gorm.DB) *gorm.DB {
			return q.Where("skill_id in ?", skillIds)
		}).
		Preload("ParticipantRole.SecondaryRole.Group").
		Preload("ParticipantRole.SecondaryRoleSkills.Skill").
		Preload("ParticipantRole.InterestRole").
		Preload("ParticipantRole.InterestRoleSkills", func(q *gorm.DB) *gorm.DB {
			return q.Where("skill_id in ?", skillIds)
		}).
		Preload("ParticipantRole.InterestRoleSkills.Skill").
		Preload("ParticipantRole.InterestRole.Group").
		Preload("SfiaResults").
		Preload("SfiaResults").
		Where("uuid = ?", assessementId).First(&assessment).Error

	if err != nil {
		zap.L().Error("error", zap.Error(err))

		return nil, dto.InternalError(err)
	}

	mainRoleResults := service.skillResult(assessment.ParticipantRole.MainRoleSkills, assessment.SfiaResults)
	secondaryRoleResults := service.skillResult(assessment.ParticipantRole.SecondaryRoleSkills, assessment.SfiaResults)
	interestRoleResults := service.skillResult(assessment.ParticipantRole.InterestRoleSkills, assessment.SfiaResults)

	res := responses.SfiaRoleResultResponse{
		MainRole: responses.SfiaRoleResult{
			Skills: mainRoleResults,
			Name:   assessment.ParticipantRole.MainRole.Name,
			Group:  assessment.ParticipantRole.MainRole.Group.Name,
			Level:  service.roleLevelResult(assessment.ParticipantRole.MainRoleId, skillIds, mainRoleResults),
		},
	}

	if assessment.ParticipantRole.SecondaryRole != nil {
		secondaryRole := responses.SfiaRoleResult{
			Skills: secondaryRoleResults,
			Name:   assessment.ParticipantRole.SecondaryRole.Name,
			Group:  assessment.ParticipantRole.SecondaryRole.Group.Name,
			Level:  service.roleLevelResult(*assessment.ParticipantRole.SecondaryRoleId, skillIds, secondaryRoleResults),
		}

		res.SecondaryRole = &secondaryRole
	}

	if assessment.ParticipantRole.InterestRole != nil {
		interestRole := responses.SfiaRoleResult{
			Skills: interestRoleResults,
			Name:   assessment.ParticipantRole.InterestRole.Name,
			Group:  assessment.ParticipantRole.InterestRole.Group.Name,
			Level:  service.roleLevelResult(*assessment.ParticipantRole.InterestRoleId, skillIds, interestRoleResults),
		}

		res.InterestRole = &interestRole
	}

	return &res, nil
}

func (service AssessmentService) skillResult(skills []models.RoleSkill, results []models.SfiaResult) []responses.SkillResult {
	skillIds := map[string]models.SfiaResult{}

	for _, result := range results {
		skillIds[result.SkillId.String()] = result
	}

	res := []responses.SkillResult{}

	for _, skill := range skills {
		result, exist := skillIds[skill.SkillId.String()]

		var score float32

		if exist {
			score = result.Score
		}

		res = append(res, responses.SkillResult{
			Code:  skill.Skill.Code,
			Score: score,
			Level: service.skillThreshold.ToLevel(score),
			Name:  skill.Skill.Name,
		})
	}

	return res
}

func (service AssessmentService) roleLevelResult(roleId uuid.UUID, skillIds []uuid.UUID, skillResults []responses.SkillResult) models.RoleLevel {
	mapSkillLevel := map[string]models.SkillLevel{}

	for _, skill := range skillResults {
		mapSkillLevel[skill.Code] = skill.Level
	}

	roleSkills := []models.RoleSkill{}

	service.db.
		Preload("Skill").
		Where("skill_id in ?", skillIds).
		Where("role_id = ?", roleId).
		Find(&roleSkills)

	checker := newRoleLevelChecker()

	level := checkRoleLevel(&checker, roleSkills, mapSkillLevel)

	if level == models.NOLEVEL {
		return models.JUNIOR
	}

	return level
}

type LevelChecker struct {
	Level            models.RoleLevel
	PrevLevel        models.RoleLevel
	IsMandatory      func(models.RoleSkill) bool
	GetRequiredLevel func(models.RoleSkill) int
	Next             *LevelChecker
}

func newRoleLevelChecker() LevelChecker {
	levelCheckerSenior := LevelChecker{
		Level:            models.SENIOR,
		PrevLevel:        models.MIDDLE,
		IsMandatory:      func(rs models.RoleSkill) bool { return rs.IsMandatoryForSenior },
		GetRequiredLevel: func(rs models.RoleSkill) int { return rs.RequirementForSenior.Int() },
		Next:             nil,
	}

	levelCheckerMiddle := LevelChecker{
		Level:            models.MIDDLE,
		PrevLevel:        models.JUNIOR,
		IsMandatory:      func(rs models.RoleSkill) bool { return rs.IsMandatoryForMiddle },
		GetRequiredLevel: func(rs models.RoleSkill) int { return rs.RequirementForMiddle.Int() },
		Next:             &levelCheckerSenior,
	}

	return LevelChecker{
		Level:            models.JUNIOR,
		PrevLevel:        models.NOLEVEL,
		IsMandatory:      func(rs models.RoleSkill) bool { return rs.IsMandatoryForJunior },
		GetRequiredLevel: func(rs models.RoleSkill) int { return rs.RequirementForJunior.Int() },
		Next:             &levelCheckerMiddle,
	}
}

func checkRoleLevel(checker *LevelChecker, requirements []models.RoleSkill, skills map[string]models.SkillLevel) models.RoleLevel {
	if checker == nil {
		// BASE CASE: sudah mencapai level terakhir yang berhasil lolos
		// return level sebelumnya (karena level sekarang = nil)
		// atau bisa juga: panic atau return dari argumen luar
		// tapi kita ganti sekarang: return previous level yang valid
		return models.SENIOR // misalnya SENIOR itu level tertinggi
	}

	checked := false

	for _, req := range requirements {
		if !checker.IsMandatory(req) {
			continue
		}

		checked = true
		skill, exists := skills[req.Skill.Code]
		if !exists || checker.GetRequiredLevel(req) > skill.Int() {
			// Tidak lolos di level ini → return level sebelumnya
			return checker.PrevLevel
		}
	}

	if !checked {
		if checker.Level != checker.PrevLevel {
			return checker.PrevLevel
		}

		return checker.Level
	}

	// Semua requirement di level ini lolos → lanjut ke level berikutnya
	// tapi kalau level berikutnya nil (akhir), return level ini
	if checker.Next == nil {
		return checker.Level // ini base case sebenarnya
	}

	return checkRoleLevel(checker.Next, requirements, skills)
}
