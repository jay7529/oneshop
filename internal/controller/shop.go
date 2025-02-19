package controller

import (
	"oneshop/internal/model"
	"oneshop/internal/verify"
	"oneshop/middleware"
	"oneshop/tools"
	"oneshop/utils"

	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Shop_Login(c *gin.Context) {
	if !verify.Shop_Login_Verify(c) {
		utils.Failed(c, "Parameter Error")
		return
	}

	row := model.Select_Shop_Id([]interface{}{c.PostForm("account"), utils.MD5crypt(c.PostForm("password"))})

	if len(row) < 1 {
		utils.Failed(c, "ログインできません。アカウントのパスワードが正しいか確認してください。")
		return
	}

	//取得token
	token, _ := middleware.GenerateToken("shop", row[0].Shop_id)
	tools.SetHkey("shop", utils.IntToString(row[0].Shop_id), token)

	// 新增登入記錄
	// model.Insert_login_log([]interface{}{row[0]["user_id"], c.PostForm("account"), c.ClientIP()})

	utils.Success(c, map[string]interface{}{"token": token}, "Login Success")
}

func Shop_Logout(c *gin.Context) {
	shop_id := verify.Shop_Token_Verify(c)
	if shop_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}

	tools.DelHkey("shop", utils.IntToString(shop_id))

	// 新增登出記錄
	// model.Insert_login_log([]interface{}{row[0]["user_id"], c.PostForm("account"), c.ClientIP()})

	utils.Success(c, nil, "Logout Success")
}

func Get_Shop_Detail(c *gin.Context) {
	shop_id := verify.Shop_Token_Verify(c)
	if shop_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}

	row := model.Get_Shop_Detail([]interface{}{shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	tools.SetHkey("shop", utils.IntToString(shop_id), newToken)
	utils.Success(c, map[string]interface{}{"token": newToken, "shop_detail": row}, "Success")
}

func Update_Shop_Detail(c *gin.Context) {
	shop_id := verify.Shop_Token_Verify(c)
	if shop_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}
	if !verify.Update_Shop_Detail_Verify(c) {
		utils.Failed(c, "Parameter Error")
		return
	}

	model.Update_Shop_Detail([]interface{}{
		c.PostForm("shopName"), c.PostForm("shopInfo"), c.PostForm("shopImage"),
		c.PostForm("corporationName"), c.PostForm("shopLocation"), c.PostForm("openTime"),
		c.PostForm("dayOff"), c.PostForm("phoneNumber"), c.PostForm("email"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	tools.SetHkey("shop", utils.IntToString(shop_id), newToken)
	utils.Success(c, map[string]interface{}{"token": newToken}, "Update Success")
}

func Upload_Shop_Image(c *gin.Context) {
	shop_id := verify.Shop_Token_Verify(c)
	if shop_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}

	utils.UploadImage(c, "shop/"+utils.IntToString(shop_id)+"/", utils.Int64ToString(time.Now().Unix()))
}

func Get_Shop_Image(c *gin.Context) {
	path := "./uploads/shop/" + c.Param("shopID") + "/" + c.Param("imageID")
	c.File(path)
}

func Insert_Shop_Car(c *gin.Context) {
	shop_id := verify.Shop_Token_Verify(c)
	if shop_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}
	if !verify.Insert_Shop_Car_Verify(c) {
		utils.Failed(c, "Parameter Error")
		return
	}

	model.Insert_Shop_Car([]interface{}{
		shop_id, c.PostForm("carName"), c.PostForm("carBrand"), c.PostForm("carImage"),
		c.PostForm("carPrice"), c.PostForm("carFee")})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	tools.SetHkey("shop", utils.IntToString(shop_id), newToken)
	utils.Success(c, map[string]interface{}{"token": newToken}, "Insert Success")
}

func Update_Shop_Car(c *gin.Context) {
	shop_id := verify.Shop_Token_Verify(c)
	if shop_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}
	if !verify.Update_Shop_Car_Verify(c) {
		utils.Failed(c, "Parameter Error")
		return
	}

	model.Update_Shop_Car([]interface{}{
		c.PostForm("carName"), c.PostForm("carBrand"), c.PostForm("carImage"),
		c.PostForm("carPrice"), c.PostForm("carFee"), c.PostForm("shelves"),
		c.PostForm("carId"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	tools.SetHkey("shop", utils.IntToString(shop_id), newToken)
	utils.Success(c, map[string]interface{}{"token": newToken}, "Update Success")
}

func Delete_Shop_Car(c *gin.Context) {
	shop_id := verify.Shop_Token_Verify(c)
	if shop_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}
	if !verify.Delete_Shop_Car_Verify(c) {
		utils.Failed(c, "Parameter Error")
		return
	}

	model.Delete_Shop_Car([]interface{}{c.GetHeader("carId"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	tools.SetHkey("shop", utils.IntToString(shop_id), newToken)
	utils.Success(c, map[string]interface{}{"token": newToken}, "Delete Success")
}

func Get_Shop_Car(c *gin.Context) {
	shop_id := verify.Shop_Token_Verify(c)
	if shop_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}
	if !verify.Get_Shop_Car_Verify(c) {
		utils.Failed(c, "Parameter Error")
		return
	}

	row := model.Get_Shop_Car([]interface{}{c.GetHeader("carId"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	tools.SetHkey("shop", utils.IntToString(shop_id), newToken)
	utils.Success(c, map[string]interface{}{"token": newToken, "car": row}, "Success")
}

func Get_Shop_Car_List(c *gin.Context) {
	shop_id := verify.Shop_Token_Verify(c)
	if shop_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}

	row := model.Get_Shop_Car_List([]interface{}{c.GetHeader("carId"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	tools.SetHkey("shop", utils.IntToString(shop_id), newToken)
	utils.Success(c, map[string]interface{}{"token": newToken, "car": row}, "Success")
}
