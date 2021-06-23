package api

import (
	"encoding/json"
	"fmt"
)

func ApplicationSwitch(application string, service string, objectId string, data []byte) (string, error) {
	switch application {
	case "iaas":
		return IaasSwitch(service, objectId, data)
	case "membership":
		return MembershipSwitch(service, objectId, data)
	case "training":
		return TrainingSwitch(service, objectId, data)
	}

	return "", fmt.Errorf("application not found")
}

func IaasSwitch(service string, objectId string, data []byte) (string, error) {
	var err error
	switch service {
	case "allocation":
		if objectId == "" {
			allocationInstance := Allocation{}
			err = json.Unmarshal(data, &allocationInstance)
			return fmt.Sprint(allocationInstance.Content[0]), err
		}

		allocationInstance := AllocationSpecific{}
		err = json.Unmarshal(data, &allocationInstance)
		return fmt.Sprint(allocationInstance.Content), err

	case "app_settings":
		appSettingsInstance := AppSettingsSpecific{}
		err = json.Unmarshal(data, &appSettingsInstance)
		return fmt.Sprint(appSettingsInstance.Content), err

	case "asn":
		if objectId == "" {
			asnInstance := Asn{}
			err = json.Unmarshal(data, &asnInstance)
			return fmt.Sprint(asnInstance.Content[0]), err
		}

		asnInstance := AsnSpecific{}
		err = json.Unmarshal(data, &asnInstance)
		return fmt.Sprint(asnInstance.Content), err

	case "cix_blacklist":
		if objectId == "" {
			cixBlacklistInstance := CixBlacklist{}
			err = json.Unmarshal(data, &cixBlacklistInstance)
			return fmt.Sprint(cixBlacklistInstance.Content[0]), err
		}

		cixBlacklistInstance := CixBlacklistSpecific{}
		err = json.Unmarshal(data, &cixBlacklistInstance)
		return fmt.Sprint(cixBlacklistInstance.Content), err

	case "cix_whitelist":
		if objectId == "" {
			cixWhitelistInstance := CixWhitelist{}
			err = json.Unmarshal(data, &cixWhitelistInstance)
			return fmt.Sprint(cixWhitelistInstance.Content[0]), err
		}

		cixWhitelistInstance := CixWhitelistSpecific{}
		err = json.Unmarshal(data, &cixWhitelistInstance)
		return fmt.Sprint(cixWhitelistInstance.Content), err

	case "cloud":
		cloudInstance := Cloud{}
		err = json.Unmarshal(data, &cloudInstance)
		return fmt.Sprint(cloudInstance.Content), err

	case "domain":
		if objectId == "" {
			domainInstance := Domain{}
			err = json.Unmarshal(data, &domainInstance)
			return fmt.Sprint(domainInstance.Content[0]), err
		}

		domainInstance := DomainSpecific{}
		err = json.Unmarshal(data, &domainInstance)
		return fmt.Sprint(domainInstance.Content), err

	case "image":
		if objectId == "" {
			imageInstance := Image{}
			err = json.Unmarshal(data, &imageInstance)
			return fmt.Sprint(imageInstance.Content[0]), err
		}

		imageInstance := ImageSpecific{}
		err = json.Unmarshal(data, &imageInstance)
		return fmt.Sprint(imageInstance.Content), err

	case "interface":
		if objectId == "" {
			interfaceInstance := Interface{}
			err = json.Unmarshal(data, &interfaceInstance)
			return fmt.Sprint(interfaceInstance.Content[0]), err
		}

		interfaceInstance := InterfaceSpecific{}
		err = json.Unmarshal(data, &interfaceInstance)
		return fmt.Sprint(interfaceInstance.Content), err

	case "ip_address":
		if objectId == "" {
			ipAddressInstance := IpAddress{}
			err = json.Unmarshal(data, &ipAddressInstance)
			return fmt.Sprint(ipAddressInstance.Content[0]), err
		}

		ipAddressInstance := IpAddressSpecific{}
		err = json.Unmarshal(data, &ipAddressInstance)
		return fmt.Sprint(ipAddressInstance.Content), err

	case "ipmi":
		if objectId == "" {
			ipmiInstance := Ipmi{}
			err = json.Unmarshal(data, &ipmiInstance)
			return fmt.Sprint(ipmiInstance.Content[0]), err
		}

		ipmiInstance := IpmiSpecific{}
		err = json.Unmarshal(data, &ipmiInstance)
		return fmt.Sprint(ipmiInstance.Content), err

	case "policy_log":
		if objectId == "" {
			return "", fmt.Errorf("Invalid URL for policy log.")
		}

		policy_logInstance := PolicyLog{}
		err = json.Unmarshal(data, &policy_logInstance)
		return fmt.Sprint(policy_logInstance.Content), err

	case "pool_ip":
		if objectId == "" {
			poolIpInstance := PoolIp{}
			err = json.Unmarshal(data, &poolIpInstance)
			return fmt.Sprint(poolIpInstance.Content[0]), err
		}

		poolIpInstance := PoolIpSpecific{}
		err = json.Unmarshal(data, &poolIpInstance)
		return fmt.Sprint(poolIpInstance.Content), err

	case "project":
		if objectId == "" {
			projectInstance := Project{}
			err = json.Unmarshal(data, &projectInstance)
			return fmt.Sprint(projectInstance.Content[0]), err
		}

		projectInstance := ProjectSpecific{}
		err = json.Unmarshal(data, &projectInstance)
		return fmt.Sprint(projectInstance.Content), err

	case "ptr_record":
		if objectId == "" {
			ptrRecordInstance := PtrRecord{}
			err = json.Unmarshal(data, &ptrRecordInstance)
			return fmt.Sprint(ptrRecordInstance.Content[0]), err
		}

		ptrRecordInstance := PtrRecordSpecific{}
		err = json.Unmarshal(data, &ptrRecordInstance)
		return fmt.Sprint(ptrRecordInstance.Content), err

	case "record":
		if objectId == "" {
			recordInstance := Record{}
			err = json.Unmarshal(data, &recordInstance)
			return fmt.Sprint(recordInstance.Content[0]), err
		}

		recordInstance := RecordSpecific{}
		err = json.Unmarshal(data, &recordInstance)
		return fmt.Sprint(recordInstance.Content), err

	case "router":
		if objectId == "" {
			routerInstance := Router{}
			err = json.Unmarshal(data, &routerInstance)
			return fmt.Sprint(routerInstance.Content[0]), err
		}

		routerInstance := RouterSpecific{}
		err = json.Unmarshal(data, &routerInstance)
		return fmt.Sprint(routerInstance.Content), err

	case "server":
		if objectId == "" {
			serverInstance := Server{}
			err = json.Unmarshal(data, &serverInstance)
			return fmt.Sprint(serverInstance.Content[0]), err
		}

		serverInstance := ServerSpecific{}
		err = json.Unmarshal(data, &serverInstance)
		return fmt.Sprint(serverInstance.Content), err

	case "server_type":
		if objectId == "" {
			serverTypeInstance := ServerType{}
			err = json.Unmarshal(data, &serverTypeInstance)
			return fmt.Sprint(serverTypeInstance.Content[0]), err
		}

		serverTypeInstance := ServerTypeSpecific{}
		err = json.Unmarshal(data, &serverTypeInstance)
		return fmt.Sprint(serverTypeInstance.Content), err

	case "storage_type":
		if objectId == "" {
			storageTypeInstance := StorageType{}
			err = json.Unmarshal(data, &storageTypeInstance)
			return fmt.Sprint(storageTypeInstance.Content[0]), err
		}

		storageTypeInstance := StorageTypeSpecific{}
		err = json.Unmarshal(data, &storageTypeInstance)
		return fmt.Sprint(storageTypeInstance.Content), err

	case "subnet":
		if objectId == "" {
			subnetInstance := Subnet{}
			err = json.Unmarshal(data, &subnetInstance)
			return fmt.Sprint(subnetInstance.Content[0]), err
		}

		subnetInstance := SubnetSpecific{}
		err = json.Unmarshal(data, &subnetInstance)
		return fmt.Sprint(subnetInstance.Content), err

	case "subnet_space":
		if objectId == "" {
			return "", fmt.Errorf("Invalid URL for subnet space.")
		}

		subnetSpaceInstance := SubnetSpace{}
		err = json.Unmarshal(data, &subnetSpaceInstance)
		return fmt.Sprint(subnetSpaceInstance.Content), err

	case "virtual_router":
		if objectId == "" {
			virtualRouterInstance := VirtualRouter{}
			err = json.Unmarshal(data, &virtualRouterInstance)
			return fmt.Sprint(virtualRouterInstance.Content[0]), err
		}

		virtualRouterInstance := VirtualRouterSpecific{}
		err = json.Unmarshal(data, &virtualRouterInstance)
		return fmt.Sprint(virtualRouterInstance.Content), err

	case "vm":
		if objectId == "" {
			vmInstance := Vm{}
			err = json.Unmarshal(data, &vmInstance)
			return fmt.Sprint(vmInstance.Content[0]), err
		}

		vmInstance := VmSpecific{}
		err = json.Unmarshal(data, &vmInstance)
		return fmt.Sprint(vmInstance.Content), err

	case "vpn":
		if objectId == "" {
			vpnInstance := Vpn{}
			err = json.Unmarshal(data, &vpnInstance)
			return fmt.Sprint(vpnInstance.Content[0]), err
		}

		vpnInstance := VpnSpecific{}
		err = json.Unmarshal(data, &vpnInstance)
		return fmt.Sprint(vpnInstance.Content), err
	}

	return "", fmt.Errorf("IAAS service not found")
}

