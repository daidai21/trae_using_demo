package model

import (
	"time"
	"gorm.io/gorm"
)

type Merchant struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"uniqueIndex;not null" json:"user_id"`
	Name        string         `gorm:"not null;size:100" json:"name"`
	Description string         `gorm:"size:500" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
