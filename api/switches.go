package api

import (
	"encoding/json"
	"fmt"
)

func Application_switch(application string, service string, object_id string, data []byte, err error) {
	switch application {
	case "iaas":
		IAAS_switch(service, object_id, data, err)

	case "membership":
		Membership_switch(service, object_id, data, err)

	case "training":
		Training_switch(service, object_id, data, err)
	}
}

func IAAS_switch(service string, object_id string, data []byte, err error) {

	switch service {

	case "allocation":
		if object_id == "" {
			allocation_instance := ALLOCATION{}
			err = json.Unmarshal(data, &allocation_instance)
			fmt.Println(allocation_instance.Content[0])
		} else {
			allocation_instance := ALLOCATION_SPECIFIC{}
			err = json.Unmarshal(data, &allocation_instance)
			fmt.Println(allocation_instance.Content)
		}

	case "app_settings":
		app_settings_instance := APP_SETTINGS_SPECIFIC{}
		err = json.Unmarshal(data, &app_settings_instance)
		fmt.Println(app_settings_instance.Content)

	case "asn":
		if object_id == "" {
			asn_instance := ASN{}
			err = json.Unmarshal(data, &asn_instance)
			fmt.Println(asn_instance.Content[0])
		} else {
			asn_instance := ASN_SPECIFIC{}
			err = json.Unmarshal(data, &asn_instance)
			fmt.Println(asn_instance.Content)
		}

	case "cix_blacklist":
		if object_id == "" {
			cix_blacklist_instance := CIX_BLACKLIST{}
			err = json.Unmarshal(data, &cix_blacklist_instance)
			fmt.Println(cix_blacklist_instance.Content[0])
		} else {
			cix_blacklist_instance := CIX_BLACKLIST{}
			err = json.Unmarshal(data, &cix_blacklist_instance)
			fmt.Println(cix_blacklist_instance.Content)
		}

	case "cix_whitelist":
		if object_id == "" {
			cix_whitelist_instance := CIX_WHITELIST{}
			err = json.Unmarshal(data, &cix_whitelist_instance)
			fmt.Println(cix_whitelist_instance.Content[0])
		} else {
			cix_whitelist_instance := CIX_WHITELIST_SPECIFIC{}
			err = json.Unmarshal(data, &cix_whitelist_instance)
			fmt.Println(cix_whitelist_instance.Content)
		}

	case "cloud":
		cloud_instance := CLOUD{}
		err = json.Unmarshal(data, &cloud_instance)
		fmt.Println(cloud_instance.Content)

	case "domain":
		if object_id == "" {
			domain_instance := DOMAIN{}
			err = json.Unmarshal(data, &domain_instance)
			fmt.Println(domain_instance.Content[0])
		} else {
			domain_instance := DOMAIN_SPECIFIC{}
			err = json.Unmarshal(data, &domain_instance)
			fmt.Println(domain_instance.Content)
		}

	case "image":
		if object_id == "" {
			image_instance := IMAGE{}
			err = json.Unmarshal(data, &image_instance)
			fmt.Println(image_instance.Content[0])
		} else {
			image_instance := IMAGE_SPECIFIC{}
			err = json.Unmarshal(data, &image_instance)
			fmt.Println(image_instance.Content)
		}

	case "interface":
		if object_id == "" {
			interface_instance := INTERFACE{}
			err = json.Unmarshal(data, &interface_instance)
			fmt.Println(interface_instance.Content[0])
		} else {
			interface_instance := INTERFACE_SPECIFIC{}
			err = json.Unmarshal(data, &interface_instance)
			fmt.Println(interface_instance.Content)
		}

	case "ip_address":
		if object_id == "" {
			ip_address_instance := IP_ADDRESS{}
			err = json.Unmarshal(data, &ip_address_instance)
			fmt.Println(ip_address_instance.Content[0])
		} else {
			ip_address_instance := IP_ADDRESS_SPECIFIC{}
			err = json.Unmarshal(data, &ip_address_instance)
			fmt.Println(ip_address_instance.Content)
		}

	case "ipmi":
		if object_id == "" {
			ipmi_instance := IPMI{}
			err = json.Unmarshal(data, &ipmi_instance)
			fmt.Println(ipmi_instance.Content[0])
		} else {
			ipmi_instance := IPMI_SPECIFIC{}
			err = json.Unmarshal(data, &ipmi_instance)
			fmt.Println(ipmi_instance.Content)
		}

	case "policy_log":
		if object_id == "" {
			fmt.Println("Invalid URL for policy log.")
		} else {
			policy_log_instance := POLICY_LOG{}
			err = json.Unmarshal(data, &policy_log_instance)
			fmt.Println(policy_log_instance.Content)
		}

	case "pool_ip":
		if object_id == "" {
			pool_ip_instance := POOL_IP{}
			err = json.Unmarshal(data, &pool_ip_instance)
			fmt.Println(pool_ip_instance.Content[0])
		} else {
			pool_ip_instance := POOL_IP_SPECIFIC{}
			err = json.Unmarshal(data, &pool_ip_instance)
			fmt.Println(pool_ip_instance.Content)
		}

	case "project":
		if object_id == "" {
			project_instance := PROJECT{}
			err = json.Unmarshal(data, &project_instance)
			fmt.Println(project_instance.Content[0])
		} else {
			project_instance := PROJECT_SPECIFIC{}
			err = json.Unmarshal(data, &project_instance)
			fmt.Println(project_instance.Content)
		}

	case "ptr_record":
		if object_id == "" {
			ptr_record_instance := PTR_RECORD{}
			err = json.Unmarshal(data, &ptr_record_instance)
			fmt.Println(ptr_record_instance.Content[0])
		} else {
			ptr_record_instance := PTR_RECORD_SPECIFIC{}
			err = json.Unmarshal(data, &ptr_record_instance)
			fmt.Println(ptr_record_instance.Content)
		}

	case "record":
		if object_id == "" {
			record_instance := RECORD{}
			err = json.Unmarshal(data, &record_instance)
			fmt.Println(record_instance.Content[0])
		} else {
			record_instance := RECORD_SPECIFIC{}
			err = json.Unmarshal(data, &record_instance)
			fmt.Println(record_instance.Content)
		}

	case "router":
		if object_id == "" {
			router_instance := ROUTER{}
			err = json.Unmarshal(data, &router_instance)
			fmt.Println(router_instance.Content[0])
		} else {
			router_instance := ROUTER_SPECIFIC{}
			err = json.Unmarshal(data, &router_instance)
			fmt.Println(router_instance.Content)
		}

	case "server":
		if object_id == "" {
			server_instance := SERVER{}
			err = json.Unmarshal(data, &server_instance)
			fmt.Println(server_instance.Content[0])
		} else {
			server_instance := SERVER_SPECIFIC{}
			err = json.Unmarshal(data, &server_instance)
			fmt.Println(server_instance.Content)
		}

	case "server_type":
		if object_id == "" {
			server_type_instance := SERVER_TYPE{}
			err = json.Unmarshal(data, &server_type_instance)
			fmt.Println(server_type_instance.Content[0])
		} else {
			server_type_instance := SERVER_TYPE_SPECIFIC{}
			err = json.Unmarshal(data, &server_type_instance)
			fmt.Println(server_type_instance.Content)
		}

	case "storage_type":
		if object_id == "" {
			storage_type_instance := STORAGE_TYPE{}
			err = json.Unmarshal(data, &storage_type_instance)
			fmt.Println(storage_type_instance.Content[0])
		} else {
			storage_type_instance := STORAGE_TYPE_SPECIFIC{}
			err = json.Unmarshal(data, &storage_type_instance)
			fmt.Println(storage_type_instance.Content)
		}

	case "subnet":
		if object_id == "" {
			subnet_instance := SUBNET{}
			err = json.Unmarshal(data, &subnet_instance)
			fmt.Println(subnet_instance.Content[0])
		} else {
			subnet_instance := SUBNET_SPECIFIC{}
			err = json.Unmarshal(data, &subnet_instance)
			fmt.Println(subnet_instance.Content)
		}

	case "subnet_space":
		if object_id == "" {
			fmt.Println("Invalid URL for subnet space.")

		} else {
			subnet_space_instance := SUBNET_SPACE{}
			err = json.Unmarshal(data, &subnet_space_instance)
			fmt.Println(subnet_space_instance.Content)
		}

	case "virtual_router":
		if object_id == "" {
			virtual_router_instance := VIRTUAL_ROUTER{}
			err = json.Unmarshal(data, &virtual_router_instance)
			fmt.Println(virtual_router_instance.Content[0])
		} else {
			virtual_router_instance := VIRTUAL_ROUTER_SPECIFIC{}
			err = json.Unmarshal(data, &virtual_router_instance)
			fmt.Println(virtual_router_instance.Content)
		}

	case "vm":
		if object_id == "" {
			vm_instance := VM{}
			err = json.Unmarshal(data, &vm_instance)
			fmt.Println(vm_instance.Content[0])
		} else {
			vm_instance := VM_SPECIFIC{}
			err = json.Unmarshal(data, &vm_instance)
			fmt.Println(vm_instance.Content)
		}

	case "vpn":
		if object_id == "" {
			vpn_instance := VPN{}
			err = json.Unmarshal(data, &vpn_instance)
			fmt.Println(vpn_instance.Content[0])
		} else {
			vpn_instance := VPN_SPECIFIC{}
			err = json.Unmarshal(data, &vpn_instance)
			fmt.Println(vpn_instance.Content)
		}
	}
}

