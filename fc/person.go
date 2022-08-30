package fullcontact

/* This contains all the Person centric response models for the Person Enrich API */

type PersonResp struct {
	FullName     string            `json:"fullName"`
	Email        string            `json:"email"`
	Phone        string            `json:"phone"`
	AgeRange     string            `json:"ageRange"`
	Gender       string            `json:"gender"`
	Location     string            `json:"location"`
	Title        string            `json:"title"`
	Organization string            `json:"organization"`
	Twitter      string            `json:"twitter"`
	Linkedin     string            `json:"linkedin"`
	Bio          string            `json:"bio"`
	Avatar       string            `json:"avatar"`
	Website      string            `json:"website"`
	Details      *Details          `json:"details"`
	Extended     map[string]string `json:"extended"`
	Updated      string            `json:"updated"`
}

type Details struct {
	Name         *PersonName  `json:"name"`
	Age          Age          `json:"age"`
	Gender       string       `json:"gender"`
	Household    *Household   `json:"household"`
	Demographics Demographics `json:"demographics"`
	Emails       []Email      `json:"emails"`
	Phones       []Phone      `json:"phones"`
	Profiles     *Profiles    `json:"profiles"`
	Identifiers  Identifiers  `json:"identifiers"`
	Surveys      Surveys      `json:"surveys"`
	Locations    []Location   `json:"locations"`
	Employment   []Employment `json:"employment"`
	Photos       []Photo      `json:"photos"`
	Education    []Education  `json:"education"`
	Urls         []Url        `json:"urls"`
	Interests    []Interest   `json:"interests"`
	Triggers     Triggers     `json:"triggers"`
	MarketTrends MarketTrends `json:"marketTrends"`
}

type Triggers struct {
	Type Type        `json:"type"`
	Date TriggerDate `json:"date"`
}
type Type struct {
	CollegeGraduate           string `json:"collegeGraduate"`
	EmptyNester               string `json:"emptyNester"`
	HomeMarketValue           string `json:"homeMarketValue"`
	Income                    string `json:"income"`
	NewAdultToFile            string `json:"newAdultToFile"`
	NewFirstChild0to2YearsOld string `json:"newFirstChild0to2YearsOld"`
	NewPreDriver              string `json:"newPreDriver"`
	NewYoungAdultToFile       string `json:"newYoungAdultToFile"`
	NicheSwitch               string `json:"nicheSwitch"`
	ValueScore                string `json:"valueScore"`
}

type TriggerDate struct {
	CollegeGraduateChange string `json:"collegeGraduateChange"`
	EmptyNesterChange     string `json:"emptyNesterChange"`
	FirstChildChange      string `json:"firstChildChange"`
	HomeMarketValueChange string `json:"homeMarketValueChange"`
	IncomeChange          string `json:"incomeChange"`
	NewAdultChange        string `json:"newAdultChange"`
	NewDriverChange       string `json:"newDriverChange"`
	NewYoungAdultToChange string `json:"newYoungAdultToChange"`
	NicheSwitch           string `json:"nicheSwitch"`
	ValueScoreChange      string `json:"valueScoreChange"`
}

