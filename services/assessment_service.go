package services

import (
	"fmt"
	"sv-sfia/dto"
	"sv-sfia/dto/requests"
	"sv-sfia/dto/responses"
	"sv-sfia/models"
	"sv-sfia/utils"
	"time"

	"github.com/google/uuid"
	"github.com/gookit/goutil/dump"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AssessmentService struct {
	db *gorm.DB
}

func newAssessmentService(db *gorm.DB) *AssessmentService {
	return &AssessmentService{
		db: db,
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

func (service AssessmentService) CreateNewAssessment(participantId uuid.UUID) (responses.AssessmentResponse, *dto.ApiError) {
	year := time.Now().Year()

	assessment := models.Assessment{
		Year:          uint16(year),
		ParticipantId: participantId,
	}

	if !service.CanAssignAssessment(participantId) {
		return responses.AssessmentResponse{}, &dto.ApiError{
			Typ:          dto.ErrorBadData,
			Err:          fmt.Errorf("already have assessment this year"),
			ErrorMessage: "Sudah mengisi assessment tahun ini",
		}
	}

	err := service.db.Create(&assessment).Error
	if err != nil {
		zap.L().Error("error create new assessment", zap.Error(err))

		return responses.AssessmentResponse{}, dto.InternalError(err)
	}

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

func (service AssessmentService) GetDujAssessment(participantId uuid.UUID) {
	participantDepartment := models.ParticipantDepartment{}

	service.db.Where("participant_id = ?", participantId).
		Find(&participantDepartment)

	dump.P(participantDepartment)
}

func (service AssessmentService) StoreSelfAssessment(participantId uuid.UUID, req requests.SelfAssessmentRequest) *dto.ApiError {
	answers := []models.SelfAssessmentAnswer{}

	assesment := service.GetAssessement(participantId)

	for _, answer := range req.Answers {
		questionId, err := utils.ParseUUid(answer.QuestionId)
		if err != nil {
			return err
		}

		answers = append(answers, models.SelfAssessmentAnswer{
			QuestionId:   questionId,
			Value:        answer.Value,
			Evidence:     answer.Evidence,
			AssessmentId: assesment.Uuid,
		})
	}

	if len(answers) > 0 {
		err := service.db.Create(answers).Error

		if err != nil {
			return dto.InternalError(err)
		}
	}

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

func (service AssessmentService) StoreToolAssessment(participantId uuid.UUID, req requests.ToolAssessmentRequest) *dto.ApiError {
	if len(req.DeletedToolIds) > 0 {
		service.db.
			Where("uuid in ?", req.DeletedToolIds).
			Delete(&models.ParticipantTool{})
	}

	tools := []models.ParticipantTool{}

	for _, tool := range req.Tools {
		toolReq := models.ParticipantTool{
			ParticipantId: participantId,
			Tool:          tool.Name,
			Level:         tool.Level,
			Evidence:      tool.Evidence,
		}

		if tool.AnswerId != nil {
			answerId, err := utils.ParseUUid(*tool.AnswerId)
			if err != nil {
				return err
			}

			service.db.Where("uuid = ?", answerId).
				Updates(&toolReq)

			continue
		}

		tools = append(tools, toolReq)
	}

	if len(tools) > 0 {
		err := service.db.Create(&tools).Error

		if err != nil {
			zap.L().Error("failed store tools assessments", zap.Error(err))

			return &dto.ApiError{
				Typ:          dto.ErrorExec,
				ErrorMessage: "Gagal menyimpan tools",
				Err:          err,
			}
		}
	}

	return nil
}
