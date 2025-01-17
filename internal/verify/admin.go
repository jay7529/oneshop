package verify

import (
	"oneshop/middleware"
	"oneshop/tools"
	"oneshop/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// token驗證
func Admin_Token_Verify(c *gin.Context) int {

	type Verify struct {
		Token string `validate:"required,min=1"`
	}
	verify := &Verify{
		Token: c.GetHeader("token"),
	}
	err := validator.New().Struct(verify)
	if err == nil {
		return 0
	}

	claim, err := middleware.ParseToken(c.GetHeader("token"))
	if err == nil && claim != nil && claim.ExpiresAt >= time.Now().Unix() &&
		claim.Identity == "admin" &&
		tools.ExistsHkey("admin", utils.IntToString(claim.ID)) &&
		c.GetHeader("token") == tools.GetHkey("admin", utils.IntToString(claim.ID)) {
		return claim.ID
	} else {
		return 0
	}
}

func Admin_Login_Verify(c *gin.Context) bool {
	type Verify struct {
		Account  string `validate:"required,max=15,min=1"`
		Password string `validate:"required,max=15,min=1"`
	}

	verify := &Verify{
		Account:  c.PostForm("account"),
		Password: c.PostForm("password"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Update_Admin_Detail_Verify(c *gin.Context) bool {
	type Verify struct {
		ShopName        string `validate:"max=20,min=1"`
		ShopInfo        string `validate:"max=100"`
		ShopImage       string ``
		CorporationName string `validate:"max=100"`
		ShopLocation    string `validate:"max=100"`
		OpenTime        string `validate:"max=100"`
		DayOff          string `validate:"max=100"`
		PhoneNumber     string `validate:"max=100"`
		Email           string `validate:"max=100"`
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
