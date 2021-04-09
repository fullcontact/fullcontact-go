package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestPermissionCreate(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCreateUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
	assert.Equal(t, "202 Accepted", resp.Status)
}

func TestPermissionDelete(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionDeleteUrl, "", 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
}

func TestPermissionCurrent(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"1\":{\"phone\":{\"ttl\":365,\"enabled\":true,\"channel\":\"phone\",\"purposeId\":1,\"purposeName\":\"Information storage & access\",\"timestamp\":1617962540547},\"web\":{\"ttl\":365,\"enabled\":true,\"channel\":\"web\",\"purposeId\":1,\"purposeName\":\"Information storage & access\",\"timestamp\":1617962540547}},\"2\":{\"mobile\":{\"ttl\":365,\"enabled\":true,\"channel\":\"mobile\",\"purposeId\":2,\"purposeName\":\"Personalized Ads Profile\",\"timestamp\":1617962540547}}}"
	fcTestClient, testServer := getTestServerAndClient(permissionCurrentUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.PermissionCurrentResponse
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, 365, response["1"]["phone"].Ttl)
	assert.Equal(t, true, response["1"]["phone"].Enabled)
	assert.Equal(t, "phone", response["1"]["phone"].Channel)
	assert.Equal(t, 1, response["1"]["phone"].PurposeId)
	assert.Equal(t, "Information storage & access", response["1"]["phone"].PurposeName)
	assert.Equal(t, 1617962540547, response["1"]["phone"].Timestamp)
	assert.Equal(t, 365, response["1"]["web"].Ttl)
	assert.Equal(t, true, response["1"]["web"].Enabled)
	assert.Equal(t, "web", response["1"]["web"].Channel)
	assert.Equal(t, 1, response["1"]["web"].PurposeId)
	assert.Equal(t, "Information storage & access", response["1"]["web"].PurposeName)
	assert.Equal(t, 1617962540547, response["1"]["web"].Timestamp)
	assert.Equal(t, true, response["2"]["mobile"].Enabled)
	assert.Equal(t, "mobile", response["2"]["mobile"].Channel)
	assert.Equal(t, 2, response["2"]["mobile"].PurposeId)
	assert.Equal(t, "Personalized Ads Profile", response["2"]["mobile"].PurposeName)
	assert.Equal(t, 1617962540547, response["2"]["mobile"].Timestamp)
}

func TestPermissionFind(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "[{\"permissionType\":\"create\",\"permissionId\":\"1c99f4fb-96a2-46f4-8fd7-64750a591e05\",\"consentPurposes\":[{\"ttl\":365,\"enabled\":true,\"channel\":\"web\",\"purposeId\":1,\"purposeName\":\"Information storage & access\",\"timestamp\":1617628580297}],\"locale\":null,\"ipAddress\":null,\"language\":null,\"collectionMethod\":\"cookiePopUp\",\"collectionLocation\":\"Can we get a snapshot of where someone is opting in/out here?\",\"policyUrl\":\"https://www.fullcontact.com/privacy/privacy-policy\",\"termsService\":\"https://www.fullcontact.com/privacy/terms-of-use\",\"timestamp\":null,\"created\":1617628580297}]"
	fcTestClient, testServer := getTestServerAndClient(permissionFindUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.PermissionFindResponse
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, "create", response[0].PermissionType)
	assert.Equal(t, "1c99f4fb-96a2-46f4-8fd7-64750a591e05", response[0].PermissionId)
	assert.Equal(t, 365, response[0].ConsentPurposes[0].Ttl)
	assert.Equal(t, true, response[0].ConsentPurposes[0].Enabled)
	assert.Equal(t, "web", response[0].ConsentPurposes[0].Channel)
	assert.Equal(t, 1, response[0].ConsentPurposes[0].PurposeId)
	assert.Equal(t, "Information storage & access", response[0].ConsentPurposes[0].PurposeName)
	assert.Equal(t, 1617628580297, response[0].ConsentPurposes[0].Timestamp)
	assert.Equal(t, "", response[0].Locale)
	assert.Equal(t, "", response[0].IpAddress)
	assert.Equal(t, "", response[0].Language)
	assert.Equal(t, "cookiePopUp", response[0].CollectionMethod)
	assert.Equal(t, "Can we get a snapshot of where someone is opting in/out here?", response[0].CollectionLocation)
	assert.Equal(t, "https://www.fullcontact.com/privacy/privacy-policy", response[0].PolicyUrl)
	assert.Equal(t, "https://www.fullcontact.com/privacy/terms-of-use", response[0].TermsService)
	assert.Equal(t, 0, response[0].Timestamp)
	assert.Equal(t, 1617628580297, response[0].Created)
}

func TestPermissionVerify(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"ttl\":365,\"enabled\":true,\"channel\":\"web\",\"purposeId\":1,\"purposeName\":\"Information storage & access\",\"timestamp\":1617962540547}"
	fcTestClient, testServer := getTestServerAndClient(permissionVerifyUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.PermissionVerifyResponse
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, 365, response.Ttl)
	assert.Equal(t, true, response.Enabled)
	assert.Equal(t, "web", response.Channel)
	assert.Equal(t, 1, response.PurposeId)
	assert.Equal(t, "Information storage & access", response.PurposeName)
	assert.Equal(t, 1617962540547, response.Timestamp)
}

func TestPermissionCreateStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCreateUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestPermissionCreateStatus204(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCreateUrl, "", 204)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestPermissionCreateStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCreateUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestPermissionCreateStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCreateUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestPermissionCreateStatus404(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCreateUrl, "", 404)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestPermissionCreateStatus500(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCreateUrl, "", 500)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestPermissionDeleteStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionDeleteUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestPermissionDeleteStatus204(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionDeleteUrl, "", 204)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestPermissionDeleteStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionDeleteUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestPermissionDeleteStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionDeleteUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestPermissionDeleteStatus404(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionDeleteUrl, "", 404)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestPermissionDeleteStatus500(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionDeleteUrl, "", 500)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestPermissionFindStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionFindUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestPermissionFindStatus204(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionFindUrl, "", 204)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestPermissionFindStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionFindUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestPermissionFindStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionFindUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestPermissionFindStatus404(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionFindUrl, "", 404)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestPermissionFindStatus500(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionFindUrl, "", 500)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestPermissionCurrentStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCurrentUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestPermissionCurrentStatus204(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCurrentUrl, "", 204)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestPermissionCurrentStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCurrentUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestPermissionCurrentStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCurrentUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestPermissionCurrentStatus404(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCurrentUrl, "", 404)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestPermissionCurrentStatus500(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionCurrentUrl, "", 500)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestPermissionVerifyStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionVerifyUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestPermissionVerifyStatus204(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionVerifyUrl, "", 204)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestPermissionVerifyStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionVerifyUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestPermissionVerifyStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionVerifyUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestPermissionVerifyStatus404(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionVerifyUrl, "", 404)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestPermissionVerifyStatus500(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(permissionVerifyUrl, "", 500)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 500, resp.StatusCode)
}
