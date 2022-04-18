package routes

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"stockApi/controller"
	"stockApi/middleware"
	"strings"
)

func Routes(router *gin.Engine) {

	router.GET("/", controller.Hello)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
	}

	user := router.Group("/user")
	{
		user.Use(middleware.IsAuth())
		user.GET("/", controller.Hello)
		user.POST("/changePassword", controller.ChangePassword)
		user.DELETE("/deleteUser", controller.DeleteUser)
	}
	wallet := router.Group("/wallet")
	{
		wallet.GET("/", controller.MicroService)
		wStock := wallet.Group("/stock")
		{
			wStock.Use(middleware.IsAuth())
			wStock.GET("/", controller.GetControler)
			wStock.POST("/", controller.CreateControler)
			wStock.PUT("/", controller.UpdateControler)
			wStock.DELETE("/", controller.DeleteControler)
		}

		wInfoStock := wallet.Group("/stock")
		wInfoStock.GET("/data/:name", func(c *gin.Context) {
			path := strings.Split(c.Request.URL.String(), "/wallet")
			log.Info("PATH ", os.Getenv("MICROSERVICE_URL")+path[1])
			c.Redirect(http.StatusMovedPermanently, os.Getenv("FETCH_URL")+os.Getenv("FETCH_PORT")+path[1])
		})

		wBroker := wallet.Group("/broker")
		{
			wBroker.Use(middleware.IsAdmin())
			wBroker.GET("/:broker", controller.GetControler)
			wBroker.GET("/", controller.GetControler)
			wBroker.POST("/", controller.CreateControler)
			wBroker.PUT("/", controller.UpdateControler)
			wBroker.DELETE("/", controller.DeleteControler)
		}

		wStockBroker := wallet.Group("/stockBroker/")
		{
			wStockBroker.Use(middleware.IsAdmin())
			wStockBroker.GET("/broker", controller.GetControler)
			wStockBroker.GET("/stock", controller.GetControler)
			wStockBroker.POST("/associate", controller.CreateControler)
			wStockBroker.DELETE("/disassociate", controller.DeleteControler)
		}

		wUser := wallet.Group("/user")
		{
			wUser.Use(middleware.IsAuth())
			wUser.GET("/", controller.GetControler)
			wUser.POST("/", controller.CreateControler)
			wUser.PUT("/", controller.UpdateControler)
			wUser.DELETE("/", controller.DeleteControler)

		}

		wPosition := wUser.Group("/positions")
		{
			wPosition.Use(middleware.IsAuth())
			wPosition.GET("/", controller.GetControler)
			wPosition.POST("/", controller.CreateControler)
			wPosition.DELETE("/", controller.DeleteControler)
			wPosition.POST("/closePosition", controller.CreateControler)

		}
		wResult := wUser.Group("/results")
		{
			wResult.Use(middleware.IsAuth())
			wResult.GET("/", func(c *gin.Context) {
				path := strings.Split(c.Request.URL.String(), "/wallet")
				c.Redirect(http.StatusMovedPermanently, os.Getenv("MICROSERVICE_URL")+path[1])
			})

		}

	}
}
