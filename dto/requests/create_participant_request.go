package requests

type CreateParticipantRequest struct {
	Pn           string `json:"pn" validate:"required"`
	Email        string `json:"email" validate:"email,required"`
	Name         string `json:"name" validate:"required"`
	DepartmentId string `json:"department_id" validate:"required"`
	// DepartmentTeamId string  `json:"department_team_id" validate:"required"`
	// DepartmentRoleId string  `json:"department_role_id" validate:"required"`
	MainRoleId      string  `json:"main_role_id" validate:"required"`
	SecondaryRoleId *string `json:"secondary_role_id"`
	InterestRoleId  *string `json:"interest_role_id"`
}

func (r *CreateParticipantRequest) Messages() map[string]string {
	return map[string]string{
		"Pn.required":               "PN is Required",
		"Email.required":            "Email is Required",
		"Email.email":               "Invalid email",
		"Name.required":             "Name is Required",
		"DepartmentId.required":     "Department is Required",
		"DepartmentTeamId.required": "Team is Required",
		"DepartmentRoleId.required": "Department Role is Required",
		"MainRoleId.required":       "Main role is required",
	}
}
