package repo

import (
	"fmt"

	"github.com/ydhnwb/opaku-dummy-backend/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	All() ([]entity.Product, error)
	FindOneProductByID(ID string) (entity.Product, error)
	FindProductsByName(name string) ([]entity.Product, error)
}

type productRepo struct {
	connection *gorm.DB
}

func NewProductRepo(connection *gorm.DB) ProductRepository {
	return &productRepo{
		connection: connection,
	}
}

func (c *productRepo) All() ([]entity.Product, error) {
	products := []entity.Product{}
	// c.connection.Find(&products)
	c.connection.Find(&products)
	return products, nil
}

func (c *productRepo) FindOneProductByID(productID string) (entity.Product, error) {
	var product entity.Product
	res := c.connection.Where("id = ?", productID).Take(&product)
	if res.Error != nil {
		return product, res.Error
	}
	return product, nil
}

func (c *productRepo) FindProductsByName(name string) ([]entity.Product, error) {
	q := fmt.Sprintf("%%%s%%", name)
	products := []entity.Product{}
	res := c.connection.Where("name ILIKE ?", q).Find(&products)
	if res.Error != nil {
		return products, res.Error
	}
	return products, nil
}
