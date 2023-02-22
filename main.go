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

	// var user models.UserModel
	// rows, err := db.Query("SELECT * FROM users")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for rows.Next() {
	// 	err2 := rows.Scan(&user.Id, &user.Created_at, &user.Username, &user.Email, &user.Nickname, &user.PasswordHash, &user.RefreshToken)
	// 	if err2 != nil {
	// 		fmt.Println(err2)
	// 	}
	// 	fmt.Println(user)
	// }
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	r := gin.Default()

	authGroup := r.Group("/auth")

	routes.AuthRoutes(authGroup)

	r.Use(middleware.JwtAuth())

	r.GET("/", func(c *gin.Context) {
		userEmail, _ := c.Get("userEmail")
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello %v", userEmail),
		})
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
