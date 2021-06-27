package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ydhnwb/opaku-dummy-backend/dto"
	"github.com/ydhnwb/opaku-dummy-backend/entity"
	"github.com/ydhnwb/opaku-dummy-backend/repo"
	_transaction "github.com/ydhnwb/opaku-dummy-backend/service/_transaction"
	"gorm.io/gorm"
)

type TransactionService interface {
	FindAllMyTransaction(userID string) (*[]_transaction.TransactionResponse, error)
	CreateTransaction(createTransactionRequest dto.CreateTransactionRequest, userID string) (*_transaction.TransactionResponse, error)
}

type transactionService struct {
	transactionRepo repo.TransactionRepository
	productRepo     repo.ProductRepository
}

func NewTransactionService(transactionRepo repo.TransactionRepository, productRepo repo.ProductRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		productRepo:     productRepo,
	}
}

func (c *transactionService) FindAllMyTransaction(userID string) (*[]_transaction.TransactionResponse, error) {
	transactions, err := c.transactionRepo.FindAllMyTransaction(userID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	result := _transaction.NewTransactionArrayResponse(transactions)
	return &result, nil
}

func (c *transactionService) CreateTransaction(createTransactionRequest dto.CreateTransactionRequest, userID string) (*_transaction.TransactionResponse, error) {
	product, err := c.productRepo.FindOneProductByID(fmt.Sprintf("%d", createTransactionRequest.ProductID))
	if err != nil {
		return nil, err
	}

	id, _ := strconv.ParseUint(userID, 10, 64)
	transactionEntity := entity.Transaction{
		Name:   product.Name,
		Price:  product.Price,
		Image:  product.Image,
		UserID: id,
	}

	res, err := c.transactionRepo.CreateTransaction(transactionEntity)
	if err != nil {
		return nil, err
	}
	result := _transaction.NewTransactionResponse(res)
	return &result, nil
}
