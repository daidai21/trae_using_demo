package model

import (
	"time"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderID   uint           `gorm:"not null;index" json:"order_id"`
	ProductID uint           `gorm:"not null;index" json:"product_id"`
	Quantity  int            `gorm:"not null;default:1" json:"quantity"`
	Price     float64        `gorm:"not null;type:decimal(10,2)" json:"price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Order     Order          `gorm:"foreignKey:OrderID" json:"order,omitempty"`
	Product   Product        `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}
