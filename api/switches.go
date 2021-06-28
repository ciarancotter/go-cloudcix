package api

import (
	"encoding/json"
	"fmt"
)

func ApplicationSwitch(application string, service string, objectId string, data []byte) (map[string]interface{}, error) {
	switch application {
	case "iaas":
		return IaasSwitch(service, objectId, data)
	case "membership":
		return MembershipSwitch(service, objectId, data)
	case "training":
		return TrainingSwitch(service, objectId, data)
	}

	return nil, fmt.Errorf("application not found")
}

func IaasSwitch(service string, objectId string, data []byte) (map[string]interface{}, error) {
	var err error
	switch service {
	case "allocation":
		if objectId == "" {
			allocationInstance := Allocation{}
			err = json.Unmarshal(data, &allocationInstance)
			response := make(map[string]interface{})
			response["allocation"] = allocationInstance.Content
			return response, err
		}

		allocationInstance := AllocationSpecific{}
		err = json.Unmarshal(data, &allocationInstance)
		response := make(map[string]interface{})
		response["allocation"] = allocationInstance.Content
		return response, err

	case "app_settings":
		appSettingsInstance := AppSettingsSpecific{}
		err = json.Unmarshal(data, &appSettingsInstance)
		response := make(map[string]interface{})
		response["app_settings"] = appSettingsInstance.Content
		return response, err

	case "asn":
		if objectId == "" {
			asnInstance := Asn{}
			err = json.Unmarshal(data, &asnInstance)
			response := make(map[string]interface{})
			response["asn"] = asnInstance.Content
			return response, err
		}

		asnInstance := AsnSpecific{}
		err = json.Unmarshal(data, &asnInstance)
		response := make(map[string]interface{})
		response["asn"] = asnInstance.Content
		return response, err

	case "cix_blacklist":
		if objectId == "" {
			cixBlacklistInstance := CixBlacklist{}
			err = json.Unmarshal(data, &cixBlacklistInstance)
			response := make(map[string]interface{})
			response["cix_blacklist"] = cixBlacklistInstance.Content
			return response, err
		}

		cixBlacklistInstance := CixBlacklistSpecific{}
		err = json.Unmarshal(data, &cixBlacklistInstance)
		response := make(map[string]interface{})
		response["cix_blacklist"] = cixBlacklistInstance.Content
		return response, err
	case "cix_whitelist":
		if objectId == "" {
			cixWhitelistInstance := CixWhitelist{}
			err = json.Unmarshal(data, &cixWhitelistInstance)
			response := make(map[string]interface{})
			response["cix_whitelist"] = cixWhitelistInstance.Content
			return response, err
		}

		cixWhitelistInstance := CixWhitelistSpecific{}
		err = json.Unmarshal(data, &cixWhitelistInstance)
		response := make(map[string]interface{})
		response["cix_whitelist"] = cixWhitelistInstance.Content
		return response, err
	case "cloud":
		cloudInstance := Cloud{}
		err = json.Unmarshal(data, &cloudInstance)
		response := make(map[string]interface{})
		response["cloud"] = cloudInstance.Content
		return response, err
	case "domain":
		if objectId == "" {
			domainInstance := Domain{}
			err = json.Unmarshal(data, &domainInstance)
			response := make(map[string]interface{})
			response["domain"] = domainInstance.Content
			return response, err
		}

		domainInstance := DomainSpecific{}
		err = json.Unmarshal(data, &domainInstance)
		response := make(map[string]interface{})
		response["domain"] = domainInstance.Content
		return response, err
	case "image":
		if objectId == "" {
			imageInstance := Image{}
			err = json.Unmarshal(data, &imageInstance)
			response := make(map[string]interface{})
			response["image"] = imageInstance.Content
			return response, err
		}

		imageInstance := ImageSpecific{}
		err = json.Unmarshal(data, &imageInstance)
		response := make(map[string]interface{})
		response["image"] = imageInstance.Content
		return response, err
	case "interface":
		if objectId == "" {
			interfaceInstance := Interface{}
			err = json.Unmarshal(data, &interfaceInstance)
			response := make(map[string]interface{})
			response["interface"] = interfaceInstance.Content
			return response, err
		}

		interfaceInstance := InterfaceSpecific{}
		err = json.Unmarshal(data, &interfaceInstance)
		response := make(map[string]interface{})
		response["interface"] = interfaceInstance.Content
		return response, err
	case "ip_address":
		if objectId == "" {
			ipAddressInstance := IpAddress{}
			err = json.Unmarshal(data, &ipAddressInstance)
			response := make(map[string]interface{})
			response["ip_address"] = ipAddressInstance.Content
			return response, err
		}

		ipAddressInstance := IpAddressSpecific{}
		err = json.Unmarshal(data, &ipAddressInstance)
		response := make(map[string]interface{})
		response["ip_address"] = ipAddressInstance.Content
		return response, err

	case "ipmi":
		if objectId == "" {
			ipmiInstance := Ipmi{}
			err = json.Unmarshal(data, &ipmiInstance)
			response := make(map[string]interface{})
			response["ipmi"] = ipmiInstance.Content
			return response, err
		}

		ipmiInstance := IpmiSpecific{}
		err = json.Unmarshal(data, &ipmiInstance)
		response := make(map[string]interface{})
		response["ipmi"] = ipmiInstance.Content
		return response, err
	case "policy_log":
		if objectId == "" {
			return nil, fmt.Errorf("Invalid URL for policy log.")
		}

		policy_logInstance := PolicyLog{}
		err = json.Unmarshal(data, &policy_logInstance)
		response := make(map[string]interface{})
		response["policy_log"] = policy_logInstance.Content
		return response, err
	case "pool_ip":
		if objectId == "" {
			poolIpInstance := PoolIp{}
			err = json.Unmarshal(data, &poolIpInstance)
			response := make(map[string]interface{})
			response["pool_ip"] = poolIpInstance.Content
			return response, err
		}

		poolIpInstance := PoolIpSpecific{}
		err = json.Unmarshal(data, &poolIpInstance)
		response := make(map[string]interface{})
		response["pool_ip"] = poolIpInstance.Content
		return response, err
	case "project":
		if objectId == "" {
			projectInstance := Project{}
			err = json.Unmarshal(data, &projectInstance)
			response := make(map[string]interface{})
			response["project"] = projectInstance.Content
			return response, err
		}

		projectInstance := ProjectSpecific{}
		err = json.Unmarshal(data, &projectInstance)
		response := make(map[string]interface{})
		response["project"] = projectInstance.Content
		return response, err
	case "ptr_record":
		if objectId == "" {
			ptrRecordInstance := PtrRecord{}
			err = json.Unmarshal(data, &ptrRecordInstance)
			response := make(map[string]interface{})
			response["ptr_record"] = ptrRecordInstance.Content
			return response, err
		}

		ptrRecordInstance := PtrRecordSpecific{}
		err = json.Unmarshal(data, &ptrRecordInstance)
		response := make(map[string]interface{})
		response["ptr_record"] = ptrRecordInstance.Content
		return response, err
	case "record":
		if objectId == "" {
			recordInstance := Record{}
			err = json.Unmarshal(data, &recordInstance)
			response := make(map[string]interface{})
			response["record"] = recordInstance.Content
			return response, err
		}

		recordInstance := RecordSpecific{}
		err = json.Unmarshal(data, &recordInstance)
		response := make(map[string]interface{})
		response["record"] = recordInstance.Content
		return response, err
	case "router":
		if objectId == "" {
			routerInstance := Router{}
			err = json.Unmarshal(data, &routerInstance)
			response := make(map[string]interface{})
			response["router"] = routerInstance.Content
			return response, err
		}

		routerInstance := RouterSpecific{}
		err = json.Unmarshal(data, &routerInstance)
		response := make(map[string]interface{})
		response["router"] = routerInstance.Content
		return response, err
	case "server":
		if objectId == "" {
			serverInstance := Server{}
			err = json.Unmarshal(data, &serverInstance)
			response := make(map[string]interface{})
			response["server"] = serverInstance.Content
			return response, err
		}

		serverInstance := ServerSpecific{}
		err = json.Unmarshal(data, &serverInstance)
		response := make(map[string]interface{})
		response["server"] = serverInstance.Content
		return response, err
	case "server_type":
		if objectId == "" {
			serverTypeInstance := ServerType{}
			err = json.Unmarshal(data, &serverTypeInstance)
			response := make(map[string]interface{})
			response["server_type"] = serverTypeInstance.Content
			return response, err
		}

		serverTypeInstance := ServerTypeSpecific{}
		err = json.Unmarshal(data, &serverTypeInstance)
		response := make(map[string]interface{})
		response["server_type"] = serverTypeInstance.Content
		return response, err
	case "storage_type":
		if objectId == "" {
			storageTypeInstance := StorageType{}
			err = json.Unmarshal(data, &storageTypeInstance)
			response := make(map[string]interface{})
			response["storage_type"] = storageTypeInstance.Content
			return response, err
		}

		storageTypeInstance := StorageTypeSpecific{}
		err = json.Unmarshal(data, &storageTypeInstance)
		response := make(map[string]interface{})
		response["storage_type"] = storageTypeInstance.Content
		return response, err
	case "subnet":
		if objectId == "" {
			subnetInstance := Subnet{}
			err = json.Unmarshal(data, &subnetInstance)
			response := make(map[string]interface{})
			response["subnet"] = subnetInstance.Content
			return response, err
		}

		subnetInstance := SubnetSpecific{}
		err = json.Unmarshal(data, &subnetInstance)
		response := make(map[string]interface{})
		response["subnet"] = subnetInstance.Content
		return response, err
	case "subnet_space":
		if objectId == "" {
			return nil, fmt.Errorf("Invalid URL for subnet space.")
		}

		subnetSpaceInstance := SubnetSpace{}
		err = json.Unmarshal(data, &subnetSpaceInstance)
		response := make(map[string]interface{})
		response["subnet_space"] = subnetSpaceInstance.Content
		return response, err
	case "virtual_router":
		if objectId == "" {
			virtualRouterInstance := VirtualRouter{}
			err = json.Unmarshal(data, &virtualRouterInstance)
			response := make(map[string]interface{})
			response["virtual_router"] = virtualRouterInstance.Content
			return response, err
		}

		virtualRouterInstance := VirtualRouterSpecific{}
		err = json.Unmarshal(data, &virtualRouterInstance)
		response := make(map[string]interface{})
		response["virtual_router"] = virtualRouterInstance.Content
		return response, err
	case "vm":
		if objectId == "" {
			vmInstance := Vm{}
			err = json.Unmarshal(data, &vmInstance)
			response := make(map[string]interface{})
			response["vm"] = vmInstance.Content
			return response, err
		}

		vmInstance := VmSpecific{}
		err = json.Unmarshal(data, &vmInstance)
		response := make(map[string]interface{})
		response["vm"] = vmInstance.Content
		return response, err
	case "vpn":
		if objectId == "" {
			vpnInstance := Vpn{}
			err = json.Unmarshal(data, &vpnInstance)
			response := make(map[string]interface{})
			response["vpn"] = vpnInstance.Content
			return response, err
		}

		vpnInstance := VpnSpecific{}
		err = json.Unmarshal(data, &vpnInstance)
		response := make(map[string]interface{})
		response["vpn"] = vpnInstance.Content
		return response, err
	}

	return nil, fmt.Errorf("IAAS service not found")
}

