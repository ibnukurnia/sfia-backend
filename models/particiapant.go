package models

const DefaultPN = "99999"

type Participant struct {
	Base
	Name     string
	Password string
	Pn       string
	Email    string
	RoleAccess string `gorm:"type:role_access_enum;default:user;not null"`

	ParticipantRole       *ParticipantRole
	ParticipantDepartment *ParticipantDepartment
	ParticipantSkills     []ParticipantSkill
	SelfAssessmentAnswers []SelfAssessmentAnswer `gorm:"foreignKey:ParticipantId"`
}
