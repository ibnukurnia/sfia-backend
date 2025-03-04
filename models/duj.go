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
