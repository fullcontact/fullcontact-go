package fullcontact

type ResolveResponse struct {
	RecordIds []string `json:"recordIds"`
	PersonIds []string `json:"personIds"`
}

type ResolveResponseWithTags struct {
	RecordIds []string         `json:"recordIds"`
	PersonIds []string         `json:"personIds"`
	Tags      map[string][]Tag `json:"tags"`
}
