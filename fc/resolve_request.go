package fullcontact

type ResolveRequestOption func(pr *ResolveRequest)

type ResolveRequest struct {
	Emails      []string    `json:"emails,omitempty"`
	Phones      []string    `json:"phones,omitempty"`
	Maid        []string    `json:"maids,omitempty"`
	Location    *Location   `json:"location,omitempty"`
	Name        *PersonName `json:"name,omitempty"`
	Profiles    []*Profile  `json:"profiles,omitempty"`
	RecordId    string      `json:"recordId,omitempty"`
	PersonId    string      `json:"personId,omitempty"`
	PartnerId   string      `json:"partnerId,omitempty"`
	LiNonId     string      `json:"li_nonid,omitempty"`
	Tags        []*Tag      `json:"tags,omitempty"`
	Placekey    string      `json:"placekey,omitempty"`
	PanoramaId  string      `json:"panoramaId,omitempty"`
	GeneratePid bool        `json:"generatePid,omitempty"`
}

func NewResolveRequest(option ...ResolveRequestOption) (*ResolveRequest, error) {
	resolveRequest := &ResolveRequest{}
	for _, opt := range option {
		opt(resolveRequest)
	}
	return resolveRequest, nil
}

func (resolveRequest *ResolveRequest) isQueryable() bool {
	return resolveRequest.Emails != nil ||
		resolveRequest.Phones != nil ||
		resolveRequest.Profiles != nil ||
		resolveRequest.Maid != nil ||
		isPopulated(resolveRequest.PersonId) ||
		isPopulated(resolveRequest.PartnerId) ||
		isPopulated(resolveRequest.LiNonId)
}

func validateResolveRequest(resolveRequest *ResolveRequest) error {
	if !resolveRequest.isQueryable() {
		if resolveRequest.Location == nil && resolveRequest.Name == nil && !isPopulated(resolveRequest.Placekey) {
			return nil
		} else if isPopulated(resolveRequest.Placekey) && resolveRequest.Name != nil {
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
			"If you want to use 'location'(or placekey) or 'name' as an input, both must be present and they must have non-blank values")
	}
	return nil
}

func validateForIdentityMap(request *ResolveRequest) error {
	if isPopulated(request.PersonId) {
		return NewFullContactError("Invalid map request, person id must be empty")
	}
	if request.Tags != nil {
		for _, tag := range request.Tags {
			if !tag.isValid() {
				return NewFullContactError("Both Key and Value must be populated for adding a Tag")
			}
		}
	}
	if !request.isQueryable() {
		return NewFullContactError("Invalid map request, Any of Email, Phone, SocialProfile, Name and Location must be present")
	}
	err := validateResolveRequest(request)
	return err
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

func WithTagForResolve(tag *Tag) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Tags == nil {
			resolveRequest.Tags = make([]*Tag, 0)
		}
		resolveRequest.Tags = append(resolveRequest.Tags, tag)
	}
}

func WithTagsForResolve(tags []*Tag) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		if resolveRequest.Tags == nil {
			resolveRequest.Tags = make([]*Tag, 0)
		}
		resolveRequest.Tags = append(resolveRequest.Tags, tags...)
	}
}

func WithPlacekeyForResolve(placekey string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		resolveRequest.Placekey = placekey
	}
}

func WithPanoramaIDForResolve(panoramaId string) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		resolveRequest.PanoramaId = panoramaId
	}
}

func WithGeneratePidForResolve(generatePid bool) ResolveRequestOption {
	return func(resolveRequest *ResolveRequest) {
		resolveRequest.GeneratePid = generatePid
	}
}
