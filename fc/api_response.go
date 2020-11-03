package fullcontact

import (
	"fmt"
	"net/http"
)

type APIResponse struct {
	RawHttpResponse         *http.Response
	PersonResponse          *PersonResp
	CompanyResponse         *CompanyResponse
	CompanySearchResponse   []*CompanySearchResponse
	ResolveResponse         *ResolveResponse
	ResolveResponseWithTags *ResolveResponseWithTags
	StatusCode              int
	Status                  string
	IsSuccessful            bool
	Err                     error
}

func (resp *APIResponse) String() string {
	return fmt.Sprintf("\nRawHttpResponse: %v,\nPersonResponse: %v,\nCompanyResponse: %v,\nCompanySearchResponse: %v,"+
		"\nResolveResponse: %v,\nResolveResponseWithTags: %v,\nStatusCode: %v,\nStatus: %v,\nIsSuccessful: %v,\nErr: %v\n",
		resp.RawHttpResponse, resp.PersonResponse, resp.CompanyResponse, resp.CompanySearchResponse, resp.ResolveResponse,
		resp.ResolveResponseWithTags, resp.StatusCode, resp.Status, resp.IsSuccessful, resp.Err)
}
