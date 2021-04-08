package fullcontact

type PermissionFindResponse struct {
	PermissionType      string						`json:"permissionType"`
	PermissionId   		string						`json:"permissionId"`
	ConsentPurposes    	[]*ConsentPurposeResponse	`json:"consentPurposes"`
	Locale     			string						`json:"locale"`
	IpAddress			string						`json:"ipAddress"`
	Language     		string						`json:"language"`
	CollectionMethod	string						`json:"collectionMethod"`
	CollectionLocation 	string						`json:"collectionLocation"`
	PolicyUrl     		string						`json:"policyUrl"`
	TermsService		string						`json:"termsService"`
	Timestamp			int							`json:"timestamp"`
	Created				int							`json:"created"`
}

type ConsentPurposeResponse struct {
	Ttl				int			`json:"ttl"`
	Enabled			bool		`json:"enabled"`
	Channel			string		`json:"channel"`
	PurposeId		int			`json:"purposeId"`
	PurposeName		string		`json:"purposeName"`
	timestamp 		int			`json:"timestamp"`
}
