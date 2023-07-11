package routes

import (
	controller "github.com/brf153/jwt-golang.git/controllers"
	"github.com/brf153/jwt-golang.git/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("./users/:user_id", controller.GetUser())
}
