package dto

type NewParticipantDto struct {
	Name            string
	Pn              string
	Email           string
	MainRoleId      string
	SecondaryRoleId *string
	InterestRoleId  *string
	DepartmentId    string
}
