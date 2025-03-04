package models

import "github.com/google/uuid"

type ParticipantSkill struct {
	Base
	ParticipantId uuid.UUID
	SkillId       uuid.UUID
	IsMastered    bool
	UsedFor       int8
	Skill         Skill
	Participant   Participant
}
