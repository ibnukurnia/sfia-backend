package response_management

type RoleDataResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
	Color string `json:"color"`
}

func DummyRoleDataResponse() []RoleDataResponse {
	return []RoleDataResponse{
		{ID: 1, Name: "Senior", Value: 172, Color: "#24E5B2"},
		{ID: 2, Name: "Middle", Value: 149, Color: "#FCD400"},
		{ID: 3, Name: "Junior", Value: 262, Color: "#FF4848"},
	}
}

type SkillExistingRoleData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
	Color string `json:"color"`
}

func DummySkillExistingRoleDataResponse() []SkillExistingRoleData {
	return []SkillExistingRoleData{
		{ID: 1, Name: "Advanced", Value: 122, Color: "#24E5B2"},
		{ID: 2, Name: "Intermediate", Value: 100, Color: "#FCD400"},
		{ID: 3, Name: "Basic", Value: 270, Color: "#FF4848"},
	}
}

type KesulitanDUJData struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
	Color string `json:"color"`
}

func DummyKesulitanDUJDataResponse() []KesulitanDUJData {
	return []KesulitanDUJData{
		{ID: 1, Name: "Sulit", Value: 157, Color: "#1A46CB"},
		{ID: 2, Name: "Mudah", Value: 271, Color: "#78DFFF"},
	}
}
