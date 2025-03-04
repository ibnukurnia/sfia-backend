package responses

// type SfiaResponse struct {
// 	Skill     string              `json:"skill"`
// 	Questions []RoleSkillQuestion `json:"questions"`
// }

type RoleSkillQuestion struct {
	Id              string                    `json:"id"`
	Question        string                    `json:"question"`
	Answers         []RoleSkillQuestionAnswer `json:"answers"`
	CurrentAnswerId *string                   `json:"current_answer_id,omitempty"`
}

type RoleSkillQuestionAnswer struct {
	Id     string `json:"id"`
	Answer string `json:"answer"`
}

// func ToResponseSfia(skills []models.Skill, currentAnswer map[string]string) []SfiaResponse {
// 	res := []SfiaResponse{}

// 	for _, skill := range skills {
// 		questions := []RoleSkillQuestion{}

// 		for _, question := range skill.SfiaQuestions {
// 			answers := []RoleSkillQuestionAnswer{}

// 			// for _, answer := range question.Answers {
// 			// 	answers = append(answers, RoleSkillQuestionAnswer{
// 			// 		Id:     answer.Uuid.String(),
// 			// 		Answer: answer.Answer,
// 			// 	})
// 			// }

// 			rs := RoleSkillQuestion{
// 				Id:       question.Uuid.String(),
// 				Question: question.Question,
// 				Answers:  answers,
// 			}

// 			if answerId, exist := currentAnswer[question.Uuid.String()]; exist {
// 				rs.CurrentAnswerId = &answerId
// 			}

// 			questions = append(questions, rs)
// 		}

// 		res = append(res, SfiaResponse{
// 			Skill:     fmt.Sprintf("%s (%s)", skill.Name, skill.Code),
// 			Questions: questions,
// 		})
// 	}

// 	return res
// }

type DujResponse struct {
	Job             string  `json:"job"`
	Detail          string  `json:"detail"`
	HaveTrouble     bool    `json:"have_trouble,omitempty"`
	TroubleCause    *string `json:"trouble_cause,omitempty"`
	CurrentJob      bool    `json:"current_job"`
	CurrentAnswerId *string `json:"current_answer_id,omitempty"`
}

type DujQuestion struct {
	Id       string `json:"id"`
	Question string `json:"question"`
}
