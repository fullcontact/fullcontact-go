package fullcontact

import (
	"net/http"
	"time"
)

type ClientOption func(fc *fullContactClient)

type fullContactClient struct {
	credentialsProvider  CredentialsProvider
	connectTimeoutMillis int
	headers              map[string]string
	httpClient           *http.Client
	retryHandler         RetryHandler
}

func NewFullContactClient(options ...ClientOption) (*fullContactClient, error) {
	c := &fullContactClient{
		headers:              make(map[string]string),
		connectTimeoutMillis: 0,
	}

	for _, opts := range options {
		opts(c)
	}
	if c.credentialsProvider == nil {
		cp, err := NewDefaultCredentialsProvider(FcApiKey)
		if err != nil {
			return nil, err
		}
		c.credentialsProvider = cp
	}

	if c.connectTimeoutMillis <= 0 {
		c.connectTimeoutMillis = 3000
	}

	if c.retryHandler == nil {
		c.retryHandler = &DefaultRetryHandler{}
	}

	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Timeout: time.Duration(c.connectTimeoutMillis) * time.Millisecond,
		}
	}

	return c, nil
}

func WithCredentialsProvider(credentialsProvider CredentialsProvider) ClientOption {
	return func(fc *fullContactClient) {
		fc.credentialsProvider = credentialsProvider
	}
}

func WithTimeout(timeout int) ClientOption {
	return func(fc *fullContactClient) {
		fc.connectTimeoutMillis = timeout
	}
}

func WithHeaders(headers map[string]string) ClientOption {
	return func(fc *fullContactClient) {
		for k, v := range headers {
			fc.headers[k] = v
		}
	}
}

func WithRetryHandler(retryHandler RetryHandler) ClientOption {
	return func(fc *fullContactClient) {
		fc.retryHandler = retryHandler
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(fc *fullContactClient) {
		fc.httpClient = httpClient
	}
}
