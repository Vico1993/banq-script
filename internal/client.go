package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type service struct {
	client *Client
}

type Client struct {
	BaseURL string
	client  *http.Client // HTTP client used to communicate with the API.

	common service

	Availability *AvailabilityService
}

// Create a new client to communicate with Otto
func NewClient(httpClient *http.Client, baseUrl string) *Client {
	c := &Client{client: httpClient, BaseURL: baseUrl}

	// Initialise the client
	c.init()

	return c
}

func (c *Client) init() {
	if c.client == nil {
		c.client = &http.Client{}
	}

	c.common.client = c

	c.Availability = (*AvailabilityService)(&c.common)
}

// Execute the request that the client received
func (c *Client) Do(req *http.Request) ([]byte, error) {
	// Fetch Request
	response, err := c.client.Do(req)
	if err != nil {
		fmt.Println("Error making the request: " + err.Error())
		return []byte{}, err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		fmt.Println("Api respond with status code: " + strconv.Itoa(response.StatusCode))
		fmt.Println(string(body))
		return []byte{}, err
	}

	return body, nil
}
