package requests

import (
	"errors"

	"github.com/google/uuid"
)

type UpdateUserAdminRequest struct {
	Id   uuid.UUID `json:"id" validate:"required"`
	Role string    `json:"role" validate:"required,oneof=user admin"`
}

func (r UpdateUserAdminRequest) Messages() map[string]string {
	return map[string]string{
		"Id.required":   "User id harus diisi",
		"Role.required": "Role harus diisi",
		"Role.oneof":    "Role tidak sesuai dengan ketentuan",
	}
}

func (r UpdateUserAdminRequest) Validate() error {
	if r.Role != "user" && r.Role != "admin" {
		return errors.New("Role harus bernilai 'user' atau 'admin'")
	}
	return nil
}

type UserAdminRequest struct {
	Search        string   `json:"search" form:"search"`
	Page          int      `json:"page" form:"page" binding:"required,min=1"`
	Limit         int      `json:"limit" form:"limit" binding:"required,min=1,max=100"`
	CorporateIDs  []string `json:"corporate_ids" form:"corporate_ids"`
	RoleIDs       []string `json:"role_ids" form:"role_ids"`
	DepartmentIDs []string `json:"department_ids" form:"department_ids"`
}
