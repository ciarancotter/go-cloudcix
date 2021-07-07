package api

type Address struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Address1         string `json:"address1"`
		Address2         string `json:"address2"`
		Address3         string `json:"address3"`
		BillingAddressID int    `json:"billing_address_id"`
		City             string `json:"city"`
		CloudRegion      bool   `json:"cloud_region"`
		Country          struct {
			Alpha2Code       string `json:"alpha_2_code"`
			Alpha3Code       string `json:"alpha_3_code"`
			EnglishName      string `json:"english_name"`
			ID               int    `json:"id"`
			PhonePrefix      string `json:"phone_prefix"`
			PrimaryLevelName string `json:"primary_level_name"`
			URI              string `json:"uri"`
		} `json:"country"`
		Currency struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
			URI    string `json:"uri"`
		} `json:"currency"`
		Email       string `json:"email"`
		FullAddress string `json:"full_address"`
		Gln         string `json:"gln"`
		ID          int    `json:"id"`
		Language    struct {
			Code        string `json:"code"`
			EnglishName string `json:"english_name"`
			ID          int    `json:"id"`
			NativeName  string `json:"native_name"`
			URI         string `json:"uri"`
		} `json:"language"`
		Link struct {
			AddressID       int    `json:"address_id"`
			Client          bool   `json:"client"`
			Compute         bool   `json:"compute"`
			ContraAddressID int    `json:"contra_address_id"`
			CreditLimit     string `json:"credit_limit"`
			Customer        bool   `json:"customer"`
			Note            string `json:"note"`
			Reference       string `json:"reference"`
			ServiceCentre   bool   `json:"service_centre"`
			Supplier        bool   `json:"supplier"`
			Territory       struct {
				ID     int `json:"id"`
				Member struct {
					APIKey   string `json:"api_key"`
					Currency struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Symbol string `json:"symbol"`
						URI    string `json:"uri"`
					} `json:"currency"`
					GlnPrefix   string `json:"gln_prefix"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Secret      bool   `json:"secret"`
					SelfManaged bool   `json:"self_managed"`
					URI         string `json:"uri"`
				} `json:"member"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"territory"`
			URI       string `json:"uri"`
			Warrantor bool   `json:"warrantor"`
		} `json:"link"`
		Linked bool `json:"linked"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name   string `json:"name"`
		Phones []struct {
			Name   string `json:"name"`
			Number string `json:"number"`
		} `json:"phones"`
		Postcode    string `json:"postcode"`
		Subdivision struct {
			AlphaCode string `json:"alpha_code"`
			Country   struct {
				Alpha2Code       string `json:"alpha_2_code"`
				Alpha3Code       string `json:"alpha_3_code"`
				EnglishName      string `json:"english_name"`
				ID               int    `json:"id"`
				PhonePrefix      string `json:"phone_prefix"`
				PrimaryLevelName string `json:"primary_level_name"`
				URI              string `json:"uri"`
			} `json:"country"`
			EnglishName string `json:"english_name"`
			ID          int    `json:"id"`
			URI         string `json:"uri"`
		} `json:"subdivision"`
		URI       string `json:"uri"`
		VatNumber string `json:"vat_number"`
		Website   string `json:"website"`
	} `json:"content"`
}