func MembershipSwitch(service string, objectId string, data []byte) (map[string]interface{}, error) {
	var err error

	switch service {
	case "address":
		addressInstance := Address{}
		err = json.Unmarshal(data, &addressInstance)
		response := make(map[string]interface{})
		response["address"] = addressInstance.Content
		return response, err
	case "app_settings":
		appSettingsInstance := AppSettings{}
		err = json.Unmarshal(data, &appSettingsInstance)
		response := make(map[string]interface{})
		response["app_setting"] = appSettingsInstance.Content
		return response, err
	case "country":
		if objectId == "" {
			countryInstance := Country{}
			err = json.Unmarshal(data, &countryInstance)
			response := make(map[string]interface{})
			response["country"] = countryInstance.Content
			return response, err
		}

		countryInstance := CountrySpecific{}
		err = json.Unmarshal(data, &countryInstance)
		response := make(map[string]interface{})
		response["country"] = countryInstance.Content
		return response, err
	case "currency":
		if objectId == "" {
			currencyInstance := Currency{}
			err = json.Unmarshal(data, &currencyInstance)
			response := make(map[string]interface{})
			response["currency"] = currencyInstance.Content
			return response, err
		}

		currencyInstance := CurrencySpecific{}
		err = json.Unmarshal(data, &currencyInstance)
		response := make(map[string]interface{})
		response["currency"] = currencyInstance.Content
		return response, err
	case "department":
		if objectId == "" {
			departmentInstance := Department{}
			err = json.Unmarshal(data, &departmentInstance)
			response := make(map[string]interface{})
			response["department"] = departmentInstance.Content
			return response, err
		}

		departmentInstance := DepartmentSpecific{}
		err = json.Unmarshal(data, &departmentInstance)
		response := make(map[string]interface{})
		response["department"] = departmentInstance.Content
		return response, err
	case "email_confirmation":
		emailInstance := EmailConfirmed{}
		err = json.Unmarshal(data, &emailInstance)
		response := make(map[string]interface{})
		response["email_confirmation"] = emailInstance.Content
		return response, err
	case "language":
		if objectId == "" {
			languageInstance := Language{}
			err = json.Unmarshal(data, &languageInstance)
			response := make(map[string]interface{})
			response["language"] = languageInstance.Content
			return response, err
		}

		languageInstance := LanguageSpecific{}
		err = json.Unmarshal(data, &languageInstance)
		response := make(map[string]interface{})
		response["language"] = languageInstance.Content
		return response, err
	case "member":
		if objectId == "" {
			memberInstance := Member{}
			err = json.Unmarshal(data, &memberInstance)
			response := make(map[string]interface{})
			response["member"] = memberInstance.Content
			return response, err
		}

		memberInstance := MemberSpecific{}
		err = json.Unmarshal(data, &memberInstance)
		response := make(map[string]interface{})
		response["member"] = memberInstance.Content
		return response, err
	case "profile":
		if objectId == "" {
			profileInstance := Profile{}
			err = json.Unmarshal(data, &profileInstance)
			response := make(map[string]interface{})
			response["profile"] = profileInstance.Content
			return response, err
		}

		profileInstance := ProfileSpecific{}
		err = json.Unmarshal(data, &profileInstance)
		response := make(map[string]interface{})
		response["profile"] = profileInstance.Content
		return response, err
	case "team":
		if objectId == "" {
			teamInstance := Team{}
			err = json.Unmarshal(data, &teamInstance)
			response := make(map[string]interface{})
			response["team"] = teamInstance.Content
			return response, err
		}

		teamInstance := TeamSpecific{}
		err = json.Unmarshal(data, &teamInstance)
		response := make(map[string]interface{})
		response["team"] = teamInstance.Content
		return response, err
	case "territory":
		if objectId == "" {
			territoryInstance := Territory{}
			err = json.Unmarshal(data, &territoryInstance)
			response := make(map[string]interface{})
			response["territory"] = territoryInstance.Content
			return response, err
		}

		territoryInstance := TerritorySpecific{}
		err = json.Unmarshal(data, &territoryInstance)
		response := make(map[string]interface{})
		response["territory"] = territoryInstance.Content
		return response, err
	case "transaction_type":
		if objectId == "" {
			transactionTypeInstance := TransactionType{}
			err = json.Unmarshal(data, &transactionTypeInstance)
			response := make(map[string]interface{})
			response["transaction_type"] = transactionTypeInstance.Content
			return response, err
		} else {
			transactionTypeInstance := TranscationTypeSpecific{}
			err = json.Unmarshal(data, &transactionTypeInstance)
			response := make(map[string]interface{})
			response["transaction_type"] = transactionTypeInstance.Content
			return response, err
		}

	case "user":
		if objectId == "" {
			userInstance := User{}
			err = json.Unmarshal(data, &userInstance)
			response := make(map[string]interface{})
			response["user"] = userInstance.Content
			return response, err
		}

		userInstance := UserSpecific{}
		err = json.Unmarshal(data, &userInstance)
		response := make(map[string]interface{})
		response["user"] = userInstance.Content
		return response, err
	}

	return nil, fmt.Errorf("Membership service not found")
}

