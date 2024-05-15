package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type AvailabilityTimeResponse struct {
	SlotsAvailable int    `json:"slotsAvailable"`
	Time           string `json:"time"`
}

type AvailabilityService service

var (
	owner      = "76426cf9"
	calendarId = "6653366"
	timezone   = "America/Toronto"
)

// Retrieve day available
// month format: YYYY-MM-DD [2024-05-01]
func (s *AvailabilityService) ListMonth(month string, appointmentTypeId string, addonId string) map[string]bool {
	queryParams := url.Values{
		"owner":             {owner},
		"appointmentTypeId": {appointmentTypeId},
		"calendarId":        {calendarId},
		"timezone":          {timezone},
		"month":             {month},
		"addonIds[]":        {addonId},
	}

	req, err := http.NewRequest(
		http.MethodGet,
		s.client.BaseURL+"/scheduling/v1/availability/month?"+queryParams.Encode(),
		strings.NewReader(
			string([]byte{}),
		),
	)

	if err != nil {
		fmt.Println("Error creating the request to query availability: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		return nil
	}

	var res map[string]bool
	_ = json.Unmarshal(body, &res)

	return res
}

// Retrieve time for each day
// date format: YYYY-MM-DD [2024-05-15]
func (s *AvailabilityService) ListTime(date string, appointmentTypeId string, addonId string) []AvailabilityTimeResponse {
	queryParams := url.Values{
		"owner":             {owner},
		"appointmentTypeId": {appointmentTypeId},
		"calendarId":        {calendarId},
		"startDate":         {date},
		"timezone":          {timezone},
		"addonIds[]":        {addonId},
	}

	req, err := http.NewRequest(
		http.MethodGet,
		s.client.BaseURL+"/scheduling/v1/availability/times?"+queryParams.Encode(),
		strings.NewReader(
			string([]byte{}),
		),
	)

	if err != nil {
		fmt.Println("Error creating the request to query availability: " + err.Error())
		return nil
	}

	body, err := s.client.Do(req)
	if err != nil {
		return nil
	}

	var res map[string][]AvailabilityTimeResponse
	_ = json.Unmarshal(body, &res)

	return res[date]
}
