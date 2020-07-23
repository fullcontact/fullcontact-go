package fullcontact

type Survey struct {
	Own             Own          `json:"own"`
	Collectibles    Collectibles `json:"collectibles"`
	CreditCards     CreditCards  `json:"creditCards"`
	DietConcerns    DietConcerns `json:"dietConcerns"`
	Hobby           Hobby        `json:"hobby"`
	Music           Music        `json:"music"`
	Reading         Reading      `json:"reading"`
	Sporting        Sporting     `json:"sporting"`
	Travel          Travel       `json:"travel"`
	Electronics     Electronics  `json:"electronics"`
	Purchase        Purchase     `json:"purchase"`
	Religious       bool         `json:"religious"`
	Grandchildren   bool         `json:"grandchildren"`
	OnlinePurchaser bool         `json:"onlinePurchaser"`
	Investments     bool         `json:"investments"`
}

type Own struct {
	OwnCat           bool `json:"ownCat"`
	OwnDog           bool `json:"ownDog"`
	OwnDigitalCamera bool `json:"ownDigitalCamera"`
	OwnHorse         bool `json:"ownHorse"`
	OwnMotorcycle    bool `json:"ownMotorcycle"`
	OwnSwimmingPool  bool `json:"ownSwimmingPool"`
	OwnAtv           bool `json:"ownAtv"`
	OwnRv            bool `json:"ownRv"`
}

type Collectibles struct {
	General           bool `json:"general"`
	Coins             bool `json:"coins"`
	Dolls             bool `json:"dolls"`
	Figurines         bool `json:"figurines"`
	Other             bool `json:"other"`
	Plates            bool `json:"plates"`
	SportsMemorabilia bool `json:"sportsMemorabilia"`
	Stamps            bool `json:"stamps"`
	FineArts          bool `json:"fineArts"`
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
	General            bool `json:"general"`
	LoseWeight         bool `json:"loseWeight"`
	EatNatural         bool `json:"eatNatural"`
	VitaminSupplements bool `json:"vitaminSupplements"`
	Healthy            bool `json:"healthy"`
}

type Hobby struct {
	Gardening                Gardening `json:"gardening"`
	General                  bool      `json:"general"`
	Baking                   bool      `json:"baking"`
	BirdWatching             bool      `json:"birdWatching"`
	Cars                     bool      `json:"cars"`
	CigarSmoking             bool      `json:"cigarSmoking"`
	GourmetCooking           bool      `json:"gourmetCooking"`
	Cooking                  bool      `json:"cooking"`
	Crafts                   bool      `json:"crafts"`
	CasinoGambling           bool      `json:"casinoGambling"`
	HomeImprovement          bool      `json:"homeImprovement"`
	HomeStudyCourses         bool      `json:"homeStudyCourses"`
	Knitting                 bool      `json:"knitting"`
	Lotteries                bool      `json:"lotteries"`
	Quilting                 bool      `json:"quilting"`
	SelfImprovementCourses   bool      `json:"selfImprovementCourses"`
	Sewing                   bool      `json:"sewing"`
	Theater                  bool      `json:"theater"`
	Woodworking              bool      `json:"woodworking"`
	WineAppreciation         bool      `json:"wineAppreciation"`
	Photography              bool      `json:"photography"`
	Exercise3xPerWeek        bool      `json:"exercise3xPerWeek"`
	ScrapBooking             bool      `json:"scrapBooking"`
	LowFatCooking            bool      `json:"lowFatCooking"`
	CareerAdvancementCourses bool      `json:"careerAdvancementCourses"`
	JewelryMaking            bool      `json:"jewelryMaking"`
	Diy                      bool      `json:"diy"`
	Green                    bool      `json:"green"`
	SocialNetworking         bool      `json:"socialNetworking"`
	Spirituality             bool      `json:"spirituality"`
}

type Gardening struct {
	General   bool `json:"general"`
	Flowers   bool `json:"flowers"`
	Organic   bool `json:"organic"`
	Vegetable bool `json:"vegetable"`
}

