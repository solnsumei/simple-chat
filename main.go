package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/solnsumei/simple-chat/models"
	"github.com/solnsumei/simple-chat/utils"
)

func init() {
	runMigrations()
}

func main() {
	config, err := utils.LoadConfigVars()

	if err != nil {
		panic("Please set .env config variables.")
	}

	if err := models.ConnectDatabase(config); err != nil {
		panic(err)
	}

	router := gin.Default()

	if err := utils.InitSocket(); err != nil {
		panic(err)
	}

	utils.SocketServer.OnConnect("/", func(conn socketio.Conn) error {
		// conn.SetContext("")
		fmt.Println("connected:", conn.ID())
		return nil
	})

	utils.SocketServer.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
	})

	utils.SocketServer.OnEvent("/", "bye", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
		s.Close()
	})

	utils.SocketServer.OnError("/", func(conn socketio.Conn, err error) {
		fmt.Println("meet error:", err)
	})

	utils.SocketServer.OnDisconnect("/", func(conn socketio.Conn, reason string) {
		fmt.Println("closed:", reason)
	})

	loadGuestRoutes(router)
	loadAuthRoutes(router)

	go utils.SocketServer.Serve()
	defer utils.SocketServer.Close()

	socketHandler(router)

	// router.StaticFS("/public", http.Dir("../asset"))

	server := &http.Server{Addr: "localhost:" + config.Port, Handler: router}

	server.ListenAndServe()
}
