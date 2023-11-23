package auth

import (
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

		claims, err := ValidateToken(tokenString, configs.Config.JWTSecret)
		if err != nil {
			ctx.Set("permission", Guest)
			ctx.Next()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Set("permission", claims.Permission)

		ctx.Next()
	}
}
