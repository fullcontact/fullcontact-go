package fullcontact


type PermissionFindResponse struct {
	PermissionType      string          	`json:"permissionType"`
	PermissionId   		string          	`json:"permissionId"`
	ConsentPurposes    	[]*ConsentPurpose   `json:"consentPurposes"`
	Locale     			string      		`json:"locale"`
	IpAddress			string				`json:"ipAddress"`
	Language     		string      		`json:"language"`
	CollectionMethod	string				`json:"collectionMethod"`
	CollectionLocation 	string      		`json:"collectionLocation"`
	PolicyUrl     		string      		`json:"policyUrl"`
	TermsService		string				`json:"termsService"`
	Timestamp			string				`json:"timestamp"`
	Created				string				`json:"created"`
}

type ConsentPurpose struct {
	PurposeId		int			`json:"purposeId"`
	Channel			string		`json:"channel"`
	Ttl				int			`json:"ttl"`
	Enabled			bool		`json:"enabled"`
	AsOfTimestamp 	int			`json:"asOfTimestamp"`
}

type PermissionVerifyResponse struct {
	Ttl				string		`json:"ttl"`
	Enabled			bool		`json:"enabled"`
	Channel			string		`json:"channel"`
	PurposeId		int			`json:"purposeId"`
	PurposeName		string		`json:"purposeName"`
	timestamp		int			`json:"timestamp"`
}

type PermissionCurrentResponse struct {
	Ttl				string		`json:"ttl"`
	Enabled			bool		`json:"enabled"`
	Channel			string		`json:"channel"`
	PurposeId		int			`json:"purposeId"`
	timestamp		int			`json:"timestamp"`
}
