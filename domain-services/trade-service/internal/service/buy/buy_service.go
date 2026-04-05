package buy

import (
	"errors"

	"ecommerce/trade-service/internal/model"
	"ecommerce/trade-service/internal/repository"

	"gorm.io/gorm"
)

type BuyService struct {
	orderRepo *repository.OrderRepository
	cartRepo  *repository.CartRepository
	db        *gorm.DB
}

func NewBuyService(orderRepo *repository.OrderRepository, cartRepo *repository.CartRepository, db *gorm.DB) *BuyService {
	return &BuyService{orderRepo: orderRepo, cartRepo: cartRepo, db: db}
}

type ProductInfo struct {
	ID    uint
	Price float64
	Stock int
	Name  string
}

func (s *BuyService) CreateOrder(userID uint, carts []*model.Cart, products map[uint]*ProductInfo) (*model.Order, error) {
	if len(carts) == 0 {
		return nil, errors.New("cart is empty")
	}

	var totalAmount float64
	orderItems := make([]*model.OrderItem, 0, len(carts))

	for _, cart := range carts {
		product, exists := products[cart.ProductID]
		if !exists {
			return nil, errors.New("product not found")
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
	err := s.db.Transaction(func(tx *gorm.DB) error {
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
