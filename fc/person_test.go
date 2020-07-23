package fullcontact

import (
	assert "github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestPersonEnrich(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson := "{\"fullName\":\"Marquita H Ross\",\"ageRange\":\"37-47\",\"gender\":\"Female\",\"location\":\"San Francisco, California, United States\",\"title\":\"Senior Petroleum Manager\",\"organization\":\"Mostow Co.\",\"twitter\":\"https://twitter.com/marqross91\",\"linkedin\":\"https://www.linkedin.com/in/marquita-ross-5b6b72192\",\"facebook\":null,\"bio\":\"Senior Petroleum Manager at Mostow Co.\",\"avatar\":\"https://img.fullcontact.com/sandbox/1gagrO2K67_oc5DLG_siVCpYVE5UvCu2Z.png\",\"website\":\"http://marquitaas8.com/\",\"details\":{\"name\":{\"given\":\"Marquita\",\"family\":\"Ross\",\"full\":\"Marquita H Ross\"},\"age\":{\"range\":\"35-44\",\"value\":42},\"gender\":\"Female\",\"household\":{\"familyInfo\":{\"totalAdults\":2,\"totalChildren\":1,\"totalPeopleInHousehold\":3},\"homeInfo\":{\"dwellingType\":\"Multi Family Dwelling/Apartment\",\"homeValueEstimate\":513,\"loanToValueEstimate\":4},\"locationInfo\":{\"seasonalAddress\":true,\"carrierRoute\":\"PO Box\",\"designatedMarketArea\":\"807\",\"coreBasedStatisticalArea\":\"41860 - San Francisco-Oakland-Hayward, CA Metropolitan Statistical Area\",\"nielsenCountySize\":\"B - All counties in the next largest set of metropolitan areas which toher account for 30% of U.S. households according to the 2000 Census. No non-metropolitan counties qualify as A or B counties.\",\"congressionalDistrict\":14,\"numericCountyCode\":222},\"presence\":{\"multigenerationalResident\":\"PRESENT\",\"children\":\"PRESENT\"},\"finance\":{\"discretionaryIncomeEstimate\":45,\"financialDebtRangeEstimate\":\"$0\",\"householdIncomeEstimate\":\"$150,000 - $199,999\",\"netWorthRange\":\"$50,000 - $74,999\",\"cashValueBalanceHouseholdEstimate\":\"$25,000 - $49,999\"}},\"demographics\":{\"gender\":\"Female\",\"age\":{\"range\":\"35-44\",\"value\":42},\"livingStatus\":\"Probable Homeowner\",\"maritalStatus\":\"MARRIED\",\"occupation\":\"Professional - Engineer/Industrial\"},\"survey\":{\"own\":{\"ownDigitalCamera\":true,\"ownDog\":true,\"ownSwimmingPool\":true,\"ownRv\":true},\"collectibles\":{\"general\":true,\"coins\":true,\"stamps\":true},\"creditCards\":{\"premium\":{\"amex\":true,\"store\":true,\"visaOrMasterCard\":true},\"regular\":{\"amex\":true,\"other\":true,\"visaOrMasterCard\":true},\"debit\":true},\"dietConcerns\":{\"general\":true,\"loseWeight\":true,\"vitaminSupplements\":true,\"healthy\":true},\"hobby\":{\"general\":true,\"baking\":true,\"cigarSmoking\":true,\"gourmetCooking\":true,\"cooking\":true,\"gardening\":{\"general\":true,\"flowers\":true},\"homeImprovement\":true,\"exercise3xPerWeek\":true,\"lowFatCooking\":true,\"diy\":true,\"spirituality\":true},\"onlinePurchaser\":true,\"investments\":true,\"music\":{\"general\":true,\"christianOrGospel\":true,\"classical\":true,\"other\":true,\"rhythmAndBlues\":true},\"reading\":{\"likesToRead\":true,\"bibleOrDevotional\":true,\"bestSellingFiction\":true,\"childrens\":true,\"history\":true,\"health\":true,\"naturalHealthRemedies\":true,\"entertainment\":true,\"worldNewsOrPolitics\":true,\"bestSellers\":true,\"magazines\":true},\"sporting\":{\"fitness\":true,\"walking\":true,\"running\":true},\"travel\":{\"general\":true,\"usBusiness\":true,\"casinoVacations\":true,\"frequentFlyer\":true},\"religious\":true,\"electronics\":{\"appleDevice\":true,\"cableTv\":true,\"highSpeedInternet\":true,\"dvdPlayer\":true,\"hdtv\":true,\"homeTheater\":true,\"other\":true},\"purchase\":{\"homeDecorating\":true,\"beautyProducts\":true,\"clubStores\":true,\"usesCoupons\":true}},\"finance\":{\"bankCard\":\"Multiple Bank Card\",\"activeLineOfCredit\":true,\"retailCard\":\"Multiple Retail Card\"},\"census\":{\"basicTractNumber\":145,\"basicBlockGroup\":2,\"year2010\":{\"educationLevel\":\"High School Diploma\",\"average\":{\"numberOfCarsInHousehold\":2.3000000000000003},\"percent\":{\"abovePovertyLevel\":74,\"belowPovertyLevel\":28,\"black\":17,\"blueCollarEmployed\":35,\"divorcedOrSeparated\":14,\"hispanic\":43,\"homesBuiltSince2000\":52,\"homeowner\":90,\"householdsWithChildren\":38,\"married\":36,\"mobileHome\":45,\"movedToAreaSince2000\":21,\"salariedProfessional\":35,\"singleFamilyHome\":43,\"vehicleOwnership\":84,\"white\":35},\"median\":{\"ageOfHouseholder\":23,\"effectiveBuyingIncome\":18,\"homeValue\":35,\"householdIncome\":18,\"householdIncomeByState\":69},\"populationDensity\":{\"centileInState\":67,\"centileInUs\":17},\"socioEconomicScore\":29}},\"buyer\":{\"catalog\":{\"payment\":{\"amx\":true,\"creditCard\":true,\"houseCharge\":true,\"masterCard\":true,\"retailCard\":true,\"visa\":true},\"apparel\":{\"general\":true,\"children\":true,\"men\":true,\"mensBigTall\":true,\"nonGenderSpecific\":true,\"womenPetiteSize\":true,\"womenPlusSize\":true},\"beauty\":true,\"books\":true,\"childrenProducts\":true,\"electronics\":true,\"furniture\":true,\"health\":true,\"jewelry\":true,\"music\":true,\"stationery\":true,\"travel\":true,\"videoEntertainment\":true},\"retail\":{\"apparel\":{\"general\":true,\"children\":true,\"women\":true},\"beauty\":true,\"books\":true,\"childrenProducts\":true,\"collectibles\":true,\"foodBeverages\":true,\"gift\":true,\"health\":true,\"homeFurnishings\":true,\"jewelry\":true,\"music\":true,\"stationery\":true,\"travel\":true}},\"emails\":[],\"phones\":[],\"profiles\":{\"twitter\":{\"username\":\"marqross91\",\"url\":\"https://twitter.com/marqross91\",\"bio\":\"Senior Petroleum Manager at Mostow Co.\",\"service\":\"twitter\"},\"pinterest\":{\"username\":\"marquitaross006\",\"url\":\"http://www.pinterest.com/marquitaross006/\",\"bio\":\"Senior Petroleum Manager at Mostow Co.\",\"service\":\"pinterest\"},\"linkedin\":{\"username\":\"marquita-ross-5b6b72192\",\"userid\":\"761326554\",\"url\":\"https://www.linkedin.com/in/marquita-ross-5b6b72192\",\"bio\":\"Senior Petroleum Manager at Mostow Co.\",\"service\":\"linkedin\"}},\"locations\":[{\"city\":\"San Francisco\",\"region\":\"California\",\"regionCode\":\"CA\",\"country\":\"United States\",\"countryCode\":\"US\",\"formatted\":\"San Francisco, California, United States\"}],\"employment\":[{\"name\":\"Mostow Co.\",\"current\":true,\"title\":\"Senior Petroleum Manager\",\"start\":{\"year\":2019,\"month\":9}}],\"photos\":[{\"label\":\"avatar\",\"value\":\"https://img.fullcontact.com/sandbox/1gagrO2K67_oc5DLG_siVCpYVE5UvCu2Z.png\"}],\"education\":[{\"name\":\"University of California, Berkeley\",\"degree\":\"Bachelors\",\"end\":{\"year\":1998}}],\"urls\":[{\"value\":\"http://marquitaas8.com/\"}],\"interests\":[]},\"isSandboxProfile\":true,\"updated\":\"1970-01-01\"}"
	fcTestClient, testServer := getTestServerAndClient(personEnrichUrl, respJson, 200)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	response := resp.PersonResponse

	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "200 OK", resp.Status)
	assert.True(t, strings.Contains(resp.String(), "RawHttpRespons"))
	assert.Equal(t, "Marquita H Ross", response.FullName)
	assert.Equal(t, "37-47", response.AgeRange)
	assert.Equal(t, "Female", response.Gender)
	assert.Equal(t, "San Francisco, California, United States", response.Location)
	assert.Equal(t, "Senior Petroleum Manager", response.Title)
	assert.Equal(t, "Mostow Co.", response.Organization)
	assert.Equal(t, "Senior Petroleum Manager at Mostow Co.", response.Bio)
	assert.Equal(t, "https://img.fullcontact.com/sandbox/1gagrO2K67_oc5DLG_siVCpYVE5UvCu2Z.png", response.Avatar)
	assert.Equal(t, "http://marquitaas8.com/", response.Website)
	assert.Equal(t, "https://twitter.com/marqross91", response.Twitter)
	assert.Equal(t, "https://www.linkedin.com/in/marquita-ross-5b6b72192", response.Linkedin)
	assert.Equal(t, "Marquita", response.Details.Name.Given)
	assert.Equal(t, "Ross", response.Details.Name.Family)
	assert.Equal(t, "Marquita H Ross", response.Details.Name.Full)
	assert.Equal(t, "35-44", response.Details.Age.Range)
	assert.Equal(t, 42, response.Details.Age.Value)
	assert.Equal(t, "Female", response.Details.Gender)
	assert.Equal(t, 2, response.Details.Household.FamilyInfo.TotalAdults)
	assert.Equal(t, 1, response.Details.Household.FamilyInfo.TotalChildren)
	assert.Equal(t, 3, response.Details.Household.FamilyInfo.TotalPeopleInHousehold)
	assert.Equal(t, "Multi Family Dwelling/Apartment", response.Details.Household.HomeInfo.DwellingType)
	assert.Equal(t, 513, response.Details.Household.HomeInfo.HomeValueEstimate)
	assert.Equal(t, 4, response.Details.Household.HomeInfo.LoanToValueEstimate)
	assert.True(t, response.Details.Household.LocationInfo.SeasonalAddress)
	assert.Equal(t, "PO Box", response.Details.Household.LocationInfo.CarrierRoute)
	assert.Equal(t, "807", response.Details.Household.LocationInfo.DesignatedMarketArea)
	assert.Equal(t, "41860 - San Francisco-Oakland-Hayward, CA Metropolitan Statistical Area", response.Details.Household.LocationInfo.CoreBasedStatisticalArea)
	assert.Equal(t, 14, response.Details.Household.LocationInfo.CongressionalDistrict)
	assert.Equal(t, 222, response.Details.Household.LocationInfo.NumericCountyCode)
	assert.Equal(t, "PRESENT", response.Details.Household.Presence.MultigenerationalResident)
	assert.Equal(t, "PRESENT", response.Details.Household.Presence.Children)
	assert.Equal(t, 45, response.Details.Household.Finance.DiscretionaryIncomeEstimate)
	assert.Equal(t, "$150,000 - $199,999", response.Details.Household.Finance.HouseholdIncomeEstimate)
	assert.Equal(t, "$25,000 - $49,999", response.Details.Household.Finance.CashValueBalanceHouseholdEstimate)
	assert.Equal(t, "Probable Homeowner", response.Details.Demographics.LivingStatus)
	assert.Equal(t, "Professional - Engineer/Industrial", response.Details.Demographics.Occupation)
	assert.Equal(t, "Multiple Bank Card", response.Details.Finance.BankCard)
	assert.Equal(t, "Multiple Retail Card", response.Details.Finance.RetailCard)
	assert.True(t, response.Details.Finance.ActiveLineOfCredit)
	assert.Equal(t, 145, response.Details.Census.BasicTractNumber)
	assert.Equal(t, 2, response.Details.Census.BasicBlockGroup)
	assert.Equal(t, "High School Diploma", response.Details.Census.Year2010.EducationLevel)
	assert.Equal(t, 90, response.Details.Census.Year2010.Percent.Homeowner)
	assert.Equal(t, 35, response.Details.Census.Year2010.Median.HomeValue)
	assert.Equal(t, 17, response.Details.Census.Year2010.PopulationDensity.CentileInUs)
	assert.Equal(t, 29, response.Details.Census.Year2010.SocioEconomicScore)
	assert.True(t, response.Details.Buyer.Catalog.Payment.CreditCard)
	assert.True(t, response.Details.Buyer.Catalog.Payment.HouseCharge)
	assert.True(t, response.Details.Buyer.Catalog.Payment.MasterCard)
	assert.True(t, response.Details.Buyer.Catalog.Apparel.Children)
	assert.True(t, response.Details.Buyer.Catalog.Apparel.Men)
	assert.True(t, response.Details.Buyer.Catalog.Beauty)
	assert.True(t, response.Details.Buyer.Catalog.VideoEntertainment)
	assert.True(t, response.Details.Buyer.Retail.Beauty)
	assert.True(t, response.Details.Buyer.Retail.Apparel.Children)
	assert.True(t, response.Details.Survey.Own.OwnDigitalCamera)
	assert.True(t, response.Details.Survey.Collectibles.Coins)
	assert.True(t, response.Details.Survey.CreditCards.Premium.Amex)
	assert.True(t, response.Details.Survey.CreditCards.Debit)
	assert.True(t, response.Details.Survey.DietConcerns.Healthy)
	assert.True(t, response.Details.Survey.Hobby.CigarSmoking)
	assert.True(t, response.Details.Survey.Hobby.Gardening.Flowers)
	assert.True(t, response.Details.Survey.Hobby.Spirituality)
	assert.True(t, response.Details.Survey.Investments)
	assert.True(t, response.Details.Survey.Music.Classical)
	assert.True(t, response.Details.Survey.Reading.History)
	assert.True(t, response.Details.Survey.Sporting.Fitness)
	assert.True(t, response.Details.Survey.Travel.CasinoVacations)
	assert.True(t, response.Details.Survey.Electronics.HomeTheater)
	assert.True(t, response.Details.Survey.Purchase.UsesCoupons)
	assert.Equal(t, "marqross91", response.Details.Profiles.Twitter.Username)
	assert.Equal(t, "Senior Petroleum Manager at Mostow Co.", response.Details.Profiles.Twitter.Bio)
	assert.Equal(t, "marquita-ross-5b6b72192", response.Details.Profiles.Linkedin.Username)
	assert.Equal(t, "http://www.pinterest.com/marquitaross006/", response.Details.Profiles.Pinterest.URL)
	assert.Equal(t, "California", response.Details.Locations[0].Region)
	assert.Equal(t, "Mostow Co.", response.Details.Employment[0].Name)
	assert.Equal(t, "https://img.fullcontact.com/sandbox/1gagrO2K67_oc5DLG_siVCpYVE5UvCu2Z.png", response.Details.Photos[0].Value)
	assert.Equal(t, "University of California, Berkeley", response.Details.Education[0].Name)
	assert.Equal(t, "http://marquitaas8.com/", response.Details.Urls[0].Value)
}

func TestPersonEnrichAutoRetry(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(personEnrichUrl, "", 429)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 429, resp.StatusCode)
}

func TestPersonEnrichStatus400(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(personEnrichUrl, "", 400)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 400, resp.StatusCode)
}

func TestPersonEnrichStatus202(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(personEnrichUrl, "", 202)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.True(t, resp.IsSuccessful)
	assert.Equal(t, 202, resp.StatusCode)
}

func TestPersonEnrichStatus401(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(personEnrichUrl, "", 401)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 401, resp.StatusCode)
}

func TestPersonEnrichStatus403(t *testing.T) {
	ch := make(chan *APIResponse)
	fcTestClient, testServer := getTestServerAndClient(personEnrichUrl, "", 403)
	defer testServer.Close()
	go fcTestClient.do(testServer.URL, nil, ch)
	resp := <-ch
	assert.False(t, resp.IsSuccessful)
	assert.Equal(t, 403, resp.StatusCode)
}
