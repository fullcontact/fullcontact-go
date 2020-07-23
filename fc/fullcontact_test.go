package fullcontact

import (
	"io"
	"net/http"
	"net/http/httptest"
)

//Remember to close the returned Test Server
func getTestServerAndClient(testType, respJson string, statusCode int) (fullContactClient, *httptest.Server) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(FCGoClientTestType, testType)
		w.WriteHeader(statusCode)
		io.WriteString(w, respJson)
	}))

	fcTestClient := fullContactClient{
		credentialsProvider: StaticCredentialsProvider{apiKey: "apikey"},
		httpClient:          &http.Client{},
		retryHandler:        &DefaultRetryHandler{}}
	return fcTestClient, testServer
}
