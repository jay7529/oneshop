package controller

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

func OnConnect(s socketio.Conn) error {
	s.Join("dddd")
	s.SetContext("")
	fmt.Println("connected:", s.ID())
	return nil
}

func OnError(s socketio.Conn, e error) {
	fmt.Println("meet error:", e)
}

func OnDisconnect(s socketio.Conn, reason string) {
	fmt.Println("closed", reason)
}

func OnSingleChat(s socketio.Conn, message string) string {
	s.SetContext(message)
	fmt.Println(message)
	return message
	// s.Emit("SINGLE_CHAT", "收到的消息是"+message.Content)
}

func OnBye(s socketio.Conn) string {
	last := s.Context().(string)
	fmt.Println(last)
	s.Emit("BYE", last)
	// s.Close()
	return last
}
