package models

import (
	"github.com/google/uuid"
)

type (
	SkillLevel string
	RoleLevel  string
)

func (e RoleLevel) NextLevel() *RoleLevel {
	if e == SENIOR {
		return nil
	}

	nextLevel := MIDDLE

	if e == MIDDLE {
		nextLevel = SENIOR
	}

	return &nextLevel
}

func (e SkillLevel) Int() int {
	if e == BASIC {
		return 1
	}

	if e == INTERMEDIATE {
		return 2
	}

	return 3

}

const (
	BASIC        SkillLevel = "basic"
	INTERMEDIATE SkillLevel = "intermediate"
	ADVANCE      SkillLevel = "advance"

	NOLEVEL RoleLevel = "not qualified"
	JUNIOR  RoleLevel = "junior"
	MIDDLE  RoleLevel = "middle"
	SENIOR  RoleLevel = "senior"
)

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
	Description   string
	RoleSkills    []RoleSkill
	SfiaQuestions []SfiaQuestion `gorm:"foreignKey:SkillId"`
}

type RoleSkill struct {
	Base
	RoleId                                                           uuid.UUID
	Role                                                             Role
	SkillId                                                          uuid.UUID
	IsMandatoryForJunior, IsMandatoryForMiddle, IsMandatoryForSenior bool
	RequirementForJunior, RequirementForMiddle, RequirementForSenior SkillLevel
	Skill                                                            Skill
}
