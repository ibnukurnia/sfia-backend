package models

import "github.com/google/uuid"

type Sfia struct {
	Base
	Description string
	SkillId     uint
	Skill       Skill
}

type SfiaQuestion struct {
	Base
	SkillId           uuid.UUID
	Question          string
	Descrtipion       string
	ParticipantAnswer *SelfAssessmentAnswer `gorm:"foreignKey:QuestionId"`
}

type SfiaAnswer struct {
	Base
	Answer string
	Value  int8
}
