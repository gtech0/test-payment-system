package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base is a gorm.Model with uuid.UUID as an id type.
// It's used as a base for other models.
type Base struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
