package response_management

type PersebaranTipeRole struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Value1 int    `json:"value1"`
	Value2 int    `json:"value2"`
	Value3 int    `json:"value3"`
}

func DummyPersebaranRole() []PersebaranTipeRole {
	return []PersebaranTipeRole{
		{ID: 1, Name: "Problem Analyst", Value1: 8, Value2: 2, Value3: 0},
		{ID: 2, Name: "Business Analyst", Value1: 5, Value2: 1, Value3: 0},
		{ID: 3, Name: "System Engineer", Value1: 2, Value2: 0, Value3: 1},
		{ID: 4, Name: "Programmer", Value1: 1, Value2: 0, Value3: 0},
		{ID: 5, Name: "Management", Value1: 1, Value2: 0, Value3: 0},
		{ID: 6, Name: "IT Solutions", Value1: 0, Value2: 0, Value3: 1},
	}
}

type PersebaranSkillPenguasaanResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Value1 int    `json:"value1"`
	Value2 int    `json:"value2"`
	Value3 int    `json:"value3"`
}

func DummyPersebaranSkill() []PersebaranSkillPenguasaanResponse {
	return []PersebaranSkillPenguasaanResponse{
		{ID: 1, Name: "DBDS", Value1: 180, Value2: 40, Value3: 30},
		{ID: 2, Name: "SWDN", Value1: 140, Value2: 50, Value3: 20},
		{ID: 3, Name: "QUAS", Value1: 200, Value2: 40, Value3: 30},
		{ID: 4, Name: "ITSP", Value1: 160, Value2: 30, Value3: 20},
		{ID: 5, Name: "STPL", Value1: 190, Value2: 40, Value3: 40},
		{ID: 6, Name: "BSMO", Value1: 170, Value2: 30, Value3: 30},
		{ID: 7, Name: "BPRE", Value1: 140, Value2: 40, Value3: 30},
		{ID: 8, Name: "USUP", Value1: 130, Value2: 30, Value3: 20},
		{ID: 9, Name: "DBAD", Value1: 160, Value2: 30, Value3: 30},
		{ID: 10, Name: "ICPM", Value1: 80, Value2: 20, Value3: 10},
		{ID: 11, Name: "TMCR", Value1: 60, Value2: 20, Value3: 10},
		{ID: 12, Name: "ITCM", Value1: 40, Value2: 10, Value3: 5},
		{ID: 13, Name: "RELM", Value1: 30, Value2: 10, Value3: 5},
		{ID: 14, Name: "SINT", Value1: 100, Value2: 30, Value3: 20},
		{ID: 15, Name: "BURM", Value1: 140, Value2: 40, Value3: 30},
		{ID: 16, Name: "ASMG", Value1: 170, Value2: 40, Value3: 40},
		{ID: 17, Name: "RLMT", Value1: 160, Value2: 40, Value3: 30},
		{ID: 18, Name: "USEV", Value1: 150, Value2: 30, Value3: 30},
		{ID: 19, Name: "SLMO", Value1: 110, Value2: 30, Value3: 20},
		{ID: 20, Name: "HSIN", Value1: 80, Value2: 20, Value3: 10},
		{ID: 21, Name: "PEMT", Value1: 50, Value2: 10, Value3: 5},
		{ID: 22, Name: "COPL", Value1: 30, Value2: 5, Value3: 3},
		{ID: 23, Name: "ITCM", Value1: 10, Value2: 2, Value3: 1},
	}
}

