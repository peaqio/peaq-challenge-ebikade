package schedule

import (
	"fmt"
	"sync"
	"time"
)

// Repository ...
type Repository interface {
	Run(func(), time.Duration)
}

// Schedule ...
type Schedule struct {
}

// NewSchedule ...
func NewSchedule() *Schedule {
	return &Schedule{}
}

// Run ...
func (s *Schedule) Run(fn func(), timeInMinute time.Duration) {
	fmt.Println("Schedule called")
	ticker := time.NewTicker(timeInMinute * time.Minute)
	quit := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func(fnc func()) {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				fnc()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}(fn)
	// close(quit)
	wg.Wait()
}
