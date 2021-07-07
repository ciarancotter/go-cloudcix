package api

type Allocation struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		AddressID    int    `json:"address_id"`
		AddressRange string `json:"address_range"`
		Asn          struct {
			AllocationsInUse int    `json:"allocations_in_use"`
			Created          string `json:"created"`
			ID               int    `json:"id"`
			MemberID         int    `json:"member_id"`
			Number           int    `json:"number"`
			Updated          string `json:"updated"`
			URI              string `json:"uri"`
		} `json:"asn"`
		Created      string `json:"created"`
		ID           int    `json:"id"`
		ModifiedBy   int    `json:"modified_by"`
		Name         string `json:"name"`
		SubnetsInUse int    `json:"subnets_in_use"`
		Updated      string `json:"updated"`
		URI          string `json:"uri"`
	} `json:"content"`
}

type AllocationSpecific struct {
	Content struct {
		AddressID    int    `json:"address_id"`
		AddressRange string `json:"address_range"`
		Asn          struct {
			AllocationsInUse int    `json:"allocations_in_use"`
			Created          string `json:"created"`
			ID               int    `json:"id"`
			MemberID         int    `json:"member_id"`
			Number           int    `json:"number"`
			Updated          string `json:"updated"`
			URI              string `json:"uri"`
		} `json:"asn"`
		Created      string `json:"created"`
		ID           int    `json:"id"`
		ModifiedBy   int    `json:"modified_by"`
		Name         string `json:"name"`
		SubnetsInUse int    `json:"subnets_in_use"`
		Updated      string `json:"updated"`
		URI          string `json:"uri"`
	} `json:"content"`
}

