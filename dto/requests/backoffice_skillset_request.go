package requests

import "github.com/google/uuid"

type AddSkillsetRequest struct {
	RoleId 				uuid.UUID `json:"role_id" validate:"required,uuid"`
	SkillsetName        string `json:"name" validate:"required"`
	SkillsetDescription string `json:"description" validate:"required"`
}

func (r AddSkillsetRequest) Messages() map[string]string {
	return map[string]string{
		"RoleId.required": 				"Id Role harus diisi",
		"SkillsetName.required":        "Nama skillset harus diisi",
		"SkillsetDescription.required":        "Deskripsi skillset harus diisi",
	}
}

type UpdateSkillsetRequest struct {
	SkillsetId 			uuid.UUID `json:"skill_id" validate:"required,uuid"`
	RoleId 				uuid.UUID `json:"role_id" validate:"required,uuid"`
	SkillsetName        string `json:"name" validate:"required"`
	SkillsetDescription string `json:"description" validate:"required"`
}

func (r UpdateSkillsetRequest) Messages() map[string]string {
	return map[string]string{
		"SkillsetId.required": 				"Id skillset harus diisi",
		"RoleId.required": 				"Id Role harus diisi",
		"SkillsetName.required":        "Nama skillset harus diisi",
		"SkillsetDescription.required":        "Deskripsi skillset harus diisi",
	}
}