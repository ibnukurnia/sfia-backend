package responses

type TrainingMasterResponse struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Jenjang  string   `json:"jenjang"`
	SkillId  string   `json:"skills_id"`
	Level    string   `json:"level"`
	Type     string   `json:"type"`
	Mode     string   `json:"mode"`
	Provider []string `json:"provider"`
	Silabus  string   `json:"silabus"`
}