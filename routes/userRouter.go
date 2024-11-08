package routes 

import (
	"github.com/gin-gonic/gin"
	 controller "QueueIt/controllers"
	 "QueueIt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.getUsers())
	incomingRoutes.GET("users/:user_id", controller.getUser())
}