package api

import (
	"log"
	"skill_test/models"
	"skill_test/utils"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllSource() ([]models.DbSource, error) {
	var result []models.DbSource

	err := r.db.Find(&result).Error
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *Repository) ChainUpdate() error {

	updates, err := r.GetAllSource()
	if err != nil {
		log.Printf("Transaction failed due to fetch error: %v", r)
		return err
	}

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Transaction rollback due to panic: %v", r)
		}
	}()

	var wg sync.WaitGroup
	errChan := make(chan error, len(updates))

	for _, source := range updates {
		wg.Add(1)
		go func(s models.DbSource) {
			defer wg.Done()
			for attempt := 1; attempt <= 3; attempt++ {
				err := tx.Model(&models.DbDestination{}).Where("id = ?", s.ID).Updates(models.DbDestination{
					Qty:          s.Qty,
					SellingPrice: s.SellingPrice,
					PromoPrice:   s.PromoPrice,
				}).Error
				if err != nil {
					if utils.IsDeadlockError(err) {
						log.Printf("Database locked, retrying attempt %d in %v...", attempt, time.Second)
						time.Sleep(time.Second)
						continue
					}
					errChan <- err
					break
				}
			}
		}(source)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		tx.Rollback()
		return <-errChan
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		return err
	}

	return nil

}
