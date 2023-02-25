package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nithinrdy/connect-gin/controllers"
)

func RefreshRoute(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.RefreshToken)
}
