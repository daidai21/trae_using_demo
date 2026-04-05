package repository

import (
	"ecommerce/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) FindAll() ([]*model.Product, error) {
	var products []*model.Product
	err := r.db.Preload("Merchant").Find(&products).Error
	return products, err
}

func (r *ProductRepository) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.Preload("Merchant").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Update(product *model.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(product *model.Product) error {
	return r.db.Delete(product).Error
}

func (r *ProductRepository) FindByMerchantID(merchantID uint) ([]*model.Product, error) {
	var products []*model.Product
	err := r.db.Where("merchant_id = ?", merchantID).Find(&products).Error
	return products, err
}

func (r *ProductRepository) DecreaseStockWithTransaction(tx *gorm.DB, productID uint, quantity int) error {
	return tx.Model(&model.Product{}).Where("id = ? AND stock >= ?", productID, quantity).UpdateColumn("stock", gorm.Expr("stock - ?", quantity)).Error
}
