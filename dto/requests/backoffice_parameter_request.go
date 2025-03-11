package requests

import "github.com/google/uuid"

type AddParameterScore struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
}

func (r AddParameterScore) Messages() map[string]string {
	return map[string]string{
		"Name.required": "Nama harus diisi",
		"Description.required": "Deskripsi harus diisi",
	}
}

type UpdateParameterScore struct {
	ParameterScoreId      uuid.UUID    `json:"id" validate:"required"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
}

func (r UpdateParameterScore) Messages() map[string]string {
	return map[string]string{
		"ParameterScoreId.required": "Parameter Score id harus diisi",
		"Name.required": "Nama harus diisi",
		"Description.required": "Deskripsi harus diisi",
	}
}

type AddParameterDifficulty struct {
	Description string   `json:"description" validate:"required"`
}

func (r AddParameterDifficulty) Messages() map[string]string {
	return map[string]string{
		"Description.required": "Deskripsi harus diisi",
	}
}

type UpdateParameterDifficulty struct {
	ParameterDifficultyId      uuid.UUID    `json:"id" validate:"required"`
	Description string    `json:"description" validate:"required"`
}

func (r UpdateParameterDifficulty) Messages() map[string]string {
	return map[string]string{
		"ParameterDifficultyId.required": "Parameter Difficulty id harus diisi",
		"Description.required": "Deskripsi harus diisi",
	}
}