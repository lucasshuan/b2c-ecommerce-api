package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lucasshuan/b2c-ecommerce-api/configs"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.Set("permission", Guest)
			ctx.Next()
			return
		}

		claims := &Claims{}
		if _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return configs.Config.JWTSecret, nil
		}); err != nil {
			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("permission", claims.Permission)

		ctx.Next()
	}
}
