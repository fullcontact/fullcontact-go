package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestNewProfile1(t *testing.T) {
	_, err := NewProfile(WithUrl("https://twitter.com/mcreedy"))
	assert.NoError(t, err)
}

func TestNewProfile2(t *testing.T) {
	_, err := NewProfile(WithService("twitter"), WithUrl("mcreedy"))
	assert.NoError(t, err)
}

func TestNewProfile3(t *testing.T) {
	_, err := NewProfile(WithService("twitter"), WithUsername("mcreedy"))
	assert.NoError(t, err)
}

func TestNewProfileInvalid1(t *testing.T) {
	_, err := NewProfile(WithUrl("https://twitter.com/mcreedy"), WithUserid("mcreedy"))
	assert.EqualError(t, err, "FullContactError: Specifying username or userid together with url is not allowed")
}

func TestNewProfileInvalid2(t *testing.T) {
	_, err := NewProfile(WithUrl("https://twitter.com/mcreedy"), WithUsername("mcreedy"))
	assert.EqualError(t, err, "FullContactError: Specifying username or userid together with url is not allowed")
}

func TestNewProfileInvalid3(t *testing.T) {
	_, err := NewProfile(WithService("twitter"))
	assert.EqualError(t, err, "FullContactError: Either url or service plus username or userid must be set on every profiles entry.")
}

func TestNewProfileInvalid4(t *testing.T) {
	_, err := NewProfile(WithService("twitter"), WithUserid("test"), WithUsername("test"))
	assert.EqualError(t, err, "FullContactError: Specifying userid together with username is not allowed")
}

func TestNewProfileInvalid5(t *testing.T) {
	_, err := NewProfile(WithUsername("test"))
	assert.EqualError(t, err, "FullContactError: Either url or service plus username or userid must be set on every profiles entry.")
}
