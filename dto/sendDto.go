package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type SendDto struct {
	From   uuid.UUID       `json:"from"`
	To     uuid.UUID       `json:"to"`
	Amount decimal.Decimal `json:"amount"`
}
