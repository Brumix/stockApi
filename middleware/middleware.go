package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stockApi/errors"
	service "stockApi/services"
)

func IsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("authorization")
		if err != nil {
			errors.ErrorMiddleware(" Cookie don't exist ", err)
			ctx.Abort()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthenticated"})
			return
		}
		err = service.ValidateCookie(cookie)
		if err == nil {
			ctx.Next()
			return
		}
		errors.ErrorMiddleware("unauthenticated ", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthenticated"})
		ctx.Abort()
		return

	}
}