# CloudCIX Go SDK

An implementation of the CloudCIX API in Go. When I refer to "PK endpoints", I
am referring to the endpoints that take an ID value after the service - for 
example, "/user/1", where 1 is the PK value. The use of the term "PK" is just 
to mimic the CloudCIX Python library. The Go SDK is logically divided into 
the reading/writing functionality, the structs, and the client functions. 
Every "resource" - that being an object - has a corresponding List, Get and Write 
function where applicable. Each function returns its corresponding struct, containing
the information needed.

## Authentication
To begin with, configure the client by setting the necessary properties of the 
object (see the `main.go` example).
The `GetToken` function generates a token. This token is then used to 
authenticate the user when performing various requests to the API.

## Read Data

The `GetData` function takes three parameters - the application, 
the service, and the PK value. 
- The application parameter specifies the application within the CloudCIX SAAS, 
  such as IAAS or Membership. 
- The service specifies what service you want to read from.
- The PK value specifies the value after the service in the URL. 
  This value will depend on the resource you are trying to access.

Due to the fact that Go is a relatively static language, it was not feasible to 
dynamically generate structs for the JSON data. Instead, the structs have been
defined already in their respective files (see `iaas.go` for IAAS structs,
`membership.go` for Membership structs, and so on). 

This makes it very easy to access the data in a highly organised, structured 
way.

## Write Data

The `WriteData` function takes a parameter called `method` to decide whether to
send a PATCH, POST or PUT request. The usage of the `WriteData` function is not
complicated.
- The `method` parameter must be spelled in capital letters. It should be either
  `PATCH`, `POST` or `PUT`.
- The `data` parameter must be of type `map[string]string`. This sounds
  complicated, but it just means that you must declare your JSON payload as
  follows: 
  ```
  my_data := map[string]string{"test": "test"}
  ```

## Delete Data 

The delete functionality is very simple. It is almost identical in structure
to a GET request. The `DeleteData` function handles DELETE requests.

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

