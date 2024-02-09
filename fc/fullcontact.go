package fullcontact

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func (fcClient *fullContactClient) newHttpRequest(url string, reqBytes []byte) (*http.Request, error) {
	var method string
	var buffer *bytes.Buffer

	if isHttpGet(url) {
		method = "GET"
		url = url + "?" + string(reqBytes)
		buffer = nil
	} else {
		method = "POST"
		buffer = bytes.NewBuffer(reqBytes)
	}

	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}
	req = fcClient.addHeaders(req)
	return req, nil

}

func (fcClient *fullContactClient) addHeaders(req *http.Request) *http.Request {
	for k, v := range fcClient.headers {
		req.Header.Add(k, v)
	}
	req.Header.Add("Authorization", "Bearer "+fcClient.credentialsProvider.getApiKey())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", userAgent)
	return req
}

func isHttpGet(url string) bool {
	// Add urls to below list for HTTP GET request
	getUrlList := []string{audienceDownloadUrl}

	for _, getUrl := range getUrlList {
		if url == getUrl {
			return true
		}
	}
	return false
}

func (fcClient *fullContactClient) do(url string, reqBytes []byte, ch chan *APIResponse) {
	req, err := fcClient.newHttpRequest(url, reqBytes)
	if err != nil {
		sendToChannel(ch, nil, url, err)
	}

	resp, err := fcClient.httpClient.Do(req) //first attempt

	if err != nil {
		fcClient.autoRetry(ch, err, resp, 0, url, reqBytes)
	} else if resp != nil && !fcClient.retryHandler.ShouldRetry(resp.StatusCode) {
		sendToChannel(ch, resp, url, nil)
	} else {
		fcClient.autoRetry(ch, nil, resp, 0, url, reqBytes)
	}
}

func (fcClient *fullContactClient) autoRetry(ch chan *APIResponse, err error, resp *http.Response, retryAttemptsDone int, url string, reqBytes []byte) {
	if retryAttemptsDone < min(fcClient.retryHandler.RetryAttempts(), 5) {
		retryAttemptsDone++
		time.Sleep(time.Duration(fcClient.retryHandler.RetryDelayMillis()*(1<<(retryAttemptsDone-1))) * time.Millisecond)
		req, err := fcClient.newHttpRequest(url, reqBytes)
		if err != nil {
			sendToChannel(ch, nil, url, err)
		}
		resp, err = fcClient.httpClient.Do(req)
		if err != nil {
			fcClient.autoRetry(ch, err, resp, retryAttemptsDone, url, reqBytes)
		} else if resp != nil && !fcClient.retryHandler.ShouldRetry(resp.StatusCode) {
			sendToChannel(ch, resp, url, nil)
		} else {
			fcClient.autoRetry(ch, nil, resp, retryAttemptsDone, url, reqBytes)
		}
	} else if err != nil {
		sendToChannel(ch, nil, url, err)
	} else {
		sendToChannel(ch, resp, url, nil)
	}

}

func sendToChannel(ch chan *APIResponse, response *http.Response, url string, err error) {
	apiResponse := &APIResponse{
		RawHttpResponse: response,
		Err:             err,
	}

	if response != nil {
		//For Testing Purposes
		testType := response.Header.Get(FCGoClientTestType)
		if isPopulated(testType) {
			url = testType
		}

		switch url {
		case personEnrichUrl:
			setPersonResponse(apiResponse)
		case companyEnrichUrl:
			setCompanyResponse(apiResponse)
		case identityMapUrl, identityResolveUrl, identityMapResolveUrl, identityDeleteUrl:
			setResolveResponse(apiResponse)
		case identityResolveWithTagsUrl:
			setResolveResponseWithTags(apiResponse)
		case tagsCreateUrl, tagsGetUrl, tagsDeleteUrl:
			setTagsResponse(apiResponse)
		case audienceCreateUrl, audienceDownloadUrl:
			setAudienceResponse(apiResponse)
		case permissionCreateUrl:
			setPermissionCreateResponse(apiResponse)
		case permissionDeleteUrl:
			setPermissionDeleteResponse(apiResponse)
		case permissionFindUrl:
			setPermissionFindResponse(apiResponse)
		case permissionCurrentUrl:
			setPermissionCurrentResponse(apiResponse)
		case permissionVerifyUrl:
			setPermissionVerifyResponse(apiResponse)
		case verifySignalsUrl:
			setVerfiySignalsResponse(apiResponse)
		case verifyMatchUrl:
			setVerfiyMatchResponse(apiResponse)
		case verifyActivityUrl:
			setVerfiyActivityResponse(apiResponse)
		}
	}
	ch <- apiResponse
	return
}

