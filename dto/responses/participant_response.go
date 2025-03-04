package responses

type AuthParticipantResponse struct {
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	Pn                    string `json:"pn"`
	IsExist               bool   `json:"is_exist"`
	AssessmentToken       string `json:"assessment_token"`
	IsDepartmentCompleted bool   `json:"is_department_completed"`
	IsRoleCompleted       bool   `json:"is_role_completed"`
}

type ParticipantRegistrationStatusResponse struct {
	IsDepartmentCompleted bool `json:"is_department_completed"`
	IsRoleCompleted       bool `json:"is_role_completed"`
}

type ParticipantToolResponse struct {
	Tool  string `json:"tool"`
	Level string `json:"level"`
}

type ParticipantTraingResponse struct {
	Roles []RoleTrainingResponse `json:"roles"`
}

type RoleTrainingResponse struct {
	Name      string             `json:"name"`
	Trainings []TrainingResponse `json:"trainings"`
}

type TrainingResponse struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Selected          bool   `json:"selected,omitempty"`
	NeedSertification bool   `json:"need_certification,omitempty"`
}
