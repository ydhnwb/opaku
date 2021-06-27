package _transaction

import (
	"github.com/ydhnwb/opaku-dummy-backend/entity"
	_user "github.com/ydhnwb/opaku-dummy-backend/service/_user"
)

type TransactionResponse struct {
	ID    uint64             `json:"id"`
	Name  string             `json:"name"`
	Price uint64             `json:"price"`
	Image string             `json:"image"`
	User  _user.UserResponse `json:"user,omitempty"`
}

func NewTransactionResponse(transaction entity.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:    transaction.ID,
		Name:  transaction.Name,
		Price: transaction.Price,
		Image: transaction.Image,
		User:  _user.NewUserResponse(transaction.User),
	}
}

func NewTransactionArrayResponse(transactions []entity.Transaction) []TransactionResponse {
	transactionRes := []TransactionResponse{}
	for _, v := range transactions {

		p := TransactionResponse{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Image: v.Image,
			User:  _user.NewUserResponse(v.User),
		}
		transactionRes = append(transactionRes, p)
	}
	return transactionRes
}
