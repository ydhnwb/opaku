package service

import (
	"errors"

	"github.com/ydhnwb/opaku-dummy-backend/repo"
	_product "github.com/ydhnwb/opaku-dummy-backend/service/_product"
	"gorm.io/gorm"
)

type ProductService interface {
	All() (*[]_product.ProductResponse, error)
	FindOneProductByID(productID string) (*_product.ProductResponse, error)
	FindProductsByName(name string) (*[]_product.ProductResponse, error)
}

type productService struct {
	productRepo repo.ProductRepository
}

func NewProductService(productRepo repo.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (c *productService) All() (*[]_product.ProductResponse, error) {
	products, err := c.productRepo.All()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	prods := _product.NewProductArrayResponse(products)
	return &prods, nil
}

func (c *productService) FindOneProductByID(productID string) (*_product.ProductResponse, error) {
	product, err := c.productRepo.FindOneProductByID(productID)

	if err != nil {
		return nil, err
	}

	res := _product.NewProductResponse(product)
	return &res, nil
}

func (c *productService) FindProductsByName(name string) (*[]_product.ProductResponse, error) {
	products, err := c.productRepo.FindProductsByName(name)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	res := _product.NewProductArrayResponse(products)
	return &res, nil
}
