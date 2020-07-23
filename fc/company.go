package fullcontact

type CompanyResponse struct {
	Name       string          `json:"name"`
	Location   string          `json:"location"`
	Twitter    string          `json:"twitter"`
	Linkedin   string          `json:"linkedin"`
	Bio        string          `json:"bio"`
	Logo       string          `json:"logo"`
	Website    string          `json:"website"`
	Locale     string          `json:"locale"`
	Category   string          `json:"category"`
	Updated    string          `json:"updated"`
	Founded    int             `json:"founded"`
	Employees  int             `json:"employees"`
	DataAddOns []DataAddOns    `json:"dataAddOns"`
	Details    *CompanyDetails `json:"details"`
}

type CompanySearchResponse struct {
	LookupDomain string                `json:"lookupDomain"`
	OrgName      string                `json:"orgName"`
	Logo         string                `json:"logo"`
	Location     CompanySearchLocation `json:"location"`
}

type CompanySearchLocation struct {
	Locality string        `json:"locality"`
	Region   CompanyRegion `json:"region"`
	Country  CompanyRegion `json:"country"`
}

type CompanyRegion struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type DataAddOns struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DocLink     string `json:"docLink"`
}

type CompanyDetails struct {
	Locales    []Locale          `json:"locales"`
	Category   []CompanyCategory `json:"categories"`
	Entity     Entity            `json:"entity"`
	Industries []Industry        `json:"industries"`
	Emails     []Email           `json:"emails"`
	Phones     []Phone           `json:"phones"`
	Locations  []*Location       `json:"locations"`
	Images     []Photo           `json:"images"`
	Urls       []Url             `json:"urls"`
	Keywords   []string          `json:"keywords"`
	KeyPeople  []People          `json:"keyPeople"`
	Traffic    CompanyTraffic    `json:"traffic"`
	Profiles   Profiles          `json:"profiles"`
}

type Locale struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type CompanyCategory struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Entity struct {
	Name      string `json:"name"`
	Founded   int    `json:"founded"`
	Employees int    `json:"employees"`
}

type Industry struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Code string `json:"code"`
}

type People struct {
	FullName string `json:"fullName"`
	Title    string `json:"title"`
	Avatar   string `json:"avatar"`
}

type CompanyTraffic struct {
	CountryRank TrafficRank `json:"countryRank"`
	LocaleRank  TrafficRank `json:"localeRank"`
}

type TrafficRank struct {
	Global Rank `json:"global"`
	Us     Rank `json:"us"`
	In     Rank `json:"in"`
	Gb     Rank `json:"gb"`
}

type Rank struct {
	Rank int    `json:"rank"`
	Name string `json:"name"`
}
