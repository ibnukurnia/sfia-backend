package responses

type DepartmentResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DepartmentTeamResponse struct {
	Id   string `json:"id" gorm:"column:uuid"`
	DepartmentId string `json:"department_id"`
	DepartmentName string `json:"department_name,omitempty" gorm:"column:department_name"`
	Name string `json:"name"`
}

type DepartmentRoleResponse struct {
	Id   string `json:"id" gorm:"column:uuid"`
	DepartmentId string `json:"department_id"`
	DepartmentName string `json:"department_name,omitempty" gorm:"column:name"`
	Name string `json:"name"`
}
