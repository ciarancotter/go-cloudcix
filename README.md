# CloudCIX Go SDK

This is an implementation of the CloudCIX REST API in Go. When I refer to "PK endpoints", I am referring to the endpoints that take an ID value - for example, "/user/1", where 1 is the PK value. The use of the term *PK* is for the sake of consistency with the CloudCIX Python library.

The Go SDK is logically divided into the Read/Write functionality, the structs of each resource, and the client functions. 
Every *resource* - that being an object in CloudCIX - has a corresponding List, Get and Write function where applicable. 

- **List** takes no PK value. It returns a struct containing a list of all instances of the resource for which the user has permission to view.
- **Get** takes a PK value, and returns a struct containing the information about the specified instance.
- **Write** allows the user to send a POST, PATCH or PUT request. It takes a `method` parameter and a `data` parameter, which specify the method and the payload to be sent respectively.

## Authentication
To begin with, configure the client by setting the necessary properties of the 
object (see the `main.go` example).
The `GetToken` function generates a token. This token is then used to 
authenticate the user when performing various requests to the API.

## Read Data
The `GetData` function takes three parameters - the application, 
the service, and the PK value. 
- The application parameter specifies the application within the CloudCIX SAAS - currently, the Go SDK supports IAAS, Membership and Training.
- The service specifies what service you want to read from. For example, the service *server* within the *IAAS* application.
- The PK value specifies the value after the service in the URL. This value will depend on the resource you are trying to access.

Due to the fact that Go is a relatively static language, it was not feasible to 
dynamically generate structs for the JSON data. Instead, the structs have been
defined already in their respective files (see `iaas.go` for IAAS structs,
`membership.go` for Membership structs, and `training.go` for training structs). 

This makes it very easy to access the data in a highly organised, structured 
way.

## Write Data

The `WriteData` function is used to send data to the CloudCIX infrastructure.
- The `method` parameter must be spelled in capital letters. It should be either
  `PATCH`, `POST` or `PUT`.
- The `data` parameter must be of type `map[string]string`. This sounds
  complicated, but it just means that you must declare your JSON payload as
  follows: 
  
  ```
  my_data := map[string]string{"test_key": "test_value"}
  ```

## Delete Data 

The `DeleteData` function handles DELETE requests.

## Usage / Underlying Code

For example, if you were to get the ClientIP value of the first IPMI instance 
(index 0) via the general (non-PK) endpoint, you would access it as follows:

```
body, err := client.ListIpmi()
fmt.Println(body.Content[0].ClientIP, err)
```

On the other hand, if you were to search by ID (which I refer to as PK 
endpoints), you would use the following snippet:

```
body, err := client.GetIpmi("123")
fmt.Println(body.Content.ClientIP)
```

All code accessing the API client itself should be in `main.go`. 
The sample I provided contains the following:

```
package main

import (
	"CloudCIX/api"
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
```

## Miscellaneous / Other
- It is recommended that you do not hardcode your credentials, and instead read 
  them from a configuration file of some sort.

## Developers
- To construct custom structs from JSON data, visit 
  [this website](https://mholt.github.io/json-to-go/). You can use this to 
  update the structs as needed by just copying the JSON from the 
  [CloudCIX REST API](https://docs.cloudcix.com/) and pasting it in.
- Have fun! ~ Ciarán

