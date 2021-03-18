package fullcontact

//Permission Request

type PermissionRequestOption func(ar *PermissionRequest)

type PermissionRequest struct {
	Emails     []string    `json:"emails,omitempty"`
	Phones     []string    `json:"phones,omitempty"`
	Maids      []string    `json:"maids,omitempty"`
	Location   *Location   `json:"location,omitempty"`
	Name       *PersonName `json:"name,omitempty"`
	Profiles   []*Profile  `json:"profiles,omitempty"`
	LiNonId    string      `json:"li_nonid,omitempty"`
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
	return permissionRequest.Emails != nil ||
		permissionRequest.Phones != nil ||
		permissionRequest.Profiles != nil ||
		permissionRequest.Maids != nil ||
		isPopulated(permissionRequest.LiNonId)
}

func validatePermissionRequest(permissionRequest *PermissionRequest) error {
	if !permissionRequest.isQueryable() {
		if permissionRequest.Location == nil && permissionRequest.Name == nil {
			return nil
		} else if permissionRequest.Location != nil && permissionRequest.Name != nil {
			// Validating Location fields
			if isPopulated(permissionRequest.Location.AddressLine1) &&
				((isPopulated(permissionRequest.Location.City) &&
					(isPopulated(permissionRequest.Location.Region) || isPopulated(permissionRequest.Location.RegionCode))) ||
					(isPopulated(permissionRequest.Location.PostalCode))) {
				// Validating Name fields
				if (isPopulated(permissionRequest.Name.Full)) ||
					(isPopulated(permissionRequest.Name.Given) && isPopulated(permissionRequest.Name.Family)) {
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

func validateForPermissionCreate(request *PermissionRequest) error {
	if !request.isQueryable(){
		return NewFullContactError("Invalid map request, Any of Email, Phone, SocialProfile, Name and Location must be present")
	}
	err := validatePermissionRequest(request)
	return err
}

func validateForPermissionDelete(request *PermissionRequest) error {
	if !request.isQueryable(){
		return NewFullContactError("Invalid map request, Any of Email, Phone, SocialProfile, Name and Location must be present")
	}
	err := validatePermissionRequest(request)
	return err
}

func validateForPermissionFind(request *PermissionRequest) error {
	if !request.isQueryable(){
		return NewFullContactError("Invalid map request, Any of Email, Phone, SocialProfile, Name and Location must be present")
	}
	err := validatePermissionRequest(request)
	return err
}

func validateForPermissionCurrent(request *PermissionRequest) error {
	if !request.isQueryable(){
		return NewFullContactError("Invalid map request, Any of Email, Phone, SocialProfile, Name and Location must be present")
	}
	err := validatePermissionRequest(request)
	return err
}

func validateForPermissionVerify(request *PermissionRequest) error {
	if !request.isQueryable(){
		return NewFullContactError("Invalid map request, Any of Email, Phone, SocialProfile, Name and Location must be present")
	}
	err := validatePermissionRequest(request)
	return err
}

func WithMaidsForPermission(maid string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Maids == nil {
			pr.Maids = make([]string, 0)
		}
		pr.Maids = append(pr.Maids, maid)
	}
}

func WithEmailForPermission(email string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Emails == nil {
			pr.Emails = make([]string, 0)
		}
		pr.Emails = append(pr.Emails, email)
	}
}

func WithEmailsForPermission(emails []string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Emails == nil {
			pr.Emails = make([]string, 0)
		}
		pr.Emails = append(pr.Emails, emails...)
	}
}

func WithPhoneForPermission(phone string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Phones == nil {
			pr.Phones = make([]string, 0)
		}
		pr.Phones = append(pr.Phones, phone)
	}
}

func WithPhonesForPermission(phones []string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Phones == nil {
			pr.Phones = make([]string, 0)
		}
		pr.Phones = append(pr.Phones, phones...)
	}
}

func WithProfileForPermission(profile *Profile) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Profiles == nil {
			pr.Profiles = make([]*Profile, 0)
		}
		pr.Profiles = append(pr.Profiles, profile)
	}
}

func WithProfilesForPermission(profiles []*Profile) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.Profiles == nil {
			pr.Profiles = make([]*Profile, 0)
		}
		pr.Profiles = append(pr.Profiles, profiles...)
	}
}

func WithNameForPermission(name *PersonName) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Name = name
	}
}

func WithLocationForPermission(location *Location) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Location = location
	}
}

func WithLiNonIdForPermission(liNonId string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.LiNonId = liNonId
	}
}
