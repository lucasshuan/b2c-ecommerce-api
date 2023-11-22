package log

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()

		latencyTime := fmt.Sprintf("%dms", endTime.Sub(startTime).Milliseconds())
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI

		AppLogger.Info("Handled HTTP request:",
			"host", c.Request.Host,
			"method", reqMethod,
			"uri", reqUri,
			"status", c.Writer.Status(),
			"ping", latencyTime,
		)
	}
}
