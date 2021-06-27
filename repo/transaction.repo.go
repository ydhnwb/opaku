package repo

import (
	"github.com/ydhnwb/opaku-dummy-backend/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAllMyTransaction(userID string) ([]entity.Transaction, error)
	CreateTransaction(transaction entity.Transaction) (entity.Transaction, error)
}

type transactionRepo struct {
	connection *gorm.DB
}

func NewTransactionRepo(connection *gorm.DB) TransactionRepository {
	return &transactionRepo{
		connection: connection,
	}
}

func (c *transactionRepo) FindAllMyTransaction(userID string) ([]entity.Transaction, error) {
	transactions := []entity.Transaction{}
	res := c.connection.Preload("User").Where("user_id = ?", userID).Find(&transactions)
	if res.Error != nil {
		return transactions, res.Error
	}
	return transactions, nil
}

func (c *transactionRepo) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	res := c.connection.Save(&transaction)
	if res.Error != nil {
		return transaction, res.Error
	}
	c.connection.Preload("User").Take(&transaction)
	return transaction, nil
}
