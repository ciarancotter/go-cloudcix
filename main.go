package main

import "github.com/ciarancotter/go-cloudcix/api"

func main() {
	my_client := api.CloudCIX_Client{
		Email:    "example_email",
		Password: "example_password",
		Api_key:  "example_api_key",
		API_URL:  "api.cloudcix.com",
	}

	my_client.Token = my_client.Get_Token()

	my_client.Read_Data("membership", "user", "insert user ID here", my_client.Token)

	// my_data := map[string]string{"test": "test"}
	// my_client.Write_Data("iaas", "project", "", my_client.Token, my_data, "POST")

	// my_client.Delete_Data("membership", "user", "insert user ID here", my_client.Token)
}
