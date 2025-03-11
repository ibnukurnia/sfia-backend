package requests

import "github.com/google/uuid"

type AddDepartment struct {
	Name        string   `json:"name" validate:"required"`
}

func (r AddDepartment) Messages() map[string]string {
	return map[string]string{
		"Name.required": "Nama harus diisi",
	}
}

type UpdateDepartment struct {
	DepartmendId      uuid.UUID    `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
}

func (r UpdateDepartment) Messages() map[string]string {
	return map[string]string{
		"DepartmendId.required": "Department id harus diisi",
		"Name.required": "Nama harus diisi",
	}
}

type AddDepartmentUnit struct {
	DepartmendId      uuid.UUID    `json:"department_id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
}

func (r AddDepartmentUnit) Messages() map[string]string {
	return map[string]string{
		"DepartmendId.required": "Department id harus diisi",
		"Name.required": "Nama harus diisi",
	}
}

type UpdateDepartmentUnit struct {
	DepartmentUnitId	  uuid.UUID    `json:"department_unit_id" validate:"required"`
	DepartmendId      uuid.UUID    `json:"department_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
}

func (r UpdateDepartmentUnit) Messages() map[string]string {
	return map[string]string{
		"DepartmentUnitId.required": "Department unit id harus diisi",
		"DepartmendId.required": "Department id harus diisi",
		"Name.required": "Nama harus diisi",
	}
}

type AddDepartmentTeam struct {
	DepartmendId      uuid.UUID    `json:"department_id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
}

func (r AddDepartmentTeam) Messages() map[string]string {
	return map[string]string{
		"DepartmendId.required": "Department id harus diisi",
		"Name.required": "Nama harus diisi",
	}
}

type UpdateDepartmentTeam struct {
	DepartmentTeamId	  uuid.UUID    `json:"department_team_id" validate:"required"`
	DepartmendId      uuid.UUID    `json:"department_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
}

func (r UpdateDepartmentTeam) Messages() map[string]string {
	return map[string]string{
		"DepartmentTeamId.required": "Department team id harus diisi",
		"DepartmendId.required": "Department id harus diisi",
		"Name.required": "Nama harus diisi",
	}
}