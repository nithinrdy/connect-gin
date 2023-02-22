package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nithinrdy/connect-gin/config"
	"github.com/nithinrdy/connect-gin/utilities"
)

type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
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
	var registerData RegisterData

	if err := c.ShouldBindWith(&registerData, binding.JSON); err != nil {
		c.JSON(400, "Invalid register data credentials (missing fields/bad data)")
		return
	}

	accessTokenString, errAccess := utilities.GenerateAccessToken(registerData.Username)
	refreshTokenString, errRefresh := utilities.GenerateRefreshToken(registerData.Username)

	if errAccess != nil || errRefresh != nil {
		c.JSON(500, "Error while generating token")
		return
	}

	_, queryErr := config.DatabaseInstance.Exec("INSERT INTO users (username, email, nickname, password_hash, refresh_token) VALUES ($1, $2, $3, $4, $5) RETURNING id", registerData.Username, registerData.Email, registerData.Nickname, registerData.Password, refreshTokenString)

	var potentialDuplicateErrors map[string]string = map[string]string{
		"email":    "pq: duplicate key value violates unique constraint \"users_email_key\"",
		"username": "pq: duplicate key value violates unique constraint \"users_username_key\"",
		"token":    "pq: duplicate key value violates unique constraint \"users_refresh_token_key\"",
	}

	if queryErr != nil {
		for key, value := range potentialDuplicateErrors {
			if value == queryErr.Error() {
				c.JSON(400, fmt.Sprintf("Duplicate %v", key))
				return
			}
		}
	}

	c.SetCookie("refresh_token", refreshTokenString, 3600*24*30, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"email_received": registerData.Email,
		"username":       registerData.Username,
		"nickname":       registerData.Nickname,
		"access_token":   accessTokenString,
	})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout",
	})
}
