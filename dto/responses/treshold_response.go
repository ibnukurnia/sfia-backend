package responses

type TresholdResponse struct {
	RoleLevel  []TresholdItem `json:"role_level"`
	SkillLevel []TresholdItem `json:"skill_level"`
}

type TresholdItem struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	TresholdFrom float32 `json:"treshold_from"`
	TresholdTo   float32 `json:"treshold_to"`
	Color        string  `json:"color"`
}

type SkillTresholdResponse struct {
	Basic        float32 `json:"basic"`
	Intermediate float32 `json:"intermediate"`
	Advance      float32 `json:"advance"`
}
