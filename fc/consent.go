package fullcontact

type ConsentPurposesOption func(consentPurposes *ConsentPurposes)

type ConsentPurposes struct {
	PurposeId		int			`json:"purposeId"`
	Channel			[]string	`json:"channel"`
	Ttl				int			`json:"ttl"`
	Enabled			bool		`json:"enabled"`
}

func NewConsentPurposes(options ...ConsentPurposesOption) (*ConsentPurposes, error) {
	consentPurpose := &ConsentPurposes{}
	for _, opts := range options {
		opts(consentPurpose)
	}

	err := validateConsentPurpose(consentPurpose)
	if err != nil {
		consentPurpose = nil
	}
	return consentPurpose, nil
}

func validateConsentPurpose(consentPurpose *ConsentPurposes) error {
	if consentPurpose.PurposeId == 0 {
		return NewFullContactError("Purpose id is required for consentPurpose")
	} else if consentPurpose.Channel == nil {
		return NewFullContactError("Channel is required for consentPurpose")
	} else if &consentPurpose.Enabled == nil {
		return NewFullContactError("Enabled is required for consentPurpose")
	}
	return nil
}

func WithConsentPurposeId(purposeId int) ConsentPurposesOption {
	return func(consentPurpose *ConsentPurposes) {
		consentPurpose.PurposeId = purposeId
	}
}

func WithConsentPurposeChannel(channel string) ConsentPurposesOption {
	return func(consentPurpose *ConsentPurposes) {
		if consentPurpose.Channel == nil {
			consentPurpose.Channel = make([]string, 0)
		}
		consentPurpose.Channel = append(consentPurpose.Channel, channel)
	}
}

func WithConsentPurposeChannels(channel []string) ConsentPurposesOption {
	return func(consentPurpose *ConsentPurposes) {
		if consentPurpose.Channel == nil {
			consentPurpose.Channel = make([]string, 0)
		}
		consentPurpose.Channel = append(consentPurpose.Channel, channel...)
	}
}

func WithConsentPurposeTtl(ttl int) ConsentPurposesOption {
	return func(consentPurpose *ConsentPurposes) {
		consentPurpose.Ttl = ttl
	}
}

func WithConsentPurposeEnabled(enabled bool) ConsentPurposesOption {
	return func(consentPurpose *ConsentPurposes) {
		consentPurpose.Enabled = enabled
	}
}
