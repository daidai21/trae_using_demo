package cart

import (
	"errors"

	"ecommerce/trade-service/internal/model"
	"ecommerce/trade-service/internal/repository"

	"gorm.io/gorm"
)

type CartService struct {
	cartRepo *repository.CartRepository
}

func NewCartService(cartRepo *repository.CartRepository) *CartService {
	return &CartService{cartRepo: cartRepo}
}

type AddToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

func (s *CartService) AddToCart(userID uint, productID uint, quantity int, productStock int) (*model.Cart, error) {
	if productStock < quantity {
		return nil, errors.New("insufficient stock")
	}

	existingCart, err := s.cartRepo.FindByUserIDAndProductID(userID, productID)
	if err == nil {
		existingCart.Quantity += quantity
		if err := s.cartRepo.Update(existingCart); err != nil {
			return nil, err
		}
		return s.cartRepo.FindByID(existingCart.ID)
	}

	cart := &model.Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	if err := s.cartRepo.Create(cart); err != nil {
		return nil, err
	}

	return s.cartRepo.FindByID(cart.ID)
}

func (s *CartService) GetCart(userID uint) ([]*model.Cart, error) {
	return s.cartRepo.FindByUserID(userID)
}

func (s *CartService) UpdateCartItem(userID uint, cartID uint, quantity int, productStock int) (*model.Cart, error) {
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

	if productStock < quantity {
		return nil, errors.New("insufficient stock")
	}

	cart.Quantity = quantity
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