type IAASAppSettings struct {
	Content struct {
		Created         string `json:"created"`
		ID              int    `json:"id"`
		IpmiCredentials string `json:"ipmi_credentials"`
		IpmiHost        string `json:"ipmi_host"`
		IpmiUsername    string `json:"ipmi_username"`
		Rage4APIKey     string `json:"rage4_api_key"`
		Rage4Email      string `json:"rage4_email"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
	} `json:"content"`
}

type Asn struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		AllocationsInUse int    `json:"allocations_in_use"`
		Created          string `json:"created"`
		ID               int    `json:"id"`
		MemberID         int    `json:"member_id"`
		Number           int    `json:"number"`
		Updated          string `json:"updated"`
		URI              string `json:"uri"`
	} `json:"content"`
}

type AsnSpecific struct {
	Content struct {
		AllocationsInUse int    `json:"allocations_in_use"`
		Created          string `json:"created"`
		ID               int    `json:"id"`
		MemberID         int    `json:"member_id"`
		Number           int    `json:"number"`
		Updated          string `json:"updated"`
		URI              string `json:"uri"`
	} `json:"content"`
}

type CixBlacklist struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Cidr       string `json:"cidr"`
		Comment    string `json:"comment"`
		Created    string `json:"created"`
		ModifiedBy int    `json:"modified_by"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type CixBlacklistSpecific struct {
	Content struct {
		Cidr       string `json:"cidr"`
		Comment    string `json:"comment"`
		Created    string `json:"created"`
		ModifiedBy int    `json:"modified_by"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type CixWhitelist struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Cidr       string `json:"cidr"`
		Comment    string `json:"comment"`
		Created    string `json:"created"`
		ModifiedBy int    `json:"modified_by"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type CixWhitelistSpecific struct {
	Content struct {
		Cidr       string `json:"cidr"`
		Comment    string `json:"comment"`
		Created    string `json:"created"`
		ModifiedBy int    `json:"modified_by"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type Cloud struct {
	Content struct {
		Project struct {
			AddressID       int    `json:"address_id"`
			Closed          bool   `json:"closed"`
			CloudURI        string `json:"cloud_uri"`
			Created         string `json:"created"`
			ID              int    `json:"id"`
			ManagerID       int    `json:"manager_id"`
			MinState        int    `json:"min_state"`
			Name            string `json:"name"`
			RegionID        int    `json:"region_id"`
			ShutDown        bool   `json:"shut_down"`
			Stable          bool   `json:"stable"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
		} `json:"project"`
		VirtualRouter struct {
			Created       string `json:"created"`
			FirewallRules []struct {
				Allow           bool   `json:"allow"`
				Created         string `json:"created"`
				DebugLogging    bool   `json:"debug_logging"`
				Description     string `json:"description"`
				Destination     string `json:"destination"`
				ID              int    `json:"id"`
				Order           int    `json:"order"`
				PciLogging      bool   `json:"pci_logging"`
				Port            int    `json:"port"`
				Protocol        string `json:"protocol"`
				Source          string `json:"source"`
				Updated         string `json:"updated"`
				VirtualRouterID int    `json:"virtual_router_id"`
			} `json:"firewall_rules"`
			ID        int `json:"id"`
			IPAddress struct {
				Address     string `json:"address"`
				Cloud       bool   `json:"cloud"`
				Created     string `json:"created"`
				Credentials string `json:"credentials"`
				ID          int    `json:"id"`
				Location    string `json:"location"`
				ModifiedBy  int    `json:"modified_by"`
				Name        string `json:"name"`
				PublicIP    struct {
					Address string `json:"address"`
					ID      int    `json:"id"`
				} `json:"public_ip"`
				Subnet struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Children []struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Allocation   struct {
							AddressID    int    `json:"address_id"`
							AddressRange string `json:"address_range"`
							Asn          struct {
								AllocationsInUse int    `json:"allocations_in_use"`
								Created          string `json:"created"`
								ID               int    `json:"id"`
								MemberID         int    `json:"member_id"`
								Number           int    `json:"number"`
								Updated          string `json:"updated"`
								URI              string `json:"uri"`
							} `json:"asn"`
							Created      string `json:"created"`
							ID           int    `json:"id"`
							ModifiedBy   int    `json:"modified_by"`
							Name         string `json:"name"`
							SubnetsInUse int    `json:"subnets_in_use"`
							Updated      string `json:"updated"`
							URI          string `json:"uri"`
						} `json:"allocation"`
						Cloud           bool   `json:"cloud"`
						Created         string `json:"created"`
						Gateway         string `json:"gateway"`
						ID              int    `json:"id"`
						IpsInUse        int    `json:"ips_in_use"`
						ModifiedBy      int    `json:"modified_by"`
						Name            string `json:"name"`
						SubnetsInUse    int    `json:"subnets_in_use"`
						Updated         string `json:"updated"`
						URI             string `json:"uri"`
						VirtualRouterID int    `json:"virtual_router_id"`
						Vlan            int    `json:"vlan"`
						Vxlan           int    `json:"vxlan"`
					} `json:"children"`
					Cloud      bool   `json:"cloud"`
					Created    string `json:"created"`
					Gateway    string `json:"gateway"`
					ID         int    `json:"id"`
					IpsInUse   int    `json:"ips_in_use"`
					ModifiedBy int    `json:"modified_by"`
					Name       string `json:"name"`
					Parent     struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Allocation   struct {
							AddressID    int    `json:"address_id"`
							AddressRange string `json:"address_range"`
							Asn          struct {
								AllocationsInUse int    `json:"allocations_in_use"`
								Created          string `json:"created"`
								ID               int    `json:"id"`
								MemberID         int    `json:"member_id"`
								Number           int    `json:"number"`
								Updated          string `json:"updated"`
								URI              string `json:"uri"`
							} `json:"asn"`
							Created      string `json:"created"`
							ID           int    `json:"id"`
							ModifiedBy   int    `json:"modified_by"`
							Name         string `json:"name"`
							SubnetsInUse int    `json:"subnets_in_use"`
							Updated      string `json:"updated"`
							URI          string `json:"uri"`
						} `json:"allocation"`
						Cloud           bool   `json:"cloud"`
						Created         string `json:"created"`
						Gateway         string `json:"gateway"`
						ID              int    `json:"id"`
						IpsInUse        int    `json:"ips_in_use"`
						ModifiedBy      int    `json:"modified_by"`
						Name            string `json:"name"`
						SubnetsInUse    int    `json:"subnets_in_use"`
						Updated         string `json:"updated"`
						URI             string `json:"uri"`
						VirtualRouterID int    `json:"virtual_router_id"`
						Vlan            int    `json:"vlan"`
						Vxlan           int    `json:"vxlan"`
					} `json:"parent"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"subnet"`
				Updated string `json:"updated"`
				VMID    int    `json:"vm_id"`
			} `json:"ip_address"`
			Project struct {
				AddressID       int    `json:"address_id"`
				Closed          bool   `json:"closed"`
				CloudURI        string `json:"cloud_uri"`
				Created         string `json:"created"`
				ID              int    `json:"id"`
				ManagerID       int    `json:"manager_id"`
				MinState        int    `json:"min_state"`
				Name            string `json:"name"`
				RegionID        int    `json:"region_id"`
				ShutDown        bool   `json:"shut_down"`
				Stable          bool   `json:"stable"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
			} `json:"project"`
			RouterID int `json:"router_id"`
			State    int `json:"state"`
			Subnets  []struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				Vlan         int    `json:"vlan"`
				Vxlan        int    `json:"vxlan"`
			} `json:"subnets"`
			Updated string `json:"updated"`
			URI     string `json:"uri"`
		} `json:"virtual_router"`
		Vms []struct {
			CPU           int      `json:"cpu"`
			Created       string   `json:"created"`
			DNS           string   `json:"dns"`
			Emails        []string `json:"emails"`
			GatewaySubnet struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Children []struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"children"`
				Cloud      bool   `json:"cloud"`
				Created    string `json:"created"`
				Gateway    string `json:"gateway"`
				ID         int    `json:"id"`
				IpsInUse   int    `json:"ips_in_use"`
				ModifiedBy int    `json:"modified_by"`
				Name       string `json:"name"`
				Parent     struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"parent"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"gateway_subnet"`
			History struct {
				CPUQuantity      int    `json:"cpu_quantity"`
				CPUSku           string `json:"cpu_sku"`
				Created          string `json:"created"`
				CustomerAddress  int    `json:"customer_address"`
				ImageQuantity    int    `json:"image_quantity"`
				ImageSku         string `json:"image_sku"`
				ModifiedBy       int    `json:"modified_by"`
				NatQuantity      int    `json:"nat_quantity"`
				NatSku           string `json:"nat_sku"`
				ProjectID        int    `json:"project_id"`
				ProjectVMName    string `json:"project_vm_name"`
				RAMQuantity      int    `json:"ram_quantity"`
				RAMSku           string `json:"ram_sku"`
				State            int    `json:"state"`
				StorageHistories []struct {
					GbQuantity  int    `json:"gb_quantity"`
					GbSku       string `json:"gb_sku"`
					StorageID   int    `json:"storage_id"`
					StorageName string `json:"storage_name"`
				} `json:"storage_histories"`
				VMID int `json:"vm_id"`
			} `json:"history"`
			ID    int `json:"id"`
			Image struct {
				AnswerFileName string `json:"answer_file_name"`
				Created        string `json:"created"`
				DisplayName    string `json:"display_name"`
				Filename       string `json:"filename"`
				ID             int    `json:"id"`
				MultipleIps    bool   `json:"multiple_ips"`
				OsVariant      string `json:"os_variant"`
				Regions        []int  `json:"regions"`
				Updated        string `json:"updated"`
				URI            string `json:"uri"`
			} `json:"image"`
			IPAddresses []struct {
				Address     string `json:"address"`
				Cloud       bool   `json:"cloud"`
				Created     string `json:"created"`
				Credentials string `json:"credentials"`
				ID          int    `json:"id"`
				Location    string `json:"location"`
				ModifiedBy  int    `json:"modified_by"`
				Name        string `json:"name"`
				PublicIP    struct {
					Address string `json:"address"`
					ID      int    `json:"id"`
				} `json:"public_ip"`
				Subnet struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Children []struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Allocation   struct {
							AddressID    int    `json:"address_id"`
							AddressRange string `json:"address_range"`
							Asn          struct {
								AllocationsInUse int    `json:"allocations_in_use"`
								Created          string `json:"created"`
								ID               int    `json:"id"`
								MemberID         int    `json:"member_id"`
								Number           int    `json:"number"`
								Updated          string `json:"updated"`
								URI              string `json:"uri"`
							} `json:"asn"`
							Created      string `json:"created"`
							ID           int    `json:"id"`
							ModifiedBy   int    `json:"modified_by"`
							Name         string `json:"name"`
							SubnetsInUse int    `json:"subnets_in_use"`
							Updated      string `json:"updated"`
							URI          string `json:"uri"`
						} `json:"allocation"`
						Cloud           bool   `json:"cloud"`
						Created         string `json:"created"`
						Gateway         string `json:"gateway"`
						ID              int    `json:"id"`
						IpsInUse        int    `json:"ips_in_use"`
						ModifiedBy      int    `json:"modified_by"`
						Name            string `json:"name"`
						SubnetsInUse    int    `json:"subnets_in_use"`
						Updated         string `json:"updated"`
						URI             string `json:"uri"`
						VirtualRouterID int    `json:"virtual_router_id"`
						Vlan            int    `json:"vlan"`
						Vxlan           int    `json:"vxlan"`
					} `json:"children"`
					Cloud      bool   `json:"cloud"`
					Created    string `json:"created"`
					Gateway    string `json:"gateway"`
					ID         int    `json:"id"`
					IpsInUse   int    `json:"ips_in_use"`
					ModifiedBy int    `json:"modified_by"`
					Name       string `json:"name"`
					Parent     struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Allocation   struct {
							AddressID    int    `json:"address_id"`
							AddressRange string `json:"address_range"`
							Asn          struct {
								AllocationsInUse int    `json:"allocations_in_use"`
								Created          string `json:"created"`
								ID               int    `json:"id"`
								MemberID         int    `json:"member_id"`
								Number           int    `json:"number"`
								Updated          string `json:"updated"`
								URI              string `json:"uri"`
							} `json:"asn"`
							Created      string `json:"created"`
							ID           int    `json:"id"`
							ModifiedBy   int    `json:"modified_by"`
							Name         string `json:"name"`
							SubnetsInUse int    `json:"subnets_in_use"`
							Updated      string `json:"updated"`
							URI          string `json:"uri"`
						} `json:"allocation"`
						Cloud           bool   `json:"cloud"`
						Created         string `json:"created"`
						Gateway         string `json:"gateway"`
						ID              int    `json:"id"`
						IpsInUse        int    `json:"ips_in_use"`
						ModifiedBy      int    `json:"modified_by"`
						Name            string `json:"name"`
						SubnetsInUse    int    `json:"subnets_in_use"`
						Updated         string `json:"updated"`
						URI             string `json:"uri"`
						VirtualRouterID int    `json:"virtual_router_id"`
						Vlan            int    `json:"vlan"`
						Vxlan           int    `json:"vxlan"`
					} `json:"parent"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"subnet"`
				Updated string `json:"updated"`
				VMID    int    `json:"vm_id"`
			} `json:"ip_addresses"`
			Name    string `json:"name"`
			Project struct {
				AddressID       int    `json:"address_id"`
				Closed          bool   `json:"closed"`
				CloudURI        string `json:"cloud_uri"`
				Created         string `json:"created"`
				ID              int    `json:"id"`
				ManagerID       int    `json:"manager_id"`
				MinState        int    `json:"min_state"`
				Name            string `json:"name"`
				RegionID        int    `json:"region_id"`
				ShutDown        bool   `json:"shut_down"`
				Stable          bool   `json:"stable"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
			} `json:"project"`
			RAM         int    `json:"ram"`
			ServerID    int    `json:"server_id"`
			State       int    `json:"state"`
			StorageType string `json:"storage_type"`
			Storages    []struct {
				Created string `json:"created"`
				Gb      int    `json:"gb"`
				ID      int    `json:"id"`
				Name    string `json:"name"`
				Primary bool   `json:"primary"`
				Updated string `json:"updated"`
				URI     string `json:"uri"`
				VMID    int    `json:"vm_id"`
			} `json:"storages"`
			Updated string `json:"updated"`
			URI     string `json:"uri"`
		} `json:"vms"`
		Vpns []struct {
			Created             string   `json:"created"`
			Description         string   `json:"description"`
			Emails              []string `json:"emails"`
			ID                  int      `json:"id"`
			IkeAuthentication   string   `json:"ike_authentication"`
			IkeDhGroups         string   `json:"ike_dh_groups"`
			IkeEncryption       string   `json:"ike_encryption"`
			IkeLifetime         int      `json:"ike_lifetime"`
			IkeMode             string   `json:"ike_mode"`
			IkePreSharedKey     string   `json:"ike_pre_shared_key"`
			IkePublicIP         string   `json:"ike_public_ip"`
			IkeVersion          string   `json:"ike_version"`
			IpsecAuthentication string   `json:"ipsec_authentication"`
			IpsecEncryption     string   `json:"ipsec_encryption"`
			IpsecEstablishTime  string   `json:"ipsec_establish_time"`
			IpsecLifetime       int      `json:"ipsec_lifetime"`
			IpsecPfsGroups      string   `json:"ipsec_pfs_groups"`
			Routes              []struct {
				Created     string `json:"created"`
				ID          int    `json:"id"`
				LocalSubnet struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Children []struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Allocation   struct {
							AddressID    int    `json:"address_id"`
							AddressRange string `json:"address_range"`
							Asn          struct {
								AllocationsInUse int    `json:"allocations_in_use"`
								Created          string `json:"created"`
								ID               int    `json:"id"`
								MemberID         int    `json:"member_id"`
								Number           int    `json:"number"`
								Updated          string `json:"updated"`
								URI              string `json:"uri"`
							} `json:"asn"`
							Created      string `json:"created"`
							ID           int    `json:"id"`
							ModifiedBy   int    `json:"modified_by"`
							Name         string `json:"name"`
							SubnetsInUse int    `json:"subnets_in_use"`
							Updated      string `json:"updated"`
							URI          string `json:"uri"`
						} `json:"allocation"`
						Cloud           bool   `json:"cloud"`
						Created         string `json:"created"`
						Gateway         string `json:"gateway"`
						ID              int    `json:"id"`
						IpsInUse        int    `json:"ips_in_use"`
						ModifiedBy      int    `json:"modified_by"`
						Name            string `json:"name"`
						SubnetsInUse    int    `json:"subnets_in_use"`
						Updated         string `json:"updated"`
						URI             string `json:"uri"`
						VirtualRouterID int    `json:"virtual_router_id"`
						Vlan            int    `json:"vlan"`
						Vxlan           int    `json:"vxlan"`
					} `json:"children"`
					Cloud      bool   `json:"cloud"`
					Created    string `json:"created"`
					Gateway    string `json:"gateway"`
					ID         int    `json:"id"`
					IpsInUse   int    `json:"ips_in_use"`
					ModifiedBy int    `json:"modified_by"`
					Name       string `json:"name"`
					Parent     struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Allocation   struct {
							AddressID    int    `json:"address_id"`
							AddressRange string `json:"address_range"`
							Asn          struct {
								AllocationsInUse int    `json:"allocations_in_use"`
								Created          string `json:"created"`
								ID               int    `json:"id"`
								MemberID         int    `json:"member_id"`
								Number           int    `json:"number"`
								Updated          string `json:"updated"`
								URI              string `json:"uri"`
							} `json:"asn"`
							Created      string `json:"created"`
							ID           int    `json:"id"`
							ModifiedBy   int    `json:"modified_by"`
							Name         string `json:"name"`
							SubnetsInUse int    `json:"subnets_in_use"`
							Updated      string `json:"updated"`
							URI          string `json:"uri"`
						} `json:"allocation"`
						Cloud           bool   `json:"cloud"`
						Created         string `json:"created"`
						Gateway         string `json:"gateway"`
						ID              int    `json:"id"`
						IpsInUse        int    `json:"ips_in_use"`
						ModifiedBy      int    `json:"modified_by"`
						Name            string `json:"name"`
						SubnetsInUse    int    `json:"subnets_in_use"`
						Updated         string `json:"updated"`
						URI             string `json:"uri"`
						VirtualRouterID int    `json:"virtual_router_id"`
						Vlan            int    `json:"vlan"`
						Vxlan           int    `json:"vxlan"`
					} `json:"parent"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"local_subnet"`
				RemoteSubnet string `json:"remote_subnet"`
				Updated      string `json:"updated"`
			} `json:"routes"`
			SendEmail       bool   `json:"send_email"`
			StifNumber      string `json:"stif_number"`
			TrafficSelector string `json:"traffic_selector"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
		} `json:"vpns"`
	} `json:"content"`
}

type Domain struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Created    string `json:"created"`
		ID         int    `json:"id"`
		MemberID   int    `json:"member_id"`
		ModifiedBy int    `json:"modified_by"`
		Name       string `json:"name"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type DomainSpecific struct {
	Content struct {
		Created    string `json:"created"`
		ID         int    `json:"id"`
		MemberID   int    `json:"member_id"`
		ModifiedBy int    `json:"modified_by"`
		Name       string `json:"name"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type Image struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		AnswerFileName string `json:"answer_file_name"`
		Created        string `json:"created"`
		DisplayName    string `json:"display_name"`
		Filename       string `json:"filename"`
		ID             int    `json:"id"`
		MultipleIps    bool   `json:"multiple_ips"`
		OsVariant      string `json:"os_variant"`
		Regions        []int  `json:"regions"`
		Updated        string `json:"updated"`
		URI            string `json:"uri"`
	} `json:"content"`
}

type ImageSpecific struct {
	Content struct {
		AnswerFileName string `json:"answer_file_name"`
		Created        string `json:"created"`
		DisplayName    string `json:"display_name"`
		Filename       string `json:"filename"`
		ID             int    `json:"id"`
		MultipleIps    bool   `json:"multiple_ips"`
		OsVariant      string `json:"os_variant"`
		Regions        []int  `json:"regions"`
		Updated        string `json:"updated"`
		URI            string `json:"uri"`
	} `json:"content"`
}

type Interface struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Created    string `json:"created"`
		Details    string `json:"details"`
		Enabled    bool   `json:"enabled"`
		Hostname   string `json:"hostname"`
		ID         int    `json:"id"`
		IPAddress  string `json:"ip_address"`
		MacAddress string `json:"mac_address"`
		ServerID   int    `json:"server_id"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type InterfaceSpecific struct {
	Content struct {
		Created    string `json:"created"`
		Details    string `json:"details"`
		Enabled    bool   `json:"enabled"`
		Hostname   string `json:"hostname"`
		ID         int    `json:"id"`
		IPAddress  string `json:"ip_address"`
		MacAddress string `json:"mac_address"`
		ServerID   int    `json:"server_id"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type IpAddress struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Address     string `json:"address"`
		Cloud       bool   `json:"cloud"`
		Created     string `json:"created"`
		Credentials string `json:"credentials"`
		ID          int    `json:"id"`
		Location    string `json:"location"`
		ModifiedBy  int    `json:"modified_by"`
		Name        string `json:"name"`
		PublicIP    struct {
			Address string `json:"address"`
			ID      int    `json:"id"`
		} `json:"public_ip"`
		Subnet struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Allocation   struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Asn          struct {
					AllocationsInUse int    `json:"allocations_in_use"`
					Created          string `json:"created"`
					ID               int    `json:"id"`
					MemberID         int    `json:"member_id"`
					Number           int    `json:"number"`
					Updated          string `json:"updated"`
					URI              string `json:"uri"`
				} `json:"asn"`
				Created      string `json:"created"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				SubnetsInUse int    `json:"subnets_in_use"`
				Updated      string `json:"updated"`
				URI          string `json:"uri"`
			} `json:"allocation"`
			Children []struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"children"`
			Cloud      bool   `json:"cloud"`
			Created    string `json:"created"`
			Gateway    string `json:"gateway"`
			ID         int    `json:"id"`
			IpsInUse   int    `json:"ips_in_use"`
			ModifiedBy int    `json:"modified_by"`
			Name       string `json:"name"`
			Parent     struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"parent"`
			SubnetsInUse    int    `json:"subnets_in_use"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
			Vlan            int    `json:"vlan"`
			Vxlan           int    `json:"vxlan"`
		} `json:"subnet"`
		Updated string `json:"updated"`
		VMID    int    `json:"vm_id"`
	} `json:"content"`
}

type IpAddressSpecific struct {
	Content struct {
		Address     string `json:"address"`
		Cloud       bool   `json:"cloud"`
		Created     string `json:"created"`
		Credentials string `json:"credentials"`
		ID          int    `json:"id"`
		Location    string `json:"location"`
		ModifiedBy  int    `json:"modified_by"`
		Name        string `json:"name"`
		PublicIP    struct {
			Address string `json:"address"`
			ID      int    `json:"id"`
		} `json:"public_ip"`
		Subnet struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Allocation   struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Asn          struct {
					AllocationsInUse int    `json:"allocations_in_use"`
					Created          string `json:"created"`
					ID               int    `json:"id"`
					MemberID         int    `json:"member_id"`
					Number           int    `json:"number"`
					Updated          string `json:"updated"`
					URI              string `json:"uri"`
				} `json:"asn"`
				Created      string `json:"created"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				SubnetsInUse int    `json:"subnets_in_use"`
				Updated      string `json:"updated"`
				URI          string `json:"uri"`
			} `json:"allocation"`
			Children []struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"children"`
			Cloud      bool   `json:"cloud"`
			Created    string `json:"created"`
			Gateway    string `json:"gateway"`
			ID         int    `json:"id"`
			IpsInUse   int    `json:"ips_in_use"`
			ModifiedBy int    `json:"modified_by"`
			Name       string `json:"name"`
			Parent     struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"parent"`
			SubnetsInUse    int    `json:"subnets_in_use"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
			Vlan            int    `json:"vlan"`
			Vxlan           int    `json:"vxlan"`
		} `json:"subnet"`
		Updated string `json:"updated"`
		VMID    int    `json:"vm_id"`
	} `json:"content"`
}

type Ipmi struct {
	Content []struct {
		ClientIP   string `json:"client_ip"`
		Created    string `json:"created"`
		CustomerIP struct {
			Address     string      `json:"address"`
			ID          int         `json:"id"`
			Cloud       bool        `json:"cloud"`
			Created     string      `json:"created"`
			Credentials string      `json:"credentials"`
			Location    string      `json:"location"`
			ModifiedBy  int         `json:"modified_by"`
			Name        string      `json:"name"`
			PublicIP    interface{} `json:"public_ip"`
			Subnet      struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool          `json:"cloud"`
				Created         string        `json:"created"`
				Gateway         interface{}   `json:"gateway"`
				ID              int           `json:"id"`
				IpsInUse        int           `json:"ips_in_use"`
				ModifiedBy      int           `json:"modified_by"`
				SubnetsInUse    int           `json:"subnets_in_use"`
				Name            string        `json:"name"`
				Updated         string        `json:"updated"`
				URI             string        `json:"uri"`
				VirtualRouterID interface{}   `json:"virtual_router_id"`
				Vlan            int           `json:"vlan"`
				Vxlan           int           `json:"vxlan"`
				Children        []interface{} `json:"children"`
				Parent          interface{}   `json:"parent"`
			} `json:"subnet"`
			Updated string      `json:"updated"`
			VMID    interface{} `json:"vm_id"`
		} `json:"customer_ip"`
		ID         int  `json:"id"`
		InUse      bool `json:"in_use"`
		ModifiedBy int  `json:"modified_by"`
		PoolIP     struct {
			Created    string `json:"created"`
			Domain     string `json:"domain"`
			ID         int    `json:"id"`
			IPAddress  string `json:"ip_address"`
			ModifiedBy int    `json:"modified_by"`
			Updated    string `json:"updated"`
			URI        string `json:"uri"`
		} `json:"pool_ip"`
		Removed string `json:"removed"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
	Metadata struct {
		TotalRecords int           `json:"total_records"`
		Page         int           `json:"page"`
		Limit        int           `json:"limit"`
		Order        string        `json:"order"`
		Warnings     []interface{} `json:"warnings"`
	} `json:"_metadata"`
}

type IpmiSpecific struct {
	Content struct {
		ClientIP   string `json:"client_ip"`
		Created    string `json:"created"`
		CustomerIP struct {
			Address     string `json:"address"`
			Cloud       bool   `json:"cloud"`
			Created     string `json:"created"`
			Credentials string `json:"credentials"`
			ID          int    `json:"id"`
			Location    string `json:"location"`
			ModifiedBy  int    `json:"modified_by"`
			Name        string `json:"name"`
			PublicIP    struct {
				Address string `json:"address"`
				ID      int    `json:"id"`
			} `json:"public_ip"`
			Subnet struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Children []struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"children"`
				Cloud      bool   `json:"cloud"`
				Created    string `json:"created"`
				Gateway    string `json:"gateway"`
				ID         int    `json:"id"`
				IpsInUse   int    `json:"ips_in_use"`
				ModifiedBy int    `json:"modified_by"`
				Name       string `json:"name"`
				Parent     struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"parent"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"subnet"`
			Updated string `json:"updated"`
			VMID    int    `json:"vm_id"`
		} `json:"customer_ip"`
		ID         int  `json:"id"`
		InUse      bool `json:"in_use"`
		ModifiedBy int  `json:"modified_by"`
		PoolIP     struct {
			Created    string `json:"created"`
			Domain     string `json:"domain"`
			ID         int    `json:"id"`
			IPAddress  string `json:"ip_address"`
			ModifiedBy int    `json:"modified_by"`
			Updated    string `json:"updated"`
			URI        string `json:"uri"`
		} `json:"pool_ip"`
		Removed string `json:"removed"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type PolicyLog struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Access             string `json:"access"`
		DestinationAddress string `json:"destination_address"`
		DestinationPort    int    `json:"destination_port"`
		ServiceName        string `json:"service_name"`
		SourceAddress      string `json:"source_address"`
		SourcePort         int    `json:"source_port"`
		Timestamp          string `json:"timestamp"`
	} `json:"content"`
}

type PoolIp struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Created    string `json:"created"`
		Domain     string `json:"domain"`
		ID         int    `json:"id"`
		IPAddress  string `json:"ip_address"`
		ModifiedBy int    `json:"modified_by"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type PoolIpSpecific struct {
	Content struct {
		Created    string `json:"created"`
		Domain     string `json:"domain"`
		ID         int    `json:"id"`
		IPAddress  string `json:"ip_address"`
		ModifiedBy int    `json:"modified_by"`
		Updated    string `json:"updated"`
		URI        string `json:"uri"`
	} `json:"content"`
}

type Project struct {
	Content []struct {
		AddressID       int    `json:"address_id"`
		Closed          bool   `json:"closed"`
		CloudURI        string `json:"cloud_uri"`
		Created         string `json:"created"`
		ID              int    `json:"id"`
		ManagerID       int    `json:"manager_id"`
		MinState        int    `json:"min_state"`
		Name            string `json:"name"`
		RegionID        int    `json:"region_id"`
		ShutDown        bool   `json:"shut_down"`
		Stable          bool   `json:"stable"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
		VirtualRouterID int    `json:"virtual_router_id"`
	} `json:"content"`
	Metadata struct {
		Page         int           `json:"page"`
		Limit        int           `json:"limit"`
		Order        string        `json:"order"`
		TotalRecords int           `json:"total_records"`
		Warnings     []interface{} `json:"warnings"`
	} `json:"_metadata"`
}

type ProjectSpecific struct {
	Content struct {
		AddressID       int    `json:"address_id"`
		Closed          bool   `json:"closed"`
		CloudURI        string `json:"cloud_uri"`
		Created         string `json:"created"`
		ID              int    `json:"id"`
		ManagerID       int    `json:"manager_id"`
		MinState        int    `json:"min_state"`
		Name            string `json:"name"`
		Note            string `json:"note"`
		RegionID        int    `json:"region_id"`
		ShutDown        bool   `json:"shut_down"`
		Stable          bool   `json:"stable"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
		VirtualRouterID int    `json:"virtual_router_id"`
	} `json:"content"`
}

type PtrRecord struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Content string `json:"content"`
		Created string `json:"created"`
		Domain  struct {
			Created    string `json:"created"`
			ID         int    `json:"id"`
			MemberID   int    `json:"member_id"`
			ModifiedBy int    `json:"modified_by"`
			Name       int    `json:"name"`
			Updated    string `json:"updated"`
			URI        string `json:"uri"`
		} `json:"domain"`
		Failover        bool   `json:"failover"`
		FailoverContent string `json:"failover_content"`
		ID              int    `json:"id"`
		IPAddress       string `json:"ip_address"`
		MemberID        int    `json:"member_id"`
		ModifiedBy      int    `json:"modified_by"`
		Name            string `json:"name"`
		Pk              int    `json:"pk"`
		Priority        int    `json:"priority"`
		TimeToLive      int    `json:"time_to_live"`
		Type            string `json:"type"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
	} `json:"content"`
}

type PtrRecordSpecific struct {
	Content struct {
		Content string `json:"content"`
		Created string `json:"created"`
		Domain  struct {
			Created    string `json:"created"`
			ID         int    `json:"id"`
			MemberID   int    `json:"member_id"`
			ModifiedBy int    `json:"modified_by"`
			Name       int    `json:"name"`
			Updated    string `json:"updated"`
			URI        string `json:"uri"`
		} `json:"domain"`
		Failover        bool   `json:"failover"`
		FailoverContent string `json:"failover_content"`
		ID              int    `json:"id"`
		IPAddress       string `json:"ip_address"`
		MemberID        int    `json:"member_id"`
		ModifiedBy      int    `json:"modified_by"`
		Name            string `json:"name"`
		Pk              int    `json:"pk"`
		Priority        int    `json:"priority"`
		TimeToLive      int    `json:"time_to_live"`
		Type            string `json:"type"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
	} `json:"content"`
}

type Record struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Content string `json:"content"`
		Created string `json:"created"`
		Domain  struct {
			Created    string `json:"created"`
			ID         int    `json:"id"`
			MemberID   int    `json:"member_id"`
			ModifiedBy int    `json:"modified_by"`
			Name       int    `json:"name"`
			Updated    string `json:"updated"`
			URI        string `json:"uri"`
		} `json:"domain"`
		Failover        bool   `json:"failover"`
		FailoverContent string `json:"failover_content"`
		ID              int    `json:"id"`
		IPAddress       string `json:"ip_address"`
		MemberID        int    `json:"member_id"`
		ModifiedBy      int    `json:"modified_by"`
		Name            string `json:"name"`
		Pk              int    `json:"pk"`
		Priority        int    `json:"priority"`
		TimeToLive      int    `json:"time_to_live"`
		Type            string `json:"type"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
	} `json:"content"`
}

type RecordSpecific struct {
	Content struct {
		Content string `json:"content"`
		Created string `json:"created"`
		Domain  struct {
			Created    string `json:"created"`
			ID         int    `json:"id"`
			MemberID   int    `json:"member_id"`
			ModifiedBy int    `json:"modified_by"`
			Name       int    `json:"name"`
			Updated    string `json:"updated"`
			URI        string `json:"uri"`
		} `json:"domain"`
		Failover        bool   `json:"failover"`
		FailoverContent string `json:"failover_content"`
		ID              int    `json:"id"`
		IPAddress       string `json:"ip_address"`
		MemberID        int    `json:"member_id"`
		ModifiedBy      int    `json:"modified_by"`
		Name            string `json:"name"`
		Pk              int    `json:"pk"`
		Priority        int    `json:"priority"`
		TimeToLive      int    `json:"time_to_live"`
		Type            string `json:"type"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
	} `json:"content"`
}

type Router struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		AssetTag            int      `json:"asset_tag"`
		Capacity            int      `json:"capacity"`
		Created             string   `json:"created"`
		Credentials         string   `json:"credentials"`
		Enabled             bool     `json:"enabled"`
		ID                  int      `json:"id"`
		ManagementInterface string   `json:"management_interface"`
		ManagementIP        string   `json:"management_ip"`
		Model               string   `json:"model"`
		OobInterface        string   `json:"oob_interface"`
		PrivateInterface    string   `json:"private_interface"`
		PublicInterface     string   `json:"public_interface"`
		PublicPortIps       []string `json:"public_port_ips"`
		RegionID            int      `json:"region_id"`
		RouterOobInterface  string   `json:"router_oob_interface"`
		SubnetIds           []int    `json:"subnet_ids"`
		Subnets             []string `json:"subnets"`
		Updated             string   `json:"updated"`
		URI                 string   `json:"uri"`
		Username            string   `json:"username"`
	} `json:"content"`
}

