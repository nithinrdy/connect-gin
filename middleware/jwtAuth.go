package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/nithinrdy/connect-gin/utilities"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authString := c.Request.Header.Get("Authorization")
		if len(authString) < 8 || authString[:7] != "Bearer " {
			c.JSON(401, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}
		tokenString := authString[7:]

		tk, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return utilities.SecretKey, nil
		})
		if err != nil {
			c.JSON(401, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		if !tk.Valid {
			c.JSON(401, gin.H{
				"message": tk.Claims.Valid().Error(),
			})
			c.Abort()
			return
		}

		userEmail, validAssertion := tk.Claims.(jwt.MapClaims)["dataToSign"]
		if !validAssertion {
			c.JSON(401, gin.H{
				"message": "Broken token (hehe)",
			})
			c.Abort()
			return
		}

		c.Set("userEmail", userEmail.(string))
		c.Next()
	}
}
