package fullcontact

type EmailVerificationResponse struct {
	Status        int                        `json:"status"`
	RequestId     string                     `json:"requestId"`
	UnknownEmails string                     `json:"unknownEmails"`
	FailedEmails  string                     `json:"FailedEmails"`
	Emails        map[string]EmailProperties `json:"emails"`
}

type EmailProperties struct {
	Message    string          `json:"message"`
	Address    string          `json:"address"`
	Username   string          `json:"username"`
	Domain     string          `json:"domain"`
	Person     string          `json:"person"`
	Company    string          `json:"company"`
	Corrected  bool            `json:"corrected"`
	SendSafely bool            `json:"sendSafely"`
	Attributes EmailAttributes `json:"attributes"`
}

type EmailAttributes struct {
	ValidSyntax bool `json:"validSyntax"`
	Deliverable bool `json:"deliverable"`
	Catchall    bool `json:"catchall"`
	Risky       bool `json:"risky"`
	Disposable  bool `json:"disposable"`
}
