package responses

type SkillResponse struct {
	Id         string `json:"id"`
	Skill      string `json:"skill"`
	IsMastered bool   `json:"is_mastered"`
	UsedFor    int    `json:"used_for"`
}

type SkillsetResponse struct {
	RoleId    string `json:"role_id" gorm:"column:role_id"`
	RoleName  string `json:"role_name" gorm:"column:role_name"`
	SkillId string `json:"skill_id" gorm:"column:skill_id"`
	SkillName string `json:"skill_name" gorm:"column:skill_name"`
	SkillDescription   string `json:"skill_description" gorm:"column:skill_description"`
}