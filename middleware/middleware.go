package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"stockApi/errors"
	service "stockApi/services"
)

func IsAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info("PROCESS ", os.Getenv("PORT"))
		cookie, err := ctx.Cookie("authorization")
		if err != nil {
			errors.ErrorMiddleware(" Cookie don't exist ", err)
			ctx.Abort()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthenticated"})
			return
		}
		claims, err := service.ValidateCookie(cookie)
		if err == nil {
			fmt.Println(claims)
			ctx.Next()
			return
		}
		errors.ErrorMiddleware("unauthenticated ", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthenticated"})
		ctx.Abort()
		return

	}
}
func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("authorization")
		if err != nil {
			errors.ErrorMiddleware(" Cookie don't exist ", err)
			ctx.Abort()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthenticated"})
			return
		}
		claims, err := service.ValidateCookie(cookie)
		if err == nil {
			log.Info(claims)
			if claims.Role == "ADMIN" {
				ctx.Next()
				return
			}
		}
		errors.ErrorMiddleware("unauthenticated ", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "unauthenticated"})
		ctx.Abort()
		return

	}
}
