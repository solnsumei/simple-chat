package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/simple-chat/config"
)

func init() {
	// runMigrations()
}

func main() {
	config, err := config.LoadConfigVars()

	if err != nil {
		panic("Please set .env config variables.")
	}

	router := gin.Default()

	loadRoutes(router)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to simple chat!"})
	})

	router.Run("localhost:" + config.PORT)
}
