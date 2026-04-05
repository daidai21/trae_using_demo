package order

import (
	"errors"

	"ecommerce/trade-service/internal/model"
	"ecommerce/trade-service/internal/repository"

	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo *repository.OrderRepository
}

func NewOrderService(orderRepo *repository.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending paid shipped delivered cancelled"`
}

func (s *OrderService) GetOrders(userID uint) ([]*model.Order, error) {
	return s.orderRepo.FindByUserID(userID)
}

func (s *OrderService) GetOrderByID(userID uint, orderID uint) (*model.Order, error) {
	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	if order.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	return order, nil
}

func (s *OrderService) UpdateOrderStatus(userID uint, orderID uint, status string) (*model.Order, error) {
	order, err := s.orderRepo.FindByID(orderID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	if order.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	order.Status = status
	if err := s.orderRepo.Update(order); err != nil {
		return nil, err
	}

	return s.orderRepo.FindByID(orderID)
}
