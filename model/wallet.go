package model

import "github.com/shopspring/decimal"

// Wallet is a model, which is used to store currency.
// Balance uses decimal.Decimal as a type for its precision.
type Wallet struct {
	Base
	Balance decimal.Decimal
}