/* FullContact V3 Person Enrich API, takes an PersonRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) PersonEnrich(personRequest *PersonRequest) chan *APIResponse {
	ch := make(chan *APIResponse)

	if personRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Person Request can't be nil"))
		return ch
	}
	err := validatePersonRequest(personRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}

	reqBytes, err := json.Marshal(personRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	// Send Asynchronous Request in Goroutine
	go fcClient.do(personEnrichUrl, reqBytes, ch)
	return ch
}

/* FullContact V3 Company Enrich API, takes an CompanyRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) CompanyEnrich(companyRequest *CompanyRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if companyRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Company Request can't be nil"))
		return ch
	}
	err := validateForCompanyEnrich(companyRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	reqBytes, err := json.Marshal(companyRequest)

	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	// Send Asynchronous Request in Goroutine
	go fcClient.do(companyEnrichUrl, reqBytes, ch)
	return ch
}

/* Resolve
FullContact Resolve API - IdentityMap, takes an ResolveRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) IdentityMap(resolveRequest *ResolveRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if resolveRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Resolve Request can't be nil"))
		return ch
	}
	err := validateForIdentityMap(resolveRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	return fcClient.resolveRequest(ch, resolveRequest, identityMapUrl)
}

/* Resolve
FullContact Resolve API - IdentityResolve, takes an ResolveRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) IdentityResolve(resolveRequest *ResolveRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if resolveRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Resolve Request can't be nil"))
		return ch
	}
	err := validateForIdentityResolve(resolveRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	return fcClient.resolveRequest(ch, resolveRequest, identityResolveUrl)
}

/* Resolve
FullContact Resolve API - IdentityMapResolve, takes an ResolveRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) IdentityMapResolve(resolveRequest *ResolveRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if resolveRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Resolve Request can't be nil"))
		return ch
	}
	err := validateForIdentityMap(resolveRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	return fcClient.resolveRequest(ch, resolveRequest, identityMapResolveUrl)
}

/* Resolve
FullContact Resolve API - IdentityResolve with Tags in response, takes an ResolveRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) IdentityResolveWithTags(resolveRequest *ResolveRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if resolveRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Resolve Request can't be nil"))
		return ch
	}
	err := validateForIdentityResolve(resolveRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	return fcClient.resolveRequest(ch, resolveRequest, identityResolveWithTagsUrl)
}

/* Resolve
FullContact Resolve API - IdentityDelete, takes an ResolveRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) IdentityDelete(resolveRequest *ResolveRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if resolveRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Resolve Request can't be nil"))
		return ch
	}
	err := validateForIdentityDelete(resolveRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	return fcClient.resolveRequest(ch, resolveRequest, identityDeleteUrl)
}

func (fcClient *fullContactClient) resolveRequest(ch chan *APIResponse, resolveRequest *ResolveRequest, url string) chan *APIResponse {
	reqBytes, err := json.Marshal(resolveRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	// Send Asynchronous Request in Goroutine
	go fcClient.do(url, reqBytes, ch)
	return ch
}

/* FullContact API for adding/creating tags for any recordId in your PIC, takes a TagsRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) TagsCreate(tagsRequest *TagsRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if tagsRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Tags Request can't be nil"))
		return ch
	}
	reqBytes, err := json.Marshal(tagsRequest)

	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	// Send Asynchronous Request in Goroutine
	go fcClient.do(tagsCreateUrl, reqBytes, ch)
	return ch
}

/* FullContact API for getting all tags for any recordId in your PIC, takes a 'recordId' and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) TagsGet(recordId string) chan *APIResponse {
	ch := make(chan *APIResponse)
	if !isPopulated(recordId) {
		go sendToChannel(ch, nil, "", NewFullContactError("recordId can't be nil"))
		return ch
	}
	reqBytes := []byte("{\"recordId\":\"" + recordId + "\"}")

	// Send Asynchronous Request in Goroutine
	go fcClient.do(tagsGetUrl, reqBytes, ch)
	return ch
}

/* FullContact API for deleting any tag(s) for any recordId in your PIC, takes a TagsRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) TagsDelete(tagsRequest *TagsRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if tagsRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Tags Request can't be nil"))
		return ch
	}
	reqBytes, err := json.Marshal(tagsRequest)

	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	// Send Asynchronous Request in Goroutine
	go fcClient.do(tagsDeleteUrl, reqBytes, ch)
	return ch
}

/* FullContact API for creating Audience based on tags from your PIC, takes a AudienceRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) AudienceCreate(audienceRequest *AudienceRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if audienceRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Audience Request can't be nil"))
		return ch
	}
	reqBytes, err := json.Marshal(audienceRequest)

	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	// Send Asynchronous Request in Goroutine
	go fcClient.do(audienceCreateUrl, reqBytes, ch)
	return ch
}

/* FullContact API for downloading Audience created using 'AudienceCreate', takes a requestId and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) AudienceDownload(requestId string) chan *APIResponse {
	ch := make(chan *APIResponse)
	if !isPopulated(requestId) {
		go sendToChannel(ch, nil, "", NewFullContactError("requestId can't be nil"))
		return ch
	}
	reqBytes := []byte("requestId=" + requestId)

	// Send Asynchronous Request in Goroutine
	go fcClient.do(audienceDownloadUrl, reqBytes, ch)
	return ch
}

/* Permission
FullContact Permission API - PermissionCreate, takes an PermissionRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) PermissionCreate(permissionRequest *PermissionRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if permissionRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Permission Request can't be nil"))
		return ch
	}
	err := validateForPermissionCreate(permissionRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}

	reqBytes, err := json.Marshal(permissionRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	// Send Asynchronous Request in Goroutine
	go fcClient.do(permissionCreateUrl, reqBytes, ch)
	return ch
}

/* FullContact Permission API - PermissionDelete, takes an PermissionRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) PermissionDelete(multifieldRequest *MultifieldRequest) chan *APIResponse {
	return fcClient.validateAndSendMultiFieldRequestAsync(permissionDeleteUrl, multifieldRequest)
}

/* FullContact Permission API - PermissionFind, takes an PermissionRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) PermissionFind(multifieldRequest *MultifieldRequest) chan *APIResponse {
	return fcClient.validateAndSendMultiFieldRequestAsync(permissionFindUrl, multifieldRequest)
}

/* FullContact Permission API - PermissionCurrent, takes an PermissionRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) PermissionCurrent(multifieldRequest *MultifieldRequest) chan *APIResponse {
	return fcClient.validateAndSendMultiFieldRequestAsync(permissionCurrentUrl, multifieldRequest)
}

/* FullContact Permission API - PermissionVerify, takes an PermissionRequest and returns a channel of type APIResponse.
Request is converted to JSON and sends a Asynchronous request */
func (fcClient *fullContactClient) PermissionVerify(permissionRequest *PermissionRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if permissionRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("Permission Request can't be nil"))
		return ch
	}
	err := validateForPermissionVerify(permissionRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}

	reqBytes, err := json.Marshal(permissionRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	// Send Asynchronous Request in Goroutine
	go fcClient.do(permissionVerifyUrl, reqBytes, ch)
	return ch
}

