package fullcontact

//Permission Request

type PermissionRequestOption func(ar *PermissionRequest)

type PermissionRequest struct {
	Timestamp  			int							`json:"timestamp,omitempty"`
	Query    			*MultifieldRequest			`json:"query,omitempty"`
	ConsentPurposes		[]*ConsentPurpose 			`json:"consentPurposes,omitempty"`
	Locale     			string      				`json:"locale,omitempty"`
	IpAddress			string						`json:"ipAddress,omitempty"`
	Language     		string      				`json:"language,omitempty"`
	CollectionMethod	string						`json:"collectionMethod,omitempty"`
	CollectionLocation 	string      				`json:"collectionLocation,omitempty"`
	Tcf					string						`json:"tcf,omitempty"`
	PolicyUrl     		string      				`json:"policyUrl,omitempty"`
	TermsService		string						`json:"termsService,omitempty"`
	PurposeId			int							`json:"purposeId,omitempty"`
	Channel				string						`json:"channel,omitempty"`
}

func NewPermissionRequest(option ...PermissionRequestOption) (*PermissionRequest, error) {
	permissionRequest := &PermissionRequest{}

	for _, opt := range option {
		opt(permissionRequest)
	}
	return permissionRequest, nil
}

func validateForPermissionCreateFields(request *PermissionRequest) error {
	if request.ConsentPurposes == nil {
		return NewFullContactError("At least 1 `consentPurpose` is Required for PermissionRequest")
	} else if !isPopulated(request.CollectionMethod) {
		return NewFullContactError("Collection Method is required for PermissionRequest")
	} else if !isPopulated(request.CollectionLocation) {
		return NewFullContactError("Collection Location is required for PermissionRequest")
	} else if !isPopulated(request.PolicyUrl) {
		return NewFullContactError("Policy URL is required for PermissionRequest")
	} else if !isPopulated(request.TermsService) {
		return NewFullContactError("Terms of Service is required for PermissionRequest")
	}
	return nil
}

func validateForPermissionVerifyFields(request *PermissionRequest) error {
	if request.PurposeId == 0 {
		return NewFullContactError("Purpose ID is required for PermissionRequest")
	} else if !isPopulated(request.Channel) {
		return NewFullContactError("Channel is required for PermissionRequest")
	}
	return nil
}

func validateForPermissionCreate(request *PermissionRequest) error {
	err := validateForPermissionCreateFields(request)
	if err != nil {
		return err
	}
	for _, consentPurpose := range request.ConsentPurposes{
		err = validateConsentPurpose(consentPurpose)
		if err != nil {
			return err
		}
	}
	err = request.Query.validate()
	return err
}

func validateForPermissionDelete(request *MultifieldRequest) error {
	err := request.validate()
	return err
}

func validateForPermissionFind(request *MultifieldRequest) error {
	err := request.validate()
	return err
}

func validateForPermissionCurrent(request *MultifieldRequest) error {
	err := request.validate()
	return err
}

func validateForPermissionVerify(request *PermissionRequest) error {
	err := validateForPermissionVerifyFields(request)
	if err != nil {
		return err
	}
	err = request.Query.validate()
	return err
}

func WithMultifieldRequestForPermission(multifieldRequest *MultifieldRequest) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Query = multifieldRequest
	}
}

func WithConsentPurposeForPermission(consentPurpose *ConsentPurpose) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.ConsentPurposes == nil {
			pr.ConsentPurposes = make([]*ConsentPurpose, 0)
		}
		pr.ConsentPurposes = append(pr.ConsentPurposes, consentPurpose)
	}
}

func WithConsentPurposesForPermission(consentPurpose []*ConsentPurpose) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		if pr.ConsentPurposes == nil {
			pr.ConsentPurposes = make([]*ConsentPurpose, 0)
		}
		pr.ConsentPurposes = append(pr.ConsentPurposes, consentPurpose...)
	}
}

func WithLocaleForPermission(locale string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Locale = locale
	}
}

func WithIpAddressForPermission(ipAddress string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.IpAddress = ipAddress
	}
}

func WithLanguageForPermission(language string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Language = language
	}
}

func WithCollectionMethodForPermission(collectionMethod string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.CollectionMethod = collectionMethod
	}
}

func WithCollectionLocationForPermission(collectionLocation string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.CollectionLocation = collectionLocation
	}
}

func WithTcfForPermission(tcf string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Tcf = tcf
	}
}

func WithPolicyUrlForPermission(policyUrl string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.PolicyUrl = policyUrl
	}
}

func WithTermsServiceForPermission(termsService string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.TermsService = termsService
	}
}

func WithPurposeIdForPermission(purposeId int) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.PurposeId = purposeId
	}
}

func WithChannelForPermission(channel string) PermissionRequestOption {
	return func(pr *PermissionRequest) {
		pr.Channel = channel
	}
}

type ConsentPurposeOption func(consentPurposes *ConsentPurpose)

type ConsentPurpose struct {
	PurposeId		int			`json:"purposeId"`
	Channel			[]string	`json:"channel"`
	Ttl				int			`json:"ttl"`
	Enabled			*bool		`json:"enabled"`
}

func NewConsentPurpose(options ...ConsentPurposeOption) *ConsentPurpose {
	consentPurpose := &ConsentPurpose{}
	for _, opts := range options {
		opts(consentPurpose)
	}
	return consentPurpose
}

func validateConsentPurpose(consentPurpose *ConsentPurpose) error {
	if consentPurpose.PurposeId == 0 {
		return NewFullContactError("Purpose id is required for consentPurpose")
	} else if consentPurpose.Channel == nil {
		return NewFullContactError("Channel is required for consentPurpose")
	} else if consentPurpose.Enabled == nil {
		return NewFullContactError("Enabled is required for consentPurpose")
	}
	return nil
}

func WithConsentPurposeId(purposeId int) ConsentPurposeOption {
	return func(consentPurpose *ConsentPurpose) {
		consentPurpose.PurposeId = purposeId
	}
}

func WithConsentPurposeChannel(channel string) ConsentPurposeOption {
	return func(consentPurpose *ConsentPurpose) {
		if consentPurpose.Channel == nil {
			consentPurpose.Channel = make([]string, 0)
		}
		consentPurpose.Channel = append(consentPurpose.Channel, channel)
	}
}

func WithConsentPurposeChannels(channel []string) ConsentPurposeOption {
	return func(consentPurpose *ConsentPurpose) {
		if consentPurpose.Channel == nil {
			consentPurpose.Channel = make([]string, 0)
		}
		consentPurpose.Channel = append(consentPurpose.Channel, channel...)
	}
}

func WithConsentPurposeTtl(ttl int) ConsentPurposeOption {
	return func(consentPurpose *ConsentPurpose) {
		consentPurpose.Ttl = ttl
	}
}

func WithConsentPurposeEnabled(enabled bool) ConsentPurposeOption {
	return func(consentPurpose *ConsentPurpose) {
		consentPurpose.Enabled = &enabled
	}
}

