package fullcontact

const (
	version                    = "1.2.0"
	userAgent                  = "FullContact_Go_Client_V" + version
	FcApiKey                   = "FC_API_KEY"
	FCGoClientTestType         = "FCGoClientTestType"
	baseUrl                    = "https://api.fullcontact.com/v3/"
	v2BaseUrl                  = "https://api.fullcontact.com/v2/"
	personEnrichUrl            = baseUrl + "person.enrich"
	companyEnrichUrl           = baseUrl + "company.enrich"
	companySearchUrl           = baseUrl + "company.search"
	identityMapUrl             = baseUrl + "identity.map"
	identityResolveUrl         = baseUrl + "identity.resolve"
	identityResolveWithTagsUrl = baseUrl + "identity.resolve?tags=true"
	identityMapResolveUrl      = baseUrl + "identity.mapResolve"
	identityDeleteUrl          = baseUrl + "identity.delete"
	tagsCreateUrl              = baseUrl + "tags.create"
	tagsGetUrl                 = baseUrl + "tags.get"
	tagsDeleteUrl              = baseUrl + "tags.delete"
	audienceCreateUrl          = baseUrl + "audience.create"
	audienceDownloadUrl        = baseUrl + "audience.download"
	emailVerificationUrl       = v2BaseUrl + "verification/email"
	permissionCreateUrl        = baseUrl + "permission.create"
	permissionDeleteUrl        = baseUrl + "permission.delete"
	permissionFindUrl          = baseUrl + "permission.find"
	permissionCurrentUrl       = baseUrl + "permission.current"
	permissionVerifyUrl        = baseUrl + "permission.verify"
)
