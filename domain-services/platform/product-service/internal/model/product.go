package model

import (
	"time"
	"gorm.io/gorm"
)

type ProductType string

const (
	ProductTypeNormal   ProductType = "normal"
	ProductTypePreSale  ProductType = "pre_sale"
	ProductTypeAuction  ProductType = "auction"
)

type Currency string

const (
	CurrencyCNY Currency = "CNY"
	CurrencyUSD Currency = "USD"
	CurrencyIDR Currency = "IDR"
)

type ProductPrice struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProductID uint      `gorm:"not null;index" json:"product_id"`
	Currency  Currency  `gorm:"not null;size:3" json:"currency"`
	Price     float64   `gorm:"not null;type:decimal(15,2)" json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	MerchantID  uint           `gorm:"not null;index" json:"merchant_id"`
	Name        string         `gorm:"not null;size:100" json:"name"`
	Description string         `gorm:"size:500" json:"description"`
	Type        ProductType    `gorm:"not null;default:normal;size:20" json:"type"`
	Currency    Currency       `gorm:"not null;default:CNY;size:3" json:"currency"`
	Price       float64        `gorm:"not null;type:decimal(15,2)" json:"price"`
	Stock       int            `gorm:"not null;default:0" json:"stock"`
	Prices      []ProductPrice `gorm:"foreignKey:ProductID" json:"prices,omitempty"`
	
	PreSalePrice       float64    `gorm:"type:decimal(15,2)" json:"pre_sale_price,omitempty"`
	Deposit            float64    `gorm:"type:decimal(15,2)" json:"deposit,omitempty"`
	PreSaleStartAt     *time.Time `json:"pre_sale_start_at,omitempty"`
	PreSaleEndAt       *time.Time `json:"pre_sale_end_at,omitempty"`
	BalanceDueAt       *time.Time `json:"balance_due_at,omitempty"`
	
	StartPrice         float64    `gorm:"type:decimal(15,2)" json:"start_price,omitempty"`
	ReservePrice       float64    `gorm:"type:decimal(15,2)" json:"reserve_price,omitempty"`
	BidIncrement       float64    `gorm:"type:decimal(15,2)" json:"bid_increment,omitempty"`
	AuctionDuration    int        `gorm:"default:3600" json:"auction_duration,omitempty"`
	
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
