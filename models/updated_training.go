package models

import (
	"time"

	"github.com/google/uuid"
)

const (
	SYNC    TrainingImplementation = "SYNC"
	ASYNC   TrainingImplementation = "ASYNC"
	ONLINE  TrainingLocation       = "ONLINE"
	OFFLINE TrainingLocation       = "OFFLINE"
)

type TrainingImplementation string

type TrainingLocation string

type UpdatedTraining struct {
	Base
	Name             string
	HasCertification bool
	GetCertification bool
	Implementation   TrainingImplementation
	Location         TrainingLocation
	Provider         string
	StartDate        time.Time
	EndDate          time.Time
	AssessmentId     uuid.UUID
}

func (m UpdatedTraining) TableName() string {
	return "participant_updated_trainings"
}
