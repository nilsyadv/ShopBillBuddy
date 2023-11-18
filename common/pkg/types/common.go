package types

import (
	"time"
)

// Base Struct Contain Common column for Every Table.
type Base struct {
	ID        string     `gorm:"type:varchar(36);primary_key" json:"id"`
	CreatedBy string     `gorm:"type:varchar(36)" json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedBy string     `gorm:"type:varchar(36)" json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedBy string     `gorm:"type:varchar(36)" json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
