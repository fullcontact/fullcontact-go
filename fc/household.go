package fullcontact

type Household struct {
	HomeInfo     HomeInfo     `json:"homeInfo"`
	Presence     Presence     `json:"presence"`
	Finance      Finance      `json:"finance"`
	LocationInfo LocationInfo `json:"locationInfo"`
	FamilyInfo   FamilyInfo   `json:"familyInfo"`
}

type HomeInfo struct {
	HomeValueEstimate   int    `json:"homeValueEstimate"`
	LoanToValueEstimate int    `json:"loanToValueEstimate"`
	YearsInHome         int    `json:"yearsInHome"`
	DwellingType        string `json:"dwellingType"`
}

type Presence struct {
	Children                  string `json:"children"`
	MultigenerationalResident string `json:"multigenerationalResident"`
}

type Finance struct {
	CashValueBalanceHouseholdEstimate string `json:"cashValueBalanceHouseholdEstimate"`
	FinancialDebtRangeEstimate        string `json:"financialDebtRangeEstimate"`
	HouseholdIncomeEstimate           string `json:"income"`
	NetWorthRange                     string `json:"netWorthRange"`
	BankCard                          string `json:"bankCard"`
	RetailCard                        string `json:"retailCard"`
	ActiveLineOfCredit                bool   `json:"activeLineOfCredit"`
	Bankruptcy                        bool   `json:"bankruptcy"`
	DiscretionaryIncomeEstimate       int    `json:"discretionaryIncomeEstimate"`
}

type LocationInfo struct {
	CarrierRoute             string `json:"carrierRoute"`
	DesignatedMarketArea     string `json:"designatedMarketArea"`
	CoreBasedStatisticalArea string `json:"coreBasedStatisticalArea"`
	NielsenCountySize        string `json:"nielsenCountySize"`
	CongressionalDistrict    string `json:"congressionalDistrict"`
	NumericCountyCode        int    `json:"numericCountyCode"`
	SeasonalAddress          bool   `json:"seasonalAddress"`
}

type FamilyInfo struct {
	TotalAdults            int `json:"totalAdults"`
	TotalChildren          int `json:"totalChildren"`
	TotalPeopleInHousehold int `json:"totalPeopleInHousehold"`
}
