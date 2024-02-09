package fullcontact

import (
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestVerfiyActivity(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"emails\":0.21,\"online\":0.31,\"social\":0.41,\"employment\":0.51}"
	fcTestClient, testServer := getTestServerAndClient(verifyActivityUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.VerifyActivityResponse

	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.True(t, strings.Contains(resp.String(), "RawHttpRespons"))
	// Validating Activity Fields
	assert.Equal(t, 0.21, response.Emails)
	assert.Equal(t, 0.31, response.Online)
	assert.Equal(t, 0.41, response.Social)
	assert.Equal(t, 0.51, response.Employment)
}

func TestVerfiyActivityAutoRetry(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyActivityUrl, "", 429)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 429, resp.StatusCode)
}

func TestVerfiyActivitytatus400(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyActivityUrl, "", 400)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestVerfiyActivityStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyActivityUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestVerfiyActivityStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyActivityUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestVerfiyActivityStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyActivityUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestVerifyMatch(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"city\":\"household\",\"region\":\"household\",\"country\":\"household\",\"continent\":false,\"postalCode\":\"household\",\"familyName\":\"household\",\"givenName\":\"unknown\",\"phone\":\"tangled\",\"email\":\"self\",\"maid\":false,\"social\":true,\"nonId\":false,\"risk\":0.78}"
	fcTestClient, testServer := getTestServerAndClient(verifyMatchUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.VerifyMatchResponse

	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.True(t, strings.Contains(resp.String(), "RawHttpRespons"))
	// Validating Match Fields
	assert.Equal(t, "household", response.City)
	assert.Equal(t, "household", response.Region)
	assert.Equal(t, "household", response.Country)
	assert.Equal(t, "household", response.PostalCode)
	assert.Equal(t, "household", response.FamilyName)
	assert.Equal(t, "unknown", response.GivenName)
	assert.Equal(t, "tangled", response.Phone)
	assert.Equal(t, "self", response.Email)
	assert.False(t, response.Continent)
	assert.False(t, response.Maid)
	assert.True(t, response.Social)
	assert.False(t, response.NonId)
	assert.Equal(t, 0.78, response.Risk)
}

func TestVerifyMatchAutoRetry(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyMatchUrl, "", 429)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 429, resp.StatusCode)
}

