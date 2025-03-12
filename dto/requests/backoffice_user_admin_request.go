package requests

import "github.com/google/uuid"

type UpdateUserAdminRequest struct {
	Id		  uuid.UUID `json:"id" validate:"required"`
	Role 	  string `json:"role" validate:"required"`
}

func (r UpdateUserAdminRequest) Messages() map[string]string {
	return map[string]string{
		"Id.required": "User id harus diisi",
		"Role.required": "Role harus diisi",
	}
}

