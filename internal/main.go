package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/subosito/gotenv"
)

var Scheduler = gocron.NewScheduler(time.UTC)

func main() {
	// load .env file if any otherwise use env set
	_ = gotenv.Load()

	_, err := Scheduler.Cron("* 8-18/3 * * 1-5").Tag("main").Do(func() {
		fmt.Println("Start checking appointment")

		task()

		fmt.Println("End checking appointment")
	})

	if err != nil {
		fmt.Println("Couldn't initiate the main job - " + err.Error())
	}

	// Starting all the job
	Scheduler.StartBlocking()
}
