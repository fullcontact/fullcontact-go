package fullcontact

type Survey struct {
	Own          Own          `json:"own"`
	Collectibles Collectibles `json:"collectibles"`
	DietConcerns DietConcerns `json:"dietConcerns"`
	Hobby        Hobby        `json:"hobby"`
	Music        Music        `json:"music"`
	Reading      Reading      `json:"reading"`
	Sporting     Sporting     `json:"sporting"`
	Travel       Travel       `json:"travel"`
	Purchase     Purchase     `json:"purchase"`
	Investments  Investments  `json:"investments"`
	Donor        SurveyDonor  `json:"donor"`
	MailOrder    MailOrder    `json:"mailOrder"`
	Other        Other        `json:"other"`
	Social       Social       `json:"social"`
}
type Social struct {
	FacebookUser  string `json:"facebookUser"`
	InstagramUser string `json:"instagramUser"`
	PinterestUser string `json:"pinterestUser"`
	TwitterUser   string `json:"twitterUser"`
}
type Other struct {
	Electronics             string `json:"electronics"`
	Grandchildren           string `json:"grandchildren"`
	MilitaryVeteran         string `json:"militaryVeteran"`
	OnlineHousehold         string `json:"onlineHousehold"`
	ScienceAndNewTechnology string `json:"scienceAndNewTechnology"`
	SwimmingPool            string `json:"swimmingPool"`
}
type MailOrder struct {
	Any                    string `json:"any"`
	Apparel                string `json:"apparel"`
	Books                  string `json:"books"`
	Buyer                  string `json:"buyer"`
	ChildrensProducts      string `json:"childrensProducts"`
	Food                   string `json:"food"`
	Gifts                  string `json:"gifts"`
	HealthOrBeautyProducts string `json:"healthOrBeautyProducts"`
	HomeFurnishing         string `json:"homeFurnishing"`
	Jewelry                string `json:"jewelry"`
	VideosOrDVD            string `json:"videosOrDVD"`
	WomensPlusApparel      string `json:"womensPlusApparel"`
}
type Investments struct {
	JuvenileLifeInsurance string `json:"juvenileLifeInsurance"`
	BurialInsurance       string `json:"burialInsurance"`
	Insurance             string `json:"insurance"`
	Investments           string `json:"investments"`
	LifeInsurance         string `json:"lifeInsurance"`
	MedicareCoverage      string `json:"medicareCoverage"`
	MutualFunds           string `json:"mutualFunds"`
	StocksOrBonds         string `json:"stocksOrBonds"`
}
type SurveyDonor struct {
	ActiveMilitary        string `json:"activeMilitary"`
	AnimalWelfare         string `json:"animalWelfare"`
	ArtsOrCultural        string `json:"artsOrCultural"`
	Cancer                string `json:"cancer"`
	Catholic              string `json:"catholic"`
	Childrens             string `json:"childrens"`
	Charitable            string `json:"charitable"`
	Humanitarian          string `json:"humanitarian"`
	NativeAmerican        string `json:"nativeAmerican"`
	OtherReligious        string `json:"otherReligious"`
	PoliticalConservative string `json:"politicalConservative"`
	PoliticalLiberal      string `json:"politicalLiberal"`
	Veteran               string `json:"veteran"`
	WildlifeEnvironmental string `json:"wildlifeEnvironmental"`
	WorldRelief           string `json:"worldRelief"`
}

type Own struct {
	SmartPhone string `json:"smartPhone"`
	Cat        string `json:"cat"`
	Dog        string `json:"dog"`
	Pets       string `json:"pets"`
}

type Collectibles struct {
	Dolls          string `json:"dolls"`
	Figurines      string `json:"figurines"`
	Any            string `json:"any"`
	ArtAntique     string `json:"artAntique"`
	ClubContinuity string `json:"clubContinuity"`
}

type CreditCards struct {
	Premium Card `json:"premium"`
	Regular Card `json:"regular"`
	Debit   bool `json:"debit"`
}

type Card struct {
	Amex             bool `json:"amex"`
	Discover         bool `json:"discover"`
	Other            bool `json:"other"`
	Store            bool `json:"store"`
	VisaOrMasterCard bool `json:"visaOrMasterCard"`
}

type DietConcerns struct {
	NaturalFoods  string `json:"naturalFoods"`
	Diet          string `json:"diet"`
	WeightControl string `json:"weightControl"`
}

type Hobby struct {
	Baking                       string `json:"baking"`
	CigarSmoking                 string `json:"cigarSmoking"`
	Crafts                       string `json:"crafts"`
	Cooking                      string `json:"cooking"`
	Gardening                    string `json:"gardening"`
	HomeStudyCourses             string `json:"homeStudyCourses"`
	Quilting                     string `json:"quilting"`
	SelfImprovementCourses       string `json:"selfImprovementCourses"`
	Woodworking                  string `json:"woodworking"`
	Photography                  string `json:"photography"`
	CareerAdvancementCourses     string `json:"careerAdvancementCourses"`
	Any                          string `json:"any"`
	AutomotiveWork               string `json:"automotiveWork"`
	BirdFeedingOrWatching        string `json:"birdFeedingOrWatching"`
	CulturalArtsOrEvents         string `json:"culturalArtsOrEvents"`
	GourmetFoods                 string `json:"gourmetFoods"`
	HomeImprovementOrDIY         string `json:"homeImprovementOrDIY"`
	MotorcycleRiding             string `json:"motorcycleRiding"`
	Scrapbooking                 string `json:"scrapbooking"`
	SewingOrNeedleworkOrKnitting string `json:"sewingOrNeedleworkOrKnitting"`
	Wine                         string `json:"wine"`
}

