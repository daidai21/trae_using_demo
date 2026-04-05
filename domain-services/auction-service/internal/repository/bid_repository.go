package repository

import (
	"ecommerce/auction-service/internal/model"
	"gorm.io/gorm"
)

type BidRepository interface {
	Create(bid *model.Bid) error
	GetByAuctionID(auctionID uint) ([]*model.Bid, error)
	GetLatestByAuctionID(auctionID uint) (*model.Bid, error)
}

type bidRepository struct {
	db *gorm.DB
}

func NewBidRepository(db *gorm.DB) BidRepository {
	return &bidRepository{db: db}
}

func (r *bidRepository) Create(bid *model.Bid) error {
	return r.db.Create(bid).Error
}

func (r *bidRepository) GetByAuctionID(auctionID uint) ([]*model.Bid, error) {
	var bids []*model.Bid
	err := r.db.Where("auction_id = ?", auctionID).Order("created_at DESC").Find(&bids).Error
	return bids, err
}

func (r *bidRepository) GetLatestByAuctionID(auctionID uint) (*model.Bid, error) {
	var bid model.Bid
	err := r.db.Where("auction_id = ?", auctionID).Order("created_at DESC").First(&bid).Error
	if err != nil {
		return nil, err
	}
	return &bid, nil
}