type RelevansiTahunLevelChartResponse struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummyRelevansiTahunLevelChartResponse() []RelevansiTahunLevelChartResponse {
	return []RelevansiTahunLevelChartResponse{
		{Name: "DBDS", Value: 150},
		{Name: "SWDN", Value: 140},
		{Name: "QUAS", Value: 130},
		{Name: "ITSP", Value: 125},
		{Name: "STPL", Value: 120},
		{Name: "BSMO", Value: 115},
		{Name: "BPRE", Value: 110},
		{Name: "USUP", Value: 108},
		{Name: "DBAD", Value: 105},
		{Name: "ICPM", Value: 100},
		{Name: "TMCR", Value: 98},
		{Name: "ITCM", Value: 95},
		{Name: "RELM", Value: 90},
		{Name: "SINT", Value: 85},
		{Name: "BURM", Value: 80},
		{Name: "ASMG", Value: 78},
		{Name: "RLMT", Value: 75},
		{Name: "USEV", Value: 70},
		{Name: "SLMO", Value: 68},
		{Name: "HSIN", Value: 65},
		{Name: "PEMT", Value: 60},
		{Name: "COPL", Value: 55},
		{Name: "ITCM", Value: 50},
	}
}

type RelevansiTahunLevelTablesResponse struct {
	ID             int    `json:"id"`
	Nama           string `json:"nama"`
	TipeRole       string `json:"tipe_role"`
	RoleSekarang   string `json:"role_sekarang"`
	CorporateTitle string `json:"corporate_title"`
	Pengalaman     string `json:"pengalaman"`
}

func DummyRelevansiTahunLevelTableResponse() []RelevansiTahunLevelTablesResponse {
	return []RelevansiTahunLevelTablesResponse{
		{ID: 1, Nama: "A. A. Gde Agung Aditya Pratama", TipeRole: "Role Tambahan", RoleSekarang: "Senior", CorporateTitle: "Mitra Pekerja", Pengalaman: ">8 th"},
		{ID: 2, Nama: "Abdul Muttaqien", TipeRole: "Role Tambahan", RoleSekarang: "Middle", CorporateTitle: "Manager", Pengalaman: ">8 th"},
		{ID: 3, Nama: "Abiq Muhammad Faesal", TipeRole: "Role Tambahan", RoleSekarang: "Senior", CorporateTitle: "Manager", Pengalaman: ">8 th"},
		{ID: 4, Nama: "Achmad Affandi", TipeRole: "Role Tambahan", RoleSekarang: "Senior", CorporateTitle: "Mitra Pekerja", Pengalaman: ">8 th"},
		{ID: 5, Nama: "Achmad Fauzi", TipeRole: "Role Tambahan", RoleSekarang: "Junior", CorporateTitle: "Mitra Pekerja", Pengalaman: ">8 th"},
		{ID: 6, Nama: "Achmad Indra Fauzan", TipeRole: "Role Tambahan", RoleSekarang: "Senior", CorporateTitle: "AVP", Pengalaman: ">8 th"},
		{ID: 7, Nama: "Ade Rissa Heliza", TipeRole: "Role Tambahan", RoleSekarang: "Junior", CorporateTitle: "Mitra Pekerja", Pengalaman: ">8 th"},
	}
}

type UnMasteredSkillChartResponse struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummyUnMasteredSkillResponse() []UnMasteredSkillChartResponse {
	return []UnMasteredSkillChartResponse{
		{Name: "DBDS", Value: 150},
		{Name: "SWDN", Value: 140},
		{Name: "QUAS", Value: 130},
		{Name: "ITSP", Value: 125},
		{Name: "STPL", Value: 120},
		{Name: "BSMO", Value: 115},
		{Name: "BPRE", Value: 110},
		{Name: "USUP", Value: 108},
		{Name: "DBAD", Value: 105},
		{Name: "ICPM", Value: 100},
		{Name: "TMCR", Value: 98},
		{Name: "ITCM", Value: 95},
		{Name: "RELM", Value: 90},
		{Name: "SINT", Value: 85},
		{Name: "BURM", Value: 80},
		{Name: "ASMG", Value: 78},
		{Name: "RLMT", Value: 75},
		{Name: "USEV", Value: 70},
		{Name: "SLMO", Value: 68},
		{Name: "HSIN", Value: 65},
		{Name: "PEMT", Value: 60},
		{Name: "COPL", Value: 55},
		{Name: "ITCM", Value: 50},
	}
}

