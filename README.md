# CloudCIX Go SDK

An implementation of the CloudCIX API in Go. When I refer to "PK endpoints", I
am referring to the endpoints that take an ID value after the service - for 
example, "/user/1", where 1 is the PK value. The use of the term "PK" is just 
to mimic the CloudCIX Python library. The Go SDK is logically divided into 
the reading/writing functionality, the structs, and the switches.

![Flowchart](https://raw.githubusercontent.com/ciarancotter/cixlogo/main/flowchart_go_sdk.png "Go SDK Flowchart")

## Authentication
To begin with, configure the client by setting the necessary properties of the 
object (see the `main.go` example).
The `Get_Token` function generates a token. This token is then used to 
authenticate the user when performing various requests to the API.

## Read Data

The `Read_Data` function takes four parameters - the application, 
the service, the PK value and the authentication token. 
- The application parameter specifies the application within the CloudCIX SAAS, 
  such as IAAS or Membership. 
- The service specifies what service you want to read from - for 
  example, you could call 
  `my_client.Read_Data("membership", "user", "some user ID", my_client.Token)` 
  to fetch the details of your chosen user ID.
- The PK value specifies the value after the service in the URL. 
  This value will depend on the resource you are trying to access.

Due to the fact that Go is a relatively static language, it was not feasible to 
dynamically generate structs for the JSON data. Instead, the structs have been
defined already in their respective files (see `iaas.go` for IAAS structs,
`membership.go` for Membership structs, and so on). 

This makes it very easy to access the data in a highly organised, structured 
way. The `Read_Data` function utilises switch statements to make this possible.

## Write Data

The `Write_Data` function takes a parameter called `method` to decide whether to
send a PATCH, POST or PUT request. The usage of the `Write_Data` function is not
complicated.
- The `method` parameter must be spelled in capital letters. It should be either
  `PATCH`, `POST` or `PUT`.
- The `data` parameter must be of type `map[string]string`. This sounds
  complicated, but it just means that you must declare your JSON payload as
  follows: 
  ```
  my_data := map[string]string{"test": "test"}
  my_client.Write_Data("iaas", "project", "", my_client.Token, my_data, "POST")
  ```

## Delete Data 

The delete functionality is very simple. It is almost identical in structure
to a GET request. The `Delete_Data` function handles DELETE requests.

## Usage / Underlying Code

For example, if you were to get the ClientIP value of the first IPMI instance 
(index 0) via the general (non-PK) endpoint, you would access it as follows:

```
ipmi_instance := IPMI{}
err = json.Unmarshal(data, &ipmi_instance)
fmt.Println(ipmi_instance.Content[0].ClientIP)
```

On the other hand, if you were to search by ID (which I refer to as PK 
endpoints), you would use the following snippet:

```
ipmi_instance := IPMI_SPECIFIC{}
err = json.Unmarshal(data, &ipmi_instance)
fmt.Println(ipmi_instance.Content)
```

However, this is all handled by the functions and switches, and all you have to
do is run the `Read_Data` function in `main.go`. At the moment, all of the 
switches just print the data. Feel free to modify these switches to your own 
needs.

All code accessing the API client itself should be in `main.go`. 
The sample I provided contains the following:

```
package main

import (
	"CloudCIX/api"
)

func main() {
	my_client := api.CloudCIX_Client{
		Email:    "example email address",
		Password: "example password",
		Api_key:  "example API key",
	}

	my_client.Token = my_client.Get_Token()

	my_client.Read_Data("membership", "user", "insert user ID here", my_client.Token)

	my_data := map[string]string{"test": "test"}
	my_client.Write_Data("iaas", "project", "", my_client.Token, my_data, "POST")

	my_client.Delete_Data("membership", "user", "insert user ID here", my_client.Token)
}
```

## Miscellaneous / Other
- `switches.go` handles all of the switch statements needed to select the 
   correct application and service. The main switch selects between 
   applications. The secondary switch selects the service within the application.
- To maintain this from a developer's perspective, just add and remove fields 
  from the structs as necessary. 
- It is recommended that you do not hardcode your credentials, and instead read 
  them from a configuration file of some sort.

## Developers
- To construct custom structs from JSON data, visit 
  [this website](https://mholt.github.io/json-to-go/). You can use this to 
  update the SDK as needed by just copying the JSON from the 
  [CloudCIX REST API](https://docs.cloudcix.com/) and pasting it in.
- Have fun! ~ Ciarán

