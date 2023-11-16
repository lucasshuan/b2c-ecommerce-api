package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nobeluc/ecommerce-api/internal/logger"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		startTime := time.Now()

		ctx.Next()
		endTime := time.Now()
		latencyTime := fmt.Sprintf("%dms", endTime.Sub(startTime).Milliseconds())
		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI

		logger.Log.Info("Handled HTTP request:",
			"host", ctx.Request.Host,
			"method", reqMethod,
			"uri", reqUri,
			"status", ctx.Writer.Status(),
			"ping", latencyTime,
		)
	}
}
