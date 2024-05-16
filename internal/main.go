package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

var Scheduler = gocron.NewScheduler(time.UTC)

func main() {
	_, err := Scheduler.Every(1).Hour().Tag("main").Do(func() {
		fmt.Println("Start checking appointment")

		task()

		fmt.Println("End checking appointment")
	})

	if err != nil {
		fmt.Println("Couldn't initiate the main job - " + err.Error())
	}

	// Run job now
	Scheduler.RunAll()

	// Starting all the job
	Scheduler.StartBlocking()
}
