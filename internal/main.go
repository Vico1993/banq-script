package main

import (
	"fmt"
	"time"
	_ "time/tzdata"

	"github.com/go-co-op/gocron"
	"github.com/subosito/gotenv"
)

var Scheduler = gocron.NewScheduler(time.UTC)

func main() {
	// load .env file if any otherwise use env set
	_ = gotenv.Load()

	loc, _ := time.LoadLocation(timezone)

	_, err := Scheduler.Every(3).Hours().Tag("main").Do(func() {
		now := time.Now().In(loc)
		// Cron("1 8-18/3 * * 1-5")
		// TODO: Quick fix, but make sure to update the Scheduler
		if now.Hour() > 8 && now.Hour() < 18 {
			fmt.Println("Start checking appointment")
			task()
			fmt.Println("End checking appointment")
		}
	})

	if err != nil {
		fmt.Println("Couldn't initiate the main job - " + err.Error())
	}

	// Starting all the job
	Scheduler.StartBlocking()
}
