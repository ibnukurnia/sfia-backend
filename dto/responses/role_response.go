package responses

type GroupedRolesResponse struct {
	Name  string         `json:"name"`
	Roles []RoleResponse `json:"roles"`
}

type RoleResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
