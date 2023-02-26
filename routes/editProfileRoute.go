package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nithinrdy/connect-gin/controllers"
)

func EditProfileRoute(routerGroup *gin.RouterGroup) {
	routerGroup.PUT("/", controllers.EditProfile)
}
