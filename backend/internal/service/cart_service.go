package service

import (
	"errors"

	"ecommerce/internal/model"
	"ecommerce/internal/repository"

	"gorm.io/gorm"
)

type CartService struct {
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}

func NewCartService(cartRepo *repository.CartRepository, productRepo *repository.ProductRepository) *CartService {
	return &CartService{cartRepo: cartRepo, productRepo: productRepo}
}

type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

func (s *CartService) AddToCart(userID uint, req *AddToCartRequest) (*model.Cart, error) {
	product, err := s.productRepo.FindByID(req.ProductID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	if product.Stock < req.Quantity {
		return nil, errors.New("insufficient stock")
	}

	existingCart, err := s.cartRepo.FindByUserIDAndProductID(userID, req.ProductID)
	if err == nil {
		existingCart.Quantity += req.Quantity
		if err := s.cartRepo.Update(existingCart); err != nil {
			return nil, err
		}
		return s.cartRepo.FindByID(existingCart.ID)
	}

	cart := &model.Cart{
		UserID:    userID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := s.cartRepo.Create(cart); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(cart.ID)
}

func (s *CartService) GetCart(userID uint) ([]*model.Cart, error) {
	return s.cartRepo.FindByUserID(userID)
}

func (s *CartService) UpdateCartItem(userID uint, cartID uint, req *UpdateCartRequest) (*model.Cart, error) {
	cart, err := s.cartRepo.FindByID(cartID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("cart item not found")
		}
		return nil, err
	}

	if cart.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	product, err := s.productRepo.FindByID(cart.ProductID)
	if err != nil {
		return nil, err
	}

	if product.Stock < req.Quantity {
		return nil, errors.New("insufficient stock")
	}

	cart.Quantity = req.Quantity
	if err := s.cartRepo.Update(cart); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(cart.ID)
}

func (s *CartService) DeleteCartItem(userID uint, cartID uint) error {
	cart, err := s.cartRepo.FindByID(cartID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("cart item not found")
		}
		return err
	}

	if cart.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.cartRepo.Delete(cart)
}
