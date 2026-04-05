package service

import (
	"errors"

	"ecommerce/internal/model"
	"ecommerce/internal/repository"

	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo   *repository.OrderRepository
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
	db          *gorm.DB
}

func NewOrderService(orderRepo *repository.OrderRepository, cartRepo *repository.CartRepository, productRepo *repository.ProductRepository, db *gorm.DB) *OrderService {
	return &OrderService{orderRepo: orderRepo, cartRepo: cartRepo, productRepo: productRepo, db: db}
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending paid shipped delivered cancelled"`
}

func (s *OrderService) CreateOrder(userID uint) (*model.Order, error) {
	carts, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	if len(carts) == 0 {
		return nil, errors.New("cart is empty")
	}

	var totalAmount float64
	orderItems := make([]*model.OrderItem, 0, len(carts))

	for _, cart := range carts {
		product, err := s.productRepo.FindByID(cart.ProductID)
		if err != nil {
			return nil, err
		}

		if product.Stock < cart.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		totalAmount += float64(cart.Quantity) * product.Price
		orderItems = append(orderItems, &model.OrderItem{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     product.Price,
		})
	}

	var order *model.Order
	err = s.db.Transaction(func(tx *gorm.DB) error {
		order = &model.Order{
			UserID:      userID,
			TotalAmount: totalAmount,
			Status:      "pending",
		}

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		for _, item := range orderItems {
			item.OrderID = order.ID
			if err := tx.Create(item).Error; err != nil {
				return err
			}
		}

		for _, cart := range carts {
			if err := tx.Model(&model.Product{}).Where("id = ? AND stock >= ?", cart.ProductID, cart.Quantity).UpdateColumn("stock", gorm.Expr("stock - ?", cart.Quantity)).Error; err != nil {
				return err
			}
		}

		if err := tx.Where("user_id = ?", userID).Delete(&model.Cart{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.orderRepo.FindByID(order.ID)
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

func (s *OrderService) UpdateOrderStatus(userID uint, orderID uint, req *UpdateOrderStatusRequest) (*model.Order, error) {
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

	order.Status = req.Status
	if err := s.orderRepo.Update(order); err != nil {
		return nil, err
	}

	return s.orderRepo.FindByID(orderID)
}
