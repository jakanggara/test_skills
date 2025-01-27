package data

import (
	"fmt"
	"log"
	"math/rand"
	"skill_test/models"
	"skill_test/utils"
	"sync"
	"time"
)

func (d *db_instance) seed_data(count uint) {
	limit := 16
	queue := make(chan models.DbSource, limit)
	var wg sync.WaitGroup
	for i := 0; i < int(count); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			name := fmt.Sprintf("Product-%v", time.Now().Format("20060102-150405"))
			name = fmt.Sprintf("%v-%v", name, rand.Intn(10000))
			d.addData(name, queue)
		}()
	}

	go func() {
		wg.Wait()
		close(queue) // Closing the channel after all writing is complete
	}()

	for data := range queue {
		if err := d.seed_data_destination(&data); err != nil {
			log.Printf("error seeding destination: %v", err)
		}
	}
}

func (d *db_instance) addData(s string, c chan models.DbSource) {
	var res models.DbSource
	var err error
	for attempt := 1; attempt <= 3; attempt++ {
		res, err = d.seed_data_source(s)
		if err != nil {
			if utils.IsDeadlockError(err) {
				log.Printf("Database locked, retrying attempt %d in %v...", attempt, time.Second)
				time.Sleep(time.Second)
				continue
			}
			log.Printf("Error seeding source: %v", err)
			break
		}

		break
	}
	c <- res
}
