package controller

import (
	"crypto/rand"
	"math/big"
	"oneshop/database"
	"oneshop/internal/model"
	"oneshop/internal/verify"
	"oneshop/middleware"
	"oneshop/utils"

	"time"

	"github.com/gin-gonic/gin"
)

func Shop_Signup(c *gin.Context) {
	if !verify.Shop_Signup_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	// 產生驗證碼
	code := ""
	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(9))
		code = code + utils.Int64ToString(n.Int64())
	}

	utils.SendEmail(
		"【OneShop】認証コード送信のお知らせ",
		`このメールは、 「OneShop」に関するご本人確認のため送付しています。
		登録を続けるには、次の認証コードを登録ページ内に入力してください。
		【認証コード】`+code+
			`
			※60分以内に手続きが完了しない場合は無効となります。

		───────────────────────────────────
		このメールにお心あたりのない方は、お手数ですがこのメールを削除してください。
		───────────────────────────────────`,
		c.PostForm("email"))

	database.Setkey(c.PostForm("email"), code, 60*time.Minute)

	utils.Success(c, "", nil, "Success")
}

func Shop_Signup_Code(c *gin.Context) {
	if !verify.Shop_Code_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}
	if !database.Existskey(c.PostForm("email")) || c.PostForm("code") != database.Getkey(c.PostForm("email")) {
		utils.Failed(c, "", "認証コードが正しくないか、有効期限が切れています。")
		return
	}

	utils.Success(c, "", nil, "SignUp Success")
}

func Shop_New_Signup(c *gin.Context) {
	if !verify.Shop_New_Signup_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}
	if !database.Existskey(c.PostForm("email")) || c.PostForm("code") != database.Getkey(c.PostForm("email")) {
		utils.Failed(c, "", "認証コードが正しくないか、有効期限が切れています。")
		return
	}

	database.Delkey(c.PostForm("email"))

	id := model.Insert_Shop([]interface{}{
		c.PostForm("email"), utils.MD5crypt(c.PostForm("password")), c.PostForm("email")})
	if id == 0 {
		utils.Failed(c, "", "SinUp Failed, Account Exists")
		return
	}
	id = model.Update_Shop_Detail_FirstTime([]interface{}{
		c.PostForm("postCode"), c.PostForm("address"), c.PostForm("phoneNumber"), c.PostForm("shopName"), c.PostForm("email"), id})
	if id == 0 {
		utils.Failed(c, "", "SinUp Failed")
		return
	}

	//取得token
	token, _ := middleware.GenerateToken("shop", id)
	utils.Success(c, token, nil, "SignUp Success")
}

func Shop_Login(c *gin.Context) {
	if !verify.Shop_Login_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	row := model.Select_Shop_Id([]interface{}{c.PostForm("account"), utils.MD5crypt(c.PostForm("password"))})

	if len(row) < 1 {
		utils.Failed(c, "", "ログインできません。アカウントのパスワードが正しいか確認してください。")
		return
	}

	// 新增登入記錄
	model.Insert_Shop_LoginLog([]interface{}{row[0].ShopId, c.ClientIP()})

	//取得token
	token, _ := middleware.GenerateToken("shop", row[0].ShopId)
	utils.Success(c, token, nil, "Login Success")
}

func Shop_Logout(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}

	database.DelHkey("shop", utils.IntToString(shop_id))

	utils.Success(c, "", nil, "Logout Success")
}

func Shop_Forget_Password(c *gin.Context) {
	if !verify.Shop_Forget_Password_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	// 產生驗證碼
	code := ""
	for i := 0; i < 6; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(9))
		code = code + utils.Int64ToString(n.Int64())
	}

	utils.SendEmail(
		"【OneShop】認証コード送信のお知らせ",
		`このメールは、 「OneShop」に関するご本人確認のため送付しています。
			登録を続けるには、次の認証コードを登録ページ内に入力してください。
			【認証コード】`+code+
			`
				※60分以内に手続きが完了しない場合は無効となります。
	
			───────────────────────────────────
			このメールにお心あたりのない方は、お手数ですがこのメールを削除してください。
			───────────────────────────────────`,
		c.PostForm("email"))

	database.Setkey(c.PostForm("email"), code, 60*time.Minute)

	utils.Success(c, "", nil, "Success")
}

func Shop_Reset_Password_Code(c *gin.Context) {
	if !verify.Shop_Code_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}
	if !database.Existskey(c.PostForm("email")) || c.PostForm("code") != database.Getkey(c.PostForm("email")) {
		utils.Failed(c, "", "認証コードが正しくないか、有効期限が切れています。")
		return
	}

	row := model.Select_Shop_Password([]interface{}{c.PostForm("email")})
	if len(row) == 0 {
		utils.Error(c, "Error")
		return
	}

	utils.Success(c, "", row[0].Password, "Success")
}

func Shop_Reset_Password(c *gin.Context) {
	if !verify.Reset_Shop_Password_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Reset_Shop_Password([]interface{}{
		utils.MD5crypt(c.PostForm("newPassword")), c.PostForm("email"), c.PostForm("oldPassword")})

	if id == 0 {
		utils.Failed(c, "", "Reset Failed")
		return
	}

	utils.Success(c, "", nil, "Reset Success")
}

