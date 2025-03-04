package requests

type SelfAssessmentRequest struct {
	Answers []SelfAssessmentAnswer `json:"answers"`
}

type SelfAssessmentAnswer struct {
	// RoleId     string  `json:"role_id"`
	// SkillId    string  `json:"skill_id"`
	QuestionId string  `json:"question_id"`
	Value      int8    `json:"value"`
	Evidence   string  `json:"evidence"`
	AnswerId   *string `json:"answer_id"`
}

func (r *SelfAssessmentRequest) Messages() map[string]string {
	return map[string]string{}
}
