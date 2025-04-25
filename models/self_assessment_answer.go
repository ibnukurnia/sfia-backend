package models

import "github.com/google/uuid"

type SelfAssessmentAnswer struct {
	Base
	AssessmentId  uuid.UUID
	ParticipantId uuid.UUID
	QuestionId    uuid.UUID
	RoleId        uuid.UUID
	SkillId       uuid.UUID
	Value         int8
	Evidence      string
	UserId        uuid.UUID
	// Question      SfiaQuestion `gorm:"foreignKey:QuestionId"`
}

func (m *SelfAssessmentAnswer) TableName() string {
	return "participant_sfia_answers"
}
