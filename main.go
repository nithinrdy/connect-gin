package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nithinrdy/connect-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	r := gin.Default()

	authGroup := r.Group("/auth")

	routes.AuthRoutes(authGroup)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
