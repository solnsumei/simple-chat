package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/simple-chat/models"
)

// FetchAllUsers from database
func FetchAllUsers(c *gin.Context) {
	fmt.Println(c.MustGet("authID"))
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
