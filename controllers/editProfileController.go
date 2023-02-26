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

type editProfileRequestBody struct {
	EditProperty string `json:"editProperty"`
	EditValue    string `json:"editValue"`
}

func EditProfile(c *gin.Context) {
	username := ""
	tokenFromClient := ""
	username = c.MustGet("username").(string)
	tokenFromClient = c.MustGet("tokenFromClient").(string)
	if username == "" || tokenFromClient == "" {
		c.JSON(401, "Unauthorized")
		return
	}

	var editDetails editProfileRequestBody
	err := c.ShouldBindWith(&editDetails, binding.JSON)
	if err != nil {
		c.JSON(400, "Error reading request body. Could be invalid data")
		return
	}

	if editDetails.EditProperty == "" || editDetails.EditValue == "" {
		c.JSON(400, "Both editProperty and editValue are required")
		return
	}

	switch editDetails.EditProperty {
	case "email":
		_, err := mail.ParseAddress(editDetails.EditValue)
		if err != nil {
			c.JSON(400, "Invalid email")
			return
		}
		_, err = config.DatabaseInstance.Exec("UPDATE users SET email = $1 WHERE username = $2", editDetails.EditValue, username)
		if err != nil {
			c.JSON(500, "Error updating email")
			return
		}

	case "password":
		if !(utilities.IsPasswordValid(editDetails.EditValue)) {
			c.JSON(400, "Invalid password")
			return
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(editDetails.EditValue), 14)
		if err != nil {
			c.JSON(500, "Error hashing password")
			return
		}
		_, err = config.DatabaseInstance.Exec("UPDATE users SET password_hash = $1 WHERE username = $2", passwordHash, username)
		if err != nil {
			c.JSON(500, "Error updating password")
			return
		}

	case "nickname":
		_, err := config.DatabaseInstance.Exec("UPDATE users SET nickname = $1 WHERE username = $2", editDetails.EditValue, username)
		if err != nil {
			c.JSON(500, "Error updating nickname")
			return
		}
	}

	rows, err := config.DatabaseInstance.Query("SELECT email, nickname FROM users WHERE username = $1", username)
	if err != nil {
		fmt.Println("test1")
		c.JSON(500, "Error fetching updated data")
		return
	}

	var email, nickname string
	for rows.Next() {
		err = rows.Scan(&email, &nickname)
		if err != nil {
			fmt.Println("test2")
			c.JSON(500, "Error fetching updated data")
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "Profile successfully edited",
		"user": gin.H{
			"email":       email,
			"username":    username,
			"nickname":    nickname,
			"accessToken": tokenFromClient,
		},
	})
}
