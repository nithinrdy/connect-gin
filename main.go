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

func main() {
	db := config.DbConn()
	defer db.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	r := gin.Default()

	routes.AuthRoutes(r.Group("/api/auth"))

	r.Use(middleware.JwtAuth())

	r.GET("/", func(c *gin.Context) {
		userEmail, _ := c.Get("userEmail")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %v", userEmail),
		})
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
