package requests

type RegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Pn       string `json:"pn" validate:"required"`
}

func (r RegisterRequest) Messages() map[string]string {
	return map[string]string{
		"Email.required":    "Email harus diisi",
		"Password.required": "Password harus diisi",
		"Name.required":     "Nama harus diisi",
		"Pn.required":       "Pn harus diisi",
	}
}
