package repo

import (
	"github.com/ydhnwb/opaku-dummy-backend/entity"
	"gorm.io/gorm"
)

type CartRepository interface {
	FindAllCart(userID string) ([]entity.Cart, error)
	FindOneByProductIDAndUserID(productID string, userID string) (entity.Cart, error)
	AddToCart(cart entity.Cart) (entity.Cart, error)
	DeleteCart(cartID string) error
}

type cartRepo struct {
	connection *gorm.DB
}

func NewCartRepo(connection *gorm.DB) CartRepository {
	return &cartRepo{
		connection: connection,
	}
}

func (c *cartRepo) FindAllCart(userID string) ([]entity.Cart, error) {
	carts := []entity.Cart{}
	c.connection.Preload("Product").Preload("User").Where("user_id = ?", userID).Find(&carts)
	return carts, nil
}

func (c *cartRepo) FindOneByProductIDAndUserID(productID string, userID string) (entity.Cart, error) {
	var cart entity.Cart
	res := c.connection.Preload("User").Preload("Product").Where("product_id = ? AND user_id = ?", productID, userID).Take(&cart)
	if res.Error != nil {
		return cart, res.Error
	}
	return cart, nil
}

func (c *cartRepo) AddToCart(cart entity.Cart) (entity.Cart, error) {
	c.connection.Save(&cart)
	c.connection.Preload("Product").Preload("User").Find(&cart)
	return cart, nil
}

func (c *cartRepo) DeleteCart(cartID string) error {
	res := c.connection.Delete(&entity.Cart{}, cartID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
