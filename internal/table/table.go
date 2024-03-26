package table

type Admin struct {
	Admin_id int
	Account  string
	Password string
}

type Admin_detail struct {
	Admin_id        int
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
	Shop_id  int
	Account  string
	Password string
}

type Shop_detail struct {
	Shop_id         int
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

type Car struct {
	Car_id    int
	Shop_id   int
	Car_Name  string
	Car_Brand string
	Car_Image string
	Car_Price int
	Car_Fee   int
	Shelves   bool
}
