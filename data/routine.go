package data

import (
	"fmt"
	"log"
	"skill_test/models"
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
			name := fmt.Sprintf("Product-%v", time.Now().Format("20060102-150405-000000"))
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
	res, err := d.seed_data_source(s)
	if err != nil {
		log.Printf("error seed source: %v", err)
	}
	c <- res
}
