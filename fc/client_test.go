package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewFullContactClient(t *testing.T) {
	cp, err := NewStaticCredentialsProvider("apikey")
	assert.NoError(t, err)

	_, err = NewFullContactClient(
		WithCredentialsProvider(cp),
		WithHeaders(map[string]string{"Reporting-Key": "FC_GoClient"}),
		WithTimeout(3000))
	assert.NoError(t, err)
}

type CustomRetryHandler struct{}

func (crh CustomRetryHandler) ShouldRetry(responseCode int) bool {
	if responseCode == 429 {
		return true
	}
	return false
}

func (crh CustomRetryHandler) RetryAttempts() int {
	return 2
}

func (crh CustomRetryHandler) RetryDelayMillis() int {
	return 2000
}

func TestNewFullContactClientWithRetryHandler(t *testing.T) {
	cp, err := NewStaticCredentialsProvider("apikey")
	assert.NoError(t, err)

	//Creating a FullContact Client
	_, err = NewFullContactClient(
		WithCredentialsProvider(cp),
		WithHeaders(map[string]string{"Reporting-Key": "FC_GoClient"}),
		WithTimeout(3000),
		WithRetryHandler(&CustomRetryHandler{}))
	assert.NoError(t, err)
}

func TestNewFullContactClientWithoutAuth(t *testing.T) {
	_, err := NewFullContactClient(
		WithHeaders(map[string]string{"Reporting-Key": "FC_GoClient"}),
		WithTimeout(3000))
	assert.EqualError(t, err, "FullContactError: Couldn't find valid API Key from ENV variable: FC_API_KEY")
}

func TestNewFullContactClientWithAuth(t *testing.T) {
	err := os.Setenv(FcApiKey, "apikey")
	defer os.Unsetenv(FcApiKey)
	if err != nil {
		_, err = NewFullContactClient(
			WithHeaders(map[string]string{"Reporting-Key": "FC_GoClient"}),
			WithTimeout(3000), WithTimeout(-1))
		assert.NoError(t, err)
	}
}
