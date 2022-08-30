package fullcontact

type Household struct {
	HomeInfo     HomeInfo     `json:"homeInfo"`
	Presence     Presence     `json:"presence"`
	Finance      Finance      `json:"finance"`
	LocationInfo LocationInfo `json:"locationInfo"`
	FamilyInfo   FamilyInfo   `json:"familyInfo"`
	ChildrenInfo ChildrenInfo `json:"childrenInfo"`
}

type HomeInfo struct {
	DwellingType                string `json:"dwellingType"`
	DwellingTypeIndicator       string `json:"dwellingTypeIndicator"`
	HouseholdEducation          string `json:"householdEducation"`
	HouseholdEducationIndicator string `json:"householdEducationIndicator"`
	MaritalStatus               string `json:"maritalStatus"`
	LengthOfResidence           string `json:"lengthOfResidence"`
	LengthOfResidenceIndicator  string `json:"lengthOfResidenceIndicator"`
	Affluents                   string `json:"affluents"`
	CurrentLoanToValue          string `json:"currentLoanToValue"`
	HomeHeatSource              string `json:"homeHeatSource"`
	HouseholdOccupation         string `json:"householdOccupation"`
	FamilyComposition           string `json:"familyComposition"`
	HomeMarketValue             string `json:"homeMarketValue"`
	OwnerToOwner                string `json:"ownerToOwner"`
	RentertoOwner               string `json:"rentertoOwner"`
	YearHomeBuilt               string `json:"yearHomeBuilt"`
}

type Presence struct {
	PresenceOfChildren          string `json:"presenceOfChildren"`
	PresenceOfChildrenIndicator string `json:"presenceOfChildrenIndicator"`
	Adult                       Adult  `json:"adult"`
	Child                       Child  `json:"child"`
}
type Adult struct {
	Age18to24  string `json:"age18to24"`
	Age25to34  string `json:"age25to34"`
	Age35to44  string `json:"age35to44"`
	Age45to54  string `json:"age45to54"`
	Age55to64  string `json:"age55to64"`
	Age65to74  string `json:"age65to74"`
	Age75above string `json:"age75above"`
}
type Child struct {
	Age0to2   string `json:"age0to2"`
	Age3to5   string `json:"age3to5"`
	Age6to10  string `json:"age6to10"`
	Age11to15 string `json:"age11to15"`
	Age16to17 string `json:"age16to17"`
}
type Finance struct {
	Income                              string `json:"income"`
	IncomeIndicator                     string `json:"incomeIndicator"`
	NarrowBandIncome                    string `json:"narrowBandIncome"`
	DiscretionarySpendingIncome         string `json:"discretionarySpendingIncome"`
	FirstMortgageAmountInThousands      string `json:"firstMortgageAmountInThousands"`
	HomeEquityLoanDate                  string `json:"homeEquityLoanDate"`
	HomeMarketValueTaxRecord            string `json:"homeMarketValueTaxRecord"`
	HomeEquityLoanInThousands           string `json:"homeEquityLoanInThousands"`
	HomeEquityLoanIndicator             string `json:"homeEquityLoanIndicator"`
	InvestmentResources                 string `json:"investmentResources"`
	LiquidResources                     string `json:"liquidResources"`
	MortgageInterestRateTypeOrRefinance string `json:"mortgageInterestRateTypeOrRefinance"`
	MortgageLiability                   string `json:"mortgageLiability"`
	MortgageLoanTypeOrRefinance         string `json:"mortgageLoanTypeOrRefinance"`
	MortgageDate                        string `json:"mortgageDate"`
	RefinanceIndicator                  string `json:"refinanceIndicator"`
	SecondMortgageAmountInThousands     string `json:"secondMortgageAmountInThousands"`
	ShortTermLiability                  string `json:"shortTermLiability"`
	IncomeIndex                         string `json:"incomeIndex"`
	NetWorth                            string `json:"netWorth"`
	WealthResources                     string `json:"wealthResources"`
	PaymentMethodCreditCard             string `json:"paymentMethodCreditCard"`
}

type LocationInfo struct {
	CarrierRoute             string `json:"carrierRoute"`
	CoreBasedStatisticalArea string `json:"coreBasedStatisticalArea"`
	NielsenCountySize        string `json:"nielsenCountySize"`
	BlockGroupNumber         string `json:"blockGroupNumber"`
	CensusTractSuffix        string `json:"censusTractSuffix"`
	CountyCode               string `json:"countyCode"`
	DsfSeasonCode            string `json:"dsfSeasonCode"`
	NielsenCountySizeCode    string `json:"nielsenCountySizeCode"`
}

type FamilyInfo struct {
	HouseholdSize    string     `json:"householdSize"`
	NumberOfAdults   string     `json:"numberOfAdults"`
	NumberOfChildren string     `json:"numberOfChildren"`
	LifeCycles       LifeCycles `json:"lifeCycles"`
}

type LifeCycles struct {
	BabyBoomers                        string `json:"babyBoomers"`
	DualIncomeNoKids                   string `json:"dualIncomeNoKids"`
	FamilyTies                         string `json:"familyTies"`
	GenerationX                        string `json:"generationX"`
	Millenials                         string `json:"millenials"`
	MillenialsButFirstLetMeTakeASelfie string `json:"millenialsButFirstLetMeTakeASelfie"`
	MillenialsGettinHitched            string `json:"millenialsGettinHitched"`
	MillenialsIAmAnAdult               string `json:"millenialsIAmAnAdult"`
	MillenialsLivesWithMom             string `json:"millenialsLivesWithMom"`
	MillenialsMomLife                  string `json:"millenialsMomLife"`
	MillenialsPuttingDownRoots         string `json:"millenialsPuttingDownRoots"`
}
