package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"stockApi/controller"
	"stockApi/middleware"
	"strings"
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
		wStock := wallet.Group("/stock")
		{
			wStock.GET("/", controller.GetControler)
			wStock.POST("/", controller.CreateControler)
			wStock.PUT("/", controller.UpdateControler)
			wStock.DELETE("/", controller.DeleteControler)

		}
		wBroker := wallet.Group("/broker")
		{
			wBroker.GET("/:broker", controller.GetControler)
			wBroker.GET("/", controller.GetControler)
			wBroker.POST("/", controller.CreateControler)
			wBroker.PUT("/", controller.UpdateControler)
			wBroker.DELETE("/", controller.DeleteControler)

		}
		wStockBroker := wallet.Group("/stockBroker/")
		{
			wStockBroker.GET("/broker", controller.GetControler)
			wStockBroker.GET("/stock", controller.GetControler)
			wStockBroker.POST("/associate", controller.CreateControler)
			wStockBroker.DELETE("/disassociate", controller.DeleteControler)
		}
		wUser := wallet.Group("/user")
		{
			wUser.GET("/", controller.GetControler)
			wUser.POST("/", controller.CreateControler)
			wUser.PUT("/", controller.UpdateControler)
			wUser.DELETE("/", controller.DeleteControler)

		}

		wPosition := wUser.Group("/positions")
		{
			wPosition.GET("/", controller.GetControler)
			wPosition.POST("/", controller.CreateControler)
			wPosition.DELETE("/", controller.DeleteControler)
			wPosition.POST("/closePosition", controller.CreateControler)

		}
		wResult := wUser.Group("/results")
		{
			wResult.GET("/", func(c *gin.Context) {
				path := strings.Split(c.Request.URL.String(), "/wallet")
				c.Redirect(http.StatusMovedPermanently, os.Getenv("MICROSERVICE_URL")+path[1])
			})

		}

	}
}
