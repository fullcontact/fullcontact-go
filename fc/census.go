package fullcontact

type Census struct {
	BasicTractNumber int      `json:"basicTractNumber"`
	BasicBlockGroup  int      `json:"basicBlockGroup"`
	Year2010         Year2010 `json:"year2010"`
}
type Year2010 struct {
	EducationLevel     string            `json:"educationLevel"`
	SocioEconomicScore int               `json:"socioEconomicScore"`
	Average            Average           `json:"average"`
	Percent            Percent           `json:"percent"`
	Median             Median            `json:"median"`
	PopulationDensity  PopulationDensity `json:"populationDensity"`
}

type Average struct {
	NumberOfCarsInHousehold float64 `json:"numberOfCarsInHousehold"`
}

type Percent struct {
	AbovePovertyLevel      int `json:"abovePovertyLevel"`
	BelowPovertyLevel      int `json:"belowPovertyLevel"`
	Black                  int `json:"black"`
	BlueCollarEmployed     int `json:"blueCollarEmployed"`
	DivorcedOrSeparated    int `json:"divorcedOrSeparated"`
	Hispanic               int `json:"hispanic"`
	HomesBuiltSince2000    int `json:"homesBuiltSince2000"`
	Homeowner              int `json:"homeowner"`
	HouseholdsWithChildren int `json:"householdsWithChildren"`
	Married                int `json:"married"`
	MobileHome             int `json:"mobileHome"`
	MovedToAreaSince2000   int `json:"movedToAreaSince2000"`
	SalariedProfessional   int `json:"salariedProfessional"`
	SingleFamilyHome       int `json:"singleFamilyHome"`
	VehicleOwnership       int `json:"vehicleOwnership"`
	White                  int `json:"white"`
}

type Median struct {
	AgeOfHouseholder       int `json:"ageOfHouseholder"`
	EffectiveBuyingIncome  int `json:"effectiveBuyingIncome"`
	HomeValue              int `json:"homeValue"`
	HouseholdIncome        int `json:"householdIncome"`
	HouseholdIncomeByState int `json:"householdIncomeByState"`
}

type PopulationDensity struct {
	CentileInState int `json:"centileInState"`
	CentileInUs    int `json:"centileInUs"`
}
