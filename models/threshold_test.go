package models

import (
	"testing"
)

func TestToLevel(t *testing.T) {
	skillThresholds := SkillLevelThreshold{
		Basic:        3.0,
		Intermediate: 5.0,
		Advance:      6.0,
	}

	cases := []struct {
		Score float32
		Level SkillLevel
	}{
		{
			Score: 1.5,
			Level: BASIC,
		},
		{
			Score: 6,
			Level: ADVANCE,
		},
		{
			Score: 3.71,
			Level: INTERMEDIATE,
		},
	}

	for _, c := range cases {
		if c.Level != skillThresholds.ToLevel(c.Score) {
			t.Fail()
		}
	}
}
