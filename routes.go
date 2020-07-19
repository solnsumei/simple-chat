package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api/v1")

	apiRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to simple chat api"})
	})
}
