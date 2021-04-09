package fullcontact

type MultifieldRequestOption func(ar *MultifieldRequest)

type MultifieldRequest struct {
	Emails     			[]string    		`json:"emails,omitempty"`
	Phones     			[]string    		`json:"phones,omitempty"`
	Maids      			[]string    		`json:"maids,omitempty"`
	Location   			*Location   		`json:"location,omitempty"`
	Name       			*PersonName 		`json:"name,omitempty"`
	Profiles   			[]*Profile  		`json:"profiles,omitempty"`
	PersonId    		string      		`json:"personId,omitempty"`
	RecordId    		string      		`json:"recordId,omitempty"`
	PartnerId    		string      		`json:"partnerId,omitempty"`
	LiNonId    			string      		`json:"li_nonid,omitempty"`
	Placekey    		string      		`json:"placekey,omitempty"`
}

func NewMultifieldRequest(option ...MultifieldRequestOption) (*MultifieldRequest, error) {
	multifieldRequest := &MultifieldRequest{}

	for _, opt := range option {
		opt(multifieldRequest)
	}
	return multifieldRequest, nil
}

func (multifieldRequest *MultifieldRequest) isQueryable() bool {
	return multifieldRequest.Emails != nil ||
		multifieldRequest.Phones != nil ||
		multifieldRequest.Profiles != nil ||
		multifieldRequest.Maids != nil ||
		isPopulated(multifieldRequest.LiNonId)
}

func (multifieldRequest *MultifieldRequest) isValidName() bool {
	return (isPopulated(multifieldRequest.Name.Full)) ||
		(isPopulated(multifieldRequest.Name.Given) && isPopulated(multifieldRequest.Name.Family))
}

func (multifieldRequest *MultifieldRequest) isValidLocation() bool {
	return isPopulated(multifieldRequest.Location.AddressLine1) &&
		((isPopulated(multifieldRequest.Location.City) &&
			(isPopulated(multifieldRequest.Location.Region) || isPopulated(multifieldRequest.Location.RegionCode))) ||
			(isPopulated(multifieldRequest.Location.PostalCode)))
}

func (multifieldRequest *MultifieldRequest) validate() error {
	if !multifieldRequest.isQueryable() {
		if multifieldRequest.Location == nil && multifieldRequest.Name == nil && !isPopulated(multifieldRequest.Placekey) {
			return nil
		} else if (multifieldRequest.Location != nil || isPopulated(multifieldRequest.Placekey)) && multifieldRequest.Name != nil {
			if (multifieldRequest.Location != nil && multifieldRequest.isValidLocation()) || isPopulated(multifieldRequest.Placekey) {
				if multifieldRequest.isValidName() {
					return nil
				} else {
					return NewFullContactError("Name data requires full name or given and family name")
				}
			} else {
				return NewFullContactError("A valid placekey is required or Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
			}
		}
		return NewFullContactError(
			"If you want to use 'location'(or placekey) or 'name' as an input, both must be present and they must have non-blank values")
	}
	return nil
}

func WithMaidsForMultifieldRequest(maid string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		if multifieldRequest.Maids == nil {
			multifieldRequest.Maids = make([]string, 0)
		}
		multifieldRequest.Maids = append(multifieldRequest.Maids, maid)
	}
}

func WithEmailForMultifieldRequest(email string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		if multifieldRequest.Emails == nil {
			multifieldRequest.Emails = make([]string, 0)
		}
		multifieldRequest.Emails = append(multifieldRequest.Emails, email)
	}
}

func WithEmailsForMultifieldRequest(emails []string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		if multifieldRequest.Emails == nil {
			multifieldRequest.Emails = make([]string, 0)
		}
		multifieldRequest.Emails = append(multifieldRequest.Emails, emails...)
	}
}

func WithPhoneForMultifieldRequest(phone string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		if multifieldRequest.Phones == nil {
			multifieldRequest.Phones = make([]string, 0)
		}
		multifieldRequest.Phones = append(multifieldRequest.Phones, phone)
	}
}

func WithPhonesForMultifieldRequest(phones []string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		if multifieldRequest.Phones == nil {
			multifieldRequest.Phones = make([]string, 0)
		}
		multifieldRequest.Phones = append(multifieldRequest.Phones, phones...)
	}
}

func WithProfileForMultifieldRequest(profile *Profile) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		if multifieldRequest.Profiles == nil {
			multifieldRequest.Profiles = make([]*Profile, 0)
		}
		multifieldRequest.Profiles = append(multifieldRequest.Profiles, profile)
	}
}

func WithProfilesForMultifieldRequest(profiles []*Profile) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		if multifieldRequest.Profiles == nil {
			multifieldRequest.Profiles = make([]*Profile, 0)
		}
		multifieldRequest.Profiles = append(multifieldRequest.Profiles, profiles...)
	}
}

func WithNameForMultifieldRequest(name *PersonName) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		multifieldRequest.Name = name
	}
}

func WithLocationForMultifieldRequest(location *Location) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		multifieldRequest.Location = location
	}
}

func WithLiNonIdForMultifieldRequest(liNonId string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		multifieldRequest.LiNonId = liNonId
	}
}

func WithPersonIdForMultifieldRequest(personId string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		multifieldRequest.PersonId = personId
	}
}

func WithRecordIdForMultifieldRequest(recordId string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		multifieldRequest.RecordId = recordId
	}
}

func WithPartnerIdForMultifieldRequest(partnerId string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		multifieldRequest.PartnerId = partnerId
	}
}

func WithPlacekeyForMultifieldRequest(placekey string) MultifieldRequestOption {
	return func(multifieldRequest *MultifieldRequest) {
		multifieldRequest.Placekey = placekey
	}
}
