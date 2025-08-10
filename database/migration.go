package database

import (
	"errors"
	"log"
	"math/rand"
	"test-payment-system/model"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func Sync() {
	if err := DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		log.Fatalf("failed to enable uuid-ossp extension: %v", err)
	}

	if err := DB.AutoMigrate(
		&model.Transaction{},
		&model.Wallet{},
	); err == nil && DB.Migrator().HasTable(&model.Wallet{}) {
		if err = DB.First(&model.Wallet{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			for i := 0; i < 10; i++ {
				var wallet model.Wallet
				wallet.Balance = decimal.NewFromInt(int64(rand.Uint32()))
				if err = DB.Create(&wallet).Error; err != nil {
					log.Fatal(err)
				}
			}
		}
	} else if err != nil {
		log.Fatal(err)
	}
}
