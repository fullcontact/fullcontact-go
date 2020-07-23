package fullcontact

/* This contains all the Person centric response models for the Person Enrich API */

type PersonResp struct {
	FullName     string   `json:"fullName"`
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	AgeRange     string   `json:"ageRange"`
	Gender       string   `json:"gender"`
	Location     string   `json:"location"`
	Title        string   `json:"title"`
	Organization string   `json:"organization"`
	Twitter      string   `json:"twitter"`
	Linkedin     string   `json:"linkedin"`
	Bio          string   `json:"bio"`
	Avatar       string   `json:"avatar"`
	Website      string   `json:"website"`
	Details      *Details `json:"details"`
	Updated      string   `json:"updated"`
}

type Details struct {
	Name         *PersonName  `json:"name"`
	Age          Age          `json:"age"`
	Gender       string       `json:"gender"`
	Household    *Household   `json:"household"`
	Demographics Demographics `json:"demographics"`
	Finance      Finance      `json:"finance"`
	Emails       []Email      `json:"emails"`
	Phones       []Phone      `json:"phones"`
	Census       *Census      `json:"census"`
	Survey       *Survey      `json:"survey"`
	Buyer        Buyer        `json:"buyer"`
	Profiles     *Profiles    `json:"profiles"`
	Identifiers  Identifiers  `json:"identifiers"`
	Automotive   Automotive   `json:"automotive"`
	Locations    []Location   `json:"locations"`
	Employment   []Employment `json:"employment"`
	Photos       []Photo      `json:"photos"`
	Education    []Education  `json:"education"`
	Urls         []Url        `json:"urls"`
	Interests    []Interest   `json:"interests"`
}

type Demographics struct {
	Gender        string `json:"gender"`
	MaritalStatus string `json:"MaritalStatus"`
	Occupation    string `json:"occupation"`
	LivingStatus  string `json:"livingStatus"`
	Age           Age    `json:"age"`
}

type Email struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	Type   string `json:"type"`
	Md5    string `json:"md5"`
	Sha256 string `json:"sha256"`
}

type Phone struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	Type   string `json:"type"`
	Md5    string `json:"md5"`
	Sha256 string `json:"sha256"`
}

type Employment struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Domain  string `json:"domain"`
	Current bool   `json:"current"`
	Start   Date   `json:"start"`
	End     Date   `json:"end"`
}

type Photo struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type Date struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

type Education struct {
	Name   string `json:"name"`
	Degree string `json:"degree"`
	End    Date   `json:"end"`
	Start  Date   `json:"start"`
}

type Url struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type Interest struct {
	Name      string   `json:"name"`
	Id        string   `json:"id"`
	Affinity  string   `json:"affinity"`
	ParentIds []string `json:"parentIds"`
	Category  string   `json:"category"`
}

type Profiles struct {
	Twitter         ProfileData `json:"twitter"`
	Linkedin        ProfileData `json:"linkedin"`
	Linkedincompany ProfileData `json:"linkedinCompany"`
	Klout           ProfileData `json:"klout"`
	Youtube         ProfileData `json:"youtube"`
	Angellist       ProfileData `json:"angellist"`
	Owler           ProfileData `json:"owler"`
	Pinterest       ProfileData `json:"pinterest"`
}

type ProfileData struct {
	URL       string  `json:"url"`
	Username  string  `json:"username"`
	UserId    string  `json:"userid"`
	Service   string  `json:"service"`
	Bio       string  `json:"bio"`
	Followers int     `json:"followers"`
	Following int     `json:"following"`
	Photos    []Photo `json:"photos"`
	Urls      []Url   `json:"urls"`
}

type Age struct {
	Range    string `json:"range"`
	Value    int    `json:"value"`
	Birthday Date   `json:"birthday"`
}

type Identifiers struct {
	Maids     []Maids  `json:"maids"`
	PersonIds []string `json:"personIds"`
	RecordIds []string `json:"recordIds"`
}

type Maids struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type Automotive struct {
	Ownership Ownership `json:"ownership"`
}

type Ownership struct {
	TotalCars string    `json:"totalCars"`
	Vehicles  []Vehicle `json:"vehicles"`
}

type Vehicle struct {
	BodyStyle    string `json:"bodyStyle"`
	FuelType     string `json:"fuelType"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	PurchaseType string `json:"purchaseType"`
	PurchaseDate int    `json:"purchaseDate"`
	Year         int    `json:"year"`
}
