package models

import "github.com/google/uuid"

type Training struct {
	Base
	Name   string
	RoleId uuid.UUID
	Role   Role
}

type TrainingMaster struct {
	Base
	Name     string
	Code     string
	Jenjang  string
	SkillsId  uuid.UUID
	Level    string
	Type     string
	Mode     string
	Provider string
	Silabus  string
}