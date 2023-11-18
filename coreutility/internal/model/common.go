package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Base Struct Contain Common column for Every Table.
type Base struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key" json:"id"`
	CreatedBy uuid.UUID  `gorm:"type:varchar(36)" json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedBy uuid.UUID  `gorm:"type:varchar(36)" json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedBy uuid.UUID  `gorm:"type:varchar(36)" json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