func MembershipSwitch(service string, objectId string, data []byte) (string, error) {
	var err error

	switch service {
	case "address":
		addressInstance := Address{}
		err = json.Unmarshal(data, &addressInstance)
		return fmt.Sprint(addressInstance.Content[0]), err

	case "app_settings":
		appSettingsInstance := AppSettings{}
		err = json.Unmarshal(data, &appSettingsInstance)
		return fmt.Sprint(appSettingsInstance.Content), err

	case "country":
		if objectId == "" {
			countryInstance := Country{}
			err = json.Unmarshal(data, &countryInstance)
			return fmt.Sprint(countryInstance.Content[0]), err
		}

		countryInstance := CountrySpecific{}
		err = json.Unmarshal(data, &countryInstance)
		return fmt.Sprint(countryInstance.Content), err

	case "currency":
		if objectId == "" {
			currencyInstance := Currency{}
			err = json.Unmarshal(data, &currencyInstance)
			return fmt.Sprint(currencyInstance.Content[0]), err
		}

		currencyInstance := CurrencySpecific{}
		err = json.Unmarshal(data, &currencyInstance)
		return fmt.Sprint(currencyInstance.Content), err

	case "department":
		if objectId == "" {
			departmentInstance := Department{}
			err = json.Unmarshal(data, &departmentInstance)
			return fmt.Sprint(departmentInstance.Content[0]), err
		}

		departmentInstance := DepartmentSpecific{}
		err = json.Unmarshal(data, &departmentInstance)
		return fmt.Sprint(departmentInstance.Content), err

	case "email_confirmation":
		emailInstance := EmailConfirmed{}
		err = json.Unmarshal(data, &emailInstance)
		return fmt.Sprint(emailInstance.Content), err

	case "language":
		if objectId == "" {
			languageInstance := Language{}
			err = json.Unmarshal(data, &languageInstance)
			return fmt.Sprint(languageInstance.Content[0]), err
		}

		languageInstance := LanguageSpecific{}
		err = json.Unmarshal(data, &languageInstance)
		return fmt.Sprint(languageInstance.Content), err

	case "member":
		if objectId == "" {
			memberInstance := Member{}
			err = json.Unmarshal(data, &memberInstance)
			return fmt.Sprint(memberInstance.Content[0]), err
		}

		memberInstance := MemberSpecific{}
		err = json.Unmarshal(data, &memberInstance)
		return fmt.Sprint(memberInstance.Content), err

	case "profile":
		if objectId == "" {
			profileInstance := Profile{}
			err = json.Unmarshal(data, &profileInstance)
			return fmt.Sprint(profileInstance.Content[0]), err
		}

		profileInstance := ProfileSpecific{}
		err = json.Unmarshal(data, &profileInstance)
		return fmt.Sprint(profileInstance.Content), err

	case "team":
		if objectId == "" {
			teamInstance := Team{}
			err = json.Unmarshal(data, &teamInstance)
			return fmt.Sprint(teamInstance.Content[0]), err
		}

		teamInstance := TeamSpecific{}
		err = json.Unmarshal(data, &teamInstance)
		return fmt.Sprint(teamInstance.Content), err

	case "territory":
		if objectId == "" {
			territoryInstance := Territory{}
			err = json.Unmarshal(data, &territoryInstance)
			return fmt.Sprint(territoryInstance.Content[0]), err
		}

		territoryInstance := TerritorySpecific{}
		err = json.Unmarshal(data, &territoryInstance)
		return fmt.Sprint(territoryInstance.Content), err

	case "transaction_type":
		if objectId == "" {
			transactionTypeInstance := TransactionType{}
			err = json.Unmarshal(data, &transactionTypeInstance)
			return fmt.Sprint(transactionTypeInstance.Content[0]), err
		} else {
			transactionTypeInstance := TranscationTypeSpecific{}
			err = json.Unmarshal(data, &transactionTypeInstance)
			return fmt.Sprint(transactionTypeInstance.Content), err
		}

	case "user":
		if objectId == "" {
			userInstance := User{}
			err = json.Unmarshal(data, &userInstance)
			return fmt.Sprint(userInstance.Content[0]), err
		}

		userInstance := UserSpecific{}
		err = json.Unmarshal(data, &userInstance)
		return fmt.Sprint(userInstance.Content), err
	}

	return "", fmt.Errorf("Membership service not found")
}

func TrainingSwitch(service string, objectId string, data []byte) (string, error) {
	var err error

	switch service {
	case "class":
		if objectId == "" {
			classInstance := Class{}
			err = json.Unmarshal(data, &classInstance)
			return fmt.Sprint(classInstance.Content[0]), err
		}

		classInstance := ClassSpecific{}
		err = json.Unmarshal(data, &classInstance)
		return fmt.Sprint(classInstance.Content), err

	case "student":
		if objectId == "" {
			studentInstance := Student{}
			err = json.Unmarshal(data, &studentInstance)
			return fmt.Sprint(studentInstance.Content[0]), err
		}

		studentInstance := StudentSpecific{}
		err = json.Unmarshal(data, &studentInstance)
		return fmt.Sprint(studentInstance.Content), err

	case "syllabus":
		if objectId == "" {
			syllabusInstance := Syllabus{}
			err = json.Unmarshal(data, &syllabusInstance)
			return fmt.Sprint(syllabusInstance.Content[0]), err
		}

		syllabusInstance := SyllabusSpecific{}
		err = json.Unmarshal(data, &syllabusInstance)
		return fmt.Sprint(syllabusInstance.Content), err
	}

	return "", fmt.Errorf("Training service not found")
}
