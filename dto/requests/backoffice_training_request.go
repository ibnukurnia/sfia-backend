package requests

import "github.com/google/uuid"

type AddTrainingMasterRequest struct {
	Name    string `json:"name" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Jenjang string `json:"jenjang" validate:"required"`
	SkillId uuid.UUID `json:"skills_id" validate:"required"`
	Level  string `json:"level" validate:"required"`
	Type  string `json:"type" validate:"required"`
	Mode  string `json:"mode" validate:"required"`
	Provider  []string `json:"provider" validate:"required"`
	Silabus  string `json:"silabus" validate:"required"`
}

func (r AddTrainingMasterRequest) Messages() map[string]string {
	return map[string]string{
		"Name.required": 				"Nama harus diisi",
		"Code.required": 				"Code harus diisi",
		"Jenjang.required": 			"Jenjang harus diisi",
		"SkillId.required": 			"Skill harus diisi",
		"Level.required": 				"Level harus diisi",
		"Type.required": 				"Type harus diisi",
		"Mode.required": 				"Mode harus diisi",
		"Provider.required": 			"Provider harus diisi",
		"Silabus.required": 			"Silabus harus diisi",
	}
}

type UpdateTrainingMasterRequest struct {
	TrainingId uuid.UUID `json:"id" validate:"required,uuid"`
	Name	string `json:"name" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Jenjang string `json:"jenjang" validate:"required"`
	SkillId uuid.UUID `json:"skills_id" validate:"required"`
	Level  string `json:"level" validate:"required"`
	Type  string `json:"type" validate:"required"`
	Mode  string `json:"mode" validate:"required"`
	Provider  []string `json:"provider" validate:"required"`
	Silabus  string `json:"silabus" validate:"required"`
}

func (r UpdateTrainingMasterRequest) Messages() map[string]string {
	return map[string]string{
		"TrainingId.required": 			"Id training harus diisi",
		"Name.required": 				"Nama harus diisi",
		"Code.required": 				"Code harus diisi",
		"Jenjang.required": 			"Jenjang harus diisi",
		"SkillId.required": 			"Skill harus diisi",
		"Level.required": 				"Level harus diisi",
		"Type.required": 				"Type harus diisi",
		"Mode.required": 				"Mode harus diisi",
		"Provider.required": 			"Provider harus diisi",
		"Silabus.required": 			"Silabus harus diisi",
	}
}