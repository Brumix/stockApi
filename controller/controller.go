package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"stockApi/errors"
	"stockApi/models"
	"stockApi/services"
	service "stockApi/services"
	"stockApi/services/core"
	"time"
)

func Hello(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, models.Hello{
		Msg: "Hyper Media Api Stocks " + os.Getenv("NODE"),
	})
}

func Register(ctx *gin.Context) {
	var newUser core.RegisterRequest
	errors.ErrorControler("Binding the userDto", ctx.Bind(&newUser))
	if err := services.RegisterUser(newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"msg": "User register with success!!"})
	return
}

func Login(ctx *gin.Context) {
	var user core.LoginRequest
	errors.ErrorControler("Binding the userDto", ctx.Bind(&user))
	response, err := services.LoginUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	cookie := &http.Cookie{
		Name:     "authorization",
		Value:    response.Token,
		Path:     "/",
		Expires:  time.Now().Add(2 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(ctx.Writer, cookie)
	ctx.JSON(http.StatusOK, gin.H{"msg": "User logged in with success!!"})
	return
}
func ChangePassword(ctx *gin.Context) {
	var replacePassword core.ChangePasswordRequest
	errors.ErrorControler("Binding the body of Delete User", ctx.Bind(&replacePassword))
	err := services.ChangePassowrd(&replacePassword)
	if err != nil {
		errors.ErrorControler("unable to replace the password", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Error changing the Password"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "Password changed!!"})

}

func DeleteUser(ctx *gin.Context) {
	var delUser core.DeleteUserRequest
	errors.ErrorControler("Binding the body of Delete User", ctx.Bind(&delUser))
	err := service.DeleteUser(&delUser)
	if err != nil {
		errors.ErrorControler("unable to Delete the User", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Error deleting the User"})
		return
	}
	cookie := &http.Cookie{
		Name:     "authorization",
		Value:    "",
		Path:     "/",
		Expires:  time.Now(),
		HttpOnly: true,
	}
	http.SetCookie(ctx.Writer, cookie)

	ctx.JSON(http.StatusOK, gin.H{"msg": "User deleted!!"})

}
