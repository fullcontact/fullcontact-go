package fullcontact

const (
	version                    = "1.3.0"
	userAgent                  = "FullContact_Go_Client_V" + version
	FcApiKey                   = "FC_API_KEY"
	FCGoClientTestType         = "FCGoClientTestType"
	baseUrl                    = "https://api.fullcontact.com/v3/"
	personEnrichUrl            = baseUrl + "person.enrich"
	companyEnrichUrl           = baseUrl + "company.enrich"
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
	permissionCreateUrl        = baseUrl + "permission.create"
	permissionDeleteUrl        = baseUrl + "permission.delete"
	permissionFindUrl          = baseUrl + "permission.find"
	permissionCurrentUrl       = baseUrl + "permission.current"
	permissionVerifyUrl        = baseUrl + "permission.verify"
	verifySignalsUrl           = baseUrl + "verify.signals"
	verifyMatchUrl             = baseUrl + "verify.match"
	verifyActivityUrl          = baseUrl + "verify.activity"
)
