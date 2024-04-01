package routers

import (
	"oneshop/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {

	admin := app.Group("/admin")
	{
		admin.POST("/login", controller.Admin_Login)

		admin.GET("/shopApply", controller.Admin_Get_Shop_Apply_List)
		admin.PUT("/shopStatus", controller.Admin_Update_Shop_Status)
		admin.GET("/detail", controller.Get_Shop_Detail)
	}

	shop := app.Group("/shop")
	{
		shop.POST("singup", controller.Shop_Singup)
		shop.POST("singupCode", controller.Shop_Singup_Code)

		shop.POST("/login", controller.Shop_Login)
		shop.DELETE("/logout", controller.Shop_Logout)

		shop.POST("/passwordCode", controller.Shop_Forget_Password)
		shop.POST("/resetPassword", controller.Shop_Reset_Password_Code)
		shop.PUT("/resetPassword", controller.Shop_Reset_Password)

		shop.PUT("/updatePassword", controller.Update_Shop_Password)

		shop.GET("/detail", controller.Get_Shop_Detail)
		shop.PUT("/detail", controller.Update_Shop_Detail)

		shop.GET("/car", controller.Get_Shop_Car_List)
		shop.POST("/car", controller.Insert_Shop_Car)
		shop.GET("/car/:car_id", controller.Get_Shop_Car)
		shop.PUT("/car/:car_id", controller.Update_Shop_Car)
		shop.DELETE("/car/:car_id", controller.Delete_Shop_Car)

		shop.POST("/image", controller.Upload_Shop_Image)
		shop.GET("/image/:shop_id/:imageId", controller.Get_Shop_Image)
	}

	user := app.Group("/user")
	{
		// user.POST("/login", controller.User_Login)

		user.GET("/shop", controller.User_Get_Shop_List)
		user.GET("/shop/:shop_id", controller.User_Get_Shop)

		user.GET("/car", controller.User_Get_Shop_Car_List)

		user.GET("/image/:shop_id/:imageId", controller.User_Get_Shop_Image)
	}
}
