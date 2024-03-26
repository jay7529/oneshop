package controller

import (
	"oneshop/internal/model"
	"oneshop/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// func User_Login(c *gin.Context) {
// 	if !verify.Account_verify(c) || !verify.Password_verify(c) {
// 		utils.Failed(c, "Parameter Error")
// 		return
// 	}

// 	row := model.Select_User_Id([]interface{}{c.PostForm("account"), crypt.MD5crypt(c.PostForm("password"))})

// 	if len(row) < 1 {
// 		utils.Failed(c, "ログインできません。アカウントのパスワードが正しいか確認してください。")
// 		return
// 	}

// 	// model.Insert_login_log([]interface{}{row[0]["user_id"], c.PostForm("account"), c.ClientIP()})

// 	rsp.Success(c, row)
// }

func User_Get_Shop_List(c *gin.Context) {

	row := model.User_Get_Shop_List([]interface{}{})

	utils.Success(c, map[string]interface{}{"shop": row}, "Success")
}

func User_Get_Shop_Image(c *gin.Context) {
	path := "./uploads/shop/" + c.Param("shopId") + "/" + c.Param("imageId")
	utils.HandlerImage(c, path)
}
