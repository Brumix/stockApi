package server

import (
	"github.com/gin-gonic/gin"
	"stockApi/routes"
)

func New() *gin.Engine {
	r := gin.Default()
	routes.Routes(r)

	return r
}
