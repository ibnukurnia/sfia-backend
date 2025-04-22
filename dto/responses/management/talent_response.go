package response_management

type StatusTalentResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
	Color string `json:"color"`
}

func DummyStatusTalentResponse() []StatusTalentResponse {
	res := []StatusTalentResponse{
		{ID: 1, Name: "Organik", Value: 100, Color: "#1A46CB"},
		{ID: 2, Name: "Insource", Value: 100, Color: "#24E5B2"},
		{ID: 3, Name: "Outsource", Value: 75, Color: "#C432D0"},
		{ID: 4, Name: "Mitra Pekerja", Value: 125, Color: "#FF4848"},
		{ID: 5, Name: "Kontrak", Value: 100, Color: "#FCD400"},
		{ID: 6, Name: "Lainnya", Value: 50, Color: "#CACACA"},
	}
	return res
}

type ManagementTalentDepartmentData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Color string `json:"color"`
}

type ManagementTalentResponse struct {
	MaxValue int                              `json:"maxValue"`
	List     []ManagementTalentDepartmentData `json:"list"`
}

func DummyDepartementTalentResponse() ManagementTalentResponse {
	res := ManagementTalentResponse{
		MaxValue: 100,
		List: []ManagementTalentDepartmentData{
			{Name: "Application Quality Improvement Department (QIM)", Value: 90, Color: "#1A46CB"},
			{Name: "Application Operation Services Department (OPA)", Value: 80, Color: "#FF4848"},
			{Name: "Core and Support Application Development Department (CSD)", Value: 60, Color: "#24E5B2"},
			{Name: "Omni Channel Application Development Department (OCD)", Value: 45, Color: "#FCD400"},
			{Name: "SuperApps Development Department (SAP)", Value: 30, Color: "#C432D0"},
		},
	}
	return res
}

type FunctionResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummyFunctionResponse() []FunctionResponse {
	return []FunctionResponse{
		{ID: 1, Name: "SDK", Value: 10},
		{ID: 2, Name: "CBA", Value: 9},
		{ID: 3, Name: "CDP", Value: 7},
		{ID: 4, Name: "CCP", Value: 7},
		{ID: 5, Name: "CBP", Value: 5},
		{ID: 6, Name: "MFP", Value: 4},
		{ID: 7, Name: "MBP", Value: 3},
		{ID: 8, Name: "AES", Value: 3},
		{ID: 9, Name: "LCP", Value: 2},
		{ID: 10, Name: "INA", Value: 2},
		{ID: 11, Name: "APQ", Value: 2},
		{ID: 12, Name: "EBA", Value: 2},
		{ID: 13, Name: "IBO", Value: 1},
		{ID: 14, Name: "MDO", Value: 1},
		{ID: 15, Name: "GDP", Value: 1},
		{ID: 16, Name: "ORP", Value: 1},
		{ID: 17, Name: "DAO", Value: 1},
		{ID: 18, Name: "WDP", Value: 1},
		{ID: 19, Name: "COE", Value: 1},
		{ID: 20, Name: "MBP", Value: 1},
		{ID: 21, Name: "AES", Value: 1},
		{ID: 22, Name: "LCP", Value: 1},
		{ID: 23, Name: "INA", Value: 1},
		{ID: 24, Name: "APQ", Value: 1},
		{ID: 25, Name: "EBA", Value: 1},
		{ID: 26, Name: "MDO", Value: 1},
	}
}

type TeamResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummyTeamResponse() []TeamResponse {
	return []TeamResponse{
		{ID: 1, Name: "Tim QC SAP", Value: 10},
		{ID: 2, Name: "Tim QC ITP CMP", Value: 9},
		{ID: 3, Name: "Tim QC MDP ESB", Value: 7},
		{ID: 4, Name: "Team Leader", Value: 7},
		{ID: 5, Name: "PQM", Value: 5},
		{ID: 6, Name: "Tim Test Automation", Value: 4},
		{ID: 7, Name: "APQ", Value: 3},
		{ID: 8, Name: "Team Member AES", Value: 3},
		{ID: 9, Name: "Tim QA CPA", Value: 1},
		{ID: 10, Name: "QA", Value: 1},
		{ID: 11, Name: "Tim QA Core Banking", Value: 1},
		{ID: 12, Name: "TIM QC RTGS BI", Value: 1},
		{ID: 13, Name: "Tim Brizzi", Value: 1},
		{ID: 14, Name: "Tim BTE Technical CBA", Value: 1},
		{ID: 15, Name: "Tim QC Remittance", Value: 1},
		{ID: 16, Name: "Tim QC Remittance", Value: 1},
		{ID: 17, Name: "Tim AES", Value: 1},
	}
}

type CorporateTitleResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummyCorporateTitleResponse() []CorporateTitleResponse {
	return []CorporateTitleResponse{
		{ID: 1, Name: "Assistant Manager", Value: 38},
	}
}

type SpecializationResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummySpecializationResponse() []SpecializationResponse {
	return []SpecializationResponse{
		{ID: 1, Name: "Core Banking", Value: 38},
		{ID: 2, Name: "Backend Developer", Value: 32},
		{ID: 3, Name: "Rekonsiliasi", Value: 30},
		{ID: 4, Name: "Tester", Value: 25},
		{ID: 5, Name: "Middleware", Value: 22},
		{ID: 6, Name: "System Engineer", Value: 18},
		{ID: 7, Name: "Backend", Value: 15},
		{ID: 8, Name: "UI/UX Designer", Value: 10},
	}
}

type YearOfExperienceResponse struct {
	ID               int    `json:"id"`
	YearOfExperience string `json:"year_of_experience"`
	Value1           int    `json:"value1"`
	Value2           int    `json:"value2"`
	Value3           int    `json:"value3"`
}

func DummyYearOfExperienceResponse() []YearOfExperienceResponse {
	return []YearOfExperienceResponse{
		{ID: 1, YearOfExperience: "1", Value1: 8, Value2: 5, Value3: 20},
		{ID: 2, YearOfExperience: "2", Value1: 6, Value2: 6, Value3: 18},
		{ID: 3, YearOfExperience: "3", Value1: 5, Value2: 5, Value3: 15},
		{ID: 4, YearOfExperience: "4", Value1: 4, Value2: 4, Value3: 13},
		{ID: 5, YearOfExperience: "5", Value1: 2, Value2: 3, Value3: 9},
		{ID: 6, YearOfExperience: "6", Value1: 2, Value2: 2, Value3: 10},
		{ID: 7, YearOfExperience: "7", Value1: 2, Value2: 3, Value3: 8},
		{ID: 8, YearOfExperience: "8", Value1: 1, Value2: 1, Value3: 3},
		{ID: 9, YearOfExperience: ">8", Value1: 2, Value2: 2, Value3: 5},
	}
}
