package fullcontact

import (
	"encoding/json"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestNewPermissionRequestForCreate(t *testing.T) {
	emails := []string{"test1@gmail.com", "test2@outlook.com"}
	profile1, err := NewProfile(WithUrl("https://twitter.com/mcreedy"))
	assert.NoError(t, err)
	profile2, err := NewProfile(WithUrl("https://twitter.com/mcreedytest"))
	assert.NoError(t, err)
	consentPurposes := NewConsentPurpose(
		WithConsentPurposeId(1),
		WithConsentPurposeChannel("web"),
		WithConsentPurposeTtl(365),
		WithConsentPurposeEnabled(true))
	assert.NoError(t, err)
	multifieldRequest, err := NewMultifieldRequest(
		WithNameForMultifieldRequest(NewPersonName(WithFull("Marian C Reed"), WithFamily("Reed"), WithGiven("Marian"))),
		WithEmailForMultifieldRequest("marianrd97@outlook.com"),
		WithEmailsForMultifieldRequest(emails),
		WithPhoneForMultifieldRequest("123-4567890"),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))),
		WithProfileForMultifieldRequest(profile1),
		WithProfileForMultifieldRequest(profile2),
		WithMaidsForMultifieldRequest("abcd-123-abcd-1234-abcdlkjhasdfgh"),
		WithMaidsForMultifieldRequest("1234-snbk-lkldiemvmruixp-2kdp-vdm"),)
	assert.NoError(t, err)
	requestJson := "{\"query\":{\"emails\":[\"marianrd97@outlook.com\",\"test1@gmail.com\",\"test2@outlook.com\"],\"phones\":[\"123-4567890\"],\"maids\":[\"abcd-123-abcd-1234-abcdlkjhasdfgh\",\"1234-snbk-lkldiemvmruixp-2kdp-vdm\"],\"location\":{\"addressLine1\":\"123/23\",\"addressLine2\":\"Some Street\",\"city\":\"Denver\",\"region\":\"Denver\",\"regionCode\":\"123123\",\"postalCode\":\"23124\"},\"name\":{\"given\":\"Marian\",\"family\":\"Reed\",\"full\":\"Marian C Reed\"},\"profiles\":[{\"url\":\"https://twitter.com/mcreedy\"},{\"url\":\"https://twitter.com/mcreedytest\"}]},\"consentPurposes\":[{\"purposeId\":1,\"channel\":[\"web\"],\"ttl\":365,\"enabled\":true}],\"locale\":\"US\",\"ipAddress\":\"127.0.0.1\",\"language\":\"en\",\"collectionMethod\":\"cookiePopUp\",\"collectionLocation\":\"Can we get a snapshot of where someone is opting in/out here?\",\"tcf\":\"some.valid.tcfv2.string\",\"policyUrl\":\"http://foo.baz\",\"termsService\":\"http://foo.tos\"}"
	pr, err := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(consentPurposes),
		WithLocaleForPermission("US"),
		WithIpAddressForPermission("127.0.0.1"),
		WithLanguageForPermission("en"),
		WithCollectionMethodForPermission("cookiePopUp"),
		WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		WithTcfForPermission("some.valid.tcfv2.string"),
		WithPolicyUrlForPermission("http://foo.baz"),
		WithTermsServiceForPermission("http://foo.tos"))
	assert.NoError(t, err)
	reqBytes, err := json.Marshal(pr)
	assert.NoError(t, err)
	assert.Equal(t, requestJson, string(reqBytes))
}

