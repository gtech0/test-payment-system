package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionGetDto struct {
	Id         uuid.UUID       `json:"id"`
	Time       time.Time       `json:"time"`
	Amount     decimal.Decimal `json:"amount"`
	SenderId   uuid.UUID       `json:"senderId"`
	ReceiverId uuid.UUID       `json:"receiverId"`
}
