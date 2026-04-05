package model

import (
	"time"
	"gorm.io/gorm"
)

type AuctionStatus string

const (
	AuctionStatusPending AuctionStatus = "pending"
	AuctionStatusLive    AuctionStatus = "live"
	AuctionStatusEnded   AuctionStatus = "ended"
	AuctionStatusSold    AuctionStatus = "sold"
	AuctionStatusUnsold  AuctionStatus = "unsold"
)

type Currency string

const (
	CurrencyCNY Currency = "CNY"
	CurrencyUSD Currency = "USD"
	CurrencyIDR Currency = "IDR"
)

type Auction struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	ProductID      uint           `gorm:"not null;index" json:"product_id"`
	SellerID       uint           `gorm:"not null;index" json:"seller_id"`
	Title          string         `gorm:"not null;size:100" json:"title"`
	Description    string         `gorm:"size:500" json:"description"`
	Currency       Currency       `gorm:"not null;default:CNY;size:3" json:"currency"`
	StartPrice     float64        `gorm:"not null;type:decimal(15,2)" json:"start_price"`
	ReservePrice   float64        `gorm:"type:decimal(15,2)" json:"reserve_price,omitempty"`
	BidIncrement   float64        `gorm:"not null;type:decimal(15,2);default:10" json:"bid_increment"`
	CurrentPrice   float64        `gorm:"not null;type:decimal(15,2)" json:"current_price"`
	HighestBidderID *uint          `gorm:"index" json:"highest_bidder_id,omitempty"`
	Status         AuctionStatus  `gorm:"not null;default:pending;size:20" json:"status"`
	StartTime      *time.Time     `json:"start_time,omitempty"`
	EndTime        *time.Time     `json:"end_time,omitempty"`
	Duration       int            `gorm:"not null;default:3600" json:"duration"`
	BidCount       int            `gorm:"default:0" json:"bid_count"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	Bids           []Bid          `gorm:"foreignKey:AuctionID" json:"bids,omitempty"`
}

type Bid struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AuctionID uint      `gorm:"not null;index" json:"auction_id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Amount    float64   `gorm:"not null;type:decimal(15,2)" json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
