package models

import "github.com/google/uuid"

type ParticipantTraining struct {
	Base
	IsNeeded          bool
	NeedCertification bool
	Name              string

	// IsPriority        bool
	Priority      *int8
	TrainingId    uuid.UUID
	ParticipantId uuid.UUID
	AssessmentId  uuid.UUID
}
