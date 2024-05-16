package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/schollz/progressbar/v3"
)

var BASE_URL = "https://app.acuityscheduling.com/api"

// 31666981 -> Découpe au laser
// 31667152 -> Brodeuse numérique
// 31778525 -> Decoupe de vinyle

// 4600904 -> Rendez-vous de 3 heures
// 4600906 -> Rendez-vous de 4 heures (jeudi. vendredi ou samedi)
// 4600908 -> Rendez-vous de 5 heures (jeudi. vendredi ou samedi)
// 4600914 -> Rendez-vous de 6 heures (jeudi. vendredi ou samedi)
// 4600920 -> Rendez-vous de 7 heures (jeudi. vendredi ou samedi)
// 4601413 -> Rendez-vous de 8 heures (samedi)

var (
	client      = NewClient(nil, BASE_URL)
	currentTime = time.Now()

	appointmentType = map[string]string{
		"31667152": "Brodeuse numérique",
		// "31778525": "Decoupe de vinyle",
		"31666981": "Découpe au laser",
	}

	addons = map[string]string{
		"4600904": "Rendez-vous de 3 heures",
		"4600906": "Rendez-vous de 4 heures (jeudi. vendredi ou samedi)",
		"4600908": "Rendez-vous de 5 heures (jeudi. vendredi ou samedi)",
		"4600914": "Rendez-vous de 6 heures (jeudi. vendredi ou samedi)",
		"4600920": "Rendez-vous de 7 heures (jeudi. vendredi ou samedi)",
		"4601413": "Rendez-vous de 8 heures (samedi)",
	}
)

func task() {
	// For all AppointmentTypeId
	for appointmentTypeId := range appointmentType {
		fmt.Println("Recherche for: " + appointmentType[appointmentTypeId])

		// For all addon
		for addonId := range addons {
			checkAvailability(appointmentTypeId, addonId)
		}

		// Small break
		makeABreak()
	}
}

func checkAvailability(appointmentTypeId string, addonId string) {
	res := client.Availability.ListMonth(
		currentTime.Format("2006-01-02"),
		appointmentTypeId,
		addonId,
	)

	addGotSlots := false
	for date, isAvailable := range res {
		// We have apointment available for this day
		if isAvailable {
			addGotSlots = true
			times := client.Availability.ListTime(
				date,
				appointmentTypeId,
				addonId,
			)

			fmt.Println(date + " slots disponible ( " + addons[addonId] + " ):")
			for _, time := range times {
				fmt.Println(time.Time)
			}
		}
	}

	if !addGotSlots {
		fmt.Println(addons[addonId] + " - ")
	}
}

// Stop the script for a short time
func makeABreak() {
	// Minimum time for the break
	min := 5
	// Maximum time for the break
	max := 15
	timeForBreak := rand.Intn(max-min) + min

	fmt.Printf("Going to make a short break: %d \n", timeForBreak)

	bar := progressbar.Default(int64(timeForBreak))
	for i := 0; i <= timeForBreak; i++ {
		_ = bar.Add(1)
		time.Sleep(1 * time.Second)
	}
}