func Update_Shop_Password(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Update_Shop_Password_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Update_Shop_Password([]interface{}{
		utils.MD5crypt(c.PostForm("oldPassword")), utils.MD5crypt(c.PostForm("newPassword")), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)

	if id == 0 {
		utils.Failed(c, newToken, "Update Failed")
		return
	}

	utils.Success(c, newToken, nil, "Update Success")
}

func Get_Shop_Detail(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}

	row := model.Get_Shop_Detail([]interface{}{shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	utils.Success(c, newToken, map[string]interface{}{"shop_detail": row}, "Success")
}

func Update_Shop_Detail(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Update_Shop_Detail_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Update_Shop_Detail([]interface{}{
		c.PostForm("shopName"), c.PostForm("shopInfo"), c.PostForm("shopImage"),
		c.PostForm("corporationName"), c.PostForm("postCode"), c.PostForm("shopLocation"), c.PostForm("shopCity"),
		c.PostForm("openTime"), c.PostForm("dayOff"), c.PostForm("phoneNumber"), c.PostForm("email"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	if id == 0 {
		utils.Failed(c, newToken, "Update Failed")
		return
	}

	utils.Success(c, newToken, nil, "Update Success")
}

func Upload_Shop_Image(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}

	utils.UploadImage(c, "shop/", utils.IntToString(shop_id)+"/", utils.Int64ToString(time.Now().Unix()))
}

func Get_Shop_Image(c *gin.Context) {
	path := "./uploads/shop/" + c.Param("shop_id") + "/" + c.Param("imageId")
	utils.HandlerImage(c, path)
}

func Insert_Shop_Car(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Insert_Shop_Car_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Insert_Shop_Car([]interface{}{
		shop_id, c.PostForm("carName"), c.PostForm("carBrand"), c.PostForm("carImage"),
		c.PostForm("carPrice"), c.PostForm("carFee"), c.PostForm("carYear")})

	newToken, _ := middleware.GenerateToken("shop", shop_id)

	if id == 0 {
		utils.Failed(c, newToken, "Insert Failed")
		return
	}

	utils.Success(c, newToken, nil, "Insert Success")
}

func Update_Shop_Car(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Update_Shop_Car_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Update_Shop_Car([]interface{}{
		c.PostForm("carName"), c.PostForm("carBrand"), c.PostForm("carImage"),
		c.PostForm("carPrice"), c.PostForm("carFee"), c.PostForm("carYear"), c.PostForm("shelves"),
		c.Param("car_id"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)

	if id == 0 {
		utils.Failed(c, newToken, "Update Failed")
		return
	}

	utils.Success(c, newToken, nil, "Update Success")
}

func Delete_Shop_Car(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Delete_Shop_Car_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Delete_Shop_Car([]interface{}{c.Param("car_id"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)

	if id == 0 {
		utils.Failed(c, newToken, "Delete Failed")
		return
	}
	utils.Success(c, newToken, nil, "Delete Success")
}

func Get_Shop_Car_List(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}

	row := model.Get_Shop_Car_List([]interface{}{shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	utils.Success(c, newToken, map[string]interface{}{"car": row}, "Success")
}

func Get_Shop_Car(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Get_Shop_Car_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	row := model.Get_Shop_Car([]interface{}{c.Param("car_id"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	utils.Success(c, newToken, map[string]interface{}{"car": row}, "Success")
}

func Insert_Shop_Staff(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Insert_Shop_Staff_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Insert_Shop_Staff([]interface{}{
		shop_id, c.PostForm("staffName"), c.PostForm("staffImage"),
		c.PostForm("staffPosition"), c.PostForm("staffIntroduction")})

	newToken, _ := middleware.GenerateToken("shop", shop_id)

	if id == 0 {
		utils.Failed(c, newToken, "Insert Failed")
		return
	}

	utils.Success(c, newToken, nil, "Insert Success")
}

func Update_Shop_Staff(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Update_Shop_Staff_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Update_Shop_Staff([]interface{}{
		c.PostForm("staffName"), c.PostForm("staffImage"),
		c.PostForm("staffPosition"), c.PostForm("staffIntroduction"),
		c.Param("staff_id"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)

	if id == 0 {
		utils.Failed(c, newToken, "Update Failed")
		return
	}

	utils.Success(c, newToken, nil, "Update Success")
}

func Delete_Shop_Staff(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}
	if !verify.Delete_Shop_Staff_Verify(c) {
		utils.Error(c, "Parameter Error")
		return
	}

	id := model.Delete_Shop_Staff([]interface{}{c.Param("staff_id"), shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)

	if id == 0 {
		utils.Failed(c, newToken, "Delete Failed")
		return
	}
	utils.Success(c, newToken, nil, "Delete Success")
}

func Get_Shop_Staff_List(c *gin.Context) {
	shop_id := middleware.VerifyToken(c, "shop")
	if shop_id == 0 {
		utils.Error(c, "Token Error")
		return
	}

	row := model.Get_Shop_Staff_List([]interface{}{shop_id})

	newToken, _ := middleware.GenerateToken("shop", shop_id)
	utils.Success(c, newToken, map[string]interface{}{"staff": row}, "Success")
}
