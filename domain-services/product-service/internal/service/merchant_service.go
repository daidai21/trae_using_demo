package service

import (
	"errors"

	"ecommerce/product-service/internal/model"
	"ecommerce/product-service/internal/repository"

	"gorm.io/gorm"
)

type MerchantService struct {
	merchantRepo *repository.MerchantRepository
}

func NewMerchantService(merchantRepo *repository.MerchantRepository) *MerchantService {
	return &MerchantService{merchantRepo: merchantRepo}
}

type CreateMerchantRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Description string `json:"description" binding:"max=500"`
}

type UpdateMerchantRequest struct {
	Name        string `json:"name" binding:"required,max=100"`
	Description string `json:"description" binding:"max=500"`
}

func (s *MerchantService) CreateMerchant(userID uint, req *CreateMerchantRequest) (*model.Merchant, error) {
	existingMerchant, _ := s.merchantRepo.FindByUserID(userID)
	if existingMerchant != nil {
		return nil, errors.New("user already has a merchant")
	}

	merchant := &model.Merchant{
		UserID:      userID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.merchantRepo.Create(merchant); err != nil {
		return nil, err
	}

	return s.merchantRepo.FindByID(merchant.ID)
}

func (s *MerchantService) GetMerchants() ([]*model.Merchant, error) {
	return s.merchantRepo.FindAll()
}

func (s *MerchantService) GetMerchantByID(id uint) (*model.Merchant, error) {
	return s.merchantRepo.FindByID(id)
}

func (s *MerchantService) UpdateMerchant(userID uint, id uint, req *UpdateMerchantRequest) (*model.Merchant, error) {
	merchant, err := s.merchantRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("merchant not found")
		}
		return nil, err
	}

	if merchant.UserID != userID {
		return nil, errors.New("you are not authorized to update this merchant")
	}

	merchant.Name = req.Name
	merchant.Description = req.Description

	if err := s.merchantRepo.Update(merchant); err != nil {
		return nil, err
	}

	return s.merchantRepo.FindByID(merchant.ID)
}