func TrainingSwitch(service string, objectId string, data []byte) (map[string]interface{}, error) {
	var err error

	switch service {
	case "class":
		if objectId == "" {
			classInstance := Class{}
			err = json.Unmarshal(data, &classInstance)
			response := make(map[string]interface{})
			response["class"] = classInstance.Content
			return response, err
		}

		classInstance := ClassSpecific{}
		err = json.Unmarshal(data, &classInstance)
		response := make(map[string]interface{})
		response["class"] = classInstance.Content
		return response, err
	case "student":
		if objectId == "" {
			studentInstance := Student{}
			err = json.Unmarshal(data, &studentInstance)
			response := make(map[string]interface{})
			response["student"] = studentInstance.Content
			return response, err
		}

		studentInstance := StudentSpecific{}
		err = json.Unmarshal(data, &studentInstance)
		response := make(map[string]interface{})
		response["student"] = studentInstance.Content
		return response, err
	case "syllabus":
		response := make(map[string]interface{})
		if objectId == "" {
			syllabusInstance := Syllabus{}
			err = json.Unmarshal(data, &syllabusInstance)
			return response, err
		}

		syllabusInstance := SyllabusSpecific{}
		err = json.Unmarshal(data, &syllabusInstance)

		response["id"] = syllabusInstance.Content.ID
		response["name"] = syllabusInstance.Content.Name
		response["description"] = syllabusInstance.Content.Description
		response["member_id"] = syllabusInstance.Content.MemberID
		response["uri"] = syllabusInstance.Content.URI

		return response, err
	}

	return nil, fmt.Errorf("Training service not found")
}
