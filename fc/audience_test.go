package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestNewAudienceRequest(t *testing.T) {
	_, err := NewAudienceRequest(WithWebhookUrlForAudience("webhookUrl"),
		WithTagForAudience(NewTag(WithTagKey("key"), WithTagValue("value"))))
	assert.NoError(t, err)
}

func TestInvalidNewAudienceRequest1(t *testing.T) {
	_, err := NewAudienceRequest(WithTagForAudience(NewTag(WithTagKey("key"), WithTagValue("value"))))
	assert.EqualError(t, err, "FullContactError: WebhookUrl is mandatory for creating Audience")
}

func TestInvalidNewAudienceRequest2(t *testing.T) {
	_, err := NewAudienceRequest(WithWebhookUrlForAudience("webhookUrl"))
	assert.EqualError(t, err, "FullContactError: At least 1 Tag is mandatory for creating Audience")
}

func TestInvalidNewAudienceRequest3(t *testing.T) {
	_, err := NewAudienceRequest(WithWebhookUrlForAudience("webhookUrl"),
		WithTagForAudience(NewTag(WithTagKey("key"))))
	assert.EqualError(t, err, "FullContactError: Both Key and Value must be populated for adding a Tag")
}
