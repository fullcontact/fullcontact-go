package fullcontact

import (
	"os"
	"strings"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPersonEnrich(t *testing.T) {
	ch := make(chan *APIResponse)
	respJson, _ := os.ReadFile("person_test.json")

	fcTestClient, testServer := getTestServerAndClient(personEnrichUrl, string(respJson), 200)
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
	assert.Equal(t, "14", response.Details.Household.LocationInfo.CongressionalDistrict)
	assert.Equal(t, 222, response.Details.Household.LocationInfo.NumericCountyCode)
	assert.Equal(t, "PRESENT", response.Details.Household.Presence.MultigenerationalResident)
	assert.Equal(t, "PRESENT", response.Details.Household.Presence.Children)

	// Finance
	assert.Equal(t, "$30,000 - $39,999", response.Details.Household.Finance.Income)
	assert.Equal(t, "$65,000 - $74,999", response.Details.Household.Finance.DiscretionarySpendingIncome)
	assert.Equal(t, "5147", response.Details.Household.Finance.FirstMortgageAmountInThousands)
	assert.Equal(t, "453", response.Details.Household.Finance.HomeMarketValueTaxRecord)
	assert.Equal(t, "$50,000 or more", response.Details.Household.Finance.ShortTermLiability)
	assert.Equal(t, "$150,000 - $249,999", response.Details.Household.Finance.NetWorth)
	assert.Equal(t, "$250,000 - $499,999", response.Details.Household.Finance.WealthResources)
	assert.Equal(t, "Y", response.Details.Household.Finance.PaymentMethodCreditCard)

	assert.Equal(t, "Probable Homeowner", response.Details.Demographics.LivingStatus)
	assert.Equal(t, "Professional - Engineer/Industrial", response.Details.Demographics.Occupation)
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
	assert.Equal(t, "9", response.Extended["epsilon_key_1"])
	assert.Equal(t, "Q", response.Extended["epsilon_key_3"])
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