/*
FullContact Verify API - VerfiySignals, takes a MultiFieldRequest and returns a channel of type APIResponse.

Request is converted to JSON and sends an Asynchronous request.
Response will be avaiable in the VerifySignalsResponse field of APIResponse
*/
func (fcClient *fullContactClient) VerifySignals(multifieldRequest *MultifieldRequest) chan *APIResponse {
	return fcClient.validateAndSendMultiFieldRequestAsync(verifySignalsUrl, multifieldRequest)
}

/*
FullContact Verify API - VerfiyMatch, takes a MultiFieldRequest and returns a channel of type APIResponse.

Request is converted to JSON and sends an Asynchronous request.
Response will be avaiable in the VerifyMatchResponse field of APIResponse
*/
func (fcClient *fullContactClient) VerifyMatch(multifieldRequest *MultifieldRequest) chan *APIResponse {
	return fcClient.validateAndSendMultiFieldRequestAsync(verifyMatchUrl, multifieldRequest)
}

/*
FullContact Verify API - VerfiyActivity, takes a MultiFieldRequest and returns a channel of type APIResponse.

Request is converted to JSON and sends an Asynchronous request.
Response will be avaiable in the VerifyActivityResponse field of APIResponse
*/
func (fcClient *fullContactClient) VerifyActivity(multifieldRequest *MultifieldRequest) chan *APIResponse {
	return fcClient.validateAndSendMultiFieldRequestAsync(verifyActivityUrl, multifieldRequest)
}