type UnMasteredSkillTableResponse struct {
	Name     string `json:"nama"`
	RoleType string `json:"tipe_role"`
	Role     string `json:"role"`
	KeySkill string `json:"skill_utama"`
}

func DummyUnMasteredSkillTablesReponse() []UnMasteredSkillTableResponse {
	return []UnMasteredSkillTableResponse{
		{
			Name:     "Arsya Wildan Rajendra",
			RoleType: "Role Utama",
			Role:     "DevOps Engineer",
			KeySkill: "Release and Deployment (RELM)",
		},
		{
			Name:     "Cahyo Yunan Adianto",
			RoleType: "Role Utama",
			Role:     "System Engineer",
			KeySkill: "Release and Deployment (RELM)",
		},
		{
			Name:     "Fitra Mefa Pratama",
			RoleType: "Role Utama",
			Role:     "Senior Engineer",
			KeySkill: "Release and Deployment (RELM)",
		},
		{
			Name:     "I Putu Indra Mahendra Priyadi",
			RoleType: "Role Utama",
			Role:     "DevOps Engineer",
			KeySkill: "Release and Deployment (RELM)",
		},
		{
			Name:     "Imam Chanafi",
			RoleType: "Role Utama",
			Role:     "Tester Leader",
			KeySkill: "Testing (TEST)",
		},
		{
			Name:     "Nurul Ayu Rahmawati",
			RoleType: "Role Utama",
			Role:     "System Engineer",
			KeySkill: "Release and Deployment (RELM)",
		},
		{
			Name:     "Rahmad Ramadhan",
			RoleType: "Role Utama",
			Role:     "System Engineer",
			KeySkill: "Release and Deployment (RELM)",
		},
		{
			Name:     "Wahyu Budiadi Wibowo",
			RoleType: "Role Utama",
			Role:     "System Engineer",
			KeySkill: "Release and Deployment (RELM)",
		},
	}
}

type YearLevelRoleMappingResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
	Color string `json:"color"`
}

func DummyDataRoleMappingResponse() []YearLevelRoleMappingResponse {
	return []YearLevelRoleMappingResponse{
		{ID: 1, Name: "Certified", Value: 75, Color: "#C432D0"},
		{ID: 2, Name: "Advanced", Value: 100, Color: "#24E5B2"},
		{ID: 3, Name: "Intermediate", Value: 100, Color: "#FCD400"},
		{ID: 4, Name: "Basic", Value: 125, Color: "#FF4848"},
	}
}

type YearLevelDistributionResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Value1 int    `json:"value1"`
	Value2 int    `json:"value2"`
	Value3 int    `json:"value3"`
}

func DummyDataYearLevelDistribution() []YearLevelDistributionResponse {
	return []YearLevelDistributionResponse{
		{ID: 1, Name: "1 th", Value1: 8, Value2: 5, Value3: 20},
		{ID: 2, Name: "2 th", Value1: 6, Value2: 6, Value3: 18},
		{ID: 3, Name: "3 th", Value1: 5, Value2: 5, Value3: 15},
		{ID: 4, Name: "4 th", Value1: 4, Value2: 4, Value3: 13},
		{ID: 5, Name: "5 th", Value1: 2, Value2: 3, Value3: 9},
		{ID: 6, Name: "6 th", Value1: 2, Value2: 2, Value3: 10},
		{ID: 7, Name: "7 th", Value1: 2, Value2: 3, Value3: 8},
		{ID: 8, Name: "8 th", Value1: 1, Value2: 1, Value3: 3},
		{ID: 9, Name: ">8 th", Value1: 2, Value2: 2, Value3: 5},
	}
}

type SkillRequirementResponse struct {
	KodeSkill string `json:"kode_skill"`
	ReqJunior string `json:"req_junior"`
	ReqMiddle string `json:"req_middle"`
	ReqSenior string `json:"req_senior"`
}

