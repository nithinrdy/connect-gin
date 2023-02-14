package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nithinrdy/connect-gin/controllers"
)

func AuthRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/login", controllers.Login)
	routerGroup.GET("/register", controllers.Register)
	routerGroup.GET("/logout", controllers.Logout)
}
