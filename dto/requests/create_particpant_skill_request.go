package requests

type StoreParticipantSkillRequest struct {
	Skills []Skill `json:"skills" validate:"required,min=1"`
}

type Skill struct {
	Id         string `json:"id"`
	IsMastered bool   `json:"is_mastered"`
	UsedFor    int    `json:"used_for"`
}

func (r *StoreParticipantSkillRequest) Messages() map[string]string {
	return map[string]string{
		"Skills.min":      "Harus pilih minimal 1 skill",
		"Skills.required": "Harus pilih minimal 1 skill",
	}
}
