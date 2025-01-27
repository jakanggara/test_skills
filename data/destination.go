package data

import "skill_test/models"

func (d *db_instance) seed_data_destination(s *models.DbSource) error {
	ss := models.DbDestination{
		ID:           s.ID,
		ProductName:  s.ProductName,
		Qty:          0,
		SellingPrice: 0,
		PromoPrice:   0,
	}

	err := d.db.Create(&ss).Error
	if err != nil {
		return err
	}

	return nil
}
