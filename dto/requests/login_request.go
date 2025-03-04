package requests

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r LoginRequest) Messages() map[string]string {
	return map[string]string{
		"Email.required":    "Email harus diisi",
		"Password.required": "Password harus diisi",
	}
}
