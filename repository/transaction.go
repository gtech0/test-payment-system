package repository

import (
	"test-payment-system/database"
	"test-payment-system/errs"
	"test-payment-system/model"
)

// TransactionRepository handles database operations related to model.Transaction.
type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (t *TransactionRepository) Create(transaction *model.Transaction) error {
	if err := database.DB.Create(transaction).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (t *TransactionRepository) FindLast(count int) ([]model.Transaction, error) {
	transactions := make([]model.Transaction, 0)
	if err := database.DB.Model(&model.Transaction{}).
		Order("created_at DESC").
		Limit(count).
		Find(&transactions).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return transactions, nil
}
