package dto

import "github.com/google/uuid"

type ParticipantRoleIdsDto struct {
	MainRoleId      uuid.UUID
	SecondaryRoleId *uuid.UUID
	InterestRoleId  *uuid.UUID
}
