package jobs

import (
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/pkg/growth"
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/pkg/rate"
	"github.com/ebikode/peaq-challenge/challenge-3/exchange/pkg/schedule"
	storage "github.com/ebikode/peaq-challenge/challenge-3/exchange/storage/mysql"
)

const (
	scheduleInMinites = 5
)

// Init Initialize all scheduled jobs
func Init(mdb *storage.EDatabase) {

	growthStorage := storage.NewGrowthRecordStorage(mdb)
	rateStorage := storage.NewRateStorage(mdb)

	// Init Services
	growthService := growth.NewService(growthStorage)
	rateService := rate.NewService(rateStorage)

	// Init market job handler
	jobHandler := NewMarketJobHandler(rateService, growthService)

	// Initialize schedules
	sched := schedule.NewSchedule()

	var runJobs = func() {

		var runMarkDataAutomation = func() {
			jobHandler.getMarketData()
		}
		go runMarkDataAutomation()
		go sched.Run(runMarkDataAutomation, scheduleInMinites)
	}
	runJobs()
}
