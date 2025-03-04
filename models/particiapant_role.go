package models

import "github.com/google/uuid"

type ParticipantRole struct {
	Base
	MainRoleId          uuid.UUID   `gorm:"not null"`
	MainRole            Role        `gorm:"foreignKey:MainRoleId;references:uuid"`
	MainRoleSkills      []RoleSkill `gorm:"foreignKey:RoleId;references:MainRoleId"`
	SecondaryRoleId     *uuid.UUID
	SecondaryRole       *Role       `gorm:"foreignKey:SecondaryRoleId;references:uuid"`
	SecondaryRoleSkills []RoleSkill `gorm:"foreignKey:RoleId;references:SecondaryRoleId"`
	InterestRoleId      *uuid.UUID
	InterestRole        *Role       `gorm:"foreignKey:InterestRoleId;references:uuid"`
	InterestRoleSkills  []RoleSkill `gorm:"foreignKey:RoleId;references:InterestRoleId"`
	ParticipantId       uuid.UUID
}
