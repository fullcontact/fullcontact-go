package fullcontact

//Permission Request

type PermissionRequestOption func(ar *PermissionRequest)

type PermissionMultifield struct {
	Emails     			[]string    		`json:"emails,omitempty"`
	Phones     			[]string    		`json:"phones,omitempty"`
	Maids      			[]string    		`json:"maids,omitempty"`
	Location   			*Location   		`json:"location,omitempty"`
	Name       			*PersonName 		`json:"name,omitempty"`
	Profiles   			[]*Profile  		`json:"profiles,omitempty"`
	LiNonId    			string      		`json:"li_nonid,omitempty"`
}

type PermissionRequest struct {
	Timestamp  			int							`json:"timestamp,omitempty"`
	Query    			PermissionMultifield		`json:"query,omitempty"`
	ConsentPurposes		[]*ConsentPurposes 			`json:"consentPurposes,omitempty"`
	Locale     			string      				`json:"locale,omitempty"`
	IpAddress			string						`json:"ipAddress,omitempty"`
	Language     		string      				`json:"language,omitempty"`
	CollectionMethod	string						`json:"collectionMethod,omitempty"`
	CollectionLocation 	string      				`json:"collectionLocation,omitempty"`
	Tcf					string						`json:"tcf,omitempty"`
	PolicyUrl     		string      				`json:"policyUrl,omitempty"`
	TermsService		string						`json:"termsService,omitempty"`
	PurposeId			int							`json:"purposeId,omitempty"`
	Channel				string						`json:"channel,omitempty"`
}

func NewPermissionRequest(option ...PermissionRequestOption) (*PermissionRequest, error) {
	permissionRequest := &PermissionRequest{}

	for _, opt := range option {
		print(option)
		opt(permissionRequest)
	}
	return permissionRequest, nil
}

func (permissionRequest *PermissionRequest) isQueryable() bool {
	return permissionRequest.Query.Emails != nil ||
		permissionRequest.Query.Phones != nil ||
		permissionRequest.Query.Profiles != nil ||
		permissionRequest.Query.Maids != nil ||
		isPopulated(permissionRequest.Query.LiNonId)
}

func validatePermissionMultifieldRequest(permissionRequest *PermissionRequest) error {
	if !permissionRequest.isQueryable() {
		if permissionRequest.Query.Location == nil && permissionRequest.Query.Name == nil {
			return nil
		} else if permissionRequest.Query.Location != nil && permissionRequest.Query.Name != nil {
			// Validating Location fields
			if isPopulated(permissionRequest.Query.Location.AddressLine1) &&
				((isPopulated(permissionRequest.Query.Location.City) &&
					(isPopulated(permissionRequest.Query.Location.Region) || isPopulated(permissionRequest.Query.Location.RegionCode))) ||
					(isPopulated(permissionRequest.Query.Location.PostalCode))) {
				// Validating Name fields
				if (isPopulated(permissionRequest.Query.Name.Full)) ||
					(isPopulated(permissionRequest.Query.Name.Given) && isPopulated(permissionRequest.Query.Name.Family)) {
					return nil
				} else {
					return NewFullContactError("Name data requires full name or given and family name")
				}
			} else {
				return NewFullContactError(
					"Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
			}
		}
		return NewFullContactError(
			"If you want to use 'location' or 'name' as an input, both must be present and they must have non-blank values")
	}
	return nil
}

func validateForPermissionCreateFields(request *PermissionRequest) error {
	if request.ConsentPurposes == nil {
		return NewFullContactError("At least 1 `consentPurpose` is Required for PermissionRequest")
	} else if !isPopulated(request.CollectionMethod) {
		return NewFullContactError("Collection Method is required for PermissionRequest")
	} else if !isPopulated(request.CollectionLocation) {
		return NewFullContactError("Collection Location is required for PermissionRequest")
	} else if !isPopulated(request.PolicyUrl) {
		return NewFullContactError("Policy URL is required for PermissionRequest")
	} else if !isPopulated(request.TermsService) {
		return NewFullContactError("Terms of Service is required for PermissionRequest")
	}
	return nil
}

func validateForPermissionVerifyFields(request *PermissionRequest) error {
	if request.PurposeId == 0 {
		return NewFullContactError("Purpose ID is required for PermissionRequest")
	} else if !isPopulated(request.Channel) {
		return NewFullContactError("Channel is required for PermissionRequest")
	}
	return nil
}

func validateForPermissionCreate(request *PermissionRequest) error {
	err := validateForPermissionCreateFields(request)
	if err != nil {
		return err
	}
	err = validatePermissionMultifieldRequest(request)
	return err
}

func validateForPermissionDelete(request *PermissionRequest) error {
	err := validatePermissionMultifieldRequest(request)
	return err
}

func validateForPermissionFind(request *PermissionRequest) error {
	err := validatePermissionMultifieldRequest(request)
	return err
}

func validateForPermissionCurrent(request *PermissionRequest) error {
	err := validatePermissionMultifieldRequest(request)
	return err
}

func validateForPermissionVerify(request *PermissionRequest) error {
	err := validateForPermissionVerifyFields(request)
	if err != nil {
		return err
	}
	err = validatePermissionMultifieldRequest(request)
	return err
}

func WithMaidsForPermission(maid string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Query.Maids == nil {
			pr.Query.Maids = make([]string, 0)
		}
		pr.Query.Maids = append(pr.Query.Maids, maid)
	}
}

