package requests

type CreateParticpantDujRequest struct {
	Jobs []DujJob `json:"jobs"`
}

type DujJob struct {
	Job             string  `json:"name"`
	Detail          string  `json:"detail"`
	CurrentJob      bool    `json:"current_job"`
	HaveTrouble     bool    `json:"have_trouble"`
	TroubleCause    *string `json:"trouble_cause"`
	CurrentAnswerId *string `json:"current_answer_id"`
}

func (r *CreateParticpantDujRequest) Messages() map[string]string {
	return map[string]string{}
}
