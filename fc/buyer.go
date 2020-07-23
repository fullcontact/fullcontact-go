package fullcontact

type Buyer struct {
	Catalog PurchaseBehavior `json:"catalog"`
	Retail  PurchaseBehavior `json:"retail"`
}
type PurchaseBehavior struct {
	Apparel             Apparel `json:"apparel"`
	Payment             Payment `json:"payment"`
	ArtsAntiques        bool    `json:"artsAntiques"`
	Automative          bool    `json:"automative"`
	Beauty              bool    `json:"beauty"`
	Books               bool    `json:"books"`
	ChildrenProducts    bool    `json:"childrenProducts"`
	Collectibles        bool    `json:"collectibles"`
	HomeOffice          bool    `json:"homeOffice"`
	Crafts              bool    `json:"crafts"`
	Electronics         bool    `json:"electronics"`
	FoodBeverages       bool    `json:"foodBeverages"`
	Furniture           bool    `json:"furniture"`
	Garden              bool    `json:"garden"`
	GeneralMerchandise  bool    `json:"generalMerchandise"`
	Gift                bool    `json:"gift"`
	Health              bool    `json:"health"`
	Holiday             bool    `json:"holiday"`
	HomeCare            bool    `json:"homeCare"`
	HomeFurnishings     bool    `json:"homeFurnishings"`
	Housewares          bool    `json:"housewares"`
	Jewelry             bool    `json:"jewelry"`
	Linens              bool    `json:"linens"`
	Music               bool    `json:"music"`
	Novelty             bool    `json:"novelty"`
	OtherMerchServices  bool    `json:"otherMerchServices"`
	PersonalCare        bool    `json:"personalCare"`
	Pets                bool    `json:"pets"`
	PhotoVideoEquipment bool    `json:"photoVideoEquipment"`
	SpecialtyFood       bool    `json:"specialtyFood"`
	SpecialtyGifts      bool    `json:"specialtyGifts"`
	SportsLeisure       bool    `json:"sportsLeisure"`
	Stationery          bool    `json:"stationery"`
	Travel              bool    `json:"travel"`
	VideoEntertainment  bool    `json:"videoEntertainment"`
	ContinuityShopper   bool    `json:"continuityShopper"`
	OnlineShopper       bool    `json:"onlineShopper"`
}

type Apparel struct {
	General           bool `json:"general"`
	Children          bool `json:"children"`
	Men               bool `json:"men"`
	MensBigTall       bool `json:"mensBigTall"`
	NonGenderSpecific bool `json:"nonGenderSpecific"`
	Teenagers         bool `json:"teenagers"`
	Women             bool `json:"women"`
	WomenPetiteSize   bool `json:"womenPetiteSize"`
	WSomenPlusSize    bool `json:"womenPlusSize"`
}

type Payment struct {
	Amx         bool `json:"amx"`
	CreditCard  bool `json:"creditCard"`
	Discover    bool `json:"discover"`
	HouseCharge bool `json:"houseCharge"`
	MasterCard  bool `json:"masterCard"`
	RetailCard  bool `json:"retailCard"`
	Visa        bool `json:"visa"`
}
