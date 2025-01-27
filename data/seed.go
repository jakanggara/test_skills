package data

import (
	"log"

	"gorm.io/gorm"
)

const (
	maxSellPrice int = 100000
	maxQty           = 100
)

type db_instance struct {
	db *gorm.DB
}

func Init(db *gorm.DB) db_instance {
	return db_instance{
		db: db,
	}
}

func Seed(db *gorm.DB, count uint) {
	dbi := Init(db)
	log.Print("seeding...")
	dbi.seed_data(count)
}
