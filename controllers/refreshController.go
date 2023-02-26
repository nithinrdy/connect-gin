package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/nithinrdy/connect-gin/config"
	"github.com/nithinrdy/connect-gin/utilities"
)

func RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(401, "Refresh token not found")
		return
	}

	rows, err := config.DatabaseInstance.Query("SELECT username, email, nickname FROM users WHERE refresh_token = $1", refreshToken)
	if err != nil {
		c.JSON(500, "Error contacting database")
		return
	}

	var correspondingUsername string
	var email string
	var nickname string

	for rows.Next() {
		err3 := rows.Scan(&correspondingUsername, &email, &nickname)
		if err3 != nil {
			c.JSON(500, "Error processing query")
			return
		}
	}

	token, err := utilities.ValidateJWT(refreshToken)
	if err != nil {
		c.JSON(401, "Invalid refresh token")
		return
	}

	if !token.Valid {
		c.JSON(401, "Invalid refresh token")
		return
	}

	username, validAssertion := token.Claims.(jwt.MapClaims)["dataToSign"]
	if !validAssertion || username != correspondingUsername {
		c.JSON(401, "Refresh token is invalid/has been tampered with")
		return
	}

	accessTokenString, errAccess := utilities.GenerateAccessToken(correspondingUsername)
	if errAccess != nil {
		c.JSON(500, "Error generating access token")
		return
	}

	c.JSON(200, gin.H{
		"message": "Token refreshed",
		"user": gin.H{
			"username":    correspondingUsername,
			"email":       email,
			"nickname":    nickname,
			"accessToken": accessTokenString,
		},
	})

}
