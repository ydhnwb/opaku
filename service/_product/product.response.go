package _product

import (
	"github.com/ydhnwb/opaku-dummy-backend/entity"
)

type ProductResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Price uint64 `json:"price"`
	Image string `json:"image"`
	// User  _user.UserResponse `json:"user,omitempty"`
}

func NewProductResponse(product entity.Product) ProductResponse {
	return ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Image: product.Image,
		// User:  _user.NewUserResponse(product.User),
	}
}

func NewProductArrayResponse(products []entity.Product) []ProductResponse {
	productRes := []ProductResponse{}
	for _, v := range products {

		p := ProductResponse{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Image: v.Image,
			// User:  _user.NewUserResponse(v.User),
		}
		productRes = append(productRes, p)
	}
	return productRes
}
