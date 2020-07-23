package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestIdentityMap(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"recordIds\": [\"21c300bcf16b079ae52025cc1c06765c\"]}"
	fcTestClient, testServer := getTestServerAndClient(identityMapUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.ResolveResponse
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, "21c300bcf16b079ae52025cc1c06765c", response.RecordIds[0])
}

func TestIdentityResolve(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"recordIds\":[\"customer123\"],\"personIds\":[\"VS1OPPPPvxHcCNPezUbvYBCDEAOdSj5AI0adsA2bLmh12345\"]}"
	fcTestClient, testServer := getTestServerAndClient(identityResolveUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.ResolveResponse
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, "customer123", response.RecordIds[0])
	assert.Equal(t, "VS1OPPPPvxHcCNPezUbvYBCDEAOdSj5AI0adsA2bLmh12345", response.PersonIds[0])
}

func TestIdentityDelete(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(identityDeleteUrl, "", 204)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 204, resp.StatusCode)
	assert.Equal(t, "204 No Content", resp.Status)
}

func TestResolveWithAutoRetry(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(identityMapUrl, "", 429)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 429, resp.StatusCode)
}

func TestIdentityMapStatus400(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(identityMapUrl, "", 400)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestIdentityMapStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(identityMapUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestIdentityMapStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(identityMapUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestIdentityMapStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(identityMapUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}
