package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func TestEmailVerification(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"status\":200,\"requestId\":\"10887309-1e58-461a-89ad-09d8e80a8f35\",\"emails\":{\"bart@fullcontact.com\":{\"message\":\"High Risk (Complainer, Fraudulent)\",\"address\":\"bart@fullcontact.com\",\"username\":\"bart\",\"domain\":\"fullcontact.com\",\"corrected\":false,\"attributes\":{\"validSyntax\":true,\"deliverable\":false,\"catchall\":false,\"risky\":true,\"disposable\":false},\"person\":\"https://api.fullcontact.com/v2/person.json?email=bart@fullcontact.com&apiKey=\",\"company\":\"https://api.fullcontact.com/v2/company/lookup.json?domain=fullcontact.com&apiKey=\",\"sendSafely\":false}}}"
	fcTestClient, testServer := getTestServerAndClient(emailVerificationUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.EmailVerificationResponse

	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "High Risk (Complainer, Fraudulent)", response.Emails["bart@fullcontact.com"].Message)
	assert.Equal(t, "bart@fullcontact.com", response.Emails["bart@fullcontact.com"].Address)
	assert.Equal(t, "bart", response.Emails["bart@fullcontact.com"].Username)
	assert.Equal(t, "fullcontact.com", response.Emails["bart@fullcontact.com"].Domain)
	assert.False(t, response.Emails["bart@fullcontact.com"].Corrected)
	assert.False(t, response.Emails["bart@fullcontact.com"].SendSafely)
	assert.True(t, response.Emails["bart@fullcontact.com"].Attributes.ValidSyntax)
	assert.True(t, response.Emails["bart@fullcontact.com"].Attributes.Risky)
	assert.False(t, response.Emails["bart@fullcontact.com"].Attributes.Deliverable)
	assert.False(t, response.Emails["bart@fullcontact.com"].Attributes.Catchall)
	assert.False(t, response.Emails["bart@fullcontact.com"].Attributes.Disposable)
}
