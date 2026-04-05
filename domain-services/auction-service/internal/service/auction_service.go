package service

import (
	"ecommerce/auction-service/internal/model"
	"ecommerce/auction-service/internal/repository"
	"ecommerce/auction-service/internal/ws"
	"errors"
	"fmt"
	"time"
)

type AuctionService interface {
	CreateAuction(auction *model.Auction, sellerID uint) error
	GetAuction(id uint) (*model.Auction, error)
	GetAllAuctions() ([]*model.Auction, error)
	GetLiveAuctions() ([]*model.Auction, error)
	StartAuction(id uint) error
	PlaceBid(auctionID uint, userID uint, amount float64) (*model.Bid, error)
	EndAuction(id uint) error
}

type auctionService struct {
	auctionRepo repository.AuctionRepository
	bidRepo     repository.BidRepository
	hub         *ws.Hub
}

func NewAuctionService(auctionRepo repository.AuctionRepository, bidRepo repository.BidRepository, hub *ws.Hub) AuctionService {
	return &auctionService{
		auctionRepo: auctionRepo,
		bidRepo:     bidRepo,
		hub:         hub,
	}
}

func (s *auctionService) CreateAuction(auction *model.Auction, sellerID uint) error {
	auction.SellerID = sellerID
	auction.Status = model.AuctionStatusPending
	auction.CurrentPrice = auction.StartPrice
	auction.BidCount = 0
	return s.auctionRepo.Create(auction)
}

func (s *auctionService) GetAuction(id uint) (*model.Auction, error) {
	return s.auctionRepo.GetByID(id)
}

func (s *auctionService) GetAllAuctions() ([]*model.Auction, error) {
	return s.auctionRepo.GetAll()
}

func (s *auctionService) GetLiveAuctions() ([]*model.Auction, error) {
	return s.auctionRepo.GetByStatus(model.AuctionStatusLive)
}

func (s *auctionService) StartAuction(id uint) error {
	auction, err := s.auctionRepo.GetByID(id)
	if err != nil {
		return err
	}
	if auction.Status != model.AuctionStatusPending {
		return errors.New("拍卖状态不正确")
	}

	now := time.Now()
	endTime := now.Add(time.Duration(auction.Duration) * time.Second)
	auction.Status = model.AuctionStatusLive
	auction.StartTime = &now
	auction.EndTime = &endTime

	if err := s.auctionRepo.Update(auction); err != nil {
		return err
	}

	s.broadcastUpdate(auction)
	return nil
}

func (s *auctionService) PlaceBid(auctionID uint, userID uint, amount float64) (*model.Bid, error) {
	auction, err := s.auctionRepo.GetByID(auctionID)
	if err != nil {
		return nil, err
	}

	if auction.Status != model.AuctionStatusLive {
		return nil, errors.New("拍卖未开始或已结束")
	}

	minBid := auction.CurrentPrice + auction.BidIncrement
	if amount < minBid {
		return nil, fmt.Errorf("出价必须至少为 %.2f", minBid)
	}

	if auction.HighestBidderID != nil && *auction.HighestBidderID == userID {
		return nil, errors.New("您已经是最高出价者")
	}

	bid := &model.Bid{
		AuctionID: auctionID,
		UserID:    userID,
		Amount:    amount,
	}
	if err := s.bidRepo.Create(bid); err != nil {
		return nil, err
	}

	auction.CurrentPrice = amount
	auction.HighestBidderID = &userID
	auction.BidCount++
	if err := s.auctionRepo.Update(auction); err != nil {
		return nil, err
	}

	s.broadcastBid(auction, bid)
	return bid, nil
}

func (s *auctionService) EndAuction(id uint) error {
	auction, err := s.auctionRepo.GetByID(id)
	if err != nil {
		return err
	}
	if auction.Status != model.AuctionStatusLive {
		return errors.New("拍卖未在进行中")
	}

	now := time.Now()
	auction.EndTime = &now
	if auction.HighestBidderID != nil && (auction.ReservePrice == 0 || auction.CurrentPrice >= auction.ReservePrice) {
		auction.Status = model.AuctionStatusSold
	} else {
		auction.Status = model.AuctionStatusUnsold
	}

	if err := s.auctionRepo.Update(auction); err != nil {
		return err
	}

	s.broadcastUpdate(auction)
	return nil
}

func (s *auctionService) broadcastUpdate(auction *model.Auction) {
	msg := &ws.Message{
		Type:      ws.MessageTypeUpdate,
		AuctionID: auction.ID,
		Data: map[string]interface{}{
			"status":           auction.Status,
			"current_price":    auction.CurrentPrice,
			"highest_bidder":   auction.HighestBidderID,
			"bid_count":        auction.BidCount,
			"start_time":       auction.StartTime,
			"end_time":         auction.EndTime,
			"online_count":     s.hub.GetOnlineCount(auction.ID),
		},
		Timestamp: time.Now().Unix(),
	}
	s.hub.Broadcast <- msg
}

func (s *auctionService) broadcastBid(auction *model.Auction, bid *model.Bid) {
	msg := &ws.Message{
		Type:      ws.MessageTypeBid,
		AuctionID: auction.ID,
		UserID:    bid.UserID,
		Data: map[string]interface{}{
			"amount":     bid.Amount,
			"user_id":    bid.UserID,
			"created_at": bid.CreatedAt,
		},
		Timestamp: time.Now().Unix(),
	}
	s.hub.Broadcast <- msg
}
