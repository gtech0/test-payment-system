package mapper

import (
	"test-payment-system/dto"
	"test-payment-system/model"
)

// TransactionMapper is used to convert model.Transaction to respective dto's.
type TransactionMapper struct{}

func NewTransactionMapper() *TransactionMapper {
	return &TransactionMapper{}
}

func (t *TransactionMapper) ToGetDto(transaction model.Transaction) dto.TransactionGetDto {
	var mappedTransaction dto.TransactionGetDto
	mappedTransaction.Id = transaction.Id
	mappedTransaction.Time = transaction.CreatedAt
	mappedTransaction.Amount = transaction.Amount
	mappedTransaction.SenderId = transaction.SenderId
	mappedTransaction.ReceiverId = transaction.ReceiverId
	return mappedTransaction
}

func (t *TransactionMapper) ToGetDtos(transactions []model.Transaction) []dto.TransactionGetDto {
	mappedTransactions := make([]dto.TransactionGetDto, 0)
	for _, transaction := range transactions {
		mappedTransactions = append(mappedTransactions, t.ToGetDto(transaction))
	}
	return mappedTransactions
}