func WithEmailForPermission(email string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Query.Emails == nil {
			pr.Query.Emails = make([]string, 0)
		}
		pr.Query.Emails = append(pr.Query.Emails, email)
	}
}

func WithEmailsForPermission(emails []string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Query.Emails == nil {
			pr.Query.Emails = make([]string, 0)
		}
		pr.Query.Emails = append(pr.Query.Emails, emails...)
	}
}

func WithPhoneForPermission(phone string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Query.Phones == nil {
			pr.Query.Phones = make([]string, 0)
		}
		pr.Query.Phones = append(pr.Query.Phones, phone)
	}
}

func WithPhonesForPermission(phones []string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Query.Phones == nil {
			pr.Query.Phones = make([]string, 0)
		}
		pr.Query.Phones = append(pr.Query.Phones, phones...)
	}
}

func WithProfileForPermission(profile *Profile) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Query.Profiles == nil {
			pr.Query.Profiles = make([]*Profile, 0)
		}
		pr.Query.Profiles = append(pr.Query.Profiles, profile)
	}
}

func WithProfilesForPermission(profiles []*Profile) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Query.Profiles == nil {
			pr.Query.Profiles = make([]*Profile, 0)
		}
		pr.Query.Profiles = append(pr.Query.Profiles, profiles...)
	}
}

func WithNameForPermission(name *PersonName) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Query.Name = name
	}
}

func WithLocationForPermission(location *Location) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Query.Location = location
	}
}

func WithLiNonIdForPermission(liNonId string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Query.LiNonId = liNonId
	}
}

func WithConsentPurposeForPermission(consentPurpose *ConsentPurposes) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.ConsentPurposes == nil {
			pr.ConsentPurposes = make([]*ConsentPurposes, 0)
		}
		pr.ConsentPurposes = append(pr.ConsentPurposes, consentPurpose)
	}
}

func WithConsentPurposesForPermission(consentPurpose []*ConsentPurposes) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.ConsentPurposes == nil {
			pr.ConsentPurposes = make([]*ConsentPurposes, 0)
		}
		pr.ConsentPurposes = append(pr.ConsentPurposes, consentPurpose...)
	}
}

func WithLocaleForPermission(locale string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Locale = locale
	}
}

func WithIpAddressForPermission(ipAddress string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.IpAddress = ipAddress
	}
}

func WithLanguageForPermission(language string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Language = language
	}
}

func WithCollectionMethodForPermission(collectionMethod string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.CollectionMethod = collectionMethod
	}
}

func WithCollectionLocationForPermission(collectionLocation string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.CollectionLocation = collectionLocation
	}
}

func WithTcfForPermission(tcf string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Tcf = tcf
	}
}

func WithPolicyUrlForPermission(policyUrl string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.PolicyUrl = policyUrl
	}
}

func WithTermsServiceForPermission(termsService string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.TermsService = termsService
	}
}

func WithPurposeIdForPermission(purposeId int) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.PurposeId = purposeId
	}
}

func WithChannelForPermission(channel string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Channel = channel
	}
}
