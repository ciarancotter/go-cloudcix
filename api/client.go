package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CloudCIX_Client struct {
	Email, Password, Api_key, Token, API_URL string
}

func (cix_client CloudCIX_Client) Get_Token() string {

	json_data := map[string]string{"api_key": cix_client.Api_key, "email": cix_client.Email, "password": cix_client.Password}
	json_value, _ := json.Marshal(json_data)
	url := "https://membership.api.cloudcix.com/auth/login/"
	response, err := http.Post(url, "application/json", bytes.NewBuffer(json_value))

	if err != nil {
		return "Failed"
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var HTTPresponse map[string]string
		json.Unmarshal(data, &HTTPresponse)
		return HTTPresponse["token"]
	}
}

func (cix_client CloudCIX_Client) Read_Data(application string, service string, object_id string, token string) {
	client := http.Client{}
	url := "https://" + application + "." + cix_client.API_URL + "/" + service + "/" + object_id
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("X-Auth-Token", token)
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Failed")
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("This data is not a JSON response.")
		} else {
			Application_switch(application, service, object_id, data, err)
		}

	}
}

func (cix_client CloudCIX_Client) Write_Data(application string, service string, object_id string, token string, data map[string]string, method string) {
	client := http.Client{}
	url := "https://" + application + ".api.cloudcix.com/" + service + "/" + object_id
	post_data, _ := json.Marshal(data)
	post_data_buffer := bytes.NewBuffer(post_data)
	request, _ := http.NewRequest(method, url, post_data_buffer)
	request.Header.Set("X-Auth-Token", token)
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Failed")
	} else {
		response_data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(response_data))
	}
}

func (cix_client CloudCIX_Client) Delete_Data(application string, service string, object_id string, token string) {
	client := http.Client{}
	url := "https://" + application + ".api.cloudcix.com/" + service + "/" + object_id
	request, _ := http.NewRequest("DELETE", url, nil)
	request.Header.Set("X-Auth-Token", token)
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Failed")
	} else {
		if err != nil {
			fmt.Println("Did not delete.")
		} else {
			fmt.Println(response)
			fmt.Println("Deleted successfully")
		}

	}
}
