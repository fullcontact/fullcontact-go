package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestCompanyEnrich(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"name\":\"FullContact Inc.\",\"location\":\"1755 Blake Street Suite 450 Denver CO, 80202 USA\",\"twitter\":\"https://twitter.com/fullcontact\",\"linkedin\":\"https://www.linkedin.com/company/fullcontact-inc-\",\"facebook\":null,\"bio\":\"FullContact is the most powerful fully-connected contact management platform for professionals and enterprises who need to master their contacts and be awesome with people.\",\"logo\":\"https://img.fullcontact.com/static/bb796b303166bd928f6c0968f15d4a4e_7ef85b2a563abd95ae07e815da2db916a5f8de4d82702388e546a66adc9eac44\",\"website\":\"https://www.fullcontact.com\",\"founded\":2010,\"employees\":351,\"locale\":\"en\",\"category\":\"Other\",\"details\":{\"locales\":[{\"code\":\"en\",\"name\":\"English\"}],\"categories\":[{\"code\":\"OTHER\",\"name\":\"Other\"}],\"industries\":[{\"type\":\"SIC\",\"name\":\"Computer Peripheral Equipment, Nec\",\"code\":\"3577\"},{\"type\":\"SIC\",\"name\":\"Computers, Peripherals, and Software\",\"code\":\"5045\"},{\"type\":\"SIC\",\"name\":\"Computer Integrated Systems Design\",\"code\":\"7373\"}],\"emails\":[{\"value\":\"support@fullcontact.com\",\"label\":\"other\"},{\"value\":\"team@fullcontact.com\",\"label\":\"sales\"},{\"value\":\"sales@fullcontact.com\",\"label\":\"work\"}],\"phones\":[{\"value\":\"+1 (720) 475-1292\",\"label\":\"other\"},{\"value\":\"+1 (888) 330-6943\",\"label\":\"other\"},{\"value\":\"+1-888-330-6943\",\"label\":\"other\"}],\"profiles\":{\"angellist\":{\"service\":\"angellist\",\"username\":\"fullcontact\",\"userid\":\"1748\",\"url\":\"https://angel.co/fullcontact\",\"bio\":\"FullContact's address book brings all of your contacts into one place and keeps them automatically up to date on the web, as well as on your iPhone and iPad. \\n\\nAdd photos to your contacts. Find them on social networks like Twitter, Facebook, LinkedIn and of course AngelList. It's the address book that busy professionals from any walk of life can appreciate, and best of all it's free. \\n\\nFor developers, the suite of FullContact APIs builds powerful, complete profiles of contacts that can be included in any application.\",\"followers\":285},\"youtube\":{\"service\":\"youtube\",\"username\":\"FullContactAPI\",\"url\":\"https://youtube.com/user/FullContactAPI\"},\"owler\":{\"service\":\"owler\",\"username\":\"fullcontact\",\"userid\":\"106145\",\"url\":\"https://www.owler.com/iaApp/106145/fullcontact-company-profile\"},\"twitter\":{\"service\":\"twitter\",\"username\":\"fullcontact\",\"url\":\"https://twitter.com/fullcontact\"},\"crunchbasecompany\":{\"service\":\"crunchbasecompany\",\"username\":\"fullcontact\",\"url\":\"http://www.crunchbase.com/organization/fullcontact\",\"bio\":\"FullContact provides a suite of cloud-based contact management solutions for businesses, developers, and individuals.\"},\"linkedincompany\":{\"service\":\"linkedincompany\",\"username\":\"fullcontact-inc-\",\"url\":\"https://www.linkedin.com/company/fullcontact-inc-\"},\"klout\":{\"service\":\"klout\",\"username\":\"FullContact\",\"url\":\"http://klout.com/FullContact\"}},\"locations\":[{\"label\":\"work\",\"addressLine1\":\"1755 Blake Street\",\"addressLine2\":\"Suite 450\",\"city\":\"Denver\",\"region\":\"CO\",\"postalCode\":\"80202\",\"country\":\"USA\",\"formatted\":\"1755 Blake Street Suite 450 Denver CO, 80202 USA\"},{\"country\":\"United States\",\"formatted\":\"     United States\"}],\"images\":[{\"value\":\"https://img.fullcontact.com/static/0772022abcec146b2ce1804934a2dcc0_377deada9adff990884ba8269633c21f099915995a9a365908fc0f4f12c37431\",\"label\":\"twitter\"},{\"value\":\"https://img.fullcontact.com/static/1bacd7306731a30d2a9f024eeb1dcff1_94d77dcdedbfe40707ac4a75ca4f4d2978bffc20b2e33a3288ea9e4d47f5af6c\",\"label\":\"logo\"},{\"value\":\"https://img.fullcontact.com/static/2ab4d453f220d5d33558a29b95d5ef28_b151428e2f8f7f87ca0b7f870eb1799c23598700baab75c45cfb8de2810cf30f\",\"label\":\"logo\"},{\"value\":\"https://img.fullcontact.com/static/675fd3bf7507596b54c3f074eef80d07_9fb5af193721963d2547cbe30a999fda2cd446a55afd9fd537bfbd35c27bfe9d\",\"label\":\"logo\"},{\"value\":\"https://img.fullcontact.com/static/eef9e3bb8d01f4a025a2c8d1857c530c_a88841c6af751e53c9fd1b575451643c782b750f31c8354361c7fee99d5a069e\",\"label\":\"other\"},{\"value\":\"https://img.fullcontact.com/static/bb796b303166bd928f6c0968f15d4a4e_7ef85b2a563abd95ae07e815da2db916a5f8de4d82702388e546a66adc9eac44\",\"label\":\"other\"}],\"urls\":[{\"value\":\"https://www.fullcontact.com\",\"label\":\"website\"},{\"value\":\"https://www.youtube.com/watch?v=RnltbT0BKMo\",\"label\":\"youtube\"},{\"value\":\"https://www.fullcontact.com/blog\",\"label\":\"blog\"}],\"keywords\":[\"CRM\",\"Contact Management\",\"Developer APIs\",\"Information Services\",\"Services\",\"Social Media\"],\"keyPeople\":[],\"traffic\":{\"countryRank\":{\"global\":{\"rank\":88991,\"name\":\"Global\"}},\"localeRank\":{\"br\":{\"rank\":20591,\"name\":\"Brazil\"},\"in\":{\"rank\":48867,\"name\":\"India\"},\"us\":{\"rank\":24385,\"name\":\"United States\"}}}},\"dataAddOns\":[{\"id\":\"keypeople\",\"name\":\"Key People\",\"enabled\":false,\"applied\":false,\"description\":\"Displays information about people of interest at this company.\",\"docLink\":\"http://docs.fullcontact.com/api/#key-people\"}],\"updated\":\"2020-04-01\"}"
	fcTestClient, testServer := getTestServerAndClient(companyEnrichUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.CompanyResponse
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, "FullContact Inc.", response.Name)
	assert.Equal(t, "1755 Blake Street Suite 450 Denver CO, 80202 USA", response.Location)
	assert.Equal(t, "https://twitter.com/fullcontact", response.Twitter)
	assert.Equal(t, "https://www.linkedin.com/company/fullcontact-inc-", response.Linkedin)
	assert.Equal(t, "FullContact is the most powerful fully-connected contact management platform for professionals and enterprises who need to master their contacts and be awesome with people.", response.Bio)
	assert.Equal(t, "https://www.fullcontact.com", response.Website)
	assert.Equal(t, 2010, response.Founded)
	assert.Equal(t, 351, response.Employees)
	assert.Equal(t, "English", response.Details.Locales[0].Name)
	assert.Equal(t, "3577", response.Details.Industries[0].Code)
	assert.Equal(t, "team@fullcontact.com", response.Details.Emails[1].Value)
	assert.Equal(t, "+1-888-330-6943", response.Details.Phones[2].Value)
	assert.Equal(t, "https://youtube.com/user/FullContactAPI", response.Details.Profiles.Youtube.URL)
	assert.Equal(t, "Denver", response.Details.Locations[0].City)
	assert.Equal(t, "https://img.fullcontact.com/static/2ab4d453f220d5d33558a29b95d5ef28_b151428e2f8f7f87ca0b7f870eb1799c23598700baab75c45cfb8de2810cf30f", response.Details.Images[2].Value)
	assert.Equal(t, "https://www.fullcontact.com", response.Details.Urls[0].Value)
	assert.Equal(t, "Contact Management", response.Details.Keywords[1])
	assert.Equal(t, 88991, response.Details.Traffic.CountryRank.Global.Rank)
	assert.Equal(t, 24385, response.Details.Traffic.LocaleRank.Us.Rank)
	assert.Equal(t, "http://docs.fullcontact.com/api/#key-people", response.DataAddOns[0].DocLink)
}

