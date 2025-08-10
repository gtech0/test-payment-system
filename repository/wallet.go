package repository

import (
	"test-payment-system/database"
	"test-payment-system/errs"
	"test-payment-system/model"

	"github.com/google/uuid"
)

// WalletRepository handles database operations related to model.Wallet.
type WalletRepository struct{}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{}
}

func (w *WalletRepository) Create(wallet *model.Wallet) error {
	if err := database.DB.Create(wallet).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (w *WalletRepository) Save(wallet *model.Wallet) error {
	if err := database.DB.Save(wallet).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (w *WalletRepository) FindById(id uuid.UUID) (*model.Wallet, error) {
	wallet := new(model.Wallet)
	if err := database.DB.Model(&model.Wallet{}).
		Where("id = ?", id).
		First(wallet).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return wallet, nil
}
