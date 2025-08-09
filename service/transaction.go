package service

import (
	"test-payment-system/dto"
	"test-payment-system/errs"
	"test-payment-system/mapper"
	"test-payment-system/model"
	"test-payment-system/repository"
)

type TransactionService struct {
	walletRepository      *repository.WalletRepository
	transactionRepository *repository.TransactionRepository
	transactionMapper     *mapper.TransactionMapper
}

func NewTransactionService() *TransactionService {
	return &TransactionService{
		walletRepository:      repository.NewWalletRepository(),
		transactionRepository: repository.NewTransactionRepository(),
		transactionMapper:     mapper.NewTransactionMapper(),
	}
}

func (t *TransactionService) Send(sendDto dto.SendDto) error {
	sender, err := t.walletRepository.FindById(sendDto.From)
	if err != nil {
		return err
	}

	receiver, err := t.walletRepository.FindById(sendDto.To)
	if err != nil {
		return err
	}

	if sender.Balance.LessThan(sendDto.Amount) {
		return errs.New("insufficient funds", 402)
	}

	sender.Balance = sender.Balance.Sub(sendDto.Amount)
	receiver.Balance = receiver.Balance.Add(sendDto.Amount)

	if err = t.walletRepository.Save(sender); err != nil {
		return err
	}

	if err = t.walletRepository.Save(receiver); err != nil {
		return err
	}

	if err = t.transactionRepository.Create(&model.Transaction{
		Amount:     sendDto.Amount,
		SenderId:   sender.Id,
		ReceiverId: receiver.Id,
	}); err != nil {
		return err
	}

	return nil
}

func (t *TransactionService) GetLast(count int) ([]dto.TransactionGetDto, error) {
	transactions, err := t.transactionRepository.FindLast(count)
	if err != nil {
		return nil, err
	}

	return t.transactionMapper.ToGetDtos(transactions), nil
}
