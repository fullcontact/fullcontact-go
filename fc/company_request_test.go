package fullcontact

import (
	"encoding/json"
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestNewCompanyRequestForEnrich(t *testing.T) {
	req, err := NewCompanyRequest(WithDomain("fullcontact.com"))
	assert.NoError(t, err)
	reqJson := "{\"domain\":\"fullcontact.com\"}"
	reqBytes, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.Equal(t, reqJson, string(reqBytes))
}

func TestNewCompanyRequestForSearch(t *testing.T) {
	req, err := NewCompanyRequest(
		WithCompanyName("Full Contact"),
		WithLocality("Denver"),
		WithRegion("Colorado"),
		WithWebhookUrlForCompany("http://www.fullcontact.com/hook"),
		WithCountry("US"),
		WithLocationForCompany("Denver, CO"),
		WithSort("employees"))
	assert.NoError(t, err)
	reqJSON := "{\"companyName\":\"Full Contact\",\"webhoookUrl\":\"http://www.fullcontact.com/hook\",\"location\":\"Denver, CO\",\"locality\":\"Denver\",\"region\":\"Colorado\",\"country\":\"US\",\"sort\":\"employees\"}"
	reqBytes, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.Equal(t, reqJSON, string(reqBytes))
}

func TestNewCompanyRequestInvalidSort(t *testing.T) {
	_, err := NewCompanyRequest(
		WithCompanyName("Full Contact"),
		WithLocality("Denver"),
		WithRegion("Colorado"),
		WithWebhookUrlForCompany("http://www.fullcontact.com/hook"),
		WithCountry("US"),
		WithLocationForCompany("Denver, CO"),
		WithSort("test"))
	assert.EqualError(t, err, "FullContactError: Sort value can only be 'traffic','relevance','employees'")
}

func TestNilCompanyEnrichRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.CompanyEnrich(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Company Request can't be nil")
}

func TestNilCompanySearchRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	ch := fcTestClient.CompanySearch(nil)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Company Request can't be nil")
}

func TestInvalidCompanyEnrichRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	cr, err := NewCompanyRequest(WithCompanyName("Fullcontact"))
	assert.NoError(t, err)
	resp := <-fcTestClient.CompanyEnrich(cr)
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Company Domain is mandatory for Company Enrich")
}

func TestInvalidCompanySearchRequest(t *testing.T) {
	fcTestClient := fullContactClient{}
	cr, err := NewCompanyRequest(WithCountry("US"))
	assert.NoError(t, err)
	resp := <-fcTestClient.CompanySearch(cr)
	assert.False(t, resp.IsSuccessful)
	assert.EqualError(t, resp.Err, "FullContactError: Company Name is mandatory for Company Search")
}
