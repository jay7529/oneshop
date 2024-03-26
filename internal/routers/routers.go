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
		shop.GET("/logout", controller.Shop_Logout)
		shop.GET("/detail", controller.Get_Shop_Detail)
		shop.PUT("/detail", controller.Update_Shop_Detail)
		shop.POST("/image", controller.Upload_Shop_Image)
		shop.GET("/shopimage/:shopID/:imageID", controller.Get_Shop_Image)
		shop.POST("/car", controller.Insert_Shop_Car)
		shop.PUT("/car", controller.Update_Shop_Car)
		shop.DELETE("/car", controller.Delete_Shop_Car)
		shop.GET("/car", controller.Get_Shop_Car)
		shop.GET("/carlist", controller.Get_Shop_Car_List)
	}

	user := app.Group("/user")
	{
		// user.POST("/login", controller.User_Login)
		user.GET("/shoplist", controller.User_Get_Shop_List)
		user.GET("/shopimage/:imageID", controller.User_Get_Shop_Image)

	}
}
