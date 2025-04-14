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
	SkillId            uuid.UUID
	Question           string
	Descrtipion        string
	ParticipantAnswers []SelfAssessmentAnswer `gorm:"foreignKey:QuestionId"`
	ParticipantAnswer  SelfAssessmentAnswer   `gorm:"foreignKey:QuestionId"`
}

type SfiaAnswer struct {
	Base
	Answer string
	Value  int8
}

type SfiaResult struct {
	Base
	AssessmentId uuid.UUID
	SkillId      uuid.UUID
	Score        float32
	Skill        Skill
}
