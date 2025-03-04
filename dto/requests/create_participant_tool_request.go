package requests

type CreateParticipantToolRequest struct {
	Tools []ToolRequest `json:"tools"`
}

type ToolRequest struct {
	Tool  string `json:"tool"`
	Level string `json:"level"`
}

func (r *CreateParticipantToolRequest) Messages() map[string]string {
	return map[string]string{}
}
