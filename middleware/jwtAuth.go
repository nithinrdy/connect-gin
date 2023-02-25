package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/nithinrdy/connect-gin/utilities"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Split(c.Request.URL.Path, "/")[1] != "api" {
			c.Next()
			return
		}
		authString := c.Request.Header.Get("Authorization")
		if len(authString) < 8 || authString[:7] != "Bearer " {
			c.JSON(401, "unauthorized")
			c.Abort()
			return
		}
		tokenString := authString[7:]

		tk, err := utilities.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}
		if !tk.Valid {
			c.JSON(401, tk.Claims.Valid().Error())
			c.Abort()
			return
		}

		username, validAssertion := tk.Claims.(jwt.MapClaims)["dataToSign"]
		if !validAssertion {
			c.JSON(401, "Broken token (hehe)")
			c.Abort()
			return
		}

		c.Set("username", username.(string))
		c.Next()
	}
}
