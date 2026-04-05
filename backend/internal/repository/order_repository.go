package repository

import (
	"ecommerce/internal/model"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) FindByID(id uint) (*model.Order, error) {
	var order model.Order
	err := r.db.Preload("User").Preload("OrderItems").Preload("OrderItems.Product").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindByUserID(userID uint) ([]*model.Order, error) {
	var orders []*model.Order
	err := r.db.Preload("OrderItems").Preload("OrderItems.Product").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) Update(order *model.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) CreateWithTransaction(tx *gorm.DB, order *model.Order) error {
	return tx.Create(order).Error
}

func (r *OrderRepository) CreateOrderItemWithTransaction(tx *gorm.DB, item *model.OrderItem) error {
	return tx.Create(item).Error
}