func TestNewPermissionRequestForFind(t *testing.T) {
	emails := []string{"test1@gmail.com", "test2@outlook.com"}
	profile1, err := NewProfile(WithUrl("https://twitter.com/mcreedy"))
	assert.NoError(t, err)
	profile2, err := NewProfile(WithUrl("https://twitter.com/mcreedytest"))
	assert.NoError(t, err)
	requestJson := "{\"emails\":[\"marianrd97@outlook.com\",\"test1@gmail.com\",\"test2@outlook.com\"],\"phones\":[\"123-4567890\"],\"maids\":[\"abcd-123-abcd-1234-abcdlkjhasdfgh\",\"1234-snbk-lkldiemvmruixp-2kdp-vdm\"],\"location\":{\"addressLine1\":\"123/23\",\"addressLine2\":\"Some Street\",\"city\":\"Denver\",\"region\":\"Denver\",\"regionCode\":\"123123\",\"postalCode\":\"23124\"},\"name\":{\"given\":\"Marian\",\"family\":\"Reed\",\"full\":\"Marian C Reed\"},\"profiles\":[{\"url\":\"https://twitter.com/mcreedy\"},{\"url\":\"https://twitter.com/mcreedytest\"}]}"
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(NewPersonName(WithFull("Marian C Reed"), WithFamily("Reed"), WithGiven("Marian"))),
		WithEmailForMultifieldRequest("marianrd97@outlook.com"),
		WithEmailsForMultifieldRequest(emails),
		WithPhoneForMultifieldRequest("123-4567890"),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))),
		WithProfileForMultifieldRequest(profile1),
		WithProfileForMultifieldRequest(profile2),
		WithMaidsForMultifieldRequest("abcd-123-abcd-1234-abcdlkjhasdfgh"),
		WithMaidsForMultifieldRequest("1234-snbk-lkldiemvmruixp-2kdp-vdm"))
	pr, err := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest))
	assert.NoError(t, err)
	reqBytes, err := json.Marshal(pr.Query)
	assert.NoError(t, err)
	assert.Equal(t, requestJson, string(reqBytes))
}

func TestNewPermissionRequestForVerify(t *testing.T) {
	emails := []string{"test1@gmail.com", "test2@outlook.com"}
	profile1, err := NewProfile(WithUrl("https://twitter.com/mcreedy"))
	assert.NoError(t, err)
	profile2, err := NewProfile(WithUrl("https://twitter.com/mcreedytest"))
	assert.NoError(t, err)
	requestJson := "{\"query\":{\"emails\":[\"marianrd97@outlook.com\",\"test1@gmail.com\",\"test2@outlook.com\"],\"phones\":[\"123-4567890\"],\"maids\":[\"abcd-123-abcd-1234-abcdlkjhasdfgh\",\"1234-snbk-lkldiemvmruixp-2kdp-vdm\"],\"location\":{\"addressLine1\":\"123/23\",\"addressLine2\":\"Some Street\",\"city\":\"Denver\",\"region\":\"Denver\",\"regionCode\":\"123123\",\"postalCode\":\"23124\"},\"name\":{\"given\":\"Marian\",\"family\":\"Reed\",\"full\":\"Marian C Reed\"},\"profiles\":[{\"url\":\"https://twitter.com/mcreedy\"},{\"url\":\"https://twitter.com/mcreedytest\"}]},\"purposeId\":8,\"channel\":\"web\"}"
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(NewPersonName(WithFull("Marian C Reed"), WithFamily("Reed"), WithGiven("Marian"))),
		WithEmailForMultifieldRequest("marianrd97@outlook.com"),
		WithEmailsForMultifieldRequest(emails),
		WithPhoneForMultifieldRequest("123-4567890"),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))),
		WithProfileForMultifieldRequest(profile1),
		WithProfileForMultifieldRequest(profile2),
		WithMaidsForMultifieldRequest("abcd-123-abcd-1234-abcdlkjhasdfgh"),
		WithMaidsForMultifieldRequest("1234-snbk-lkldiemvmruixp-2kdp-vdm"))
	pr, err := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithPurposeIdForPermission(8),
		WithChannelForPermission("web"))
	assert.NoError(t, err)
	reqBytes, err := json.Marshal(pr)
	assert.NoError(t, err)
	assert.Equal(t, requestJson, string(reqBytes))
}

func TestNewPermissionRequestWithoutNameAndLocation(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithEmailForMultifieldRequest("marianrd97@outlook.com"))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithNameOnlyWithQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithEmailForMultifieldRequest("marianrd97@outlook.com"),
		WithNameForMultifieldRequest(&PersonName{
			Full: "Marian C Reed",
		}))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithNameOnlyWithoutQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{
			Full: "Marian C Reed",
		}))
	err := validateForPermissionFind(multifieldRequest)
	assert.EqualError(t, err, "FullContactError: If you want to use 'location'(or placekey) or 'name' as an input, both must be present and they must have non-blank values")
}

func TestNewPermissionRequestWithLocationOnlyWithQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithEmailForMultifieldRequest("marianrd97@outlook.com"),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationOnlyWithoutQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.EqualError(t, err, "FullContactError: If you want to use 'location'(or placekey) or 'name' as an input, both must be present and they must have non-blank values")
}

