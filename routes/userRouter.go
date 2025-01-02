package routes

import (
	"github.com/arjunbetor/Golang-Backend/controllers"
	"github.com/arjunbetor/Golang-Backend/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Context) {
	incomingRoutes.Use(middlewares.Authenticated())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:userId", controllers.GetUserById())
}
