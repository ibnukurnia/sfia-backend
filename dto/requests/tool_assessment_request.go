package requests

type ToolAssessmentRequest struct {
	Tools          []ToolAssessmentAnswer `json:"tools"`
	DeletedToolIds []string               `json:"deleted_tool_ids"`
}

type ToolAssessmentAnswer struct {
	Name     string  `json:"name"`
	Level    string  `json:"level"`
	Evidence string  `json:"evidence"`
	AnswerId *string `json:"id"`
}

func (r *ToolAssessmentRequest) Messages() map[string]string {
	return map[string]string{}
}
