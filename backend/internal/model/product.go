package model

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	MerchantID  uint           `gorm:"not null;index" json:"merchant_id"`
	Name        string         `gorm:"not null;size:100" json:"name"`
	Description string         `gorm:"size:500" json:"description"`
	Price       float64        `gorm:"not null;type:decimal(10,2)" json:"price"`
	Stock       int            `gorm:"not null;default:0" json:"stock"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Merchant    Merchant       `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
}
