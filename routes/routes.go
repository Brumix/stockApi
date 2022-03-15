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

	user := router.Group("/user")
	user.Use(middleware.IsAuth())
	{
		user.GET("/", controller.Hello)
		user.POST("/changePassword", controller.ChangePassword)
		user.DELETE("/deleteUser", controller.DeleteUser)
	}
	wallet := router.Group("/wallet")
	{
		wallet.GET("/", controller.MicroService)
		wUser := wallet.Group("/user")
		{
			wUser.GET("/", controller.CreateUser)

		}
	}
}
