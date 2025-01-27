package api

import "log"

type Service struct {
	r     *Repository
	Tasks chan bool
}

func InitService(repo *Repository, queueSize int) *Service {
	return &Service{
		r:     repo,
		Tasks: make(chan bool, queueSize),
	}
}

func (s *Service) Run() {
	go func() {
		for range s.Tasks {
			log.Println("Processing update task...")
			err := s.r.ChainUpdate()
			if err != nil {
				log.Printf("Update failed: %v", err)
			} else {
				log.Println("Update completed successfully!")
			}
		}
	}()
}
func (w *Service) Queue() {
	w.Tasks <- true

}
