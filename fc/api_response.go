package fullcontact

import (
	"fmt"
	"net/http"
)

type APIResponse struct {
	RawHttpResponse           *http.Response
	PersonResponse            *PersonResp
	CompanyResponse           *CompanyResponse
	ResolveResponse           *ResolveResponse
	VerifySignalsResponse     *VerifySignalsResponse
	VerifyMatchResponse       *VerifyMatchResponse
	VerifyActivityResponse    *VerifyActivityResponse
	ResolveResponseWithTags   *ResolveResponseWithTags
	TagsResponse              *TagsResponse
	AudienceResponse          *AudienceResponse
	PermissionFindResponse    []*PermissionFindResponse
	PermissionCurrentResponse map[string]map[string]ConsentPurposeResponse
	PermissionVerifyResponse  *ConsentPurposeResponse
	StatusCode                int
	Status                    string
	IsSuccessful              bool
	Err                       error
}

func (resp *APIResponse) String() string {
	return fmt.Sprintf("\nRawHttpResponse: %v,\nPersonResponse: %v,\nCompanyResponse: %v,"+
		"\nResolveResponse: %v,\nResolveResponseWithTags: %v,\nTagsResponse: %v,\nAudienceResponse: %v,"+
		"\nPermissionFindResponse: %v,\nPermissionCurrentResponse: %v,\nPermissionVerifyResponse: %v,"+
		"\nVerifySignalsResponse: %v,\nVerifyMatchResponse: %v,\nVerifyActivityResponse: %v,"+
		"\nStatusCode: %v,\nStatus: %v,\nIsSuccessful: %v,\nErr: %v\n",
		resp.RawHttpResponse, resp.PersonResponse, resp.CompanyResponse, resp.ResolveResponse,
		resp.ResolveResponseWithTags, resp.TagsResponse, resp.AudienceResponse,
		resp.PermissionFindResponse, resp.PermissionCurrentResponse, resp.PermissionVerifyResponse,
		resp.VerifySignalsResponse, resp.VerifyMatchResponse, resp.VerifyActivityResponse,
		resp.StatusCode, resp.Status, resp.IsSuccessful, resp.Err)
}
