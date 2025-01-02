package routes

import (
	controllesr "github.com/arjunbetor/Golang-Backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Context) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}
