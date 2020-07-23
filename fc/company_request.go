package fullcontact

type CompanyRequestOption func(cr *CompanyRequest)

type CompanyRequest struct {
	Domain      string `json:"domain,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	WebhookUrl  string `json:"webhoookUrl,omitempty"`
	Location    string `json:"location,omitempty"`
	Locality    string `json:"locality,omitempty"`
	Region      string `json:"region,omitempty"`
	Country     string `json:"country,omitempty"`
	Sort        string `json:"sort,omitempty"`
}

func NewCompanyRequest(options ...CompanyRequestOption) (*CompanyRequest, error) {
	cr := &CompanyRequest{}

	for _, opts := range options {
		opts(cr)
	}
	err := validateCompanyRequest(cr)
	return cr, err
}

func validateCompanyRequest(cr *CompanyRequest) error {
	if isPopulated(cr.Sort) && cr.Sort != "traffic" && cr.Sort != "relevance" && cr.Sort != "employees" {
		return NewFullContactError("Sort value can only be 'traffic','relevance','employees'")
	}
	return nil
}

func validateForCompanyEnrich(request *CompanyRequest) error {
	if !isPopulated(request.Domain) {
		return NewFullContactError("Company Domain is mandatory for Company Enrich")
	}
	return nil
}

func validateForCompanySearch(request *CompanyRequest) error {
	if !isPopulated(request.CompanyName) {
		return NewFullContactError("Company Name is mandatory for Company Search")
	}
	return nil
}

func WithDomain(domain string) CompanyRequestOption {
	return func(cr *CompanyRequest) {
		cr.Domain = domain
	}
}

func WithCompanyName(companyName string) CompanyRequestOption {
	return func(cr *CompanyRequest) {
		cr.CompanyName = companyName
	}
}

func WithWebhookUrlForCompany(webhookUrl string) CompanyRequestOption {
	return func(cr *CompanyRequest) {
		cr.WebhookUrl = webhookUrl
	}
}

func WithLocationForCompany(location string) CompanyRequestOption {
	return func(cr *CompanyRequest) {
		cr.Location = location
	}
}

func WithLocality(locality string) CompanyRequestOption {
	return func(cr *CompanyRequest) {
		cr.Locality = locality
	}
}

func WithRegion(region string) CompanyRequestOption {
	return func(cr *CompanyRequest) {
		cr.Region = region
	}
}

func WithCountry(country string) CompanyRequestOption {
	return func(cr *CompanyRequest) {
		cr.Country = country
	}
}

func WithSort(sort string) CompanyRequestOption {
	return func(cr *CompanyRequest) {
		cr.Sort = sort
	}
}
