package models

import "github.com/google/uuid"

type Training struct {
	Base
	Name   string
	RoleId uuid.UUID
	Role   Role
}
