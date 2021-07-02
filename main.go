package main

import (
	"fmt"

	"github.com/CloudCIX/go-cloudcix/api"
)

func main() {

	var err error
	// Credentials to use the API.
	client := api.CloudCIXClient{
		Email:    "jimbobjoe@example.com",
		Password: "verysecurepassword",
		ApiKey:   "1234567890abcdefg",
		ApiUrl:   "api.cloudcix.com",
	}

	client.Token, err = client.GetToken()
	if err != nil {
		// Something has gone wrong with retrieving token - print error
		panic(err)
	}

	// Returns a struct containing the information for the syllabus with ID 123.
	body, err := client.GetSyllabus("123")
	fmt.Println(body.Content.ID, err)
}
