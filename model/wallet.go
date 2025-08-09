package model

import "github.com/shopspring/decimal"

type Wallet struct {
	Base
	Balance decimal.Decimal
	//Transactions   []*Transaction `gorm:"foreignKey:SenderId,ReceiverId"`
}
