package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Transaction is a model, which is used to record transaction data.
// SenderId and ReceiverId are ids of wallets.
// Amount is a value transferred from sender to receiver.
type Transaction struct {
	Base
	Amount     decimal.Decimal
	SenderId   uuid.UUID `gorm:"type:uuid"`
	ReceiverId uuid.UUID `gorm:"type:uuid"`
}
