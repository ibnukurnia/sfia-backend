package models

import "github.com/google/uuid"

type AssessmentStatus string

const (
	NEW           AssessmentStatus = "NEW"
	SFIA          AssessmentStatus = "SFIA"
	DUJ           AssessmentStatus = "DUJ"
	TOOL          AssessmentStatus = "TOOL"
	UPDATETRANING AssessmentStatus = "UPDATE_TRAINING"
	TRANING       AssessmentStatus = "TRAINING"
	DONE          AssessmentStatus = "DONE"
)

type Assessment struct {
	Base
	ParticipantId         uuid.UUID
	Year                  uint16
	Status                AssessmentStatus
	Participant           Participant
	ParticipantSkills     []ParticipantSkill
	ParticipantRole       ParticipantRole
	ParticipantDepartment ParticipantDepartment
	SfiaResults           []SfiaResult
}
