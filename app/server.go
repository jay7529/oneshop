package app

import (
	"oneshop/middleware"
	"oneshop/routers"

	"github.com/gin-gonic/gin"
)

func InitGin() {

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(middleware.Cors)
	app.TrustedPlatform = "Client-IP"

	routers.RegisterRouter(app)

	// // socket io
	// 	server := socketio.NewServer(nil)
	// go server.Serve()
	// defer server.Close()
	// routers.RegisterSocket(app, server)

	app.Run()

}