type RouterSpecific struct {
	Content struct {
		AssetTag            int      `json:"asset_tag"`
		Capacity            int      `json:"capacity"`
		Created             string   `json:"created"`
		Credentials         string   `json:"credentials"`
		Enabled             bool     `json:"enabled"`
		ID                  int      `json:"id"`
		ManagementInterface string   `json:"management_interface"`
		ManagementIP        string   `json:"management_ip"`
		Model               string   `json:"model"`
		OobInterface        string   `json:"oob_interface"`
		PrivateInterface    string   `json:"private_interface"`
		PublicInterface     string   `json:"public_interface"`
		PublicPortIps       []string `json:"public_port_ips"`
		RegionID            int      `json:"region_id"`
		RouterOobInterface  string   `json:"router_oob_interface"`
		SubnetIds           []int    `json:"subnet_ids"`
		Subnets             []string `json:"subnets"`
		Updated             string   `json:"updated"`
		URI                 string   `json:"uri"`
		Username            string   `json:"username"`
	} `json:"content"`
}

type Server struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		AssetTag   int    `json:"asset_tag"`
		Cores      int    `json:"cores"`
		Created    string `json:"created"`
		Enabled    bool   `json:"enabled"`
		Gb         int    `json:"gb"`
		ID         int    `json:"id"`
		Interfaces []struct {
			Created    string `json:"created"`
			Details    string `json:"details"`
			Enabled    bool   `json:"enabled"`
			Hostname   string `json:"hostname"`
			ID         int    `json:"id"`
			IPAddress  string `json:"ip_address"`
			MacAddress string `json:"mac_address"`
			ServerID   int    `json:"server_id"`
			Updated    string `json:"updated"`
			URI        string `json:"uri"`
		} `json:"interfaces"`
		Model       string `json:"model"`
		RAM         int    `json:"ram"`
		RegionID    int    `json:"region_id"`
		StorageType struct {
			Created string `json:"created"`
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Regions []int  `json:"regions"`
			Updated string `json:"updated"`
			URI     string `json:"uri"`
		} `json:"storage_type"`
		Type struct {
			Created string `json:"created"`
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Updated string `json:"updated"`
			URI     string `json:"uri"`
		} `json:"type"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
		Vcpus   int    `json:"vcpus"`
	} `json:"content"`
}

type ServerSpecific struct {
	Content struct {
		AssetTag   int    `json:"asset_tag"`
		Cores      int    `json:"cores"`
		Created    string `json:"created"`
		Enabled    bool   `json:"enabled"`
		Gb         int    `json:"gb"`
		ID         int    `json:"id"`
		Interfaces []struct {
			Created    string `json:"created"`
			Details    string `json:"details"`
			Enabled    bool   `json:"enabled"`
			Hostname   string `json:"hostname"`
			ID         int    `json:"id"`
			IPAddress  string `json:"ip_address"`
			MacAddress string `json:"mac_address"`
			ServerID   int    `json:"server_id"`
			Updated    string `json:"updated"`
			URI        string `json:"uri"`
		} `json:"interfaces"`
		Model       string `json:"model"`
		RAM         int    `json:"ram"`
		RegionID    int    `json:"region_id"`
		StorageType struct {
			Created string `json:"created"`
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Regions []int  `json:"regions"`
			Updated string `json:"updated"`
			URI     string `json:"uri"`
		} `json:"storage_type"`
		Type struct {
			Created string `json:"created"`
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Updated string `json:"updated"`
			URI     string `json:"uri"`
		} `json:"type"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
		Vcpus   int    `json:"vcpus"`
	} `json:"content"`
}

