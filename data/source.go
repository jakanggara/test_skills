package data

import (
	"math/rand"
	"skill_test/models"

	"github.com/google/uuid"
)

func (d *db_instance) seed_data_source(n string) (models.DbSource, error) {
	sp := uint(rand.Intn(maxSellPrice))
	s := models.DbSource{
		ID:           uuid.New(),
		ProductName:  n,
		Qty:          uint(rand.Intn(maxQty)),
		SellingPrice: sp,
		PromoPrice:   sp / 2,
	}

	err := d.db.Create(&s).Error
	if err != nil {
		return s, err
	}

	return s, nil
}