func DummyDataSkillRequirements() []SkillRequirementResponse {
	return []SkillRequirementResponse{
		{KodeSkill: "ADMN", ReqJunior: "Basic", ReqMiddle: "Intermediate", ReqSenior: "Intermediate"},
		{KodeSkill: "ARCH", ReqJunior: "", ReqMiddle: "", ReqSenior: "Intermediate"},
		{KodeSkill: "ASUP", ReqJunior: "Basic", ReqMiddle: "Intermediate", ReqSenior: ""},
		{KodeSkill: "AUDT", ReqJunior: "Basic", ReqMiddle: "Intermediate", ReqSenior: "Advance"},
		{KodeSkill: "AVMT", ReqJunior: "Basic", ReqMiddle: "Intermediate", ReqSenior: ""},
		{KodeSkill: "AVMT", ReqJunior: "", ReqMiddle: "Intermediate", ReqSenior: "Advance"},
		{KodeSkill: "BPRE", ReqJunior: "Basic", ReqMiddle: "Intermediate", ReqSenior: ""},
		{KodeSkill: "BPTS", ReqJunior: "Intermediate", ReqMiddle: "", ReqSenior: "Advance"},
		{KodeSkill: "BPTS", ReqJunior: "Intermediate", ReqMiddle: "", ReqSenior: "Advance"},
		{KodeSkill: "BURM", ReqJunior: "Intermediate", ReqMiddle: "Basic", ReqSenior: "Intermediate"},
		{KodeSkill: "BURM", ReqJunior: "", ReqMiddle: "Basic", ReqSenior: ""},
		{KodeSkill: "CHMG", ReqJunior: "Basic", ReqMiddle: "", ReqSenior: "Intermediate"},
		{KodeSkill: "CHMG", ReqJunior: "Basic", ReqMiddle: "Intermediate", ReqSenior: ""},
		{KodeSkill: "CHMG", ReqJunior: "Basic", ReqMiddle: "Basic", ReqSenior: "Basic"},
	}
}

type CrossRoleRecommendationResponse struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummyDataCrossRoleRecommendations() []CrossRoleRecommendationResponse {
	return []CrossRoleRecommendationResponse{
		{Name: "DBDS", Value: 150},
		{Name: "SWDN", Value: 140},
		{Name: "QUAS", Value: 130},
		{Name: "ITSP", Value: 125},
		{Name: "STPL", Value: 120},
		{Name: "BSMO", Value: 115},
		{Name: "BPRE", Value: 110},
		{Name: "USUP", Value: 108},
		{Name: "DBAD", Value: 105},
		{Name: "ICPM", Value: 100},
		{Name: "TMCR", Value: 98},
		{Name: "ITCM", Value: 95},
		{Name: "RELM", Value: 90},
		{Name: "SINT", Value: 85},
		{Name: "BURM", Value: 80},
		{Name: "ASMG", Value: 78},
		{Name: "RLMT", Value: 75},
		{Name: "USEV", Value: 70},
		{Name: "SLMO", Value: 68},
		{Name: "HSIN", Value: 65},
		{Name: "PEMT", Value: 60},
		{Name: "COPL", Value: 55},
		{Name: "ITCM", Value: 50},
	}
}

type CrossRoleRecommendationTableResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	RoleType       string `json:"role_type"`
	CurrentRole    string `json:"current_role"`
	CorporateTitle string `json:"corporate_title"`
	Experience     string `json:"experience"`
}

