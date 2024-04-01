package verify

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Admin_Login_Verify(c *gin.Context) bool {
	type Verify struct {
		Account  string `validate:"required,max=15,min=1"`
		Password string `validate:"required,max=15,min=6"`
	}

	verify := &Verify{
		Account:  c.PostForm("account"),
		Password: c.PostForm("password"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}

func Admin_Update_Shop_Status_Verify(c *gin.Context) bool {
	type Verify struct {
		ShopId string `validate:"required,max=15,min=1"`
		Status string `validate:"required,max=15,min=1"`
	}

	verify := &Verify{
		ShopId: c.PostForm("shop_id"),
		Status: c.PostForm("status"),
	}

	err := validator.New().Struct(verify)
	return err == nil
}
