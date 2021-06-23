package main

import (
	"fmt"

	"github.com/CloudCIX/go-cloudcix/api"
)

func main() {
	client := api.CloudCIXClient{
		Email:    "example email",
		Password: "example password",
		ApiKey:   "example API key",
		ApiUrl:   "api.cloudcix.com",
	}

	var err error
	client.Token, err = client.GetToken()
	if err != nil {
		// Something has gone wrong with retrieving token - print error
		panic(err)
	}

	data, err := client.ReadData("membership", "user", "insert user ID here", client.Token)
	if err != nil {
		// Couldn't read data - handle error to find out why
		panic(err)
	}

	fmt.Println(data)

	// my_data := map[string]string{"test": "test"}
	// my_client.Write_Data("iaas", "project", "", my_client.Token, my_data, "POST")

	// my_client.Delete_Data("membership", "user", "insert user ID here", my_client.Token)
}
