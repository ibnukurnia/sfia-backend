package models

type SkillLevelThreshold struct {
	Base
	Basic, Intermediate, Advance float32
}

func (m SkillLevelThreshold) ToLevel(score float32) SkillLevel {
	if m.Basic > score {
		return BASIC
	}

	if score >= m.Intermediate {
		return ADVANCE
	}

	return INTERMEDIATE
}
