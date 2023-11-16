package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/internal/logger"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		startTime := time.Now()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		clientIP := ctx.ClientIP()

		ctx.Next()

		logger.Log.Info("HTTP",
			"METHOD", reqMethod,
			"URI", reqUri,
			"LATENCY", latencyTime,
			"CLIENT_IP", clientIP,
		)
	}
}