type ServerType struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Created string `json:"created"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type ServerTypeSpecific struct {
	Content struct {
		Created string `json:"created"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type StorageType struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Created string `json:"created"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Regions []int  `json:"regions"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type StorageTypeSpecific struct {
	Content struct {
		Created string `json:"created"`
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Regions []int  `json:"regions"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type Subnet struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		AddressID    int    `json:"address_id"`
		AddressRange string `json:"address_range"`
		Allocation   struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Asn          struct {
				AllocationsInUse int    `json:"allocations_in_use"`
				Created          string `json:"created"`
				ID               int    `json:"id"`
				MemberID         int    `json:"member_id"`
				Number           int    `json:"number"`
				Updated          string `json:"updated"`
				URI              string `json:"uri"`
			} `json:"asn"`
			Created      string `json:"created"`
			ID           int    `json:"id"`
			ModifiedBy   int    `json:"modified_by"`
			Name         string `json:"name"`
			SubnetsInUse int    `json:"subnets_in_use"`
			Updated      string `json:"updated"`
			URI          string `json:"uri"`
		} `json:"allocation"`
		Children []struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Allocation   struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Asn          struct {
					AllocationsInUse int    `json:"allocations_in_use"`
					Created          string `json:"created"`
					ID               int    `json:"id"`
					MemberID         int    `json:"member_id"`
					Number           int    `json:"number"`
					Updated          string `json:"updated"`
					URI              string `json:"uri"`
				} `json:"asn"`
				Created      string `json:"created"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				SubnetsInUse int    `json:"subnets_in_use"`
				Updated      string `json:"updated"`
				URI          string `json:"uri"`
			} `json:"allocation"`
			Cloud           bool   `json:"cloud"`
			Created         string `json:"created"`
			Gateway         string `json:"gateway"`
			ID              int    `json:"id"`
			IpsInUse        int    `json:"ips_in_use"`
			ModifiedBy      int    `json:"modified_by"`
			Name            string `json:"name"`
			SubnetsInUse    int    `json:"subnets_in_use"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
			Vlan            int    `json:"vlan"`
			Vxlan           int    `json:"vxlan"`
		} `json:"children"`
		Cloud      bool   `json:"cloud"`
		Created    string `json:"created"`
		Gateway    string `json:"gateway"`
		ID         int    `json:"id"`
		IpsInUse   int    `json:"ips_in_use"`
		ModifiedBy int    `json:"modified_by"`
		Name       string `json:"name"`
		Parent     struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Allocation   struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Asn          struct {
					AllocationsInUse int    `json:"allocations_in_use"`
					Created          string `json:"created"`
					ID               int    `json:"id"`
					MemberID         int    `json:"member_id"`
					Number           int    `json:"number"`
					Updated          string `json:"updated"`
					URI              string `json:"uri"`
				} `json:"asn"`
				Created      string `json:"created"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				SubnetsInUse int    `json:"subnets_in_use"`
				Updated      string `json:"updated"`
				URI          string `json:"uri"`
			} `json:"allocation"`
			Cloud           bool   `json:"cloud"`
			Created         string `json:"created"`
			Gateway         string `json:"gateway"`
			ID              int    `json:"id"`
			IpsInUse        int    `json:"ips_in_use"`
			ModifiedBy      int    `json:"modified_by"`
			Name            string `json:"name"`
			SubnetsInUse    int    `json:"subnets_in_use"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
			Vlan            int    `json:"vlan"`
			Vxlan           int    `json:"vxlan"`
		} `json:"parent"`
		SubnetsInUse    int    `json:"subnets_in_use"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
		VirtualRouterID int    `json:"virtual_router_id"`
		Vlan            int    `json:"vlan"`
		Vxlan           int    `json:"vxlan"`
	} `json:"content"`
}