type AddressSpecific struct {
	Content struct {
		Address1         string `json:"address1"`
		Address2         string `json:"address2"`
		Address3         string `json:"address3"`
		BillingAddressID int    `json:"billing_address_id"`
		City             string `json:"city"`
		CloudRegion      bool   `json:"cloud_region"`
		Country          struct {
			Alpha2Code       string `json:"alpha_2_code"`
			Alpha3Code       string `json:"alpha_3_code"`
			EnglishName      string `json:"english_name"`
			ID               int    `json:"id"`
			PhonePrefix      string `json:"phone_prefix"`
			PrimaryLevelName string `json:"primary_level_name"`
			URI              string `json:"uri"`
		} `json:"country"`
		Currency struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
			URI    string `json:"uri"`
		} `json:"currency"`
		Email       string `json:"email"`
		FullAddress string `json:"full_address"`
		Gln         string `json:"gln"`
		ID          int    `json:"id"`
		Language    struct {
			Code        string `json:"code"`
			EnglishName string `json:"english_name"`
			ID          int    `json:"id"`
			NativeName  string `json:"native_name"`
			URI         string `json:"uri"`
		} `json:"language"`
		Link struct {
			AddressID       int    `json:"address_id"`
			Client          bool   `json:"client"`
			Compute         bool   `json:"compute"`
			ContraAddressID int    `json:"contra_address_id"`
			CreditLimit     string `json:"credit_limit"`
			Customer        bool   `json:"customer"`
			Note            string `json:"note"`
			Reference       string `json:"reference"`
			ServiceCentre   bool   `json:"service_centre"`
			Supplier        bool   `json:"supplier"`
			Territory       struct {
				ID     int `json:"id"`
				Member struct {
					APIKey   string `json:"api_key"`
					Currency struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Symbol string `json:"symbol"`
						URI    string `json:"uri"`
					} `json:"currency"`
					GlnPrefix   string `json:"gln_prefix"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Secret      bool   `json:"secret"`
					SelfManaged bool   `json:"self_managed"`
					URI         string `json:"uri"`
				} `json:"member"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"territory"`
			URI       string `json:"uri"`
			Warrantor bool   `json:"warrantor"`
		} `json:"link"`
		Linked bool `json:"linked"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name   string `json:"name"`
		Phones []struct {
			Name   string `json:"name"`
			Number string `json:"number"`
		} `json:"phones"`
		Postcode    string `json:"postcode"`
		Subdivision struct {
			AlphaCode string `json:"alpha_code"`
			Country   struct {
				Alpha2Code       string `json:"alpha_2_code"`
				Alpha3Code       string `json:"alpha_3_code"`
				EnglishName      string `json:"english_name"`
				ID               int    `json:"id"`
				PhonePrefix      string `json:"phone_prefix"`
				PrimaryLevelName string `json:"primary_level_name"`
				URI              string `json:"uri"`
			} `json:"country"`
			EnglishName string `json:"english_name"`
			ID          int    `json:"id"`
			URI         string `json:"uri"`
		} `json:"subdivision"`
		URI       string `json:"uri"`
		VatNumber string `json:"vat_number"`
		Website   string `json:"website"`
	} `json:"content"`
}

type MembershipAppSettings struct {
	Content struct {
		Created        string `json:"created"`
		ID             int    `json:"id"`
		MinioAccessKey string `json:"minio_access_key"`
		MinioSecretKey string `json:"minio_secret_key"`
		MinioURL       string `json:"minio_url"`
		Updated        string `json:"updated"`
		URI            string `json:"uri"`
	} `json:"content"`
}

type Country struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Alpha2Code       string `json:"alpha_2_code"`
		Alpha3Code       string `json:"alpha_3_code"`
		EnglishName      string `json:"english_name"`
		ID               int    `json:"id"`
		PhonePrefix      string `json:"phone_prefix"`
		PrimaryLevelName string `json:"primary_level_name"`
		URI              string `json:"uri"`
	} `json:"content"`
}

type CountrySpecific struct {
	Content struct {
		Alpha2Code       string `json:"alpha_2_code"`
		Alpha3Code       string `json:"alpha_3_code"`
		EnglishName      string `json:"english_name"`
		ID               int    `json:"id"`
		PhonePrefix      string `json:"phone_prefix"`
		PrimaryLevelName string `json:"primary_level_name"`
		URI              string `json:"uri"`
	} `json:"content"`
}

type Currency struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		URI    string `json:"uri"`
	} `json:"content"`
}

type CurrencySpecific struct {
	Content struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
		URI    string `json:"uri"`
	} `json:"content"`
}

type Department struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		ID     int `json:"id"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"content"`
}

type DepartmentSpecific struct {
	Content struct {
		ID     int `json:"id"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"content"`
}

type EmailConfirmed struct {
	Content struct {
		Valid bool `json:"valid"`
	} `json:"content"`
}

type Language struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Code        string `json:"code"`
		EnglishName string `json:"english_name"`
		ID          int    `json:"id"`
		NativeName  string `json:"native_name"`
		URI         string `json:"uri"`
	} `json:"content"`
}

type LanguageSpecific struct {
	Content struct {
		Code        string `json:"code"`
		EnglishName string `json:"english_name"`
		ID          int    `json:"id"`
		NativeName  string `json:"native_name"`
		URI         string `json:"uri"`
	} `json:"content"`
}

