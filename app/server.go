package app

import (
	"oneshop/internal/routers"
	"oneshop/middleware"

	"github.com/gin-gonic/gin"
)

func InitGin() {

	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(middleware.Cors)
	app.TrustedPlatform = "Client-IP"

	// server := socketio.NewServer(nil)
	// go server.Serve()
	// defer server.Close()

	routers.RegisterRouter(app)
	// routers.RegisterSocket(app, server)

	app.Run(":8000")

}