func TestVerifyMatchtatus400(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyMatchUrl, "", 400)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestVerifyMatchStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyMatchUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestVerifyMatchStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyMatchUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestVerifyMatchStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifyMatchUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestVerifySignals(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"panoIds\": [ {\"id\": \"tes2ch30-pifn-cbvi-30yy-nia-zex7aw5u\",\"firstSeenMs\": 1350021600,\"lastSeenMs\": 1640415600,\"observations\": 100,\"confidence\": 0.87},{\"id\": \"tes20000-pifn-cbvi-30yy-nia-zex7aw5u\",\"firstSeenMs\": 1640415600,\"lastSeenMs\": 1640415700,\"observations\": 1000,\"confidence\": 0.99}],\"personIds\": [ \"c0VAsuEb4DRPuXmEXLutGitk-Hq9xUMautmqzyfuHpZyl3\",\"220VAsuEb4DRPuXmEXLutGitk-Hq9xUMautmqzyfuHpZyl3\"],\"phones\":[{\"label\":\"work\",\"value\":\"+19702255555\",\"firstSeenMs\":1350021600,\"lastSeenMs\":1640415600,\"observations\":100,\"confidence\":0.65},{\"label\":\"home\",\"value\":\"+19702244444\",\"firstSeenMs\":1350021500,\"lastSeenMs\":1350021600,\"observations\":99,\"confidence\":0.45}],\"emails\":[{\"md5\":\"5bacd323eae243ca2b8a84cd1c2b14aa\",\"sha1\":\"c1d89b652016ff2f5c4e2545b0f0676d9a7467aa\",\"sha256\":\"cc608758cdb416e0ebaa83f1d4f013ed98a1f5373188d2a32aff38bf51ab31ebaa\",\"firstSeenMs\":1458923376000,\"lastSeenMs\":1616590864668,\"observations\":1,\"confidence\":0.1},{\"md5\":\"2ab684f5a377204230ba72706f1d3eaa\",\"sha1\":\"93479c6277876005f4eb16f84ad07fb0381983aa\",\"sha256\":\"c32a659aa924a55d4df202ed7d2ddefc1aefe5a6e7b369e15a5c222fe58c93aa\",\"firstSeenMs\":1641932405000,\"lastSeenMs\":1653991146899,\"observations\":1,\"confidence\":1}],\"maids\":[{\"id\":\"454d83f8-516f-4f7d-ad1d-f0794e9d684\",\"type\":\"idfa\",\"firstSeenMs\":1627516800000,\"lastSeenMs\":1650982797000,\"observations\":2,\"confidence\":0.1},{\"id\":\"d2d91bf4-efe1-4e22-b10e-698b284c5b4\",\"type\":\"aaid\",\"firstSeenMs\":1633005310000,\"lastSeenMs\":1648771200000,\"observations\":3,\"confidence\":0.1}],\"name\":{\"givenName\":\"Jane\",\"familyName\":\"Doe\"},\"nonIds\":[{\"id\":\"o5baZ5TFr20zRO9gKnyWzocvGaD-i8JphXwC6g\",\"firstSeenMs\":1646438400000,\"lastSeenMs\":1646438400000,\"observations\":2,\"confidence\":0.1},{\"id\":\"-uyC3knqTJ5HZyWhQuI8gZzPH_Ts4_nBAyN1sQ\",\"firstSeenMs\":1648583171000,\"lastSeenMs\":1650982797000,\"observations\":2,\"confidence\":0.1}],\"ipAddresses\":[{\"id\":\"100.100.100.100\",\"firstSeenMs\":1627516800000,\"lastSeenMs\":1650982797000,\"confidence\":0.1},{\"id\":\"100.100.100.101\",\"firstSeenMs\":1633005310000,\"lastSeenMs\":1648771200000,\"confidence\":0.1}],\"socialProfiles\":{\"twitterUrl\":\"https://twitter.com/JaneDoeFullContact\",\"linkedInUrl\":\"https://www.linkedin.com/in/JaneDoeFullContact\"},\"demographics\":{\"age\":33,\"ageRange\":\"30-39\",\"locationFormatted\":\"Denver, Colorado, United States\",\"gender\":\"Female\"},\"employment\":{\"current\":true,\"company\":\"FullContact Inc\",\"title\":\"Quality Assurance Ghost\"}}"
	fcTestClient, testServer := getTestServerAndClient(verifySignalsUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.VerifySignalsResponse

	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.True(t, strings.Contains(resp.String(), "RawHttpRespons"))
	// Validating PersonIds Fields
	assert.Equal(t, 2, len(response.PersonIds))
	assert.Equal(t, "c0VAsuEb4DRPuXmEXLutGitk-Hq9xUMautmqzyfuHpZyl3", response.PersonIds[0])
	assert.Equal(t, "220VAsuEb4DRPuXmEXLutGitk-Hq9xUMautmqzyfuHpZyl3", response.PersonIds[1])
	// Validating Email Fields
	assert.Equal(t, 2, len(response.Emails))
	assert.Equal(t, "5bacd323eae243ca2b8a84cd1c2b14aa", response.Emails[0].Md5)
	assert.Equal(t, "c1d89b652016ff2f5c4e2545b0f0676d9a7467aa", response.Emails[0].Sha1)
	assert.Equal(t, "cc608758cdb416e0ebaa83f1d4f013ed98a1f5373188d2a32aff38bf51ab31ebaa", response.Emails[0].Sha256)
	assert.Equal(t, int64(1458923376000), response.Emails[0].FirstSeenMs)
	assert.Equal(t, int64(1616590864668), response.Emails[0].LastSeenMs)
	assert.Equal(t, 1, response.Emails[0].Observations)
	assert.Equal(t, 0.1, response.Emails[0].Confidence)
	assert.Equal(t, "2ab684f5a377204230ba72706f1d3eaa", response.Emails[1].Md5)
	assert.Equal(t, "93479c6277876005f4eb16f84ad07fb0381983aa", response.Emails[1].Sha1)
	assert.Equal(t, "c32a659aa924a55d4df202ed7d2ddefc1aefe5a6e7b369e15a5c222fe58c93aa", response.Emails[1].Sha256)
	assert.Equal(t, int64(1641932405000), response.Emails[1].FirstSeenMs)
	assert.Equal(t, int64(1653991146899), response.Emails[1].LastSeenMs)
	assert.Equal(t, 1, response.Emails[1].Observations)
	assert.Equal(t, 1.0, response.Emails[1].Confidence)
	// Validating Phone Fields
	assert.Equal(t, 2, len(response.Phones))
	assert.Equal(t, "work", response.Phones[0].Label)
	assert.Equal(t, "+19702255555", response.Phones[0].Value)
	assert.Equal(t, int64(1350021600), response.Phones[0].FirstSeenMs)
	assert.Equal(t, int64(1640415600), response.Phones[0].LastSeenMs)
	assert.Equal(t, 100, response.Phones[0].Observations)
	assert.Equal(t, 0.65, response.Phones[0].Confidence)
	assert.Equal(t, "home", response.Phones[1].Label)
	assert.Equal(t, "+19702244444", response.Phones[1].Value)
	assert.Equal(t, int64(1350021500), response.Phones[1].FirstSeenMs)
	assert.Equal(t, int64(1350021600), response.Phones[1].LastSeenMs)
	assert.Equal(t, 99, response.Phones[1].Observations)
	assert.Equal(t, 0.45, response.Phones[1].Confidence)
	// Validating Maids Fields
	assert.Equal(t, 2, len(response.Maids))
	assert.Equal(t, "454d83f8-516f-4f7d-ad1d-f0794e9d684", response.Maids[0].Id)
	assert.Equal(t, "idfa", response.Maids[0].Type)
	assert.Equal(t, int64(1627516800000), response.Maids[0].FirstSeenMs)
	assert.Equal(t, int64(1650982797000), response.Maids[0].LastSeenMs)
	assert.Equal(t, 2, response.Maids[0].Observations)
	assert.Equal(t, 0.1, response.Maids[0].Confidence)
	assert.Equal(t, "d2d91bf4-efe1-4e22-b10e-698b284c5b4", response.Maids[1].Id)
	assert.Equal(t, "aaid", response.Maids[1].Type)
	assert.Equal(t, int64(1633005310000), response.Maids[1].FirstSeenMs)
	assert.Equal(t, int64(1648771200000), response.Maids[1].LastSeenMs)
	assert.Equal(t, 3, response.Maids[1].Observations)
	assert.Equal(t, 0.1, response.Maids[1].Confidence)
	// Validating PanoIds Fields
	assert.Equal(t, 2, len(response.PanoIds))
	assert.Equal(t, "tes2ch30-pifn-cbvi-30yy-nia-zex7aw5u", response.PanoIds[0].Id)
	assert.Equal(t, int64(1350021600), response.PanoIds[0].FirstSeenMs)
	assert.Equal(t, int64(1640415600), response.PanoIds[0].LastSeenMs)
	assert.Equal(t, 100, response.PanoIds[0].Observations)
	assert.Equal(t, 0.87, response.PanoIds[0].Confidence)
	assert.Equal(t, "tes20000-pifn-cbvi-30yy-nia-zex7aw5u", response.PanoIds[1].Id)
	assert.Equal(t, int64(1640415600), response.PanoIds[1].FirstSeenMs)
	assert.Equal(t, int64(1640415700), response.PanoIds[1].LastSeenMs)
	assert.Equal(t, 1000, response.PanoIds[1].Observations)
	assert.Equal(t, 0.99, response.PanoIds[1].Confidence)
	// Validating Name
	assert.Equal(t, "Jane", response.Name.GivenName)
	assert.Equal(t, "Doe", response.Name.FamilyName)
	// Validating NonIds
	assert.Equal(t, 2, len(response.NonIds))
	assert.Equal(t, "o5baZ5TFr20zRO9gKnyWzocvGaD-i8JphXwC6g", response.NonIds[0].Id)
	assert.Equal(t, int64(1646438400000), response.NonIds[0].FirstSeenMs)
	assert.Equal(t, int64(1646438400000), response.NonIds[0].LastSeenMs)
	assert.Equal(t, 2, response.NonIds[0].Observations)
	assert.Equal(t, 0.1, response.NonIds[0].Confidence)
	assert.Equal(t, "-uyC3knqTJ5HZyWhQuI8gZzPH_Ts4_nBAyN1sQ", response.NonIds[1].Id)
	assert.Equal(t, int64(1648583171000), response.NonIds[1].FirstSeenMs)
	assert.Equal(t, int64(1650982797000), response.NonIds[1].LastSeenMs)
	assert.Equal(t, 2, response.NonIds[1].Observations)
	assert.Equal(t, 0.1, response.NonIds[1].Confidence)
	// Validating IPAddresses
	assert.Equal(t, 2, len(response.IpAddresses))
	assert.Equal(t, "100.100.100.100", response.IpAddresses[0].Id)
	assert.Equal(t, int64(1627516800000), response.IpAddresses[0].FirstSeenMs)
	assert.Equal(t, int64(1650982797000), response.IpAddresses[0].LastSeenMs)
	assert.Equal(t, 0.1, response.IpAddresses[0].Confidence)
	assert.Equal(t, "100.100.100.101", response.IpAddresses[1].Id)
	assert.Equal(t, int64(1633005310000), response.IpAddresses[1].FirstSeenMs)
	assert.Equal(t, int64(1648771200000), response.IpAddresses[1].LastSeenMs)
	assert.Equal(t, 0.1, response.IpAddresses[1].Confidence)
	// Validating Social Profiles
	assert.Equal(t, "https://twitter.com/JaneDoeFullContact", response.SocialProfiles.TwitterUrl)
	assert.Equal(t, "https://www.linkedin.com/in/JaneDoeFullContact", response.SocialProfiles.LinkedInUrl)
	// Validating Demographics
	assert.Equal(t, 33, response.Demographics.Age)
	assert.Equal(t, "30-39", response.Demographics.AgeRange)
	assert.Equal(t, "Female", response.Demographics.Gender)
	assert.Equal(t, "Denver, Colorado, United States", response.Demographics.LocationFormatted)
	// Validating Employment
	assert.Equal(t, true, response.Employment.Current)
	assert.Equal(t, "FullContact Inc", response.Employment.Company)
	assert.Equal(t, "Quality Assurance Ghost", response.Employment.Title)

}

func TestVerifySignalsAutoRetry(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifySignalsUrl, "", 429)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 429, resp.StatusCode)
}

func TestVerifySignalstatus400(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifySignalsUrl, "", 400)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestVerifySignalsStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifySignalsUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestVerifySignalsStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifySignalsUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestVerifySignalsStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(verifySignalsUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}
