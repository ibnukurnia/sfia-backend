package models

import "github.com/google/uuid"

type ParticipantDepartment struct {
	Base
	ParticipantId    uuid.UUID
	DepartmentId     uuid.UUID
	DepartmentTeamId uuid.UUID
	DepartmentUnitId uuid.UUID
	DepartmentUnit   DepartmentUnit `gorm:"foreignKey:DepartmentUnitId"`
}
