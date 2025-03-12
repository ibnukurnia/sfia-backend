package requests

import "github.com/google/uuid"

type AddAdminDujRequest struct {
	JobDescription        string   `json:"job_description" validate:"required"`
	Roles []uuid.UUID `json:"roles"`
	Skills []uuid.UUID `json:"skills"`
}

func (r AddAdminDujRequest) Messages() map[string]string {
	return map[string]string{
		"JobDescription.required": "Nama harus diisi",
	}
}

type UpdateAdminDujRequest struct {
	DujId      string    `json:"id" validate:"required"`
	JobDescription        string   `json:"job_description" validate:"required"`
	Roles []uuid.UUID `json:"roles"`
	Skills []uuid.UUID `json:"skills"`
}

func (r UpdateAdminDujRequest) Messages() map[string]string {
	return map[string]string{
		"DujId.required": "DUJ id harus diisi",
		"JobDescription.required": "Nama harus diisi",
	}
}