type MarketTrends struct {
	Switchers   Switchers              `json:"switchers"`
	Seekers     Seekers                `json:"seekers"`
	Enthusiasts MarketTrendEnthusiasts `json:"enthusiasts"`
	Attendees   Attendees              `json:"attendees"`
	Buyers      Buyers                 `json:"buyers"`
	Chef        Chef                   `json:"chef"`
	Customers   Customers              `json:"customers"`
	Consumers   MarketTrendConsumers   `json:"consumers"`
	Donor       Donor                  `json:"donor"`
	Owners      Owners                 `json:"owners"`
	Travellers  Travellers             `json:"travellers"`
	Readers     Readers                `json:"readers"`
	Voters      Voters                 `json:"voters"`
	Subscribers Subscribers            `json:"subscribers"`
	Planners    Planners               `json:"planners"`
	Purchasers  Purchasers             `json:"purchasers"`
	Spenders    Spenders               `json:"spenders"`
	Shoppers    Shoppers               `json:"shoppers"`
	Stores      Stores                 `json:"stores"`
	Users       Users                  `json:"users"`
}
type Users struct {
	App                                          string `json:"app"`
	BrandMotivatedLaundry                        string `json:"brandMotivatedLaundry"`
	BrandMotivatedPersonalCareProduct            string `json:"brandMotivatedPersonalCareProduct"`
	BrandedRetailCreditCard                      string `json:"brandedRetailCreditCard"`
	ConvenienceDrivenPersonalCareProduct         string `json:"convenienceDrivenPersonalCareProduct"`
	CreditCardBalanceTransfer                    string `json:"creditCardBalanceTransfer"`
	DebitCardRewards                             string `json:"debitCardRewards"`
	DebitCard                                    string `json:"debitCard"`
	GroceryStoreApp                              string `json:"groceryStoreApp"`
	HeavyCoupon                                  string `json:"heavyCoupon"`
	MobileBanking                                string `json:"mobileBanking"`
	MobileShoppingList                           string `json:"mobileShoppingList"`
	NaturalProductPersonalCareProduct            string `json:"naturalProductPersonalCareProduct"`
	OnlineBroker                                 string `json:"onlineBroker"`
	OnlineSavings                                string `json:"onlineSavings"`
	PaperShoppingList                            string `json:"paperShoppingList"`
	PremiumNaturalPersonalCareProduct            string `json:"premiumNaturalPersonalCareProduct"`
	PriceMotivatedLaundryProduct                 string `json:"priceMotivatedLaundryProduct"`
	PriceMotivatedPersonalCareProduct            string `json:"priceMotivatedPersonalCareProduct"`
	PrimaryCellPhone                             string `json:"primaryCellPhone"`
	PublicTransportation                         string `json:"publicTransportation"`
	RestaurantApp                                string `json:"restaurantApp"`
	RestaurantLoyaltyApp                         string `json:"restaurantLoyaltyApp"`
	RewardsCardCashBack                          string `json:"rewardsCardCashBack"`
	RewardsCardOther                             string `json:"rewardsCardOther"`
	SmartPhone                                   string `json:"smartPhone"`
	SociallyActiveonFacebook                     string `json:"sociallyActiveonFacebook"`
	SociallyActiveonFacebookBrandLikers          string `json:"sociallyActiveonFacebookBrandLikers"`
	SociallyActiveonFacebookCategoryRecommenders string `json:"sociallyActiveonFacebookCategoryRecommenders"`
	SociallyActiveonPinterest                    string `json:"sociallyActiveonPinterest"`
	SociallyActiveonTwitter                      string `json:"sociallyActiveonTwitter"`
	SolarRoofingInterest                         string `json:"solarRoofingInterest"`
	TargetCartwheel                              string `json:"targetCartwheel"`
	UberOrLyft                                   string `json:"uberOrLyft"`
	VehicleServiceCenter                         string `json:"vehicleServiceCenter"`
	WalmartSavingCatcher                         string `json:"walmartSavingCatcher"`
}
type Voters struct {
	Democratic  string `json:"democratic"`
	Independent string `json:"independent"`
	Likely      string `json:"likely"`
	Republican  string `json:"republican"`
	Swing       string `json:"swing"`
}
type Travellers struct {
	Business      string `json:"business"`
	International string `json:"international"`
	RVTripTakers  string `json:"rVTripTakers"`
}
type Subscribers struct {
	CableTVPremium            string `json:"cableTVPremium"`
	FinancialHealthNewsletter string `json:"financialHealthNewsletter"`
	OnlineMagazineNewspaper   string `json:"onlineMagazineNewspaper"`
	RetailerEmail             string `json:"retailerEmail"`
	SatelliteRadio            string `json:"satelliteRadio"`
}
type Stores struct {
	StockUpAtGroceryStores string `json:"stockUpAtGroceryStores"`
	StockUpAtWalmart       string `json:"stockUpAtWalmart"`
}
type Spenders struct {
	PandemicDecreasedSpenders string `json:"pandemicDecreasedSpenders"`
	VacationSpenders          string `json:"vacationSpenders"`
}
type Shoppers struct {
	BargainHotel               string `json:"bargainHotel"`
	Bargain                    string `json:"bargain"`
	EverydayLowPrice           string `json:"everydayLowPrice"`
	FinancialInstitution       string `json:"financialInstitution"`
	Highend                    string `json:"highend"`
	MultiRetailer              string `json:"multiRetailer"`
	OneStop                    string `json:"oneStop"`
	PriceMatchers              string `json:"priceMatchers"`
	PrivateLabel               string `json:"privateLabel"`
	QuickShopAtWalmartOrTarget string `json:"quickShopAtWalmartOrTarget"`
	StockUp                    string `json:"stockUp"`
	WhatsOnSale                string `json:"whatsOnSale"`
}
type Readers struct {
	AvidBook         string `json:"avidBook"`
	BibleDevotional  string `json:"bibleDevotional"`
	Book             string `json:"book"`
	Entertainment    string `json:"entertainment"`
	Label            string `json:"label"`
	RetailerCircular string `json:"retailerCircular"`
	Romance          string `json:"romance"`
	Sports           string `json:"sports"`
}
type Purchasers struct {
	ACAHealthInsurance            string `json:"aCAHealthInsurance"`
	AutoLoan                      string `json:"autoLoan"`
	BrandDrivenHomeCleaners       string `json:"brandDrivenHomeCleaners"`
	FrequentMobile                string `json:"frequentMobile"`
	FrequentOnlineMusic           string `json:"frequentOnlineMusic"`
	GreenProduct                  string `json:"greenProduct"`
	HomeWarranty                  string `json:"homeWarranty"`
	Impulse                       string `json:"impulse"`
	MedicareAdvantagePlan         string `json:"medicareAdvantagePlan"`
	MidmarketTermLifeInsurance    string `json:"midmarketTermLifeInsurance"`
	MidmarketWholeLifeInsurance   string `json:"midmarketWholeLifeInsurance"`
	NewLuxuryVehicle              string `json:"newLuxuryVehicle"`
	NewnonLuxuryVehicle           string `json:"newnonLuxuryVehicle"`
	OrganicFood                   string `json:"organicFood"`
	OrganicProduct                string `json:"organicProduct"`
	PetInsurance                  string `json:"petInsurance"`
	RV                            string `json:"rV"`
	Vehicle                       string `json:"vehicle"`
	WebSurferBrickMortar          string `json:"webSurferBrickMortar"`
	WebandBrickMortarViewerOnline string `json:"webandBrickMortarViewerOnline"`
}
type Planners struct {
	BudgetMeal string `json:"budgetMeal"`
	Meal       string `json:"meal"`
	PreShop    string `json:"preShop"`
}
type Owners struct {
	UnderScore401k            string `json:"_401k"`
	AndroidSmartPhone         string `json:"androidSmartPhone"`
	AppleSmartPhone           string `json:"appleSmartPhone"`
	Boat                      string `json:"boat"`
	EducationSavingsPlan      string `json:"educationSavingsPlan"`
	MultiPolicyInsurance      string `json:"multiPolicyInsurance"`
	Pet                       string `json:"pet"`
	PrepaidCard               string `json:"prepaidCard"`
	SecondHome                string `json:"secondHome"`
	SmartTV                   string `json:"smartTV"`
	Tablet                    string `json:"tablet"`
	Timeshare                 string `json:"timeshare"`
	VeterinarianInfluencedPet string `json:"veterinarianInfluencedPet"`
}
type Donor struct {
	AnimalWelfare                     string `json:"animalWelfare"`
	Cancer                            string `json:"cancer"`
	ChildrensCauses                   string `json:"childrensCauses"`
	ConsistentReligious               string `json:"consistentReligious"`
	Environmental                     string `json:"environmental"`
	HighDollarOtherCausesnonReligious string `json:"highDollarOtherCausesnonReligious"`
	HighDollarReligiousCauses         string `json:"highDollarReligiousCauses"`
	LiberalCauses                     string `json:"liberalCauses"`
	University                        string `json:"university"`
	Veteran                           string `json:"veteran"`
}
type MarketTrendConsumers struct {
	AirlineUpgraders                             string `json:"airlineUpgraders"`
	AutoInsurance                                string `json:"autoInsurance"`
	AutoInsuranceAgentSold                       string `json:"autoInsuranceAgentSold"`
	AutoInsuranceCallCenterSold                  string `json:"autoInsuranceCallCenterSold"`
	CableBundlecableinternethomephone            string `json:"cableBundlecableinternethomephone"`
	CasinoGamer                                  string `json:"casinoGamer"`
	CoinsCollector                               string `json:"coinsCollector"`
	ConservativeInvestmentStyle                  string `json:"conservativeInvestmentStyle"`
	ContactlessPayApplication                    string `json:"contactlessPayApplication"`
	CordCutters                                  string `json:"cordCutters"`
	CreditCardAttritionHouseholds                string `json:"creditCardAttritionHouseholds"`
	CreditCardRevolvers                          string `json:"creditCardRevolvers"`
	CreditUnionMember                            string `json:"creditUnionMember"`
	DietConsciousHouseholds                      string `json:"dietConsciousHouseholds"`
	DoitYourselfer                               string `json:"doitYourselfer"`
	EmployerProvidedHealthInsurancePolicyHolders string `json:"employerProvidedHealthInsurancePolicyHolders"`
	FrequentOnlineMovieViewers                   string `json:"frequentOnlineMovieViewers"`
	FreshFoodDelivery                            string `json:"freshFoodDelivery"`
	FutureInvestors                              string `json:"futureInvestors"`
	Gamers                                       string `json:"gamers"`
	GigEconomyEmployees                          string `json:"gigEconomyEmployees"`
	GroceryStoreFrequenters                      string `json:"groceryStoreFrequenters"`
	HomeEntertainers                             string `json:"homeEntertainers"`
	HomeRemodelers                               string `json:"homeRemodelers"`
	HotelLoyaltyProgramMembers                   string `json:"hotelLoyaltyProgramMembers"`
	IntheMarkettoGetaHomeLoan                    string `json:"intheMarkettoGetaHomeLoan"`
	IntheMarkettoPurchaseaHome                   string `json:"intheMarkettoPurchaseaHome"`
	IntendtoPurchaseaSamsungmobiledevice         string `json:"intendtoPurchaseaSamsungmobiledevice"`
	Intendtopurchase5GService                    string `json:"intendtopurchase5GService"`
	InvestmentTrustBankingPreference             string `json:"investmentTrustBankingPreference"`
	LikelyCruiser                                string `json:"likelyCruiser"`
	LikelyMortgageRefinancers                    string `json:"likelyMortgageRefinancers"`
	LikelyPlannedGivers                          string `json:"likelyPlannedGivers"`
	LikelytoSufferfromInsomnia                   string `json:"likelytoSufferfromInsomnia"`
	LikelytoUseanInvestmentBroker                string `json:"likelytoUseanInvestmentBroker"`
	LikelytohaveaMortgage                        string `json:"likelytohaveaMortgage"`
	LongRoadTripTakers                           string `json:"longRoadTripTakers"`
	LongTermCare                                 string `json:"longTermCare"`
	LowInterestCreditCard                        string `json:"lowInterestCreditCard"`
	LowSodium                                    string `json:"lowSodium"`
	MealCombo                                    string `json:"mealCombo"`
	MealKitDelivery                              string `json:"mealKitDelivery"`
	MedicaidPotentialQualifiedHousehold          string `json:"medicaidPotentialQualifiedHousehold"`
	MedicareDualEligibleHousehold                string `json:"medicareDualEligibleHousehold"`
	MedicarePlanDPrescriptionDrugHealth          string `json:"medicarePlanDPrescriptionDrugHealth"`
	MobileBrowsers                               string `json:"mobileBrowsers"`
	MovieLoyaltyProgramMembers                   string `json:"movieLoyaltyProgramMembers"`
	NaturalGreenProductHomeCleaners              string `json:"naturalGreenProductHomeCleaners"`
	Non401kMutualFundInvestors                   string `json:"non401kMutualFundInvestors"`
	Non401kStocksBondsInvestors                  string `json:"non401kStocksBondsInvestors"`
	PandemicInRestaurantDiners                   string `json:"pandemicInRestaurantDiners"`
	PandemicLuxurySpenders                       string `json:"pandemicLuxurySpenders"`
	PandemicRiskTolerant                         string `json:"pandemicRiskTolerant"`
	PaychecktoPaycheck                           string `json:"paychecktoPaycheck"`
	PersonalTraveler                             string `json:"personalTraveler"`
	PlantoPurchaseHomeSecuritySystems            string `json:"plantoPurchaseHomeSecuritySystems"`
	PlantogetFitnessMembership                   string `json:"plantogetFitnessMembership"`
	RentersandAutoInsuranceJointPolicyHolders    string `json:"rentersandAutoInsuranceJointPolicyHolders"`
	RetailTexters                                string `json:"retailTexters"`
	RetiredbutStillWorking                       string `json:"retiredbutStillWorking"`
	SatelliteBundlesatelliteinethomeorwireless   string `json:"satelliteBundlesatelliteinethomeorwireless"`
	SelfPayHealthInsurance                       string `json:"selfPayHealthInsurance"`
	SeniorCaregivers                             string `json:"seniorCaregivers"`
	SeniorLivingSearchers                        string `json:"seniorLivingSearchers"`
	SociallyInfluenced                           string `json:"sociallyInfluenced"`
	TechnologyEarlyAdopters                      string `json:"technologyEarlyAdopters"`
	TeleMedicine                                 string `json:"teleMedicine"`
	TermLife                                     string `json:"termLife"`
	Underbanked                                  string `json:"underbanked"`
	UninsuredforHealth                           string `json:"uninsuredforHealth"`
	UpcomingRetirees65andOlder                   string `json:"upcomingRetirees65andOlder"`
	Vegetarians                                  string `json:"vegetarians"`
	VehicleDIYrs                                 string `json:"vehicleDIYrs"`
	WeeklyOnlineBankers                          string `json:"weeklyOnlineBankers"`
	WellnessHouseholdsHealth                     string `json:"wellnessHouseholdsHealth"`
	WholeLife                                    string `json:"wholeLife"`
	WiredLineVideoConnectors                     string `json:"wiredLineVideoConnectors"`
	WorkforSmallCompanyOfferingHealthInsurance   string `json:"workforSmallCompanyOfferingHealthInsurance"`
}
type Customers struct {
	AT_TCellPhoneCustomer                             string `json:"aT_TCellPhoneCustomer"`
	AmazonPrime                                       string `json:"amazonPrime"`
	Annuity                                           string `json:"annuity"`
	AutoInsurancePremiumDiscountviaTelematicsCustomer string `json:"autoInsurancePremiumDiscountviaTelematicsCustomer"`
	CateringDelivery                                  string `json:"cateringDelivery"`
	CateringPickUp                                    string `json:"cateringPickUp"`
	CertificatesofDeposit                             string `json:"certificatesofDeposit"`
	ClicktoCartHomeDelivery                           string `json:"clicktoCartHomeDelivery"`
	ClicktoCartPickUp                                 string `json:"clicktoCartPickUp"`
	CommunityBankCustomer                             string `json:"communityBankCustomer"`
	Convenience                                       string `json:"convenience"`
	Deposit                                           string `json:"deposit"`
	DirectMediaPreference                             string `json:"directMediaPreference"`
	FinancialAdvisor                                  string `json:"financialAdvisor"`
	FrequentATM                                       string `json:"frequentATM"`
	GroceryLoyaltyCard                                string `json:"groceryLoyaltyCard"`
	InterestCheckingPreference                        string `json:"interestCheckingPreference"`
	InternationalWirelessorLandline                   string `json:"internationalWirelessorLandline"`
	InternetResearchPreference                        string `json:"internetResearchPreference"`
	Lending                                           string `json:"lending"`
	LoyalFinancialInstitution                         string `json:"loyalFinancialInstitution"`
	MensBigandTallApparel                             string `json:"mensBigandTallApparel"`
	NationalBankCustomer                              string `json:"nationalBankCustomer"`
	NewRoof                                           string `json:"newRoof"`
	OnlineDeliveryRestaurant                          string `json:"onlineDeliveryRestaurant"`
	OnlinePickUpRestaurant                            string `json:"onlinePickUpRestaurant"`
	PlantoPurchaseSmartHomeProducts                   string `json:"plantoPurchaseSmartHomeProducts"`
	QSRCash                                           string `json:"qSRCash"`
	QuantumUpgrade                                    string `json:"quantumUpgrade"`
	RegionalBankCustomer                              string `json:"regionalBankCustomer"`
	RestaurantLoyaltyCard                             string `json:"restaurantLoyaltyCard"`
	SelfInsuredDental                                 string `json:"selfInsuredDental"`
	SmartHome                                         string `json:"smartHome"`
	SocialMediaPreference                             string `json:"socialMediaPreference"`
	SprintCellPhoneCustomer                           string `json:"sprintCellPhoneCustomer"`
	StudentLoan                                       string `json:"studentLoan"`
	SubscriptionorAutoShipmentClothAccessories        string `json:"subscriptionorAutoShipmentClothAccessories"`
	SubscriptionorAutoShipmentFoodorBeverage          string `json:"subscriptionorAutoShipmentFoodorBeverage"`
	SubscriptionorAutoShipmentHouseholdProduct        string `json:"subscriptionorAutoShipmentHouseholdProduct"`
	SubscriptionorAutoShipmentPersonalCare            string `json:"subscriptionorAutoShipmentPersonalCare"`
	SubscriptionorAutoShipmentPetProducts             string `json:"subscriptionorAutoShipmentPetProducts"`
	SubscriptionorAutoShipment                        string `json:"subscriptionorAutoShipment"`
	TMobileCellPhoneCustomer                          string `json:"tMobileCellPhoneCustomer"`
	VOIPLandline                                      string `json:"vOIPLandline"`
	VerizonCellPhoneCustomer                          string `json:"verizonCellPhoneCustomer"`
	WiredService                                      string `json:"wiredService"`
	WomensPlusSizeApparel                             string `json:"womensPlusSizeApparel"`
}
type Chef struct {
	Experimental   string `json:"experimental"`
	Master         string `json:"master"`
	RealIngredient string `json:"realIngredient"`
}
type Buyers struct {
	AutoInsuranceSelfServeOnline   string `json:"autoInsuranceSelfServeOnline"`
	ChristmasOrnamentsCollectibles string `json:"christmasOrnamentsCollectibles"`
	HeavyFiberFocusedFood          string `json:"heavyFiberFocusedFood"`
	HeavyGlutenFreeFood            string `json:"heavyGlutenFreeFood"`
	HeavyLowFatFood                string `json:"heavyLowFatFood"`
	OnlineHomeCleaningProduct      string `json:"onlineHomeCleaningProduct"`
	OnlineInsurance                string `json:"onlineInsurance"`
	OnlineLaundryProduct           string `json:"onlineLaundryProduct"`
	OnlinePersonalCareProduct      string `json:"onlinePersonalCareProduct"`
	OnlinePetFood                  string `json:"onlinePetFood"`
}
type Attendees struct {
	AmusementPark            string `json:"amusementPark"`
	CulturalArtsEvents       string `json:"culturalArtsEvents"`
	LiveMusicConcert         string `json:"liveMusicConcert"`
	ProfessionalSportsEvents string `json:"professionalSportsEvents"`
}
type MarketTrendEnthusiasts struct {
	ArtHouseMovie           string `json:"artHouseMovie"`
	BarandLoungeFood        string `json:"barandLoungeFood"`
	Baseball                string `json:"baseball"`
	Basketball              string `json:"basketball"`
	BrandLoyalists          string `json:"brandLoyalists"`
	BreakfastDining         string `json:"breakfastDining"`
	CarryOut                string `json:"carryOut"`
	CasualDining            string `json:"casualDining"`
	ChristianorGospelMusic  string `json:"christianorGospelMusic"`
	CigarPipe               string `json:"cigarPipe"`
	Coffee                  string `json:"coffee"`
	ConvenienceHomeCleaners string `json:"convenienceHomeCleaners"`
	CountryMusic            string `json:"countryMusic"`
	DinnerDining            string `json:"dinnerDining"`
	DiscountMovie           string `json:"discountMovie"`
	DomesticBeer            string `json:"domesticBeer"`
	ExtremeFitness          string `json:"extremeFitness"`
	FantasySports           string `json:"fantasySports"`
	FineDining              string `json:"fineDining"`
	Football                string `json:"football"`
	FreeStreaming           string `json:"freeStreaming"`
	FrequentMovie           string `json:"frequentMovie"`
	GamingnonMobileDevices  string `json:"gamingnonMobileDevices"`
	HardCider               string `json:"hardCider"`
	HardSeltzer             string `json:"hardSeltzer"`
	Hockey                  string `json:"hockey"`
	HomeShoppingNetwork     string `json:"homeShoppingNetwork"`
	ImportBeer              string `json:"importBeer"`
	Kroger                  string `json:"kroger"`
	LatinMusic              string `json:"latinMusic"`
	Liquor                  string `json:"liquor"`
	LunchDining             string `json:"lunchDining"`
	Meditation              string `json:"meditation"`
	MobileGaming            string `json:"mobileGaming"`
	OpeningWeekendMovie     string `json:"openingWeekendMovie"`
	PaidStreaming           string `json:"paidStreaming"`
	PetsAreFamily           string `json:"petsAreFamily"`
	QuickServiceRestaurant  string `json:"quickServiceRestaurant"`
	RedWine                 string `json:"redWine"`
	Soccer                  string `json:"soccer"`
	Target                  string `json:"target"`
	ValueChains             string `json:"valueChains"`
	Walmart                 string `json:"walmart"`
	WhiteWine               string `json:"whiteWine"`
	YogaPilates             string `json:"yogaPilates"`
}
type Seekers struct {
	FreshFood                  string `json:"freshFood"`
	HomeCleaningNewProduct     string `json:"homeCleaningNewProduct"`
	LaundryNewProduct          string `json:"laundryNewProduct"`
	OnlineDegreeEducation      string `json:"onlineDegreeEducation"`
	PersonalCareNewProduct     string `json:"personalCareNewProduct"`
	PremiumNaturalHomeCleaners string `json:"premiumNaturalHomeCleaners"`
	Scent                      string `json:"scent"`
	UnscentedProduct           string `json:"unscentedProduct"`
}
type Switchers struct {
	AlcoholBeverage                string `json:"alcoholBeverage"`
	BeerBrand                      string `json:"beerBrand"`
	BreakfastMeatBrand             string `json:"breakfastMeatBrand"`
	ChocolateCandyBrand            string `json:"chocolateCandyBrand"`
	CocaColaBrand                  string `json:"cocaColaBrand"`
	ColdCerealBrand                string `json:"coldCerealBrand"`
	EnergyDrink                    string `json:"energyDrink"`
	FrozenFoodBrand                string `json:"frozenFoodBrand"`
	HouseholdCleaningProductsBrand string `json:"householdCleaningProductsBrand"`
	Insurance                      string `json:"insurance"`
	Job                            string `json:"job"`
	MobilePhoneService             string `json:"mobilePhoneService"`
	NaturalCheeseSlicesBrand       string `json:"naturalCheeseSlicesBrand"`
	NaturalShreddedCheeseBrand     string `json:"naturalShreddedCheeseBrand"`
	NutritionalHealthBarBrand      string `json:"nutritionalHealthBarBrand"`
	RefrigeratedLunchMeatBrand     string `json:"refrigeratedLunchMeatBrand"`
	SnackBarGranolaBarBrand        string `json:"snackBarGranolaBarBrand"`
	SoftDrinksBrand                string `json:"softDrinksBrand"`
	SpiritsBrand                   string `json:"spiritsBrand"`
	YogurtBrand                    string `json:"yogurtBrand"`
}