func TestNewPermissionRequestWithLocationWithoutAddressLine1WithQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithEmailForMultifieldRequest("marianrd97@outlook.com"),
		WithNameForMultifieldRequest(NewPersonName(WithFull("Test Name"))),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationWithoutAddressLine1WithoutQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(NewPersonName(WithFull("Test Name"))),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("Denver"),
			WithRegionCode("123123"),
			WithPostalCode("23124"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.EqualError(t, err, "FullContactError: A valid placekey is required or Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPermissionRequestWithLocationOnlyAddressLine1WithQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithEmailForMultifieldRequest("marianrd97@outlook.com"),
		WithNameForMultifieldRequest(NewPersonName(WithFull("Test Name"))),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationOnlyAddressLine1WithoutQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(NewPersonName(WithFull("Test Name"))),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.EqualError(t, err, "FullContactError: A valid placekey is required or Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPermissionRequestWithLocationWithAddressLine1AndCityWithQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithEmailForMultifieldRequest("marianrd97@outlook.com"),
		WithNameForMultifieldRequest(NewPersonName(WithFull("Test Name"))),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationWithAddressLine1AndCityWithoutQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(NewPersonName(WithFull("Test Name"))),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.EqualError(t, err, "FullContactError: A valid placekey is required or Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPermissionRequestWithLocationWithAddressLine1AndRegionWithQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithEmailForMultifieldRequest("marianrd97@outlook.com"),
		WithNameForMultifieldRequest(NewPersonName(WithFull("Test Name"))),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithRegionCode("123123"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithLocationWithAddressLine1AndRegionWithoutQueryable(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(NewPersonName(WithFull("Test Name"))),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithRegionCode("123123"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.EqualError(t, err, "FullContactError: A valid placekey is required or Location data requires addressLine1 and postalCode or addressLine1, city and regionCode (or region)")
}

func TestNewPermissionRequestWithValidLocation1(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Full: "Marian C Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("12343"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithValidLocation2(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Full: "Marian C Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithCity("Denver"),
			WithRegionCode("123123"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithValidLocation3(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Full: "Marian C Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithAddressLine2("Some Street"),
			WithCity("Denver"),
			WithRegionForLocation("123123"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithValidName(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestWithValidNameWithValidPlacekey(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithPlacekeyForMultifieldRequest("226@5z4-zvy-ffz"))
	err := validateForPermissionFind(multifieldRequest)
	assert.NoError(t, err)
}

func TestNewPermissionRequestForCreateWithConsentPurpose(t *testing.T) {
	consentPurposes := NewConsentPurpose(
		WithConsentPurposeId(1),
		WithConsentPurposeChannel("web"),
		WithConsentPurposeTtl(365),
		WithConsentPurposeEnabled(true))
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(consentPurposes),
		WithCollectionMethodForPermission("cookiePopUp"),
		WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		WithPolicyUrlForPermission("http://foo.baz"),
		WithTermsServiceForPermission("http://foo.tos"))
	err := validateForPermissionCreate(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestForCreateWithoutConsentPurposeId(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(NewConsentPurpose(
			WithConsentPurposeChannel("web"),
			WithConsentPurposeTtl(365),
			WithConsentPurposeEnabled(true))),
		WithCollectionMethodForPermission("cookiePopUp"),
		WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		WithPolicyUrlForPermission("http://foo.baz"),
		WithTermsServiceForPermission("http://foo.tos"))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Purpose id is required for consentPurpose")
}

func TestNewPermissionRequestForCreateWithoutConsentPurposeChannel(t *testing.T) {
	consentPurposes := NewConsentPurpose(
		WithConsentPurposeId(1),
		WithConsentPurposeTtl(365),
		WithConsentPurposeEnabled(true))
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(consentPurposes),
		WithCollectionMethodForPermission("cookiePopUp"),
		WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		WithPolicyUrlForPermission("http://foo.baz"),
		WithTermsServiceForPermission("http://foo.tos"))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Channel is required for consentPurpose")
}

func TestNewPermissionRequestForCreateWithoutConsentPurposeEnabled(t *testing.T) {
	consentPurposes := NewConsentPurpose(
		WithConsentPurposeId(1),
		WithConsentPurposeTtl(365),
		WithConsentPurposeEnabled(true))
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(consentPurposes),
		WithCollectionMethodForPermission("cookiePopUp"),
		WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		WithPolicyUrlForPermission("http://foo.baz"),
		WithTermsServiceForPermission("http://foo.tos"))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Channel is required for consentPurpose")
}

func TestNewPermissionRequestForCreateWithoutConsentPurposeTtl(t *testing.T) {
	consentPurposes := NewConsentPurpose(
		WithConsentPurposeId(1),
		WithConsentPurposeChannel("web"),
		WithConsentPurposeEnabled(true))
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	_, err := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(consentPurposes),
		WithCollectionMethodForPermission("cookiePopUp"),
		WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		WithPolicyUrlForPermission("http://foo.baz"),
		WithTermsServiceForPermission("http://foo.tos"))
	assert.NoError(t, err)
}

func TestNewPermissionRequestForCreateWithoutCollectionMethod(t *testing.T) {
	consentPurposes := NewConsentPurpose(
		WithConsentPurposeId(1),
		WithConsentPurposeChannel("web"),
		WithConsentPurposeTtl(365),
		WithConsentPurposeEnabled(true))
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(consentPurposes),
		WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		WithPolicyUrlForPermission("http://foo.baz"),
		WithTermsServiceForPermission("http://foo.tos"))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Collection Method is required for PermissionRequest")
}

func TestNewPermissionRequestForCreateWithoutCollectionLocation(t *testing.T) {
	consentPurposes := NewConsentPurpose(
		WithConsentPurposeId(1),
		WithConsentPurposeChannel("web"),
		WithConsentPurposeTtl(365),
		WithConsentPurposeEnabled(true))
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(consentPurposes),
		WithCollectionMethodForPermission("cookiePopUp"),
		WithPolicyUrlForPermission("http://foo.baz"),
		WithTermsServiceForPermission("http://foo.tos"))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Collection Location is required for PermissionRequest")
}

func TestNewPermissionRequestForCreateWithoutPolicyUrl(t *testing.T) {
	consentPurposes := NewConsentPurpose(
		WithConsentPurposeId(1),
		WithConsentPurposeChannel("web"),
		WithConsentPurposeTtl(365),
		WithConsentPurposeEnabled(true))
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(consentPurposes),
		WithCollectionMethodForPermission("cookiePopUp"),
		WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		WithTermsServiceForPermission("http://foo.tos"))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Policy URL is required for PermissionRequest")
}

func TestNewPermissionRequestForCreateWithoutTermsService(t *testing.T) {
	consentPurposes := NewConsentPurpose(
		WithConsentPurposeId(1),
		WithConsentPurposeChannel("web"),
		WithConsentPurposeTtl(365),
		WithConsentPurposeEnabled(true))
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithConsentPurposeForPermission(consentPurposes),
		WithCollectionMethodForPermission("cookiePopUp"),
		WithCollectionLocationForPermission("Can we get a snapshot of where someone is opting in/out here?"),
		WithPolicyUrlForPermission("http://foo.baz"))
	err := validateForPermissionCreate(pr)
	assert.EqualError(t, err, "FullContactError: Terms of Service is required for PermissionRequest")
}

func TestNewPermissionRequestForVerifyWithPurposeIdAndChannel(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithPurposeIdForPermission(1),
		WithChannelForPermission("email"))
	err := validateForPermissionVerify(pr)
	assert.NoError(t, err)
}

func TestNewPermissionRequestForVerifyWithoutPurposeId(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithChannelForPermission("email"))
	err := validateForPermissionVerify(pr)
	assert.EqualError(t, err, "FullContactError: Purpose ID is required for PermissionRequest")
}

func TestNewPermissionRequestForVerifyWithoutChannel(t *testing.T) {
	multifieldRequest, _ := NewMultifieldRequest(
		WithNameForMultifieldRequest(&PersonName{Given: "Marian", Family: "Reed"}),
		WithLocationForMultifieldRequest(NewLocation(
			WithAddressLine1("123/23"),
			WithPostalCode("23432"))))
	pr, _ := NewPermissionRequest(
		WithMultifieldRequestForPermission(multifieldRequest),
		WithPurposeIdForPermission(1))
	err := validateForPermissionVerify(pr)
	assert.EqualError(t, err, "FullContactError: Channel is required for PermissionRequest")
}

func TestNilPermissionCreateRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionCreate(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission Request can't be nil")
}

func TestNilPermissionDeleteRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionDelete(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission Request can't be nil")
}

func TestNilPermissionFindRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionFind(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission Request can't be nil")
}

func TestNilPermissionCurrentRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionCurrent(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission Request can't be nil")
}

func TestNilPermissionVerifyRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.PermissionVerify(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Permission Request can't be nil")
}
