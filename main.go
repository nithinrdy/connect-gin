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
