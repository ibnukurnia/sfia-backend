package models

const (
	ADMIN   RoleAccess = "admin"
	USER    RoleAccess = "user"
	MANAGER RoleAccess = "manager"
)

type RoleAccess string

type User struct {
	Base
	Name       string
	Email      string
	Pn         string
	Password   string
	RoleAccess RoleAccess
}
