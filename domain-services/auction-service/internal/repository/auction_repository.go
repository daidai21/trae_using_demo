package repository

import (
	"ecommerce/auction-service/internal/model"
	"gorm.io/gorm"
)

type AuctionRepository interface {
	Create(auction *model.Auction) error
	GetByID(id uint) (*model.Auction, error)
	GetAll() ([]*model.Auction, error)
	GetByStatus(status model.AuctionStatus) ([]*model.Auction, error)
	Update(auction *model.Auction) error
	Delete(id uint) error
}

type auctionRepository struct {
	db *gorm.DB
}

func NewAuctionRepository(db *gorm.DB) AuctionRepository {
	return &auctionRepository{db: db}
}

func (r *auctionRepository) Create(auction *model.Auction) error {
	return r.db.Create(auction).Error
}

func (r *auctionRepository) GetByID(id uint) (*model.Auction, error) {
	var auction model.Auction
	err := r.db.Preload("Bids").First(&auction, id).Error
	if err != nil {
		return nil, err
	}
	return &auction, nil
}

func (r *auctionRepository) GetAll() ([]*model.Auction, error) {
	var auctions []*model.Auction
	err := r.db.Find(&auctions).Error
	return auctions, err
}

func (r *auctionRepository) GetByStatus(status model.AuctionStatus) ([]*model.Auction, error) {
	var auctions []*model.Auction
	err := r.db.Where("status = ?", status).Find(&auctions).Error
	return auctions, err
}

func (r *auctionRepository) Update(auction *model.Auction) error {
	return r.db.Save(auction).Error
}

func (r *auctionRepository) Delete(id uint) error {
	return r.db.Delete(&model.Auction{}, id).Error
}
