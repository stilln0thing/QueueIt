package routes

import (
	"QueueIt/controllers"
	"github.com/gin-gonic/gin"
)

func BusinessRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("business", controllers.CreateBusiness())
	incomingRoutes.GET("business", controllers.GetAllBusinesses())
	incomingRoutes.GET("business/:business_id", controllers.GetBusiness())
 }




