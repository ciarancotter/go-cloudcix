package main

import (
	"fmt"

	"github.com/CloudCIX/go-cloudcix/api"
)

func main() {

	var err error
	// Credentials to use the API.
	client := api.CloudCIXClient{
		Email:    "example@example.com",
		Password: "3x4mpl3",
		ApiKey:   "example_api_key",
		ApiUrl:   "api.cloudcix.com",
	}

	client.Token, err = client.GetToken()
	if err != nil {
		// Something has gone wrong with retrieving token - print error
		panic(err)
	}

	// Returns a struct containing the information for the syllabus with ID 123.
	body, err := client.ListVm()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(body.Content[0], err)
}
