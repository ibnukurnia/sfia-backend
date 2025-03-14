package requests

type AddToolsRequest struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url"`
}

func (r AddToolsRequest) Messages() map[string]string {
	return map[string]string{
		"Name.required": "Nama harus diisi",
	}
}

type UpdateToolsRequest struct {
	ToolsId string `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Url     string `json:"url"`
}

func (r UpdateToolsRequest) Messages() map[string]string {
	return map[string]string{
		"ToolsId.required": "Id tools harus diisi",
		"Name.required":    "Nama harus diisi",
	}
}