type Member struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		APIKey   string `json:"api_key"`
		Currency struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
			URI    string `json:"uri"`
		} `json:"currency"`
		GlnPrefix   string `json:"gln_prefix"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Secret      bool   `json:"secret"`
		SelfManaged bool   `json:"self_managed"`
		URI         string `json:"uri"`
	} `json:"content"`
}

type MemberSpecific struct {
	Content struct {
		APIKey   string `json:"api_key"`
		Currency struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
			URI    string `json:"uri"`
		} `json:"currency"`
		GlnPrefix   string `json:"gln_prefix"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Secret      bool   `json:"secret"`
		SelfManaged bool   `json:"self_managed"`
		URI         string `json:"uri"`
	} `json:"content"`
}

type Profile struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		ID     int `json:"id"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"content"`
}

type ProfileSpecific struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		ID     int `json:"id"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"content"`
}

type Team struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		ID     int `json:"id"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name  string `json:"name"`
		URI   string `json:"uri"`
		Users struct {
			Address struct {
				Address1         string `json:"address1"`
				Address2         string `json:"address2"`
				Address3         string `json:"address3"`
				BillingAddressID int    `json:"billing_address_id"`
				City             string `json:"city"`
				CloudRegion      bool   `json:"cloud_region"`
				Country          struct {
					Alpha2Code       string `json:"alpha_2_code"`
					Alpha3Code       string `json:"alpha_3_code"`
					EnglishName      string `json:"english_name"`
					ID               int    `json:"id"`
					PhonePrefix      string `json:"phone_prefix"`
					PrimaryLevelName string `json:"primary_level_name"`
					URI              string `json:"uri"`
				} `json:"country"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				Email       string `json:"email"`
				FullAddress string `json:"full_address"`
				Gln         string `json:"gln"`
				ID          int    `json:"id"`
				Language    struct {
					Code        string `json:"code"`
					EnglishName string `json:"english_name"`
					ID          int    `json:"id"`
					NativeName  string `json:"native_name"`
					URI         string `json:"uri"`
				} `json:"language"`
				Link struct {
					AddressID       int    `json:"address_id"`
					Client          bool   `json:"client"`
					Compute         bool   `json:"compute"`
					ContraAddressID int    `json:"contra_address_id"`
					CreditLimit     string `json:"credit_limit"`
					Customer        bool   `json:"customer"`
					Note            string `json:"note"`
					Reference       string `json:"reference"`
					ServiceCentre   bool   `json:"service_centre"`
					Supplier        bool   `json:"supplier"`
					Territory       struct {
						ID     int `json:"id"`
						Member struct {
							APIKey   string `json:"api_key"`
							Currency struct {
								ID     int    `json:"id"`
								Name   string `json:"name"`
								Symbol string `json:"symbol"`
								URI    string `json:"uri"`
							} `json:"currency"`
							GlnPrefix   string `json:"gln_prefix"`
							ID          int    `json:"id"`
							Name        string `json:"name"`
							Secret      bool   `json:"secret"`
							SelfManaged bool   `json:"self_managed"`
							URI         string `json:"uri"`
						} `json:"member"`
						Name string `json:"name"`
						URI  string `json:"uri"`
					} `json:"territory"`
					URI       string `json:"uri"`
					Warrantor bool   `json:"warrantor"`
				} `json:"link"`
				Linked bool `json:"linked"`
				Member struct {
					APIKey   string `json:"api_key"`
					Currency struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Symbol string `json:"symbol"`
						URI    string `json:"uri"`
					} `json:"currency"`
					GlnPrefix   string `json:"gln_prefix"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Secret      bool   `json:"secret"`
					SelfManaged bool   `json:"self_managed"`
					URI         string `json:"uri"`
				} `json:"member"`
				Name   string `json:"name"`
				Phones []struct {
					Name   string `json:"name"`
					Number string `json:"number"`
				} `json:"phones"`
				Postcode    string `json:"postcode"`
				Subdivision struct {
					AlphaCode string `json:"alpha_code"`
					Country   struct {
						Alpha2Code       string `json:"alpha_2_code"`
						Alpha3Code       string `json:"alpha_3_code"`
						EnglishName      string `json:"english_name"`
						ID               int    `json:"id"`
						PhonePrefix      string `json:"phone_prefix"`
						PrimaryLevelName string `json:"primary_level_name"`
						URI              string `json:"uri"`
					} `json:"country"`
					EnglishName string `json:"english_name"`
					ID          int    `json:"id"`
					URI         string `json:"uri"`
				} `json:"subdivision"`
				URI       string `json:"uri"`
				VatNumber string `json:"vat_number"`
				Website   string `json:"website"`
			} `json:"address"`
			Administrator bool `json:"administrator"`
			Department    struct {
				ID     int `json:"id"`
				Member struct {
					APIKey   string `json:"api_key"`
					Currency struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Symbol string `json:"symbol"`
						URI    string `json:"uri"`
					} `json:"currency"`
					GlnPrefix   string `json:"gln_prefix"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Secret      bool   `json:"secret"`
					SelfManaged bool   `json:"self_managed"`
					URI         string `json:"uri"`
				} `json:"member"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"department"`
			Email                 string `json:"email"`
			EmailValidated        bool   `json:"email_validated"`
			ExpiryDate            string `json:"expiry_date"`
			ExternalNotifications []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"external_notifications"`
			FirstName             string `json:"first_name"`
			FirstOtp              string `json:"first_otp"`
			GlobalActive          bool   `json:"global_active"`
			GlobalUser            bool   `json:"global_user"`
			ID                    int    `json:"id"`
			Image                 string `json:"image"`
			InternalNotifications []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"internal_notifications"`
			IsPrivate bool   `json:"is_private"`
			JobTitle  string `json:"job_title"`
			Language  struct {
				Code        string `json:"code"`
				EnglishName string `json:"english_name"`
				ID          int    `json:"id"`
				NativeName  string `json:"native_name"`
				URI         string `json:"uri"`
			} `json:"language"`
			LastLogin string `json:"last_login"`
			Member    struct {
				APIKey   string `json:"api_key"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				GlnPrefix   string `json:"gln_prefix"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Secret      bool   `json:"secret"`
				SelfManaged bool   `json:"self_managed"`
				URI         string `json:"uri"`
			} `json:"member"`
			Otp    bool `json:"otp"`
			Phones []struct {
				Name   string `json:"name"`
				Number string `json:"number"`
			} `json:"phones"`
			Profile struct {
				ID     int `json:"id"`
				Member struct {
					APIKey   string `json:"api_key"`
					Currency struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Symbol string `json:"symbol"`
						URI    string `json:"uri"`
					} `json:"currency"`
					GlnPrefix   string `json:"gln_prefix"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Secret      bool   `json:"secret"`
					SelfManaged bool   `json:"self_managed"`
					URI         string `json:"uri"`
				} `json:"member"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"profile"`
			Robot     bool   `json:"robot"`
			Signature string `json:"signature"`
			StartDate string `json:"start_date"`
			Surname   string `json:"surname"`
			Timezone  string `json:"timezone"`
			URI       string `json:"uri"`
		} `json:"users"`
	} `json:"content"`
}

type TeamSpecific struct {
	Content struct {
		ID     int `json:"id"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name  string `json:"name"`
		URI   string `json:"uri"`
		Users struct {
			Address struct {
				Address1         string `json:"address1"`
				Address2         string `json:"address2"`
				Address3         string `json:"address3"`
				BillingAddressID int    `json:"billing_address_id"`
				City             string `json:"city"`
				CloudRegion      bool   `json:"cloud_region"`
				Country          struct {
					Alpha2Code       string `json:"alpha_2_code"`
					Alpha3Code       string `json:"alpha_3_code"`
					EnglishName      string `json:"english_name"`
					ID               int    `json:"id"`
					PhonePrefix      string `json:"phone_prefix"`
					PrimaryLevelName string `json:"primary_level_name"`
					URI              string `json:"uri"`
				} `json:"country"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				Email       string `json:"email"`
				FullAddress string `json:"full_address"`
				Gln         string `json:"gln"`
				ID          int    `json:"id"`
				Language    struct {
					Code        string `json:"code"`
					EnglishName string `json:"english_name"`
					ID          int    `json:"id"`
					NativeName  string `json:"native_name"`
					URI         string `json:"uri"`
				} `json:"language"`
				Link struct {
					AddressID       int    `json:"address_id"`
					Client          bool   `json:"client"`
					Compute         bool   `json:"compute"`
					ContraAddressID int    `json:"contra_address_id"`
					CreditLimit     string `json:"credit_limit"`
					Customer        bool   `json:"customer"`
					Note            string `json:"note"`
					Reference       string `json:"reference"`
					ServiceCentre   bool   `json:"service_centre"`
					Supplier        bool   `json:"supplier"`
					Territory       struct {
						ID     int `json:"id"`
						Member struct {
							APIKey   string `json:"api_key"`
							Currency struct {
								ID     int    `json:"id"`
								Name   string `json:"name"`
								Symbol string `json:"symbol"`
								URI    string `json:"uri"`
							} `json:"currency"`
							GlnPrefix   string `json:"gln_prefix"`
							ID          int    `json:"id"`
							Name        string `json:"name"`
							Secret      bool   `json:"secret"`
							SelfManaged bool   `json:"self_managed"`
							URI         string `json:"uri"`
						} `json:"member"`
						Name string `json:"name"`
						URI  string `json:"uri"`
					} `json:"territory"`
					URI       string `json:"uri"`
					Warrantor bool   `json:"warrantor"`
				} `json:"link"`
				Linked bool `json:"linked"`
				Member struct {
					APIKey   string `json:"api_key"`
					Currency struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Symbol string `json:"symbol"`
						URI    string `json:"uri"`
					} `json:"currency"`
					GlnPrefix   string `json:"gln_prefix"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Secret      bool   `json:"secret"`
					SelfManaged bool   `json:"self_managed"`
					URI         string `json:"uri"`
				} `json:"member"`
				Name   string `json:"name"`
				Phones []struct {
					Name   string `json:"name"`
					Number string `json:"number"`
				} `json:"phones"`
				Postcode    string `json:"postcode"`
				Subdivision struct {
					AlphaCode string `json:"alpha_code"`
					Country   struct {
						Alpha2Code       string `json:"alpha_2_code"`
						Alpha3Code       string `json:"alpha_3_code"`
						EnglishName      string `json:"english_name"`
						ID               int    `json:"id"`
						PhonePrefix      string `json:"phone_prefix"`
						PrimaryLevelName string `json:"primary_level_name"`
						URI              string `json:"uri"`
					} `json:"country"`
					EnglishName string `json:"english_name"`
					ID          int    `json:"id"`
					URI         string `json:"uri"`
				} `json:"subdivision"`
				URI       string `json:"uri"`
				VatNumber string `json:"vat_number"`
				Website   string `json:"website"`
			} `json:"address"`
			Administrator bool `json:"administrator"`
			Department    struct {
				ID     int `json:"id"`
				Member struct {
					APIKey   string `json:"api_key"`
					Currency struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Symbol string `json:"symbol"`
						URI    string `json:"uri"`
					} `json:"currency"`
					GlnPrefix   string `json:"gln_prefix"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Secret      bool   `json:"secret"`
					SelfManaged bool   `json:"self_managed"`
					URI         string `json:"uri"`
				} `json:"member"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"department"`
			Email                 string `json:"email"`
			EmailValidated        bool   `json:"email_validated"`
			ExpiryDate            string `json:"expiry_date"`
			ExternalNotifications []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"external_notifications"`
			FirstName             string `json:"first_name"`
			FirstOtp              string `json:"first_otp"`
			GlobalActive          bool   `json:"global_active"`
			GlobalUser            bool   `json:"global_user"`
			ID                    int    `json:"id"`
			Image                 string `json:"image"`
			InternalNotifications []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"internal_notifications"`
			IsPrivate bool   `json:"is_private"`
			JobTitle  string `json:"job_title"`
			Language  struct {
				Code        string `json:"code"`
				EnglishName string `json:"english_name"`
				ID          int    `json:"id"`
				NativeName  string `json:"native_name"`
				URI         string `json:"uri"`
			} `json:"language"`
			LastLogin string `json:"last_login"`
			Member    struct {
				APIKey   string `json:"api_key"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				GlnPrefix   string `json:"gln_prefix"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Secret      bool   `json:"secret"`
				SelfManaged bool   `json:"self_managed"`
				URI         string `json:"uri"`
			} `json:"member"`
			Otp    bool `json:"otp"`
			Phones []struct {
				Name   string `json:"name"`
				Number string `json:"number"`
			} `json:"phones"`
			Profile struct {
				ID     int `json:"id"`
				Member struct {
					APIKey   string `json:"api_key"`
					Currency struct {
						ID     int    `json:"id"`
						Name   string `json:"name"`
						Symbol string `json:"symbol"`
						URI    string `json:"uri"`
					} `json:"currency"`
					GlnPrefix   string `json:"gln_prefix"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
					Secret      bool   `json:"secret"`
					SelfManaged bool   `json:"self_managed"`
					URI         string `json:"uri"`
				} `json:"member"`
				Name string `json:"name"`
				URI  string `json:"uri"`
			} `json:"profile"`
			Robot     bool   `json:"robot"`
			Signature string `json:"signature"`
			StartDate string `json:"start_date"`
			Surname   string `json:"surname"`
			Timezone  string `json:"timezone"`
			URI       string `json:"uri"`
		} `json:"users"`
	} `json:"content"`
}

