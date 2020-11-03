package fullcontact

type ResolveResponse struct {
	RecordIds  []string `json:"recordIds"`
	PersonIds  []string `json:"personIds"`
	PartnerIds []string `json:"partnerIds"`
}

type ResolveResponseWithTags struct {
	RecordIds  []string         `json:"recordIds"`
	PersonIds  []string         `json:"personIds"`
	PartnerIds []string         `json:"partnerIds"`
	Tags       map[string][]Tag `json:"tags"`
}