/*
	This function will perform the `MultifieldRequest` validations and if
	there are no errors then it'll be marshalled and a `MultifieldRequest` will be
	made to the specified `url`

	Returns a channel frm which the request response can be obtained
*/
func (fcClient *fullContactClient) validateAndSendMultiFieldRequestAsync(url string, multifieldRequest *MultifieldRequest) chan *APIResponse {
	ch := make(chan *APIResponse)
	if multifieldRequest == nil {
		go sendToChannel(ch, nil, "", NewFullContactError("MultiFieldRequest can't be nil"))
		return ch
	}

	err := multifieldRequest.validate()
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}

	reqBytes, err := json.Marshal(multifieldRequest)
	if err != nil {
		go sendToChannel(ch, nil, "", err)
		return ch
	}
	// Send Asynchronous Request in Goroutine
	go fcClient.do(url, reqBytes, ch)
	return ch
}

func setPersonResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var person PersonResp
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &person)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
	apiResponse.PersonResponse = &person
}

func setCompanyResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var companyResponse CompanyResponse
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &companyResponse)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
	apiResponse.CompanyResponse = &companyResponse
}

func setResolveResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var resolveResponse ResolveResponse
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &resolveResponse)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 204) || (apiResponse.StatusCode == 404)
	apiResponse.ResolveResponse = &resolveResponse
}

func setResolveResponseWithTags(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var resolveResponse ResolveResponseWithTags
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &resolveResponse)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 204) || (apiResponse.StatusCode == 404)
	apiResponse.ResolveResponseWithTags = &resolveResponse
}

func setTagsResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var tagsResponse TagsResponse
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &tagsResponse)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 204) || (apiResponse.StatusCode == 404)
	apiResponse.TagsResponse = &tagsResponse
}

func setAudienceResponse(apiResponse *APIResponse) {
	contentType := apiResponse.RawHttpResponse.Header.Get("Content-Type")
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var audienceResponse AudienceResponse
	if isPopulated(string(bodyBytes)) {
		if contentType == "application/octet-stream" {
			audienceResponse.AudienceBytes = bodyBytes
		} else {
			err = json.Unmarshal(bodyBytes, &audienceResponse)
			if err != nil {
				apiResponse.Err = err
				return
			}
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
	apiResponse.AudienceResponse = &audienceResponse
}

func setPermissionCreateResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
}

func setPermissionDeleteResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
}

func setPermissionFindResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var response []*PermissionFindResponse
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
	apiResponse.PermissionFindResponse = response
}

func setPermissionVerifyResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var response ConsentPurposeResponse
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
	apiResponse.PermissionVerifyResponse = &response
}

func setPermissionCurrentResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var response map[string]map[string]ConsentPurposeResponse
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
	apiResponse.PermissionCurrentResponse = response
}

func setVerfiySignalsResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var response VerifySignalsResponse
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
	apiResponse.VerifySignalsResponse = &response
}

func setVerfiyMatchResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var response VerifyMatchResponse
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
	apiResponse.VerifyMatchResponse = &response
}

func setVerfiyActivityResponse(apiResponse *APIResponse) {
	bodyBytes, err := ioutil.ReadAll(apiResponse.RawHttpResponse.Body)
	defer apiResponse.RawHttpResponse.Body.Close()

	// Reset the buffer so that it can be re-read by the caller.
	apiResponse.RawHttpResponse.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err != nil {
		apiResponse.Err = err
		return
	}
	var response VerifyActivityResponse
	if isPopulated(string(bodyBytes)) {
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			apiResponse.Err = err
			return
		}
	}
	apiResponse.Status = apiResponse.RawHttpResponse.Status
	apiResponse.StatusCode = apiResponse.RawHttpResponse.StatusCode
	apiResponse.IsSuccessful = (apiResponse.StatusCode == 200) || (apiResponse.StatusCode == 202) || (apiResponse.StatusCode == 404)
	apiResponse.VerifyActivityResponse = &response
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