type Music struct {
	General           bool `json:"general"`
	ChristianOrGospel bool `json:"christianOrGospel"`
	Classical         bool `json:"classical"`
	Country           bool `json:"country"`
	Jazz              bool `json:"jazz"`
	Other             bool `json:"other"`
	RhythmAndBlues    bool `json:"rhythmAndBlues"`
	Rock              bool `json:"rock"`
	SoftRock          bool `json:"softRock"`
	Swing             bool `json:"swing"`
	Alternative       bool `json:"alternative"`
}

type Reading struct {
	LikesToRead           bool `json:"likesToRead"`
	Astrology             bool `json:"astrology"`
	BibleOrDevotional     bool `json:"bibleOrDevotional"`
	BestSellingFiction    bool `json:"bestSellingFiction"`
	Audiobooks            bool `json:"audiobooks"`
	Childrens             bool `json:"childrens"`
	Cooking               bool `json:"cooking"`
	Computer              bool `json:"computer"`
	CountryLifestyle      bool `json:"countryLifestyle"`
	Fashion               bool `json:"fashion"`
	History               bool `json:"history"`
	InteriorDecorating    bool `json:"interiorDecorating"`
	Health                bool `json:"health"`
	Military              bool `json:"military"`
	Mystery               bool `json:"mystery"`
	NaturalHealthRemedies bool `json:"naturalHealthRemedies"`
	Entertainment         bool `json:"entertainment"`
	Romance               bool `json:"romance"`
	ScienceFiction        bool `json:"scienceFiction"`
	Technology            bool `json:"technology"`
	Sports                bool `json:"sports"`
	WorldNewsOrPolitics   bool `json:"worldNewsOrPolitics"`
	Suspense              bool `json:"suspense"`
	BestSellers           bool `json:"bestSellers"`
	BookClub              bool `json:"bookClub"`
	Comics                bool `json:"comics"`
	Financial             bool `json:"financial"`
	HomeAndGarden         bool `json:"homeAndGarden"`
	SelfImprovement       bool `json:"selfImprovement"`
	Travel                bool `json:"travel"`
	Magazines             bool `json:"magazines"`
}

type Sporting struct {
	Other            bool `json:"other"`
	CampingOrHiking  bool `json:"campingOrHiking"`
	Baseball         bool `json:"baseball"`
	Boating          bool `json:"boating"`
	Basketball       bool `json:"basketball"`
	Fishing          bool `json:"fishing"`
	AmericanFootball bool `json:"americanFootball"`
	Fitness          bool `json:"fitness"`
	Golf             bool `json:"golf"`
	Hockey           bool `json:"hockey"`
	Hunting          bool `json:"hunting"`
	Nascar           bool `json:"nascar"`
	SnowSkiing       bool `json:"snowSkiing"`
	Walking          bool `json:"walking"`
	Running          bool `json:"running"`
	Scuba            bool `json:"scuba"`
	Tennis           bool `json:"tennis"`
	WeightLifting    bool `json:"weightLifting"`
	Biking           bool `json:"biking"`
	ExtremeSports    bool `json:"extremeSports"`
	Motocross        bool `json:"motocross"`
	Skateboarding    bool `json:"skateboarding"`
	Snowboarding     bool `json:"snowboarding"`
	Rollerblading    bool `json:"rollerblading"`
	Interests        bool `json:"interests"`
}

type Travel struct {
	General                 bool `json:"general"`
	UsBusiness              bool `json:"usBusiness"`
	InternationalBusiness   bool `json:"internationalBusiness"`
	UsPersonal              bool `json:"usPersonal"`
	InternationalPersonal   bool `json:"internationalPersonal"`
	CasinoVacations         bool `json:"casinoVacations"`
	FamilyVacations         bool `json:"familyVacations"`
	FrequentFlyer           bool `json:"frequentFlyer"`
	Timeshare               bool `json:"timeshare"`
	VacationCruises         bool `json:"vacationCruises"`
	AttractionsOrThemeParks bool `json:"attractionsOrThemeParks"`
	Rv                      bool `json:"rv"`
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
	HomeDecorating          bool `json:"homeDecorating"`
	BeautyProducts          bool `json:"beautyProducts"`
	ClubStores              bool `json:"clubStores"`
	FastFoods               bool `json:"fastFoods"`
	SpecialtyBeautyProducts bool `json:"specialtyBeautyProducts"`
	UsesCoupons             bool `json:"usesCoupons"`
}
