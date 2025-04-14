package models

import "github.com/google/uuid"

type ParticipantSkill struct {
	Base
	AssessmentId  uuid.UUID
	ParticipantId uuid.UUID
	SkillId       uuid.UUID
	IsMastered    bool
	UsedFor       int8
	Skill         Skill
	Participant   Participant
}
