package responses

type UserAdminResponse struct {
	Id   string `json:"id" gorm:"column:uuid"`
	Name string `json:"name" gorm:"column:name"`
	Role string `json:"role" gorm:"column:role_name"`
	IdKaryawan string `json:"id_karyawan" gorm:"column:pn"`
	CorporateTitle string `json:"corporation_title" gorm:"column:corporation_title"`
	RoleAccess string `json:"role_access" gorm:"column:role_access"`
}