package models

import "github.com/google/uuid"

type DujAnswer struct {
	Base
	ParticipantId uuid.UUID
	Job           string
	Detail        string
	CurrentJob    bool
	HaveTrouble   bool
	TroubleCause  *string
}

func (m DujAnswer) TableName() string {
	return "participant_duj_answers"
}
