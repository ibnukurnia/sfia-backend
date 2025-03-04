package responses

type SkillResponse struct {
	Id         string `json:"id"`
	Skill      string `json:"skill"`
	IsMastered bool   `json:"is_mastered"`
	UsedFor    int    `json:"used_for"`
}
