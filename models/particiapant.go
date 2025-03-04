package models

const DefaultPN = "99999"

type Participant struct {
	Base
	Name     string
	Password string
	Pn       string
	Email    string

	ParticipantRole       *ParticipantRole
	ParticipantDepartment *ParticipantDepartment
	ParticipantSkills     []ParticipantSkill
	SelfAssessmentAnswers []SelfAssessmentAnswer `gorm:"foreignKey:ParticipantId"`
}
