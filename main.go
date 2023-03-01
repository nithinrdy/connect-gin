package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nithinrdy/connect-gin/middleware"
	"github.com/nithinrdy/connect-gin/routes"

	"github.com/gin-gonic/gin"
	"github.com/nithinrdy/connect-gin/config"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Next()
	}
}

func main() {
	db := config.DbConn()
	defer db.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	r := gin.Default()

	r.Use(CORSMiddleware())

	routes.AuthRoutes(r.Group("/api/auth"))
	routes.RefreshRoute(r.Group("/api/refresh"))

	r.Use(middleware.JwtAuth())

	routes.EditProfileRoute(r.Group("/api/editProfile"))

	r.GET("/", func(c *gin.Context) {
		username, _ := c.Get("username")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %v", username),
		})
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