func TestCompanySearch(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "[{\"lookupDomain\":\"fullcontact.com\",\"orgName\":\"FullContact Inc.\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://img.fullcontact.com/static/bb796b303166bd928f6c0968f15d4a4e_7ef85b2a563abd95ae07e815da2db916a5f8de4d82702388e546a66adc9eac44\",\"location\":{\"locality\":\"Denver\",\"region\":{\"name\":\"CO\"},\"country\":{\"name\":\"USA\"}}},{\"lookupDomain\":\"fullcontact.bg\",\"orgName\":\"FULLCONTACT\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://img.fullcontact.com/static/f689ec8bde5e30116c2845990b9aa783_3fab9351d45942f8ba73dc8cbcea45f0263f0d4e55391ee91f4d50f7f2908cee\",\"location\":{\"locality\":\"Sofia\",\"country\":{\"name\":\"Bulgaria\",\"code\":\"BG\"}}},{\"lookupDomain\":\"fullcontact.name\",\"orgName\":\"FullContact, Inc.\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://img.fullcontact.com/static/f5f644029c2970c80dd5d8a40eb9bc45_3fe5aecdad350930d6197bd8116fed8ee4009ba1ad0a244262d511a01842912b\"},{\"lookupDomain\":\"tryfullcontact.com\",\"orgName\":\"FullContact, Inc.\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://img.fullcontact.com/static/d82fa6b4ff10b53af598b274bb9cafbc_f739aa5936447143d4fdcb242a112c62dcfba5c13400dfbcdc9b685ade1409f8\",\"location\":{\"locality\":\"Denver\",\"region\":{\"name\":\"Colorado\",\"code\":\"CO\"},\"country\":{\"name\":\"United States\",\"code\":\"US\"}}},{\"lookupDomain\":\"fullcontact.nl\",\"orgName\":\"Fullcontact Network SDN.BHD\",\"location\":{\"country\":{\"name\":\"Netherlands\",\"code\":\"NL\"}}},{\"lookupDomain\":\"vpn-r01.brewster.com\",\"orgName\":\"FullContact, Inc.\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://d2ojpxxtu63wzl.cloudfront.net/static/21bc03dc04aed84fd385ec474a1be1be_c5b7ec055d326b51df0f3c969471986777d4d3456e3e079a46ceeb0fc47d91c6\"},{\"lookupDomain\":\"fullcontactapi.com\",\"orgName\":\"FullContact, Inc.\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://d2ojpxxtu63wzl.cloudfront.net/static/b5a14b4259f5b0f8e24823e81c90f5a8_9c359854b685f9b7b03a6322b6c901893e00af5d820a0dc3980a77eb9f3e16a2\"},{\"lookupDomain\":\"blog.fullcontact.com\",\"orgName\":\"FullContact, Inc.\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://img.fullcontact.com/static/6a43815ada39ec64b22d4589c11360b0_ac6d3981bba92f9df1c232eb98ad255a821acb3bbf66334c079b27f438b93653\",\"location\":{\"locality\":\"Denver\",\"region\":{\"name\":\"Colorado\",\"code\":\"CO\"},\"country\":{\"name\":\"United States\",\"code\":\"US\"}}},{\"lookupDomain\":\"fullcontact.net\",\"orgName\":\"FullContact, Inc.\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://img.fullcontact.com/static/675fd3bf7507596b54c3f074eef80d07_9fb5af193721963d2547cbe30a999fda2cd446a55afd9fd537bfbd35c27bfe9d\"},{\"lookupDomain\":\"prod-rr-r01.brewster.com\",\"orgName\":\"FullContact, Inc.\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://d2ojpxxtu63wzl.cloudfront.net/static/91a060eb8afdbd9c89fe23ed3a658a84_c1315c9e876cb92fee969e95c92e5fc5930d1b8f5d563d34b5e6fc4c58d8fa7e\"},{\"lookupDomain\":\"portal.fullcontact.com\",\"orgName\":\"FullContact\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://d2ojpxxtu63wzl.cloudfront.net/static/91a060eb8afdbd9c89fe23ed3a658a84_c1315c9e876cb92fee969e95c92e5fc5930d1b8f5d563d34b5e6fc4c58d8fa7e\"},{\"lookupDomain\":\"fullcontact.xxx\",\"orgName\":\"FullContact, Inc.\",\"logo\":\"https://d2ojpxxtu63wzl.cloudfront.net/v1/thumbnail?size=128&url=https://img.fullcontact.com/static/f5f644029c2970c80dd5d8a40eb9bc45_3fe5aecdad350930d6197bd8116fed8ee4009ba1ad0a244262d511a01842912b\"}]"
	fcTestClient, testServer := getTestServerAndClient(companySearchUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, "fullcontact.com", resp.CompanySearchResponse[0].LookupDomain)
	assert.Equal(t, "FullContact Inc.", resp.CompanySearchResponse[0].OrgName)
	assert.True(t, strings.Contains(resp.CompanySearchResponse[0].Logo, "img.fullcontact.com"))
	assert.Equal(t, "Denver", resp.CompanySearchResponse[0].Location.Locality)
	assert.Equal(t, "CO", resp.CompanySearchResponse[0].Location.Region.Name)
	assert.Equal(t, "USA", resp.CompanySearchResponse[0].Location.Country.Name)
}
func TestCompanyEnrichAutoRetry(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companyEnrichUrl, "", 429)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 429, resp.StatusCode)
}

func TestCompanyEnrichStatus400(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companyEnrichUrl, "", 400)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestCompanyEnrichStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companyEnrichUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestCompanyEnrichStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companyEnrichUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestCompanyEnrichStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companyEnrichUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestCompanySearchAutoRetry(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companySearchUrl, "", 429)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 429, resp.StatusCode)
}

func TestCompanySearchStatus400(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companySearchUrl, "", 400)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestCompanySearchStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companySearchUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestCompanySearchStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companySearchUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestCompanySearchStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(companySearchUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}
