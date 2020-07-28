package utils

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

// SocketServer definition
var SocketServer *socketio.Server

// InitSocket handler for use by app
func InitSocket() error {
	var err error
	SocketServer, err = socketio.NewServer(nil)

	if err != nil {
		return err
	}

	return nil
}

// SocketEvents from websocket
func SocketEvents() {
	SocketServer.OnConnect("/", func(conn socketio.Conn) error {
		// conn.SetContext("")
		fmt.Println("connected:", conn.ID())
		return nil
	})

	SocketServer.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
	})

	SocketServer.OnEvent("/", "bye", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
		log.Println(s.Close())
	})

	SocketServer.OnError("/", func(conn socketio.Conn, err error) {
		fmt.Println("meet error:", err)
	})

	SocketServer.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		fmt.Println("closed:", reason)
	})
}
