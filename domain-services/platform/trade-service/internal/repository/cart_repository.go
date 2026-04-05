package repository

import (
	"ecommerce/trade-service/internal/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) Create(cart *model.Cart) error {
	return r.db.Create(cart).Error
}

func (r *CartRepository) FindByUserID(userID uint) ([]*model.Cart, error) {
	var carts []*model.Cart
	err := r.db.Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

func (r *CartRepository) FindByID(id uint) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.First(&cart, id).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) FindByUserIDAndProductID(userID uint, productID uint) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepository) Update(cart *model.Cart) error {
	return r.db.Save(cart).Error
}

func (r *CartRepository) Delete(cart *model.Cart) error {
	return r.db.Delete(cart).Error
}

func (r *CartRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}
