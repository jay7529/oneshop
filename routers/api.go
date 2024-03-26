package routers

import (
	"oneshop/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) {

	admin := app.Group("/admin")
	{
		admin.POST("/login", controller.Admin_Login)

		admin.GET("/detail", controller.Get_Admin_Detail)
		admin.PUT("/detail", controller.Update_Admin_Detail)

		admin.POST("/image", controller.Upload_Admin_Image)
	}

	shop := app.Group("/shop")
	{
		shop.POST("/login", controller.Shop_Login)
		shop.DELETE("/logout", controller.Shop_Logout)

		shop.GET("/detail", controller.Get_Shop_Detail)
		shop.PUT("/detail", controller.Update_Shop_Detail)

		shop.GET("/car", controller.Get_Shop_Car_List)
		shop.POST("/car", controller.Insert_Shop_Car)
		shop.GET("/car/:carId", controller.Get_Shop_Car)
		shop.PUT("/car/:carId", controller.Update_Shop_Car)
		shop.DELETE("/car/:carId", controller.Delete_Shop_Car)

		shop.POST("/image", controller.Upload_Shop_Image)
		shop.GET("/image/:shopId/:imageId", controller.Get_Shop_Image)
	}

	user := app.Group("/user")
	{
		// user.POST("/login", controller.User_Login)
		user.GET("/shop", controller.User_Get_Shop_List)
		user.GET("/image/:shopId/:imageId", controller.User_Get_Shop_Image)
	}
}
