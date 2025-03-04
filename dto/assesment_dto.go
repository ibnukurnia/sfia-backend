package dto

type FindAssessementByRolesDto struct {
	MainRoleId      uint
	SecondaryRoleId *uint
	InterestRoleId  *uint
}

type RoleAssesmentsDto struct {
	MainRoleAssesment      []assesmentDto
	SecondaryRoleAssesment []assesmentDto
	InterestRoleAssesment  []assesmentDto
}

type assesmentDto struct {
	Title     string
	Questions []assesmentQuestion
}

type assesmentQuestion struct {
	Id       uint
	Question string
	Answer   []assesmentAnswer
}

type assesmentAnswer struct {
	Id     uint
	Answer string
}
