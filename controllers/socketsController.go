package controllers

import (
	"encoding/json"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/solnsumei/simple-chat/models"
	"github.com/solnsumei/simple-chat/utils"
	"log"
	"strconv"
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
		conn.SetContext("")
		fmt.Println("connected:", conn.ID())
		conn.Join("bbm")
		return nil
	})

	SocketServer.OnEvent("/", "message", func(conn socketio.Conn, msg string) {
		var input utils.MessageInput

		if err := json.Unmarshal([]byte(msg), &input); err != nil {
			fmt.Println(err)
			return
		}

		if err := input.Validate(); err != nil {
			fmt.Println(err)
			return
		}

		chatID, _ := strconv.Atoi(input.ChatID)
		receiverID, _ := strconv.Atoi(input.ReceiverID)

		message := models.Message{
			Body: input.Message,
			ChatID: uint(chatID),
			ReceiverID: uint(receiverID),
			IsRead: false,
		}

		if err := models.DB.Create(&message).Error; err != nil {
			fmt.Println(err)
		}

		event := "message" + input.ReceiverID
		conn.SetContext("")
		SocketServer.BroadcastToRoom("/", "bbm", event, msg)
		// fmt.Println("<<<<<<", input)
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
