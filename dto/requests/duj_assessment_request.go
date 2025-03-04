package requests

type DujAssessmentRequest struct {
	Jobs []job `json:"jobs"`
}

type job struct {
	Name         string  `json:"name"`
	CurrentJob   bool    `json:"current_job"`
	HaveTrouble  bool    `json:"have_trouble"`
	TroubleCause *string `json:"trouble_cause"`
	Id           *string `json:"id"`
}

func (r *DujAssessmentRequest) Messages() map[string]string {
	return map[string]string{}
}
