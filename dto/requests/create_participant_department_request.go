package requests

type CreateParticipantDepartmentRequest struct {
	DepartmentId     string `json:"department_id" validate:"required"`
	DepartmentTeamId string `json:"department_team_id" validate:"required"`
	DepartmentUnitId string `json:"department_Unit_id" validate:"required"`
}

func (r *CreateParticipantDepartmentRequest) Messages() map[string]string {
	return map[string]string{
		"DepartmentId.required":     "Departemen tidak boleh kosong",
		"DepartmentTeamId.required": "Team tidak boleh kosong",
		"DepartmentUnitId.required": "Fungsi tidak boleh kosong",
	}
}
