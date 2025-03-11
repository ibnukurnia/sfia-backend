package models

import "github.com/google/uuid"

type Department struct {
	Base
	Name            string
	DepartmentTeams []DepartmentTeam
	DepartmentUnits []DepartmentUnit
}

type DepartmentTeam struct {
	Base
	DepartmentId uuid.UUID
	Name 	   string	
	Department
}

type DepartmentUnit struct {
	Base
	DepartmentId uuid.UUID
	Name 	   string
	Department
	Dujs []Duj
}