type Territory struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		ID     int `json:"id"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"content"`
}

type TerritorySpecific struct {
	Content struct {
		ID     int `json:"id"`
		Member struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"content"`
}

type TransactionType struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"content"`
}

type TransactionTypeSpecific struct {
	Content struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URI  string `json:"uri"`
	} `json:"content"`
}

type User struct {
	Metadata struct {
		Limit        int      `json:"limit"`
		Order        string   `json:"order"`
		Page         int      `json:"page"`
		TotalRecords int      `json:"total_records"`
		Warnings     []string `json:"warnings"`
	} `json:"_metadata"`
	Content []struct {
		Address struct {
			Address1         string `json:"address1"`
			Address2         string `json:"address2"`
			Address3         string `json:"address3"`
			BillingAddressID int    `json:"billing_address_id"`
			City             string `json:"city"`
			CloudRegion      bool   `json:"cloud_region"`
			Country          struct {
				Alpha2Code       string `json:"alpha_2_code"`
				Alpha3Code       string `json:"alpha_3_code"`
				EnglishName      string `json:"english_name"`
				ID               int    `json:"id"`
				PhonePrefix      string `json:"phone_prefix"`
				PrimaryLevelName string `json:"primary_level_name"`
				URI              string `json:"uri"`
			} `json:"country"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			Email       string `json:"email"`
			FullAddress string `json:"full_address"`
			Gln         string `json:"gln"`
			ID          int    `json:"id"`
			Language    struct {
				Code        string `json:"code"`
				EnglishName string `json:"english_name"`
				ID          int    `json:"id"`
				NativeName  string `json:"native_name"`
				URI         string `json:"uri"`
			} `json:"language"`
			Link struct {
				AddressID       int    `json:"address_id"`
				Client          bool   `json:"client"`
				Compute         bool   `json:"compute"`
				ContraAddressID int    `json:"contra_address_id"`
				CreditLimit     string `json:"credit_limit"`
				Customer        bool   `json:"customer"`
				Note            string `json:"note"`
				Reference       string `json:"reference"`
				ServiceCentre   bool   `json:"service_centre"`
				Supplier        bool   `json:"supplier"`
				Territory       struct {
					ID     int `json:"id"`
					Member struct {
						APIKey   string `json:"api_key"`
						Currency struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Symbol string `json:"symbol"`
							URI    string `json:"uri"`
						} `json:"currency"`
						GlnPrefix   string `json:"gln_prefix"`
						ID          int    `json:"id"`
						Name        string `json:"name"`
						Secret      bool   `json:"secret"`
						SelfManaged bool   `json:"self_managed"`
						URI         string `json:"uri"`
					} `json:"member"`
					Name string `json:"name"`
					URI  string `json:"uri"`
				} `json:"territory"`
				URI       string `json:"uri"`
				Warrantor bool   `json:"warrantor"`
			} `json:"link"`
			Linked bool `json:"linked"`
			Member struct {
				APIKey   string `json:"api_key"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				GlnPrefix   string `json:"gln_prefix"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Secret      bool   `json:"secret"`
				SelfManaged bool   `json:"self_managed"`
				URI         string `json:"uri"`
			} `json:"member"`
			Name   string `json:"name"`
			Phones []struct {
				Name   string `json:"name"`
				Number string `json:"number"`
			} `json:"phones"`
			Postcode    string `json:"postcode"`
			Subdivision struct {
				AlphaCode string `json:"alpha_code"`
				Country   struct {
					Alpha2Code       string `json:"alpha_2_code"`
					Alpha3Code       string `json:"alpha_3_code"`
					EnglishName      string `json:"english_name"`
					ID               int    `json:"id"`
					PhonePrefix      string `json:"phone_prefix"`
					PrimaryLevelName string `json:"primary_level_name"`
					URI              string `json:"uri"`
				} `json:"country"`
				EnglishName string `json:"english_name"`
				ID          int    `json:"id"`
				URI         string `json:"uri"`
			} `json:"subdivision"`
			URI       string `json:"uri"`
			VatNumber string `json:"vat_number"`
			Website   string `json:"website"`
		} `json:"address"`
		Administrator bool `json:"administrator"`
		Department    struct {
			ID     int `json:"id"`
			Member struct {
				APIKey   string `json:"api_key"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				GlnPrefix   string `json:"gln_prefix"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Secret      bool   `json:"secret"`
				SelfManaged bool   `json:"self_managed"`
				URI         string `json:"uri"`
			} `json:"member"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"department"`
		Email                 string `json:"email"`
		EmailValidated        bool   `json:"email_validated"`
		ExpiryDate            string `json:"expiry_date"`
		ExternalNotifications []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"external_notifications"`
		FirstName             string `json:"first_name"`
		FirstOtp              string `json:"first_otp"`
		GlobalActive          bool   `json:"global_active"`
		GlobalUser            bool   `json:"global_user"`
		ID                    int    `json:"id"`
		Image                 string `json:"image"`
		InternalNotifications []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"internal_notifications"`
		IsPrivate bool   `json:"is_private"`
		JobTitle  string `json:"job_title"`
		Language  struct {
			Code        string `json:"code"`
			EnglishName string `json:"english_name"`
			ID          int    `json:"id"`
			NativeName  string `json:"native_name"`
			URI         string `json:"uri"`
		} `json:"language"`
		LastLogin string `json:"last_login"`
		Member    struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Otp    bool `json:"otp"`
		Phones []struct {
			Name   string `json:"name"`
			Number string `json:"number"`
		} `json:"phones"`
		Profile struct {
			ID     int `json:"id"`
			Member struct {
				APIKey   string `json:"api_key"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				GlnPrefix   string `json:"gln_prefix"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Secret      bool   `json:"secret"`
				SelfManaged bool   `json:"self_managed"`
				URI         string `json:"uri"`
			} `json:"member"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"profile"`
		Robot     bool   `json:"robot"`
		Signature string `json:"signature"`
		StartDate string `json:"start_date"`
		Surname   string `json:"surname"`
		Timezone  string `json:"timezone"`
		URI       string `json:"uri"`
	} `json:"content"`
}

type UserSpecific struct {
	Content struct {
		Address struct {
			Address1         string `json:"address1"`
			Address2         string `json:"address2"`
			Address3         string `json:"address3"`
			BillingAddressID int    `json:"billing_address_id"`
			City             string `json:"city"`
			CloudRegion      bool   `json:"cloud_region"`
			Country          struct {
				Alpha2Code       string `json:"alpha_2_code"`
				Alpha3Code       string `json:"alpha_3_code"`
				EnglishName      string `json:"english_name"`
				ID               int    `json:"id"`
				PhonePrefix      string `json:"phone_prefix"`
				PrimaryLevelName string `json:"primary_level_name"`
				URI              string `json:"uri"`
			} `json:"country"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			Email       string `json:"email"`
			FullAddress string `json:"full_address"`
			Gln         string `json:"gln"`
			ID          int    `json:"id"`
			Language    struct {
				Code        string `json:"code"`
				EnglishName string `json:"english_name"`
				ID          int    `json:"id"`
				NativeName  string `json:"native_name"`
				URI         string `json:"uri"`
			} `json:"language"`
			Link struct {
				AddressID       int    `json:"address_id"`
				Client          bool   `json:"client"`
				Compute         bool   `json:"compute"`
				ContraAddressID int    `json:"contra_address_id"`
				CreditLimit     string `json:"credit_limit"`
				Customer        bool   `json:"customer"`
				Note            string `json:"note"`
				Reference       string `json:"reference"`
				ServiceCentre   bool   `json:"service_centre"`
				Supplier        bool   `json:"supplier"`
				Territory       struct {
					ID     int `json:"id"`
					Member struct {
						APIKey   string `json:"api_key"`
						Currency struct {
							ID     int    `json:"id"`
							Name   string `json:"name"`
							Symbol string `json:"symbol"`
							URI    string `json:"uri"`
						} `json:"currency"`
						GlnPrefix   string `json:"gln_prefix"`
						ID          int    `json:"id"`
						Name        string `json:"name"`
						Secret      bool   `json:"secret"`
						SelfManaged bool   `json:"self_managed"`
						URI         string `json:"uri"`
					} `json:"member"`
					Name string `json:"name"`
					URI  string `json:"uri"`
				} `json:"territory"`
				URI       string `json:"uri"`
				Warrantor bool   `json:"warrantor"`
			} `json:"link"`
			Linked bool `json:"linked"`
			Member struct {
				APIKey   string `json:"api_key"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				GlnPrefix   string `json:"gln_prefix"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Secret      bool   `json:"secret"`
				SelfManaged bool   `json:"self_managed"`
				URI         string `json:"uri"`
			} `json:"member"`
			Name   string `json:"name"`
			Phones []struct {
				Name   string `json:"name"`
				Number string `json:"number"`
			} `json:"phones"`
			Postcode    string `json:"postcode"`
			Subdivision struct {
				AlphaCode string `json:"alpha_code"`
				Country   struct {
					Alpha2Code       string `json:"alpha_2_code"`
					Alpha3Code       string `json:"alpha_3_code"`
					EnglishName      string `json:"english_name"`
					ID               int    `json:"id"`
					PhonePrefix      string `json:"phone_prefix"`
					PrimaryLevelName string `json:"primary_level_name"`
					URI              string `json:"uri"`
				} `json:"country"`
				EnglishName string `json:"english_name"`
				ID          int    `json:"id"`
				URI         string `json:"uri"`
			} `json:"subdivision"`
			URI       string `json:"uri"`
			VatNumber string `json:"vat_number"`
			Website   string `json:"website"`
		} `json:"address"`
		Administrator bool `json:"administrator"`
		Department    struct {
			ID     int `json:"id"`
			Member struct {
				APIKey   string `json:"api_key"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				GlnPrefix   string `json:"gln_prefix"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Secret      bool   `json:"secret"`
				SelfManaged bool   `json:"self_managed"`
				URI         string `json:"uri"`
			} `json:"member"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"department"`
		Email                 string `json:"email"`
		EmailValidated        bool   `json:"email_validated"`
		ExpiryDate            string `json:"expiry_date"`
		ExternalNotifications []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"external_notifications"`
		FirstName             string `json:"first_name"`
		FirstOtp              string `json:"first_otp"`
		GlobalActive          bool   `json:"global_active"`
		GlobalUser            bool   `json:"global_user"`
		ID                    int    `json:"id"`
		Image                 string `json:"image"`
		InternalNotifications []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"internal_notifications"`
		IsPrivate bool   `json:"is_private"`
		JobTitle  string `json:"job_title"`
		Language  struct {
			Code        string `json:"code"`
			EnglishName string `json:"english_name"`
			ID          int    `json:"id"`
			NativeName  string `json:"native_name"`
			URI         string `json:"uri"`
		} `json:"language"`
		LastLogin string `json:"last_login"`
		Member    struct {
			APIKey   string `json:"api_key"`
			Currency struct {
				ID     int    `json:"id"`
				Name   string `json:"name"`
				Symbol string `json:"symbol"`
				URI    string `json:"uri"`
			} `json:"currency"`
			GlnPrefix   string `json:"gln_prefix"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Secret      bool   `json:"secret"`
			SelfManaged bool   `json:"self_managed"`
			URI         string `json:"uri"`
		} `json:"member"`
		Otp    bool `json:"otp"`
		Phones []struct {
			Name   string `json:"name"`
			Number string `json:"number"`
		} `json:"phones"`
		Profile struct {
			ID     int `json:"id"`
			Member struct {
				APIKey   string `json:"api_key"`
				Currency struct {
					ID     int    `json:"id"`
					Name   string `json:"name"`
					Symbol string `json:"symbol"`
					URI    string `json:"uri"`
				} `json:"currency"`
				GlnPrefix   string `json:"gln_prefix"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Secret      bool   `json:"secret"`
				SelfManaged bool   `json:"self_managed"`
				URI         string `json:"uri"`
			} `json:"member"`
			Name string `json:"name"`
			URI  string `json:"uri"`
		} `json:"profile"`
		Robot     bool   `json:"robot"`
		Signature string `json:"signature"`
		StartDate string `json:"start_date"`
		Surname   string `json:"surname"`
		Timezone  string `json:"timezone"`
		URI       string `json:"uri"`
	} `json:"content"`
}
