package fullcontact

type LocationOption func(location *Location)

type Location struct {
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	City         string `json:"city,omitempty"`
	Region       string `json:"region,omitempty"`
	RegionCode   string `json:"regionCode,omitempty"`
	Country      string `json:"country,omitempty"`
	CountryCode  string `json:"countryCode,omitempty"`
	Formatted    string `json:"formatted,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
	Type         string `json:"type,omitempty"`
	Label        string `json:"type,omitempty"`
}

func NewLocation(options ...LocationOption) *Location {
	location := &Location{}
	for _, opts := range options {
		opts(location)
	}
	return location
}

func WithAddressLine1(addressLine1 string) LocationOption {
	return func(location *Location) {
		location.AddressLine1 = addressLine1
	}
}

func WithAddressLine2(addressLine2 string) LocationOption {
	return func(location *Location) {
		location.AddressLine2 = addressLine2
	}
}

func WithCity(city string) LocationOption {
	return func(location *Location) {
		location.City = city
	}
}

func WithRegionForLocation(region string) LocationOption {
	return func(location *Location) {
		location.Region = region
	}
}

func WithRegionCode(regionCode string) LocationOption {
	return func(location *Location) {
		location.RegionCode = regionCode
	}
}

func WithPostalCode(postalCode string) LocationOption {
	return func(location *Location) {
		location.PostalCode = postalCode
	}
}

func WithCountryForLocation(country string) LocationOption {
	return func(location *Location) {
		location.Country = country
	}
}

func WithCountryCode(countryCode string) LocationOption {
	return func(location *Location) {
		location.CountryCode = countryCode
	}
}

func WithFormatted(formatted string) LocationOption {
	return func(location *Location) {
		location.Formatted = formatted
	}
}
