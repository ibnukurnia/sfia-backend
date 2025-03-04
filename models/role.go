package models

import "github.com/google/uuid"

type Role struct {
	Base
	Name       string
	GroupId    uuid.UUID
	Group      RoleGroup `gorm:"reference:GroupId"`
	RoleSkills []RoleSkill
	Trainings  []Training
}

type Skill struct {
	Base
	Name          string
	Code          string
	RoleSkills    []RoleSkill
	SfiaQuestions []SfiaQuestion `gorm:"foreignKey:SkillId"`
}

type RoleSkill struct {
	Base
	RoleId  uuid.UUID
	Role    Role
	SkillId uuid.UUID
	Skill   Skill
}
