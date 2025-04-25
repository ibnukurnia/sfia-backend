package requests

type AddCorporateTitleRequest struct {
	Name string `json:"name" validate:"required"`
}

func (r AddCorporateTitleRequest) Messages() map[string]string {
	return map[string]string{
		"Name.required": "Nama harus diisi",
	}
}

type UpdateCorporateTitleRequest struct {
	CorporateTitleId string `json:"corporate_title_id" validate:"required"`
	Name             string `json:"name" validate:"required"`
}

func (r UpdateCorporateTitleRequest) Messages() map[string]string {
	return map[string]string{
		"CorporateTitleId.required": "CorporateTitle id harus diisi",
		"Name.required":             "Nama harus diisi",
	}
}

type GetCorporateTitlesRequest struct {
	Search string `json:"search" form:"search"`
}

func (r GetCorporateTitlesRequest) Messages() map[string]string {
	return map[string]string{}
}
