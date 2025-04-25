package responses

import "sv-sfia/models"

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	Name  string            `json:"name"`
	Role  models.RoleAccess `json:"role"`
	Pn    string            `json:"pn"`
	Email string            `json:"email"`
}

type ParticipantAuthResponse struct {
	Token             string      `json:"token"`
	Participant       Participant `json:"participant"`
	IsProfileComplete bool        `json:"is_profile_complete"`
}

type Participant struct {
	Name string `json:"name"`
	Pn   string `json:"pn"`
}

type ParticipantRole struct {
	Id    string `json:"id"`
	Group string `json:"group"`
	Role  string `json:"role"`
}

type RoleInformation struct {
	Main      *ParticipantRole `json:"main"`
	Secondary *ParticipantRole `json:"secondary"`
	Interest  *ParticipantRole `json:"interest"`
}

func NewParticipantRoleInformation(m models.ParticipantRole) *RoleInformation {
	roleInformation := RoleInformation{}

	mainRole := ParticipantRole{
		Id:    m.MainRoleId.String(),
		Role:  m.MainRole.Name,
		Group: m.MainRole.Group.Name,
	}

	roleInformation.Main = &mainRole

	if m.SecondaryRole != nil {
		roleInformation.Secondary = &ParticipantRole{
			Id:    m.SecondaryRoleId.String(),
			Role:  m.SecondaryRole.Name,
			Group: m.SecondaryRole.Group.Name,
		}
	}

	if m.InterestRole != nil {
		roleInformation.Interest = &ParticipantRole{
			Id:    m.InterestRoleId.String(),
			Role:  m.InterestRole.Name,
			Group: m.InterestRole.Group.Name,
		}
	}

	return &roleInformation
}

type ParticipantProfileResponse struct {
	Name               string           `json:"name"`
	Pn                 string           `json:"pn"`
	Email              string           `json:"email"`
	DepartmentId       *string          `json:"department_id"`
	DepartmentRoleId   *string          `json:"department_role_id"`
	DepartmentTeamId   *string          `json:"department_team_id"`
	RoleInformation    *RoleInformation `json:"role_information"`
	IsProfileCompleted bool             `json:"is_profile_completed"`
}
