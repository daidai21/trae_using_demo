package repository

import (
	"ecommerce/internal/model"

	"gorm.io/gorm"
)

type MerchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

func (r *MerchantRepository) Create(merchant *model.Merchant) error {
	return r.db.Create(merchant).Error
}

func (r *MerchantRepository) FindAll() ([]*model.Merchant, error) {
	var merchants []*model.Merchant
	err := r.db.Preload("User").Find(&merchants).Error
	return merchants, err
}

func (r *MerchantRepository) FindByID(id uint) (*model.Merchant, error) {
	var merchant model.Merchant
	err := r.db.Preload("User").First(&merchant, id).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) FindByUserID(userID uint) (*model.Merchant, error) {
	var merchant model.Merchant
	err := r.db.Where("user_id = ?", userID).First(&merchant).Error
	if err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (r *MerchantRepository) Update(merchant *model.Merchant) error {
	return r.db.Save(merchant).Error
}
