package response_management

type ChartData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func DummyChartData() []ChartData {
	return []ChartData{
		{Name: "Programmer", Value: 150},
		{Name: "SAD", Value: 130},
		{Name: "System Engineer", Value: 130},
		{Name: "Management", Value: 100},
		{Name: "QA Tester", Value: 100},
		{Name: "Business Analyst", Value: 110},
		{Name: "QA Analyst", Value: 110},
		{Name: "Tester Leader", Value: 100},
		{Name: "DevOps Engineer", Value: 90},
		{Name: "Scrum Master", Value: 90},
		{Name: "IT Governance", Value: 80},
		{Name: "Admin", Value: 80},
		{Name: "Problem Analyst", Value: 70},
		{Name: "UI/UX Designer", Value: 70},
		{Name: "IT Human Capital", Value: 60},
		{Name: "Operator", Value: 60},
		{Name: "IT Solutions", Value: 55},
		{Name: "Technical Writer", Value: 55},
		{Name: "DB Admin", Value: 50},
		{Name: "Senior Engineer", Value: 50},
		{Name: "Liaison Officer", Value: 45},
		{Name: "Performance Analyst", Value: 40},
		{Name: "IT Procurement", Value: 35},
	}
}

type CountKomposisiData struct {
	Title string `json:"title"`
	Image string `json:"image"`
	Count int    `json:"count"`
	Color string `json:"color"`
}

func DummyCountKomposisiData() []CountKomposisiData {
	return []CountKomposisiData{
		{Title: "Role Utama", Image: "role_utama", Count: 126, Color: "#FF5A5A"},
		{Title: "Role Utama & Role Minat", Image: "role_utama_minat", Count: 81, Color: "#3D9BE9"},
		{Title: "Role Utama & Role Tambahan", Image: "role_utama_tambahan", Count: 110, Color: "#FDCB13"},
		{Title: "Role Utama & Role Tambahan + Role Minat", Image: "role_utama_tambahan_minat", Count: 111, Color: "#12C293"},
	}
}

type SampleData struct {
	Name       string  `json:"name"`
	Admin      float64 `json:"Admin"`
	Business   float64 `json:"Business"`
	DBAdmin    float64 `json:"DBAdmin"`
	Programmer float64 `json:"Programmer"`
}

func DummySampleData() []SampleData {
	return []SampleData{
		{Name: "1 th", Admin: 15, Business: 2, DBAdmin: 1, Programmer: 8},
		{Name: "1.5 th", Admin: 13, Business: 2, DBAdmin: 1.5, Programmer: 8},
		{Name: "2 th", Admin: 12, Business: 2, DBAdmin: 2, Programmer: 8},
		{Name: "2.5 th", Admin: 11, Business: 2, DBAdmin: 2, Programmer: 8},
		{Name: "3 th", Admin: 10, Business: 2, DBAdmin: 2, Programmer: 8},
		{Name: "3.5 th", Admin: 9.5, Business: 2, DBAdmin: 2.5, Programmer: 7.5},
		{Name: "4 th", Admin: 9, Business: 2, DBAdmin: 3, Programmer: 7},
		{Name: "4.5 th", Admin: 8, Business: 2, DBAdmin: 3, Programmer: 6.5},
		{Name: "5 th", Admin: 7, Business: 2, DBAdmin: 3, Programmer: 6},
		{Name: "5.5 th", Admin: 6, Business: 2, DBAdmin: 3, Programmer: 5.5},
		{Name: "6 th", Admin: 5, Business: 2, DBAdmin: 3, Programmer: 5},
		{Name: "6.5 th", Admin: 4.5, Business: 2, DBAdmin: 3, Programmer: 4.5},
		{Name: "7 th", Admin: 4, Business: 2, DBAdmin: 3, Programmer: 4},
		{Name: "7.5 th", Admin: 3.5, Business: 2, DBAdmin: 3, Programmer: 3.5},
		{Name: "8 th", Admin: 3, Business: 2, DBAdmin: 3, Programmer: 3},
		{Name: ">8 th", Admin: 1, Business: 1, DBAdmin: 1, Programmer: 3},
	}
}
