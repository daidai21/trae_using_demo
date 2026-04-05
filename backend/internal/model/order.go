package model

import (
	"time"
	"gorm.io/gorm"
)

type Order struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	TotalAmount float64        `gorm:"not null;type:decimal(10,2)" json:"total_amount"`
	Status      string         `gorm:"not null;size:20;default:'pending'" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	User        User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	OrderItems  []OrderItem    `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}
