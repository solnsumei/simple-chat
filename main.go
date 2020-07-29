package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	utils.SocketEvents()

	loadGuestRoutes(router)
	loadAuthRoutes(router)

	go utils.SocketServer.Serve()
	defer utils.SocketServer.Close()

	socketHandler(router)

	// router.StaticFS("/public", http.Dir("../asset"))

	server := &http.Server{Addr: "localhost:" + config.Port, Handler: router}

	log.Fatal(server.ListenAndServe())
}
