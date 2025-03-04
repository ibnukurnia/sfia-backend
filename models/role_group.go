package models

type RoleGroup struct {
	Base
	Name  string
	Roles []Role `gorm:"foreignKey:GroupId"`
}
