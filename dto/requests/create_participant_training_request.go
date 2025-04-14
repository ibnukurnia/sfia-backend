package requests

type CreateParticipantTrainingRequest struct {
	Trainings []Training `json:"trainings"`
}

type Training struct {
	Id                *string `json:"id"`
	Name              string  `json:"name"`
	NeedCertification bool    `json:"need_certification"`
	IsNeeded          bool    `json:"is_needed"`
	Priority          *int8   `json:"priority"`
	RoleId            *string `json:"role_id"`
}

func (r CreateParticipantTrainingRequest) Messages() map[string]string {
	return map[string]string{}
}
