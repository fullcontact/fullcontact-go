package fullcontact

import (
	"encoding/json"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestMarshallNewPermissionRequest(t *testing.T) {
	emails := []string{"test1@gmail.com", "test2@outlook.com"}
	profile1, err := NewProfile(WithUrl("https://twitter.com/mcreedy"))
	assert.NoError(t, err)
	profile2, err := NewProfile(WithUrl("https://twitter.com/mcreedytest"))
	assert.NoError(t, err)
	requestJson := "{\"emails\":[\"marianrd97@outlook.com\",\"test1@gmail.com\",\"test2@outlook.com\"],\"phones\":[\"123-4567890\"],\"dataFilter\":[\"individual\",\"social\"],\"maids\":[\"abcd-123-abcd-1234-abcdlkjhasdfgh\",\"1234-snbk-lkldiemvmruixp-2kdp-vdm\"],\"location\":{\"addressLine1\":\"123/23\",\"addressLine2\":\"Some Street\",\"city\":\"Denver\",\"region\":\"Denver\",\"regionCode\":\"123123\",\"postalCode\":\"23124\"},\"name\":{\"given\":\"Marian\",\"family\":\"Reed\",\"full\":\"Marian C Reed\"},\"profiles\":[{\"url\":\"https://twitter.com/mcreedy\"},{\"url\":\"https://twitter.com/mcreedytest\"}],\"webhookUrl\":\"http://www.fullcontact.com/hook\",\"recordId\":\"customer123\",\"personId\":\"VS1OPPPPvxHcCNPezUbvYBCDEAOdSj5AI0adsA2bLmh12345\",\"confidence\":\"HIGH\"}"
	pr, err := NewPermissionRequest(
		WithNameForPermission(NewPersonName(WithFull("Marian C Reed"), WithFamily("Reed"), WithGiven("Marian"))),
		WithEmailForPermission("marianrd97@outlook.com"),
		WithEmailsForPermission(emails),
		WithPhoneForPermission("123-4567890"),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))),
		WithProfileForPermission(profile1),
		WithProfileForPermission(profile2),
		WithMaidsForPermission("abcd-123-abcd-1234-abcdlkjhasdfgh"),
		WithMaidsForPermission("1234-snbk-lkldiemvmruixp-2kdp-vdm"))
	assert.NoError(t, err)
	reqBytes, err := json.Marshal(pr)
	assert.NoError(t, err)
	assert.Equal(t, requestJson, string(reqBytes))
}

func TestNewPermissionRequestWithoutNameAndLocation(t *testing.T) {
	pr, _ := NewPermissionRequest(WithEmailForPermission("marianrd97@outlook.com"))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithNameOnlyWithQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithEmailForPermission("marianrd97@outlook.com"),
		WithNameForPermission(&PersonName{
			Full: "Marian C Reed",
		}))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithNameOnlyWithoutQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithNameForPermission(&PersonName{
			Full: "Marian C Reed",
		}))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: If you want to use 'location' or 'name' as an input, both must be present and they must have non-blank values")
}

func TestNewPermissionRequestWithLocationOnlyWithQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithEmailForPermission("marianrd97@outlook.com"),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationOnlyWithoutQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: If you want to use 'location' or 'name' as an input, both must be present and they must have non-blank values")
}

func TestNewPermissionRequestWithLocationWithoutAddressLine1WithQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithEmailForPermission("marianrd97@outlook.com"),
		WithNameForPermission(NewPersonName(WithFull("Test Name"))),
		WithLocationForPermission(NewLocation(
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationWithoutAddressLine1WithoutQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithNameForPermission(NewPersonName(WithFull("Test Name"))),
		WithLocationForPermission(NewLocation(
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPermissionRequestWithLocationOnlyAddressLine1WithQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithEmailForPermission("marianrd97@outlook.com"),
		WithNameForPermission(NewPersonName(WithFull("Test Name"))),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"))))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationOnlyAddressLine1WithoutQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithNameForPermission(NewPersonName(WithFull("Test Name"))),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"))))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPermissionRequestWithLocationWithAddressLine1AndCityWithQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithEmailForPermission("marianrd97@outlook.com"),
		WithNameForPermission(NewPersonName(WithFull("Test Name"))),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"))))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationWithAddressLine1AndCityWithoutQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithNameForPermission(NewPersonName(WithFull("Test Name"))),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"))))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPermissionRequestWithLocationWithAddressLine1AndRegionWithQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithEmailForPermission("marianrd97@outlook.com"),
		WithNameForPermission(NewPersonName(WithFull("Test Name"))),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithRegionCode("123123"))))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationWithAddressLine1AndRegionWithoutQueryable(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithNameForPermission(NewPersonName(WithFull("Test Name"))),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithRegionCode("123123"))))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPermissionRequestWithValidLocation1(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithNameForPermission(&PersonName{Full: "Marian C Reed"}),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("12343"))))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithValidLocation2(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithNameForPermission(&PersonName{Full: "Marian C Reed"}),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"),
			WithRegionCode("123123"))))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithValidLocation3(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithNameForPermission(&PersonName{Full: "Marian C Reed"}),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("123123"))))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithValidName(t *testing.T) {
	pr, _ := NewPermissionRequest(
		WithNameForPermission(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForPermission(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}


func TestNilPermissionCreateRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionCreate(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission create request can't be nil")
}

func TestNilPermissionDeleteRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionDelete(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission delete request can't be nil")
}

func TestNilPermissionFindRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionFind(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission find request can't be nil")
}

func TestNilPermissionCurrentRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionCurrent(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission current request can't be nil")
}

func TestNilPermissionVerifyRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionVerify(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission verify request can't be nil")
}
