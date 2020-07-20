package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/simple-chat/models"
	"github.com/solnsumei/simple-chat/utils"
)

func init() {
	// runMigrations()
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

	loadGuestRoutes(router)
	loadAuthRoutes(router)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to simple chat!"})
	})

	router.Run("localhost:" + config.Port)
}
