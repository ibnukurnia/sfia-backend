package responses

type AdminDujResponse struct {
	Id   string `json:"id" gorm:"column:uuid"`
	JobDescription string `json:"job_description" gorm:"column:job_description"`
	Roles []RoleItem `json:"roles"`
	Skillset []SkillItem `json:"skills"`
}

type RoleItem struct {
	Id   string `json:"id" gorm:"column:role_id"`
	Name string `json:"name" gorm:"column:role_name"`
}

type SkillItem struct {
	Id   string `json:"id" gorm:"column:skills_id"`
	Name string `json:"name" gorm:"column:skills_name"`
}
