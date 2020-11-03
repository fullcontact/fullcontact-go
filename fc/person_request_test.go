package fullcontact

import (
	"encoding/json"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestMarshallNewPersonRequest(t *testing.T) {
	emails := []string{"test1@gmail.com", "test2@outlook.com"}
	profile1, err := NewProfile(WithUrl("https://twitter.com/mcreedy"))
	assert.NoError(t, err)
	profile2, err := NewProfile(WithUrl("https://twitter.com/mcreedytest"))
	assert.NoError(t, err)
	requestJson := "{\"emails\":[\"marianrd97@outlook.com\",\"test1@gmail.com\",\"test2@outlook.com\"],\"phones\":[\"123-4567890\"],\"dataFilters\":[\"individual\",\"social\"],\"maids\":[\"abcd-123-abcd-1234-abcdlkjhasdfgh\",\"1234-snbk-lkldiemvmruixp-2kdp-vdm\"],\"location\":{\"addressLine1\":\"123/23\",\"addressLine2\":\"Some Street\",\"city\":\"Denver\",\"region\":\"Denver\",\"regionCode\":\"123123\",\"postalCode\":\"23124\"},\"name\":{\"given\":\"Marian\",\"family\":\"Reed\",\"full\":\"Marian C Reed\"},\"profiles\":[{\"url\":\"https://twitter.com/mcreedy\"},{\"url\":\"https://twitter.com/mcreedytest\"}],\"webhookUrl\":\"http://www.fullcontact.com/hook\",\"recordId\":\"customer123\",\"personId\":\"VS1OPPPPvxHcCNPezUbvYBCDEAOdSj5AI0adsA2bLmh12345\",\"confidence\":\"HIGH\"}"
	pr, err := NewPersonRequest(
		WithName(NewPersonName(WithFull("Marian C Reed"), WithFamily("Reed"), WithGiven("Marian"))),
		WithEmail("marianrd97@outlook.com"),
		WithEmails(emails),
		WithPhone("123-4567890"),
		WithConfidence("HIGH"),
		WithInfer(false),
		WithDataFilter("individual"),
		WithDataFilter("social"),
		WithLocation(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))),
		WithProfile(profile1),
		WithProfile(profile2),
		WithMaid("abcd-123-abcd-1234-abcdlkjhasdfgh"),
		WithMaid("1234-snbk-lkldiemvmruixp-2kdp-vdm"),
		WithWebhookUrl("http://www.fullcontact.com/hook"),
		WithRecordId("customer123"),
		WithPersonId("VS1OPPPPvxHcCNPezUbvYBCDEAOdSj5AI0adsA2bLmh12345"))
	assert.NoError(t, err)
	reqBytes, err := json.Marshal(pr)
	assert.NoError(t, err)
	assert.Equal(t, requestJson, string(reqBytes))
}

func TestNewPersonRequestWithoutNameAndLocation(t *testing.T) {
	_, err := NewPersonRequest(WithEmail("marianrd97@outlook.com"))
	assert.NoError(t, err)
}

func TestNewPersonRequestWithNameOnly(t *testing.T) {
	_, err := NewPersonRequest(WithEmail("marianrd97@outlook.com"),
		WithName(&PersonName{
			Full: "Marian C Reed",
		}))
	assert.EqualError(t, err, "FullContactError: If you want to use 'location' or 'name' as an input, both must be present and they must have non-blank values")
}

func TestNewPersonRequestWithLocationOnly(t *testing.T) {
	_, err := NewPersonRequest(
		WithEmail("marianrd97@outlook.com"),
		WithLocation(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	assert.EqualError(t, err, "FullContactError: If you want to use 'location' or 'name' as an input, both must be present and they must have non-blank values")
}

func TestNewPersonRequestWithLocationWithoutAddressLine1(t *testing.T) {
	_, err := NewPersonRequest(
		WithEmail("marianrd97@outlook.com"),
		WithName(NewPersonName(WithFull("Test Name"))),
		WithLocation(NewLocation(
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPersonRequestWithLocationOnlyAddressLine1(t *testing.T) {
	_, err := NewPersonRequest(
		WithEmail("marianrd97@outlook.com"),
		WithName(NewPersonName(WithFull("Test Name"))),
		WithLocation(NewLocation(
			WithAddressLine1("123/23"))))
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPersonRequestWithLocationWithAddressLine1AndCity(t *testing.T) {
	_, err := NewPersonRequest(
		WithEmail("marianrd97@outlook.com"),
		WithName(NewPersonName(WithFull("Test Name"))),
		WithLocation(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"))))
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPersonRequestWithLocationWithAddressLine1AndRegion(t *testing.T) {
	_, err := NewPersonRequest(
		WithEmail("marianrd97@outlook.com"),
		WithName(NewPersonName(WithFull("Test Name"))),
		WithLocation(NewLocation(
			WithAddressLine1("123/23"),
			WithRegionCode("123123"))))
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPersonRequestWithValidLocation1(t *testing.T) {
	_, err := NewPersonRequest(
		WithName(&PersonName{Full: "Marian C Reed"}),
		WithLocation(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("12343"))))
	assert.NoError(t, err)
}

func TestNewPersonRequestWithValidLocation2(t *testing.T) {
	_, err := NewPersonRequest(
		WithName(&PersonName{Full: "Marian C Reed"}),
		WithLocation(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"),
			WithRegionCode("123123"))))
	assert.NoError(t, err)
}

func TestNewPersonRequestWithValidLocation3(t *testing.T) {
	_, err := NewPersonRequest(
		WithName(&PersonName{Full: "Marian C Reed"}),
		WithLocation(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("123123"))))
	assert.NoError(t, err)
}

func TestNewPersonRequestWithValidName(t *testing.T) {
	_, err := NewPersonRequest(
		WithName(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocation(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	assert.NoError(t, err)
}

func TestWithConfidence(t *testing.T) {
	_, err := NewPersonRequest(WithConfidence("test"))
	assert.EqualError(t, err, "FullContactError: Confidence value can only be 'LOW', 'MED', 'HIGH', 'MAX'")
}

func TestNilPersonRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PersonEnrich(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Person Request can't be nil")
}
