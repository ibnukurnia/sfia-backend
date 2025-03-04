package requests

type PrecheckParticipantRequest struct {
	Pn    string `json:"pn" validate:"required"`
	Email string `json:"email" validate:"required"`
	Name  string `json:"name" validate:"required"`
}

func (r *PrecheckParticipantRequest) Messages() map[string]string {
	return map[string]string{
		"Pn.required":    "Nomor Pn tidak boleh kosong",
		"Email.required": "Email tidak boleh kosong",
		"Name.required":  "Nama tidak boleh kosong",
	}
}
