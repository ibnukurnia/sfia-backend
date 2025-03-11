package responses

type ParameterResponse struct {
	ParameterScore []ParameterScoreItem `json:"parameter_score"`
	ParameterDifficulty []ParameterDifficultyItem `json:"parameter_difficulties"`
}

type ParameterScoreItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type ParameterDifficultyItem struct {
	Id   string `json:"id"`
	Description string `json:"description"`
}