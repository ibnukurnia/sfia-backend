package models

import "github.com/google/uuid"

type ParticipantDepartment struct {
	Base
	ParticipantId    uuid.UUID
	DepartmentId     uuid.UUID
	DepartmentTeamId uuid.UUID
	DepartmentUnitId uuid.UUID
	AssessmentId     uuid.UUID
	Department       Department     `gorm:"foreignKey:DepartmentId"`
	DepartmentUnit   DepartmentUnit `gorm:"foreignKey:DepartmentUnitId"`
	DepartmentTeam   DepartmentTeam `gorm:"foreignKey:DepartmentTeamId"`
}
