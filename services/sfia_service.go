package services

import (
	"sv-sfia/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SfiaService struct {
	db *gorm.DB
}

func newSfiaService(db *gorm.DB) *SfiaService {
	return &SfiaService{db: db}
}

// func (service SfiaService) getCurrentSelfAssessmentAnswers(participantId uuid.UUID) ([]models.SelfAssessmentAnswer, error) {
// 	answers := []models.SelfAssessmentAnswer{}

// 	err := service.db.Where("participant_id = ?", participantId).
// 		Find(&answers).Error

// 	return answers, err
// }

// func (service SfiaService) GetSfias(participantId uuid.UUID, skillIds []string) ([]responses.SfiaResponse, *dto.ApiError) {
// 	// assessments := []models.Sfia{}
// 	skills := []models.ParticipantSkill{}

// 	service.db.
// 		Where("participant_id = ?", participantId).
// 		Preload("Skill").
// 		Find(&skills)

// 	dump.P(skills)

// 	return nil, nil
// }

func (service SfiaService) StoreSelfAssesmentAnswers(participantId uuid.UUID) *dto.ApiError {
	return nil
	// currentAnswers := map[string]string{}

	// answers, err := service.getCurrentSelfAssessmentAnswers(participantId)
	// if err != nil {
	// 	zap.L().Error("error query current sfia answer", zap.Error(err))

	// 	return &dto.ApiError{
	// 		Typ:          dto.ErrorInternal,
	// 		Err:          err,
	// 		ErrorMessage: "internal server error",
	// 	}
	// }

	// for _, answer := range answers {
	// 	currentAnswers[answer.QuestionId.String()] = answer.AnswerId.String()
	// }

	// for _, answer := range request.Answers {
	// 	answerId, err := uuid.Parse(answer.AnswerId)

	// 	if err != nil {
	// 		zap.L().Error("error parsing uuid", zap.Error(err))

	// 		return &dto.ApiError{
	// 			Typ:          dto.ErrorInternal,
	// 			Err:          err,
	// 			ErrorMessage: "internal server error",
	// 		}
	// 	}

	// 	questionId, err := uuid.Parse(answer.QuestionId)
	// 	if err != nil {
	// 		zap.L().Error("error parsing uuid", zap.Error(err))

	// 		return &dto.ApiError{
	// 			Typ:          dto.ErrorInternal,
	// 			Err:          err,
	// 			ErrorMessage: "internal server error",
	// 		}
	// 	}

	// 	if currentAnswerId, exist := currentAnswers[answer.QuestionId]; exist {
	// 		if currentAnswerId == answer.AnswerId {
	// 			continue
	// 		}

	// 		err := service.db.Model(&models.SelfAssessmentAnswer{}).
	// 			Where("participant_id = ?", participantId).
	// 			Where("question_id = ?", questionId).
	// 			Update("answer_id", answerId).Error

	// 		if err != nil {
	// 			zap.L().Error("error updating sfia answer", zap.Error(err))

	// 			return &dto.ApiError{
	// 				Typ:          dto.ErrorInternal,
	// 				Err:          err,
	// 				ErrorMessage: "internal server error",
	// 			}
	// 		}

	// 		continue
	// 	}

	// 	err = service.db.Create(&models.SelfAssessmentAnswer{
	// 		ParticipantId: participantId,
	// 		AnswerId:      answerId,
	// 		QuestionId:    questionId,
	// 	}).Error

	// 	if err != nil {
	// 		zap.L().Error("error storing sfia answer", zap.Error(err))

	// 		return &dto.ApiError{
	// 			Typ:          dto.ErrorInternal,
	// 			Err:          err,
	// 			ErrorMessage: "internal server error",
	// 		}
	// 	}
	// }

	// return nil
}
