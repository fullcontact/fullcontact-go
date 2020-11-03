package fullcontact

const (
	version                    = "1.0.0"
	userAgent                  = "FullContact_Go_Client_V" + version
	FcApiKey                   = "FC_API_KEY"
	FCGoClientTestType         = "FCGoClientTestType"
	baseUrl                    = "https://api.fullcontact.com/v3/"
	personEnrichUrl            = baseUrl + "person.enrich"
	companyEnrichUrl           = baseUrl + "company.enrich"
	companySearchUrl           = baseUrl + "company.search"
	identityMapUrl             = baseUrl + "identity.map"
	identityResolveUrl         = baseUrl + "identity.resolve"
	identityResolveWithTagsUrl = baseUrl + "identity.resolve?tags=true"
	identityDeleteUrl          = baseUrl + "identity.delete"
)
