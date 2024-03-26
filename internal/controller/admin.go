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

func Admin_Login(c *gin.Context) {
	if !verify.Admin_Login_Verify(c) {
		utils.Failed(c, "Parameter Error")
		return
	}

	row := model.Select_Admin_Id([]interface{}{c.PostForm("account"), utils.MD5crypt(c.PostForm("password"))})

	if len(row) < 1 {
		utils.Failed(c, "ログインできません。アカウントのパスワードが正しいか確認してください。")
		return
	}

	//取得token
	token, _ := middleware.GenerateToken("admin", row[0].Admin_id)

	// model.Insert_login_log([]interface{}{row[0]["user_id"], c.PostForm("account"), c.ClientIP()})

	utils.Success(c, map[string]interface{}{"token": token}, "Login Success")
}

func Get_Admin_Detail(c *gin.Context) {
	admin_id := verify.Admin_Token_Verify(c)
	if admin_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}

	row := model.Select_Admin_Detail([]interface{}{admin_id})

	newToken, _ := middleware.GenerateToken("admin", admin_id)
	tools.SetHkey("shop", utils.IntToString(admin_id), newToken)
	utils.Success(c, map[string]interface{}{"token": newToken, "admin": row}, "Success")
}

func Update_Admin_Detail(c *gin.Context) {
	admin_id := verify.Admin_Token_Verify(c)
	if admin_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}
	if !verify.Update_Admin_Detail_Verify(c) {
		utils.Failed(c, "Parameter Error")
		return
	}
	// p.UploadImage(c, "admin/"+convert.IntToString(admin_id)+"/", "test")
	model.Update_Admin_Detail([]interface{}{
		c.PostForm("shopName"), c.PostForm("shopInfo"), c.PostForm("shopImage"),
		c.PostForm("corporationName"), c.PostForm("shopLocation"), c.PostForm("openTime"),
		c.PostForm("dayOff"), c.PostForm("phoneNumber"), c.PostForm("email"), admin_id})

	newToken, _ := middleware.GenerateToken("admin", admin_id)
	tools.SetHkey("shop", utils.IntToString(admin_id), newToken)
	utils.Success(c, map[string]interface{}{"token": newToken}, "Update Success")
}

func Upload_Admin_Image(c *gin.Context) {
	admin_id := verify.Admin_Token_Verify(c)
	if admin_id == 0 {
		utils.Failed(c, "Token Error")
		return
	}

	utils.UploadImage(c, "admin/"+utils.IntToString(admin_id)+"/", utils.Int64ToString(time.Now().Unix()))
}
