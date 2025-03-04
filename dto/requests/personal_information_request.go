package requests

type PersonalInformationRequest struct {
	DepartmentId     string `json:"department_id" validate:"required"`
	DepartmentRoleId string `json:"department_role_id" validate:"required"`
	DepartmentTeamId string `json:"department_team_id" validate:"required"`
}

func (r PersonalInformationRequest) Messages() map[string]string {
	return map[string]string{
		"DepartmentId.required":     "Departemen harus diisi",
		"DepartmentRoleId.required": "Departemen role harus diisi",
		"DepartmentTeamId.required": "Departemen tim harus diisi",
	}
}
