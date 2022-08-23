package fullcontact

type VerifySignalsResponse struct {
	Emails         []VerifiedEmail       `json:"emails"`
	PersonIds      []string              `json:"personIds"`
	Phones         []VerifiedPhone       `json:"phones"`
	Maids          []VerifiedIdentifier  `json:"maids"`
	Name           VerifiedName          `json:"name"`
	PanoIds        []VerifiedPanoId      `json:"panoIds"`
	NonIds         []VerifiedNonId       `json:"nonIds"`
	IpAddresses    []VerifiedIpAddress   `json:"ipAddresses"`
	SocialProfiles VerifiedSocialProfile `json:"socialProfiles"`
	Demographics   VerifiedDemographics  `json:"demographics"`
	Employment     VerifiedEmployment    `json:"employment"`
	Message        string                `json:"message"`
}

type VerifyMatchResponse struct {
	City       bool    `json:"city"`
	Region     bool    `json:"region"`
	Country    bool    `json:"country"`
	Continent  bool    `json:"continent"`
	PostalCode bool    `json:"postalCode"`
	FamilyName bool    `json:"familyName"`
	GivenName  bool    `json:"givenName"`
	Phone      bool    `json:"phone"`
	Email      bool    `json:"email"`
	Maid       bool    `json:"maid"`
	Social     bool    `json:"social"`
	NonId      bool    `json:"nonId"`
	Risk       float64 `json:"risk"`
}

type VerifyActivityResponse struct {
	Emails     float64 `json:"emails"`
	Online     float64 `json:"online"`
	Social     float64 `json:"social"`
	Employment float64 `json:"employment"`
}

type VerifiedEmail struct {
	Md5          string  `json:"md5"`
	Sha1         string  `json:"sha1"`
	Sha256       string  `json:"sha256"`
	FirstSeenMs  int64   `json:"firstSeenMs"`
	LastSeenMs   int64   `json:"lastSeenMs"`
	Observations int     `json:"observations"`
	Confidence   float64 `json:"confidence"`
}

type VerifiedPhone struct {
	Label        string  `json:"label"`
	Value        string  `json:"value"`
	FirstSeenMs  int64   `json:"firstSeenMs"`
	LastSeenMs   int64   `json:"lastSeenMs"`
	Observations int     `json:"observations"`
	Confidence   float64 `json:"confidence"`
}

type VerifiedIdentifier struct {
	Id           string  `json:"id"`
	Type         string  `json:"type"`
	FirstSeenMs  int64   `json:"firstSeenMs"`
	LastSeenMs   int64   `json:"lastSeenMs"`
	Observations int     `json:"observations"`
	Confidence   float64 `json:"confidence"`
}

type VerifiedName struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

type VerifiedPanoId struct {
	Id           string  `json:"id"`
	FirstSeenMs  int64   `json:"firstSeenMs"`
	LastSeenMs   int64   `json:"lastSeenMs"`
	Observations int     `json:"observations"`
	Confidence   float64 `json:"confidence"`
}

type VerifiedNonId struct {
	Id           string  `json:"id"`
	FirstSeenMs  int64   `json:"firstSeenMs"`
	LastSeenMs   int64   `json:"lastSeenMs"`
	Observations int     `json:"observations"`
	Confidence   float64 `json:"confidence"`
}

type VerifiedDemographics struct {
	Age               int    `json:"age"`
	AgeRange          string `json:"ageRange"`
	Gender            string `json:"gender"`
	LocationFormatted string `json:"locationFormatted"`
}

type VerifiedSocialProfile struct {
	TwitterUrl  string `json:"twitterUrl"`
	LinkedInUrl string `json:"linkedInUrl"`
}

type VerifiedIpAddress struct {
	Id          string  `json:"id"`
	FirstSeenMs int64   `json:"firstSeenMs"`
	LastSeenMs  int64   `json:"lastSeenMs"`
	Confidence  float64 `json:"confidence"`
}

type VerifiedEmployment struct {
	Current bool   `json:"current"`
	Company string `json:"company"`
	Title   string `json:"title"`
}
