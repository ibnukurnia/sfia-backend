package requests

type CreateParticipantUpdatedTrainingRequest struct {
	Trainings []UpdatedTraining `json:"trainings"`
}

type UpdatedTraining struct {
	Name             string `json:"name"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	Implementation   string `json:"implementation"`
	Location         string `json:"location"`
	HasCertification bool   `json:"has_certification"`
	GetCertification bool   `json:"get_certification"`
	Provider         string `json:"provider"`
}

func (r CreateParticipantUpdatedTrainingRequest) Messages() map[string]string {
	return map[string]string{}
}
