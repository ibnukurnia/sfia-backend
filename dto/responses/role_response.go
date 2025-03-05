package responses

type GroupedRolesResponse struct {
	Name  string         `json:"name"`
	Roles []RoleResponse `json:"roles"`
}

type RoleResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RoleGroupResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RoleListResponse struct {
	RoleId    string `json:"role_id" gorm:"column:uuid"`
	RoleName  string `json:"role_name" gorm:"column:name"`
	GroupName string `json:"role_group_name" gorm:"column:group_name"`
	GroupId   string `json:"role_group_id" gorm:"column:group_id"`
}