func Membership_switch(service string, object_id string, data []byte, err error) {
	switch service {
	case "address":
		address_instance := ADDRESS{}
		err = json.Unmarshal(data, &address_instance)
		fmt.Println(address_instance.Content[0])

	case "app_settings":
		app_settings_instance := APP_SETTINGS{}
		err = json.Unmarshal(data, &app_settings_instance)
		fmt.Println(app_settings_instance.Content)

	case "country":
		if object_id == "" {
			country_instance := COUNTRY{}
			err = json.Unmarshal(data, &country_instance)
			fmt.Println(country_instance.Content[0])
		} else {
			country_instance := COUNTRY_SPECIFIC{}
			err = json.Unmarshal(data, &country_instance)
			fmt.Println(country_instance.Content)
		}

	case "currency":
		if object_id == "" {
			currency_instance := CURRENCY{}
			err = json.Unmarshal(data, &currency_instance)
			fmt.Println(currency_instance.Content[0])
		} else {
			currency_instance := CURRENCY_SPECIFIC{}
			err = json.Unmarshal(data, &currency_instance)
			fmt.Println(currency_instance.Content)
		}

	case "department":
		if object_id == "" {
			department_instance := DEPARTMENT{}
			err = json.Unmarshal(data, &department_instance)
			fmt.Println(department_instance.Content[0])
		} else {
			department_instance := DEPARTMENT_SPECIFIC{}
			err = json.Unmarshal(data, &department_instance)
			fmt.Println(department_instance.Content)
		}

	case "email_confirmation":
		email_instance := EMAIL_CONFIRMED{}
		err = json.Unmarshal(data, &email_instance)
		fmt.Println(email_instance.Content)

	case "language":
		if object_id == "" {
			language_instance := LANGUAGE{}
			err = json.Unmarshal(data, &language_instance)
			fmt.Println(language_instance.Content[0])
		} else {
			language_instance := LANGUAGE_SPECIFIC{}
			err = json.Unmarshal(data, &language_instance)
			fmt.Println(language_instance.Content)
		}

	case "member":
		if object_id == "" {
			member_instance := MEMBER{}
			err = json.Unmarshal(data, &member_instance)
			fmt.Println(member_instance.Content[0])
		} else {
			member_instance := MEMBER_SPECIFIC{}
			err = json.Unmarshal(data, &member_instance)
			fmt.Println(member_instance.Content)
		}

	case "profile":
		if object_id == "" {
			profile_instance := PROFILE{}
			err = json.Unmarshal(data, &profile_instance)
			fmt.Println(profile_instance.Content[0])
		} else {
			profile_instance := PROFILE_SPECIFIC{}
			err = json.Unmarshal(data, &profile_instance)
			fmt.Println(profile_instance.Content)
		}

	case "team":
		if object_id == "" {
			team_instance := TEAM{}
			err = json.Unmarshal(data, &team_instance)
			fmt.Println(team_instance.Content[0])
		} else {
			team_instance := TEAM_SPECIFIC{}
			err = json.Unmarshal(data, &team_instance)
			fmt.Println(team_instance.Content)
		}

	case "territory":
		if object_id == "" {
			territory_instance := TERRITORY{}
			err = json.Unmarshal(data, &territory_instance)
			fmt.Println(territory_instance.Content[0])
		} else {
			territory_instance := TERRITORY_SPECIFIC{}
			err = json.Unmarshal(data, &territory_instance)
			fmt.Println(territory_instance.Content)
		}

	case "transaction_type":
		if object_id == "" {
			transaction_type_instance := TRANSACTION_TYPE{}
			err = json.Unmarshal(data, &transaction_type_instance)
			fmt.Println(transaction_type_instance.Content[0])
		} else {
			transaction_type_instance := TRANSACTION_TYPE_SPECIFIC{}
			err = json.Unmarshal(data, &transaction_type_instance)
			fmt.Println(transaction_type_instance.Content)
		}

	case "user":
		if object_id == "" {
			user_instance := USER{}
			err = json.Unmarshal(data, &user_instance)
			fmt.Println(user_instance.Content[0])
		} else {
			user_instance := USER_SPECIFIC{}
			err = json.Unmarshal(data, &user_instance)
			fmt.Println(user_instance.Content)
		}
	}
}

func Training_switch(service string, object_id string, data []byte, err error) {
	switch service {
	case "class":
		if object_id == "" {
			class_instance := CLASS{}
			err = json.Unmarshal(data, &class_instance)
			fmt.Println(class_instance.Content[0])
		} else {
			class_instance := CLASS_SPECIFIC{}
			err = json.Unmarshal(data, &class_instance)
			fmt.Println(class_instance.Content)
		}

	case "student":
		if object_id == "" {
			student_instance := STUDENT{}
			err = json.Unmarshal(data, &student_instance)
			fmt.Println(student_instance.Content[0])
		} else {
			student_instance := STUDENT_SPECIFIC{}
			err = json.Unmarshal(data, &student_instance)
			fmt.Println(student_instance.Content)
		}

	case "syllabus":
		if object_id == "" {
			syllabus_instance := SYLLABUS{}
			err = json.Unmarshal(data, &syllabus_instance)
			fmt.Println(syllabus_instance.Content[0])
		} else {
			syllabus_instance := SYLLABUS_SPECIFIC{}
			err = json.Unmarshal(data, &syllabus_instance)
			fmt.Println(syllabus_instance.Content)
		}
	}
}
