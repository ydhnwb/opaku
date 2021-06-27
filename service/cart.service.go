package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/opaku-dummy-backend/dto"
	"github.com/ydhnwb/opaku-dummy-backend/entity"
	"github.com/ydhnwb/opaku-dummy-backend/repo"
	"github.com/ydhnwb/opaku-dummy-backend/service/_cart"
	"gorm.io/gorm"
)

type CartService interface {
	FindAllCart(userID string) (*[]_cart.CartResponse, error)
	AddToCart(addToCartRequest dto.AddToCartRequest, userID uint64) (*_cart.CartResponse, error)
	DeleteCart(cartID string) error
}

type cartService struct {
	cartRepo repo.CartRepository
}

func NewCartService(cartRepo repo.CartRepository) CartService {
	return &cartService{
		cartRepo: cartRepo,
	}
}

func (c *cartService) FindAllCart(userID string) (*[]_cart.CartResponse, error) {
	carts, err := c.cartRepo.FindAllCart(userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	res := _cart.NewCartArrayResponse(carts)
	return &res, nil
}

func (c *cartService) AddToCart(addToCartRequest dto.AddToCartRequest, userID uint64) (*_cart.CartResponse, error) {
	cartEntity := entity.Cart{}
	err := smapping.FillStruct(&cartEntity, smapping.MapFields(&addToCartRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	existingCart, err := c.cartRepo.FindOneByProductIDAndUserID(fmt.Sprintf("%v", addToCartRequest.ProductID), fmt.Sprintf("%v", userID))
	if err == nil {
		msg := fmt.Errorf("%s already exists!", existingCart.Product.Name)
		return nil, errors.New(msg.Error())
	}

	cartEntity.UserID = userID
	res, err := c.cartRepo.AddToCart(cartEntity)
	if err != nil {
		return nil, err
	}
	result := _cart.NewCartResponse(res)
	return &result, nil
}

func (c *cartService) DeleteCart(cartID string) error {
	err := c.cartRepo.DeleteCart(cartID)
	if err != nil {
		return err
	}
	return nil
}
