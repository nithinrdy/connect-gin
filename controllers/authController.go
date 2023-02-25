package controllers

import (
	"fmt"

	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/nithinrdy/connect-gin/config"
	"github.com/nithinrdy/connect-gin/utilities"
	"golang.org/x/crypto/bcrypt"
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
	err := c.ShouldBindWith(&loginCreds, binding.JSON)
	if err != nil {
		c.JSON(400, "Invalid login credentials (missing fields/bad data)")
		return
	}

	rows, err2 := config.DatabaseInstance.Query("SELECT username, password_hash FROM users WHERE email = $1", loginCreds.Email)
	if err2 != nil {
		c.JSON(500, "Error contacting database")
		return
	}

	var username string
	var passwordHash string

	for rows.Next() {
		err3 := rows.Scan(&username, &passwordHash)
		if err3 != nil {
			c.JSON(500, "Error processing password")
			return
		}
	}

	err4 := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginCreds.Password))
	if err4 != nil {
		c.JSON(400, "Invalid login credentials")
		return
	}

	accessTokenString, errAccess := utilities.GenerateAccessToken(username)
	refreshTokenString, errRefresh := utilities.GenerateRefreshToken(username)

	if errAccess != nil || errRefresh != nil {
		c.JSON(500, "Error generating tokens")
		return
	}

	_, err5 := config.DatabaseInstance.Exec("UPDATE users SET refresh_token = $1 WHERE email = $2", refreshTokenString, loginCreds.Email)
	if err5 != nil {
		c.JSON(500, "Error updating refresh token")
		return
	}

	c.SetCookie("refresh_token", refreshTokenString, 3600*24*30, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"message":      "Logged in",
		"username":     username,
		"email":        loginCreds.Email,
		"access_token": accessTokenString,
	})
}

func Register(c *gin.Context) {
	var registerData RegisterData

	if err := c.ShouldBindWith(&registerData, binding.JSON); err != nil {
		c.JSON(400, "Invalid register data credentials (missing fields/bad data)")
		return
	}

	_, err := mail.ParseAddress(registerData.Email)
	if err != nil {
		c.JSON(400, "Invalid email")
		return
	}

	if !utilities.IsPasswordValid(registerData.Password) {
		c.JSON(400, "Invalid password")
		return
	}

	if len(registerData.Username) < 4 {
		c.JSON(400, "Username too short")
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), 14)
	if err != nil {
		c.JSON(500, "Error while hashing password")
		return
	}
	registerData.Password = string(hashedPass)

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
		"message":        "Registered",
		"email_received": registerData.Email,
		"username":       registerData.Username,
		"nickname":       registerData.Nickname,
		"access_token":   accessTokenString,
	})
}

func Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.Status(204)
		return
	}

	_, err2 := config.DatabaseInstance.Exec("UPDATE users SET refresh_token = '' WHERE refresh_token = $1", refreshToken)
	if err2 != nil {
		c.JSON(500, "Error resetting refresh token")
		return
	}
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"message": "Logged out",
	})
}
