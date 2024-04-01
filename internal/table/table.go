package table

type Admin struct {
	AdminId  int
	Account  string
	Password string
}

type Admin_detail struct {
	AdminId         int
	ShopName        string
	ShopInfo        string
	ShopImage       string
	CorporationName string
	ShopLocation    string
	OpenTime        string
	DayOff          string
	PhoneNumber     string
	Email           string
}

type Shop struct {
	ShopId   int
	Account  string
	Password string
	Status   int
}

type Shop_detail struct {
	ShopId          int
	ShopName        string
	ShopInfo        string
	ShopImage       string
	CorporationName string
	ShopLocation    string
	ShopCity        string
	OpenTime        string
	DayOff          string
	PhoneNumber     string
	Email           string
}

type Car struct {
	CarId    int
	ShopId   int
	CarName  string
	CarBrand string
	CarImage string
	CarPrice int
	CarFee   int
	CarYear  int
	Shelves  bool
}