type Gardening struct {
	General   bool `json:"general"`
	Flowers   bool `json:"flowers"`
	Organic   bool `json:"organic"`
	Vegetable bool `json:"vegetable"`
}

type Music struct {
	ChristianOrGospel string `json:"christianOrGospel"`
	Classical         string `json:"classical"`
	Country           string `json:"country"`
	Jazz              string `json:"jazz"`
	Any               string `json:"any"`
	RhythmAndBlues    string `json:"rhythmAndBlues"`
	RockNRoll         string `json:"rockNRoll"`
}

type Reading struct {
	BibleOrDevotional  string `json:"bibleOrDevotional"`
	BestSellingFiction string `json:"bestSellingFiction"`
	Childrens          string `json:"childrens"`
	Fashion            string `json:"fashion"`
	Military           string `json:"military"`
	Entertainment      string `json:"entertainment"`
	Romance            string `json:"romance"`
	Sports             string `json:"sports"`
	Books              string `json:"books"`
	CookingOrCulinary  string `json:"cookingOrCulinary"`
	CountryOrLifestyle string `json:"countryOrLifestyle"`
	Interior           string `json:"interior"`
	MedicalOrHealth    string `json:"medicalOrHealth"`
	WorldNews          string `json:"worldNews"`
}

type Sporting struct {
	CampingOrHiking                      string `json:"campingOrHiking"`
	Fishing                              string `json:"fishing"`
	Golf                                 string `json:"golf"`
	Nascar                               string `json:"nascar"`
	BoatingOrSailing                     string `json:"boatingOrSailing"`
	Cycling                              string `json:"cycling"`
	FitnessExcercise                     string `json:"fitnessExcercise"`
	BigGameHunting                       string `json:"bigGameHunting"`
	HuntingOrShooting                    string `json:"huntingOrShooting"`
	SportsMerchandiseOrActivewearRecency string `json:"sportsMerchandiseOrActivewearRecency"`
	RunningOrJogging                     string `json:"runningOrJogging"`
	SkiingOrSnowboarding                 string `json:"skiingOrSnowboarding"`
	SportsParticipation                  string `json:"sportsParticipation"`
	WalkingForHealth                     string `json:"walkingForHealth"`
	YogaOrPilates                        string `json:"yogaOrPilates"`
}

type Travel struct {
	Timeshare          string `json:"timeshare"`
	Business           string `json:"business"`
	CruiseShipVacation string `json:"cruiseShipVacation"`
	International      string `json:"international"`
	Leisure            string `json:"leisure"`
	RvVacations        string `json:"rvVacations"`
	TravelInTheUSA     string `json:"travelInTheUSA"`
	Traveler           string `json:"traveler"`
}

type Electronics struct {
	AppleDevice       bool `json:"appleDevice"`
	CableTv           bool `json:"cableTv"`
	HighSpeedInternet bool `json:"highSpeedInternet"`
	Dvr               bool `json:"dvr"`
	DvdPlayer         bool `json:"dvdPlayer"`
	Hdtv              bool `json:"hdtv"`
	HomeTheater       bool `json:"homeTheater"`
	SatelliteRadio    bool `json:"satelliteRadio"`
	SatelliteTv       bool `json:"satelliteTv"`
	VideoGameSystems  bool `json:"videoGameSystems"`
	Other             bool `json:"other"`
}

type Purchase struct {
	ArtsCraftsRecency             string `json:"artsCraftsRecency"`
	BeautyAndSpaRecency           string `json:"beautyAndSpaRecency"`
	BeveragesRecency              string `json:"beveragesRecency"`
	BooksRecency                  string `json:"booksRecency"`
	ClubContinuity                string `json:"clubContinuity"`
	GardenAndBackyardRecency      string `json:"gardenAndBackyardRecency"`
	HomeDecorRecency              string `json:"homeDecorRecency"`
	Sweepstakes                   string `json:"sweepstakes"`
	MaleApparelRecency            string `json:"maleApparelRecency"`
	MusicVideosRecency            string `json:"musicVideosRecency"`
	SpecialtyFoodsAndGiftsRecency string `json:"specialtyFoodsAndGiftsRecency"`
	SportsAndOutdoorRecency       string `json:"sportsAndOutdoorRecency"`
	ToolsAndElectronicsRecency    string `json:"toolsAndElectronicsRecency"`
	FemaleAndMaleRecency          string `json:"femaleAndMaleRecency"`
	FemaleBrandAndFitRecency      string `json:"femaleBrandAndFitRecency"`
}
