package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewStaticCredentialsProvider(t *testing.T) {
	cp, err := NewStaticCredentialsProvider("apikey")
	assert.NoError(t, err)
	assert.Equal(t, "apikey", cp.getApiKey())
}

func TestNewStaticCredentialsProviderWithEmptyKey(t *testing.T) {
	_, err := NewStaticCredentialsProvider("")
	assert.EqualError(t, err, "FullContactError: API Key can't be empty")
}

func TestNewDefaultCredentialsProvider(t *testing.T) {
	err := os.Setenv("FC_API_KEY", "apikey")
	defer os.Unsetenv("FC_API_KEY")
	if err != nil {
		cp, err := NewDefaultCredentialsProvider(FcApiKey)
		assert.NoError(t, err)
		assert.Equal(t, "apikey", cp.getApiKey())
	}
}

func TestNewDefaultCredentialsProviderWithoutEnv(t *testing.T) {
	_, err := NewDefaultCredentialsProvider(FcApiKey)
	assert.EqualError(t, err, "FullContactError: Couldn't find valid API Key from ENV variable: FC_API_KEY")
}
