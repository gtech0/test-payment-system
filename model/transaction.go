package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	Base
	Amount     decimal.Decimal
	SenderId   uuid.UUID `gorm:"type:uuid"`
	ReceiverId uuid.UUID `gorm:"type:uuid"`
}