type SubnetSpace struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		AddressID       int    `json:"address_id"`
		AddressRange    string `json:"address_range"`
		Cloud           bool   `json:"cloud"`
		ID              int    `json:"id"`
		IpsInUse        int    `json:"ips_in_use"`
		ModifiedBy      int    `json:"modified_by"`
		Name            string `json:"name"`
		SubnetsInUse    int    `json:"subnets_in_use"`
		VirtualRouterID int    `json:"virtual_router_id"`
		Vlan            int    `json:"vlan"`
		Vxlan           int    `json:"vxlan"`
	} `json:"content"`
}

type SubnetSpecific struct {
	Content struct {
		AddressID    int    `json:"address_id"`
		AddressRange string `json:"address_range"`
		Allocation   struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Asn          struct {
				AllocationsInUse int    `json:"allocations_in_use"`
				Created          string `json:"created"`
				ID               int    `json:"id"`
				MemberID         int    `json:"member_id"`
				Number           int    `json:"number"`
				Updated          string `json:"updated"`
				URI              string `json:"uri"`
			} `json:"asn"`
			Created      string `json:"created"`
			ID           int    `json:"id"`
			ModifiedBy   int    `json:"modified_by"`
			Name         string `json:"name"`
			SubnetsInUse int    `json:"subnets_in_use"`
			Updated      string `json:"updated"`
			URI          string `json:"uri"`
		} `json:"allocation"`
		Children []struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Allocation   struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Asn          struct {
					AllocationsInUse int    `json:"allocations_in_use"`
					Created          string `json:"created"`
					ID               int    `json:"id"`
					MemberID         int    `json:"member_id"`
					Number           int    `json:"number"`
					Updated          string `json:"updated"`
					URI              string `json:"uri"`
				} `json:"asn"`
				Created      string `json:"created"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				SubnetsInUse int    `json:"subnets_in_use"`
				Updated      string `json:"updated"`
				URI          string `json:"uri"`
			} `json:"allocation"`
			Cloud           bool   `json:"cloud"`
			Created         string `json:"created"`
			Gateway         string `json:"gateway"`
			ID              int    `json:"id"`
			IpsInUse        int    `json:"ips_in_use"`
			ModifiedBy      int    `json:"modified_by"`
			Name            string `json:"name"`
			SubnetsInUse    int    `json:"subnets_in_use"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
			Vlan            int    `json:"vlan"`
			Vxlan           int    `json:"vxlan"`
		} `json:"children"`
		Cloud      bool   `json:"cloud"`
		Created    string `json:"created"`
		Gateway    string `json:"gateway"`
		ID         int    `json:"id"`
		IpsInUse   int    `json:"ips_in_use"`
		ModifiedBy int    `json:"modified_by"`
		Name       string `json:"name"`
		Parent     struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Allocation   struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Asn          struct {
					AllocationsInUse int    `json:"allocations_in_use"`
					Created          string `json:"created"`
					ID               int    `json:"id"`
					MemberID         int    `json:"member_id"`
					Number           int    `json:"number"`
					Updated          string `json:"updated"`
					URI              string `json:"uri"`
				} `json:"asn"`
				Created      string `json:"created"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				SubnetsInUse int    `json:"subnets_in_use"`
				Updated      string `json:"updated"`
				URI          string `json:"uri"`
			} `json:"allocation"`
			Cloud           bool   `json:"cloud"`
			Created         string `json:"created"`
			Gateway         string `json:"gateway"`
			ID              int    `json:"id"`
			IpsInUse        int    `json:"ips_in_use"`
			ModifiedBy      int    `json:"modified_by"`
			Name            string `json:"name"`
			SubnetsInUse    int    `json:"subnets_in_use"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
			Vlan            int    `json:"vlan"`
			Vxlan           int    `json:"vxlan"`
		} `json:"parent"`
		SubnetsInUse    int    `json:"subnets_in_use"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
		VirtualRouterID int    `json:"virtual_router_id"`
		Vlan            int    `json:"vlan"`
		Vxlan           int    `json:"vxlan"`
	} `json:"content"`
}

type VirtualRouter struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Created       string `json:"created"`
		FirewallRules []struct {
			Allow           bool   `json:"allow"`
			Created         string `json:"created"`
			DebugLogging    bool   `json:"debug_logging"`
			Description     string `json:"description"`
			Destination     string `json:"destination"`
			ID              int    `json:"id"`
			Order           int    `json:"order"`
			PciLogging      bool   `json:"pci_logging"`
			Port            int    `json:"port"`
			Protocol        string `json:"protocol"`
			Source          string `json:"source"`
			Updated         string `json:"updated"`
			VirtualRouterID int    `json:"virtual_router_id"`
		} `json:"firewall_rules"`
		ID        int `json:"id"`
		IPAddress struct {
			Address     string `json:"address"`
			Cloud       bool   `json:"cloud"`
			Created     string `json:"created"`
			Credentials string `json:"credentials"`
			ID          int    `json:"id"`
			Location    string `json:"location"`
			ModifiedBy  int    `json:"modified_by"`
			Name        string `json:"name"`
			PublicIP    struct {
				Address string `json:"address"`
				ID      int    `json:"id"`
			} `json:"public_ip"`
			Subnet struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Children []struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"children"`
				Cloud      bool   `json:"cloud"`
				Created    string `json:"created"`
				Gateway    string `json:"gateway"`
				ID         int    `json:"id"`
				IpsInUse   int    `json:"ips_in_use"`
				ModifiedBy int    `json:"modified_by"`
				Name       string `json:"name"`
				Parent     struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"parent"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"subnet"`
			Updated string `json:"updated"`
			VMID    int    `json:"vm_id"`
		} `json:"ip_address"`
		Project struct {
			AddressID       int    `json:"address_id"`
			Closed          bool   `json:"closed"`
			CloudURI        string `json:"cloud_uri"`
			Created         string `json:"created"`
			ID              int    `json:"id"`
			ManagerID       int    `json:"manager_id"`
			MinState        int    `json:"min_state"`
			Name            string `json:"name"`
			RegionID        int    `json:"region_id"`
			ShutDown        bool   `json:"shut_down"`
			Stable          bool   `json:"stable"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
		} `json:"project"`
		RouterID int `json:"router_id"`
		State    int `json:"state"`
		Subnets  []struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			ID           int    `json:"id"`
			ModifiedBy   int    `json:"modified_by"`
			Name         string `json:"name"`
			Vlan         int    `json:"vlan"`
			Vxlan        int    `json:"vxlan"`
		} `json:"subnets"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type VirtualRouterSpecific struct {
	Content struct {
		Created       string `json:"created"`
		FirewallRules []struct {
			Allow           bool   `json:"allow"`
			Created         string `json:"created"`
			DebugLogging    bool   `json:"debug_logging"`
			Description     string `json:"description"`
			Destination     string `json:"destination"`
			ID              int    `json:"id"`
			Order           int    `json:"order"`
			PciLogging      bool   `json:"pci_logging"`
			Port            int    `json:"port"`
			Protocol        string `json:"protocol"`
			Source          string `json:"source"`
			Updated         string `json:"updated"`
			VirtualRouterID int    `json:"virtual_router_id"`
		} `json:"firewall_rules"`
		ID        int `json:"id"`
		IPAddress struct {
			Address     string `json:"address"`
			Cloud       bool   `json:"cloud"`
			Created     string `json:"created"`
			Credentials string `json:"credentials"`
			ID          int    `json:"id"`
			Location    string `json:"location"`
			ModifiedBy  int    `json:"modified_by"`
			Name        string `json:"name"`
			PublicIP    struct {
				Address string `json:"address"`
				ID      int    `json:"id"`
			} `json:"public_ip"`
			Subnet struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Children []struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"children"`
				Cloud      bool   `json:"cloud"`
				Created    string `json:"created"`
				Gateway    string `json:"gateway"`
				ID         int    `json:"id"`
				IpsInUse   int    `json:"ips_in_use"`
				ModifiedBy int    `json:"modified_by"`
				Name       string `json:"name"`
				Parent     struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"parent"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"subnet"`
			Updated string `json:"updated"`
			VMID    int    `json:"vm_id"`
		} `json:"ip_address"`
		Project struct {
			AddressID       int    `json:"address_id"`
			Closed          bool   `json:"closed"`
			CloudURI        string `json:"cloud_uri"`
			Created         string `json:"created"`
			ID              int    `json:"id"`
			ManagerID       int    `json:"manager_id"`
			MinState        int    `json:"min_state"`
			Name            string `json:"name"`
			RegionID        int    `json:"region_id"`
			ShutDown        bool   `json:"shut_down"`
			Stable          bool   `json:"stable"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
		} `json:"project"`
		RouterID int `json:"router_id"`
		State    int `json:"state"`
		Subnets  []struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			ID           int    `json:"id"`
			ModifiedBy   int    `json:"modified_by"`
			Name         string `json:"name"`
			Vlan         int    `json:"vlan"`
			Vxlan        int    `json:"vxlan"`
		} `json:"subnets"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type Vm struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		CPU           int      `json:"cpu"`
		Created       string   `json:"created"`
		DNS           string   `json:"dns"`
		Emails        []string `json:"emails"`
		GatewaySubnet struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Allocation   struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Asn          struct {
					AllocationsInUse int    `json:"allocations_in_use"`
					Created          string `json:"created"`
					ID               int    `json:"id"`
					MemberID         int    `json:"member_id"`
					Number           int    `json:"number"`
					Updated          string `json:"updated"`
					URI              string `json:"uri"`
				} `json:"asn"`
				Created      string `json:"created"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				SubnetsInUse int    `json:"subnets_in_use"`
				Updated      string `json:"updated"`
				URI          string `json:"uri"`
			} `json:"allocation"`
			Children []struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"children"`
			Cloud      bool   `json:"cloud"`
			Created    string `json:"created"`
			Gateway    string `json:"gateway"`
			ID         int    `json:"id"`
			IpsInUse   int    `json:"ips_in_use"`
			ModifiedBy int    `json:"modified_by"`
			Name       string `json:"name"`
			Parent     struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"parent"`
			SubnetsInUse    int    `json:"subnets_in_use"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
			Vlan            int    `json:"vlan"`
			Vxlan           int    `json:"vxlan"`
		} `json:"gateway_subnet"`
		History struct {
			CPUQuantity      int    `json:"cpu_quantity"`
			CPUSku           string `json:"cpu_sku"`
			Created          string `json:"created"`
			CustomerAddress  int    `json:"customer_address"`
			ImageQuantity    int    `json:"image_quantity"`
			ImageSku         string `json:"image_sku"`
			ModifiedBy       int    `json:"modified_by"`
			NatQuantity      int    `json:"nat_quantity"`
			NatSku           string `json:"nat_sku"`
			ProjectID        int    `json:"project_id"`
			ProjectVMName    string `json:"project_vm_name"`
			RAMQuantity      int    `json:"ram_quantity"`
			RAMSku           string `json:"ram_sku"`
			State            int    `json:"state"`
			StorageHistories []struct {
				GbQuantity  int    `json:"gb_quantity"`
				GbSku       string `json:"gb_sku"`
				StorageID   int    `json:"storage_id"`
				StorageName string `json:"storage_name"`
			} `json:"storage_histories"`
			VMID int `json:"vm_id"`
		} `json:"history"`
		ID    int `json:"id"`
		Image struct {
			AnswerFileName string `json:"answer_file_name"`
			Created        string `json:"created"`
			DisplayName    string `json:"display_name"`
			Filename       string `json:"filename"`
			ID             int    `json:"id"`
			MultipleIps    bool   `json:"multiple_ips"`
			OsVariant      string `json:"os_variant"`
			Regions        []int  `json:"regions"`
			Updated        string `json:"updated"`
			URI            string `json:"uri"`
		} `json:"image"`
		IPAddresses []struct {
			Address     string `json:"address"`
			Cloud       bool   `json:"cloud"`
			Created     string `json:"created"`
			Credentials string `json:"credentials"`
			ID          int    `json:"id"`
			Location    string `json:"location"`
			ModifiedBy  int    `json:"modified_by"`
			Name        string `json:"name"`
			PublicIP    struct {
				Address string `json:"address"`
				ID      int    `json:"id"`
			} `json:"public_ip"`
			Subnet struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Children []struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"children"`
				Cloud      bool   `json:"cloud"`
				Created    string `json:"created"`
				Gateway    string `json:"gateway"`
				ID         int    `json:"id"`
				IpsInUse   int    `json:"ips_in_use"`
				ModifiedBy int    `json:"modified_by"`
				Name       string `json:"name"`
				Parent     struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"parent"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"subnet"`
			Updated string `json:"updated"`
			VMID    int    `json:"vm_id"`
		} `json:"ip_addresses"`
		Name    string `json:"name"`
		Project struct {
			AddressID       int    `json:"address_id"`
			Closed          bool   `json:"closed"`
			CloudURI        string `json:"cloud_uri"`
			Created         string `json:"created"`
			ID              int    `json:"id"`
			ManagerID       int    `json:"manager_id"`
			MinState        int    `json:"min_state"`
			Name            string `json:"name"`
			RegionID        int    `json:"region_id"`
			ShutDown        bool   `json:"shut_down"`
			Stable          bool   `json:"stable"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
		} `json:"project"`
		RAM         int    `json:"ram"`
		ServerID    int    `json:"server_id"`
		State       int    `json:"state"`
		StorageType string `json:"storage_type"`
		Storages    []struct {
			Created string `json:"created"`
			Gb      int    `json:"gb"`
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Primary bool   `json:"primary"`
			Updated string `json:"updated"`
			URI     string `json:"uri"`
			VMID    int    `json:"vm_id"`
		} `json:"storages"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type VmSpecific struct {
	Content struct {
		CPU           int      `json:"cpu"`
		Created       string   `json:"created"`
		DNS           string   `json:"dns"`
		Emails        []string `json:"emails"`
		GatewaySubnet struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Allocation   struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Asn          struct {
					AllocationsInUse int    `json:"allocations_in_use"`
					Created          string `json:"created"`
					ID               int    `json:"id"`
					MemberID         int    `json:"member_id"`
					Number           int    `json:"number"`
					Updated          string `json:"updated"`
					URI              string `json:"uri"`
				} `json:"asn"`
				Created      string `json:"created"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				SubnetsInUse int    `json:"subnets_in_use"`
				Updated      string `json:"updated"`
				URI          string `json:"uri"`
			} `json:"allocation"`
			Children []struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"children"`
			Cloud      bool   `json:"cloud"`
			Created    string `json:"created"`
			Gateway    string `json:"gateway"`
			ID         int    `json:"id"`
			IpsInUse   int    `json:"ips_in_use"`
			ModifiedBy int    `json:"modified_by"`
			Name       string `json:"name"`
			Parent     struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"parent"`
			SubnetsInUse    int    `json:"subnets_in_use"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
			Vlan            int    `json:"vlan"`
			Vxlan           int    `json:"vxlan"`
		} `json:"gateway_subnet"`
		History struct {
			CPUQuantity      int    `json:"cpu_quantity"`
			CPUSku           string `json:"cpu_sku"`
			Created          string `json:"created"`
			CustomerAddress  int    `json:"customer_address"`
			ImageQuantity    int    `json:"image_quantity"`
			ImageSku         string `json:"image_sku"`
			ModifiedBy       int    `json:"modified_by"`
			NatQuantity      int    `json:"nat_quantity"`
			NatSku           string `json:"nat_sku"`
			ProjectID        int    `json:"project_id"`
			ProjectVMName    string `json:"project_vm_name"`
			RAMQuantity      int    `json:"ram_quantity"`
			RAMSku           string `json:"ram_sku"`
			State            int    `json:"state"`
			StorageHistories []struct {
				GbQuantity  int    `json:"gb_quantity"`
				GbSku       string `json:"gb_sku"`
				StorageID   int    `json:"storage_id"`
				StorageName string `json:"storage_name"`
			} `json:"storage_histories"`
			VMID int `json:"vm_id"`
		} `json:"history"`
		ID    int `json:"id"`
		Image struct {
			AnswerFileName string `json:"answer_file_name"`
			Created        string `json:"created"`
			DisplayName    string `json:"display_name"`
			Filename       string `json:"filename"`
			ID             int    `json:"id"`
			MultipleIps    bool   `json:"multiple_ips"`
			OsVariant      string `json:"os_variant"`
			Regions        []int  `json:"regions"`
			Updated        string `json:"updated"`
			URI            string `json:"uri"`
		} `json:"image"`
		IPAddresses []struct {
			Address     string `json:"address"`
			Cloud       bool   `json:"cloud"`
			Created     string `json:"created"`
			Credentials string `json:"credentials"`
			ID          int    `json:"id"`
			Location    string `json:"location"`
			ModifiedBy  int    `json:"modified_by"`
			Name        string `json:"name"`
			PublicIP    struct {
				Address string `json:"address"`
				ID      int    `json:"id"`
			} `json:"public_ip"`
			Subnet struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Children []struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"children"`
				Cloud      bool   `json:"cloud"`
				Created    string `json:"created"`
				Gateway    string `json:"gateway"`
				ID         int    `json:"id"`
				IpsInUse   int    `json:"ips_in_use"`
				ModifiedBy int    `json:"modified_by"`
				Name       string `json:"name"`
				Parent     struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"parent"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"subnet"`
			Updated string `json:"updated"`
			VMID    int    `json:"vm_id"`
		} `json:"ip_addresses"`
		Name    string `json:"name"`
		Project struct {
			AddressID       int    `json:"address_id"`
			Closed          bool   `json:"closed"`
			CloudURI        string `json:"cloud_uri"`
			Created         string `json:"created"`
			ID              int    `json:"id"`
			ManagerID       int    `json:"manager_id"`
			MinState        int    `json:"min_state"`
			Name            string `json:"name"`
			RegionID        int    `json:"region_id"`
			ShutDown        bool   `json:"shut_down"`
			Stable          bool   `json:"stable"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
		} `json:"project"`
		RAM         int    `json:"ram"`
		ServerID    int    `json:"server_id"`
		State       int    `json:"state"`
		StorageType string `json:"storage_type"`
		Storages    []struct {
			Created string `json:"created"`
			Gb      int    `json:"gb"`
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Primary bool   `json:"primary"`
			Updated string `json:"updated"`
			URI     string `json:"uri"`
			VMID    int    `json:"vm_id"`
		} `json:"storages"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}

type Vpn struct {
	Content []struct {
		Created             string `json:"created"`
		Description         string `json:"description"`
		ID                  int    `json:"id"`
		IkeAuthentication   string `json:"ike_authentication"`
		IkeDhGroups         string `json:"ike_dh_groups"`
		IkeEncryption       string `json:"ike_encryption"`
		IkeLifetime         int    `json:"ike_lifetime"`
		IkeMode             string `json:"ike_mode"`
		IkePreSharedKey     string `json:"ike_pre_shared_key"`
		IkePublicIP         string `json:"ike_public_ip"`
		IkeVersion          string `json:"ike_version"`
		IpsecAuthentication string `json:"ipsec_authentication"`
		IpsecEncryption     string `json:"ipsec_encryption"`
		IpsecEstablishTime  string `json:"ipsec_establish_time"`
		IpsecPfsGroups      string `json:"ipsec_pfs_groups"`
		IpsecLifetime       int    `json:"ipsec_lifetime"`
		Routes              []struct {
			Created     string `json:"created"`
			ID          int    `json:"id"`
			LocalSubnet struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int64  `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool          `json:"cloud"`
				Created         string        `json:"created"`
				Gateway         interface{}   `json:"gateway"`
				ID              int           `json:"id"`
				IpsInUse        int           `json:"ips_in_use"`
				ModifiedBy      int           `json:"modified_by"`
				SubnetsInUse    int           `json:"subnets_in_use"`
				Name            string        `json:"name"`
				Updated         string        `json:"updated"`
				URI             string        `json:"uri"`
				VirtualRouterID int           `json:"virtual_router_id"`
				Vlan            int           `json:"vlan"`
				Vxlan           int           `json:"vxlan"`
				Children        []interface{} `json:"children"`
				Parent          interface{}   `json:"parent"`
			} `json:"local_subnet"`
			RemoteSubnet string `json:"remote_subnet"`
			Updated      string `json:"updated"`
		} `json:"routes"`
		SendEmail       bool   `json:"send_email"`
		StifNumber      int    `json:"stif_number"`
		TrafficSelector bool   `json:"traffic_selector"`
		Updated         string `json:"updated"`
		URI             string `json:"uri"`
		VirtualRouterID int    `json:"virtual_router_id"`
	} `json:"content"`
	Metadata struct {
		Page         int           `json:"page"`
		Limit        int           `json:"limit"`
		Order        string        `json:"order"`
		TotalRecords int           `json:"total_records"`
		Warnings     []interface{} `json:"warnings"`
	} `json:"_metadata"`
}

