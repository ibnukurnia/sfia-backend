package models

import (
	"github.com/google/uuid"
)

type Duj struct {
	Base
	Name             string
	Detail           string
	DepartmentUnitId uuid.UUID
}

func (model Duj) TableName() string {
	return "department_unit_jobs"
}

type DujAdmin struct {
	Base
	JobDescription string
}

func (model DujAdmin) TableName() string {
	return "duj"
}

type DujRole struct {
	Base
	DujID  uuid.UUID
	RoleID uuid.UUID
}

type DujSkill struct {
	Base
	DujID   uuid.UUID
	SkillID uuid.UUID
}