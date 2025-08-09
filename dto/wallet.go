package dto

import "github.com/shopspring/decimal"

type WalletBalanceDto struct {
	Balance decimal.Decimal `json:"balance"`
}
