package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

var Log *log.Logger

type LoggerWriter struct {
	logger *log.Logger
}

func (lw *LoggerWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func Init() {
	Log = log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
	Log.SetLevel(log.DebugLevel)

	customWriter := &LoggerWriter{
		logger: Log,
	}

	gin.DefaultWriter = customWriter
}
