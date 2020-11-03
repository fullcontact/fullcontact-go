package fullcontact

type ResolveRequestOption func(pr *ResolveRequest)

type ResolveRequest struct {
	Emails    []string    `json:"emails,omitempty"`
	Phones    []string    `json:"phones,omitempty"`
	Maid      []string    `json:"maids,omitempty"`
	Location  *Location   `json:"location,omitempty"`
	Name      *PersonName `json:"name,omitempty"`
	Profiles  []*Profile  `json:"profiles,omitempty"`
	RecordId  string      `json:"recordId,omitempty"`
	PersonId  string      `json:"personId,omitempty"`
	PartnerId string      `json:"partnerId,omitempty"`
	LiNonId   string      `json:"li_nonid,omitempty"`
}

func NewResolveRequest(option ...ResolveRequestOption) (*ResolveRequest, error) {
	resolveRequest := &ResolveRequest{}

	for _, opt := range option {
		opt(resolveRequest)
	}
	err := validateResolveRequest(resolveRequest)
	if err != nil {
		resolveRequest = nil
	}
	return resolveRequest, err
}

func validateResolveRequest(resolveRequest *ResolveRequest) error {
	if resolveRequest.Location == nil && resolveRequest.Name == nil {
		return nil
	} else if resolveRequest.Location != nil && resolveRequest.Name != nil {
		// Validating Location fields
		if isPopulated(resolveRequest.Location.AddressLine1) &&
			((isPopulated(resolveRequest.Location.City) &&
				(isPopulated(resolveRequest.Location.Region) || isPopulated(resolveRequest.Location.RegionCode))) ||
				(isPopulated(resolveRequest.Location.PostalCode))) {
			// Validating Name fields
			if (isPopulated(resolveRequest.Name.Full)) ||
				(isPopulated(resolveRequest.Name.Given) && isPopulated(resolveRequest.Name.Family)) {
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

func validateForIdentityMap(request *ResolveRequest) error {
	if isPopulated(request.PersonId) {
		return NewFullContactError("Invalid map request, person id must be empty")
	}
	if (request.Name != nil && request.Location != nil) ||
		request.Emails != nil ||
		request.Phones != nil ||
		request.Profiles != nil {
		return nil
	} else {
		return NewFullContactError("Invalid map request, Any of Email, Phone, SocialProfile, Name and Location must be present")
	}
}

func validateForIdentityResolve(request *ResolveRequest) error {
	if isPopulated(request.RecordId) && isPopulated(request.PersonId) {
		return NewFullContactError("Both record id and person id are populated, please select one")
	}
	return nil
}

func validateForIdentityDelete(request *ResolveRequest) error {
	if !isPopulated(request.RecordId) {
		return NewFullContactError("recordId param must be specified")
	}
	return nil
}

func WithEmailForResolve(email string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Emails == nil {
			resolveRequest.Emails = make([]string, 0)
		}
		resolveRequest.Emails = append(resolveRequest.Emails, email)
	}
}

func WithEmailsForResolve(emails []string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Emails == nil {
			resolveRequest.Emails = make([]string, 0)
		}
		resolveRequest.Emails = append(resolveRequest.Emails, emails...)
	}
}

func WithPhoneForResolve(phone string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Phones == nil {
			resolveRequest.Phones = make([]string, 0)
		}
		resolveRequest.Phones = append(resolveRequest.Phones, phone)
	}
}

func WithPhonesForResolve(phones []string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Phones == nil {
			resolveRequest.Phones = make([]string, 0)
		}
		resolveRequest.Phones = append(resolveRequest.Phones, phones...)
	}
}

func WithMaidForResolve(maid string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Maid == nil {
			resolveRequest.Maid = make([]string, 0)
		}
		resolveRequest.Maid = append(resolveRequest.Maid, maid)
	}
}

func WithMaidsForResolve(maids []string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Maid == nil {
			resolveRequest.Maid = make([]string, 0)
		}
		resolveRequest.Maid = append(resolveRequest.Maid, maids...)
	}
}

func WithLocationForResolve(location *Location) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		resolveRequest.Location = location
	}
}

func WithRecordIdForResolve(recordId string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		resolveRequest.RecordId = recordId
	}
}

func WithPersonIdForResolve(personId string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		resolveRequest.PersonId = personId
	}
}

func WithPartnerIdForResolve(partnerId string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		resolveRequest.PartnerId = partnerId
	}
}

func WithLiNonIdForResolve(liNonId string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		resolveRequest.LiNonId = liNonId
	}
}

func WithNameForResolve(name *PersonName) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		resolveRequest.Name = name
	}
}

func WithProfileForResolve(profile *Profile) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Profiles == nil {
			resolveRequest.Profiles = make([]*Profile, 0)
		}
		resolveRequest.Profiles = append(resolveRequest.Profiles, profile)
	}
}

func WithProfilesForResolve(profiles []*Profile) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Profiles == nil {
			resolveRequest.Profiles = make([]*Profile, 0)
		}
		resolveRequest.Profiles = append(resolveRequest.Profiles, profiles...)
	}
}
