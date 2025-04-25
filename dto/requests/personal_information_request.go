package requests

type PersonalInformationRequest struct {
	DepartmentId     string `json:"department_id" validate:"required"`
	DepartmentUnitId string `json:"department_unit_id" validate:"required"`
	DepartmentTeamId string `json:"department_team_id" validate:"required"`
}

func (r PersonalInformationRequest) Messages() map[string]string {
	return map[string]string{
		"DepartmentId.required":     "Departemen harus diisi",
		"DepartmentUnitId.required": "Fungsi harus diisi",
		"DepartmentTeamId.required": "Tim harus diisi",
	}
}
