package models

import "github.com/google/uuid"

type DujAnswer struct {
	Base
	ParticipantId uuid.UUID
	AssessmentId  uuid.UUID
	JobId         uuid.UUID
	CurrentJob    bool
	HaveTrouble   bool
	TroubleCause  *string
	Job           Duj `gorm:"foreignKey:JobId"`
}

func (m DujAnswer) TableName() string {
	return "participant_duj_answers"
}
