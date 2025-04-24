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
