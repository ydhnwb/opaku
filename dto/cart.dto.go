package dto

type AddToCartRequest struct {
	ProductID uint64 `json:"product_id" form:"product_id" binding:"required"`
	// UserID    uint64 `json:"user_id" form:"user_id" binding:"required"`
}
