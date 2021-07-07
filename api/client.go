package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// The client object that everything operates under.
type CloudCIXClient struct {
	Email, Password, ApiKey, Token, ApiUrl string
}

// Fetches the token.
func (cix CloudCIXClient) GetToken() (string, error) {

	// Reads the information from the configuration.
	json_data := map[string]string{"api_key": cix.ApiKey, "email": cix.Email, "password": cix.Password}
	json_value, err := json.Marshal(json_data)
	if err != nil {
		return "", err
	}

	// Sends the request to fetch the token using the supplied credentials.
	url := fmt.Sprintf("https://membership.%s/auth/login/", cix.ApiUrl)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(json_value))
	if err != nil {
		return "", err
	}

	// Reads the response.
	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	// Formats the response and returns the token.
	var resp map[string]string
	err = json.Unmarshal(data, &resp)
	return resp["token"], err
}

// A general function for reading data using the API.
func (cix CloudCIXClient) GetData(application string, service string, objectId string) ([]byte, error) {

	// Initialises the client and generates the token.
	client := http.Client{}
	token, err := cix.GetToken()
	if err != nil {
		fmt.Print(err)
	}

	// Creates the GET request based on the supplied parameters.
	var url string

	if objectId == "" {
		url = "https://" + application + "." + cix.ApiUrl + "/" + service + "/"
	} else {
		url = "https://" + application + "." + cix.ApiUrl + "/" + service + "/" + objectId + "/"
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Sets the token in the request header and sends the request.
	request.Header.Set("X-Auth-Token", token)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Reads and returns the response.
	data, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	return data, err
}

// A general function for sending data using the API.
func (cix CloudCIXClient) WriteData(application string, service string, objectId string, data map[string]string, method string) ([]byte, error) {

	// Initialises the client using the supplied credentials.
	client := http.Client{}
	token, err := cix.GetToken()
	if err != nil {
		fmt.Print(err)
	}

	// Formats the data to be sent via POST, PATCH or PUT.

	var url string

	if objectId == "" {
		url = "https://" + application + "." + cix.ApiUrl + "/" + service + "/"
	} else {
		url = "https://" + application + "." + cix.ApiUrl + "/" + service + "/" + objectId + "/"
	}

	post_data, _ := json.Marshal(data)

	// Due to the way the library works, we have to send it as a buffer.
	post_data_buffer := bytes.NewBuffer(post_data)
	request, err := http.NewRequest(method, url, post_data_buffer)
	if err != nil {
		return nil, err
	}

	// Sets the token and content headers, then makes the request.
	request.Header.Set("X-Auth-Token", token)
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	// The API usually returns a response containing the created or updated object.
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, err
}

func (cix CloudCIXClient) DeleteData(application string, service string, objectId string) ([]byte, error) {

	client := http.Client{}
	token, err := cix.GetToken()
	if err != nil {
		fmt.Print(err)
	}

	url := "https://" + application + "." + cix.ApiUrl + "/" + service + "/" + objectId + "/"

	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("X-Auth-Token", token)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, err
}

// The format of these functions are relatively constant. Each object (which I will refer to as a Resource) will
// have three functions associated with them - List, Get and Write.
// The List function will return a list of all accessible instances of the specified Resource.
// The Get function takes an ID parameter, and will return the details of the Resource with the corresponding ID.
// The Write function takes a Method parameter, which dictates whether it will be a POST, PATCH or PUT request.
// It also takes a data parameter, which dictates what data will be sent in the request.
// These functions return their respective structs.

// Start of Training functions.
func (cix CloudCIXClient) ListClass() (Class, error) {
	data, err := cix.GetData("training", "class", "")
	if err != nil {
		fmt.Println(err)
	}

	ClassInstance := Class{}
	err = json.Unmarshal(data, &ClassInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ClassInstance, err
}

func (cix CloudCIXClient) GetClass(objectId string) (ClassSpecific, error) {
	data, err := cix.GetData("training", "class", objectId)
	if err != nil {
		fmt.Println(err)
	}

	ClassInstance := ClassSpecific{}
	err = json.Unmarshal(data, &ClassInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ClassInstance, err
}

func (cix CloudCIXClient) WriteClass(objectId string, payload map[string]string, method string) (ClassSpecific, error) {
	data, err := cix.WriteData("training", "class", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	ClassInstance := ClassSpecific{}
	err = json.Unmarshal(data, &ClassInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ClassInstance, err
}

func (cix CloudCIXClient) ListSyllabus() (Syllabus, error) {
	data, err := cix.GetData("training", "syllabus", "")
	if err != nil {
		fmt.Println(err)
	}

	SyllabusInstance := Syllabus{}
	err = json.Unmarshal(data, &SyllabusInstance)
	if err != nil {
		fmt.Println(err)
	}

	return SyllabusInstance, err
}

func (cix CloudCIXClient) GetSyllabus(objectId string) (SyllabusSpecific, error) {
	data, err := cix.GetData("training", "syllabus", objectId)
	if err != nil {
		fmt.Println(err)
	}

	SyllabusInstance := SyllabusSpecific{}
	err = json.Unmarshal(data, &SyllabusInstance)
	if err != nil {
		fmt.Println(err)
	}

	return SyllabusInstance, err
}

func (cix CloudCIXClient) WriteSyllabus(objectId string, payload map[string]string, method string) (SyllabusSpecific, error) {
	data, err := cix.WriteData("training", "syllabus", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	SyllabusInstance := SyllabusSpecific{}
	err = json.Unmarshal(data, &SyllabusInstance)
	if err != nil {
		fmt.Println(err)
	}

	return SyllabusInstance, err

}

func (cix CloudCIXClient) ListStudent() (Student, error) {
	data, err := cix.GetData("training", "student", "")
	if err != nil {
		fmt.Println(err)
	}

	StudentInstance := Student{}
	err = json.Unmarshal(data, &StudentInstance)
	if err != nil {
		fmt.Println(err)
	}

	return StudentInstance, err
}

func (cix CloudCIXClient) GetStudent(objectId string) (StudentSpecific, error) {
	data, err := cix.GetData("training", "student", objectId)
	if err != nil {
		fmt.Println(err)
	}

	StudentInstance := StudentSpecific{}
	err = json.Unmarshal(data, &StudentInstance)
	if err != nil {
		fmt.Println(err)
	}

	return StudentInstance, err
}

func (cix CloudCIXClient) WriteStudent(objectId string, payload map[string]string, method string) (StudentSpecific, error) {
	data, err := cix.WriteData("training", "student", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	StudentInstance := StudentSpecific{}
	err = json.Unmarshal(data, &StudentInstance)
	if err != nil {
		fmt.Println(err)
	}

	return StudentInstance, err
}

// Start of Membership functions.

func (cix CloudCIXClient) ListAddress() (Address, error) {
	data, err := cix.GetData("membership", "address", "")
	if err != nil {
		fmt.Println(err)
	}

	AddressInstance := Address{}
	err = json.Unmarshal(data, &AddressInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AddressInstance, err
}

func (cix CloudCIXClient) GetAddress(objectId string) (AddressSpecific, error) {
	data, err := cix.GetData("membership", "address", objectId)
	if err != nil {
		fmt.Println(err)
	}

	AddressInstance := AddressSpecific{}
	err = json.Unmarshal(data, &AddressInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AddressInstance, err
}

func (cix CloudCIXClient) WriteAddress(objectId string, payload map[string]string, method string) (AddressSpecific, error) {
	data, err := cix.WriteData("membership", "address", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	AddressInstance := AddressSpecific{}
	err = json.Unmarshal(data, &AddressInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AddressInstance, err
}

func (cix CloudCIXClient) GetMembershipAppSettings(objectId string) (MembershipAppSettings, error) {
	data, err := cix.GetData("membership", "app_settings", objectId)
	if err != nil {
		fmt.Println(err)
	}

	AppSettingsInstance := MembershipAppSettings{}
	err = json.Unmarshal(data, &AppSettingsInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AppSettingsInstance, err
}

func (cix CloudCIXClient) WriteMembershipAppSettings(objectId string, payload map[string]string, method string) (MembershipAppSettings, error) {
	data, err := cix.WriteData("membership", "app_settings", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	AppSettingsInstance := MembershipAppSettings{}
	err = json.Unmarshal(data, &AppSettingsInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AppSettingsInstance, err
}

func (cix CloudCIXClient) ListCountry() (Country, error) {
	data, err := cix.GetData("membership", "country", "")
	if err != nil {
		fmt.Println(err)
	}

	CountryInstance := Country{}
	err = json.Unmarshal(data, &CountryInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CountryInstance, err
}

func (cix CloudCIXClient) GetCountry(objectId string) (CountrySpecific, error) {
	data, err := cix.GetData("membership", "country", objectId)
	if err != nil {
		fmt.Println(err)
	}

	CountryInstance := CountrySpecific{}
	err = json.Unmarshal(data, &CountryInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CountryInstance, err
}

func (cix CloudCIXClient) ListCurrency() (Currency, error) {
	data, err := cix.GetData("membership", "currency", "")
	if err != nil {
		fmt.Println(err)
	}

	CurrencyInstance := Currency{}
	err = json.Unmarshal(data, &CurrencyInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CurrencyInstance, err
}

func (cix CloudCIXClient) GetCurrency(objectId string) (CurrencySpecific, error) {
	data, err := cix.GetData("membership", "currency", objectId)
	if err != nil {
		fmt.Println(err)
	}

	CurrencyInstance := CurrencySpecific{}
	err = json.Unmarshal(data, &CurrencyInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CurrencyInstance, err
}

func (cix CloudCIXClient) ListDepartment() (Department, error) {
	data, err := cix.GetData("membership", "department", "")
	if err != nil {
		fmt.Println(err)
	}

	DepartmentInstance := Department{}
	err = json.Unmarshal(data, &DepartmentInstance)
	if err != nil {
		fmt.Println(err)
	}

	return DepartmentInstance, err
}

func (cix CloudCIXClient) GetDepartment(objectId string) (DepartmentSpecific, error) {
	data, err := cix.GetData("membership", "department", objectId)
	if err != nil {
		fmt.Println(err)
	}

	DepartmentInstance := DepartmentSpecific{}
	err = json.Unmarshal(data, &DepartmentInstance)
	if err != nil {
		fmt.Println(err)
	}

	return DepartmentInstance, err
}

func (cix CloudCIXClient) WriteDepartment(objectId string, payload map[string]string, method string) (DepartmentSpecific, error) {
	data, err := cix.WriteData("membership", "department", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	DepartmentInstance := DepartmentSpecific{}
	err = json.Unmarshal(data, &DepartmentInstance)
	if err != nil {
		fmt.Println(err)
	}

	return DepartmentInstance, err
}

func (cix CloudCIXClient) ListLanguage() (Language, error) {
	data, err := cix.GetData("membership", "language", "")
	if err != nil {
		fmt.Println(err)
	}

	LanguageInstance := Language{}
	err = json.Unmarshal(data, &LanguageInstance)
	if err != nil {
		fmt.Println(err)
	}

	return LanguageInstance, err
}

func (cix CloudCIXClient) GetLanguage(objectId string) (LanguageSpecific, error) {
	data, err := cix.GetData("membership", "language", objectId)
	if err != nil {
		fmt.Println(err)
	}

	LanguageInstance := LanguageSpecific{}
	err = json.Unmarshal(data, &LanguageInstance)
	if err != nil {
		fmt.Println(err)
	}

	return LanguageInstance, err
}

func (cix CloudCIXClient) GetMember(objectId string) (MemberSpecific, error) {
	data, err := cix.GetData("membership", "member", objectId)
	if err != nil {
		fmt.Println(err)
	}

	MemberInstance := MemberSpecific{}
	err = json.Unmarshal(data, &MemberInstance)
	if err != nil {
		fmt.Println(err)
	}

	return MemberInstance, err
}

func (cix CloudCIXClient) WriteMember(objectId string, payload map[string]string, method string) (MemberSpecific, error) {
	data, err := cix.WriteData("membership", "member", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	MemberInstance := MemberSpecific{}
	err = json.Unmarshal(data, &MemberInstance)
	if err != nil {
		fmt.Println(err)
	}

	return MemberInstance, err
}

func (cix CloudCIXClient) ListProfile() (Profile, error) {
	data, err := cix.GetData("membership", "profile", "")
	if err != nil {
		fmt.Println(err)
	}

	ProfileInstance := Profile{}
	err = json.Unmarshal(data, &ProfileInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ProfileInstance, err
}

func (cix CloudCIXClient) GetProfile(objectId string) (ProfileSpecific, error) {
	data, err := cix.GetData("membership", "profile", objectId)
	if err != nil {
		fmt.Println(err)
	}

	ProfileInstance := ProfileSpecific{}
	err = json.Unmarshal(data, &ProfileInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ProfileInstance, err
}

func (cix CloudCIXClient) WriteProfile(objectId string, payload map[string]string, method string) (ProfileSpecific, error) {
	data, err := cix.WriteData("membership", "profile", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	ProfileInstance := ProfileSpecific{}
	err = json.Unmarshal(data, &ProfileInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ProfileInstance, err
}

func (cix CloudCIXClient) ListTeam() (Team, error) {
	data, err := cix.GetData("membership", "team", "")
	if err != nil {
		fmt.Println(err)
	}

	TeamInstance := Team{}
	err = json.Unmarshal(data, &TeamInstance)
	if err != nil {
		fmt.Println(err)
	}

	return TeamInstance, err
}

func (cix CloudCIXClient) GetTeam(objectId string) (TeamSpecific, error) {
	data, err := cix.GetData("membership", "team", objectId)
	if err != nil {
		fmt.Println(err)
	}

	TeamInstance := TeamSpecific{}
	err = json.Unmarshal(data, &TeamInstance)
	if err != nil {
		fmt.Println(err)
	}

	return TeamInstance, err
}

func (cix CloudCIXClient) WriteTeam(objectId string, payload map[string]string, method string) (TeamSpecific, error) {
	data, err := cix.WriteData("membership", "team", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	TeamInstance := TeamSpecific{}
	err = json.Unmarshal(data, &TeamInstance)
	if err != nil {
		fmt.Println(err)
	}

	return TeamInstance, err
}

func (cix CloudCIXClient) ListTerritory() (Territory, error) {
	data, err := cix.GetData("membership", "territory", "")
	if err != nil {
		fmt.Println(err)
	}

	TerritoryInstance := Territory{}
	err = json.Unmarshal(data, &TerritoryInstance)
	if err != nil {
		fmt.Println(err)
	}

	return TerritoryInstance, err
}

func (cix CloudCIXClient) GetTerritory(objectId string) (TerritorySpecific, error) {
	data, err := cix.GetData("membership", "territory", objectId)
	if err != nil {
		fmt.Println(err)
	}

	TerritoryInstance := TerritorySpecific{}
	err = json.Unmarshal(data, &TerritoryInstance)
	if err != nil {
		fmt.Println(err)
	}

	return TerritoryInstance, err
}

func (cix CloudCIXClient) WriteTerritory(objectId string, payload map[string]string, method string) (TerritorySpecific, error) {
	data, err := cix.WriteData("membership", "territory", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	TerritoryInstance := TerritorySpecific{}
	err = json.Unmarshal(data, &TerritoryInstance)
	if err != nil {
		fmt.Println(err)
	}

	return TerritoryInstance, err
}

func (cix CloudCIXClient) ListTransactionType() (TransactionType, error) {
	data, err := cix.GetData("membership", "transaction_type", "")
	if err != nil {
		fmt.Println(err)
	}

	TransactionTypeInstance := TransactionType{}
	err = json.Unmarshal(data, &TransactionTypeInstance)
	if err != nil {
		fmt.Println(err)
	}

	return TransactionTypeInstance, err
}

func (cix CloudCIXClient) GetTransactionType(objectId string) (TransactionTypeSpecific, error) {
	data, err := cix.GetData("membership", "transaction_type", objectId)
	if err != nil {
		fmt.Println(err)
	}

	TransactionTypeInstance := TransactionTypeSpecific{}
	err = json.Unmarshal(data, &TransactionTypeInstance)
	if err != nil {
		fmt.Println(err)
	}

	return TransactionTypeInstance, err
}

func (cix CloudCIXClient) ListUser() (User, error) {
	data, err := cix.GetData("membership", "user", "")
	if err != nil {
		fmt.Println(err)
	}

	UserInstance := User{}
	err = json.Unmarshal(data, &UserInstance)
	if err != nil {
		fmt.Println(err)
	}

	return UserInstance, err
}

func (cix CloudCIXClient) GetUser(objectId string) (UserSpecific, error) {
	data, err := cix.GetData("membership", "user", objectId)
	if err != nil {
		fmt.Println(err)
	}

	UserInstance := UserSpecific{}
	err = json.Unmarshal(data, &UserInstance)
	if err != nil {
		fmt.Println(err)
	}

	return UserInstance, err
}

func (cix CloudCIXClient) WriteUser(objectId string, payload map[string]string, method string) (UserSpecific, error) {
	data, err := cix.WriteData("membership", "user", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	UserInstance := UserSpecific{}
	err = json.Unmarshal(data, &UserInstance)
	if err != nil {
		fmt.Println(err)
	}

	return UserInstance, err
}

//Start of IAAS functions.

func (cix CloudCIXClient) ListAllocation() (Allocation, error) {
	data, err := cix.GetData("iaas", "allocation", "")
	if err != nil {
		fmt.Println(err)
	}

	AllocationInstance := Allocation{}
	err = json.Unmarshal(data, &AllocationInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AllocationInstance, err
}

func (cix CloudCIXClient) GetAllocation(objectId string) (AllocationSpecific, error) {
	data, err := cix.GetData("iaas", "allocation", objectId)
	if err != nil {
		fmt.Println(err)
	}
	AllocationInstance := AllocationSpecific{}
	err = json.Unmarshal(data, &AllocationInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AllocationInstance, err
}

func (cix CloudCIXClient) WriteAllocation(objectId string, payload map[string]string, method string) (AllocationSpecific, error) {
	data, err := cix.WriteData("iaas", "allocation", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	AllocationInstance := AllocationSpecific{}
	err = json.Unmarshal(data, &AllocationInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AllocationInstance, err

}

func (cix CloudCIXClient) GetIAASAppSettings(objectId string) (IAASAppSettings, error) {
	data, err := cix.GetData("iaas", "app_settings", objectId)
	if err != nil {
		fmt.Println(err)
	}

	AppSettingsInstance := IAASAppSettings{}
	err = json.Unmarshal(data, &AppSettingsInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AppSettingsInstance, err
}

func (cix CloudCIXClient) WriteIAASAppSettings(objectId string, payload map[string]string, method string) (IAASAppSettings, error) {
	data, err := cix.WriteData("iaas", "app_settings", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	AppSettingsInstance := IAASAppSettings{}
	err = json.Unmarshal(data, &AppSettingsInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AppSettingsInstance, err
}

func (cix CloudCIXClient) ListAsn() (Asn, error) {
	data, err := cix.GetData("iaas", "asn", "")
	if err != nil {
		fmt.Println(err)
	}
	AsnInstance := Asn{}
	err = json.Unmarshal(data, &AsnInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AsnInstance, err
}

func (cix CloudCIXClient) GetAsn(objectId string) (AsnSpecific, error) {
	data, err := cix.GetData("iaas", "asn", objectId)
	if err != nil {
		fmt.Println(err)
	}

	AsnInstance := AsnSpecific{}
	err = json.Unmarshal(data, &AsnInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AsnInstance, err
}

func (cix CloudCIXClient) WriteAsn(objectId string, payload map[string]string, method string) (AsnSpecific, error) {
	data, err := cix.WriteData("iaas", "asn", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	AsnInstance := AsnSpecific{}
	err = json.Unmarshal(data, &AsnInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AsnInstance, err

}

func (cix CloudCIXClient) ListCixBlacklist() (CixBlacklist, error) {
	data, err := cix.GetData("iaas", "cix_blacklist", "")
	if err != nil {
		fmt.Println(err)
	}

	CixBlacklistInstance := CixBlacklist{}
	err = json.Unmarshal(data, &CixBlacklistInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CixBlacklistInstance, err
}

func (cix CloudCIXClient) GetCixBlacklist(objectId string) (CixBlacklistSpecific, error) {
	data, err := cix.GetData("iaas", "cix_blacklist", objectId)
	if err != nil {
		fmt.Println(err)
	}

	CixBlacklistInstance := CixBlacklistSpecific{}
	err = json.Unmarshal(data, &CixBlacklistInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CixBlacklistInstance, err
}

func (cix CloudCIXClient) WriteCixBlacklist(objectId string, payload map[string]string, method string) (CixBlacklistSpecific, error) {
	data, err := cix.WriteData("iaas", "cix_blacklist", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	CixBlacklistInstance := CixBlacklistSpecific{}
	err = json.Unmarshal(data, &CixBlacklistInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CixBlacklistInstance, err

}

func (cix CloudCIXClient) ListCixWhitelist() (CixWhitelist, error) {
	data, err := cix.GetData("iaas", "cix_whitelist", "")
	if err != nil {
		fmt.Println(err)
	}

	CixWhitelistInstance := CixWhitelist{}
	err = json.Unmarshal(data, &CixWhitelistInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CixWhitelistInstance, err
}

func (cix CloudCIXClient) GetCixWhitelist(objectId string) (CixWhitelistSpecific, error) {
	data, err := cix.GetData("iaas", "cix_whitelist", objectId)
	if err != nil {
		fmt.Println(err)
	}

	CixWhitelistInstance := CixWhitelistSpecific{}
	err = json.Unmarshal(data, &CixWhitelistInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CixWhitelistInstance, err
}

func (cix CloudCIXClient) WriteCiXWhitelist(objectId string, payload map[string]string, method string) (CixWhitelistSpecific, error) {
	data, err := cix.WriteData("iaas", "cix_whitelist", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	CixWhitelistInstance := CixWhitelistSpecific{}
	err = json.Unmarshal(data, &CixWhitelistInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CixWhitelistInstance, err
}

func (cix CloudCIXClient) ListCloud() (Cloud, error) {
	data, err := cix.GetData("iaas", "cloud", "")
	if err != nil {
		fmt.Println(err)
	}

	CloudInstance := Cloud{}
	err = json.Unmarshal(data, &CloudInstance)
	if err != nil {
		fmt.Println(err)
	}

	return CloudInstance, err
}

func (cix CloudCIXClient) WriteCloud(objectId string, payload map[string]string, method string) (Cloud, error) {
	data, err := cix.WriteData("iaas", "cloud", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	cloudInstance := Cloud{}
	err = json.Unmarshal(data, &cloudInstance)
	if err != nil {
		fmt.Println(err)
	}

	return cloudInstance, err

}

func (cix CloudCIXClient) ListDomain() (Domain, error) {
	data, err := cix.GetData("iaas", "domain", "")
	if err != nil {
		fmt.Println(err)
	}

	DomainInstance := Domain{}
	err = json.Unmarshal(data, &DomainInstance)
	if err != nil {
		fmt.Println(err)
	}

	return DomainInstance, err
}

func (cix CloudCIXClient) GetDomain(objectId string) (DomainSpecific, error) {
	data, err := cix.GetData("iaas", "domain", objectId)
	if err != nil {
		fmt.Println(err)
	}

	DomainInstance := DomainSpecific{}
	err = json.Unmarshal(data, &DomainInstance)
	if err != nil {
		fmt.Println(err)
	}

	return DomainInstance, err
}

func (cix CloudCIXClient) WriteDomain(objectId string, payload map[string]string, method string) (DomainSpecific, error) {
	data, err := cix.WriteData("iaas", "domain", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	domainInstance := DomainSpecific{}
	err = json.Unmarshal(data, &domainInstance)
	if err != nil {
		fmt.Println(err)
	}

	return domainInstance, err

}

func (cix CloudCIXClient) ListImage() (Image, error) {
	data, err := cix.GetData("iaas", "image", "")
	if err != nil {
		fmt.Println(err)
	}

	ImageInstance := Image{}
	err = json.Unmarshal(data, &ImageInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ImageInstance, err
}

func (cix CloudCIXClient) GetImage(objectId string) (ImageSpecific, error) {
	data, err := cix.GetData("iaas", "image", objectId)
	if err != nil {
		fmt.Println(err)
	}

	ImageInstance := ImageSpecific{}
	err = json.Unmarshal(data, &ImageInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ImageInstance, err
}

func (cix CloudCIXClient) WriteImage(objectId string, payload map[string]string, method string) (ImageSpecific, error) {
	data, err := cix.WriteData("iaas", "image", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	imageInstance := ImageSpecific{}
	err = json.Unmarshal(data, &imageInstance)
	if err != nil {
		fmt.Println(err)
	}

	return imageInstance, err

}

func (cix CloudCIXClient) ListInterface() (Interface, error) {
	data, err := cix.GetData("iaas", "interface", "")
	if err != nil {
		fmt.Println(err)
	}

	InterfaceInstance := Interface{}
	err = json.Unmarshal(data, &InterfaceInstance)
	if err != nil {
		fmt.Println(err)
	}

	return InterfaceInstance, err
}

func (cix CloudCIXClient) GetInterface(objectId string) (InterfaceSpecific, error) {
	data, err := cix.GetData("iaas", "interface", objectId)
	if err != nil {
		fmt.Println(err)
	}

	InterfaceInstance := InterfaceSpecific{}
	err = json.Unmarshal(data, &InterfaceInstance)
	if err != nil {
		fmt.Println(err)
	}

	return InterfaceInstance, err
}

func (cix CloudCIXClient) WriteInterface(objectId string, payload map[string]string, method string) (InterfaceSpecific, error) {
	data, err := cix.WriteData("iaas", "interface", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	interfaceInstance := InterfaceSpecific{}
	err = json.Unmarshal(data, &interfaceInstance)
	if err != nil {
		fmt.Println(err)
	}

	return interfaceInstance, err

}

func (cix CloudCIXClient) ListIpAddress() (IpAddress, error) {
	data, err := cix.GetData("iaas", "ip_address", "")
	if err != nil {
		fmt.Println(err)
	}

	IpAddressInstance := IpAddress{}
	err = json.Unmarshal(data, &IpAddressInstance)
	if err != nil {
		fmt.Println(err)
	}

	return IpAddressInstance, err
}

func (cix CloudCIXClient) GetIpAddress(objectId string) (IpAddressSpecific, error) {
	data, err := cix.GetData("iaas", "ip_address", objectId)
	if err != nil {
		fmt.Println(err)
	}

	IpAddressInstance := IpAddressSpecific{}
	err = json.Unmarshal(data, &IpAddressInstance)
	if err != nil {
		fmt.Println(err)
	}

	return IpAddressInstance, err
}

func (cix CloudCIXClient) WriteIpAddress(objectId string, payload map[string]string, method string) (IpAddressSpecific, error) {
	data, err := cix.WriteData("iaas", "ip_address", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	IpAddressInstance := IpAddressSpecific{}
	err = json.Unmarshal(data, &IpAddressInstance)
	if err != nil {
		fmt.Println(err)
	}

	return IpAddressInstance, err

}

func (cix CloudCIXClient) ListIpmi() (Ipmi, error) {
	data, err := cix.GetData("iaas", "ipmi", "")
	if err != nil {
		fmt.Println(err)
	}

	IpmiInstance := Ipmi{}
	err = json.Unmarshal(data, &IpmiInstance)
	if err != nil {
		fmt.Println(err)
	}

	return IpmiInstance, err
}

func (cix CloudCIXClient) GetIpmi(objectId string) (IpmiSpecific, error) {
	data, err := cix.GetData("iaas", "ipmi", objectId)
	if err != nil {
		fmt.Println(err)
	}

	IpmiInstance := IpmiSpecific{}
	err = json.Unmarshal(data, &IpmiInstance)
	if err != nil {
		fmt.Println(err)
	}

	return IpmiInstance, err
}

func (cix CloudCIXClient) WriteIpmi(objectId string, payload map[string]string, method string) (IpmiSpecific, error) {
	data, err := cix.WriteData("iaas", "ipmi", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	IpmiInstance := IpmiSpecific{}
	err = json.Unmarshal(data, &IpmiInstance)
	if err != nil {
		fmt.Println(err)
	}

	return IpmiInstance, err

}

func (cix CloudCIXClient) ListPoolIp() (PoolIp, error) {
	data, err := cix.GetData("iaas", "pool_ip", "")
	if err != nil {
		fmt.Println(err)
	}

	PoolIpInstance := PoolIp{}
	err = json.Unmarshal(data, &PoolIpInstance)
	if err != nil {
		fmt.Println(err)
	}

	return PoolIpInstance, err
}

func (cix CloudCIXClient) GetPoolIp(objectId string) (PoolIpSpecific, error) {
	data, err := cix.GetData("iaas", "pool_ip", objectId)
	if err != nil {
		fmt.Println(err)
	}

	PoolIpInstance := PoolIpSpecific{}
	err = json.Unmarshal(data, &PoolIpInstance)
	if err != nil {
		fmt.Println(err)
	}

	return PoolIpInstance, err
}

func (cix CloudCIXClient) WritePoolIp(objectId string, payload map[string]string, method string) (PoolIpSpecific, error) {
	data, err := cix.WriteData("iaas", "pool_ip", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	PoolIpInstance := PoolIpSpecific{}
	err = json.Unmarshal(data, &PoolIpInstance)
	if err != nil {
		fmt.Println(err)
	}

	return PoolIpInstance, err

}

func (cix CloudCIXClient) ListProject() (Project, error) {
	data, err := cix.GetData("iaas", "project", "")
	if err != nil {
		fmt.Println(err)
	}

	ProjectInstance := Project{}
	err = json.Unmarshal(data, &ProjectInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ProjectInstance, err
}

func (cix CloudCIXClient) GetProject(objectId string) (ProjectSpecific, error) {
	data, err := cix.GetData("iaas", "project", objectId)
	if err != nil {
		fmt.Println(err)
	}

	ProjectInstance := ProjectSpecific{}
	err = json.Unmarshal(data, &ProjectInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ProjectInstance, err
}

func (cix CloudCIXClient) WriteProject(objectId string, payload map[string]string, method string) (ProjectSpecific, error) {
	data, err := cix.WriteData("iaas", "project", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	ProjectInstance := ProjectSpecific{}
	err = json.Unmarshal(data, &ProjectInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ProjectInstance, err

}

func (cix CloudCIXClient) ListPtrRecord() (PtrRecord, error) {
	data, err := cix.GetData("iaas", "ptr_record", "")
	if err != nil {
		fmt.Println(err)
	}

	PtrRecordInstance := PtrRecord{}
	err = json.Unmarshal(data, &PtrRecordInstance)
	if err != nil {
		fmt.Println(err)
	}

	return PtrRecordInstance, err
}

func (cix CloudCIXClient) GetPtrRecord(objectId string) (PtrRecordSpecific, error) {
	data, err := cix.GetData("iaas", "ptr_record", objectId)
	if err != nil {
		fmt.Println(err)
	}

	PtrRecordInstance := PtrRecordSpecific{}
	err = json.Unmarshal(data, &PtrRecordInstance)
	if err != nil {
		fmt.Println(err)
	}

	return PtrRecordInstance, err
}

func (cix CloudCIXClient) WritePtrRecord(objectId string, payload map[string]string, method string) (PtrRecordSpecific, error) {
	data, err := cix.WriteData("iaas", "ptr_record", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	PtrRecordInstance := PtrRecordSpecific{}
	err = json.Unmarshal(data, &PtrRecordInstance)
	if err != nil {
		fmt.Println(err)
	}

	return PtrRecordInstance, err

}

func (cix CloudCIXClient) ListRecord() (Record, error) {
	data, err := cix.GetData("iaas", "record", "")
	if err != nil {
		fmt.Println(err)
	}

	RecordInstance := Record{}
	err = json.Unmarshal(data, &RecordInstance)
	if err != nil {
		fmt.Println(err)
	}

	return RecordInstance, err
}

func (cix CloudCIXClient) GetRecord(objectId string) (RecordSpecific, error) {
	data, err := cix.GetData("iaas", "record", objectId)
	if err != nil {
		fmt.Println(err)
	}

	RecordInstance := RecordSpecific{}
	err = json.Unmarshal(data, &RecordInstance)
	if err != nil {
		fmt.Println(err)
	}

	return RecordInstance, err
}

func (cix CloudCIXClient) WriteRecord(objectId string, payload map[string]string, method string) (RecordSpecific, error) {
	data, err := cix.WriteData("iaas", "record", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	RecordInstance := RecordSpecific{}
	err = json.Unmarshal(data, &RecordInstance)
	if err != nil {
		fmt.Println(err)
	}

	return RecordInstance, err

}

func (cix CloudCIXClient) ListRouter() (Router, error) {
	data, err := cix.GetData("iaas", "router", "")
	if err != nil {
		fmt.Println(err)
	}

	RouterInstance := Router{}
	err = json.Unmarshal(data, &RouterInstance)
	if err != nil {
		fmt.Println(err)
	}

	return RouterInstance, err
}

func (cix CloudCIXClient) GetRouter(objectId string) (RouterSpecific, error) {
	data, err := cix.GetData("iaas", "router", objectId)
	if err != nil {
		fmt.Println(err)
	}

	RouterInstance := RouterSpecific{}
	err = json.Unmarshal(data, &RouterInstance)
	if err != nil {
		fmt.Println(err)
	}

	return RouterInstance, err
}

func (cix CloudCIXClient) WriteRouter(objectId string, payload map[string]string, method string) (RouterSpecific, error) {
	data, err := cix.WriteData("iaas", "router", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	RouterInstance := RouterSpecific{}
	err = json.Unmarshal(data, &RouterInstance)
	if err != nil {
		fmt.Println(err)
	}

	return RouterInstance, err

}

func (cix CloudCIXClient) ListServer() (Server, error) {
	data, err := cix.GetData("iaas", "server", "")
	if err != nil {
		fmt.Println(err)
	}

	ServerInstance := Server{}
	err = json.Unmarshal(data, &ServerInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ServerInstance, err
}

func (cix CloudCIXClient) GetServer(objectId string) (ServerSpecific, error) {
	data, err := cix.GetData("iaas", "server", objectId)
	if err != nil {
		fmt.Println(err)
	}

	ServerInstance := ServerSpecific{}
	err = json.Unmarshal(data, &ServerInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ServerInstance, err
}

func (cix CloudCIXClient) WriteServer(objectId string, payload map[string]string, method string) (ServerSpecific, error) {
	data, err := cix.WriteData("iaas", "server", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	ServerInstance := ServerSpecific{}
	err = json.Unmarshal(data, &ServerInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ServerInstance, err

}

func (cix CloudCIXClient) ListServerType() (ServerType, error) {
	data, err := cix.GetData("iaas", "server_type", "")
	if err != nil {
		fmt.Println(err)
	}

	ServerTypeInstance := ServerType{}
	err = json.Unmarshal(data, &ServerTypeInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ServerTypeInstance, err
}

func (cix CloudCIXClient) GetServerType(objectId string) (ServerTypeSpecific, error) {
	data, err := cix.GetData("iaas", "server_type", objectId)
	if err != nil {
		fmt.Println(err)
	}

	ServerTypeInstance := ServerTypeSpecific{}
	err = json.Unmarshal(data, &ServerTypeInstance)
	if err != nil {
		fmt.Println(err)
	}

	return ServerTypeInstance, err
}

func (cix CloudCIXClient) ListStorageType() (StorageType, error) {
	data, err := cix.GetData("iaas", "storage_type", "")
	if err != nil {
		fmt.Println(err)
	}

	StorageTypeInstance := StorageType{}
	err = json.Unmarshal(data, &StorageTypeInstance)
	if err != nil {
		fmt.Println(err)
	}

	return StorageTypeInstance, err
}

func (cix CloudCIXClient) GetStorageType(objectId string) (StorageTypeSpecific, error) {
	data, err := cix.GetData("iaas", "storage_type", objectId)
	if err != nil {
		fmt.Println(err)
	}

	StorageTypeInstance := StorageTypeSpecific{}
	err = json.Unmarshal(data, &StorageTypeInstance)
	if err != nil {
		fmt.Println(err)
	}

	return StorageTypeInstance, err
}

func (cix CloudCIXClient) WriteStorageType(objectId string, payload map[string]string, method string) (StorageTypeSpecific, error) {
	data, err := cix.WriteData("iaas", "storage_type", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	StorageTypeInstance := StorageTypeSpecific{}
	err = json.Unmarshal(data, &StorageTypeInstance)
	if err != nil {
		fmt.Println(err)
	}

	return StorageTypeInstance, err

}

func (cix CloudCIXClient) ListSubnet() (Subnet, error) {
	data, err := cix.GetData("iaas", "subnet", "")
	if err != nil {
		fmt.Println(err)
	}

	SubnetInstance := Subnet{}
	err = json.Unmarshal(data, &SubnetInstance)
	if err != nil {
		fmt.Println(err)
	}

	return SubnetInstance, err
}

func (cix CloudCIXClient) GetSubnet(objectId string) (SubnetSpecific, error) {
	data, err := cix.GetData("iaas", "subnet", objectId)
	if err != nil {
		fmt.Println(err)
	}

	SubnetInstance := SubnetSpecific{}
	err = json.Unmarshal(data, &SubnetInstance)
	if err != nil {
		fmt.Println(err)
	}

	return SubnetInstance, err
}

func (cix CloudCIXClient) WriteSubnet(objectId string, payload map[string]string, method string) (SubnetSpecific, error) {
	data, err := cix.WriteData("iaas", "subnet", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	SubnetInstance := SubnetSpecific{}
	err = json.Unmarshal(data, &SubnetInstance)
	if err != nil {
		fmt.Println(err)
	}

	return SubnetInstance, err

}

func (cix CloudCIXClient) ListVirtualRouter() (VirtualRouter, error) {
	data, err := cix.GetData("iaas", "virtual_router", "")
	if err != nil {
		fmt.Println(err)
	}

	VirtualRouterInstance := VirtualRouter{}
	err = json.Unmarshal(data, &VirtualRouterInstance)
	if err != nil {
		fmt.Println(err)
	}

	return VirtualRouterInstance, err
}

func (cix CloudCIXClient) GetVirtualRouter(objectId string) (VirtualRouterSpecific, error) {
	data, err := cix.GetData("iaas", "virtual_router", objectId)
	if err != nil {
		fmt.Println(err)
	}

	VirtualRouterInstance := VirtualRouterSpecific{}
	err = json.Unmarshal(data, &VirtualRouterInstance)
	if err != nil {
		fmt.Println(err)
	}

	return VirtualRouterInstance, err
}

func (cix CloudCIXClient) WriteVirtualRouter(objectId string, payload map[string]string, method string) (VirtualRouterSpecific, error) {
	data, err := cix.WriteData("iaas", "virtual_router", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	VirtualRouterInstance := VirtualRouterSpecific{}
	err = json.Unmarshal(data, &VirtualRouterInstance)
	if err != nil {
		fmt.Println(err)
	}

	return VirtualRouterInstance, err

}

func (cix CloudCIXClient) ListVm() (Vm, error) {
	data, err := cix.GetData("iaas", "vm", "")
	if err != nil {
		fmt.Println(err)
	}

	VmInstance := Vm{}
	err = json.Unmarshal(data, &VmInstance)
	if err != nil {
		fmt.Println(err)
	}

	return VmInstance, err
}

func (cix CloudCIXClient) GetVm(objectId string) (VmSpecific, error) {
	data, err := cix.GetData("iaas", "vm", objectId)
	if err != nil {
		fmt.Println(err)
	}

	VmInstance := VmSpecific{}
	err = json.Unmarshal(data, &VmInstance)
	if err != nil {
		fmt.Println(err)
	}

	return VmInstance, err
}

func (cix CloudCIXClient) WriteVm(objectId string, payload map[string]string, method string) (VmSpecific, error) {
	data, err := cix.WriteData("iaas", "vm", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	VmInstance := VmSpecific{}
	err = json.Unmarshal(data, &VmInstance)
	if err != nil {
		fmt.Println(err)
	}

	return VmInstance, err

}

func (cix CloudCIXClient) ListVpn() (Vpn, error) {
	data, err := cix.GetData("iaas", "vpn", "")
	if err != nil {
		fmt.Println(err)
	}

	VpnInstance := Vpn{}
	err = json.Unmarshal(data, &VpnInstance)
	if err != nil {
		fmt.Println(err)
	}

	return VpnInstance, err
}

func (cix CloudCIXClient) GetVpn(objectId string) (VpnSpecific, error) {
	data, err := cix.GetData("iaas", "vpn", objectId)
	if err != nil {
		fmt.Println(err)
	}

	VpnInstance := VpnSpecific{}
	err = json.Unmarshal(data, &VpnInstance)
	if err != nil {
		fmt.Println(err)
	}

	return VpnInstance, err
}

func (cix CloudCIXClient) WriteVpn(objectId string, payload map[string]string, method string) (VpnSpecific, error) {
	data, err := cix.WriteData("iaas", "vpn", objectId, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	VpnInstance := VpnSpecific{}
	err = json.Unmarshal(data, &VpnInstance)
	if err != nil {
		fmt.Println(err)
	}

	return VpnInstance, err

}
