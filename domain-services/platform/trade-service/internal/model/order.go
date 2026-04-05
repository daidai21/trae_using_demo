package model

import (
	"time"
	"gorm.io/gorm"
)

type Currency string

const (
	CurrencyCNY Currency = "CNY"
	CurrencyUSD Currency = "USD"
	CurrencyIDR Currency = "IDR"
)

type OrderStatus string

const (
	OrderStatusPending         OrderStatus = "pending"
	OrderStatusPaid            OrderStatus = "paid"
	OrderStatusShipped         OrderStatus = "shipped"
	OrderStatusCompleted       OrderStatus = "completed"
	OrderStatusCancelled       OrderStatus = "cancelled"
	OrderStatusPendingDeposit  OrderStatus = "pending_deposit"
	OrderStatusDepositPaid     OrderStatus = "deposit_paid"
	OrderStatusPendingBalance  OrderStatus = "pending_balance"
)

type Order struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	UserID          uint           `gorm:"not null;index" json:"user_id"`
	Currency        Currency       `gorm:"not null;default:CNY;size:3" json:"currency"`
	TotalAmount     float64        `gorm:"not null;type:decimal(15,2)" json:"total_amount"`
	Status          OrderStatus    `gorm:"not null;size:20;default:'pending'" json:"status"`
	IsPreSale       bool           `gorm:"default:false" json:"is_pre_sale"`
	DepositAmount   float64        `gorm:"type:decimal(15,2)" json:"deposit_amount,omitempty"`
	BalanceAmount   float64        `gorm:"type:decimal(15,2)" json:"balance_amount,omitempty"`
	BalanceDueAt    *time.Time     `json:"balance_due_at,omitempty"`
	TaxAmount       float64        `gorm:"type:decimal(15,2);default:0" json:"tax_amount"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	OrderItems      []OrderItem    `gorm:"foreignKey:OrderID" json:"order_items,omitempty"`
}
