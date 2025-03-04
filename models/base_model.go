package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	Uuid      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;column:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Base) BeforeCreate(tx *gorm.DB) error {
	m.Uuid = uuid.New()

	return nil
}
