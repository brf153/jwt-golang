package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/brf153/jwt-golang.git/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())

	// engine.Use(incomingRoutes.RegisterRoutes())
}
