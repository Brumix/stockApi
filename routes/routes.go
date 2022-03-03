package routes

import (
	"github.com/gin-gonic/gin"
	"stockApi/controller"
	"stockApi/middleware"
)

func Routes(router *gin.Engine) {

	router.GET("/", controller.Hello)

	auth := router.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}

	stocks := router.Group("/stocks")
	stocks.Use(middleware.IsAuth())
	{
		stocks.GET("/", controller.Hello)
		stocks.POST("/changePassword", controller.ChangePassword)
		stocks.DELETE("/deleteUser", controller.DeleteUser)
	}
}
