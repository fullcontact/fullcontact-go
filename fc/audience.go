package fullcontact

//Audience Request

type AudienceRequestOption func(ar *AudienceRequest)

type AudienceRequest struct {
	WebhookURL string `json:"webhookUrl,omitempty"`
	Tags       []*Tag `json:"tags,omitempty"`
}

func NewAudienceRequest(option ...AudienceRequestOption) (*AudienceRequest, error) {
	audienceRequest := &AudienceRequest{}

	for _, opt := range option {
		opt(audienceRequest)
	}
	err := validateAudienceRequest(audienceRequest)
	if err != nil {
		audienceRequest = nil
	}
	return audienceRequest, err
}

func validateAudienceRequest(audienceRequest *AudienceRequest) error {
	if !isPopulated(audienceRequest.WebhookURL) {
		return NewFullContactError("WebhookUrl is mandatory for creating Audience")
	}
	if len(audienceRequest.Tags) < 1 {
		return NewFullContactError("At least 1 Tag is mandatory for creating Audience")
	}
	for _, tag := range audienceRequest.Tags {
		if !tag.isValid() {
			return NewFullContactError("Both Key and Value must be populated for adding a Tag")
		}
	}
	return nil
}

func WithWebhookUrlForAudience(webhookUrl string) AudienceRequestOption {
	return func(audienceRequest *AudienceRequest) {
		audienceRequest.WebhookURL = webhookUrl
	}
}

func WithTagForAudience(tag *Tag) AudienceRequestOption {
	return func(audienceRequest *AudienceRequest) {
		if audienceRequest.Tags == nil {
			audienceRequest.Tags = make([]*Tag, 0)
		}
		audienceRequest.Tags = append(audienceRequest.Tags, tag)
	}
}

func WithTagsForAudience(tags []*Tag) AudienceRequestOption {
	return func(audienceRequest *AudienceRequest) {
		if audienceRequest.Tags == nil {
			audienceRequest.Tags = make([]*Tag, 0)
		}
		audienceRequest.Tags = append(audienceRequest.Tags, tags...)
	}
}

//Audience Response

type AudienceResponse struct {
	RequestId     string `json:"requestId"`
	AudienceBytes []byte
}
