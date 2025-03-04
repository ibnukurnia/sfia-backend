package models

import "github.com/google/uuid"

type ParticipantTool struct {
	Base
	ParticipantId uuid.UUID
	Tool          string
	Level         string
	Evidence      string
}
