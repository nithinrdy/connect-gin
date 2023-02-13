package controllers

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register",
	})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout",
	})
}