type VpnSpecific struct {
	Content struct {
		CPU           int      `json:"cpu"`
		Created       string   `json:"created"`
		DNS           string   `json:"dns"`
		Emails        []string `json:"emails"`
		GatewaySubnet struct {
			AddressID    int    `json:"address_id"`
			AddressRange string `json:"address_range"`
			Allocation   struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Asn          struct {
					AllocationsInUse int    `json:"allocations_in_use"`
					Created          string `json:"created"`
					ID               int    `json:"id"`
					MemberID         int    `json:"member_id"`
					Number           int    `json:"number"`
					Updated          string `json:"updated"`
					URI              string `json:"uri"`
				} `json:"asn"`
				Created      string `json:"created"`
				ID           int    `json:"id"`
				ModifiedBy   int    `json:"modified_by"`
				Name         string `json:"name"`
				SubnetsInUse int    `json:"subnets_in_use"`
				Updated      string `json:"updated"`
				URI          string `json:"uri"`
			} `json:"allocation"`
			Children []struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"children"`
			Cloud      bool   `json:"cloud"`
			Created    string `json:"created"`
			Gateway    string `json:"gateway"`
			ID         int    `json:"id"`
			IpsInUse   int    `json:"ips_in_use"`
			ModifiedBy int    `json:"modified_by"`
			Name       string `json:"name"`
			Parent     struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Cloud           bool   `json:"cloud"`
				Created         string `json:"created"`
				Gateway         string `json:"gateway"`
				ID              int    `json:"id"`
				IpsInUse        int    `json:"ips_in_use"`
				ModifiedBy      int    `json:"modified_by"`
				Name            string `json:"name"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"parent"`
			SubnetsInUse    int    `json:"subnets_in_use"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
			Vlan            int    `json:"vlan"`
			Vxlan           int    `json:"vxlan"`
		} `json:"gateway_subnet"`
		History struct {
			CPUQuantity      int    `json:"cpu_quantity"`
			CPUSku           string `json:"cpu_sku"`
			Created          string `json:"created"`
			CustomerAddress  int    `json:"customer_address"`
			ImageQuantity    int    `json:"image_quantity"`
			ImageSku         string `json:"image_sku"`
			ModifiedBy       int    `json:"modified_by"`
			NatQuantity      int    `json:"nat_quantity"`
			NatSku           string `json:"nat_sku"`
			ProjectID        int    `json:"project_id"`
			ProjectVMName    string `json:"project_vm_name"`
			RAMQuantity      int    `json:"ram_quantity"`
			RAMSku           string `json:"ram_sku"`
			State            int    `json:"state"`
			StorageHistories []struct {
				GbQuantity  int    `json:"gb_quantity"`
				GbSku       string `json:"gb_sku"`
				StorageID   int    `json:"storage_id"`
				StorageName string `json:"storage_name"`
			} `json:"storage_histories"`
			VMID int `json:"vm_id"`
		} `json:"history"`
		ID    int `json:"id"`
		Image struct {
			AnswerFileName string `json:"answer_file_name"`
			Created        string `json:"created"`
			DisplayName    string `json:"display_name"`
			Filename       string `json:"filename"`
			ID             int    `json:"id"`
			MultipleIps    bool   `json:"multiple_ips"`
			OsVariant      string `json:"os_variant"`
			Regions        []int  `json:"regions"`
			Updated        string `json:"updated"`
			URI            string `json:"uri"`
		} `json:"image"`
		IPAddresses []struct {
			Address     string `json:"address"`
			Cloud       bool   `json:"cloud"`
			Created     string `json:"created"`
			Credentials string `json:"credentials"`
			ID          int    `json:"id"`
			Location    string `json:"location"`
			ModifiedBy  int    `json:"modified_by"`
			Name        string `json:"name"`
			PublicIP    struct {
				Address string `json:"address"`
				ID      int    `json:"id"`
			} `json:"public_ip"`
			Subnet struct {
				AddressID    int    `json:"address_id"`
				AddressRange string `json:"address_range"`
				Allocation   struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Asn          struct {
						AllocationsInUse int    `json:"allocations_in_use"`
						Created          string `json:"created"`
						ID               int    `json:"id"`
						MemberID         int    `json:"member_id"`
						Number           int    `json:"number"`
						Updated          string `json:"updated"`
						URI              string `json:"uri"`
					} `json:"asn"`
					Created      string `json:"created"`
					ID           int    `json:"id"`
					ModifiedBy   int    `json:"modified_by"`
					Name         string `json:"name"`
					SubnetsInUse int    `json:"subnets_in_use"`
					Updated      string `json:"updated"`
					URI          string `json:"uri"`
				} `json:"allocation"`
				Children []struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"children"`
				Cloud      bool   `json:"cloud"`
				Created    string `json:"created"`
				Gateway    string `json:"gateway"`
				ID         int    `json:"id"`
				IpsInUse   int    `json:"ips_in_use"`
				ModifiedBy int    `json:"modified_by"`
				Name       string `json:"name"`
				Parent     struct {
					AddressID    int    `json:"address_id"`
					AddressRange string `json:"address_range"`
					Allocation   struct {
						AddressID    int    `json:"address_id"`
						AddressRange string `json:"address_range"`
						Asn          struct {
							AllocationsInUse int    `json:"allocations_in_use"`
							Created          string `json:"created"`
							ID               int    `json:"id"`
							MemberID         int    `json:"member_id"`
							Number           int    `json:"number"`
							Updated          string `json:"updated"`
							URI              string `json:"uri"`
						} `json:"asn"`
						Created      string `json:"created"`
						ID           int    `json:"id"`
						ModifiedBy   int    `json:"modified_by"`
						Name         string `json:"name"`
						SubnetsInUse int    `json:"subnets_in_use"`
						Updated      string `json:"updated"`
						URI          string `json:"uri"`
					} `json:"allocation"`
					Cloud           bool   `json:"cloud"`
					Created         string `json:"created"`
					Gateway         string `json:"gateway"`
					ID              int    `json:"id"`
					IpsInUse        int    `json:"ips_in_use"`
					ModifiedBy      int    `json:"modified_by"`
					Name            string `json:"name"`
					SubnetsInUse    int    `json:"subnets_in_use"`
					Updated         string `json:"updated"`
					URI             string `json:"uri"`
					VirtualRouterID int    `json:"virtual_router_id"`
					Vlan            int    `json:"vlan"`
					Vxlan           int    `json:"vxlan"`
				} `json:"parent"`
				SubnetsInUse    int    `json:"subnets_in_use"`
				Updated         string `json:"updated"`
				URI             string `json:"uri"`
				VirtualRouterID int    `json:"virtual_router_id"`
				Vlan            int    `json:"vlan"`
				Vxlan           int    `json:"vxlan"`
			} `json:"subnet"`
			Updated string `json:"updated"`
			VMID    int    `json:"vm_id"`
		} `json:"ip_addresses"`
		Name    string `json:"name"`
		Project struct {
			AddressID       int    `json:"address_id"`
			Closed          bool   `json:"closed"`
			CloudURI        string `json:"cloud_uri"`
			Created         string `json:"created"`
			ID              int    `json:"id"`
			ManagerID       int    `json:"manager_id"`
			MinState        int    `json:"min_state"`
			Name            string `json:"name"`
			RegionID        int    `json:"region_id"`
			ShutDown        bool   `json:"shut_down"`
			Stable          bool   `json:"stable"`
			Updated         string `json:"updated"`
			URI             string `json:"uri"`
			VirtualRouterID int    `json:"virtual_router_id"`
		} `json:"project"`
		RAM         int    `json:"ram"`
		ServerID    int    `json:"server_id"`
		State       int    `json:"state"`
		StorageType string `json:"storage_type"`
		Storages    []struct {
			Created string `json:"created"`
			Gb      int    `json:"gb"`
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Primary bool   `json:"primary"`
			Updated string `json:"updated"`
			URI     string `json:"uri"`
			VMID    int    `json:"vm_id"`
		} `json:"storages"`
		Updated string `json:"updated"`
		URI     string `json:"uri"`
	} `json:"content"`
}