func DummyDataCrossRoleRecommendationTable() []CrossRoleRecommendationTableResponse {
	return []CrossRoleRecommendationTableResponse{
		{ID: 1, Name: "A. A. Gde Agung Aditya Pratama", RoleType: "Additional Role", CurrentRole: "Senior", CorporateTitle: "Mitra Pekerja", Experience: ">8 th"},
		{ID: 2, Name: "Abdul Muttaqien", RoleType: "Additional Role", CurrentRole: "Middle", CorporateTitle: "Manager", Experience: ">8 th"},
		{ID: 3, Name: "Abiq Muhammad Faesal", RoleType: "Additional Role", CurrentRole: "Senior", CorporateTitle: "Manager", Experience: ">8 th"},
		{ID: 4, Name: "Achmad Affandi", RoleType: "Additional Role", CurrentRole: "Senior", CorporateTitle: "Mitra Pekerja", Experience: ">8 th"},
		{ID: 5, Name: "Achmad Fauzi", RoleType: "Additional Role", CurrentRole: "Junior", CorporateTitle: "Mitra Pekerja", Experience: ">8 th"},
		{ID: 6, Name: "Achmad Indra Fauzan", RoleType: "Additional Role", CurrentRole: "Senior", CorporateTitle: "AVP", Experience: ">8 th"},
		{ID: 7, Name: "Ade Rissa Heliza", RoleType: "Additional Role", CurrentRole: "Junior", CorporateTitle: "Mitra Pekerja", Experience: ">8 th"},
	}
}

type CrossSkillRecommendationChartResponse struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummyDataCrossSkillRecommendationChart() []CrossSkillRecommendationChartResponse {
	return []CrossSkillRecommendationChartResponse{
		{Name: "DBDS", Value: 150},
		{Name: "SWDN", Value: 140},
		{Name: "QUAS", Value: 130},
		{Name: "ITSP", Value: 125},
		{Name: "STPL", Value: 120},
		{Name: "BSMO", Value: 115},
		{Name: "BPRE", Value: 110},
		{Name: "USUP", Value: 108},
		{Name: "DBAD", Value: 105},
		{Name: "ICPM", Value: 100},
		{Name: "TMCR", Value: 98},
		{Name: "ITCM", Value: 95},
		{Name: "RELM", Value: 90},
		{Name: "SINT", Value: 85},
		{Name: "BURM", Value: 80},
		{Name: "ASMG", Value: 78},
		{Name: "RLMT", Value: 75},
		{Name: "USEV", Value: 70},
		{Name: "SLMO", Value: 68},
		{Name: "HSIN", Value: 65},
		{Name: "PEMT", Value: 60},
		{Name: "COPL", Value: 55},
		{Name: "ITCM", Value: 50},
	}
}

type CrossSkillRecommendationTableResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	RoleType       string `json:"role_type"`
	CurrentRole    string `json:"current_role"`
	CorporateTitle string `json:"corporate_title"`
	Experience     string `json:"experience"`
}

func DummyDataCrossSkillRecommendationTable() []CrossSkillRecommendationTableResponse {
	return []CrossSkillRecommendationTableResponse{
		{ID: 1, Name: "A. A. Gde Agung Aditya Pratama", RoleType: "Additional Role", CurrentRole: "Senior", CorporateTitle: "Employee Partner", Experience: ">8 th"},
		{ID: 2, Name: "Abdul Muttaqien", RoleType: "Additional Role", CurrentRole: "Middle", CorporateTitle: "Manager", Experience: ">8 th"},
		{ID: 3, Name: "Abiq Muhammad Faesal", RoleType: "Additional Role", CurrentRole: "Senior", CorporateTitle: "Manager", Experience: ">8 th"},
		{ID: 4, Name: "Achmad Affandi", RoleType: "Additional Role", CurrentRole: "Senior", CorporateTitle: "Employee Partner", Experience: ">8 th"},
		{ID: 5, Name: "Achmad Fauzi", RoleType: "Additional Role", CurrentRole: "Junior", CorporateTitle: "Employee Partner", Experience: ">8 th"},
		{ID: 6, Name: "Achmad Indra Fauzan", RoleType: "Additional Role", CurrentRole: "Senior", CorporateTitle: "AVP", Experience: ">8 th"},
		{ID: 7, Name: "Ade Rissa Heliza", RoleType: "Additional Role", CurrentRole: "Junior", CorporateTitle: "Employee Partner", Experience: ">8 th"},
	}
}
