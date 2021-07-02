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
func (cix CloudCIXClient) GetData(application string, service string, object_id string) ([]byte, error) {

	// Initialises the client and generates the token.
	client := http.Client{}
	token, err := cix.GetToken()
	if err != nil {
		fmt.Print(err)
	}

	// Creates the GET request based on the supplied parameters.
	url := "https://" + application + "." + cix.ApiUrl + "/" + service + "/" + object_id
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
func (cix CloudCIXClient) WriteData(application string, service string, object_id string, data map[string]string, method string) ([]byte, error) {

	// Initialises the client using the supplied credentials.
	client := http.Client{}
	token, err := cix.GetToken()
	if err != nil {
		fmt.Print(err)
	}

	// Formats the data to be sent via POST, PATCH or PUT.
	url := "https://" + application + "." + cix.ApiUrl + "/" + service + "/" + object_id
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

func (cix CloudCIXClient) GetClass(object_id string) (ClassSpecific, error) {
	data, err := cix.GetData("training", "class", object_id)
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

func (cix CloudCIXClient) WriteClass(object_id string, payload map[string]string, method string) (ClassSpecific, error) {
	data, err := cix.WriteData("training", "class", object_id, payload, method)
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

func (cix CloudCIXClient) GetSyllabus(object_id string) (SyllabusSpecific, error) {
	data, err := cix.GetData("training", "syllabus", object_id)
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

func (cix CloudCIXClient) WriteSyllabus(object_id string, payload map[string]string, method string) (SyllabusSpecific, error) {
	data, err := cix.WriteData("training", "syllabus", object_id, payload, method)
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

func (cix CloudCIXClient) GetStudent(object_id string) (StudentSpecific, error) {
	data, err := cix.GetData("training", "student", "object_id")
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

func (cix CloudCIXClient) WriteStudent(object_id string, payload map[string]string, method string) (StudentSpecific, error) {
	data, err := cix.WriteData("training", "student", object_id, payload, method)
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

func (cix CloudCIXClient) GetAddress(object_id string) (AddressSpecific, error) {
	data, err := cix.GetData("membership", "address", object_id)
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

func (cix CloudCIXClient) WriteAddress(object_id string, payload map[string]string, method string) (AddressSpecific, error) {
	data, err := cix.WriteData("membership", "address", object_id, payload, method)
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

func (cix CloudCIXClient) GetAppSettings(object_id string) (AppSettings, error) {
	data, err := cix.GetData("membership", "app_settings", object_id)
	if err != nil {
		fmt.Println(err)
	}

	AppSettingsInstance := AppSettings{}
	err = json.Unmarshal(data, &AppSettingsInstance)
	if err != nil {
		fmt.Println(err)
	}

	return AppSettingsInstance, err
}

func (cix CloudCIXClient) WriteAppSettings(object_id string, payload map[string]string, method string) (AppSettings, error) {
	data, err := cix.WriteData("membership", "app_settings", object_id, payload, method)
	if err != nil {
		fmt.Println(err)
	}

	AppSettingsInstance := AppSettings{}
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

func (cix CloudCIXClient) GetCountry(object_id string) (CountrySpecific, error) {
	data, err := cix.GetData("membership", "country", object_id)
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

func (cix CloudCIXClient) GetCurrency(object_id string) (CurrencySpecific, error) {
	data, err := cix.GetData("membership", "currency", object_id)
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

func (cix CloudCIXClient) GetDepartment(object_id string) (DepartmentSpecific, error) {
	data, err := cix.GetData("membership", "department", object_id)
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

func (cix CloudCIXClient) WriteDepartment(object_id string, payload map[string]string, method string) (DepartmentSpecific, error) {
	data, err := cix.WriteData("membership", "department", object_id, payload, method)
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

func (cix CloudCIXClient) GetLanguage(object_id string) (LanguageSpecific, error) {
	data, err := cix.GetData("membership", "language", object_id)
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

func (cix CloudCIXClient) GetMember(object_id string) (MemberSpecific, error) {
	data, err := cix.GetData("membership", "member", object_id)
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

func (cix CloudCIXClient) WriteMember(object_id string, payload map[string]string, method string) (MemberSpecific, error) {
	data, err := cix.WriteData("membership", "member", object_id, payload, method)
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

func (cix CloudCIXClient) GetProfile(object_id string) (ProfileSpecific, error) {
	data, err := cix.GetData("membership", "profile", object_id)
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

func (cix CloudCIXClient) WriteProfile(object_id string, payload map[string]string, method string) (ProfileSpecific, error) {
	data, err := cix.WriteData("membership", "profile", object_id, payload, method)
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

func (cix CloudCIXClient) GetTeam(object_id string) (TeamSpecific, error) {
	data, err := cix.GetData("membership", "team", object_id)
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

func (cix CloudCIXClient) WriteTeam(object_id string, payload map[string]string, method string) (TeamSpecific, error) {
	data, err := cix.WriteData("membership", "team", object_id, payload, method)
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

func (cix CloudCIXClient) GetTerritory(object_id string) (TerritorySpecific, error) {
	data, err := cix.GetData("membership", "territory", object_id)
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

func (cix CloudCIXClient) WriteTerritory(object_id string, payload map[string]string, method string) (TerritorySpecific, error) {
	data, err := cix.WriteData("membership", "territory", object_id, payload, method)
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

func (cix CloudCIXClient) GetTransactionType(object_id string) (TransactionTypeSpecific, error) {
	data, err := cix.GetData("membership", "transaction_type", object_id)
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

func (cix CloudCIXClient) GetUser(object_id string) (UserSpecific, error) {
	data, err := cix.GetData("membership", "user", object_id)
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

func (cix CloudCIXClient) WriteUser(object_id string, payload map[string]string, method string) (UserSpecific, error) {
	data, err := cix.WriteData("membership", "user", object_id, payload, method)
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

func (cix CloudCIXClient) GetAllocation(object_id string) (AllocationSpecific, error) {
	data, err := cix.GetData("iaas", "allocation", "object_id")
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

func (cix CloudCIXClient) WriteAllocation(object_id string, payload map[string]string, method string) (AllocationSpecific, error) {
	data, err := cix.WriteData("iaas", "allocation", object_id, payload, method)
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

func (cix CloudCIXClient) GetAsn(object_id string) (AsnSpecific, error) {
	data, err := cix.GetData("iaas", "asn", "object_id")
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

func (cix CloudCIXClient) WriteAsn(object_id string, payload map[string]string, method string) (AsnSpecific, error) {
	data, err := cix.WriteData("iaas", "asn", object_id, payload, method)
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

func (cix CloudCIXClient) GetCixBlacklist(object_id string) (CixBlacklistSpecific, error) {
	data, err := cix.GetData("iaas", "cix_blacklist", "object_id")
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

func (cix CloudCIXClient) WriteCixBlacklist(object_id string, payload map[string]string, method string) (CixBlacklistSpecific, error) {
	data, err := cix.WriteData("iaas", "cix_blacklist", object_id, payload, method)
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

func (cix CloudCIXClient) GetCixWhitelist(object_id string) (CixWhitelistSpecific, error) {
	data, err := cix.GetData("iaas", "cix_whitelist", "object_id")
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

func (cix CloudCIXClient) WriteCiXWhitelist(object_id string, payload map[string]string, method string) (CixWhitelistSpecific, error) {
	data, err := cix.WriteData("iaas", "cix_whitelist", object_id, payload, method)
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

func (cix CloudCIXClient) WriteCloud(object_id string, payload map[string]string, method string) (Cloud, error) {
	data, err := cix.WriteData("iaas", "cloud", object_id, payload, method)
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

func (cix CloudCIXClient) GetDomain(object_id string) (DomainSpecific, error) {
	data, err := cix.GetData("iaas", "domain", "object_id")
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

func (cix CloudCIXClient) WriteDomain(object_id string, payload map[string]string, method string) (DomainSpecific, error) {
	data, err := cix.WriteData("iaas", "domain", object_id, payload, method)
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

func (cix CloudCIXClient) GetImage(object_id string) (ImageSpecific, error) {
	data, err := cix.GetData("iaas", "image", "object_id")
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

func (cix CloudCIXClient) WriteImage(object_id string, payload map[string]string, method string) (ImageSpecific, error) {
	data, err := cix.WriteData("iaas", "image", object_id, payload, method)
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

func (cix CloudCIXClient) GetInterface(object_id string) (InterfaceSpecific, error) {
	data, err := cix.GetData("iaas", "interface", "object_id")
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

func (cix CloudCIXClient) WriteInterface(object_id string, payload map[string]string, method string) (InterfaceSpecific, error) {
	data, err := cix.WriteData("iaas", "interface", object_id, payload, method)
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

func (cix CloudCIXClient) GetIpAddress(object_id string) (IpAddressSpecific, error) {
	data, err := cix.GetData("iaas", "ip_address", "object_id")
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

func (cix CloudCIXClient) WriteIpAddress(object_id string, payload map[string]string, method string) (IpAddressSpecific, error) {
	data, err := cix.WriteData("iaas", "ip_address", object_id, payload, method)
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

func (cix CloudCIXClient) GetIpmi(object_id string) (IpmiSpecific, error) {
	data, err := cix.GetData("iaas", "ipmi", "object_id")
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

func (cix CloudCIXClient) WriteIpmi(object_id string, payload map[string]string, method string) (IpmiSpecific, error) {
	data, err := cix.WriteData("iaas", "ipmi", object_id, payload, method)
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

func (cix CloudCIXClient) GetPoolIp(object_id string) (PoolIpSpecific, error) {
	data, err := cix.GetData("iaas", "pool_ip", "object_id")
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

func (cix CloudCIXClient) WritePoolIp(object_id string, payload map[string]string, method string) (PoolIpSpecific, error) {
	data, err := cix.WriteData("iaas", "pool_ip", object_id, payload, method)
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

func (cix CloudCIXClient) GetProject(object_id string) (ProjectSpecific, error) {
	data, err := cix.GetData("iaas", "project", "object_id")
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

func (cix CloudCIXClient) WriteProject(object_id string, payload map[string]string, method string) (ProjectSpecific, error) {
	data, err := cix.WriteData("iaas", "project", object_id, payload, method)
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

func (cix CloudCIXClient) GetPtrRecord(object_id string) (PtrRecordSpecific, error) {
	data, err := cix.GetData("iaas", "ptr_record", "object_id")
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

func (cix CloudCIXClient) WritePtrRecord(object_id string, payload map[string]string, method string) (PtrRecordSpecific, error) {
	data, err := cix.WriteData("iaas", "ptr_record", object_id, payload, method)
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

func (cix CloudCIXClient) GetRecord(object_id string) (RecordSpecific, error) {
	data, err := cix.GetData("iaas", "record", "object_id")
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

func (cix CloudCIXClient) WriteRecord(object_id string, payload map[string]string, method string) (RecordSpecific, error) {
	data, err := cix.WriteData("iaas", "record", object_id, payload, method)
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

func (cix CloudCIXClient) GetRouter(object_id string) (RouterSpecific, error) {
	data, err := cix.GetData("iaas", "router", "object_id")
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

func (cix CloudCIXClient) WriteRouter(object_id string, payload map[string]string, method string) (RouterSpecific, error) {
	data, err := cix.WriteData("iaas", "router", object_id, payload, method)
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

func (cix CloudCIXClient) GetServer(object_id string) (ServerSpecific, error) {
	data, err := cix.GetData("iaas", "server", "object_id")
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

func (cix CloudCIXClient) WriteServer(object_id string, payload map[string]string, method string) (ServerSpecific, error) {
	data, err := cix.WriteData("iaas", "server", object_id, payload, method)
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

func (cix CloudCIXClient) GetServerType(object_id string) (ServerTypeSpecific, error) {
	data, err := cix.GetData("iaas", "server_type", "object_id")
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

func (cix CloudCIXClient) GetStorageType(object_id string) (StorageTypeSpecific, error) {
	data, err := cix.GetData("iaas", "storage_type", "object_id")
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

func (cix CloudCIXClient) WriteStorageType(object_id string, payload map[string]string, method string) (StorageTypeSpecific, error) {
	data, err := cix.WriteData("iaas", "storage_type", object_id, payload, method)
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

func (cix CloudCIXClient) GetSubnet(object_id string) (SubnetSpecific, error) {
	data, err := cix.GetData("iaas", "subnet", "object_id")
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

func (cix CloudCIXClient) WriteSubnet(object_id string, payload map[string]string, method string) (SubnetSpecific, error) {
	data, err := cix.WriteData("iaas", "subnet", object_id, payload, method)
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

func (cix CloudCIXClient) GetVirtualRouter(object_id string) (VirtualRouterSpecific, error) {
	data, err := cix.GetData("iaas", "virtual_router", "object_id")
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

func (cix CloudCIXClient) WriteVirtualRouter(object_id string, payload map[string]string, method string) (VirtualRouterSpecific, error) {
	data, err := cix.WriteData("iaas", "virtual_router", object_id, payload, method)
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

func (cix CloudCIXClient) GetVm(object_id string) (VmSpecific, error) {
	data, err := cix.GetData("iaas", "vm", "object_id")
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

func (cix CloudCIXClient) WriteVm(object_id string, payload map[string]string, method string) (VmSpecific, error) {
	data, err := cix.WriteData("iaas", "vm", object_id, payload, method)
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

func (cix CloudCIXClient) GetVpn(object_id string) (VpnSpecific, error) {
	data, err := cix.GetData("iaas", "vpn", "object_id")
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

func (cix CloudCIXClient) WriteVpn(object_id string, payload map[string]string, method string) (VpnSpecific, error) {
	data, err := cix.WriteData("iaas", "vpn", object_id, payload, method)
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
