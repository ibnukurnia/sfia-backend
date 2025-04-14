package responses

import (
	"fmt"
	"sv-sfia/models"
)

type AssessmentStatusResponse struct {
}

type SelfAssessmentResponse struct {
	Options []option `json:"options"`
	Skills  []skill  `json:"skills"`
}

type option struct {
	Option string `json:"option"`
	Value  int8   `json:"value"`
}

type skill struct {
	Name     string          `json:"name"`
	Roles    []string        `json:"roles"`
	Question []skillQuestion `json:"questions"`
}

type skillQuestion struct {
	Id            string                       `json:"id"`
	Question      string                       `json:"question"`
	CurrentAnswer *selfAssessmentCurrentAnswer `json:"current_answer,omitempty"`
}

type selfAssessmentCurrentAnswer struct {
	Id       string `json:"id"`
	Value    int8   `json:"value"`
	Evidence string `json:"evidence"`
}

func NewSfiaResponse(skills []models.Skill, answers []models.SfiaAnswer) SelfAssessmentResponse {
	r := []skill{}
	a := []option{}

	for _, participantSkill := range skills {
		roles := []string{}

		for _, roleSkill := range participantSkill.RoleSkills {
			roles = append(roles, roleSkill.Role.Name)
		}

		questions := []skillQuestion{}
		for _, question := range participantSkill.SfiaQuestions {
			questions = append(questions, skillQuestion{
				Id:       question.Uuid.String(),
				Question: question.Question,
			})
		}

		r = append(r, skill{
			Name:     fmt.Sprintf("%s (%s)", participantSkill.Name, participantSkill.Code),
			Roles:    roles,
			Question: questions,
		})
	}

	for _, v := range answers {
		a = append(a, option{
			Option: v.Answer,
			Value:  v.Value,
		})
	}

	return SelfAssessmentResponse{
		Skills:  r,
		Options: a,
	}
}

type ToolAssessmentResponse struct {
	Tools []toolAssessment `json:"tools"`
}

type toolAssessment struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
}

func NewToolAssessmentResponse(tools []models.ParticipantTool) ToolAssessmentResponse {
	toolResponse := []toolAssessment{}

	for _, tool := range tools {
		toolResponse = append(toolResponse, toolAssessment{
			Id:    tool.Uuid.String(),
			Name:  tool.Name,
			Level: tool.Level,
		})
	}

	return ToolAssessmentResponse{
		Tools: toolResponse,
	}
}

type DujAssessmentResponse struct {
	Jobs []job `json:"jobs"`
}

func NewDujAssessmentResponse(m []models.Duj) DujAssessmentResponse {
	jobs := []job{}

	for _, v := range m {
		jobs = append(jobs, job{
			Id:     v.Uuid.String(),
			Detail: v.Detail,
		})
	}

	return DujAssessmentResponse{
		Jobs: jobs,
	}
}

func NewDujAssessmentResponseCurrentAnswer(m []models.DujAnswer) DujAssessmentResponse {
	jobs := []job{}

	for _, v := range m {
		jobs = append(jobs, job{
			Id:           v.Uuid.String(),
			CurrentJob:   &v.CurrentJob,
			HaveTrouble:  &v.HaveTrouble,
			TroubleCause: v.TroubleCause,
		})
	}

	return DujAssessmentResponse{
		Jobs: jobs,
	}
}

type job struct {
	Id           string  `json:"id,omitempty"`
	Detail       string  `json:"detail"`
	CurrentJob   *bool   `json:"current_job,omitempty"`
	HaveTrouble  *bool   `json:"have_trouble,omitempty"`
	TroubleCause *string `json:"trouble_cause,omitempty"`
}

type AssessmentResponse struct {
	Id     string                  `json:"id"`
	Status models.AssessmentStatus `json:"status"`
}

type ListAssessmentResponse struct {
	Id     string                  `json:"id"`
	Year   uint16                  `json:"year"`
	Status models.AssessmentStatus `json:"status"`
}

type AssessmentResumeResponse struct {
	GeneralInformation GeneralInformationResume `json:"general_information"`
	RoleInformation    RoleInformationResume    `json:"role_information"`
	SkillResults       []SkillResult            `json:"skill_results"`
}

type GeneralInformationResume struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Pn             string `json:"pn"`
	Department     string `json:"department"`
	DepartmentUnit string `json:"department_unit"`
	DepartmentTeam string `json:"department_team"`
}

type RoleInformationResume struct {
	MainRole      RoleResume  `json:"main_role"`
	SecondaryRole *RoleResume `json:"secondary_role"`
	InterestRole  *RoleResume `json:"interest_role"`
}

type RoleResume struct {
	Name  string           `json:"name"`
	Group string           `json:"group"`
	Level models.RoleLevel `json:"level"`
}

type SkillResult struct {
	Code  string            `json:"code"`
	Score float32           `json:"score"`
	Level models.SkillLevel `json:"level"`
	Name  string            `json:"name"`
}

type SfiaRoleResult struct {
	Name   string           `json:"name"`
	Group  string           `json:"group"`
	Skills []SkillResult    `json:"skills"`
	Level  models.RoleLevel `json:"level"`
}

type SfiaRoleResultResponse struct {
	MainRole      SfiaRoleResult  `json:"main_role"`
	SecondaryRole *SfiaRoleResult `json:"secondary_role"`
	InterestRole  *SfiaRoleResult `json:"interest_role"`
}
