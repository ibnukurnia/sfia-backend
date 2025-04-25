package requests

type UpdateParticipantRoleRequest struct {
	MainRoleId      string  `json:"main_role_id" validate:"required"`
	SecondaryRoleId *string `json:"secondary_role_id"`
	InterestRoleId  *string `json:"interest_role_id"`
}

func (r *UpdateParticipantRoleRequest) Messages() map[string]string {
	return map[string]string{
		"MainRoleId.required": "Role Utama Harus Diisi",
	}
}
