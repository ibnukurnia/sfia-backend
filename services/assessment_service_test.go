package services

import (
	"sv-sfia/models"
	"testing"
)

var roleLevelChecker = newRoleLevelChecker()

func TestRoleDoesntPassJuniorLevel(t *testing.T) {
	skillLevels := map[string]models.SkillLevel{
		"PROG": models.ADVANCE,
		// "DTAN": models.INTERMEDIATE,
	}

	roleRequirements := []models.RoleSkill{
		{
			Skill: models.Skill{
				Code: "PROG",
			},
			IsMandatoryForJunior: false,
			IsMandatoryForMiddle: false,
			IsMandatoryForSenior: false,
			RequirementForJunior: models.INTERMEDIATE,
			RequirementForMiddle: models.INTERMEDIATE,
			RequirementForSenior: models.ADVANCE,
		},
		// {
		// 	Skill: models.Skill{
		// 		Code: "DTAN",
		// 	},
		// 	IsMandatoryForJunior: false,
		// 	IsMandatoryForMiddle: false,
		// 	RequirementForJunior: models.INTERMEDIATE,
		// 	RequirementForMiddle: models.ADVANCE,
		// },
		// {
		// 	Skill: models.Skill{
		// 		Code: "DBDS",
		// 	},
		// 	IsMandatoryForJunior: false,
		// 	IsMandatoryForMiddle: false,
		// 	RequirementForJunior: models.BASIC,
		// 	RequirementForMiddle: models.INTERMEDIATE,
		// },
	}

	level := checkRoleLevel(&roleLevelChecker, roleRequirements, skillLevels)

	t.Log(level)
}

// func TestPassedRoleLevelJunior(t *testing.T) {
// 	roleRequirements := map[string]models.RoleSkill{
// 		"PROG": {
// 			IsMandatoryForJunior: true,
// 			IsMandatoryForMiddle: true,
// 			RequirementForJunior: models.INTERMEDIATE,
// 			RequirementForMiddle: models.INTERMEDIATE,
// 		},
// 		"DTAN": {
// 			IsMandatoryForJunior: true,
// 			IsMandatoryForMiddle: true,
// 			RequirementForJunior: models.BASIC,
// 			RequirementForMiddle: models.INTERMEDIATE,
// 		},
// 		"DBDS": {
// 			IsMandatoryForJunior: false,
// 			IsMandatoryForMiddle: true,
// 			RequirementForJunior: models.BASIC,
// 			RequirementForMiddle: models.BASIC,
// 		},
// 	}

// 	skills := []responses.SkillResult{
// 		{
// 			Name:  "Programming",
// 			Code:  "PROG",
// 			Level: models.INTERMEDIATE,
// 			Score: 3.7,
// 		},
// 		{
// 			Name:  "Data Modeling And Design",
// 			Code:  "DTAN",
// 			Level: models.BASIC,
// 			Score: 1,
// 		},
// 		{
// 			Name:  "Database Design",
// 			Code:  "DBDS",
// 			Level: models.BASIC,
// 			Score: 1,
// 		},
// 	}

// 	// should be junior
// 	level := service.getRoleLevel(skills, roleRequirements)

// 	t.Log(level)

// 	if level != models.JUNIOR {
// 		t.Fail()
// 	}
// }

// func TestDoesntPassedRoleLevelMiddle(t *testing.T) {
// 	skillLevels := map[string]models.SkillLevel{
// 		"PROG": models.INTERMEDIATE,
// 		"DTAN": models.BASIC,
// 		"DBDS": models.BASIC,
// 	}

// 	roleRequirements := []models.RoleSkill{
// 		{
// 			Skill: models.Skill{
// 				Code: "PROG",
// 			},
// 			IsMandatoryForJunior: true,
// 			IsMandatoryForMiddle: true,
// 			RequirementForJunior: models.INTERMEDIATE,
// 			RequirementForMiddle: models.INTERMEDIATE,
// 		},
// 		{
// 			Skill: models.Skill{
// 				Code: "DTAN",
// 			},
// 			IsMandatoryForJunior: true,
// 			IsMandatoryForMiddle: true,
// 			RequirementForJunior: models.BASIC,
// 			RequirementForMiddle: models.INTERMEDIATE,
// 		},
// 		{
// 			Skill: models.Skill{
// 				Code: "DBDS",
// 			},
// 			IsMandatoryForJunior: false,
// 			IsMandatoryForMiddle: true,
// 			RequirementForJunior: models.BASIC,
// 			RequirementForMiddle: models.BASIC,
// 		},
// 	}

// }
