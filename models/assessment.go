package models

import "github.com/google/uuid"

type AssessmentStatus string

const SFIA AssessmentStatus = "SFIA"
const DUJ AssessmentStatus = "DUJ"
const TOOL AssessmentStatus = "TOOL"

type Assessment struct {
	Base
	ParticipantId uuid.UUID
	Year          uint16
}
