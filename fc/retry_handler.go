package fullcontact

type RetryHandler interface {
	ShouldRetry(responseCode int) bool
	RetryAttempts() int
	RetryDelayMillis() int
}

type DefaultRetryHandler struct{}

func (drh *DefaultRetryHandler) ShouldRetry(responseCode int) bool {
	if responseCode == 429 || responseCode == 503 {
		return true
	}
	return false
}

func (drh *DefaultRetryHandler) RetryAttempts() int {
	return 1
}

func (drh *DefaultRetryHandler) RetryDelayMillis() int {
	return 1000
}
