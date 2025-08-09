package service

import (
	"github.com/google/uuid"
	"test-payment-system/dto"
	"test-payment-system/mapper"
	"test-payment-system/repository"
)

type WalletService struct {
	walletRepository      *repository.WalletRepository
	walletMapper          *mapper.WalletMapper
	transactionRepository *repository.TransactionRepository
}

func NewWalletService() *WalletService {
	return &WalletService{
		walletRepository:      repository.NewWalletRepository(),
		walletMapper:          mapper.NewWalletMapper(),
		transactionRepository: repository.NewTransactionRepository(),
	}
}

func (w *WalletService) GetBalance(id uuid.UUID) (dto.WalletBalanceDto, error) {
	wallet, err := w.walletRepository.FindById(id)
	if err != nil {
		return dto.WalletBalanceDto{}, err
	}

	return w.walletMapper.ToBalanceDto(wallet), nil
}
