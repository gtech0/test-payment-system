package mapper

import (
	"test-payment-system/dto"
	"test-payment-system/model"
)

type WalletMapper struct{}

func NewWalletMapper() *WalletMapper {
	return &WalletMapper{}
}

func (w *WalletMapper) ToBalanceDto(wallet *model.Wallet) dto.WalletBalanceDto {
	var mappedBalance dto.WalletBalanceDto
	mappedBalance.Balance = wallet.Balance
	return mappedBalance
}
