package dto

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type TransactionGetDto struct {
	Id         uuid.UUID       `json:"id"`
	Time       time.Time       `json:"time"`
	Amount     decimal.Decimal `json:"amount"`
	SenderId   uuid.UUID       `json:"senderId"`
	ReceiverId uuid.UUID       `json:"receiverId"`
}
