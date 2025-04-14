package requests

type departmentInformation struct {
	DepartmentId     string `json:"department_id" validate:"required"`
	DepartmentTeamId string `json:"department_team_id" validate:"required"`
	DepartmentRoleId string `json:"department_role_id" validate:"required"`
}

type roleInformation struct {
	MainRoleId      string  `json:"main_role_id" validate:"required"`
	SecondaryRoleId *string `json:"secondary_role_id"`
	InterestRoleId  *string `json:"interest_role_id"`
}

type skillInformation struct {
	Id      string `json:"id" validate:"required"`
	UsedFor int8   `json:"used_for" valudate:"required"`
}

type CreateAssessmentRequest struct {
	Department departmentInformation `json:"department"`
	Role       roleInformation       `json:"role"`
	Skills     []skillInformation    `json:"skills"`
}

func (r *CreateAssessmentRequest) Messages() map[string]string {
	return map[string]string{
		"Department.DepartmentId.required":     "Department is Required",
		"Department.DepartmentTeamId.required": "Team is Required",
		"Department.DepartmentRoleId.required": "Department Role is Required",
		"Role.MainRoleId.required":             "Main role is required",
	}
}

type CreateAssessmentRoleRequest struct {
	MainRoleId      string  `json:"main_role_id" validate:"required"`
	SecondaryRoleId *string `json:"secondary_role_id"`
	InterestRoleId  *string `json:"interest_role_id"`
	AssessmentId    string  `json:"assessment_id" validate:"required"`
}

func (r *CreateAssessmentRoleRequest) Messages() map[string]string {
	return map[string]string{
		"MainRoleId.required":   "Main role is required",
		"AssessmentId.required": "Assessment id is required",
	}
}
