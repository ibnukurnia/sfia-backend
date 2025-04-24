package response_management

type ManagementAplikasiResponse struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummyManagementAplikasiResponse() []ManagementAplikasiResponse {
	return []ManagementAplikasiResponse{
		{Name: "BRISPOT", Value: 100},
		{Name: "NDS", Value: 90},
		{Name: "RECON", Value: 80},
		{Name: "BRIMO", Value: 75},
		{Name: "BRILINK", Value: 70},
		{Name: "BRINETs", Value: 65},
		{Name: "Jira", Value: 60},
		{Name: "BRICaMS", Value: 55},
		{Name: "BRISURF", Value: 50},
		{Name: "Proswitching", Value: 45},
		{Name: "BRILINK", Value: 30},
		{Name: "BRINETs", Value: 25},
		{Name: "Jira", Value: 20},
		{Name: "BRICaMS", Value: 15},
		{Name: "BRISURF", Value: 10},
		{Name: "Proswitching", Value: 10},
	}
}
