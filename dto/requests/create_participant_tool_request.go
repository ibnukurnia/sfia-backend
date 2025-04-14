package requests

type CreateParticipantToolRequest struct {
	Tools []ToolRequest `json:"tools"`
}

type ToolRequest struct {
	Name     string  `json:"name"`
	Id       *string `json:"id"`
	Level    string  `json:"level"`
	Evidence string  `json:"evidence"`
}

func (r *CreateParticipantToolRequest) Messages() map[string]string {
	return map[string]string{}
}
