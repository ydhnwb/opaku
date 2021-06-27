package _cart

import (
	"github.com/ydhnwb/opaku-dummy-backend/entity"
	"github.com/ydhnwb/opaku-dummy-backend/service/_product"
	"github.com/ydhnwb/opaku-dummy-backend/service/_user"
)

type CartResponse struct {
	ID      uint64                   `json:"id"`
	Product _product.ProductResponse `json:"product,omitempty"`
	User    _user.UserResponse       `json:"user,omitempty"`
}

func NewCartResponse(cart entity.Cart) CartResponse {
	return CartResponse{
		ID:      cart.ID,
		Product: _product.NewProductResponse(cart.Product),
		User:    _user.NewUserResponse(cart.User),
	}
}

func NewCartArrayResponse(carts []entity.Cart) []CartResponse {
	cartRes := []CartResponse{}
	for _, v := range carts {

		p := CartResponse{
			ID:      v.ID,
			Product: _product.NewProductResponse(v.Product),
			User:    _user.NewUserResponse(v.User),
		}
		cartRes = append(cartRes, p)
	}
	return cartRes
}
