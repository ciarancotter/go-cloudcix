package main

import (
	"github.com/CloudCIX/go-cloudcix/api"
)

func main() {
	my_client := api.CloudCIX_Client{
		Email:    "example email",
		Password: "example password",
		Api_key:  "example API key",
		API_URL:  "api.cloudcix.com",
	}

	my_client.Token = my_client.Get_Token()

	my_client.Read_Data("membership", "user", "insert user ID here", my_client.Token)

	// my_data := map[string]string{"test": "test"}
	// my_client.Write_Data("iaas", "project", "", my_client.Token, my_data, "POST")

	// my_client.Delete_Data("membership", "user", "insert user ID here", my_client.Token)
}
