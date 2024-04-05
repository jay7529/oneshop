package verify

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Shop_Signup_Verify(c *gin.Context) bool {
	type Verify struct {
		Email string `validate:"required,email"`
	}

	verify := &Verify{
		Email: c.PostForm("email"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Shop_Code_Verify(c *gin.Context) bool {
	type Verify struct {
		Email string `validate:"required,email"`
		Code  string `validate:"required,max=6,min=6"`
	}

	verify := &Verify{
		Email: c.PostForm("email"),
		Code:  c.PostForm("code"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Shop_New_Signup_Verify(c *gin.Context) bool {
	type Verify struct {
		Email       string `validate:"required,email"`
		Code        string `validate:"required,max=6,min=6"`
		Password    string `validate:"required,max=32,min=6"`
		ShopName    string `validate:"required,min=1"`
		PostCode    string `validate:"required,min=6"`
		Address     string `validate:"required,max=32,min=1"`
		PhoneNumber string `validate:"required,min=6"`
	}

	verify := &Verify{
		Email:       c.PostForm("email"),
		Code:        c.PostForm("code"),
		Password:    c.PostForm("password"),
		ShopName:    c.PostForm("shopName"),
		PostCode:    c.PostForm("postCode"),
		Address:     c.PostForm("address"),
		PhoneNumber: c.PostForm("phoneNumber"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Shop_Login_Verify(c *gin.Context) bool {
	type Verify struct {
		Account  string `validate:"required,max=100,min=1"`
		Password string `validate:"required,max=20,min=6"`
	}

	verify := &Verify{
		Account:  c.PostForm("account"),
		Password: c.PostForm("password"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Shop_Forget_Password_Verify(c *gin.Context) bool {
	type Verify struct {
		Email string `validate:"required,max=100,min=1"`
	}

	verify := &Verify{
		Email: c.PostForm("email"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Reset_Shop_Password_Verify(c *gin.Context) bool {
	type Verify struct {
		Email       string `validate:"required,max=100,min=1"`
		OldPassword string `validate:"required,max=32,min=6"`
		NewPassword string `validate:"required,necsfield=OldPassword,max=32,min=6"`
	}

	verify := &Verify{
		Email:       c.PostForm("email"),
		OldPassword: c.PostForm("oldPassword"),
		NewPassword: c.PostForm("newPassword"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Update_Shop_Password_Verify(c *gin.Context) bool {
	type Verify struct {
		OldPassword string `validate:"required,max=32,min=6"`
		NewPassword string `validate:"required,necsfield=OldPassword,max=32,min=6"`
	}

	verify := &Verify{
		OldPassword: c.PostForm("oldPassword"),
		NewPassword: c.PostForm("newPassword"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Update_Shop_Detail_Verify(c *gin.Context) bool {
	type Verify struct {
		ShopName        string `validate:"required,max=20,min=1"`
		ShopInfo        string `validate:"required,max=100"`
		ShopImage       string `validate:"required,max=100"`
		CorporationName string `validate:"required,max=100"`
		PostCode        string `validate:"required,min=6"`
		ShopLocation    string `validate:"required,max=100"`
		ShopCity        string `validate:"required,max=20"`
		OpenTime        string `validate:"required,max=100"`
		DayOff          string `validate:"required,max=100"`
		PhoneNumber     string `validate:"required,max=30"`
		Email           string `validate:"required,max=100"`
	}

	verify := &Verify{
		ShopName:        c.PostForm("shopName"),
		ShopInfo:        c.PostForm("shopInfo"),
		ShopImage:       c.PostForm("shopImage"),
		CorporationName: c.PostForm("corporationName"),
		PostCode:        c.PostForm("postCode"),
		ShopLocation:    c.PostForm("shopLocation"),
		ShopCity:        c.PostForm("shopCity"),
		OpenTime:        c.PostForm("openTime"),
		DayOff:          c.PostForm("dayOff"),
		PhoneNumber:     c.PostForm("phoneNumber"),
		Email:           c.PostForm("email"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Insert_Shop_Car_Verify(c *gin.Context) bool {
	type Verify struct {
		CarName  string `validate:"required,max=50,min=1"`
		CarBrand string `validate:"required,max=50"`
		CarImage string `validate:"required,max=100"`
		CarPrice string `validate:"required,max=100"`
		CarFee   string `validate:"required,max=100"`
		CarYear  string `validate:"required,max=100"`
	}

	verify := &Verify{
		CarName:  c.PostForm("carName"),
		CarBrand: c.PostForm("carBrand"),
		CarImage: c.PostForm("carImage"),
		CarPrice: c.PostForm("carPrice"),
		CarFee:   c.PostForm("carFee"),
		CarYear:  c.PostForm("carYear"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Update_Shop_Car_Verify(c *gin.Context) bool {
	type Verify struct {
		CarId    string `validate:"required,min=1"`
		CarName  string `validate:"required,max=50,min=1"`
		CarBrand string `validate:"required,max=50"`
		CarImage string `validate:"required,max=100"`
		CarPrice string `validate:"required,max=100"`
		CarFee   string `validate:"required,max=100"`
		CarYear  string `validate:"required,max=100"`
		Shelves  string `validate:"required,max=1"`
	}

	verify := &Verify{
		CarId:    c.Param("car_id"),
		CarName:  c.PostForm("carName"),
		CarBrand: c.PostForm("carBrand"),
		CarImage: c.PostForm("carImage"),
		CarPrice: c.PostForm("carPrice"),
		CarFee:   c.PostForm("carFee"),
		CarYear:  c.PostForm("carYear"),
		Shelves:  c.PostForm("shelves"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Delete_Shop_Car_Verify(c *gin.Context) bool {
	type Verify struct {
		CarId string `validate:"required"`
	}

	verify := &Verify{
		CarId: c.Param("car_id"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Get_Shop_Car_Verify(c *gin.Context) bool {
	type Verify struct {
		CarId string `validate:"required"`
	}

	verify := &Verify{
		CarId: c.Param("car_id"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Insert_Shop_Staff_Verify(c *gin.Context) bool {
	type Verify struct {
		StaffName         string `validate:"required,max=50,min=1"`
		StaffImage        string `validate:"required,max=50"`
		StaffPosition     string `validate:"required,max=50"`
		StaffIntroduction string `validate:"required,max=200"`
	}

	verify := &Verify{
		StaffName:         c.PostForm("staffName"),
		StaffImage:        c.PostForm("staffImage"),
		StaffPosition:     c.PostForm("staffPosition"),
		StaffIntroduction: c.PostForm("staffIntroduction"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Update_Shop_Staff_Verify(c *gin.Context) bool {
	type Verify struct {
		StaffId           string `validate:"required"`
		StaffName         string `validate:"required,max=50,min=1"`
		StaffImage        string `validate:"required,max=50"`
		StaffPosition     string `validate:"required,max=50"`
		StaffIntroduction string `validate:"required,max=200"`
	}

	verify := &Verify{
		StaffId:           c.Param("staff_id"),
		StaffName:         c.PostForm("staffName"),
		StaffImage:        c.PostForm("staffImage"),
		StaffPosition:     c.PostForm("staffPosition"),
		StaffIntroduction: c.PostForm("staffIntroduction"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Delete_Shop_Staff_Verify(c *gin.Context) bool {
	type Verify struct {
		StaffId string `validate:"required"`
	}

	verify := &Verify{
		StaffId: c.Param("staff_id"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}
