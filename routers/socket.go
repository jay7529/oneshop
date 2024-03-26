package routers

import (
	"fmt"
	"net/http"
	"oneshop/internal/controller"

	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

func RegisterSocket(engine *gin.Engine, server *socketio.Server) {
	engine.GET("/socket.io/*any", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)
	}))
	engine.POST("/socket.io/*any", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)
	}))
	engine.GET("/aa", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		server.BroadcastToNamespace("/", "chat", "test")
	}))
	server.OnConnect("/", controller.OnConnect)
	server.OnError("/", controller.OnError)
	server.OnDisconnect("/", controller.OnDisconnect)
	server.OnEvent("/", "chat", controller.OnSingleChat)
	server.OnEvent("/", "bye", controller.OnBye)
	server.OnEvent("/", "join", func(s socketio.Conn, room string) {
		s.Join("dddd")
		// msg := "4154"
		fmt.Println("/:join", room, s.Namespace(), s.Rooms())
		server.BroadcastToNamespace("/", "chat", "msg")
	})
}
