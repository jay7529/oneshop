package controller

import (
	"oneshop/internal/model"
	"oneshop/internal/verify"
	"oneshop/middleware"
	"oneshop/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Admin_Login(c *gin.Context) {
	if !verify.Admin_Login_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	row := model.Select_Admin_Id([]interface{}{c.PostForm("account"), utils.MD5crypt(c.PostForm("password"))})

	if len(row) < 1 {
		utils.Failed(c, "", "ログインできません。アカウントのパスワードが正しいか確認してください。")
		return
	}

	model.Insert_Admin_LoginLog([]interface{}{row[0].AdminId, c.ClientIP()})

	//取得token
	token, _ := middleware.GenerateToken("admin", row[0].AdminId)
	utils.Success(c, token, nil, "Login Success")
}

func Admin_Update_Shop_Status(c *gin.Context) {
	admin_id := middleware.VerifyToken(c, "admin")
	if admin_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Admin_Update_Shop_Status_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Update_Shop_Status([]interface{}{c.PostForm("status"), c.PostForm("shop_id")})
	if id == 0 {
		newToken, _ := middleware.GenerateToken("admin", admin_id)
		utils.Success(c, newToken, nil, "Update Success")
		return
	}

	newToken, _ := middleware.GenerateToken("admin", admin_id)
	utils.Success(c, newToken, nil, "Update Success")
}

func Admin_Get_Shop_List(c *gin.Context) {
	admin_id := middleware.VerifyToken(c, "admin")
	if admin_id == 0 {
		utils.Error(c, "Token Error")
		return
	}

	row := model.Select_Shop_List([]interface{}{c.Param("shop_id")})

	newToken, _ := middleware.GenerateToken("admin", admin_id)
	utils.Success(c, newToken, map[string]interface{}{"shop": row}, "Success")
}
