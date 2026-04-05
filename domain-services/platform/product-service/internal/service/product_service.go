package service

import (
	"errors"

	"ecommerce/product-service/internal/model"
	"ecommerce/product-service/internal/repository"

	"gorm.io/gorm"
)

type ProductService struct {
	productRepo  *repository.ProductRepository
	merchantRepo *repository.MerchantRepository
}

func NewProductService(productRepo *repository.ProductRepository, merchantRepo *repository.MerchantRepository) *ProductService {
	return &ProductService{productRepo: productRepo, merchantRepo: merchantRepo}
}

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required,max=100"`
	Description string  `json:"description" binding:"max=500"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,min=0"`
}

type UpdateProductRequest struct {
	Name        string  `json:"name" binding:"required,max=100"`
	Description string  `json:"description" binding:"max=500"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"required,min=0"`
}

func (s *ProductService) CreateProduct(userID uint, req *CreateProductRequest) (*model.Product, error) {
	merchant, err := s.merchantRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("you are not a merchant")
		}
		return nil, err
	}

	product := &model.Product{
		MerchantID:  merchant.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	return s.productRepo.FindByID(product.ID)
}

func (s *ProductService) GetProducts() ([]*model.Product, error) {
	return s.productRepo.FindAll()
}

func (s *ProductService) GetProductByID(id uint) (*model.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *ProductService) UpdateProduct(userID uint, id uint, req *UpdateProductRequest) (*model.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	merchant, err := s.merchantRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("you are not a merchant")
		}
		return nil, err
	}

	if product.MerchantID != merchant.ID {
		return nil, errors.New("you are not authorized to update this product")
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	return s.productRepo.FindByID(product.ID)
}

func (s *ProductService) DeleteProduct(userID uint, id uint) error {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	merchant, err := s.merchantRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("you are not a merchant")
		}
		return err
	}

	if product.MerchantID != merchant.ID {
		return errors.New("you are not authorized to delete this product")
	}

	return s.productRepo.Delete(product)
}
