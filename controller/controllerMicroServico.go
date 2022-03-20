package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stockApi/models"
	service "stockApi/services"
	"strings"
)

func MicroService(ctx *gin.Context) {
	var result models.Hello
	if err := service.HelloMicroService(&result); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": result.Msg})
}

func GetControler(ctx *gin.Context) {
	var requestBody models.UserRequest
	requestBody.Email = ctx.Query("email")
	path := strings.Split(ctx.Request.URL.String(), "/wallet")
	response, errResponse := service.ParseReponse(path[1], http.MethodGet, requestBody)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": errResponse.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateControler(ctx *gin.Context) {
	var requestBody interface{}
	err := ctx.Bind(&requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	path := strings.Split(ctx.Request.URL.String(), "/wallet")
	response, err := service.ParseReponse(path[1], http.MethodPost, requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func UpdateControler(ctx *gin.Context) {
	var requestBody interface{}
	err := ctx.Bind(&requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	path := strings.Split(ctx.Request.URL.String(), "/wallet")
	response, err := service.ParseReponse(path[1], http.MethodPut, requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func DeleteControler(ctx *gin.Context) {
	var requestBody interface{}
	err := ctx.Bind(&requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	path := strings.Split(ctx.Request.URL.String(), "/wallet")
	response, errResponse := service.ParseReponse(path[1], http.MethodDelete, requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": errResponse.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
