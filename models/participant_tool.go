package models

import "github.com/google/uuid"

type ParticipantTool struct {
	Base
	ParticipantId uuid.UUID
	ToolId        uuid.UUID
	AssessmentId  uuid.UUID
	Name          string `gorm:"column:tool"`
	Level         string
	Evidence      string
	Tool          Tool
}