type Demographics struct {
	Gender      string         `json:"gender"`
	Consumers   Consumers      `json:"consumers"`
	HomeInfo    HomeInfoPerson `json:"homeInfo"`
	Enthusiasts Enthusiasts    `json:"enthusiasts"`
	MaritalInfo MaritalInfo    `json:"maritalInfo"`
	Age         Age            `json:"age"`
}

type MaritalInfo struct {
	MaritalStatus          string `json:"maritalStatus"`
	MaritalStatusIndicator string `json:"maritalStatusIndicator"`
}

type Enthusiasts struct {
	Niches         string `json:"niches"`
	PoliticalParty string `json:"politicalParty"`
}

type HomeInfoPerson struct {
	HomeOwner          string `json:"homeOwner"`
	HomeOwnerIndicator string `json:"homeOwnerIndicator"`
}

type Consumers struct {
	ValueScore string `json:"valueScore"`
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
	Range            string `json:"range"`
	Value            int    `json:"value"`
	Birthday         Date   `json:"birthday"`
	AgeIn2YearRanges string `json:"ageIn2YearRanges"`
	ExactAge         string `json:"exactAge"`
}

type ChildrenInfo struct {
	Gender    ChildrenInfoDetails `json:"gender"`
	BirthDate ChildrenInfoDetails `json:"birthDate"`
}

type ChildrenInfoDetails struct {
	FirstChild  string `json:"firstChild"`
	SecondChild string `json:"secondChild"`
	ThirdChild  string `json:"thirdChild"`
	FourthChild string `json:"fourthChild"`
}

type Identifiers struct {
	Maids      []Maids  `json:"maids"`
	PersonIds  []string `json:"personIds"`
	RecordIds  []string `json:"recordIds"`
	LiNonId    []string `json:"li_nonid"`
	PartnerIds []string `json:"partnerIds"`
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
