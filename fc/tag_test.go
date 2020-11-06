package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestValidNewTagsRequest(t *testing.T) {
	_, err := NewTagsRequest(WithRecordIdForTags("k1"),
		WithTag(NewTag(WithTagKey("key"), WithTagValue("value"))))
	assert.NoError(t, err)
}

func TestNewTagsRequestInvalid1(t *testing.T) {
	_, err := NewTagsRequest(WithTag(NewTag(WithTagKey("key"), WithTagValue("value"))))
	assert.EqualError(t, err, "FullContactError: RecordId must be present for creating Tags")
}

func TestNewTagsRequestInvalid2(t *testing.T) {
	_, err := NewTagsRequest(WithRecordIdForTags("k1"),
		WithTag(NewTag(WithTagValue("value"))))
	assert.EqualError(t, err, "FullContactError: Both Key and Value must be populated for adding a Tag")
}

func TestNewTagsRequestInvalid3(t *testing.T) {
	_, err := NewTagsRequest(WithRecordIdForTags("k1"),
		WithTag(NewTag(WithTagKey("key"))))
	assert.EqualError(t, err, "FullContactError: Both Key and Value must be populated for adding a Tag")
}

func TestTagsCreate(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"recordId\":\"k3\",\"tags\":[{\"key\":\"gender\",\"value\":\"female\"}]}"
	fcTestClient, testServer := getTestServerAndClient(tagsCreateUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.TagsResponse

	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, "k3", response.RecordId)
	assert.Equal(t, "gender", response.Tags[0].Key)
	assert.Equal(t, "female", response.Tags[0].Value)
}

func TestTagsGet(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"recordId\":\"k2\",\"partnerId\":null,\"tags\":[{\"key\":\"gender\",\"value\":\"male\"},{\"key\":\"gender\",\"value\":\"female\"}]}"
	fcTestClient, testServer := getTestServerAndClient(tagsGetUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.TagsResponse

	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, "k2", response.RecordId)
	assert.Equal(t, "gender", response.Tags[0].Key)
	assert.Equal(t, "male", response.Tags[0].Value)
	assert.Equal(t, "", response.PartnerId)
}

func TestTagsDelete(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(tagsDeleteUrl, "", 204)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch

	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 204, resp.StatusCode)
	assert.Equal(t, "204 No Content", resp.Status)
}

func TestTagsCreateStatus400(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(tagsCreateUrl, "", 400)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestTagsCreateStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(tagsCreateUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestTagsCreateStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(tagsCreateUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}

func TestTagsGetStatus404(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"status\":404,\"message\":\"No records found for identifier: k1\"}"
	fcTestClient, testServer := getTestServerAndClient(tagsGetUrl, respJson, 404)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Equal(t, "404 Not Found", resp.Status)
}
