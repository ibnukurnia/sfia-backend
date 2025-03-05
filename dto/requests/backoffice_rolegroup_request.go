package requests

type AddRoleGroupRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r AddRoleGroupRequest) Messages() map[string]string {
	return map[string]string{
		"Name.required": "Nama harus diisi",
	}
}

type UpdateRoleGroupRequest struct {
	RoleGroupId string `json:"role_group_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

func (r UpdateRoleGroupRequest) Messages() map[string]string {
	return map[string]string{
		"RoleGroupId.required": "Id Role Group harus diisi",
		"Name.required":        "Nama harus diisi",
	}
}
