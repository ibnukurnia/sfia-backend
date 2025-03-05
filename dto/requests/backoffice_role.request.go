package requests

import "github.com/google/uuid"

type AddRoleRequest struct {
	Name        string   `json:"name" validate:"required"`
	RoleGroupId uuid.UUID `json:"role_group_id" validate:"required"`
}

func (r AddRoleRequest) Messages() map[string]string {
	return map[string]string{
		"Name.required": "Nama harus diisi",
		"RoleGroupId.required": "Role Group id harus diisi",
	}
}

type UpdateRoleRequest struct {
	RoleId      string    `json:"role_id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	RoleGroupId uuid.UUID `json:"role_group_id" validate:"required"`
}

func (r UpdateRoleRequest) Messages() map[string]string {
	return map[string]string{
		"RoleId.required": "Role id harus diisi",
		"Name.required": "Nama harus diisi",
		"RoleGroupId.required": "Role Group id harus diisi",
	}
}