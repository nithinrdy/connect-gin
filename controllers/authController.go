package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nithinrdy/connect-gin/utilities"
)

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginCreds LoginCredentials

	if err := c.ShouldBindWith(&loginCreds, binding.JSON); err != nil {
		c.JSON(400, "Invalid login credentials (missing fields/bad data)")
		return
	}

	accessTokenString, errAccess := utilities.GenerateAccessToken(loginCreds.Email)
	refreshTokenString, errRefresh := utilities.GenerateRefreshToken(loginCreds.Email)

	if errAccess != nil || errRefresh != nil {
		c.JSON(500, "Error while generating token")
		return
	}

	c.JSON(200, gin.H{
		"email_received":    loginCreds.Email,
		"password_received": loginCreds.Password,
		"access_token":      accessTokenString,
		"refresh_token":     refreshTokenString,
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
