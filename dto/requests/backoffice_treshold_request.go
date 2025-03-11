package requests

import "github.com/google/uuid"

type AddTreshold struct {
	Name        string `json:"name" validate:"required"`
	Category 	string `json:"category" validate:"required"`
	TresholdFrom float32 `json:"treshold_from"`
	TresholdTo float32 `json:"treshold_to"`
	Color string `json:"color" validate:"required"`
}

func (r AddTreshold) Messages() map[string]string {
	return map[string]string{
		"Name.required": 				"Nama harus diisi",
		"Category.required": 			"Kategori harus diisi",
		"Color.required":        		"Color harus diisi",
	}
}

type UpdateTreshold struct {
	TresholdId 		uuid.UUID `json:"id" validate:"required,uuid"`
	Name        string `json:"name" validate:"required"`
	Category 	string `json:"category"`
	TresholdFrom float32 `json:"treshold_from"`
	TresholdTo float32 `json:"treshold_to" validate:"required"`
	Color string `json:"color" validate:"required"`
}

func (r UpdateTreshold) Messages() map[string]string {
	return map[string]string{
		"TresholdId.required": 			"Id treshold harus diisi",
		"Name.required": 				"Nama harus diisi",
		"Category.required": 			"Kategori harus diisi",
		"Color.required":        		"Color harus diisi",
	}
}