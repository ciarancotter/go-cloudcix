package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CloudCIXClient struct {
	Email, Password, ApiKey, Token, ApiUrl string
}

func (cixClient CloudCIXClient) GetToken() (string, error) {
	json_data := map[string]string{"api_key": cixClient.ApiKey, "email": cixClient.Email, "password": cixClient.Password}
	json_value, err := json.Marshal(json_data)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://membership.%s/auth/login/", cixClient.ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(json_value))
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	var resp map[string]string
	err = json.Unmarshal(data, &resp)
	return resp["token"], err
}

func (cixClient CloudCIXClient) ReadData(application string, service string, object_id string, token string) (string, error) {
	client := http.Client{}
	url := "https://" + application + "." + cixClient.ApiUrl + "/" + service + "/" + object_id
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("X-Auth-Token", token)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	return ApplicationSwitch(application, service, object_id, data)
}

func (cixClient CloudCIXClient) WriteData(application string, service string, object_id string, token string, data map[string]string, method string) (string, error) {
	client := http.Client{}
	url := "https://" + application + "." + cixClient.ApiUrl + "/" + service + "/" + object_id
	post_data, _ := json.Marshal(data)
	post_data_buffer := bytes.NewBuffer(post_data)
	request, err := http.NewRequest(method, url, post_data_buffer)
	if err != nil {
		return "", err
	}

	request.Header.Set("X-Auth-Token", token)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (cix_client CloudCIXClient) DeleteData(application string, service string, object_id string, token string) (string, error) {
	client := http.Client{}
	url := "https://" + application + "." + cix_client.ApiUrl + "/" + service + "/" + object_id
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("X-Auth-Token", token)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	return string(body), nil
}
