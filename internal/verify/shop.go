package verify

import (
	"oneshop/database"
	"oneshop/middleware"
	"oneshop/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// token驗證
func Shop_Token_Verify(c *gin.Context) int {

	claim, err := middleware.ParseToken(c.GetHeader("token"))
	if err == nil && claim != nil &&
		claim.ExpiresAt >= time.Now().Unix() &&
		claim.Identity == "shop" &&
		database.ExistsHkey("shop", utils.IntToString(claim.ID)) &&
		c.GetHeader("token") == database.GetHkey("shop", utils.IntToString(claim.ID)) {
		return claim.ID
	} else {
		return 0
	}
}

func Shop_Login_Verify(c *gin.Context) bool {
	type Verify struct {
		account  string `validate:"required,max=15,min=1"`
		password string `validate:"required,max=15,min=1"`
	}

	verify := &Verify{
		account:  c.PostForm("account"),
		password: c.PostForm("password"),
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
		ShopLocation    string `validate:"required,max=100"`
		OpenTime        string `validate:"required,max=100"`
		DayOff          string `validate:"required,max=100"`
		PhoneNumber     string `validate:"required,max=100"`
		Email           string `validate:"required,max=100"`
	}

	verify := &Verify{
		ShopName:        c.PostForm("shopName"),
		ShopInfo:        c.PostForm("shopInfo"),
		ShopImage:       c.PostForm("shopImage"),
		CorporationName: c.PostForm("corporationName"),
		ShopLocation:    c.PostForm("shopLocation"),
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
		CarId:    c.Param("carId"),
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
		CarId: c.Param("carId"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Get_Shop_Car_Verify(c *gin.Context) bool {
	type Verify struct {
		CarId string `validate:"required"`
	}

	verify := &Verify{
		CarId: c.Param("carId"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}
