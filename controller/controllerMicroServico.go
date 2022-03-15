package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stockApi/models"
	service "stockApi/services"
)

func MicroService(ctx *gin.Context) {
	var result models.Hello
	if err := service.HelloMicroService(&result); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": result.Msg})
}

func CreateUser(ctx *gin.Context) {
	var result models.Hello
	if err := service.HelloMicroService(&result); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": result.Msg + "test"})
}
