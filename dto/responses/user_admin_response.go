package responses

import "sv-sfia/utils"

type UserAdminResponse struct {
	Id             string `json:"id" gorm:"column:uuid"`
	Name           string `json:"name" gorm:"column:name"`
	Role           string `json:"role" gorm:"column:role_name"`
	IdKaryawan     string `json:"id_karyawan" gorm:"column:pn"`
	CorporateTitle string `json:"corporate_title" gorm:"column:corporation_title"`
	Department     string `json:"department" gorm:"column:department"`
	RoleAccess     string `json:"role_access" gorm:"column:role_access"`
}

type UserAdminPaginatedResponse struct {
	Data      []UserAdminResponse `json:"data"`
	Paginator utils.Paginator     `json:"paginator"`
}
